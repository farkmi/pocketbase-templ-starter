package httperrors

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

var (
	ErrBadRequestNameRequired = NewHTTPError(http.StatusBadRequest, "MISSING_QUERY_PARAM_NAME", "Name is required")
)

type HTTPError struct {

	// HTTP status code returned for the error
	// Example: 403
	// Required: true
	// Maximum: 599
	// Minimum: 100
	Code int `json:"status"`

	// More detailed, human-readable, optional explanation of the error
	// Example: User is lacking permission to access this resource
	Detail string `json:"detail,omitempty"`

	// Short, human-readable description of the error
	// Example: Forbidden
	// Required: true
	Title string `json:"title"`

	// Type of error returned, should be used for client-side error handling
	// Example: generic
	// Required: true
	Type string `json:"type"`
}

func NewHTTPError(code int, errorType string, title string) *HTTPError {
	return &HTTPError{
		Code:  code,
		Type:  errorType,
		Title: title,
	}
}

func NewHTTPErrorWithDetail(code int, errorType string, title string, detail string) *HTTPError {
	return &HTTPError{
		Code:   code,
		Type:   errorType,
		Title:  title,
		Detail: detail,
	}
}

func ErrBadRequestInvalidEmail(e *core.RequestEvent) error {
	return e.String(http.StatusBadRequest, "Invalid email format")
}
