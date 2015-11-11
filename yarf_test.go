package yarf

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockResource struct {
	Resource
}

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
