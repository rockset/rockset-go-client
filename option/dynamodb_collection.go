package option

import (
	"github.com/rockset/rockset-go-client/openapi"
)

type DynamoDBCollection func(o *openapi.CreateCollectionRequest)

func WithDynamoDBCollectionEventTimeInfo(field string) DynamoDBCollection {
	return func(o *openapi.CreateCollectionRequest) {
		o.EventTimeInfo = &openapi.EventTimeInfo{
			Field:    field,
			Format:   nil,
			TimeZone: nil,
		}
	}
}

/*
	createParams.FieldMappings
	createParams.FieldSchemas
	createParams.InvertedIndexGroupEncodingOptions
*/
