package rockset_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func TestError_IsNotFoundError(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	_, err = rc.GetCollection(ctx, persistentWorkspace, "notfound")
	require.Error(t, err)

	var re rockset.Error
	if errors.As(err, &re) {
		require.True(t, re.IsNotFoundError())
	}
}

func TestError_Retryable(t *testing.T) {
	for _, status := range rockset.RetryableErrors {
		var err rockset.Error
		require.True(t, errors.As(fakeError(status), &err))
		assert.True(t, err.Retryable())
	}
}
