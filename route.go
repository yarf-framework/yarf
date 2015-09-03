package yarf

import (
	"strings"
)

// Route struct stores the expected route path and the RestResource that handles that route.
type Route struct {
	path    string
	handler RestResource
}

// Match returns true/false indicating if a request URL matches the route.
// Route matchs are exact, that means, there are not optional parameters.
// To implement optional parameters you can define different routes handled by the same RestResource.
// When a route matches the request URL, this method will parse and fill
// the parameters parsed during the process into the Context object.
func (r *Route) Match(url string, c *Context) bool {
	// Init params
	params := make(map[string]string)

	// Copy route path value
	route := r.path

	// Clean initial and trailing "/" from request and routes
	for strings.HasPrefix(route, "/") {
		route = strings.TrimPrefix(route, "/")
	}
	for strings.HasPrefix(url, "/") {
		url = strings.TrimPrefix(url, "/")
	}
	for strings.HasSuffix(route, "/") {
		route = strings.TrimSuffix(route, "/")
	}
	for strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	// Split parts
	routeParts := strings.Split(route, "/")
	urlParts := strings.Split(url, "/")

	// YARF router only accepts exact route matches, so check for part count.
	if len(urlParts) != len(routeParts) {
		return false
	}

	// Check for param matching
	if route != url {
		for i, r := range routeParts {
			// Check part
			if r != urlParts[i] && r != "" && r[:1] != ":" {
				return false
			}

			// Check param
			if r != "" && r[:1] == ":" {
				// Empty params aren't params
				if urlParts[i] == "" {
					return false
				}

				params[r[1:]] = urlParts[i]
			}
		}
	}

	// Success match. Store params and return true.
	for key, value := range params {
		c.Params.Set(key, value)
	}
	return true
}

// Dispatch executes the right RestResource method based on the HTTP request in the Context object.
// Accepts method override, based on request header: X-HTTP-Method-Override
func (r *Route) Dispatch(c *Context) error {
	// Init error status
	var err error

	// Get HTTP method requested
	method := strings.ToUpper(c.Request.Method)

	// Check for method overriding
	mo := strings.ToUpper(c.Request.Header.Get("X-HTTP-Method-Override"))
	if mo == "PUT" || mo == "PATCH" || mo == "DELETE" {
		method = mo
	}

	// Add Context to handler
	r.handler.SetContext(c)

	// Method dispatch
	switch method {
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
	}

	// Return error status
	return err
}
