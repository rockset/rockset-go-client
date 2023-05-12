package rockset_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type SuiteWorkspace struct {
	suite.Suite
	rc *rockset.RockClient
	ws string
}

func TestSuiteWorkspace(t *testing.T) {
	rc, randomName := vcrClient(t)

	s := SuiteWorkspace{
		rc: rc,
		ws: randomName("ws"),
	}
	suite.Run(t, &s)
}

func (s *SuiteWorkspace) SetupSuite() {
	ctx := testCtx()
	_, err := s.rc.CreateWorkspace(ctx, s.ws,
		option.WithWorkspaceDescription(description()))
	s.Require().NoError(err)
	err = s.rc.WaitUntilWorkspaceAvailable(ctx, s.ws)
	s.Require().NoError(err)
}

func (s *SuiteWorkspace) TearDownSuite() {
	ctx := testCtx()
	err := s.rc.DeleteWorkspace(ctx, s.ws)
	s.Require().NoError(err)
	err = s.rc.WaitUntilWorkspaceGone(ctx, s.ws)
	s.Require().NoError(err)
}

func (s *SuiteWorkspace) TestGetWorkspace() {
	ctx := testCtx()
	ws, err := s.rc.GetWorkspace(ctx, s.ws)
	s.Require().NoError(err)
	s.Require().Equal(s.ws, ws.GetName())
}

func (s *SuiteWorkspace) TestGetPersistentWorkspace() {
	ctx := testCtx()
	ws, err := s.rc.GetWorkspace(ctx, persistentWorkspace)
	s.Require().NoError(err)
	s.Require().Equal(persistentWorkspace, ws.GetName())
}

func (s *SuiteWorkspace) TestGetNonExistingWorkspace() {
	ctx := testCtx()
	_, err := s.rc.GetWorkspace(ctx, randomString(16))
	s.Require().Error(err)
	var re rockset.Error
	s.Require().True(errors.As(err, &re))
	s.Assert().True(re.IsNotFoundError())
}

func (s *SuiteWorkspace) TestListWorkspace() {
	ctx := testCtx()
	list, err := s.rc.ListWorkspaces(ctx)
	s.Require().NoError(err)
	s.Require().NotEmpty(list)
}
