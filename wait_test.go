package rockset

import (
	"context"
	"errors"
	"github.com/rockset/rockset-go-client/fakes"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/stretchr/testify/suite"
	"net/http"
	"strings"
	"testing"
	"time"
)

func fakeError(code int) error {
	s := strings.Join(strings.Fields(http.StatusText(code)), "")
	return Error{
		ErrorModel: &openapi.ErrorModel{
			Type: &s,
		},
		Cause: errors.New("fake error"),
	}
}

type WaitSuite struct {
	suite.Suite
}

func TestWaitSuite(t *testing.T) {
	s := WaitSuite{}
	suite.Run(t, &s)
}

func (s *WaitSuite) TestQueryHasStatus() {
	ctx := context.TODO()

	f := &fakes.FakeQueryInfoGetter{}
	for i, status := range []QueryState{QueryQueued, QueryRunning, QueryCompleted} {
		v := string(status)
		f.GetQueryInfoReturnsOnCall(i, openapi.QueryInfo{Status: &v}, nil)
	}

	fn := queryHasStatus(ctx, f, "id", []QueryState{QueryCompleted})
	err := ExponentialRetry{MaxBackoff: time.Millisecond, WaitInterval: time.Microsecond}.RetryWithCheck(ctx, fn)
	s.Require().NoError(err)
	s.Assert().Equal(3, f.GetQueryInfoCallCount())
}

func (s *WaitSuite) TestQueryHasStatus_error() {
	ctx := context.TODO()

	f := &fakes.FakeQueryInfoGetter{}
	f.GetQueryInfoReturnsOnCall(0, openapi.QueryInfo{Status: openapi.PtrString(string(QueryQueued))}, nil)
	f.GetQueryInfoReturnsOnCall(1, openapi.QueryInfo{}, fakeError(http.StatusBadRequest))

	fn := queryHasStatus(ctx, f, "id", []QueryState{QueryCompleted})
	err := ExponentialRetry{MaxBackoff: time.Millisecond, WaitInterval: time.Microsecond}.RetryWithCheck(ctx, fn)
	s.Assert().Error(err)
	s.Assert().Equal(2, f.GetQueryInfoCallCount())
}

func (s *WaitSuite) TestCollectionHasState() {
	ctx := context.TODO()

	f := &fakes.FakeCollectionGetter{}
	for i, status := range []string{"foo", "bar", "foobar"} {
		s := status
		f.GetCollectionReturnsOnCall(i, openapi.Collection{Status: &s}, nil)
	}

	fn := collectionHasState(ctx, f, "ws", "c", "foobar")
	err := ExponentialRetry{MaxBackoff: time.Millisecond, WaitInterval: time.Microsecond}.RetryWithCheck(ctx, fn)
	s.Require().NoError(err)
	s.Assert().Equal(3, f.GetCollectionCallCount())
}

func (s *WaitSuite) TestCollectionHasNewDocs() {
	ctx := context.TODO()

	f := &fakes.FakeCollectionGetter{}
	for i, count := range []int64{10, 20, 20, 30, 40} {
		c := count
		f.GetCollectionReturnsOnCall(i, openapi.Collection{Stats: &openapi.CollectionStats{DocCount: &c}}, nil)
	}

	d := docWaiter{getter: f}

	fn := d.collectionHasNewDocs(ctx, "ws", "c", 30)
	err := ExponentialRetry{MaxBackoff: time.Millisecond, WaitInterval: time.Microsecond}.RetryWithCheck(ctx, fn)
	s.Require().NoError(err)
	s.Assert().Equal(5, f.GetCollectionCallCount())
}
