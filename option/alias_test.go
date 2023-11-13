package option_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/option"
)

func TestAliasOptions(t *testing.T) {
	options := []option.AliasOption{
		option.WithAliasDescription("foo"),
	}

	var opts option.AliasOptions
	for _, o := range options {
		o(&opts)
	}

	assert.Equal(t, "foo", *opts.Description)
}

func TestListAliasesOptions(t *testing.T) {
	options := []option.ListAliasesOption{
		option.WithAliasWorkspace("foo"),
	}

	var opts option.ListAliasesOptions
	for _, o := range options {
		o(&opts)
	}

	assert.Equal(t, "foo", opts.Workspace)
}
