package yarf

import (
	"net/http"
	"net/http/httptest"
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
	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	res := httptest.NewRecorder()
	c := NewContext(req, res)

	if r.Get(c) == nil {
		t.Error("Get() should return a MethodNotImplementedError type by default.")
	}
	if r.Post(c) == nil {
		t.Error("Post() should return a MethodNotImplementedError type by default.")
	}
	if r.Put(c) == nil {
		t.Error("Put() should return a MethodNotImplementedError type by default.")
	}
	if r.Delete(c) == nil {
		t.Error("Delete() should return a MethodNotImplementedError type by default.")
	}
	if r.Patch(c) == nil {
		t.Error("Patch() should return a MethodNotImplementedError type by default.")
	}
	if r.Head(c) == nil {
		t.Error("Head() should return a MethodNotImplementedError type by default.")
	}
	if r.Options(c) == nil {
		t.Error("Options() should return a MethodNotImplementedError type by default.")
	}
	if r.Connect(c) == nil {
		t.Error("Connect() should return a MethodNotImplementedError type by default.")
	}
	if r.Trace(c) == nil {
		t.Error("Trace() should return a MethodNotImplementedError type by default.")
	}

}
