package yarf

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockResource struct {
	Resource
}

// Get returns default success for testing
func (r *MockResource) Get(c *Context) error {
	c.Render("MockResource")
	return nil
}

// Post returns default error for testing
func (r *MockResource) Post(c *Context) error {
	return errors.New("POST ERROR")
}

// MockMiddleware is an empty middleware for mocking tests
type MockMiddleware struct {
	Middleware
}

func TestYarfCache(t *testing.T) {
	y := New()

	if len(y.cache.storage) > 0 {
		t.Error("yarf.cache.storage should be empty after initialization")
	}

	r := new(MockResource)
	y.Add("/test", r)

	req, _ := http.NewRequest("GET", "http://localhost:8080/route/not/match", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if len(y.cache.storage) > 0 {
		t.Error("yarf.cache.storage should be empty after non-matching request")
	}

	req, _ = http.NewRequest("GET", "http://localhost:8080/test", nil)
	y.ServeHTTP(res, req)

	if len(y.cache.storage) != 1 {
		t.Error("yarf.cache.storage should have 1 item after matching request")
	}

	for i := 0; i < 100; i++ {
		y.ServeHTTP(res, req)
	}

	if len(y.cache.storage) != 1 {
		t.Error("yarf.cache.storage should have 1 item after multiple matching requests to a single route")
	}
}

func TestYarfUseCacheFalse(t *testing.T) {
	r := new(MockResource)
	y := New()
	y.UseCache = false
	y.Add("/test", r)

	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if len(y.cache.storage) > 0 {
		t.Error("yarf.cache.storage should be empty after matching request with yarf.UseCache = false")
	}
}

func TestRace(t *testing.T) {
	g := RouteGroup("/test")
	g.Add("/one/:param", &MockResource{})
	g.Add("/two/:param", &MockResource{})

	y := New()
	y.AddGroup(g)

	one, _ := http.NewRequest("GET", "http://localhost:8080/test/one/1", nil)
	two, _ := http.NewRequest("GET", "http://localhost:8080/test/two/2", nil)

	for i := 0; i < 1000; i++ {
		res1 := httptest.NewRecorder()
		res2 := httptest.NewRecorder()

		go y.ServeHTTP(res1, one)
		go y.ServeHTTP(res2, two)
	}
}

func TestNotFoundResponse(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/test", r)

	req, _ := http.NewRequest("GET", "http://localhost:8080/route/not/match", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Code != 404 {
		t.Error("Non matching route should return 404 response")
	}
}

func TestCustomErrorResponse(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/test", r)

	req, _ := http.NewRequest("POST", "http://localhost:8080/test", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Code != 500 {
		t.Error("Cutom error should return HTTP 500 code")
	}
}

func TestLogger(t *testing.T) {
	y := New()

	// Init logger
	var buf bytes.Buffer
	y.Logger = log.New(&buf, "", log.Lshortfile)

	// Mock route and resource
	r := new(MockResource)
	y.Add("/test", r)

	// Request
	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	req.RemoteAddr = "127.0.0.1"
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	// Check log
	if !strings.HasSuffix(strings.TrimSpace(buf.String()), "127.0.0.1 | GET | http://localhost:8080/test") {
		t.Errorf("Wrong log expected: %s", buf.String())
	}
}

func TestFollow(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/test", r)

	y.Follow = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FOLLOW"))
	})

	req, _ := http.NewRequest("GET", "http://localhost:8080/route/not/match", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Code == 404 {
		t.Error("Non matching route should follow to Follow method")
	}

	if res.Body.String() != "FOLLOW" {
		t.Error("Follow method should have rendered 'FOLLOW'")
	}
}
