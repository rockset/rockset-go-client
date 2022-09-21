package rockset_test

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

const CIUser = "pme+circleci@rockset.com"

type UserTestSuite struct {
	suite.Suite
	rc    *rockset.RockClient
	email string
}

func TestUserTestSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := UserTestSuite{
		rc:    rc,
		email: "pme+testuser@rockset.com",
	}
	suite.Run(t, &s)
}

func (s *UserTestSuite) TearDownSuite() {
	ctx := testCtx()

	err := s.rc.DeleteUser(ctx, s.email)
	s.Require().NoError(err)
}

func (s *UserTestSuite) TestCreateUser() {
	ctx := testCtx()

	_, err := s.rc.CreateUser(ctx, s.email, []string{"read-only"})
	s.Require().NoError(err)
}

func (s *UserTestSuite) TestGetCurrentUser() {
	ctx := testCtx()

	user, err := s.rc.GetCurrentUser(ctx)
	s.Require().NoError(err)
	s.Assert().Equal(CIUser, user.Email)
}

func (s *UserTestSuite) TestGetUser() {
	ctx := testCtx()

	user, err := s.rc.GetUser(ctx, s.email)
	s.Require().NoError(err)
	s.Assert().Equal(s.email, user.Email)
}

func (s *UserTestSuite) TestListUsers() {
	ctx := testCtx()

	users, err := s.rc.ListUsers(ctx)
	s.Require().NoError(err)

	var found bool
	for _, user := range users {
		if user.Email == CIUser {
			found = true
		}
	}
	s.Assert().True(found)
}
