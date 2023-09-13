package rockset

import (
	"context"
	"net/http"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
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
	return queryWrapper(ctx, rc, "", sql, options...)
}

func queryWrapper(ctx context.Context, rc *RockClient, vID, sql string,
	options ...option.QueryOption) (openapi.QueryResponse, error) {
	var err error
	var httpResp *http.Response
	var response *openapi.QueryResponse
	var plainQuery openapi.ApiQueryRequest
	var viQuery openapi.ApiQueryVirtualInstanceRequest

	if vID == "" {
		plainQuery = rc.QueriesApi.Query(ctx)
	} else {
		viQuery = rc.VirtualInstancesApi.QueryVirtualInstance(ctx, vID)
	}

	queryRequest := openapi.NewQueryRequestWithDefaults()
	queryRequest.Sql = openapi.QueryRequestSql{Query: sql}
	queryRequest.Sql.Parameters = []openapi.QueryParameter{}

	for _, o := range options {
		o(queryRequest)
	}

	err = rc.Retry(ctx, func() error {
		if vID == "" {
			response, httpResp, err = plainQuery.Body(*queryRequest).Execute()
		} else {
			response, httpResp, err = viQuery.Body(*queryRequest).Execute()
		}

		return NewErrorWithStatusCode(err, httpResp)
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

		return NewErrorWithStatusCode(err, httpResp)
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

		return NewErrorWithStatusCode(err, httpResp)
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

		return NewErrorWithStatusCode(err, httpResp)
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

		return NewErrorWithStatusCode(err, httpResp)
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

		return NewErrorWithStatusCode(err, httpResp)
	})

	return *response.Data, nil
}
