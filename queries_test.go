package rockset_test

import (
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/suite"
	"testing"
)

// for anyone poking around in the code, rockset.sleep() only works for this test org as no sane person would want
// to add a sleep in their query
const slowQuery = `script {{{ import * as rockset from "/rockset"; export function delay(x) { rockset.sleep(x); return x; } }}} select _script.delay(2000, q) from unnest([1] as q)`

type QueryIntegrationSuite struct {
	suite.Suite
}

func TestQueryIntegration(t *testing.T) {

	s := QueryIntegrationSuite{}
	suite.Run(t, &s)
}

func (s *QueryIntegrationSuite) TestQuery() {
	ctx := testCtx()
	rc, _ := vcrClient(s.T())

	_, err := rc.Query(ctx, "SELECT 1")
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestListQueries() {
	ctx := testCtx()
	rc, _ := vcrClient(s.T())

	_, err := rc.ListActiveQueries(ctx)
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestAsyncQuery() {
	ctx := testCtx()
	rc, _ := vcrClient(s.T())

	resp, err := rc.Query(ctx, slowQuery,
		option.WithAsyncClientTimeout(100),
		option.WithAsyncMaxInitialResults(10),
	)
	s.Require().NoError(err)

	err = rc.WaitUntilQueryCompleted(ctx, *resp.QueryId)
	s.Require().NoError(err)

	_, err = rc.GetQueryResults(ctx, *resp.QueryId)
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestCancelQuery() {
	ctx := testCtx()
	rc, _ := vcrClient(s.T())

	resp, err := rc.Query(ctx, slowQuery,
		option.WithAsyncClientTimeout(100),
		option.WithAsyncMaxInitialResults(10),
	)
	s.Require().NoError(err)

	info, err := rc.CancelQuery(ctx, *resp.QueryId)
	s.Require().NoError(err)
	s.Require().Equal("CANCELLED", info.GetStatus())
	// TODO this creates a new client for each test instead of on for the whole test suite,
	//      as this test fails due to the status being COMPLETED instead of CANCELLED
}

func (s *QueryIntegrationSuite) TestValidateQuery() {
	ctx := testCtx()
	rc, _ := vcrClient(s.T())

	_, err := rc.ValidateQuery(ctx, "SELECT 1")
	s.Require().NoError(err)
}
