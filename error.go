package rockset

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rockset/rockset-go-client/openapi"
)

type Error struct {
	openapi.ErrorModel
}

// AsError tries to convert err to a rockset.Error, and returns true if it was possible
func AsError(err error, re *Error) bool {
	var ge openapi.GenericOpenAPIError
	if !errors.As(err, &ge) {
		return false
	}

	if m, ok := ge.Model().(openapi.ErrorModel); ok {
		re.ErrorModel = m
		return true
	}

	return false
}

func (e Error) Error() string {
	return e.GetMessage()
}

// RateLimited checks if the error came from a http.StatusTooManyRequests, which is used for rate limiting.
func (e Error) RateLimited() bool {
	return e.GetType() == statusWithoutSpace(http.StatusTooManyRequests)
}

// IsNotFoundError returns true when the error is from an underlying 404 response from the Rockset REST API.
func (e Error) IsNotFoundError() bool {
	return e.GetType() == statusWithoutSpace(http.StatusNotFound)
}

// RetryableErrors is the errors which will cause the operation to be retried
var RetryableErrors = []int{
	http.StatusTooManyRequests,    // 429
	http.StatusServiceUnavailable, // 503
	http.StatusGatewayTimeout,     // 504
}

// Retryable returns true if the error is in RetryableErrors
func (e Error) Retryable() bool {
	for _, t := range RetryableErrors {
		if statusWithoutSpace(t) == e.GetType() {
			return true
		}
	}

	return false
}

// strips space from http status codes as the openapi.ErrorModel Type field doesn't have spaces
func statusWithoutSpace(code int) string {
	return strings.Join(strings.Fields(http.StatusText(code)), "")
}
