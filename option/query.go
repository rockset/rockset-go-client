package option

import "github.com/rockset/rockset-go-client/openapi"

type QueryState string

const (
	QueryQueued    QueryState = "QUEUED"
	QueryRunning   QueryState = "RUNNING"
	QueryError     QueryState = "ERROR"
	QueryCompleted QueryState = "COMPLETED"
	QueryCancelled QueryState = "CANCELLED"
)

type QueryOption func(request *openapi.QueryRequest)

// WithDefaultRowLimit sets the row limit to use. Limits specified in the query text will override this default.
func WithDefaultRowLimit(limit int32) QueryOption {
	return func(o *openapi.QueryRequest) {
		o.Sql.DefaultRowLimit = &limit
	}
}

// WithRowLimit sets the row limit to use. Limits specified in the query text will override this default.
// Deprecated, use WithDefaultRowLimit instead.
func WithRowLimit(limit int32) QueryOption {
	return func(o *openapi.QueryRequest) {
		o.Sql.DefaultRowLimit = &limit
	}
}

// TODO add WithVirtualInstance for the regular query api and make it use rockset.ExecuteQueryOnVirtualInstance

func WithParameter(name, valueType, value string) QueryOption {
	return func(o *openapi.QueryRequest) {
		o.Sql.Parameters = append(o.Sql.Parameters, openapi.QueryParameter{
			Name:  name,
			Type:  valueType,
			Value: value,
		})
	}
}

// WithAsync means that the query will run asynchronously for up to 30 minutes.
// The query request will immediately return with a query id that can be used to retrieve the query status and results.
// If not specified, the query will return with results once completed or timeout after 2 minutes.
// (To return results directly for shorter queries while still allowing a timeout of up to 30 minutes,
// use WithAsyncClientTimeout())
func WithAsync() QueryOption {
	return func(o *openapi.QueryRequest) {
		o.Async = openapi.PtrBool(true)
	}
}

// WithTimeout is the maximum amount of time that Rockset will attempt to complete query execution before
// aborting the query and returning an error. The query timeout defaults to a maximum of 2 minutes.
// If WithAsync is set, the query timeout defaults to a maximum of 30 minutes.
func WithTimeout(timeout int64) QueryOption {
	return func(o *openapi.QueryRequest) {
		o.TimeoutMs = &timeout
	}
}

// WithMaxInitialResults is the maximum number of results you will receive as a client. If the query exceeds this limit,
// the remaining results can be requested using a returned pagination cursor.
func WithMaxInitialResults(maxInitialResults int64) QueryOption {
	return func(o *openapi.QueryRequest) {
		o.MaxInitialResults = &maxInitialResults
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
