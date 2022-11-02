package rockset_test

import (
	"errors"
	"github.com/rockset/rockset-go-client/openapi"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client"
)

func fakeError(code int) error {
	t := strings.Join(strings.Fields(http.StatusText(code)), "")
	return rockset.Error{
		ErrorModel: &openapi.ErrorModel{
			Type: &t,
		},
		Cause: errors.New(http.StatusText(code)),
	}
}
func TestExponentialRetry_DefaultRetry(t *testing.T) {
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

	assert.NoError(t, err)
}

func TestExponentialRetry_DefaultRetryFailure(t *testing.T) {
	ctx := testCtx()

	err := rockset.ExponentialRetry{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
	}.Retry(ctx, func() error {
		return errors.New("other error")
	})

	assert.Error(t, err)
}

type retryableErr struct{ err string }

func (r retryableErr) Error() string   { return r.err }
func (r retryableErr) Retryable() bool { return true }

func TestExponentialRetry_DefaultRetryFn(t *testing.T) {
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

	assert.NoError(t, err)
}

func TestExponentialRetry_RetryFn(t *testing.T) {
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

	assert.NoError(t, err)
}

func TestExponentialRetry_RetryWithCheck(t *testing.T) {
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

	assert.NoError(t, err)
}
