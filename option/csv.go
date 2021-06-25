package option

import "github.com/rockset/rockset-go-client/openapi"

type CSV func(*openapi.CsvParams)

func WithEncoding(encoding string) CSV {
	return func(o *openapi.CsvParams) {
		o.Encoding = &encoding
	}
}

// WithSeparator sets the CSV separator. Defaults to ,
func WithSeparator(separator string) CSV {
	return func(o *openapi.CsvParams) {
		o.Separator = &separator
	}
}

// WithQuoteChar sets the CSV quote character. Defaults to "
func WithQuoteChar(quote string) CSV {
	return func(o *openapi.CsvParams) {
		o.QuoteChar = &quote
	}
}

// WithEscapeChar sets the CSV escape character. Defaults to \
func WithEscapeChar(escape string) CSV {
	return func(o *openapi.CsvParams) {
		o.EscapeChar = &escape
	}
}

// WithFirstLineAsColumnNames enables the use of the first row as the column names
func WithFirstLineAsColumnNames() CSV {
	return func(o *openapi.CsvParams) {
		o.FirstLineAsColumnNames = openapi.PtrBool(true)
	}
}
