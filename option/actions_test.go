package option_test

import (
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGlobalAction(t *testing.T) {
	a := option.GetGlobalAction("ALL_GLOBAL_ACTIONS")
	assert.Equal(t, option.AllGlobalActions, a)
	assert.Equal(t, "ALL_GLOBAL_ACTIONS", a.String())
}
