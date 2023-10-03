package rockset_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/openapi"
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
	ctx := test.Context()
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
	ctx := test.Context()

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
	ctx := test.Context()
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
	ctx := test.Context()
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
	ctx := test.Context()

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
