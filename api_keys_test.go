package rockset_test

import (
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPIKey(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	ctx := testCtx()
	const keyName = "admin"
	const roleName = "admin"

	key, err := rc.CreateAPIKey(ctx, keyName, option.WithRole(roleName))
	require.NoError(t, err)

	key, err = rc.GetAPIKey(ctx, keyName)
	require.NoError(t, err)
	assert.Equal(t, roleName, key.GetRole())

	err = rc.DeleteAPIKey(ctx, keyName)
	require.NoError(t, err)
}
