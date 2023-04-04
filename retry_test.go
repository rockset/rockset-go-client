package rockset_test

import (
	"errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/rockset/rockset-go-client"
)

func fakeError(code int) error {
	t := strings.Join(strings.Fields(http.StatusText(code)), "")
	return rockset.Error{
		ErrorModel: &openapi.ErrorModel{
			Type: &t,
		},
		StatusCode: code,
	}
}

type ExponentialRetrySuite struct {
	suite.Suite
}

func TestExponentialRetrySuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, &ExponentialRetrySuite{})
}

func (s *ExponentialRetrySuite) TestDefaultRetry() {
	ctx := testCtx()
	var count int

	err := rockset.ExponentialRetry{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
	}.Retry(ctx, func() error {
		count++
		if count > 2 {
			return nil
		}
		return fakeError(http.StatusTooManyRequests)
	})

	s.Assert().NoError(err)
}

func (s *ExponentialRetrySuite) TestDefaultRetryWithFailure() {
	ctx := testCtx()

	err := rockset.ExponentialRetry{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
	}.Retry(ctx, func() error {
		return errors.New("other error")
	})

	s.Assert().Error(err)
}

// custom error that implements RetryableError
type retryableErr struct{ err string }

func (r retryableErr) Error() string   { return r.err }
func (r retryableErr) Retryable() bool { return true }

func (s *ExponentialRetrySuite) TestDefaultRetryFn() {
	ctx := testCtx()
	var count int

	err := rockset.ExponentialRetry{
		MaxBackoff:          time.Second,
		WaitInterval:        time.Millisecond,
		RetryableErrorCheck: rockset.DefaultRetryableErrorCheck,
	}.Retry(ctx, func() error {
		count++
		if count > 2 {
			return nil
		}
		return retryableErr{"retryable error"}
	})

	s.Assert().NoError(err)
}

func (s *ExponentialRetrySuite) TestExponentialRetry_RetryFn() {
	ctx := testCtx()
	var count int

	err := rockset.ExponentialRetry{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
		RetryableErrorCheck: func(err error) bool {
			return errors.Is(err, io.ErrUnexpectedEOF)
		},
	}.Retry(ctx, func() error {
		count++
		if count > 2 {
			return nil
		}
		return io.ErrUnexpectedEOF
	})

	s.Assert().NoError(err)
}

func (s *ExponentialRetrySuite) TestExponentialRetry_RetryWithCheck() {
	ctx := testCtx()

	var i int
	err := rockset.ExponentialRetry{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
	}.RetryWithCheck(ctx, func() (bool, error) {
		i++
		if i <= 5 {
			return true, nil
		}

		return false, nil
	})

	s.Assert().NoError(err)
}
