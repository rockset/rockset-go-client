package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type RoleTestSuite struct {
	suite.Suite
	name string
	rc   *rockset.RockClient
}

func (s *RoleTestSuite) SetupSuite() {
	skipUnlessIntegrationTest(s.T())
	s.name = "foobar"
	rc, err := rockset.NewClient()
	s.NoError(err)
	s.rc = rc
}

func (s *RoleTestSuite) TearDownSuite() {
	ctx := testCtx()
	_ = s.rc.DeleteRole(ctx, s.name)
}

func TestRoleTestSuite(t *testing.T) {
	suite.Run(t, &RoleTestSuite{})
}

func (s *RoleTestSuite) TestCRUD() {
	ctx := testCtx()

	role, err := s.rc.CreateRole(ctx, s.name, option.WithRoleDescription("custom role"))
	s.NoError(err)

	r, err := s.rc.GetRole(ctx, s.name)
	s.NoError(err)
	s.Equal(r.GetRoleName(), role.GetRoleName())

	roles, err := s.rc.ListRoles(ctx)
	s.NoError(err)

	var found bool
	for _, r := range roles {
		if role.GetRoleName() == r.GetRoleName() {
			found = true
		}
	}
	s.True(found)

	role, err = s.rc.UpdateRole(ctx, s.name,
		option.WithGlobalPrivilege(option.ListRolesGlobal),
		option.WithIntegrationPrivilege(option.CreateCollectionIntegration, "test"),
		option.WithWorkspacePrivilege(option.QueryDataWs, "ws"),
		option.WithWorkspacePrivilege(option.CreateViewWs, "ws", option.WithCluster("rs2")),
	)
	s.NoError(err)

	err = s.rc.DeleteRole(ctx, s.name)
	s.NoError(err)
}
