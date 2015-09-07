package yarf

import (
	"net/http"
	"net/url"
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
	responseStatus int

	// The aggregated response body to be writen to the response.
	responseContent string
}

// Context constructor.
// Instantiates a new *Context object with default values and returns it.
func NewContext(r *http.Request, rw http.ResponseWriter) *Context {
	c := new(Context)
	c.Request = r
	c.Response = rw
	c.responseStatus = 200
	c.Params = url.Values{}

	return c
}

// RequestContext implements Context related methods to interact with the Context object.
// It's used to composite into Resource and Middleware to satisfy the interfaces.
type RequestContext struct {
	Context *Context
}

// Context setter
func (rc *RequestContext) SetContext(c *Context) {
	rc.Context = c
}

// Render takes a string and aggregates it to the Context.responseContent.
func (rc *RequestContext) Render(content string) {
	rc.Context.responseContent += content
}

// Status sets the HTTP status code to be returned on the response.
func (rc *RequestContext) Status(code int) {
	rc.Context.responseStatus = code
}

// Param is a wrapper for rc.Context.Params.Get()
func (rc *RequestContext) Param(name string) string {
	return rc.Context.Params.Get(name)
}
