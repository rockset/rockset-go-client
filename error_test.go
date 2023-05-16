package rockset_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func TestError_IsNotFoundError(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := testCtx()

	_, err := rc.GetCollection(ctx, persistentWorkspace, "notfound")
	require.Error(t, err)

	var re rockset.Error
	require.True(t, errors.As(err, &re))
	assert.True(t, re.IsNotFoundError())
	assert.Equal(t, http.StatusNotFound, re.StatusCode)
}

func TestError_NilHTTPResponse(t *testing.T) {
	err := rockset.NewErrorWithStatusCode(errors.New("test error"), nil)

	var re rockset.Error
	require.True(t, errors.As(err, &re))
	assert.False(t, re.IsNotFoundError())
	assert.Equal(t, 0, re.StatusCode)
}
