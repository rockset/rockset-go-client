package ha_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/ha"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (s *HASuite) TestHA_Integration() {
	test.SkipUnlessIntegrationTest(s.T())
	const apikeyName = "ROCKSET_APIKEY_USE1A1" //nolint
	apikey := os.Getenv(apikeyName)
	if apikey == "" {
		s.T().Skipf("skipping test as %s is not set", apikeyName)
	}

	ctx := test.Context()

	use1a1, err := rockset.NewClient(rockset.WithAPIServer("https://api.use1a1.rockset.com"),
		rockset.WithAPIKey(apikey))
	s.NoError(err)

	rs2 := test.Client(s.T())

	client := ha.New(use1a1, rs2)
	res, errs := client.Query(ctx, "SELECT * FROM commons._events LIMIT 10")
	s.Len(errs, 0)

	s.Equal("commons._events", (res.Collections)[0])
}

func (s *HASuite) TestContextFail() {
	a := createMock("1", time.Second, nil)
	b := createMock("2", time.Second, nil)
	client := ha.New(a, b)

	ctx := test.Context()
	c, cancel := context.WithTimeout(ctx, time.Millisecond)
	defer cancel()

	_, err := client.Query(c, "SELECT 1")
	s.Len(err, 1)
}

func (s *HASuite) TestHA() {
	for _, t := range s.tests {
		mocks := createMocks(t.queries)
		client := ha.New(mocks...)
		ctx := test.Context()
		s.Run(t.name, func() {
			res, errs := client.Query(ctx, "SELECT 1")
			s.Equal(len(t.exErrors), len(errs))
			if res.QueryId == nil {
				s.Equal("", t.exID)
				return
			}
			s.Equal(*res.QueryId, t.exID)
		})
	}
}

func createMock(queryID string, delay time.Duration, err error) ha.Querier {
	var mQ mockQuerier
	call := mQ.On("Query", mock.Anything, mock.Anything, mock.Anything)
	call.After(delay)
	call.Return(openapi.QueryResponse{
		QueryId: &queryID,
	}, err)

	return &mQ
}

func createMocks(queries []query) []ha.Querier {
	var mocks []ha.Querier
	for _, q := range queries {
		mocks = append(mocks, createMock(q.id, q.delay, q.err))
	}
	return mocks
}

// TODO replace with counterfeiter
type mockQuerier struct {
	mock.Mock
}

type queryTest struct {
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

func (m *mockQuerier) Query(ctx context.Context, query string,
	options ...option.QueryOption) (openapi.QueryResponse, error) {
	args := m.Called(ctx, query, options)
	return args.Get(0).(openapi.QueryResponse), args.Error(1)
}

type HASuite struct {
	suite.Suite
	tests []queryTest
}

func TestHaSuite(t *testing.T) {
	s := new(HASuite)
	suite.Run(t, s)
}

func (s *HASuite) SetupTest() {

}

func (s *HASuite) SetupSuite() {
	tests := []queryTest{
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
		{
			name:     "BothFail",
			exID:     "",
			exErrors: []error{errors.New("fail"), errors.New("fail")},
			queries: []query{
				{id: "1", delay: time.Millisecond, err: errors.New("fail")},
				{id: "2", delay: time.Millisecond * 5, err: errors.New("fail")},
			},
		},
	}

	s.tests = tests
}
