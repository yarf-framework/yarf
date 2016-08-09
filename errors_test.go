package yarf

import (
	"testing"
)

func TestErrors(t *testing.T) {
	e := ErrorUnexpected()
	if e == nil {
		t.Error("ErrorUnexpected() should return an object. Nil value returned.")
	}
	if e.Error() != "Unexpected error" {
		t.Errorf("e.Error() should return 'Unexpected error'. '%s' returned.", e.Error())
	}
	if e.Body() != "" {
		t.Errorf("e.Body() should return ''. '%s' returned.", e.Body())
	}
	if e.Code() != 500 {
		t.Errorf("e.Code() should return '500'. '%d' returned.", e.Code())
	}
	if e.ID() != 0 {
		t.Errorf("e.ID() should return '0'. '%d' returned.", e.ID())
	}

	eni := ErrorMethodNotImplemented()
	if eni == nil {
		t.Error("ErrorMethodNotImplemented() should return an object. Nil value returned.")
	}

	enf := ErrorNotFound()
	if enf == nil {
		t.Error("ErrorNotFound() should return an object. Nil value returned.")
	}
}
