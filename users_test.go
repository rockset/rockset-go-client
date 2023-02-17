package rockset_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

const CIUser = "pme+circleci@rockset.com"

type UserIntegrationSuite struct {
	suite.Suite
	rc    *rockset.RockClient
	email string
}

func TestUserIntegration(t *testing.T) {
	skipUnlessIntegrationTest(t)

	// we convert the email address to lowercase in Rockset, so must do it here or the test will fail
	name := strings.ToLower(randomName("test"))
	s := UserIntegrationSuite{
		rc:    testClient(t),
		email: fmt.Sprintf("pme+%s@rockset.com", name),
	}
	suite.Run(t, &s)
}

func (s *UserIntegrationSuite) TearDownSuite() {
	ctx := testCtx()

	err := s.rc.DeleteUser(ctx, s.email)
	s.Require().NoError(err)
}

func (s *UserIntegrationSuite) TestCreateUser() {
	ctx := testCtx()

	user, err := s.rc.CreateUser(ctx, s.email, []string{rockset.ReadOnlyRole})
	s.Require().NoError(err)
	s.Assert().Equal(s.email, user.GetEmail())
}

func (s *UserIntegrationSuite) TestGetCurrentUser() {
	ctx := testCtx()

	user, err := s.rc.GetCurrentUser(ctx)
	s.Require().NoError(err)
	s.Assert().Contains([]string{"ACTIVE", "NEW"}, user.GetState())
}

func (s *UserIntegrationSuite) TestGetUser() {
	ctx := testCtx()

	user, err := s.rc.GetUser(ctx, s.email)
	s.Require().NoError(err)
	s.Assert().Equal(s.email, user.Email)
	s.Assert().Equal("NEW", user.GetState())
}

func (s *UserIntegrationSuite) TestListUsers() {
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

func (s *UserIntegrationSuite) TestUpdateUser() {
	ctx := testCtx()

	_, err := s.rc.UpdateUser(ctx, s.email, []string{rockset.MemberRole},
		option.WithUserFirstName("first"), option.WithUserLastName("last"))
	s.Require().NoError(err)
	// TODO: can't assert the fist and last name until the user hac accepted
	// s.Assert().Equal("first", user.GetFirstName())
	// s.Assert().Equal("last", user.GetLastName())
}
