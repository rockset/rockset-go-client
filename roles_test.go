package rockset_test

import (
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRole_CRUD(t *testing.T) {
	skipUnlessIntegrationTest(t)

	name := "foobar"
	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	role, err := rc.CreateRole(ctx, name, option.WithRoleDescription("custom role"))
	require.NoError(t, err)

	// delete in defer, so we clean up in case of error
	defer func() {
		err := rc.DeleteRole(ctx, name)
		assert.NoError(t, err)
	}()

	roles, err := rc.ListRoles(ctx)
	require.NoError(t, err)

	var found bool
	for _, r := range roles {
		if role.GetRoleName() == r.GetRoleName() {
			found = true
		}
	}
	assert.True(t, found)

	role, err = rc.UpdateRole(ctx, name,
		option.WithGlobalPrivilege(option.ListRolesGlobal),
		option.WithIntegrationPrivilege(option.CreateCollectionIntegration, "test"),
		option.WithWorkspacePrivilege(option.QueryData, "ws"),
		option.WithWorkspacePrivilege(option.CreateView, "ws", option.WithCluster("rs2")),
	)
	require.NoError(t, err)
}
