package wait_test

import (
	"net/http"
	"time"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/retry"
	"github.com/rockset/rockset-go-client/wait/fake"
)

var NotFoundErr = rockerr.Error{
	ErrorModel: openapi.NewErrorModel(),
	StatusCode: http.StatusNotFound,
}

// return a fake Rockset client with an ExponentialRetry that doesn't back off
func fakeRocksetClient() fake.FakeRocksetter {
	return fake.FakeRocksetter{
		RetryWithCheckStub: retry.Exponential{
			MaxBackoff:   time.Millisecond,
			WaitInterval: time.Millisecond,
		}.RetryWithCheck,
	}
}
