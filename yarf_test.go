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

// Put returns default success for testing
func (r *MockResource) Put(c *Context) error {
	c.Render("PUT")
	return nil
}

// Delete returns default success for testing
func (r *MockResource) Delete(c *Context) error {
	c.Render("DELETE")
	return nil
}

// Delete returns default success for testing
func (r *MockResource) Options(c *Context) error {
	c.Render("OPTIONS")
	return nil
}

// Patch returns default success for testing
func (r *MockResource) Patch(c *Context) error {
	c.Render("PATCH")
	return nil
}

// Head returns default success for testing
func (r *MockResource) Head(c *Context) error {
	c.Render("HEAD")
	return nil
}

// Trace returns default success for testing
func (r *MockResource) Trace(c *Context) error {
	c.Render("TRACE")
	return nil
}

// Connect returns default success for testing
func (r *MockResource) Connect(c *Context) error {
	c.Render("CONNECT")
	return nil
}

// MockMiddleware is an empty middleware for mocking tests
type MockMiddleware struct {
	Middleware
}

func TestGet(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	// Add some empty middleware
	y.Insert(new(MockMiddleware))

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "MockResource" {
		t.Errorf("GET should render 'MockResource'. '%s' rendered.", res.Body.String())
	}
}

func TestPost(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	req, _ := http.NewRequest("POST", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "" {
		t.Errorf("POST should render '' (error). '%s' rendered.", res.Body.String())
	}
}

func TestPut(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	req, _ := http.NewRequest("PUT", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "PUT" {
		t.Errorf("PUT should render 'PUT'. '%s' rendered.", res.Body.String())
	}
}

func TestDelete(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	req, _ := http.NewRequest("DELETE", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "DELETE" {
		t.Errorf("DELETE should render 'DELETE'. '%s' rendered.", res.Body.String())
	}
}

func TestPatch(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	req, _ := http.NewRequest("PATCH", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "PATCH" {
		t.Errorf("PATCH should render 'PATCH'. '%s' rendered.", res.Body.String())
	}
}

func TestOptions(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	req, _ := http.NewRequest("OPTIONS", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "OPTIONS" {
		t.Errorf("OPTIONS should render 'OPTIONS'. '%s' rendered.", res.Body.String())
	}
}

func TestHead(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	req, _ := http.NewRequest("HEAD", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "HEAD" {
		t.Errorf("HEAD should render 'HEAD'. '%s' rendered.", res.Body.String())
	}
}

func TestTrace(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	req, _ := http.NewRequest("TRACE", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "TRACE" {
		t.Errorf("TRACE should render 'TRACE'. '%s' rendered.", res.Body.String())
	}
}

func TestConnect(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/", r)

	req, _ := http.NewRequest("CONNECT", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Body.String() != "CONNECT" {
		t.Errorf("CONNECT should render 'CONNECT'. '%s' rendered.", res.Body.String())
	}
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

func TestNotFound(t *testing.T) {
	y := New()

	r := new(MockResource)
	y.Add("/test", r)

	y.NotFound = func(c *Context) {
		c.Render("NOT FOUND")
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/route/not/match", nil)
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	if res.Code == 404 {
		t.Error("Non matching route should follow to Follow method")
	}

	if res.Body.String() != "NOT FOUND" {
		t.Error("NotFound method should have rendered 'NOT FOUND'")
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
	req, _ := http.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "127.0.0.1"
	req.Host = "localhost:8080"
	res := httptest.NewRecorder()
	y.ServeHTTP(res, req)

	// Check log
	if !strings.HasSuffix(strings.TrimSpace(buf.String()), "127.0.0.1 | GET | localhost:8080/test") {
		t.Errorf("Wrong log expected: %s", buf.String())
	}

	// Test error log and debug
	req, _ = http.NewRequest("POST", "http://localhost:8080/test", nil)
	req.RemoteAddr = "127.0.0.1"
	y.Debug = true
	y.ServeHTTP(res, req)

	// Check log
	if !strings.HasSuffix(strings.TrimSpace(buf.String()), "127.0.0.1 | ERROR: 500 | POST ERROR | POST ERROR") {
		t.Errorf("Wrong error log: %s", buf.String())
	}

	// Check debug render
	if res.Body.String() != "MockResourcePOST ERROR" {
		t.Errorf("Error debug should have rendered '%s', got '%s'", "MockResourcePOST ERROR", res.Body.String())
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
