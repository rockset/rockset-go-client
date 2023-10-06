package errors_test

import (
	"errors"
	"github.com/rockset/rockset-go-client/openapi"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/internal/test"
)

func TestError_IsNotFoundError(t *testing.T) {
	test.SkipUnlessIntegrationTest(t)
	rc := test.Client(t)
	ctx := test.Context()

	_, err := rc.GetCollection(ctx, "notfound", "notfound")
	require.Error(t, err)

	var re rockerr.Error
	require.True(t, errors.As(err, &re))
	assert.True(t, re.IsNotFoundError())
	assert.Equal(t, http.StatusNotFound, re.StatusCode)
}

func TestError_NilHTTPResponse(t *testing.T) {
	e := errors.New("test error")
	err := rockerr.NewWithStatusCode(e, nil)

	var re rockerr.Error
	require.True(t, errors.As(err, &re))
	assert.False(t, re.IsNotFoundError())
	assert.False(t, re.RateLimited())
	assert.False(t, re.Retryable())
	assert.Equal(t, e, re.Unwrap())
	assert.Equal(t, "test error", err.Error())
}

func TestError_HTTPResponse(t *testing.T) {
	e := errors.New("test error")
	err := rockerr.NewWithStatusCode(e, &http.Response{StatusCode: http.StatusTooManyRequests})

	var re rockerr.Error
	require.True(t, errors.As(err, &re))

	re.ErrorModel = &openapi.ErrorModel{
		Message: openapi.PtrString("api error"),
	}

	assert.False(t, re.IsNotFoundError())
	assert.True(t, re.RateLimited())
	assert.True(t, re.Retryable())
	assert.Equal(t, e, re.Unwrap())
	assert.Equal(t, "api error", re.Error())
	assert.Equal(t, "test error", err.Error())
}

func TestError_Nil(t *testing.T) {
	err := rockerr.NewWithStatusCode(nil, nil)

	assert.Nil(t, err)
}
