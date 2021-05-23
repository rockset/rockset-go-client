package option

import (
	"time"

	"github.com/rockset/rockset-go-client/openapi"
)

type CollectionOption func(o *openapi.CreateCollectionRequest)

// WithCollectionRetention sets the retention in seconds for documents.
func WithCollectionRetention(d time.Duration) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		s := int64(d.Seconds())
		o.RetentionSecs = &s
	}
}

type EventTimeInfoFormat string

const (
	MillisecondsSinceEpoch EventTimeInfoFormat = "milliseconds_since_epoch"
	SecondsSinceEpoch      EventTimeInfoFormat = "seconds_since_epoch"
)

// WithCollectionClusteringKey adds a clustering key. Can be specified multiple times.
func WithCollectionClusteringKey(fieldName, fieldType string, keys []string) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		if o.ClusteringKey == nil {
			o.ClusteringKey = &[]openapi.FieldPartition{}
		}

		*o.ClusteringKey = append(*o.ClusteringKey, openapi.FieldPartition{
			FieldName: &fieldName,
			Type:      &fieldType,
			Keys:      &keys,
		})
	}
}

type FieldOption func(*openapi.FieldOptions)

type IndexMode string

func (i IndexMode) String() string { return string(i) }

const (
	IndexModeIndex   IndexMode = "index"
	IndexModeNoIndex IndexMode = "no_index"
)

func WithIndexMode(mode IndexMode) FieldOption {
	m := mode.String()
	return func(o *openapi.FieldOptions) {
		o.IndexMode = &m
	}
}

type RangeIndexMode string

func (r RangeIndexMode) String() string { return string(r) }

const (
	RangeIndexModeV1Index RangeIndexMode = "v1_index"
	RangeIndexModeNoIndex RangeIndexMode = "no_index"
)

func WithRangeIndexMode(mode RangeIndexMode) FieldOption {
	m := mode.String()
	return func(o *openapi.FieldOptions) {
		o.RangeIndexMode = &m
	}
}

type TypeIndexMode string

func (t TypeIndexMode) String() string { return string(t) }

const (
	TypeIndexModeIndex   TypeIndexMode = "index"
	TypeIndexModeNoIndex TypeIndexMode = "no_index"
)

func WithTypeIndexMode(mode TypeIndexMode) FieldOption {
	m := mode.String()
	return func(o *openapi.FieldOptions) {
		o.TypeIndexMode = &m
	}
}

type ColumnIndexMode string

func (c ColumnIndexMode) String() string { return string(c) }

const (
	ColumnIndexModeStore   ColumnIndexMode = "store"
	ColumnIndexModeNoStore ColumnIndexMode = "no_store"
)

func WithColumnIndexMode(mode ColumnIndexMode) FieldOption {
	// store or no_store
	m := mode.String()
	return func(o *openapi.FieldOptions) {
		o.ColumnIndexMode = &m
	}
}

func WithCollectionFieldSchema(fieldName string, options ...FieldOption) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		opts := openapi.FieldOptions{}
		for _, o := range options {
			o(&opts)
		}

		s := openapi.FieldSchema{
			FieldName:    &fieldName,
			FieldOptions: &opts,
		}

		if o.FieldSchemas == nil {
			o.FieldSchemas = &[]openapi.FieldSchema{}
		}
		*o.FieldSchemas = append(*o.FieldSchemas, s)
	}
}

type FieldMissingAction string

func (f FieldMissingAction) String() string { return string(f) }

const (
	FieldMissingSkip FieldMissingAction = "SKIP"
	FieldMissingPass FieldMissingAction = "PASS"
)

type InputFieldFn func(field *openapi.InputField)

func InputField(fieldName string, ifMissing FieldMissingAction, drop bool, parameterName string) InputFieldFn {
	missing := ifMissing.String()
	return func(field *openapi.InputField) {
		field.FieldName = &fieldName
		field.IfMissing = &missing
		field.IsDrop = &drop
		field.Param = &parameterName
	}
}

type OnError string

func (e OnError) String() string { return string(e) }

const (
	OnErrorSkip OnError = "SKIP"
	OnErrorFail OnError = "FAIL"
)

type OutputFieldFn func(field *openapi.OutputField)

func OutputField(fieldName string, sql string, onError OnError) OutputFieldFn {
	return func(field *openapi.OutputField) {
		e := onError.String()
		field.FieldName = &fieldName
		field.Value = &openapi.SqlExpression{Sql: &sql}
		field.OnError = &e
	}
}

// WithCollectionFieldMapping adds a field mapping to the collection.
// If dropAll is true, the input and output fields are not set.
func WithCollectionFieldMapping(name string, dropAll bool, outputField OutputFieldFn,
	inputFields ...InputFieldFn) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		mapping := openapi.FieldMappingV2{
			Name: &name,
		}

		if dropAll {
			mapping.IsDropAllFields = &dropAll
		} else {
			out := openapi.OutputField{}
			outputField(&out)
			mapping.OutputField = &out

			inputs := make([]openapi.InputField, len(inputFields))
			for i, fn := range inputFields {
				var in openapi.InputField
				fn(&in)
				inputs[i] = in
			}
			mapping.InputFields = &inputs
		}

		if o.FieldMappings == nil {
			o.FieldMappings = &[]openapi.FieldMappingV2{}
		}
		*o.FieldMappings = append(*o.FieldMappings, mapping)
	}
}

/*
	createParams.FieldSchemas
	createParams.InvertedIndexGroupEncodingOptions
*/
