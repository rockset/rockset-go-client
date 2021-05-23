package rockset_test

import (
	"testing"
	"time"

	"github.com/rockset/rockset-go-client"
	"github.com/stretchr/testify/assert"
)

func TestExponentialRetry_Retry(t *testing.T) {
	ctx := testCtx()

	var i int
	err := rockset.ExponentialRetry{
		MaxBackoff:   time.Second,
		WaitInterval: time.Millisecond,
	}.Retry(ctx, func() (bool, error) {
		i++
		if i <= 5 {
			return true, nil
		}

		return false, nil
	})

	assert.NoError(t, err)
}
