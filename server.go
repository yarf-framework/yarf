package yarf

import (
	"net/http"
	"time"
)

// The Server struct wraps the net/http Server object.
// It's the main point for the framework and it centralizes most of the request functionality.
// All YARF configuration actions are handled by the server.
type Server struct {
	httpServer *http.Server // http.Server container

	routes []*Route // Server routes

	preDispatch []MiddlewareResource // Run-before middleware resources

	postDispatch []MiddlewareResource // Run-after middleware resources
}

// Creates a new server and returns a pointer to it.
// It receives a net/http Server to be used. If nil, a default server will be initiated.
func New(srv *http.Server) *Server {
	s := new(Server)

	// If http.Server is nil, lets start with a default server configuration
	if srv == nil {
		srv = &http.Server{
			Addr:         ":8080",
			Handler:      nil,
			ReadTimeout:  time.Duration(10) * time.Second,
			WriteTimeout: time.Duration(10) * time.Second,
		}
	}

	s.httpServer = srv

	return s
}

// Add inserts a new resource with it's associated route.
func (s *Server) Add(url string, r RestResource) {
	s.routes = append(s.routes, &Route{handler: r, path: url})
}

// AddBefore inserts a MiddlewareResource into the pre-dispatch middleware list
func (s *Server) AddBefore(m MiddlewareResource) {
	s.preDispatch = append(s.preDispatch, m)
}

// AddAfter inserts a MiddlewareResource into the post-dispatch middleware list
func (s *Server) AddAfter(m MiddlewareResource) {
	s.postDispatch = append(s.postDispatch, m)
}

// ServeHTTP Implements http.Handler interface into Server.
// Initializes a Context object and handles middleware and route actions.
// If an error is returned by any of the actions, the flow is stopped and a response is sent.
func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error

	// Set initial context data.
	// The Context pointer will be affected by the middleware and resources.
	ctx := NewContext(req, res)

	// Pre-Dispatch Middleware
	for _, m := range s.preDispatch {
		err = m.PreDispatch(ctx)
		if err != nil {
			// Return response on error
			s.Response(ctx, err)
			return
		}
	}

	// Route dispatch
	for _, r := range s.routes {
		if r.Match(req.URL.Path, ctx) {
			err = r.Dispatch(ctx)
			if err != nil {
				// Return response on error
				s.Response(ctx, err)
				return
			}
		}
	}

	// Post-Dispatch Middleware
	for _, m := range s.postDispatch {
		err = m.PostDispatch(ctx)
		if err != nil {
			// Return response on error
			s.Response(ctx, err)
			return
		}
	}

	// Return response
	s.Response(ctx, nil)
}

// Response writes the corresponding response to the HTTP response writer.
// It will handle the error status and the response body to be sent.
func (s *Server) Response(c *Context, err error) {
	// Error handling
	if err != nil {

	}

	// Write HTTP status
	c.Response.WriteHeader(c.responseStatus)

	// Write body
	c.Response.Write([]byte(c.responseContent))
}

// Start puts everything in place, sets http handlers and initiates a new http server.
func (s *Server) Start() {
	// Handle every request.
	http.Handle("/", s)

	// Run
	s.httpServer.ListenAndServe()
}

// TODO: StartTLS()
// Will start a HTTPS server based on previously configured TLS options and certs.
