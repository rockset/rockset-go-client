package rockset

import (
	"context"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
	"net/http"
)

type QueryState string

const (
	QueryQueued    QueryState = "QUEUED"
	QueryRunning   QueryState = "RUNNING"
	QueryError     QueryState = "ERROR"
	QueryCompleted QueryState = "COMPLETED"
	QueryCancelled QueryState = "CANCELLED"
)

// Query executes a sql query with optional option.QueryOption
func (rc *RockClient) Query(ctx context.Context, sql string,
	options ...option.QueryOption) (openapi.QueryResponse, error) {
	var err error
	var httpResp *http.Response
	var response *openapi.QueryResponse

	q := rc.QueriesApi.Query(ctx)
	rq := openapi.NewQueryRequestWithDefaults()
	rq.Sql = openapi.QueryRequestSql{Query: sql}
	rq.Sql.Parameters = []openapi.QueryParameter{}

	for _, o := range options {
		o(rq)
	}

	err = rc.Retry(ctx, func() error {
		response, httpResp, err = q.Body(*rq).Execute()

		return NewErrorWithStatusCode(err, httpResp.StatusCode)
	})

	if err != nil {
		return openapi.QueryResponse{}, err
	}

	return *response, nil
}

// ValidateQuery validates a sql query with optional option.QueryOption
func (rc *RockClient) ValidateQuery(ctx context.Context, sql string,
	options ...option.QueryOption) (openapi.ValidateQueryResponse, error) {
	var err error
	var httpResp *http.Response
	var r *openapi.ValidateQueryResponse

	q := rc.QueriesApi.Validate(ctx)

	rq := openapi.NewQueryRequestWithDefaults()
	rq.Sql = openapi.QueryRequestSql{Query: sql}
	rq.Sql.Parameters = []openapi.QueryParameter{}

	for _, o := range options {
		o(rq)
	}

	err = rc.Retry(ctx, func() error {
		r, httpResp, err = q.Body(*rq).Execute()

		return NewErrorWithStatusCode(err, httpResp.StatusCode)
	})

	if err != nil {
		return openapi.ValidateQueryResponse{}, err
	}

	return *r, nil
}

// GetQueryInfo retrieves information about a query.
func (rc *RockClient) GetQueryInfo(ctx context.Context, queryID string) (openapi.QueryInfo, error) {
	var err error
	var httpResp *http.Response
	var response *openapi.GetQueryResponse

	q := rc.QueriesApi.GetQuery(ctx, queryID)

	err = rc.Retry(ctx, func() error {
		response, httpResp, err = q.Execute()

		return NewErrorWithStatusCode(err, httpResp.StatusCode)
	})

	if err != nil {
		return openapi.QueryInfo{}, err
	}

	return *response.Data, nil
}

// GetQueryResults retrieves the results of a completed async query.
func (rc *RockClient) GetQueryResults(ctx context.Context, queryID string) (openapi.QueryPaginationResponse, error) {
	var err error
	var httpResp *http.Response
	var response *openapi.QueryPaginationResponse

	q := rc.QueriesApi.GetQueryResults(ctx, queryID)

	err = rc.Retry(ctx, func() error {
		response, httpResp, err = q.Execute()

		return NewErrorWithStatusCode(err, httpResp.StatusCode)
	})

	if err != nil {
		return openapi.QueryPaginationResponse{}, err
	}

	return *response, nil
}

// ListActiveQueries lists all active queries, i.e. queued or running.
func (rc *RockClient) ListActiveQueries(ctx context.Context) ([]openapi.QueryInfo, error) {
	var err error
	var httpResp *http.Response
	var response *openapi.ListQueriesResponse

	q := rc.QueriesApi.ListActiveQueries(ctx)

	err = rc.Retry(ctx, func() error {
		response, httpResp, err = q.Execute()

		return NewErrorWithStatusCode(err, httpResp.StatusCode)
	})

	return response.Data, nil
}

// CancelQuery cancels a queued or running query.
func (rc *RockClient) CancelQuery(ctx context.Context, queryID string) (openapi.QueryInfo, error) {
	var err error
	var httpResp *http.Response
	var response *openapi.CancelQueryResponse

	q := rc.QueriesApi.CancelQuery(ctx, queryID)

	err = rc.Retry(ctx, func() error {
		response, httpResp, err = q.Execute()

		return NewErrorWithStatusCode(err, httpResp.StatusCode)
	})

	return *response.Data, nil
}
