package yarf

import (
	"testing"
)

func TestResourceInterface(t *testing.T) {
	var r interface{}
	r = new(Resource)

	if _, ok := r.(ResourceHandler); !ok {
		t.Error("Resource type doesn't implement ResourceHandler interface")
	}
}
