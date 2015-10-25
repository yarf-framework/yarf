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

func TestStatus(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.Status(201)

	if c.Response.(*httptest.ResponseRecorder).Code != 201 {
		t.Errorf("Status %d set to Status() method, %d found", 201, c.Response.(*httptest.ResponseRecorder).Code)
	}
}

func TestParam(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.Params.Set("name", "Joe")

	if c.Param("name") != "Joe" {
		t.Errorf("Param 'name' set to '%s', '%s' retrieved.", "Joe", c.Param("name"))
	}
}

func TestGetClientIP(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)

	if c.GetClientIP() != req.RemoteAddr {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", req.RemoteAddr, c.GetClientIP())
	}
}

func TestRender(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.Render("TEST")

	if c.Response.(*httptest.ResponseRecorder).Body.String() != "TEST" {
		t.Errorf("'%s' sent to Render() method, '%s' found on Response object", "TEST", c.Response.(*httptest.ResponseRecorder).Body.String())
	}
}

func TestRenderJSON(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.RenderJSON("TEST")

	if c.Response.(*httptest.ResponseRecorder).Body.String() != "\"TEST\"" {
		t.Errorf("'%s' sent to RenderJSON() method, '%s' found on Response object", "TEST", c.Response.(*httptest.ResponseRecorder).Body.String())
	}
}

func TestRenderJSONIndent(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.RenderJSONIndent("TEST")

	if c.Response.(*httptest.ResponseRecorder).Body.String() != "\"TEST\"" {
		t.Errorf("'%s' sent to RenderJSONIndent() method, '%s' found on Response object", "TEST", c.Response.(*httptest.ResponseRecorder).Body.String())
	}
}

func TestRenderXML(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.RenderXML("TEST")

	if c.Response.(*httptest.ResponseRecorder).Body.String() != "<string>TEST</string>" {
		t.Errorf("'%s' sent to RenderXML() method, '%s' found on Response object", "TEST", c.Response.(*httptest.ResponseRecorder).Body.String())
	}
}

func TestRenderXMLIndent(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.RenderXMLIndent("TEST")

	if c.Response.(*httptest.ResponseRecorder).Body.String() != "<string>TEST</string>" {
		t.Errorf("'%s' sent to RenderXMLIndent() method, '%s' found on Response object", "TEST", c.Response.(*httptest.ResponseRecorder).Body.String())
	}
}
