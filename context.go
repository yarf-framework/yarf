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

	return c
}
