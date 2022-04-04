package rockset_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func TestHA_Integration(t *testing.T) {
	skipUnlessIntegrationTest(t)
	const apikeyName = "ROCKSET_APIKEY_USE1A1" //nolint
	apikey := os.Getenv(apikeyName)
	if apikey == "" {
		t.Skipf("skipping test as %s is not set", apikeyName)
	}

	ctx := testCtx()

	use1a1, err := rockset.NewClient(rockset.WithAPIServer("https://api.use1a1.rockset.com"),
		rockset.WithAPIKey(apikey))
	require.NoError(t, err)

	rs2, err := rockset.NewClient()
	require.NoError(t, err)

	ha := rockset.NewHA(use1a1, rs2)
	res, errs := ha.Query(ctx, "SELECT * FROM commons._events LIMIT 10")
	assert.Len(t, errs, 0)

	assert.Equal(t, "commons._events", (res.Collections)[0])
}

func (s *HaSuite) TestHA_OK_FirstFastest() {
	f0 := newFakeQuerier("0", time.Millisecond, nil)
	f1 := newFakeQuerier("1", 2*time.Millisecond, nil)

	ha := rockset.NewHA(f0, f1)

	res, errs := ha.Query(s.ctx, "SELECT 1")
	s.Nil(errs, errs)
	// require.Nil(t, errs)
	s.Equal("0", *res.QueryId)
	// assert.Equal(t, "0", *res.QueryId)
}

func (s *HaSuite) TestHA_OK_SecondFastest() {
	f0 := newFakeQuerier("0", 2*time.Millisecond, nil)
	f1 := newFakeQuerier("1", time.Millisecond, nil)

	ha := rockset.NewHA(f0, f1)

	res, errs := ha.Query(s.ctx, "SELECT 1")
	s.Nil(errs, errs)
	s.Equal("1", *res.QueryId, "Query Id %d != 1", *res.QueryId)
}

func (s *HaSuite) TestHA_OK_FirstFails() {
	ctx := testCtx()

	f0 := newFakeQuerier("0", time.Millisecond, errors.New("failed"))
	f1 := newFakeQuerier("1", 2*time.Millisecond, nil)

	ha := rockset.NewHA(f0, f1)

	res, errs := ha.Query(ctx, "SELECT 1")
	s.Nil(errs, errs)
	s.Equal("1", *res.QueryId, "Response Query Id %d != 1", *res.QueryId)
	s.Assert().Equal("1", *res.QueryId)
	// assert.Equal(t, "1", *res.QueryId)
}

func (s *HaSuite) TestHA_Fail_BothFail() {
	ctx := testCtx()

	f0 := newFakeQuerier("0", time.Millisecond, errors.New("fail0"))
	f1 := newFakeQuerier("1", 2*time.Millisecond, errors.New("fail1"))

	ha := rockset.NewHA(f0, f1)

	_, errs := ha.Query(ctx, "SELECT 1")
	s.Len(errs, 2)
	s.Equal("fail0", errs[0].Error())
	s.Equal("fail1", errs[1].Error())
}

func (s *HaSuite) TestHA_Fail_ContextCancelled() {
	ctx := testCtx()

	f0 := newFakeQuerier("0", 10*time.Millisecond, nil)
	f1 := newFakeQuerier("1", 10*time.Millisecond, nil)

	ha := rockset.NewHA(f0, f1)

	c, cancel := context.WithTimeout(ctx, time.Millisecond)
	defer cancel()

	_, errs := ha.Query(c, "SELECT 1")
	s.Len(errs, 1)
}

type fakeQuerier struct {
	delay    time.Duration
	err      error
	response openapi.QueryResponse
}

func newFakeQuerier(qid string, delay time.Duration, err error) *fakeQuerier {
	return &fakeQuerier{
		delay: delay,
		err:   err,
		response: openapi.QueryResponse{
			QueryId: &qid,
		},
	}
}

func (f *fakeQuerier) Query(ctx context.Context, query string,
	options ...option.QueryOption) (openapi.QueryResponse, error) {
	select {
	case <-time.After(f.delay):
		if f.err != nil {
			return openapi.QueryResponse{}, f.err
		}
		return f.response, nil
	case <-ctx.Done():
		return openapi.QueryResponse{}, ctx.Err()
	}
}

type HaSuite struct {
	suite.Suite
	ctx context.Context
}

func TestHaSuite(t *testing.T) {
	s := new(HaSuite)
	suite.Run(t, s)
}

func (s *HaSuite) SetupAllSuite() {
	s.ctx = testCtx()
}
