package option_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func TestCSV(t *testing.T) {
	options := []option.CSV{
		option.WithSeparator("foo"),
		option.WithQuoteChar("foo"),
		option.WithEscapeChar("foo"),
		option.WithFirstLineAsColumnNames(),
	}

	var opts openapi.CsvParams
	for _, o := range options {
		o(&opts)
	}

	assert.Equal(t, "foo", *opts.Separator)
	assert.Equal(t, "foo", *opts.QuoteChar)
	assert.Equal(t, "foo", *opts.EscapeChar)
	assert.True(t, *opts.FirstLineAsColumnNames)
}
