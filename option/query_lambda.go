package option

import "github.com/rockset/rockset-go-client/openapi"

type QueryLambdaState string

func (q QueryLambdaState) String() string { return string(q) }

const (
	QueryLambdaActive     QueryLambdaState = "ACTIVE"
	QueryLambdaInvalidSQL QueryLambdaState = "INVALID_SQL"
)

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

// WithQueryLambdaRowLimit sets the maximum number of rows to retrieve.
func WithQueryLambdaRowLimit(limit int32) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.DefaultRowLimit = &limit
	}
}

// WithQueryLambdaParameter sets a query lambda parameter.
func WithQueryLambdaParameter(name, valueType, value string) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.Parameters = append(o.Parameters, openapi.QueryParameter{
			Name:  name,
			Type:  valueType,
			Value: value,
		})
	}
}

// WithQueryLambdaRequest sets the query lambda request.
func WithQueryLambdaRequest(r openapi.ExecuteQueryLambdaRequest) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.ExecuteQueryLambdaRequest = r
	}
}

type ListQueryLambdaOptions struct {
	Workspace *string
}

type ListQueryLambdaOption func(request *ListQueryLambdaOptions)

func WithQueryLambdaWorkspace(name string) ListQueryLambdaOption {
	return func(o *ListQueryLambdaOptions) {
		o.Workspace = &name
	}
}

type CreateQueryLambdaOptions struct {
	Description     *string
	QueryParameters []openapi.QueryParameter
}

type CreateQueryLambdaOption func(request *CreateQueryLambdaOptions)

func WithQueryLambdaDescription(desc string) CreateQueryLambdaOption {
	return func(o *CreateQueryLambdaOptions) {
		o.Description = &desc
	}
}

func WithDefaultParameter(name, paramType, value string) CreateQueryLambdaOption {
	return func(o *CreateQueryLambdaOptions) {
		o.QueryParameters = append(o.QueryParameters, openapi.QueryParameter{
			Name:  name,
			Type:  paramType,
			Value: value,
		})
	}
}
