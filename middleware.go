package yarf

import ()

// MiddlewareHandler interface provides the methods for request filters
// that needs to run before, or after, every request Resource is executed.
// Both methods receives a Context pointer in case the middleware needs to modify Context data.
type MiddlewareHandler interface {
	PreDispatch() error
	PostDispatch() error
	SetContext(*Context)
}

// Middleware struct is the default implementation of a Middleware and does nothing.
// Users can either implement both methods or composite this struct into their own.
// Both methods needs to be present to satisfy the MiddlewareHandler interface.
type Middleware struct {
	RequestContext
}

// PreDispatch includes code to be executed before every Resource request.
func (m *Middleware) PreDispatch() error {
	return nil
}

// PostDispatch includes code to be executed after every Resource request.
func (m *Middleware) PostDispatch() error {
	return nil
}
