package yarf

import (
    "testing"
)

func TestErrors(t *testing.T) {
	var e YError
	
	e = ErrorUnexpected()
	if e == nil {
        t.Error("ErrorUnexpected() should return an object. Nil value returned.")	    
	}
	
	e = ErrorMethodNotImplemented()
	if e == nil {
        t.Error("ErrorMethodNotImplemented() should return an object. Nil value returned.")	    
	}
	
	e = ErrorNotFound()
	if e == nil {
        t.Error("ErrorNotFound() should return an object. Nil value returned.")	    
	}
}