package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type RoleTestSuite struct {
	suite.Suite
	rc   *rockset.RockClient
	name string
}

func TestRoleTestSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	suite.Run(t, &RoleTestSuite{
		rc:   rc,
		name: randomName(t, "role"),
	})
}

func (s *RoleTestSuite) TearDownSuite() {
	ctx := testCtx()
	err := s.rc.DeleteRole(ctx, s.name)
	s.NoError(err)
}

func (s *RoleTestSuite) TestCreateRole() {
	ctx := testCtx()

	role, err := s.rc.CreateRole(ctx, s.name,
		option.WithRoleDescription(description()),
		option.WithWorkspacePrivilege(option.ListResourcesWs, persistentWorkspace),
	)
	s.NoError(err)
	s.Equal(s.name, role.GetRoleName())
}

func (s *RoleTestSuite) TestGetRole() {
	ctx := testCtx()

	role, err := s.rc.GetRole(ctx, s.name)
	s.NoError(err)
	s.Equal(s.name, role.GetRoleName())
}

func (s *RoleTestSuite) TestGetMissingRole() {
	ctx := testCtx()

	_, err := s.rc.GetRole(ctx, "non-existing-role")
	s.Error(err)
}

func (s *RoleTestSuite) TestListRoles() {
	ctx := testCtx()

	roles, err := s.rc.ListRoles(ctx)
	s.NoError(err)

	var found bool
	for _, role := range roles {
		if role.GetRoleName() == s.name {
			found = true
		}
	}
	s.True(found)
}

func (s *RoleTestSuite) TestUpdate() {
	ctx := testCtx()

	role, err := s.rc.UpdateRole(ctx, s.name,
		option.WithGlobalPrivilege(option.ListRolesGlobal),
		option.WithRoleDescription("go client test role with updated name"),
		option.WithIntegrationPrivilege(option.CreateCollectionIntegration, "test"),
		option.WithWorkspacePrivilege(option.QueryDataWs, "commons"),
		option.WithWorkspacePrivilege(option.CreateViewWs, "commons", option.WithCluster("usw2a1")),
	)
	s.NoError(err)
	s.Equal(s.name, role.GetRoleName())
}
