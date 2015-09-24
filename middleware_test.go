package yarf

import (
	"testing"
)

func TestMiddlewareInterface(t *testing.T) {
	var m interface{}
	m = new(Middleware)

	if _, ok := m.(MiddlewareHandler); !ok {
		t.Error("Middleware type doesn't implement MiddlewareHandler interface")
	}
}
