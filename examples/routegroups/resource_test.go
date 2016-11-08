package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/yarf-framework/yarf"
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

func TestResourceHelloV2(t *testing.T) {
    h := new(HelloV2)
    
    c := new(yarf.Context)
    c.Request, _ = http.NewRequest("GET", "/", nil)
    c.Response = httptest.NewRecorder()
    
    err := h.Get(c)
    if err != nil {
        t.Error(err.Error())
    }
}