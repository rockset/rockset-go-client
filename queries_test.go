package rockset_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

// for anyone poking around in the code, rockset.sleep() only works for this test org as no sane person would want
// to add a sleep in their query
const slowQuery = `script {{{ import * as rockset from "/rockset"; export function delay(x) ` +
	`{ rockset.sleep(x); return x; } }}} select _script.delay(2000, q) from unnest([1] as q)`

type QueryIntegrationSuite struct {
	suite.Suite
	rc *rockset.RockClient
}

func TestQueryIntegration(t *testing.T) {
	s := QueryIntegrationSuite{}
	suite.Run(t, &s)
}

func (s *QueryIntegrationSuite) BeforeTest(suiteName, testName string) {
	s.rc, _ = vcrTestClient(s.T(), fmt.Sprintf("%s/%s", suiteName, testName))
}

func (s *QueryIntegrationSuite) TestQuery() {
	ctx := test.Context()

	_, err := s.rc.Query(ctx, "SELECT 1")
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestListQueries() {
	ctx := test.Context()

	_, err := s.rc.ListActiveQueries(ctx)
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestAsyncQuery() {
	ctx := test.Context()

	resp, err := s.rc.Query(ctx, slowQuery,
		option.WithAsync(),
	)
	s.Require().NoError(err)

	err = s.rc.Wait.UntilQueryDone(ctx, *resp.QueryId)
	s.Require().NoError(err)

	_, err = s.rc.GetQueryResults(ctx, *resp.QueryId)
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestAsyncQueryWithClientTimeout() {
	ctx := test.Context()

	resp, err := s.rc.Query(ctx, slowQuery,
		option.WithAsync(),
		option.WithAsyncClientTimeout(100),
		option.WithMaxInitialResults(10),
	)
	s.Require().NoError(err)

	err = s.rc.Wait.UntilQueryDone(ctx, *resp.QueryId)
	s.Require().NoError(err)

	_, err = s.rc.GetQueryResults(ctx, *resp.QueryId)
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestQueryWithTimeout() {
	ctx := test.Context()

	_, err := s.rc.Query(ctx, slowQuery,
		option.WithTimeout(100*time.Millisecond),
	)
	s.Require().Error(err)

	var re rockerr.Error
	s.Require().True(errors.As(err, &re))
	s.Assert().Equal("QUERY_TIMEOUT", *re.Type)
}

func (s *QueryIntegrationSuite) TestCancelQuery() {
	ctx := test.Context()

	resp, err := s.rc.Query(ctx, slowQuery,
		option.WithAsync(),
		option.WithAsyncClientTimeout(100),
		option.WithMaxInitialResults(10),
	)
	s.Require().NoError(err)

	info, err := s.rc.CancelQuery(ctx, *resp.QueryId)
	s.Require().NoError(err)
	s.Require().Equal("CANCELLED", info.GetStatus())
	// TODO this creates a new client for each test instead of on for the whole test suite,
	//      as this test fails due to the status being COMPLETED instead of CANCELLED
}

func (s *QueryIntegrationSuite) TestValidateQuery() {
	ctx := test.Context()

	_, err := s.rc.ValidateQuery(ctx, "SELECT 1")
	s.Require().NoError(err)
}

func (s *QueryIntegrationSuite) TestQueryWithVirtualInstance() {
	ctx := test.Context()

	vis, err := s.rc.ListVirtualInstances(ctx)
	s.Require().NoError(err)
	var mainVI string
	for _, vi := range vis {
		if vi.GetName() == "main" {
			mainVI = vi.GetId()
		}
	}
	s.Require().NotEmpty(mainVI)

	_, err = s.rc.Query(ctx, "SELECT 1", option.WithVirtualInstance(mainVI))
	s.Require().NoError(err)
}
