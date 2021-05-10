package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (rc *RockClient) Query(ctx context.Context, sql string,
	options ...option.QueryOption) (openapi.QueryResponse, error) {
	q := rc.QueriesApi.Query(ctx)
	rq := openapi.NewQueryRequestWithDefaults()
	rq.Sql = openapi.NewQueryRequestSql(sql)
	rq.Sql.Parameters = &[]openapi.QueryParameter{}

	for _, o := range options {
		o(rq.Sql)
	}

	r, _, err := q.Body(*rq).Execute()
	if err != nil {
		return openapi.QueryResponse{}, err
	}

	return r, nil
}

func (rc *RockClient) ValidateQuery(ctx context.Context, sql string,
	options ...option.QueryOption) (openapi.ValidateQueryResponse, error) {
	q := rc.QueriesApi.Validate(ctx)

	rq := openapi.NewQueryRequestWithDefaults()
	rq.Sql = openapi.NewQueryRequestSql(sql)
	rq.Sql.Parameters = &[]openapi.QueryParameter{}

	for _, o := range options {
		o(rq.Sql)
	}

	r, _, err := q.Body(*rq).Execute()
	if err != nil {
		return openapi.ValidateQueryResponse{}, err
	}

	return r, nil
}
