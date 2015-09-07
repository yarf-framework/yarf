package yarf

import (
	"net/http"
)

// Framework version string
const Version = "0.1b"

// yarf is the main entry point for the framework and it centralizes most of the functionality.
// All YARF configuration actions are handled by the yarf.
type yarf struct {
	routes []Router // yarf routes

	middleware []MiddlewareHandler // Middleware resources
}

// Creates a new yarf and returns a pointer to it.
// Performs needed initializations
func Yarf() *yarf {
	y := new(yarf)

	// No initialization routines yet...
	return y
}

// Add inserts a new resource with it's associated route.
func (y *yarf) Add(url string, r ResourceHandler) {
	y.routes = append(y.routes, Route(url, r))
}

// AddGroup inserts a route group into the routes list.
func (y *yarf) AddGroup(g *routeGroup) {
	y.routes = append(y.routes, g)
}

// Insert adds a MiddlewareHandler into the middleware list
func (y *yarf) Insert(m MiddlewareHandler) {
	y.middleware = append(y.middleware, m)
}

// ServeHTTP Implements http.Handler interface into yarf.
// Initializes a Context object and handles middleware and route actions.
// If an error is returned by any of the actions, the flow is stopped and a response is sent.
func (y *yarf) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Init error status
	var err error

	// Set initial context data.
	// The Context pointer will be affected by the middleware and resources.
	ctx := NewContext(req, res)

	// Route match
	for _, r := range y.routes {
		if r.Match(req.URL.Path, ctx) {
			// Pre-Dispatch Middleware
			for _, m := range y.middleware {
				// Add context to middleware
				m.SetContext(ctx)

				// Dispatch
				err = m.PreDispatch()
				if err != nil {
					// Stop on error
					break
				}
			}

			// Route dispatch
			if err == nil {
				err = r.Dispatch(ctx)

				if err == nil {
					// Post-Dispatch Middleware
					for _, m := range y.middleware {
						err = m.PostDispatch()
						if err != nil {
							// Stop on error
							break
						}
					}
				}
			}

			// Return result
			y.Response(ctx, err)
			return
		}
	}

	// Return 404
	y.Response(ctx, ErrorNotFound())
}

// Response writes the corresponding response to the HTTP response writer.
// It will handle the error status and the response body to be sent.
func (y *yarf) Response(c *Context, err error) {
	// Error handling
	if err != nil {
		if _, ok := err.(YarfError); !ok {
		    err = ErrorUnexpected()
		}
		
		// Replace context content with error data.
		c.responseStatus = err.(YarfError).Code()
		c.responseContent = err.(YarfError).Body()
	}

	// Write HTTP status
	c.Response.WriteHeader(c.responseStatus)

	// Write body
	c.Response.Write([]byte(c.responseContent))
}

// Start initiates a new http yarf and start listening.
// It's a wrapper for http.ListenAndServe(addr, router)
func (y *yarf) Start(address string) {
	// Run
	http.ListenAndServe(address, y)
}

// StartTLS initiats a new http yarf and starts listening and HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(address, cert, key, yarf)
func (y *yarf) StartTLS(address, cert, key string) {
	// Run
	http.ListenAndServeTLS(address, cert, key, y)
}
