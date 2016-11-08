package middleware

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
    
    err := h.PreDispatch(c)
    if err != nil {
        t.Error(err.Error())
    }
    
    err = h.PostDispatch(c)
    if err != nil {
        t.Error(err.Error())
    }
}
