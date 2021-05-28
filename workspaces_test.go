package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func TestRockClient_GetWorkspace(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	rc, err := rockset.NewClient(rockset.FromEnv())
	require.NoError(t, err)

	_, err = rc.GetWorkspace(ctx, "commons")
	require.NoError(t, err)
}

func TestRockClient_ListWorkspaces(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	rc, err := rockset.NewClient(rockset.FromEnv())
	require.NoError(t, err)

	workspaces, err := rc.ListWorkspaces(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, workspaces)
}
