package rockset_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client"
)

func TestExponentialRetry_Retry(t *testing.T) {
	ctx := testCtx()

	err := rockset.ExponentialRetry{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
	}.Retry(ctx, func() error {
		// TODO this should return a openapi.GenericOpenAPIError for the test to be useful
		//   but it isn't possible. Need to see how to test the Retry() method. Possibly it can be solved
		//   by merging it with RetryWithCheck()
		return errors.New("test error")
	})

	assert.Error(t, err)
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
