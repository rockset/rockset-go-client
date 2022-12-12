package rockset

import (
	"context"
	"fmt"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

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

	rc := resourceIsAvailable(ctx, func(ctx context.Context) error {
		defer func() { counter++ }()

		switch counter {
		case 0:
			return Error{ErrorModel: &openapi.ErrorModel{}, StatusCode: http.StatusNotFound, Cause: fmt.Errorf("resource not present")}
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

	rc := resourceIsGone(ctx, func(ctx context.Context) error {
		defer func() { counter++ }()

		switch counter {
		case 0:
			return nil
		case 1:
			return Error{ErrorModel: &openapi.ErrorModel{}, StatusCode: http.StatusNotFound, Cause: fmt.Errorf("resource not present")}
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

func (s *WaitTestSuite) TestResourceHasState() {
	ctx := context.TODO()
	var counter int

	rc := resourceHasState(ctx, []string{"foo", "bar"}, func(ctx context.Context) (string, error) {
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
