package yarf

import (
	"errors"
	"strings"
)

// Router interface provides the methods used to handle route and GroupRoute objects.
type Router interface {
	Path() string
	Match(string, *Context) bool
	Dispatch(*Context) error
}

// GroupRouter interface adds methods to work with children routers
type GroupRouter interface {
	Router
	Add(string, ResourceHandler)
	AddGroup(*GroupRoute)
	Insert(MiddlewareHandler)
}

// route struct stores the expected route path and the ResourceHandler that handles that route.
type route struct {
	prefix string // Group route prefix, if any.

	path string // Original route

	routeParts []string // Parsed Route split into parts

	handler ResourceHandler // Handler for the route
}

// Route returns a new route object initialized with the provided data.
// Params:
//	- url string 		// The route path to handle
//	- h	ResourceHandler	// The ResourceHandler object that will process the requests to the url.
//
func Route(url string, h ResourceHandler) Router {
	return &route{
		path:       url,
		handler:    h,
		routeParts: prepareURL(url),
	}
}

// Path is a route.path getter.
func (r *route) Path() string {
	return r.prefix + r.path
}

// Match returns true/false indicating if a request URL matches the route and
// sets the Context Params for matching parts in the original route.
// Route matchs are exact, that means, there are not optional parameters.
// To implement optional parameters you can define different routes handled by the same ResourceHandler.
// When a route matches the request URL, this method will parse and fill
// the parameters parsed during the process into the Context object.
func (r *route) Match(url string, c *Context) bool {
	requestParts := prepareURL(url)

	// YARF router only accepts exact route matches, so check for part count.
	// Unless it's a catch-all route
	if len(r.routeParts) == 0 || (len(r.routeParts) > 0 && r.routeParts[len(r.routeParts)-1] != "*") {
		if len(r.routeParts) != len(requestParts) {
			return false
		}
	}

	// check that requestParts matches routeParts
	if !matches(r.routeParts, requestParts) {
		return false
	}

	// It matches. Store route into Context.
	c.Route = r

	//storeParams(c, r.routeParts, requestParts)

	return true
}

// Dispatch executes the right ResourceHandler method based on the HTTP request in the Context object.
func (r *route) Dispatch(c *Context) error {
	// Method dispatch
	switch c.Request.Method {
	case "GET":
		return r.handler.Get(c)

	case "POST":
		return r.handler.Post(c)

	case "PUT":
		return r.handler.Put(c)

	case "PATCH":
		return r.handler.Patch(c)

	case "DELETE":
		return r.handler.Delete(c)

	case "OPTIONS":
		return r.handler.Options(c)

	case "HEAD":
		return r.handler.Head(c)

	case "TRACE":
		return r.handler.Trace(c)

	case "CONNECT":
		return r.handler.Connect(c)

	}

	// Return method not implemented
	return ErrorMethodNotImplemented()
}

// GroupRoute stores routes grouped under a single url prefix.
type GroupRoute struct {
	prefix string // The url prefix path for all routes in the group

	routeParts []string // parsed Route split into parts

	middleware []MiddlewareHandler // Group middleware resources

	routes []Router // Group routes
}

// RouteGroup creates a new GroupRoute object and initializes it with the provided url prefix.
// The object implements Router interface to being able to handle groups as routes.
// Groups can be nested into each other,
// so it's possible to add a GroupRoute as a route inside another GroupRoute.
// Includes methods to work with middleware.
func RouteGroup(url string) *GroupRoute {
	return &GroupRoute{
		prefix:     url,
		routeParts: prepareURL(url),
	}
}

// Path is a g.prefix getter
func (g *GroupRoute) Path() string {
	return g.prefix
}

// Match loops through all routes inside the group and find for one that matches the request.
// After a match is found, the route matching is stored into Context.groupDispatch
// to being able to dispatch it directly after a match without looping again.
// Outside the box, works exactly the same as route.Match()
func (g *GroupRoute) Match(url string, c *Context) bool {
	urlParts := prepareURL(url)

	// check if urlParts matches routeParts
	if !matches(g.routeParts, urlParts) {
		return false
	}

	// Remove prefix part form the request URL
	rURL := strings.Join(urlParts[len(g.routeParts):], "/")

	// Now look for a match inside the routes collection
	for _, r := range g.routes {
		if r.Match(rURL, c) {
			/*
				// store the matching Router and params after a match is found
				c.groupDispatch = append(c.groupDispatch, r)
				storeParams(c, g.routeParts, urlParts)
			*/

			// Store matching router into context
			//c.Route = r it was stored on route match

			return true
		}
	}

	return false
}

/*
// Dispatch loops through all routes inside the group and dispatches the one that matches the request.
*/
// Dispatch takes the matched route stored in context and dispatches it.
// Outside the box, works exactly the same as route.Dispatch().
func (g *GroupRoute) Dispatch(c *Context) (err error) {
	if c.Route == nil {
		g.endDispatch(c)
		return errors.New("No matching route found")
	}
	/*
		if len(c.groupDispatch) == 0 {
			g.endDispatch(c)
			return errors.New("No matching route found")
		}
	*/

	// Pre-dispatch middleware
	for _, m := range g.middleware {
		// Dispatch
		err = m.PreDispatch(c)
		if err != nil {
			g.endDispatch(c)
			return
		}
	}

	/*
		// pop, dispatch last route
		n := len(c.groupDispatch) - 1
		route := c.groupDispatch[n]
		c.groupDispatch = c.groupDispatch[:n]
		err = route.Dispatch(c)
		if err != nil {
			g.endDispatch(c)
			return
		}
	*/

	// Dispatch matched route
	err = c.Route.Dispatch(c)
	if err != nil {
		g.endDispatch(c)
		return
	}

	// Post-dispatch middleware
	for _, m := range g.middleware {
		// Dispatch
		err = m.PostDispatch(c)
		if err != nil {
			g.endDispatch(c)
			return
		}
	}

	// End dispatch if no errors blocking...
	g.endDispatch(c)

	// Return success
	return
}

func (g *GroupRoute) endDispatch(c *Context) (err error) {
	// End dispatch middleware
	for _, m := range g.middleware {
		e := m.End(c)
		if e != nil {
			// If there are any error, only return the last to be sure we go through all middlewares.
			err = e
		}
	}

	return
}

// Add inserts a new resource with it's associated route into the group object.
func (g *GroupRoute) Add(url string, h ResourceHandler) {
	g.routes = append(g.routes, &route{
		prefix:     g.prefix,
		path:       url,
		handler:    h,
		routeParts: prepareURL(url),
	})
}

// AddGroup inserts a GroupRoute into the routes list of the group object.
// This makes possible to nest groups.
func (g *GroupRoute) AddGroup(r *GroupRoute) {
	g.routes = append(g.routes, r)
}

// Insert adds a MiddlewareHandler into the middleware list of the group object.
func (g *GroupRoute) Insert(m MiddlewareHandler) {
	g.middleware = append(g.middleware, m)
}

// prepareUrl trims leading and trailing slahses, splits url parts, and removes empty parts
func prepareURL(url string) []string {
	return removeEmpty(strings.Split(url, "/"))
}

// removeEmpty removes blank strings from parts in one pass, shifting elements
// of the array down, and returns the altered array.
func removeEmpty(parts []string) []string {
	x := parts[:0]

	for _, p := range parts {
		if p != "" {
			x = append(x, p)
		}
	}

	return x
}

// matches returns true if requestParts matches routeParts up through len(routeParts)
// ignoring params in routeParts
func matches(routeParts, requestParts []string) bool {
	routeCount := len(routeParts)

	// Check for catch-all wildcard
	if len(routeParts) > 0 && routeParts[len(routeParts)-1] == "*" {
		routeCount--
	}

	if len(requestParts) < routeCount {
		return false
	}

	// Check for part matching, ignoring params and * wildcards
	for i, p := range routeParts {
		// Skip wildcard
		if p == "*" {
			continue
		}
		if p != requestParts[i] && p[0] != ':' {
			return false
		}
	}

	return true
}

/*
// storeParams writes parts from requestParts that correspond with param names in
// routeParts into c.Params.
func storeParams(c *Context, routeParts, requestParts []string) {
	for i, p := range routeParts {
		if p[0] == ':' {
			c.Params.Set(p[1:], requestParts[i])
		}
	}
}
*/
