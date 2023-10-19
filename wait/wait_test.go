package wait_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/retry"
	"github.com/rockset/rockset-go-client/wait"
	"github.com/rockset/rockset-go-client/wait/fake"
)

var ErrNotFound = rockerr.Error{
	ErrorModel: openapi.NewErrorModel(),
	StatusCode: http.StatusNotFound,
}

func stringPtr[T fmt.Stringer](status T) *string {
	s := status.String()
	return &s
}

var emptyString = ""

// return a fake Rockset client with an ExponentialRetry that doesn't back off
func fakeRocksetClient() fake.FakeResourceGetter {
	return fake.FakeResourceGetter{
		RetryWithCheckStub: retry.Exponential{
			MaxBackoff:   time.Millisecond,
			WaitInterval: time.Millisecond,
		}.RetryWithCheck,
	}
}

type WaitTestSuite struct {
	suite.Suite
}

func TestWaitSuite(t *testing.T) {
	s := WaitTestSuite{}
	suite.Run(t, &s)
}

func (s *WaitTestSuite) TestResourceIsAvailable() {
	ctx := context.TODO()
	var counter int

	rc := wait.ResourceIsAvailable(ctx, func(ctx context.Context) error {
		defer func() { counter++ }()

		switch counter {
		case 0:
			return rockerr.Error{
				ErrorModel: &openapi.ErrorModel{},
				StatusCode: http.StatusNotFound,
				Cause:      fmt.Errorf("resource not present"),
			}
		case 1:
			return nil
		default:
			return fmt.Errorf("fail")
		}
	})

	retry, err := rc()
	assert.True(s.T(), retry)
	assert.NoError(s.T(), err)

	retry, err = rc()
	assert.False(s.T(), retry)
	assert.NoError(s.T(), err)

	retry, err = rc()
	assert.False(s.T(), retry)
	assert.Error(s.T(), err)
}

func (s *WaitTestSuite) TestResourceIsGone() {
	ctx := context.TODO()
	var counter int

	rc := wait.ResourceIsGone(ctx, func(ctx context.Context) error {
		defer func() { counter++ }()

		switch counter {
		case 0:
			return nil
		case 1:
			return rockerr.Error{
				ErrorModel: &openapi.ErrorModel{},
				StatusCode: http.StatusNotFound,
				Cause:      fmt.Errorf("resource not present"),
			}
		default:
			return fmt.Errorf("fail")
		}
	})

	r, err := rc()
	assert.True(s.T(), r)
	assert.NoError(s.T(), err)

	r, err = rc()
	assert.False(s.T(), r)
	assert.NoError(s.T(), err)

	r, err = rc()
	assert.False(s.T(), r)
	assert.Error(s.T(), err)
}

type state string

func (s state) String() string { return string(s) }
func (s *WaitTestSuite) TestResourceHasState() {
	ctx := context.TODO()
	var counter int

	rc := wait.ResourceHasState(ctx, []state{"foo", "bar"}, nil, func(ctx context.Context) (state, error) {
		defer func() { counter++ }()

		switch counter {
		case 0:
			return "baz", nil
		case 1:
			return "bar", nil
		default:
			return "", fmt.Errorf("fail")
		}
	})

	r, err := rc()
	assert.True(s.T(), r)
	assert.NoError(s.T(), err)

	r, err = rc()
	assert.False(s.T(), r)
	assert.NoError(s.T(), err)

	r, err = rc()
	assert.False(s.T(), r)
	assert.Error(s.T(), err)
}

func (s *WaitTestSuite) TestResourceHasState_badState() {
	ctx := context.TODO()
	var counter int

	rc := wait.ResourceHasState(ctx, []state{"foo", "bar"}, []state{"err"}, func(ctx context.Context) (state, error) {
		defer func() { counter++ }()

		switch counter {
		case 0:
			return "baz", nil
		case 1:
			return "bar", nil
		default:
			return "err", nil
		}
	})

	r, err := rc()
	assert.True(s.T(), r)
	assert.NoError(s.T(), err)

	r, err = rc()
	assert.False(s.T(), r)
	assert.NoError(s.T(), err)

	r, err = rc()
	assert.False(s.T(), r)
	assert.ErrorIs(s.T(), err, rockerr.ErrBadWaitState)
}
