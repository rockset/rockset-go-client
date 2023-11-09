package rockset

import (
	"context"
	"fmt"
	"net/http"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// LatestTag is the query lambda tag for the latest version.
const LatestTag = "latest"

// CreateQueryLambda creates a new query lambda.
//
// https://docs.rockset.com/rest-api/#createquerylambda
func (rc *RockClient) CreateQueryLambda(ctx context.Context, workspace, name, sql string,
	options ...option.CreateQueryLambdaOption) (openapi.QueryLambdaVersion, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.QueryLambdaVersionResponse

	q := rc.QueryLambdasApi.CreateQueryLambda(ctx, workspace)
	req := openapi.NewCreateQueryLambdaRequestWithDefaults()

	req.Name = name
	req.Sql = openapi.QueryLambdaSql{
		Query: sql,
	}

	opts := option.CreateQueryLambdaOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Description != nil {
		req.Description = opts.Description
	}

	if len(opts.QueryParameters) > 0 {
		req.Sql.DefaultParameters = opts.QueryParameters
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.QueryLambdaVersion{}, err
	}

	return resp.GetData(), nil
}

// DeleteQueryLambda deletes a query lambda.
//
// https://docs.rockset.com/rest-api/#deletequerylambda
func (rc *RockClient) DeleteQueryLambda(ctx context.Context, workspace, name string) error {
	var err error
	var httpResp *http.Response

	q := rc.QueryLambdasApi.DeleteQueryLambda(ctx, workspace, name)

	err = rc.Retry(ctx, func() error {
		_, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return err
	}

	return nil
}

// UpdateQueryLambda updates an existing query lambda.
//
// https://docs.rockset.com/rest-api/#updatequerylambda
func (rc *RockClient) UpdateQueryLambda(ctx context.Context, workspace, name, sql string,
	options ...option.CreateQueryLambdaOption) (openapi.QueryLambdaVersion, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.QueryLambdaVersionResponse

	q := rc.QueryLambdasApi.UpdateQueryLambda(ctx, workspace, name)
	req := openapi.NewUpdateQueryLambdaRequestWithDefaults()

	req.Sql = &openapi.QueryLambdaSql{
		Query: sql,
	}

	opts := option.CreateQueryLambdaOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Description != nil {
		req.Description = opts.Description
	}

	if len(opts.QueryParameters) > 0 {
		req.Sql.DefaultParameters = opts.QueryParameters
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.QueryLambdaVersion{}, err
	}

	return resp.GetData(), nil
}

// CreateQueryLambdaTag creates a new tag for a specific query lambda version, or update the tag if it already exists.
// Note that the tag propagation takes time, and the wait.UntilQueryLambdaTagPropagated() method should be called
// after updating an existing tag, to avoid eventual consistency issues.
//
// If strong consistency of the version is required when executing a query lambda,
// option.WithVersion() must be used, e.g.
//
//	r, err := rc.ExecuteQueryLambda(ctx, ws, ql, option.WithVersion("f79fc3eae5c823bb"))
//
// https://docs.rockset.com/documentation/reference/createquerylambdatag
func (rc *RockClient) CreateQueryLambdaTag(ctx context.Context, workspace, name, version,
	tag string) (openapi.QueryLambdaTag, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.QueryLambdaTagResponse

	q := rc.QueryLambdasApi.CreateQueryLambdaTag(ctx, workspace, name)
	req := openapi.NewCreateQueryLambdaTagRequestWithDefaults()

	req.TagName = tag
	req.Version = version

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.QueryLambdaTag{}, err
	}

	return resp.GetData(), nil
}

// DeleteQueryLambdaVersion deletes a query lambda version.
//
// https://docs.rockset.com/rest-api/#deletequerylambdaversion
func (rc *RockClient) DeleteQueryLambdaVersion(ctx context.Context, workspace, name, version string) error {
	var err error
	var httpResp *http.Response

	q := rc.QueryLambdasApi.DeleteQueryLambdaVersion(ctx, workspace, name, version)
	err = rc.Retry(ctx, func() error {
		_, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return err
	}

	return nil
}

// DeleteQueryLambdaTag deletes a query lambda tag.
//
// https://docs.rockset.com/rest-api/#deletequerylambdatag
func (rc *RockClient) DeleteQueryLambdaTag(ctx context.Context, workspace, name, tag string) error {
	var err error
	var httpResp *http.Response

	q := rc.QueryLambdasApi.DeleteQueryLambdaTag(ctx, workspace, name, tag)
	err = rc.Retry(ctx, func() error {
		_, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return err
	}

	return nil
}

// ExecuteQueryLambda executes a query lambda with optional query options. If neither option.WithTag() nor
// option.WithVersion() is used, it makes the call using option.WithTag(LatestTag)
//
// If strong consistency of the version is required when executing a query lambda,
// option.WithVersion() must be used, e.g.
//
//	r, err := rc.ExecuteQueryLambda(ctx, ws, ql, option.WithVersion("f79fc3eae5c823bb"))
//
// https://docs.rockset.com/documentation/reference/executequerylambdabytag
// https://docs.rockset.com/documentation/reference/executequerylambda
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
	var httpResp *http.Response
	var resp *openapi.QueryResponse
	if req.Version != "" {
		q := rc.QueryLambdasApi.ExecuteQueryLambda(ctx, workspace, name, req.Version)
		err = rc.Retry(ctx, func() error {
			resp, httpResp, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()

			return rockerr.NewWithStatusCode(err, httpResp)
		})
	} else if req.Tag != "" {
		q := rc.QueryLambdasApi.ExecuteQueryLambdaByTag(ctx, workspace, name, req.Tag)
		err = rc.Retry(ctx, func() error {
			resp, httpResp, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()

			return rockerr.NewWithStatusCode(err, httpResp)
		})
	} else {
		q := rc.QueryLambdasApi.ExecuteQueryLambdaByTag(ctx, workspace, name, LatestTag)
		err = rc.Retry(ctx, func() error {
			resp, httpResp, err = q.Body(req.ExecuteQueryLambdaRequest).Execute()

			return rockerr.NewWithStatusCode(err, httpResp)
		})
	}

	if err != nil {
		return openapi.QueryResponse{}, err
	}

	return *resp, nil
}

// GetQueryLambdaVersionByTag gets the query lambda version for a tag.
func (rc *RockClient) GetQueryLambdaVersionByTag(ctx context.Context,
	workspace, name, tag string) (openapi.QueryLambdaTag, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.QueryLambdaTagResponse

	q := rc.QueryLambdasApi.GetQueryLambdaTagVersion(ctx, workspace, name, tag)
	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
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
	var httpResp *http.Response
	var resp *openapi.QueryLambdaVersionResponse

	q := rc.QueryLambdasApi.GetQueryLambdaVersion(ctx, workspace, name, version)
	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
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
	var httpResp *http.Response
	var resp *openapi.ListQueryLambdasResponse

	opts := option.ListQueryLambdaOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Workspace == nil {
		q := rc.QueryLambdasApi.ListAllQueryLambdas(ctx)
		err = rc.Retry(ctx, func() error {
			resp, httpResp, err = q.Execute()

			return rockerr.NewWithStatusCode(err, httpResp)
		})
	} else {
		q := rc.QueryLambdasApi.ListQueryLambdasInWorkspace(ctx, *opts.Workspace)
		err = rc.Retry(ctx, func() error {
			resp, httpResp, err = q.Execute()

			return rockerr.NewWithStatusCode(err, httpResp)
		})
	}

	if err != nil {
		return nil, err
	}

	return resp.GetData(), err
}

// ListQueryLambdaVersions lists all versions for a query lambda.
func (rc *RockClient) ListQueryLambdaVersions(ctx context.Context,
	workspace, name string) ([]openapi.QueryLambdaVersion, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListQueryLambdaVersionsResponse

	q := rc.QueryLambdasApi.ListQueryLambdaVersions(ctx, workspace, name)
	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), err
}

// ListQueryLambdaTags lists all tags for a query lambda.
func (rc *RockClient) ListQueryLambdaTags(ctx context.Context, workspace,
	queryLambda string) ([]openapi.QueryLambdaTag, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListQueryLambdaTagsResponse

	q := rc.QueryLambdasApi.ListQueryLambdaTags(ctx, workspace, queryLambda)
	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), err
}
