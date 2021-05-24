package rockset_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func TestError_IsNotFoundError(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient(rockset.FromEnv())
	require.NoError(t, err)

	_, err = rc.GetCollection(ctx, "commons", "notfound")
	require.Error(t, err)

	var re rockset.Error
	if errors.As(err, &re) {
		require.True(t, re.IsNotFoundError())
	}
}
