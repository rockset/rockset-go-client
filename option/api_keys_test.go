package option_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/option"
)

func TestAPIKeyOptions(t *testing.T) {
	options := []option.APIKeyOption{
		option.ForUser("foo"),
		option.State(option.KeySuspended),
	}

	var opts option.APIKeyOptions
	for _, o := range options {
		o(&opts)
	}

	assert.Equal(t, "foo", *opts.User)
	assert.Equal(t, option.KeySuspended, *opts.State)
}

func TestAPIKeyRoleOptions(t *testing.T) {
	options := []option.APIKeyRoleOption{
		option.WithRole("foo"),
	}

	var opts option.APIKeyRoleOptions
	for _, o := range options {
		o(&opts)
	}

	assert.Equal(t, "foo", *opts.Role)
}
