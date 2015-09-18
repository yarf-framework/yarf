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

	// Error handling
	if err != nil {
		if _, ok := err.(YError); !ok {
			err = ErrorUnexpected()
		}

		// Replace context content with error data.
		c.Response.WriteHeader(err.(YError).Code())
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
