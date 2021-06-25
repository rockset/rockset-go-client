package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// Query executes a sql query with optional option.QueryOption
func (rc *RockClient) Query(ctx context.Context, sql string,
	options ...option.QueryOption) (openapi.QueryResponse, error) {
	var err error
	var r openapi.QueryResponse

	q := rc.QueriesApi.Query(ctx)
	rq := openapi.NewQueryRequestWithDefaults()
	rq.Sql = openapi.NewQueryRequestSql(sql)
	rq.Sql.Parameters = &[]openapi.QueryParameter{}

	for _, o := range options {
		o(rq.Sql)
	}

	err = rc.Retry(ctx, func() error {
		r, _, err = q.Body(*rq).Execute()
		return err
	})

	if err != nil {
		return openapi.QueryResponse{}, err
	}

	return r, nil
}

// ValidateQuery validates a sql query with optional option.QueryOption
func (rc *RockClient) ValidateQuery(ctx context.Context, sql string,
	options ...option.QueryOption) (openapi.ValidateQueryResponse, error) {
	var err error
	var r openapi.ValidateQueryResponse

	q := rc.QueriesApi.Validate(ctx)

	rq := openapi.NewQueryRequestWithDefaults()
	rq.Sql = openapi.NewQueryRequestSql(sql)
	rq.Sql.Parameters = &[]openapi.QueryParameter{}

	for _, o := range options {
		o(rq.Sql)
	}

	err = rc.Retry(ctx, func() error {
		r, _, err = q.Body(*rq).Execute()
		return err
	})

	if err != nil {
		return openapi.ValidateQueryResponse{}, err
	}

	return r, nil
}
