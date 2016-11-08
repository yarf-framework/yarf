package main

import (
	"github.com/yarf-framework/yarf"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResourceHello(t *testing.T) {
	h := new(Hello)

	c := new(yarf.Context)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	c.Response = httptest.NewRecorder()

	err := h.Get(c)
	if err != nil {
		t.Error(err.Error())
	}
}
