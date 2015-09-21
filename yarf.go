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
// All configuration actions are handled by this object.
type yarf struct {
	// UseCache indicates if the route cache should be used.
	UseCache bool

	// Debug enables/disables the debug mode.
	// On debug mode, extra error information is sent to the client.
	Debug bool

	// Silent mode attempts to prevent all messages that aren't part of a resource response to get to the client.
	// Specially useful to hide error messages.
	Silent bool

	routes []Router // yarf routes

	cache map[string]routeCache // Cached routes storage

	middleware []MiddlewareHandler // Middleware resources
}

// New creates a new yarf and returns a pointer to it.
// Performs needed initializations
func New() *yarf {
	y := new(yarf)

	// Init cache
	y.UseCache = true
	y.cache = make(map[string]routeCache)

	// Return object
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
	// Set initial context data.
	// The Context pointer will be affected by the middleware and resources.
	c := NewContext(req, res)

	// Cached routes
	if y.UseCache {
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
	c.Response.WriteHeader(404)
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

	// Call error handler
	y.errorHandler(err, c)
}

// errorHandler deals with request errors.
func (y *yarf) errorHandler(err error, c *Context) {
	// Return if no error or silent mode
	if err == nil || y.Silent {
		return
	}

	// Check error type
	if _, ok := err.(YError); !ok {
		// Create custom 500 error
		err = &CustomError{
			httpCode:  500,
			errorCode: 0,
			errorMsg:  err.Error(),
			errorBody: err.Error(),
		}
	}

	// Write error data to response.
	c.Response.WriteHeader(err.(YError).Code())

	if y.Debug {
		c.Response.Write([]byte(err.(YError).Body()))
	}
}

// Start initiates a new http yarf server and start listening.
// It's a shortcut for http.ListenAndServe(address, y)
func (y *yarf) Start(address string) {
	http.ListenAndServe(address, y)
}

// StartTLS initiates a new http yarf server and starts listening to HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(address, cert, key, yarf)
func (y *yarf) StartTLS(address, cert, key string) {
	http.ListenAndServeTLS(address, cert, key, y)
}
