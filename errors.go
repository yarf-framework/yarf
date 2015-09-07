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
// All custom errors should composite the CustomError in order to let know the framework what to do with each one.
type CustomError struct {
	httpCode  int    // HTTP status code to be used as this error response.
	errorCode int    // Internal YARF error code for further reference.
	errorMsg  string // YARF error message.
	errorBody string // Error content to be rendered to the client response.
}

// Implements the error interface returning the ErrorMsg value of each error.
func (e *CustomError) Error() string {
	return e.errorMsg
}

// Code returns the error's HTTP code to be used in the response.
func (e *CustomError) Code() int {
	return e.httpCode
}

// ID returns the error's ID for further reference.
func (e *CustomError) ID() int {
	return e.errorCode
}

// Msg returns the error's message, used to implement the Error interface.
func (e *CustomError) Msg() string {
	return e.errorMsg
}

// Body returns the error's content body, if needed, to be returned in the HTTP response.
func (e *CustomError) Body() string {
	return e.errorBody
}

// UnexpectedError is used when the origin of the error can't be discovered
type UnexpectedError struct {
	CustomError
}

// ErrorUnexpected creates UnexpectedError
func ErrorUnexpected() *UnexpectedError {
	e := new(UnexpectedError)

	e.httpCode = http.StatusInternalServerError
	e.errorCode = 0
	e.errorMsg = "Unexpected error"

	return e
}

// MethodNotImplementedError is used to communicate that a specific HTTP method isn't implemented by a resource.
type MethodNotImplementedError struct {
	CustomError
}

// ErrorMethodNotImplemented creates MethodNotImplementedError
func ErrorMethodNotImplemented() *MethodNotImplementedError {
	e := new(MethodNotImplementedError)

	e.httpCode = http.StatusMethodNotAllowed
	e.errorCode = 1
	e.errorMsg = "Method not implemented"

	return e
}

// NotFoundError is the HTTP 404 error equivalent.
type NotFoundError struct {
	CustomError
}

// ErrorNotFound creates NotFoundError
func ErrorNotFound() *NotFoundError {
	e := new(NotFoundError)

	e.httpCode = http.StatusNotFound
	e.errorCode = 2
	e.errorMsg = "Not found"

	return e
}
