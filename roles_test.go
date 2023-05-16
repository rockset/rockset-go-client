package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type RoleIntegrationSuite struct {
	suite.Suite
	rc   *rockset.RockClient
	name string
}

func TestRoleIntegration(t *testing.T) {
	rc, randomName := vcrTestClient(t, t.Name())

	suite.Run(t, &RoleIntegrationSuite{
		rc:   rc,
		name: randomName("role"),
	})
}

func (s *RoleIntegrationSuite) TearDownSuite() {
	ctx := testCtx()
	err := s.rc.DeleteRole(ctx, s.name)
	s.NoError(err)
}

func (s *RoleIntegrationSuite) TestCreateRole() {
	ctx := testCtx()

	role, err := s.rc.CreateRole(ctx, s.name,
		option.WithRoleDescription(description()),
		option.WithWorkspacePrivilege(option.ListResourcesWs, persistentWorkspace),
	)
	s.NoError(err)
	s.Equal(s.name, role.GetRoleName())
}

func (s *RoleIntegrationSuite) TestGetRole() {
	ctx := testCtx()

	role, err := s.rc.GetRole(ctx, s.name)
	s.NoError(err)
	s.Equal(s.name, role.GetRoleName())
}

func (s *RoleIntegrationSuite) TestGetMissingRole() {
	ctx := testCtx()

	_, err := s.rc.GetRole(ctx, "non-existing-role")
	s.Error(err)
}

func (s *RoleIntegrationSuite) TestListRoles() {
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

func (s *RoleIntegrationSuite) TestUpdate() {
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
