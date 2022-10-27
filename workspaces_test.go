package rockset_test

import (
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
	skipUnlessIntegrationTest(t)

	s := SuiteWorkspace{
		rc: testClient(t),
		ws: "integration",
	}
	suite.Run(t, &s)
}

func (s *SuiteWorkspace) SetupSuite() {
	ctx := testCtx()
	_, err := s.rc.CreateWorkspace(ctx, s.ws,
		option.WithWorkspaceDescription("created by the go client test suite"))
	s.Require().NoError(err)
	err = s.rc.WaitUntilWorkspaceAvailable(ctx, s.ws)
	s.Require().NoError(err)
}

func (s *SuiteWorkspace) TearDownSuite() {
	ctx := testCtx()
	err := s.rc.DeleteWorkspace(ctx, s.ws)
	s.Require().NoError(err)
}

func (s *SuiteWorkspace) TestGetWorkspace() {
	ctx := testCtx()
	ws, err := s.rc.GetWorkspace(ctx, s.ws)
	s.Require().NoError(err)
	s.Require().Equal(s.ws, ws.GetName())
}

func (s *SuiteWorkspace) TestListWorkspace() {
	ctx := testCtx()
	list, err := s.rc.ListWorkspaces(ctx)
	s.Require().NoError(err)
	s.Require().NotEmpty(list)
}
