package yarf

import (
	"net/http"
	"net/url"
	"encoding/json"
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

	// The HTTP status code to be writen to the response.
	ResponseStatus int

	// The aggregated response body to be writen to the response.
	ResponseContent string
}

// NewContext instantiates a new *Context object with default values and returns it.
func NewContext(r *http.Request, rw http.ResponseWriter) *Context {
	c := new(Context)
	c.Request = r
	c.Response = rw
	c.ResponseStatus = 200
	c.Params = url.Values{}

	return c
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
	rc.Context.ResponseStatus = code
}

// Param is a wrapper for rc.Context.Params.Get()
func (rc *RequestContext) Param(name string) string {
	return rc.Context.Params.Get(name)
}

// Render takes a string and aggregates it to the Context.responseContent.
// This is the default renderer, which sets a text/plain Content-Type header.
// Check other Render[Type] functions for different types.
func (rc *RequestContext) Render(content string) {
	// Set header
	rc.Context.Response.Header().Set("Content-Type", "text/plain")

	// Set content
	rc.Context.ResponseContent += content
}

// RenderJSON takes a interface{} object and writes the JSON encoded string of it. 
func (rc *RequestContext) RenderJSON(data interface{}) {
    // Set header
	rc.Context.Response.Header().Set("Content-Type", "application/json")
	
	// Set content
	encoded, err := json.Marshal(data)
	if err != nil {
		rc.Context.ResponseContent += err.Error()
	} else {
		rc.Context.ResponseContent += string(encoded)
	}
}
