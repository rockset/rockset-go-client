package rockset

import (
	"context"
	"fmt"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// LatestTag is the query lambda tag for the latest version.
const LatestTag = "latest"

// ExecuteQueryLambda executes a query lambda with optional query options.
func (rc *RockClient) ExecuteQueryLambda(ctx context.Context, workspace, name string,
	options ...option.QueryLambdaOption) (openapi.QueryResponse, error) {

	req := option.ExecuteQueryLambdaRequest{
		ExecuteQueryLambdaRequest: openapi.ExecuteQueryLambdaRequest{},
	}
	for _, o := range options {
		o(&req)
	}

	if req.Tag != "" && req.Version != "" {
		return openapi.QueryResponse{},
			fmt.Errorf("can't specify both version %s and tag %s", req.Version, req.Tag)
	}

	var err error
	var resp openapi.QueryResponse
	if req.Version != "" {
		q := rc.QueryLambdasApi.ExecuteQueryLambda(ctx, workspace, name, req.Version)
		resp, _, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()
	} else if req.Tag != "" {
		q := rc.QueryLambdasApi.ExecuteQueryLambdaByTag(ctx, workspace, name, req.Tag)
		resp, _, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()
	} else {
		q := rc.QueryLambdasApi.ExecuteQueryLambdaByTag(ctx, workspace, name, LatestTag)
		resp, _, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()
	}

	if err != nil {
		return openapi.QueryResponse{}, err
	}

	return resp, nil
}
