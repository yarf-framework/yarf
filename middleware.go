package yarf

// MiddlewareHandler interface provides the methods for request filters
// that needs to run before, or after, every request Resource is executed.
type MiddlewareHandler interface {
	PreDispatch(*Context) error
	PostDispatch(*Context) error
	End(*Context) error
}

// Middleware struct is the default implementation of a Middleware and does nothing.
// Users can either implement both methods or composite this struct into their own.
// Both methods needs to be present to satisfy the MiddlewareHandler interface.
type Middleware struct{}

// PreDispatch includes code to be executed before every Resource request.
func (m *Middleware) PreDispatch(c *Context) error {
	return nil
}

// PostDispatch includes code to be executed after every Resource request.
func (m *Middleware) PostDispatch(c *Context) error {
	return nil
}

// End will be executed ALWAYS after every request, even if there were errors present.
func (m *Middleware) End(c *Context) error {
	return nil
}
