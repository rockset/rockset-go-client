package rockset_test

import (
	"testing"

	"github.com/rockset/rockset-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRockClient_GetOrganization(t *testing.T) {
	ctx := testCtx()

	rc, err := rockset.NewClient(rockset.FromEnv(), rockset.WithHTTPDebug())
	require.NoError(t, err)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	assert.Equal(t, "rockset-circleci", org.GetId())
}
