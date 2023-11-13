package option_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/option"
)

func TestListCollectionOptions(t *testing.T) {
	options := []option.ListCollectionOption{
		option.WithWorkspace("foo"),
	}

	var opts option.ListCollectionOptions
	for _, o := range options {
		o(&opts)
	}

	assert.Equal(t, "foo", *opts.Workspace)
}
