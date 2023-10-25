package option_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/option"
)

func TestGetGlobalAction(t *testing.T) {
	a := option.GetGlobalAction("ALL_GLOBAL_ACTIONS")
	assert.Equal(t, option.AllGlobalActions, a)
	assert.Equal(t, "ALL_GLOBAL_ACTIONS", a.String())
}
