package rockset

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// CreateView creates a new view, with an optional description.
//
// REST API documentation https://docs.rockset.com/rest-api/#createview
func (rc *RockClient) CreateView(ctx context.Context, workspace, view, query string,
	options ...option.ViewOption) (openapi.View, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.CreateViewResponse

	q := rc.ViewsApi.CreateView(ctx, workspace)
	req := openapi.NewCreateViewRequest(view, query)

	opts := option.ViewOptions{}
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
		return openapi.View{}, err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", httpResp.Status).Str("state", resp.Data.GetState()).Msg("view created")

	return resp.GetData(), nil
}

// UpdateView updates an existing view, with an optional description.
//
// REST API documentation https://docs.rockset.com/rest-api/#updateview
func (rc *RockClient) UpdateView(ctx context.Context, workspace, view, query string,
	options ...option.ViewOption) (openapi.View, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.UpdateViewResponse

	q := rc.ViewsApi.UpdateView(ctx, workspace, view)
	req := openapi.NewUpdateViewRequest(query)

	opts := option.ViewOptions{}
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
		return openapi.View{}, err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", httpResp.Status).Msg("view updated")

	return resp.GetData(), nil
}

// DeleteView marks the view for deletion, which will take place in the background. Use the
// WaitUntilViewGone() call to block until the view has been deleted.
//
// REST API documentation https://docs.rockset.com/rest-api/#deleteview
func (rc *RockClient) DeleteView(ctx context.Context, workspace, view string) error {
	var err error
	var httpResp *http.Response

	q := rc.ViewsApi.DeleteView(ctx, workspace, view)

	err = rc.Retry(ctx, func() error {
		_, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", httpResp.Status).Msg("view deleted")

	return nil
}

// ListViews list views. Use option.WithViewWorkspace() to limit the request to just one workspace.
//
// REST API documentation https://docs.rockset.com/rest-api/#listviews
func (rc *RockClient) ListViews(ctx context.Context, options ...option.ListViewOption) ([]openapi.View, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListViewsResponse

	opts := option.ListViewOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Workspace == "" {
		q := rc.ViewsApi.ListViews(ctx)

		err = rc.Retry(ctx, func() error {
			resp, httpResp, err = q.Execute()

			return rockerr.NewWithStatusCode(err, httpResp)
		})
	} else {
		q := rc.ViewsApi.WorkspaceViews(ctx, opts.Workspace)

		err = rc.Retry(ctx, func() error {
			resp, httpResp, err = q.Execute()

			return rockerr.NewWithStatusCode(err, httpResp)
		})
	}

	if err != nil {
		return nil, err
	}
	log := zerolog.Ctx(ctx)
	log.Debug().Int("count", len(resp.GetData())).Msg("list views")

	return resp.GetData(), nil
}

// GetView gets details about a view.
//
// REST API documentation https://docs.rockset.com/rest-api/#getview
func (rc *RockClient) GetView(ctx context.Context, workspace, name string) (openapi.View, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.GetViewResponse
	log := zerolog.Ctx(ctx)

	q := rc.ViewsApi.GetView(ctx, workspace, name)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.View{}, err
	}

	data := resp.GetData()
	log.Debug().Str("name", data.GetName()).Msg("get view successful")

	return data, nil
}
