package rockset_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

type SuiteWorkspace struct {
	suite.Suite
	rc *rockset.RockClient
	ws string
}

func TestSuiteWorkspace(t *testing.T) {
	rc, randomName := vcrTestClient(t, t.Name())

	s := SuiteWorkspace{
		rc: rc,
		ws: randomName("ws"),
	}
	suite.Run(t, &s)
}

func (s *SuiteWorkspace) SetupSuite() {
	ctx := test.Context()
	_, err := s.rc.CreateWorkspace(ctx, s.ws,
		option.WithWorkspaceDescription(test.Description()))
	s.Require().NoError(err)
	err = s.rc.Wait.UntilWorkspaceAvailable(ctx, s.ws)
	s.Require().NoError(err)
}

func (s *SuiteWorkspace) TearDownSuite() {
	ctx := test.Context()
	err := s.rc.DeleteWorkspace(ctx, s.ws)
	s.Require().NoError(err)
	err = s.rc.Wait.UntilWorkspaceGone(ctx, s.ws)
	s.Require().NoError(err)
}

func (s *SuiteWorkspace) TestGetWorkspace() {
	ctx := test.Context()
	ws, err := s.rc.GetWorkspace(ctx, s.ws)
	s.Require().NoError(err)
	s.Require().Equal(s.ws, ws.GetName())
}

func (s *SuiteWorkspace) TestGetPersistentWorkspace() {
	ctx := test.Context()
	ws, err := s.rc.GetWorkspace(ctx, test.PersistentWorkspace)
	s.Require().NoError(err)
	s.Require().Equal(test.PersistentWorkspace, ws.GetName())
}

func (s *SuiteWorkspace) TestGetNonExistingWorkspace() {
	ctx := test.Context()
	_, err := s.rc.GetWorkspace(ctx, test.RandomString(16))
	s.Require().Error(err)
	var re rockerr.Error
	s.Require().True(errors.As(err, &re))
	s.Assert().True(re.IsNotFoundError())
}

func (s *SuiteWorkspace) TestListWorkspace() {
	ctx := test.Context()
	list, err := s.rc.ListWorkspaces(ctx)
	s.Require().NoError(err)
	s.Require().NotEmpty(list)
}
