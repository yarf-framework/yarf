package yarf

import (
	"net/http"
	"net/url"
)

// Framework version string
const Version = "0.3"

// routeCache stores previously matched and parsed routes
type routeCache struct {
	route  Router
	params url.Values
}

// yarf is the main entry point for the framework and it centralizes most of the functionality.
// All YARF configuration actions are handled by the yarf.
type yarf struct {
	routes []Router // yarf routes

	useCache bool // Indicates if the route cache should be used

	cache map[string]routeCache // Cached routes storage

	middleware []MiddlewareHandler // Middleware resources
}

// New creates a new yarf and returns a pointer to it.
// Performs needed initializations
func New() *yarf {
	y := new(yarf)

	// Init cache
	y.useCache = true
	y.cache = make(map[string]routeCache)

	// Return object
	return y
}

// UseCache sets the useCache flag to indicate if route caching should be enabled.
// The route cache improves the performance, but it also consumes (not that much) memory space.
// If you're running out of RAM memory and/or your app has too many possible routes that may not fit,
// you should disable the route cache.
func (y *yarf) UseCache(b bool) {
	y.useCache = b
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
	// Set initial context data.
	// The Context pointer will be affected by the middleware and resources.
	c := NewContext(req, res)

	// Cached routes
	if y.useCache {
		if cache, ok := y.cache[req.URL.Path]; ok {
			// Set context params
			c.Params = cache.params

			// Dispatch and stop
			y.dispatch(cache.route, c)
			return
		}
	}

	// Route match
	for _, r := range y.routes {
		if r.Match(req.URL.Path, c) {
			// Store cache
			y.cache[req.URL.Path] = routeCache{r, c.Params}

			// Dispatch and stop
			y.dispatch(r, c)
			return
		}
	}

	// Return 404
	y.Response(c, ErrorNotFound())
}

// Dispatch performs the middleware and route handler actions for a given route and context.
func (y *yarf) dispatch(r Router, c *Context) {
	// Init error status
	var err error

	// Pre-Dispatch Middleware
	for _, m := range y.middleware {
		// Add context to middleware
		m.SetContext(c)

		// Dispatch
		err = m.PreDispatch()
		if err != nil {
			// Stop on error
			break
		}
	}

	// Route dispatch
	if err == nil {
		err = r.Dispatch(c)

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
	y.Response(c, err)
}

// Response writes the corresponding response to the HTTP response writer.
// It will handle the error status and the response body to be sent.
func (y *yarf) Response(c *Context, err error) {
	// Error handling
	if err != nil {
		if _, ok := err.(YError); !ok {
			err = ErrorUnexpected()
		}

		// Replace context content with error data.
		c.ResponseStatus = err.(YError).Code()
		c.ResponseContent = err.(YError).Body()
	}

	// Write HTTP status
	c.Response.WriteHeader(c.ResponseStatus)

	// Write body
	c.Response.Write([]byte(c.ResponseContent))
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
