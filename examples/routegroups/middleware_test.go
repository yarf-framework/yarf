package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/yarf-framework/yarf"
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

func TestExtra(t *testing.T) {
    e := new(ExtraMiddleware)
    
    c := new(yarf.Context)
    c.Request, _ = http.NewRequest("GET", "/", nil)
    c.Response = httptest.NewRecorder()
    
    err := e.PreDispatch(c)
    if err != nil {
        t.Error(err.Error())
    }
    
    err = e.PostDispatch(c)
    if err != nil {
        t.Error(err.Error())
    }
}