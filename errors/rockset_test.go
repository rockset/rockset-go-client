package errors_test

import (
	"errors"
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
	err := rockerr.NewWithStatusCode(errors.New("test error"), nil)

	var re rockerr.Error
	require.True(t, errors.As(err, &re))
	assert.False(t, re.IsNotFoundError())
	assert.Equal(t, 0, re.StatusCode)
}
