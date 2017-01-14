package yarf

import (
	"log"
	"net/http"
)

// Version string
const Version = "0.8.3"

// Yarf is the main entry point for the framework and it centralizes most of the functionality.
// All configuration actions are handled by this object.
type Yarf struct {
	// UseCache indicates if the route cache should be used.
	UseCache bool

	// Debug enables/disables the debug mode.
	// On debug mode, extra error information is sent to the client.
	Debug bool

	// Silent mode attempts to prevent all messages that aren't part of a resource response to get to the client.
	// Specially useful to hide error messages.
	Silent bool

	// PanicHandler can store a func() that will be defered by each request to be able to recover().
	// If you need to log, send information or do anything about a panic, this is your place.
	PanicHandler func()

	GroupRouter

	// Cached routes storage
	cache *Cache

	// Logger object will be used if present
	Logger *log.Logger

	// Follow defines a standard http.Handler implementation to follow if no route matches.
	Follow http.Handler

	// NotFound defines a function interface to execute when a NotFound (404) error is thrown.
	NotFound func(c *Context)
}

// New creates a new yarf and returns a pointer to it.
// Performs needed initializations
func New() *Yarf {
	y := new(Yarf)

	// Init cache
	y.UseCache = true
	y.cache = NewCache()
	y.GroupRouter = RouteGroup("")

	// Return object
	return y
}

// ServeHTTP Implements http.Handler interface into yarf.
// Initializes a Context object and handles middleware and route actions.
// If an error is returned by any of the actions, the flow is stopped and a response is sent.
// If no route matches, tries to forward the request to the Yarf.Follow (http.Handler type) property if set.
// Otherwise it returns a 404 response.
func (y *Yarf) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if y.PanicHandler != nil {
		defer y.PanicHandler()
	}

	// Set initial context data.
	// The Context pointer will be affected by the middleware and resources.
	c := NewContext(req, res)

	// Cached routes
	if y.UseCache {
		if cache, ok := y.cache.Get(req.URL.Path); ok {
			// Set context params
			c.Params = cache.params
			c.groupDispatch = cache.route

			// Dispatch and stop
			err := y.Dispatch(c)
			y.finish(c, err)
			return
		}
	}

	// Route match
	if y.Match(req.URL.Path, c) {
		if y.UseCache {
			y.cache.Set(req.URL.Path, RouteCache{c.groupDispatch, c.Params})
		}
		err := y.Dispatch(c)
		y.finish(c, err)
		return
	}

	// Follow only when route doesn't match.
	// Returned 404 errors won't follow.
	if y.Follow != nil {
		// Log follow
		y.finish(c, nil)

		// Follow
		y.Follow.ServeHTTP(c.Response, c.Request)

		// End here
		return
	}

	// Return 404
	y.finish(c, ErrorNotFound())
}

// Finish handles the end of the execution.
// It checks for errors and follow actions to execute.
// It also handles the custom 404 error handler.
func (y *Yarf) finish(c *Context, err error) {
	// If a logger is present, lets log everything.
	if y.Logger != nil {
	    req := "http"
	    if c.Request.TLS != nil {
	        req += "s"
	    }
	    req += "://" + c.Request.Host + c.Request.URL.String()
	    
		y.Logger.Printf(
			"%s | %s | %s ",
			c.GetClientIP(),
			c.Request.Method,
			req,
		)
	}

	// Return if no error or silent mode
	if err == nil || y.Silent {
		return
	}

	// Check error type
	yerr, ok := err.(YError)
	if !ok {
		// Create default 500 error
		yerr = &CustomError{
			HTTPCode:  500,
			ErrorCode: 0,
			ErrorMsg:  err.Error(),
			ErrorBody: err.Error(),
		}
	}

	// Custom 404
	if yerr.Code() == 404 && y.NotFound != nil {
		y.NotFound(c)
		return
	}

	// Log errors
	if y.Logger != nil {
		y.Logger.Printf(
			"%s | ERROR: %d | %s | %s",
			c.GetClientIP(),
			yerr.Code(),
			yerr.Body(),
			yerr.Msg(),
		)
	}

	// Write error data to response.
	c.Response.WriteHeader(yerr.Code())

	// Render errors if debug enabled
	if y.Debug {
		c.Render(yerr.Body())
	}
}

// Start initiates a new http yarf server and start listening.
// It's a shortcut for http.ListenAndServe(address, y)
func (y *Yarf) Start(address string) {
	http.ListenAndServe(address, y)
}

// StartTLS initiates a new http yarf server and starts listening to HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(address, cert, key, yarf)
func (y *Yarf) StartTLS(address, cert, key string) {
	http.ListenAndServeTLS(address, cert, key, y)
}
