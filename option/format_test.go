package option_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func TestFormat(t *testing.T) {
	options := []option.Format{
		option.WithCSVFormat([]string{"foo"}, []option.ColumnType{option.ColumnTypeBool},
			option.WithEscapeChar("e")),
		option.WithJSONFormat(),
		option.WithXMLFormat(openapi.XmlParams{
			Encoding: openapi.PtrString("foo"),
		}),
	}

	var opts openapi.FormatParams
	for _, o := range options {
		o(&opts)
	}

	assert.Equal(t, "foo", opts.Csv.ColumnNames[0])
	assert.Equal(t, option.ColumnTypeBool.String(), opts.Csv.ColumnTypes[0])
	assert.Equal(t, "e", *opts.Csv.EscapeChar)
	assert.Equal(t, "foo", *opts.Xml.Encoding)
	assert.True(t, *opts.Json)
}
