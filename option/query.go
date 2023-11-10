package option

import (
	"time"

	"github.com/rockset/rockset-go-client/openapi"
)

type QueryState string

func (q QueryState) String() string { return string(q) }

const (
	QueryQueued    QueryState = "QUEUED"
	QueryRunning   QueryState = "RUNNING"
	QueryError     QueryState = "ERROR"
	QueryCompleted QueryState = "COMPLETED"
	QueryCancelled QueryState = "CANCELLED"
)

type QueryOptions struct {
	*openapi.QueryRequest
	VirtualInstance *string
}

type QueryOption func(request *QueryOptions)

// WithDefaultRowLimit sets the row limit to use. Limits specified in the query text will override this default.
func WithDefaultRowLimit(limit int32) QueryOption {
	return func(o *QueryOptions) {
		o.Sql.DefaultRowLimit = &limit
	}
}

// WithRowLimit sets the row limit to use. Limits specified in the query text will override this default.
// Deprecated, use WithDefaultRowLimit instead.
func WithRowLimit(limit int32) QueryOption {
	return func(o *QueryOptions) {
		o.Sql.DefaultRowLimit = &limit
	}
}

// WithVirtualInstance makes the query execute on a specific virtual instance, instead of the default (main).
func WithVirtualInstance(id string) QueryOption {
	return func(o *QueryOptions) {
		o.VirtualInstance = &id
	}
}

// WithParameter adds a query parameter. The type field is deprecated, and type is inferred from the
// JSON object value of the field if Type isn't set. Literal value of the field if Type is set.
func WithParameter(name, valueType, value string) QueryOption {
	return func(o *QueryOptions) {
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
	return func(o *QueryOptions) {
		o.Async = openapi.PtrBool(true)
	}
}

// WithTimeout is the maximum amount of time that Rockset will attempt to complete query execution before
// aborting the query and returning an error. The query timeout defaults to a maximum of 2 minutes.
// If WithAsync is set, the query timeout defaults to a maximum of 30 minutes.
func WithTimeout(timeout time.Duration) QueryOption {
	return func(o *QueryOptions) {
		t := timeout.Milliseconds()
		o.TimeoutMs = &t
	}
}

// WithMaxInitialResults is the maximum number of results you will receive as a client. If the query exceeds this limit,
// the remaining results can be requested using a returned pagination cursor.
func WithMaxInitialResults(maxInitialResults int64) QueryOption {
	return func(o *QueryOptions) {
		o.MaxInitialResults = &maxInitialResults
	}
}

// WithAsyncClientTimeout is maximum amount of time that the client is willing to wait for the query to complete.
// If the query is not complete by this timeout, a response will be returned with a query_id that can be used to
// check the status of the query and retrieve results once the query has completed.
func WithAsyncClientTimeout(timeout time.Duration) QueryOption {
	return func(o *QueryOptions) {
		if o.AsyncOptions == nil {
			o.AsyncOptions = &openapi.AsyncQueryOptions{}
		}
		t := timeout.Milliseconds()
		o.AsyncOptions.ClientTimeoutMs = &t
	}
}

// WithDebugThreshold enables the option to log debug information if query execution takes longer than the duration.
// If the query text includes the DEBUG hint and this parameter is also provided, only this value will
// be used and the DEBUG hint will be ignored.
func WithDebugThreshold(timeout time.Duration) QueryOption {
	return func(o *QueryOptions) {
		t := timeout.Milliseconds()
		o.DebugThresholdMs = &t
	}
}

// WithAsyncTimeout maximum amount of time that Rockset will attempt to complete query execution before
// aborting the query and returning an error.
// Deprecated: this option has no effect
func WithAsyncTimeout(timeout time.Duration) QueryOption {
	return func(o *QueryOptions) {
		if o.AsyncOptions == nil {
			o.AsyncOptions = &openapi.AsyncQueryOptions{}
		}
		t := timeout.Milliseconds()
		o.AsyncOptions.TimeoutMs = &t
	}
}

// WithAsyncMaxInitialResults maximum number of results you will receive as a client. If the query exceeds this limit,
// the remaining results can be requested using a returned pagination cursor. In addition, there is a maximum response
// size of 100MiB so fewer than max_results may be returned.
// Deprecated: this option has no effect
func WithAsyncMaxInitialResults(count int64) QueryOption {
	return func(o *QueryOptions) {
		if o.AsyncOptions == nil {
			o.AsyncOptions = &openapi.AsyncQueryOptions{}
		}
		o.AsyncOptions.MaxInitialResults = &count
	}
}

type QueryResultOptions struct {
	Cursor *string
	Docs   *int32
	Offset *int32
}
type QueryResultOption func(request *QueryResultOptions)

// WithQueryResultCursor sets the cursor to use to fetch next page of the query results.
// If not set, will default to the first page
func WithQueryResultCursor(cursor string) QueryResultOption {
	return func(o *QueryResultOptions) {
		o.Cursor = &cursor
	}
}

// WithQueryResultDocs sets the max number of documents to fetch.
func WithQueryResultDocs(count int32) QueryResultOption {
	return func(o *QueryResultOptions) {
		o.Docs = &count
	}
}

// WithQueryResultOffset sets the offset from the cursor of the first document to be returned.
func WithQueryResultOffset(offset int32) QueryResultOption {
	return func(o *QueryResultOptions) {
		o.Offset = &offset
	}
}
