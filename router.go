package yarf

import (
	"errors"
	"strings"
)

// Router interface provides the methods used to handle route and routeGroup objects.
type Router interface {
	Match(string, *Context) bool
	Dispatch(*Context) error
}

// route struct stores the expected route path and the ResourceHandler that handles that route.
type route struct {
	path string // Original route

	routeParts []string // parsed Route split into parts

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

// Match returns true/false indicating if a request URL matches the route and
// sets the Context Params for matching parts in the original route.
// Route matchs are exact, that means, there are not optional parameters.
// To implement optional parameters you can define different routes handled by the same ResourceHandler.
// When a route matches the request URL, this method will parse and fill
// the parameters parsed during the process into the Context object.
func (r *route) Match(url string, c *Context) bool {
	requestParts := prepareURL(url)

	// YARF router only accepts exact route matches, so check for part count.
	if len(r.routeParts) != len(requestParts) {
		return false
	}

	// check that requestParts matches routeParts
	if !matches(r.routeParts, requestParts) {
		return false
	}

	storeParams(c, r.routeParts, requestParts)

	return true
}

// Dispatch executes the right ResourceHandler method based on the HTTP request in the Context object.
func (r *route) Dispatch(c *Context) (err error) {
	// Add Context to handler
	r.handler.SetContext(c)

	// Method dispatch
	switch c.Request.Method {
	case "GET":
		err = r.handler.Get()

	case "POST":
		err = r.handler.Post()

	case "PUT":
		err = r.handler.Put()

	case "PATCH":
		err = r.handler.Patch()

	case "DELETE":
		err = r.handler.Delete()

	case "OPTIONS":
		err = r.handler.Options()

	case "HEAD":
		err = r.handler.Head()

	case "TRACE":
		err = r.handler.Trace()

	case "CONNECT":
		err = r.handler.Connect()

	default:
		err = ErrorMethodNotImplemented()
	}

	// Return error status
	return
}

// routeGroup stores routes grouped under a single url prefix.
type routeGroup struct {
	prefix string // The url prefix path for all routes in the group

	routeParts []string // parsed Route split into parts

	middleware []MiddlewareHandler // Group middleware resources

	routes []Router // Group routes

	lastMatch Router // Stores last matched route to be dispatched.
}

// RouteGroup creates a new routeGroup object and initializes it with the provided url prefix.
// The object implements Router interface to being able to handle groups as routes.
// Groups can be nested into each other,
// so it's possible to add a routeGroup as a route inside another routeGroup.
// Includes methods to work with middleware.
func RouteGroup(url string) *routeGroup {
	return &routeGroup{
		prefix:     url,
		routeParts: prepareURL(url),
	}
}

// Match loops through all routes inside the group and find for one that matches the request.
// After a match is found, the route matching is stored into lastMatch
// to being able to dispatch it directly after a match without looping again.
// Outside the box, works exactly the same as route.Match()
func (g *routeGroup) Match(url string, c *Context) bool {
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
			// store the matching Router and params after a match is found
			g.lastMatch = r
			storeParams(c, g.routeParts, urlParts)
			return true
		}
	}

	return false
}

// Dispatch loops through all routes inside the group and dispatch the one that matches the request.
// Outside the box, works exactly the same as route.Dispatch().
func (g *routeGroup) Dispatch(c *Context) (err error) {
	if g.lastMatch == nil {
		return errors.New("No matching route found")
	}

	// Pre-dispatch middleware
	for _, m := range g.middleware {
		// Add context to middleware
		m.SetContext(c)

		// Dispatch
		err = m.PreDispatch()
		if err != nil {
			return
		}
	}

	// Dispatch route
	err = g.lastMatch.Dispatch(c)
	if err != nil {
		return
	}

	// Post-dispatch middleware
	for _, m := range g.middleware {
		// Dispatch
		err = m.PostDispatch()
		if err != nil {
			return
		}
	}

	// Return success
	return
}

// Add inserts a new resource with it's associated route into the group object.
func (g *routeGroup) Add(url string, h ResourceHandler) {
	g.routes = append(g.routes, Route(url, h))
}

// AddGroup inserts a route group into the routes list of the group object.
// This makes possible to nest groups.
func (g *routeGroup) AddGroup(r *routeGroup) {
	g.routes = append(g.routes, r)
}

// Insert adds a MiddlewareHandler into the middleware list of the group object.
func (g *routeGroup) Insert(m MiddlewareHandler) {
	g.middleware = append(g.middleware, m)
}

// prepareUrl trims leading and trailing slahses, splits url parts, and removes empty parts
func prepareURL(url string) []string {
	return removeEmpty(strings.Split(trimSlash(url), "/"))
}

// trimSlash cleans all leading and trailing "/" from request url
func trimSlash(url string) string {
	for len(url) > 0 && url[0] == '/' {
		url = url[1:]
	}
	for len(url) > 0 && url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	return url
}

// removeEmpty removes blank strings from parts in one pass, shifting elements
// of the array down, and returns the altered array.
func removeEmpty(parts []string) []string {
	i := 0
	for j, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = parts[j]
		i++
	}
	return parts[:i]
}

// matches returns true if requestParts matches routeParts up through len(routeParts)
// ignoring params in routeParts
func matches(routeParts, requestParts []string) bool {
	if len(requestParts) < len(routeParts) {
		return false
	}

	// Check for part matching, ignoring params
	for i, p := range routeParts {
		if p != requestParts[i] && p[0] != ':' {
			return false
		}
	}

	return true
}

// storeParams writes parts from requestParts that correspond with param names in
// routeParts into c.Params.
func storeParams(c *Context, routeParts, requestParts []string) {
	for i, p := range routeParts {
		if p[0] == ':' {
			c.Params.Set(p[1:], requestParts[i])
		}
	}
}
