package yarf

import (
	"net/http"
)

// Framework version string
const Version = "0.1b"

// Yarf is the main entry point for the framework and it centralizes most of the functionality.
// All YARF configuration actions are handled by the Yarf.
type Yarf struct {
	routes []*Route // Yarf routes

	preDispatch []MiddlewareResource // Run-before middleware resources

	postDispatch []MiddlewareResource // Run-after middleware resources
}

// Creates a new Yarf and returns a pointer to it.
// Performs needed initializations
func New() *Yarf {
	y := new(Yarf)

	// No initialization routines yet...
	return y
}

// Add inserts a new resource with it's associated route.
func (y *Yarf) Add(url string, r RestResource) {
	y.routes = append(y.routes, &Route{handler: r, path: url})
}

// AddBefore inserts a MiddlewareResource into the pre-dispatch middleware list
func (y *Yarf) AddBefore(m MiddlewareResource) {
	y.preDispatch = append(y.preDispatch, m)
}

// AddAfter inserts a MiddlewareResource into the post-dispatch middleware list
func (y *Yarf) AddAfter(m MiddlewareResource) {
	y.postDispatch = append(y.postDispatch, m)
}

// ServeHTTP Implements http.Handler interface into Yarf.
// Initializes a Context object and handles middleware and route actions.
// If an error is returned by any of the actions, the flow is stopped and a response is sent.
func (y *Yarf) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error

	// Set initial context data.
	// The Context pointer will be affected by the middleware and resources.
	ctx := NewContext(req, res)

	// Pre-Dispatch Middleware
	for _, m := range y.preDispatch {
		err = m.PreDispatch(ctx)
		if err != nil {
			// Return response on error
			y.Response(ctx, err)
			return
		}
	}

	// Route dispatch
	for _, r := range y.routes {
		if r.Match(req.URL.Path, ctx) {
			err = r.Dispatch(ctx)
			if err != nil {
				// Return response on error
				y.Response(ctx, err)
				return
			}
		}
	}

	// Post-Dispatch Middleware
	for _, m := range y.postDispatch {
		err = m.PostDispatch(ctx)
		if err != nil {
			// Return response on error
			y.Response(ctx, err)
			return
		}
	}

	// Return response
	y.Response(ctx, nil)
}

// Response writes the corresponding response to the HTTP response writer.
// It will handle the error status and the response body to be sent.
func (y *Yarf) Response(c *Context, err error) {
	// Error handling
	if err != nil {

	}

	// Write HTTP status
	c.Response.WriteHeader(c.responseStatus)

	// Write body
	c.Response.Write([]byte(c.responseContent))
}

// Start initiates a new http Yarf and start listening.
// It's a wrapper for http.ListenAndServe(addr, router)
func (y *Yarf) Start(address string) {
	// Run
	http.ListenAndServe(address, y)
}

// StartTLS initiats a new http Yarf and starts listening and HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(address, cert, key, Yarf)
func (y *Yarf) StartTLS(address, cert, key string) {
	// Run
	http.ListenAndServeTLS(address, cert, key, y)
}
