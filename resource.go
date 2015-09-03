package yarf

import ()

// The Resource type is the representation of each REST resource of the application.
// It implements the RestResource interface and allows the developer to extend the methods needed.
// All resources being used by a YARF application have to composite this Resource struct.
type Resource struct {
	RequestContext
}

// Implementations for all HTTP methods.
// The default implementation will return a 405 HTTP error indicating that the method isn't allowed.
// Once a resource composites the Resource type, it will implement/overwrite the methods needed.

// HTTP GET
func (r *Resource) Get() error {
	return ErrorMethodNotImplemented()
}

// HTTP POST
func (r *Resource) Post() error {
	return ErrorMethodNotImplemented()
}

// HTTP PUT
func (r *Resource) Put() error {
	return ErrorMethodNotImplemented()
}

// HTTP PATCH
func (r *Resource) Patch() error {
	return ErrorMethodNotImplemented()
}

// HTTP DELETE
func (r *Resource) Delete() error {
	return ErrorMethodNotImplemented()
}

// HTTP OPTIONS
func (r *Resource) Options() error {
	return ErrorMethodNotImplemented()
}

// HTTP HEAD
func (r *Resource) Head() error {
	return ErrorMethodNotImplemented()
}

// HTTP TRACE
func (r *Resource) Trace() error {
	return ErrorMethodNotImplemented()
}

// HTTP CONNECT
func (r *Resource) Connect() error {
	return ErrorMethodNotImplemented()
}

// The RestResource interface defines how Resources through the application have to be defined.
// Ideally, the developer will composite the Resource struct into their own resources,
// but it's possible to implement each one by their own.
type RestResource interface {
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
