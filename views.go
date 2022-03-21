package rockset

import (
	"context"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rs/zerolog"
	"net/http"
)

// CreateView creates a new view, with an optional description.
//
// REST API documentation https://docs.rockset.com/rest-api/#createview
func (rc *RockClient) CreateView(ctx context.Context, workspace, view, query string,
	options ...option.ViewOption) (openapi.View, error) {
	var err error
	var resp *openapi.CreateViewResponse
	var hr *http.Response

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
		resp, hr, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.View{}, err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", hr.Status).Str("state", resp.Data.GetState()).Msg("view created")

	return resp.GetData(), nil
}

// UpdateView updates an existing view, with an optional description.
//
// REST API documentation https://docs.rockset.com/rest-api/#updateview
func (rc *RockClient) UpdateView(ctx context.Context, workspace, view, query string,
	options ...option.ViewOption) (openapi.View, error) {
	var err error
	var resp *openapi.UpdateViewResponse
	var hr *http.Response

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
		resp, hr, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.View{}, err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", hr.Status).Msg("view updated")

	return resp.GetData(), nil
}

// DeleteView deletes a view.
//
// REST API documentation https://docs.rockset.com/rest-api/#deleteview
func (rc *RockClient) DeleteView(ctx context.Context, workspace, view string) error {
	var err error
	var hr *http.Response

	q := rc.ViewsApi.DeleteView(ctx, workspace, view)

	err = rc.Retry(ctx, func() error {
		_, hr, err = q.Execute()
		return err
	})

	if err != nil {
		return err
	}

	log := zerolog.Ctx(ctx)
	log.Trace().Str("status", hr.Status).Msg("view deleted")

	return nil
}

// ListViews list views.
//
// REST API documentation https://docs.rockset.com/rest-api/#listviews
func (rc *RockClient) ListViews(ctx context.Context, options ...option.ListViewOption) ([]openapi.View, error) {
	var err error
	var resp *openapi.ListViewsResponse

	opts := option.ListViewOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Workspace == "" {
		q := rc.ViewsApi.ListViews(ctx)

		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Execute()
			return err
		})
	} else {
		q := rc.ViewsApi.WorkspaceViews(ctx, opts.Workspace)

		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Execute()
			return err
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
	var resp *openapi.GetViewResponse
	log := zerolog.Ctx(ctx)

	q := rc.ViewsApi.GetView(ctx, workspace, name)

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Execute()
		return err
	})

	if err != nil {
		return openapi.View{}, err
	}

	data := resp.GetData()
	log.Debug().Str("name", data.GetName()).Msg("get view successful")

	return data, nil
}
