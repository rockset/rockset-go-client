package option

import "github.com/rockset/rockset-go-client/openapi"

type ExecuteQueryLambdaRequest struct {
	openapi.ExecuteQueryLambdaRequest
	Tag     string
	Version string
}

type QueryLambdaOption func(request *ExecuteQueryLambdaRequest)

func WithTag(tag string) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.Tag = tag
	}
}

func WithVersion(version string) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.Version = version
	}
}

func WithRowLimit2(limit int32) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.DefaultRowLimit = &limit
	}
}

func WithWarnings2() QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.GenerateWarnings = openapi.PtrBool(true)
	}
}

func WithParameter2(name, valueType, value string) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		*o.Parameters = append(*o.Parameters, openapi.QueryParameter{
			Name:  name,
			Type:  valueType,
			Value: value,
		})
	}
}
