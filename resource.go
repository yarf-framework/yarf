package yarf

import ()

// The ResourceHandler interface defines how Resources through the application have to be defined.
// Ideally, the developer will composite the Resource struct into their own resources,
// but it's possible to implement each one by their own.
type ResourceHandler interface {
	// HTTP methods
	Get() error
	Post() error
	Put() error
	Patch() error
	Delete() error
	Options() error
	Head() error
	Trace() error
	Connect() error

	// Context setter
	SetContext(*Context)
}

// The Resource type is the representation of each REST resource of the application.
// It implements the ResourceHandler interface and allows the developer to extend the methods needed.
// All resources being used by a YARF application have to composite this Resource struct.
type Resource struct {
	RequestContext
}

// Implementations for all HTTP methods.
// The default implementation will return a 405 HTTP error indicating that the method isn't allowed.
// Once a resource composites the Resource type, it will implement/overwrite the methods needed.

// Default HTTP GET implementation
func (r *Resource) Get() error {
	return ErrorMethodNotImplemented()
}

// Default HTTP POST implementation
func (r *Resource) Post() error {
	return ErrorMethodNotImplemented()
}

// Default HTTP PUT implementation
func (r *Resource) Put() error {
	return ErrorMethodNotImplemented()
}

// Default HTTP PATCH implementation
func (r *Resource) Patch() error {
	return ErrorMethodNotImplemented()
}

// Default HTTP DELETE implementation
func (r *Resource) Delete() error {
	return ErrorMethodNotImplemented()
}

// Default HTTP OPTIONS implementation
func (r *Resource) Options() error {
	return ErrorMethodNotImplemented()
}

// Default HTTP HEAD implementation
func (r *Resource) Head() error {
	return ErrorMethodNotImplemented()
}

// Default HTTP TRACE
func (r *Resource) Trace() error {
	return ErrorMethodNotImplemented()
}

// Default HTTP CONNECT
func (r *Resource) Connect() error {
	return ErrorMethodNotImplemented()
}
