package rockset

import (
	"context"
	"net/http"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// Query executes a sql query with optional option.QueryOption. If you want to execute the query on a specific
// virtual instance instead of the main virtual instance,
// either add the option.WithVirtualInstance() option, or use ExecuteQueryOnVirtualInstance() method.
//
//	result, err := rc.Query(ctx, "SELECT * FROM _commons._events", option.WithVirtualInstance(vID)
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

	queryRequest := openapi.NewQueryRequestWithDefaults()
	queryRequest.Sql = openapi.QueryRequestSql{Query: sql}
	queryRequest.Sql.Parameters = []openapi.QueryParameter{}

	request := option.QueryOptions{
		QueryRequest: queryRequest,
	}

	for _, o := range options {
		o(&request)
	}

	if vID != "" {
		// a VI was specified in ExecuteQueryOnVirtualInstance
		viQuery = rc.VirtualInstancesApi.QueryVirtualInstance(ctx, vID)
	} else if request.VirtualInstance != nil {
		// an option was used to specify the VI
		viQuery = rc.VirtualInstancesApi.QueryVirtualInstance(ctx, *request.VirtualInstance)
	} else {
		// no VI specified, execute on main VI
		plainQuery = rc.QueriesApi.Query(ctx)
	}

	err = rc.Retry(ctx, func() error {
		if vID != "" || request.VirtualInstance != nil {
			response, httpResp, err = viQuery.Body(*queryRequest).Execute()
		} else {
			response, httpResp, err = plainQuery.Body(*queryRequest).Execute()
		}

		return rockerr.NewWithStatusCode(err, httpResp)
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

	request := option.QueryOptions{
		QueryRequest: rq,
	}

	for _, o := range options {
		o(&request)
	}

	err = rc.Retry(ctx, func() error {
		r, httpResp, err = q.Body(*rq).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
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

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.QueryInfo{}, err
	}

	return *response.Data, nil
}

// GetQueryResults retrieves the results of a completed async query.
func (rc *RockClient) GetQueryResults(ctx context.Context, queryID string, options ...option.QueryResultOption) (openapi.QueryPaginationResponse, error) {
	var err error
	var httpResp *http.Response
	var response *openapi.QueryPaginationResponse

	q := rc.QueriesApi.GetQueryResults(ctx, queryID)
	var opts option.QueryResultOptions
	for _, opt := range options {
		opt(&opts)
	}
	if opts.Offset != nil {
		q = q.Offset(*opts.Offset)
	}
	if opts.Docs != nil {
		q = q.Docs(*opts.Docs)
	}
	if opts.Cursor != nil {
		q = q.Cursor(*opts.Cursor)
	}

	err = rc.Retry(ctx, func() error {
		response, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
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

		return rockerr.NewWithStatusCode(err, httpResp)
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

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	return *response.Data, nil
}
