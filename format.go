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
	ColumnTypeString
	ColumnTypeTime
	ColumnTypeDate
	ColumnTypeDatetime
	ColumnTypeTimestamp
	ColumnTypeBool
	ColumnTypeInt
)

// WithCSVFormat is used by the create collection calls, to set the format to CSV.
// The columnNames and columnTypes must be of equal length, and it takes a list of optional option.CSV options.
//
//   WithCSVFormat(
//     []string{"foo", "bar"},
//     []ColumnType{ColumnTypeBoolean, ColumnTypeString},
//     option.WithSeparator(";")
//   )
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

// WithJSONFormat sets the format to JSON.
func WithJSONFormat() Format {
	return func(f *openapi.FormatParams) {
		f.Json = openapi.PtrBool(true)
	}
}

// WithXMLFormat sets the format XML.
func WithXMLFormat(xml openapi.XmlParams) Format {
	// TODO it looks lie all xml fields are optional. Add options for each of them.
	return func(f *openapi.FormatParams) {
		f.Xml = &xml
	}
}

// KafkaFormat is the definition of the Kafka format
type KafkaFormat string

// String returns the string representation of the Kafka format
func (f KafkaFormat) String() string {
	return string(f)
}

const (
	// KafkaFormatJSON is the JSON format for Kafka
	KafkaFormatJSON KafkaFormat = "JSON"
	// KafkaFormatAVRO is the AVRO format for Kafka
	KafkaFormatAVRO KafkaFormat = "AVRO"
)
