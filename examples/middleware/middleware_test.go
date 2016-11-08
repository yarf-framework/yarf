package main

import (
	"github.com/yarf-framework/yarf"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	h := new(HelloMiddleware)

	c := new(yarf.Context)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	c.Response = httptest.NewRecorder()

	err := h.PreDispatch(c)
	if err != nil {
		t.Error(err.Error())
	}

	err = h.PostDispatch(c)
	if err != nil {
		t.Error(err.Error())
	}
}
