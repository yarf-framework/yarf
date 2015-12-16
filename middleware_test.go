package yarf

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareInterface(t *testing.T) {
	var m interface{}
	m = new(Middleware)

	if _, ok := m.(MiddlewareHandler); !ok {
		t.Error("Middleware type doesn't implement MiddlewareHandler interface")
	}
}

func TestMiddlewareDefaultResponse(t *testing.T) {
	m := new(Middleware)

	// Create a dummy request.
	request, _ := http.NewRequest(
		"GET",
		"http://127.0.0.1:8080/",
		nil,
	)
	response := httptest.NewRecorder()

	c := NewContext(request, response)

	if m.PreDispatch(c) != nil {
		t.Error("Default PreDispatch() implementation should return nil")
	}
	if m.PostDispatch(c) != nil {
		t.Error("Default PostDispatch() implementation should return nil")
	}
}
