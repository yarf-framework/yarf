package yarf

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

/*
func TestParam(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	c.Params.Set("name", "Joe")

	if c.Param("name") != "Joe" {
		t.Errorf("Param 'name' set to '%s', '%s' retrieved.", "Joe", c.Param("name"))
	}
}
*/

func TestGetClientIP(t *testing.T) {
	req, res := createRequestResponse()

	c := NewContext(req, res)
	if c.GetClientIP() != req.RemoteAddr {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", req.RemoteAddr, c.GetClientIP())
	}

	// Now check proxy headers
	req.Header.Set("Forwarded", "192.168.1.6")
	c = NewContext(req, res)
	if c.GetClientIP() != "192.168.1.6" {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", "192.168.1.6", c.GetClientIP())
	}

	req.Header.Set("Forwarded-For", "192.168.1.5")
	c = NewContext(req, res)
	if c.GetClientIP() != "192.168.1.5" {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", "192.168.1.5", c.GetClientIP())
	}

	req.Header.Set("X-Forwarded", "192.168.1.4")
	c = NewContext(req, res)
	if c.GetClientIP() != "192.168.1.4" {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", "192.168.1.4", c.GetClientIP())
	}

	req.Header.Set("X-Forwarded-For", "192.168.1.3")
	c = NewContext(req, res)
	if c.GetClientIP() != "192.168.1.3" {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", "192.168.1.3", c.GetClientIP())
	}

	req.Header.Set("Real-Ip", "192.168.1.2")
	c = NewContext(req, res)
	if c.GetClientIP() != "192.168.1.2" {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", "192.168.1.2", c.GetClientIP())
	}

	req.Header.Set("X-Real-Ip", "192.168.1.1")
	c = NewContext(req, res)
	if c.GetClientIP() != "192.168.1.1" {
		t.Errorf("IP %s set to request, %s retrieved by GetClientIP()", "192.168.1.1", c.GetClientIP())
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

/*
func TestParams(t *testing.T) {
	p := Params{}
	key := "CHECK"
	value := "THIS"
	check := ""

	// Get empty
	check = p.Get(key)
	if check != "" {
		t.Errorf("Expected to get an empty string for unassigned %s key. But got %s", key, check)
	}

	// Set key/value
	p.Set(key, value)
	check = p.Get(key)
	if check != value {
		t.Errorf("Expected to get %s for %s key. But got %s", value, key, check)
	}

	// Delete key
	p.Del(key)
	check = p.Get(key)
	if check != "" {
		t.Errorf("Expected to get an empty string for deleted %s key. But got %s", key, check)
	}
}
*/

func TestFormValue(t *testing.T) {
	req, _ := http.NewRequest(
		"POST",
		"http://127.0.0.1:8080/",
		strings.NewReader("param=value"), // POST param
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res := httptest.NewRecorder()

	c := NewContext(req, res)

	if c.FormValue("param") != "value" {
		t.Error("FormValue param isn't retreiving correct value")
	}
}
