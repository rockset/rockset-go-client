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
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()
			return err
		})
	} else if req.Tag != "" {
		q := rc.QueryLambdasApi.ExecuteQueryLambdaByTag(ctx, workspace, name, req.Tag)
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()
			return err
		})
	} else {
		q := rc.QueryLambdasApi.ExecuteQueryLambdaByTag(ctx, workspace, name, LatestTag)
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()
			return err
		})
	}

	if err != nil {
		return openapi.QueryResponse{}, err
	}

	return resp, nil
}

// GetQueryLambdaVersionByTag gets the query lambda version for a tag.
func (rc *RockClient) GetQueryLambdaVersionByTag(ctx context.Context,
	workspace, name, tag string) (openapi.QueryLambdaTag, error) {
	var err error
	var resp openapi.QueryLambdaTagResponse

	q := rc.QueryLambdasApi.GetQueryLambdaTagVersion(ctx, workspace, name, tag)
	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Execute()
		return err
	})

	if err != nil {
		return openapi.QueryLambdaTag{}, err
	}

	return resp.GetData(), err
}

// GetQueryLambdaVersion get the query lambda information for a specific version.
func (rc *RockClient) GetQueryLambdaVersion(ctx context.Context,
	workspace, name, version string) (openapi.QueryLambdaVersion, error) {
	var err error
	var resp openapi.QueryLambdaVersionResponse

	q := rc.QueryLambdasApi.GetQueryLambdaVersion(ctx, workspace, name, version)
	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Execute()
		return err
	})

	if err != nil {
		return openapi.QueryLambdaVersion{}, err
	}

	return resp.GetData(), err
}

// ListQueryLambdas lists all query lambdas, unless the option.WithQueryLambdaWorkspace is used.
//
// https://docs.rockset.com/rest-api/#listallquerylambdas
func (rc *RockClient) ListQueryLambdas(ctx context.Context,
	options ...option.ListQueryLambdaOption) ([]openapi.QueryLambda, error) {
	var err error
	var resp openapi.ListQueryLambdasResponse

	opts := option.ListQueryLambdaOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Workspace == nil {
		q := rc.QueryLambdasApi.ListAllQueryLambdas(ctx)
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Execute()
			return err
		})
	} else {
		q := rc.QueryLambdasApi.ListQueryLambdasInWorkspace(ctx, *opts.Workspace)
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Execute()
			return err
		})
	}

	if err != nil {
		return nil, err
	}

	return resp.GetData(), err
}

// ListQueryLambdaTagVersions lists all query lambdas for the tag.
func (rc *RockClient) ListQueryLambdaTagVersions(ctx context.Context,
	tag string) ([]openapi.QueryLambdaVersion, error) {
	var err error
	var resp openapi.ListQueryLambdaVersionsResponse

	q := rc.QueryLambdasApi.ListQueryLambdaTagVersions(ctx, tag)
	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), err
}

// ListQueryLambdaVersions lists all versions for a query lambda.
func (rc *RockClient) ListQueryLambdaVersions(ctx context.Context,
	workspace, name string) ([]openapi.QueryLambdaVersion, error) {
	var err error
	var resp openapi.ListQueryLambdaVersionsResponse

	q := rc.QueryLambdasApi.ListQueryLambdaVersions(ctx, workspace, name)
	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), err
}

// ListQueryLambdaTags lists all tags for the organization, or for a specific query lambda if the
// option.WithQueryLambda is used.
func (rc *RockClient) ListQueryLambdaTags(ctx context.Context,
	options ...option.ListQueryLambdaTagsOption) ([]openapi.QueryLambdaTag, error) {
	var err error
	var resp openapi.ListQueryLambdaTagsResponse

	opts := option.ListQueryLambdaTagsOptions{}

	for _, o := range options {
		o(&opts)
	}

	if opts.Workspace == nil {
		q := rc.QueryLambdasApi.ListOrganizationTags(ctx)
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Execute()
			return err
		})
	} else {
		q := rc.QueryLambdasApi.ListQueryLambdaTags(ctx, *opts.Workspace, *opts.QueryLambda)
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Execute()
			return err
		})
	}

	if err != nil {
		return nil, err
	}

	return resp.GetData(), err
}
