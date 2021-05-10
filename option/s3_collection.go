package option

import (
	"github.com/rockset/rockset-go-client/openapi"
	"time"
)

type S3CollectionOption func(o *openapi.CreateCollectionRequest)

func WithS3CollectionDescription(desc string) S3CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.Description = &desc
	}
}

// WithS3CollectionRetention sets the retention in seconds for documents.
func WithS3CollectionRetention(d time.Duration) S3CollectionOption {
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

// WithS3CollectionEventTimeInfo configures event data.
// fieldName is name of the field containing event time.
// format of time field, can be one of: milliseconds_since_epoch, seconds_since_epoch
// ts is the default time zone, in standard IANA format
func WithS3CollectionEventTimeInfo(fieldName, tz string, format EventTimeInfoFormat) S3CollectionOption {
	// TODO: should tz be validated?
	f := string(format)
	return func(o *openapi.CreateCollectionRequest) {
		o.EventTimeInfo = &openapi.EventTimeInfo{
			Field:    fieldName,
			Format:   &f,
			TimeZone: &tz,
		}
	}
}

// WithS3CollectionClusteringKey adds a clustering key. Can be specified multiple times.
func WithS3CollectionClusteringKey(fieldName, fieldType string, keys []string) S3CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		*o.ClusteringKey = append(*o.ClusteringKey, openapi.FieldPartition{
			FieldName: &fieldName,
			Type:      &fieldType,
			Keys:      &keys,
		})
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

// WithS3CollectionFieldMapping adds a field mapping to the collection.
// If dropAll is true, the input and output fields are not set.
func WithS3CollectionFieldMapping(name string, dropAll bool, outputField OutputFieldFn,
	inputFields ...InputFieldFn) S3CollectionOption {

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

		*o.FieldMappings = append(*o.FieldMappings, mapping)
	}
}

/*
	createParams.FieldSchemas
	createParams.InvertedIndexGroupEncodingOptions
*/
