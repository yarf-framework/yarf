package yarf

import (
	"net/http"
)

// YError is the interface used to handle error responses inside the framework.
type YError interface {
	Code() int    // HTTP response code for this error
	ID() int      // Error code ID.
	Msg() string  // Error description
	Body() string // Error body content to be returned to the client if needed.
}

// CustomError is the standard error response format used through the framework.
// Implements Error and YError interfaces
type CustomError struct {
	HTTPCode  int    // HTTP status code to be used as this error response.
	ErrorCode int    // Internal YARF error code for further reference.
	ErrorMsg  string // YARF error message.
	ErrorBody string // Error content to be rendered to the client response.
}

// Implements the error interface returning the ErrorMsg value of each error.
func (e *CustomError) Error() string {
	return e.ErrorMsg
}

// Code returns the error's HTTP code to be used in the response.
func (e *CustomError) Code() int {
	return e.HTTPCode
}

// ID returns the error's ID for further reference.
func (e *CustomError) ID() int {
	return e.ErrorCode
}

// Msg returns the error's message, used to implement the Error interface.
func (e *CustomError) Msg() string {
	return e.ErrorMsg
}

// Body returns the error's content body, if needed, to be returned in the HTTP response.
func (e *CustomError) Body() string {
	return e.ErrorBody
}

// UnexpectedError is used when the origin of the error can't be discovered
type UnexpectedError struct {
	CustomError
}

// ErrorUnexpected creates UnexpectedError
func ErrorUnexpected() *UnexpectedError {
	e := new(UnexpectedError)
	e.HTTPCode = http.StatusInternalServerError
	e.ErrorCode = 0
	e.ErrorMsg = "Unexpected error"

	return e
}

// MethodNotImplementedError is used to communicate that a specific HTTP method isn't implemented by a resource.
type MethodNotImplementedError struct {
	CustomError
}

// ErrorMethodNotImplemented creates MethodNotImplementedError
func ErrorMethodNotImplemented() *MethodNotImplementedError {
	e := new(MethodNotImplementedError)
	e.HTTPCode = http.StatusMethodNotAllowed
	e.ErrorCode = 1
	e.ErrorMsg = "Method not implemented"

	return e
}

// NotFoundError is the HTTP 404 error equivalent.
type NotFoundError struct {
	CustomError
}

// ErrorNotFound creates NotFoundError
func ErrorNotFound() *NotFoundError {
	e := new(NotFoundError)
	e.HTTPCode = http.StatusNotFound
	e.ErrorCode = 2
	e.ErrorMsg = "Not found"

	return e
}
