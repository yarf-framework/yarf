package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/yarf-framework/yarf"
)

func TestHello(t *testing.T) {
    h := new(Hello)
    
    c := new(yarf.Context)
    c.Request, _ = http.NewRequest("GET", "/", nil)
    c.Response = httptest.NewRecorder()
    
    err := h.Get(c)
    if err != nil {
        t.Error(err.Error())
    }
}

func TestMain(t *testing.T) {
    // Run main method in a goroutine to make sure it runs. 
    // Then let it die and just capture panics.
    go func() {
        defer func() {
            if r := recover(); r != nil {
                t.Errorf("PANIC: %s", r)
            }
        }()
        
        main()
    }()   
}
