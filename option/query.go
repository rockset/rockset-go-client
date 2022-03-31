package option

import "github.com/rockset/rockset-go-client/openapi"

type QueryOption func(request *openapi.QueryRequestSql)

func WithWarnings() QueryOption {
	return func(o *openapi.QueryRequestSql) {
		o.GenerateWarnings = openapi.PtrBool(true)
	}
}

func WithRowLimit(limit int32) QueryOption {
	return func(o *openapi.QueryRequestSql) {
		o.DefaultRowLimit = &limit
	}
}

func WithParameter(name, valueType, value string) QueryOption {
	return func(o *openapi.QueryRequestSql) {
		o.Parameters = append(o.Parameters, openapi.QueryParameter{
			Name:  name,
			Type:  valueType,
			Value: value,
		})
	}
}
