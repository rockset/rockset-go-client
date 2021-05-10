package rockset

import (
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

type Format func(params *openapi.FormatParams)

type ColumnType int

// String returns the string representation of the ColumnType
func (c ColumnType) String() string {
	return [...]string{
		"Unknown", "Boolean", "Integer", "Float", "String",
		"Time", "Date", "Datetime", "Timestamp", "Bool", "Int",
	}[c]
}

const (
	ColumnTypeUnknown ColumnType = iota
	ColumnTypeBoolean
	ColumnTypeInteger
	ColumnTypeFloat
	ColumnTypeSTRING
	ColumnTypeTime
	ColumnTypeDate
	ColumnTypeDatetime
	ColumnTypeTimestamp
	ColumnTypeBool
	ColumnTypeInt
)

func WithCSVFormat(columnNames []string, columnTypes []ColumnType, options ...option.CSV) Format {
	types := make([]string, len(columnTypes))
	for i, t := range columnTypes {
		types[i] = t.String()
	}
	csv := openapi.CsvParams{
		ColumnNames: &columnNames,
		ColumnTypes: &types,
	}

	for _, o := range options {
		o(&csv)
	}

	return func(f *openapi.FormatParams) {
		f.Csv = &csv
	}
}

func WithJSONFormat() Format {
	return func(f *openapi.FormatParams) {
		f.Json = openapi.PtrBool(true)
	}
}

func WithXMLFormat(xml openapi.XmlParams) Format {
	return func(f *openapi.FormatParams) {
		f.Xml = &xml
	}
}
