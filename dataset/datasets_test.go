package dataset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/dataset"
)

func TestDatasets(t *testing.T) {
	for _, ds := range dataset.All() {
		assert.NotEmpty(t, dataset.Lookup(ds), "could not find %s", ds)
	}
}
