package errors

import (
	"errors"
	"net/http"

	"github.com/rockset/rockset-go-client/openapi"
)

// Error is an error returned by the Rockset REST API.
type Error struct {
	// ErrorModel is the ErrorModel of the underlying error
	*openapi.ErrorModel
	// Cause is the underlying cause of the error
	Cause error
	// StatusCode is the HTTP status code from the REST API call.
	StatusCode int
}

// New wraps err in an Error that provides better error messages than the openapi.GenericOpenAPIError
func New(err error) Error {
	var re = Error{Cause: err}

	var ge *openapi.GenericOpenAPIError
	if errors.As(err, &ge) {
		if m, ok := ge.Model().(openapi.ErrorModel); ok {
			re.ErrorModel = &m
		}
	}

	return re
}

// NewWithStatusCode wraps err in an Error that provides better error messages than the openapi.GenericOpenAPIError,
// and can be retried if the HTTP response StatusCode is in RetryableErrors. If err is nil,
// NewWithStatusCode() returns nil.
func NewWithStatusCode(err error, response *http.Response) error {
	if err == nil {
		return nil
	}

	e := New(err)
	if response != nil {
		e.StatusCode = response.StatusCode
	}

	return e
}

// Unwrap returns the cause of the Error
func (e Error) Unwrap() error {
	return e.Cause
}

// Error returns a string representation of the error
func (e Error) Error() string {
	if e.ErrorModel == nil {
		return e.Cause.Error()
	}

	return e.GetMessage()
}

// RateLimited checks if the error came from a http.StatusTooManyRequests, which is used for rate limiting.
func (e Error) RateLimited() bool {
	if e.ErrorModel == nil {
		return false
	}

	return e.StatusCode == http.StatusTooManyRequests
}

// IsNotFoundError returns true when the error is from an underlying 404 response from the Rockset REST API.
func (e Error) IsNotFoundError() bool {
	if e.ErrorModel == nil {
		return false
	}

	return e.StatusCode == http.StatusNotFound
}

// RetryableErrors are the errors which will cause the operation to be retried
var RetryableErrors = []int{
	http.StatusTooManyRequests,    // 429
	http.StatusServiceUnavailable, // 503
	http.StatusGatewayTimeout,     // 504
}

// Retryable returns true if the error is in RetryableErrors
func (e Error) Retryable() bool {
	if e.ErrorModel == nil {
		return false
	}

	for _, t := range RetryableErrors {
		if t == e.StatusCode {
			return true
		}
	}

	return false
}
