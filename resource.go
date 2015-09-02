package yarf

import ()


// The Resource type is the representation of each REST resource of the application.
// It implements the RestResource interface by default and allow the developer to extend the methods needed.
// All resources being used by a YARF application have to composite this Resource struct.
type Resource struct {
    Context *Context // Include the request context as a property of the Resource.
}

// Context getter
func (r *Resource) Context() *Context {
    return r.Context
}

// Context setter
func (r *Resource) SetContext(c *Context) {
    r.Context = c
}

// Render takes a string and aggregates it to the Context.responseContent.
func (r *Resource) Render(content string) {
    r.Context.responseContent += content
}

// Status sets the HTTP status code to be returned on the response.
func (r *Resource) Status(code int) {
    r.Context.responseStatusCode = code
}


// Implementations for all HTTP methods.
// The default implementation will return a 405 HTTP error indicating that the method isn't allowed.
// Once a resource composites the Resource type, it will implement/overwrite the methods needed.

// HTTP GET
func (r *Resource) Get() error {
    return MethodNotImplementedError()
}

// HTTP POST
func (r *Resource) Post() error {
    return MethodNotImplementedError()
}

// HTTP PUT
func (r *Resource) Put() error {
    return MethodNotImplementedError()
}

// HTTP DELETE
func (r *Resource) Delete() error {
    return MethodNotImplementedError()
}

// HTTP OPTIONS
func (r *Resource) Options() error {
    return MethodNotImplementedError()
}

// HTTP HEAD
func (r *Resource) Head() error {
    return MethodNotImplementedError()
}

// HTTP TRACE
func (r *Resource) Trace() error {
    return MethodNotImplementedError()
}

// HTTP CONNECT
func (r *Resource) Connect() error {
    return MethodNotImplementedError()
}

// The RestResource interface defines how Resources through the application have to be defined.
// Ideally, the developer will composite the Resource struct into their own resources,
// but it's possible to implement each one by their own.
type RestResource interface {
    // HTTP methods
    Get() error
    Post() error
    Put() error
    Delete() error
    Options() error
    Head() error
    Trace() error
    Connect() error
}
