package rockset_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (s *HASuite) TestHA_Integration() {
	skipUnlessIntegrationTest(t)
	const apikeyName = "ROCKSET_APIKEY_USE1A1" //nolint
	apikey := os.Getenv(apikeyName)
	if apikey == "" {
		t.Skipf("skipping test as %s is not set", apikeyName)
	}

	ctx := testCtx()

	use1a1, err := rockset.NewClient(rockset.WithAPIServer("https://api.use1a1.rockset.com"),
		rockset.WithAPIKey(apikey))
	s.NoError(err)

	rs2, err := rockset.NewClient()
	s.NoError(err)

	ha := rockset.NewHA(use1a1, rs2)
	res, errs := ha.Query(ctx, "SELECT * FROM commons._events LIMIT 10")
	s.Len(errs, 0)

	assert.Equal(t, "commons._events", (res.Collections)[0])
}

func createMock(queryID string, delay time.Duration, err error) rockset.Querier {
	var mQ mockQuerier
	call := mQ.On("Query", mock.Anything, mock.Anything, mock.Anything)
	call.After(delay)
	call.Return(openapi.QueryResponse{
		QueryId: &queryID,
	}, err)

	return &mQ
}

func createMocks(queries []query) []rockset.Querier {
	mocks := []rockset.Querier{}
	for _, q := range queries {
		mocks = append(mocks, createMock(q.id, q.delay, q.err))
	}
	return mocks
}

func (s *HASuite) TestHA() {
	for _, t := range s.tests {
		mocks := createMocks(t.queries)
		ha := rockset.NewHA(mocks...)
		ctx := testCtx()
		s.Run(t.name, func() {
			res, errs := ha.Query(ctx, "SELECT 1")
			s.Equal(len(t.exErrors), len(errs))
			s.Equal(*res.QueryId, t.exID)
		})
	}
}

type mockQuerier struct {
	mock.Mock
}

func (m *mockQuerier) Query(ctx context.Context, query string, options ...option.QueryOption) (openapi.QueryResponse, error) {

	args := m.Called(ctx, query, options)
	return args.Get(0).(openapi.QueryResponse), args.Error(1)
}

type HASuite struct {
	suite.Suite
	tests []test
}

func TestHaSuite(t *testing.T) {
	s := new(HASuite)
	suite.Run(t, s)
}

func (s *HASuite) SetupTest() {

}

func (s *HASuite) SetupSuite() {
	tests := []test{
		{
			name:     "FirstFastest",
			exID:     "1",
			exErrors: []error{},
			queries: []query{
				{id: "1", delay: time.Millisecond, err: nil},
				{id: "2", delay: time.Millisecond * 5, err: nil},
			},
		},
		{
			name:     "SecondFastest",
			exID:     "2",
			exErrors: []error{},
			queries: []query{
				{id: "1", delay: time.Millisecond * 5, err: nil},
				{id: "2", delay: time.Millisecond, err: nil},
			},
		},
		{
			name:     "FirstFail",
			exID:     "2",
			exErrors: []error{},
			queries: []query{
				{id: "1", delay: time.Millisecond, err: errors.New("fail")},
				{id: "2", delay: time.Millisecond * 5, err: nil},
			},
		},
	}

	s.tests = tests
}

type test struct {
	name     string
	queries  []query
	exID     string
	exErrors []error
}

type query struct {
	id    string
	delay time.Duration
	err   error
}
