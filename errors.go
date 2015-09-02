package yarf

import (
	"net/http"
)

// YarfError is the standard error response format used through the framework.
// All custom errors should composite the YarfError in order to let know the framework what to do with each one.
type YarfError struct {
	HttpCode  int    // HTTP status code to be used as this error response.
	ErrorCode int    // Internal YARF error code for further reference.
	ErrorMsg  string // YARF error message.
}

// Implements the error interface returning the ErrorMsg value of each error.
func (e *YarfError) Error() string {
	return e.ErrorMsg
}

//MethodNotImplementedError is used to indicate that the requested HTTP method isn't allowed for the actual resource.
type MethodNotImplementedError struct {
	YarfError
}

func ErrorMethodNotImplemented() *MethodNotImplementedError {
	e := new(MethodNotImplementedError)

	e.HttpCode = http.StatusMethodNotAllowed
	e.ErrorCode = 1
	e.ErrorMsg = "Method not implemented"

	return e
}
