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

func TestResourceInterfaceMethods(t *testing.T) {
	r := new(Resource)

	if r.Get() == nil {
		t.Error("Get() should return a MethodNotImplementedError type by default.")
	}
	if r.Post() == nil {
		t.Error("Post() should return a MethodNotImplementedError type by default.")
	}
	if r.Put() == nil {
		t.Error("Put() should return a MethodNotImplementedError type by default.")
	}
	if r.Delete() == nil {
		t.Error("Delete() should return a MethodNotImplementedError type by default.")
	}
	if r.Patch() == nil {
		t.Error("Patch() should return a MethodNotImplementedError type by default.")
	}
	if r.Head() == nil {
		t.Error("Head() should return a MethodNotImplementedError type by default.")
	}
	if r.Options() == nil {
		t.Error("Options() should return a MethodNotImplementedError type by default.")
	}
	if r.Connect() == nil {
		t.Error("Connect() should return a MethodNotImplementedError type by default.")
	}
	if r.Trace() == nil {
		t.Error("Trace() should return a MethodNotImplementedError type by default.")
	}

}
