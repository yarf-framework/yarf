package yarf

import ()

// The ResourceHandler interface defines how Resources through the application have to be defined.
// Ideally, the developer will composite the Resource struct into their own resources,
// but it's possible to implement each one by their own.
type ResourceHandler interface {
	// HTTP methods
	Get(*Context) error
	Post(*Context) error
	Put(*Context) error
	Patch(*Context) error
	Delete(*Context) error
	Options(*Context) error
	Head(*Context) error
	Trace(*Context) error
	Connect(*Context) error
}

// The Resource type is the representation of each REST resource of the application.
// It implements the ResourceHandler interface and allows the developer to extend the methods needed.
// All resources being used by a YARF application have to composite this Resource struct.
type Resource struct{}

// Implementations for all HTTP methods.
// The default implementation will return a 405 HTTP error indicating that the method isn't allowed.
// Once a resource composites the Resource type, it will implement/overwrite the methods needed.

// Get is the default HTTP GET implementation.
// It returns a NotImplementedError
func (r *Resource) Get(c *Context) error {
	return ErrorMethodNotImplemented()
}

// Post is the default HTTP POST implementation.
// It returns a NotImplementedError
func (r *Resource) Post(c *Context) error {
	return ErrorMethodNotImplemented()
}

// Put is the default HTTP PUT implementation.
// It returns a NotImplementedError
func (r *Resource) Put(c *Context) error {
	return ErrorMethodNotImplemented()
}

// Patch is the default HTTP PATCH implementation.
// It returns a NotImplementedError
func (r *Resource) Patch(c *Context) error {
	return ErrorMethodNotImplemented()
}

// Delete is the default HTTP DELETE implementation.
// It returns a NotImplementedError
func (r *Resource) Delete(c *Context) error {
	return ErrorMethodNotImplemented()
}

// Options is the default HTTP OPTIONS implementation.
// It returns a NotImplementedError
func (r *Resource) Options(c *Context) error {
	return ErrorMethodNotImplemented()
}

// Head is the default HTTP HEAD implementation.
// It returns a NotImplementedError
func (r *Resource) Head(c *Context) error {
	return ErrorMethodNotImplemented()
}

// Trace is the default HTTP TRACE implementation.
// It returns a NotImplementedError
func (r *Resource) Trace(c *Context) error {
	return ErrorMethodNotImplemented()
}

// Connect is the default HTTP CONNECT implementation.
// It returns a NotImplementedError
func (r *Resource) Connect(c *Context) error {
	return ErrorMethodNotImplemented()
}
