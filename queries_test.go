package rockset_test

import (
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

const slowQuery = `script {{{ import * as rockset from "/rockset"; export function delay(x) { rockset.sleep(x); return x; } }}} select _script.delay(2000, q) from unnest([1] as q)`

type QueryTestSuite struct {
	suite.Suite
	rc *rockset.RockClient
}

func TestQueryTestSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := QueryTestSuite{
		rc: rc,
	}
	suite.Run(t, &s)
}

func (s *QueryTestSuite) TestQuery() {
	ctx := testCtx()

	_, err := s.rc.Query(ctx, "SELECT 1")
	s.Require().NoError(err)
}

func (s *QueryTestSuite) TestListQueries() {
	ctx := testCtx()

	_, err := s.rc.ListActiveQueries(ctx)
	s.Require().NoError(err)
}

func (s *QueryTestSuite) TestAsyncQuery() {
	ctx := testCtx()

	resp, err := s.rc.Query(ctx, slowQuery,
		option.WithAsyncClientTimeout(100),
		option.WithAsyncMaxInitialResults(10),
	)
	s.Require().NoError(err)

	err = s.rc.WaitForQuery(ctx, *resp.QueryId)
	s.Require().NoError(err)

	_, err = s.rc.GetQueryResults(ctx, *resp.QueryId)
	s.Require().NoError(err)
}

func (s *QueryTestSuite) TestCancelQuery() {
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

func (s *QueryTestSuite) TestValidateQuery() {
	ctx := testCtx()

	_, err := s.rc.ValidateQuery(ctx, "SEL:ECT 1")
	s.Require().NoError(err)
}
