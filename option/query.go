package option

import "github.com/rockset/rockset-go-client/openapi"

type QueryOption func(request *openapi.QueryRequest)

func WithWarnings() QueryOption {
	return func(o *openapi.QueryRequest) {
		o.Sql.GenerateWarnings = openapi.PtrBool(true)
	}
}

func WithRowLimit(limit int32) QueryOption {
	return func(o *openapi.QueryRequest) {
		o.Sql.DefaultRowLimit = &limit
	}
}

func WithParameter(name, valueType, value string) QueryOption {
	return func(o *openapi.QueryRequest) {
		o.Sql.Parameters = append(o.Sql.Parameters, openapi.QueryParameter{
			Name:  name,
			Type:  valueType,
			Value: value,
		})
	}
}

// WithAsyncClientTimeout is maximum amount of time that the client is willing to wait for the query to complete.
// If the query is not complete by this timeout, a response will be returned with a query_id that can be used to
// check the status of the query and retrieve results once the query has completed.
func WithAsyncClientTimeout(timeout int64) QueryOption {
	return func(o *openapi.QueryRequest) {
		if o.AsyncOptions == nil {
			o.AsyncOptions = &openapi.AsyncQueryOptions{}
		}
		o.AsyncOptions.ClientTimeoutMs = &timeout
	}
}

// WithAsyncTimeout maximum amount of time that Rockset will attempt to complete query execution before
// aborting the query and returning an error.
func WithAsyncTimeout(timeout int64) QueryOption {
	return func(o *openapi.QueryRequest) {
		if o.AsyncOptions == nil {
			o.AsyncOptions = &openapi.AsyncQueryOptions{}
		}
		o.AsyncOptions.TimeoutMs = &timeout
	}
}

// WithAsyncMaxInitialResults maximum number of results you will receive as a client. If the query exceeds this limit,
// the remaining results can be requested using a returned pagination cursor. In addition, there is a maximum response
// size of 100MiB so fewer than max_results may be returned.
func WithAsyncMaxInitialResults(timeout int64) QueryOption {
	return func(o *openapi.QueryRequest) {
		if o.AsyncOptions == nil {
			o.AsyncOptions = &openapi.AsyncQueryOptions{}
		}
		o.AsyncOptions.MaxInitialResults = &timeout
	}
}
