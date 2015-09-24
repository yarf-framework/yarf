package yarf

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// Context is the data/status storage of every YARF request.
// Every request will instantiate a new Context object and fill in with all the request data.
// Each request Context will be shared along the entire request life to ensure accesibility of its data at all levels.
type Context struct {
	// The *http.Request object as received by the HandleFunc.
	Request *http.Request

	// The http.ResponseWriter object as received by the HandleFunc.
	Response http.ResponseWriter

	// Parameters received through URL route
	Params url.Values
}

// NewContext instantiates a new *Context object with default values and returns it.
func NewContext(r *http.Request, rw http.ResponseWriter) *Context {
	return &Context{
		Request:  r,
		Response: rw,
		Params:   url.Values{},
	}
}

// RequestContext implements Context related methods to interact with the Context object.
// It's used to composite into Resource and Middleware to satisfy the interfaces.
type RequestContext struct {
	Context *Context
}

// SetContext setter
func (rc *RequestContext) SetContext(c *Context) {
	rc.Context = c
}

// Status sets the HTTP status code to be returned on the response.
func (rc *RequestContext) Status(code int) {
	rc.Context.Response.WriteHeader(code)
}

// Param is a wrapper for rc.Context.Params.Get()
func (rc *RequestContext) Param(name string) string {
	return rc.Context.Params.Get(name)
}

// GetClientIP retrieves the client IP address from the request information.
// It detects common proxy headers to return the actual client's IP and not the proxy's.
func (rc *RequestContext) GetClientIP() (ip string) {
	var pIps string
	var pIpList []string

	if pIps = rc.Context.Request.Header.Get("X-Real-Ip"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else if pIps = rc.Context.Request.Header.Get("Real-Ip"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else if pIps = rc.Context.Request.Header.Get("X-Forwarded-For"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else if pIps = rc.Context.Request.Header.Get("X-Forwarded"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else if pIps = rc.Context.Request.Header.Get("Forwarded-For"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else if pIps = rc.Context.Request.Header.Get("Forwarded"); pIps != "" {
		pIpList = strings.Split(pIps, ",")
		ip = strings.TrimSpace(pIpList[0])

	} else {
		ip = rc.Context.Request.RemoteAddr
	}

	return strings.Split(ip, ":")[0]
}

// Render writes a string to the http.ResponseWriter.
// This is the default renderer that just sends the string to the client.
// Check other Render[Type] functions for different types.
func (rc *RequestContext) Render(content string) {
	// Write response
	rc.Context.Response.Write([]byte(content))
}

// RenderJSON takes a interface{} object and writes the JSON encoded string of it.
func (rc *RequestContext) RenderJSON(data interface{}) {
	// Set content
	encoded, err := json.Marshal(data)
	if err != nil {
		rc.Context.Response.Write([]byte(err.Error()))
	} else {
		rc.Context.Response.Write(encoded)
	}
}
