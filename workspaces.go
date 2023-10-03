package rockset

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// CreateWorkspace creates a new workspace, with an optional description.
//
// REST API documentation https://docs.rockset.com/rest-api/#createworkspace
func (rc *RockClient) CreateWorkspace(ctx context.Context, workspace string,
	options ...option.WorkspaceOption) (openapi.Workspace, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.CreateWorkspaceResponse

	q := rc.WorkspacesApi.CreateWorkspace(ctx)
	req := openapi.NewCreateWorkspaceRequest(workspace)

	opts := option.WorkspaceOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Description != nil {
		req.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.Workspace{}, err
	}

	return resp.GetData(), nil
}

// GetWorkspace gets a workspace.
//
// REST API documentation https://docs.rockset.com/rest-api/#getworkspace
func (rc *RockClient) GetWorkspace(ctx context.Context, workspace string) (openapi.Workspace, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.GetWorkspaceResponse
	log := zerolog.Ctx(ctx)

	q := rc.WorkspacesApi.GetWorkspace(ctx, workspace)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.Workspace{}, err
	}

	data := resp.GetData()
	log.Debug().Str("name", data.GetName()).Msg("get workspace successful")

	return data, nil
}

// ListWorkspaces list all workspaces.
//
// REST API documentation https://docs.rockset.com/rest-api/#listworkspaces
func (rc *RockClient) ListWorkspaces(ctx context.Context) ([]openapi.Workspace, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListWorkspacesResponse

	q := rc.WorkspacesApi.ListWorkspaces(ctx)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

// DeleteWorkspace deletes a workspace.
//
// REST API documentation https://docs.rockset.com/rest-api/#deleteworkspace
func (rc *RockClient) DeleteWorkspace(ctx context.Context, name string) error {
	var err error
	var httpResp *http.Response

	q := rc.WorkspacesApi.DeleteWorkspace(ctx, name)

	err = rc.Retry(ctx, func() error {
		_, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return err
	}

	return nil
}
