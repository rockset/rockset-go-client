package rockset_test

import (
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

// for anyone poking around in the code, rockset.sleep() only works for this test org as no sane person would want
// to add a sleep in their query
const slowQuery = `script {{{ import * as rockset from "/rockset"; export function delay(x) { rockset.sleep(x); return x; } }}} select _script.delay(2000, q) from unnest([1] as q)`

type QueryIntegrationSuite struct {
	suite.Suite
	rc *rockset.RockClient
}

func TestQueryIntegration(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := QueryIntegrationSuite{
		rc: rc,
	}
	suite.Run(t, &s)
}

func (s *QueryIntegrationSuite) TestQuery() {
	ctx := testCtx()

	_, err := s.rc.Query(ctx, "SELECT 1")
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestListQueries() {
	ctx := testCtx()

	_, err := s.rc.ListActiveQueries(ctx)
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestAsyncQuery() {
	ctx := testCtx()

	resp, err := s.rc.Query(ctx, slowQuery,
		option.WithAsyncClientTimeout(100),
		option.WithAsyncMaxInitialResults(10),
	)
	s.Require().NoError(err)

	err = s.rc.WaitUntilQueryCompleted(ctx, *resp.QueryId)
	s.Require().NoError(err)

	_, err = s.rc.GetQueryResults(ctx, *resp.QueryId)
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestCancelQuery() {
	ctx := testCtx()

	resp, err := s.rc.Query(ctx, slowQuery,
		option.WithAsyncClientTimeout(100),
		option.WithAsyncMaxInitialResults(10),
	)
	s.Require().NoError(err)

	info, err := s.rc.CancelQuery(ctx, *resp.QueryId)
	s.Require().NoError(err)
	s.Require().Equal("CANCELLED", info.GetStatus())
}

func (s *QueryIntegrationSuite) TestValidateQuery() {
	ctx := testCtx()

	_, err := s.rc.ValidateQuery(ctx, "SELECT 1")
	s.Require().NoError(err)
}
