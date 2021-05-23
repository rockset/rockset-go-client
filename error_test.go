package rockset_test

import (
	"errors"
	"testing"

	"github.com/rockset/rockset-go-client"
	"github.com/stretchr/testify/require"
)

func TestError_Nil(t *testing.T) {
	var re rockset.Error
	require.False(t, rockset.AsError(nil, &re))
}

func TestError_Plain(t *testing.T) {
	var re rockset.Error
	require.False(t, rockset.AsError(errors.New("dummy"), &re))
}

func TestError_IsNotFoundError(t *testing.T) {
	ctx := testCtx()

	rc, err := rockset.NewClient(rockset.FromEnv())
	require.NoError(t, err)

	_, err = rc.GetCollection(ctx, "commons", "notfound")
	require.Error(t, err)

	var re rockset.Error
	if rockset.AsError(err, &re) {
		require.True(t, re.IsNotFoundError())
	}
}
