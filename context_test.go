package yarf

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func createRequestResponse() (request *http.Request, response *httptest.ResponseRecorder) {
	// Create a dummy request.
	request, _ = http.NewRequest(
		"GET",
		"http://127.0.0.1:8080/",
		nil,
	)

	request.RemoteAddr = "200.201.202.203"
	request.Header.Set("User-Agent", "yarf/1.0")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response = httptest.NewRecorder()

	return
}

func TestNewContext(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	if c.Request != req {
		t.Error("Request object provided to NewContext() wasn't set correctly on Context object")
	}
	if c.Response != res {
		t.Error("Response object provided to NewContext() wasn't set correctly on Context object")
	}
}

func TestSetContext(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	rc := new(RequestContext)
	rc.SetContext(c)

	if rc.Context != c {
		t.Error("Context object provided to SetContext() wasn't set correctly on RequestContext object")
	}
}

func TestStatus(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	rc := new(RequestContext)
	rc.SetContext(c)

	rc.Status(201)

	if rc.Context.Response.(*httptest.ResponseRecorder).Code != 201 {
		t.Errorf("Status %d set to Status() method, %d found", 201, rc.Context.Response.(*httptest.ResponseRecorder).Code)
	}
}

func TestParam(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.Params.Set("name", "Joe")

	rc := new(RequestContext)
	rc.SetContext(c)

	if rc.Param("name") != "Joe" {
		t.Errorf("Param 'name' set to '%s', '%s' retrieved.", "Joe", rc.Param("name"))
	}
}

func TestGetClientIP(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	rc := new(RequestContext)
	rc.SetContext(c)

	if rc.GetClientIP() != req.RemoteAddr {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", req.RemoteAddr, rc.GetClientIP())
	}
}

func TestRender(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	rc := new(RequestContext)
	rc.SetContext(c)

	rc.Render("TEST")

	if rc.Context.Response.(*httptest.ResponseRecorder).Body.String() != "TEST" {
		t.Errorf("'%s' sent to Render() method, '%s' found on Response object", "TEST", rc.Context.Response.(*httptest.ResponseRecorder).Body.String())
	}
}

func TestRenderJSON(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	rc := new(RequestContext)
	rc.SetContext(c)

	rc.RenderJSON("TEST")

	if rc.Context.Response.(*httptest.ResponseRecorder).Body.String() != "\"TEST\"" {
		t.Errorf("'%s' sent to RenderJSON() method, '%s' found on Response object", "TEST", rc.Context.Response.(*httptest.ResponseRecorder).Body.String())
	}
}

func TestRenderJSONIndent(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	rc := new(RequestContext)
	rc.SetContext(c)

	rc.RenderJSONIndent("TEST")

	if rc.Context.Response.(*httptest.ResponseRecorder).Body.String() != "\"TEST\"" {
		t.Errorf("'%s' sent to RenderJSONIndent() method, '%s' found on Response object", "TEST", rc.Context.Response.(*httptest.ResponseRecorder).Body.String())
	}
}

func TestRenderXML(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	rc := new(RequestContext)
	rc.SetContext(c)

	rc.RenderXML("TEST")

	if rc.Context.Response.(*httptest.ResponseRecorder).Body.String() != "<string>TEST</string>" {
		t.Errorf("'%s' sent to RenderXML() method, '%s' found on Response object", "TEST", rc.Context.Response.(*httptest.ResponseRecorder).Body.String())
	}
}

func TestRenderXMLIndent(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	rc := new(RequestContext)
	rc.SetContext(c)

	rc.RenderXMLIndent("TEST")

	if rc.Context.Response.(*httptest.ResponseRecorder).Body.String() != "<string>TEST</string>" {
		t.Errorf("'%s' sent to RenderXMLIndent() method, '%s' found on Response object", "TEST", rc.Context.Response.(*httptest.ResponseRecorder).Body.String())
	}
}
