package retry_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/retry"
)

func fakeError(code int) error {
	t := strings.Join(strings.Fields(http.StatusText(code)), "")
	return rockerr.Error{
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
	suite.Run(t, &ExponentialRetrySuite{})
}

func (s *ExponentialRetrySuite) TestDefaultRetry() {
	ctx := context.TODO()
	var count int

	exp := retry.NewExponential()
	exp.JitterFraction = 0.04
	exp.MaxBackoff = time.Second
	exp.WaitInterval = time.Millisecond

	err := exp.Retry(ctx, func() error {
		count++
		if count > 2 {
			return nil
		}
		return fakeError(http.StatusTooManyRequests)
	})

	s.Assert().NoError(err)
}

func (s *ExponentialRetrySuite) TestDefaultRetryWithFailure() {
	ctx := context.TODO()

	err := retry.Exponential{
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
	ctx := context.TODO()
	var count int

	err := retry.Exponential{
		MaxBackoff:          time.Second,
		WaitInterval:        time.Millisecond,
		RetryableErrorCheck: retry.DefaultRetryableErrorCheck,
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
	ctx := context.TODO()
	var count int

	err := retry.Exponential{
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

func (s *ExponentialRetrySuite) TestExponentialRetry_RetryFnCancelled() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond)

	err := retry.Exponential{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
		RetryableErrorCheck: func(_ error) bool {
			// always retry, forever
			return true
		},
	}.Retry(ctx, func() error {
		cancel()
		return errors.New("error")
	})

	s.Assert().ErrorIs(err, context.Canceled)
}

func (s *ExponentialRetrySuite) TestExponentialRetry_RetryWithCheck() {
	ctx := context.TODO()

	var i int
	err := retry.Exponential{
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

func (s *ExponentialRetrySuite) TestExponentialRetry_RetryWithCheckError() {
	ctx := context.TODO()

	err := retry.Exponential{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
	}.RetryWithCheck(ctx, func() (bool, error) {
		return false, errors.New("error")
	})

	s.Assert().Error(err)
}

func (s *ExponentialRetrySuite) TestExponentialRetry_RetryWithCheckCancelled() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond)

	err := retry.Exponential{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
	}.RetryWithCheck(ctx, func() (bool, error) {
		cancel()
		time.Sleep(time.Millisecond)
		return true, nil
	})

	s.Assert().ErrorIs(err, context.Canceled)
}
