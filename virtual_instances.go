package rockset

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// CreateVirtualInstance creates a new virtual instance.
// Note that not supplying option.WithMountRefreshInterval() or option.WithContinuousMountRefresh() will
// create a virtual instance that will never refresh the mounts.
//
// REST API documentation https://rockset.com/docs/rest-api/#createvirtualinstance
func (rc *RockClient) CreateVirtualInstance(ctx context.Context, name string,
	options ...option.VirtualInstanceOption) (openapi.VirtualInstance, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.CreateVirtualInstanceResponse

	opts := option.VirtualInstanceOptions{}
	for _, o := range options {
		o(&opts)
	}

	zl := zerolog.Ctx(ctx)
	req := openapi.CreateVirtualInstanceRequest{Name: name}
	if opts.Description != nil {
		req.Description = opts.Description
	}
	if opts.Size != nil {
		req.Type = opts.Size
	}
	if opts.AutoSuspend != nil {
		s := int32(opts.AutoSuspend.Seconds())
		zl.Info().Int32("s", s).Dur("d", *opts.AutoSuspend).Msg("auto suspend")
		req.AutoSuspendSeconds = &s
	}
	if opts.EnableRemountOnResume != nil {
		req.EnableRemountOnResume = opts.EnableRemountOnResume
	}

	q := rc.VirtualInstancesApi.CreateVirtualInstance(ctx)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.VirtualInstance{}, err
	}

	return resp.GetData(), nil
}

// DeleteVirtualInstance deletes a virtual instance.
//
// REST API documentation https://rockset.com/docs/rest-api/#deletevirtualinstance
func (rc *RockClient) DeleteVirtualInstance(ctx context.Context, vID string) (openapi.VirtualInstance, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.DeleteVirtualInstanceResponse

	q := rc.VirtualInstancesApi.DeleteVirtualInstance(ctx, vID)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.VirtualInstance{}, err
	}

	return resp.GetData(), nil
}

// GetVirtualInstance gets a virtual instance by the virtual instance uuid.
//
// REST API documentation https://docs.rockset.com/rest-api/#getvirtualinstance
func (rc *RockClient) GetVirtualInstance(ctx context.Context, vID string) (openapi.VirtualInstance, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.GetVirtualInstanceResponse

	q := rc.VirtualInstancesApi.GetVirtualInstance(ctx, vID)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.VirtualInstance{}, err
	}

	return resp.GetData(), nil
}

// ListVirtualInstances lists all virtual instances.
//
// REST API documentation https://docs.rockset.com/rest-api/#listvirtualinstances
func (rc *RockClient) ListVirtualInstances(ctx context.Context) ([]openapi.VirtualInstance, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListVirtualInstancesResponse

	q := rc.VirtualInstancesApi.ListVirtualInstances(ctx)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

// UpdateVirtualInstance updates the properties of a virtual instance.
//
// REST API documentation https://docs.rockset.com/rest-api/#setvirtualinstance
func (rc *RockClient) UpdateVirtualInstance(ctx context.Context, vID string,
	options ...option.VirtualInstanceOption) (openapi.VirtualInstance, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.UpdateVirtualInstanceResponse

	opts := option.VirtualInstanceOptions{}
	for _, o := range options {
		o(&opts)
	}

	q := rc.VirtualInstancesApi.SetVirtualInstance(ctx, vID)
	req := openapi.NewUpdateVirtualInstanceRequest()

	if opts.Name != nil {
		req.Name = opts.Name
	}
	if opts.Description != nil {
		req.Description = opts.Description
	}
	if opts.Size != nil {
		req.NewSize = opts.Size
	}
	if opts.AutoSuspend != nil {
		s := int32(opts.AutoSuspend.Seconds())
		req.AutoSuspendSeconds = &s
	}
	if opts.EnableRemountOnResume != nil {
		req.EnableRemountOnResume = opts.EnableRemountOnResume
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.VirtualInstance{}, err
	}

	return resp.GetData(), nil
}

// SuspendVirtualInstance suspends a virtual instance.
//
// REST API documentation https://docs.rockset.com/rest-api/#suspendvirtualinstance
func (rc *RockClient) SuspendVirtualInstance(ctx context.Context, vID string) (openapi.VirtualInstance, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.SuspendVirtualInstanceResponse

	q := rc.VirtualInstancesApi.SuspendVirtualInstance(ctx, vID)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.VirtualInstance{}, err
	}

	return resp.GetData(), nil
}

// ResumeVirtualInstance resumes a virtual instance.
//
// REST API documentation https://docs.rockset.com/rest-api/#resumevirtualinstance
func (rc *RockClient) ResumeVirtualInstance(ctx context.Context, vID string) (openapi.VirtualInstance, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ResumeVirtualInstanceResponse

	q := rc.VirtualInstancesApi.ResumeVirtualInstance(ctx, vID)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.VirtualInstance{}, err
	}

	return resp.GetData(), nil
}

// ExecuteQueryOnVirtualInstance executes the SQL query on a specific virtual instance instead of the main
// virtual instance.
//
// REST API documentation https://rockset.com/docs/rest-api/#queryvirtualinstance
func (rc *RockClient) ExecuteQueryOnVirtualInstance(ctx context.Context, vID string, sql string,
	options ...option.QueryOption) (openapi.QueryResponse, error) {
	return queryWrapper(ctx, rc, vID, sql, options...)
}

// ListVirtualInstanceQueries lists actively queued and running queries for a particular Virtual Instance.
//
// REST API documentation
func (rc *RockClient) ListVirtualInstanceQueries(ctx context.Context, vID string) ([]openapi.QueryInfo, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListQueriesResponse

	q := rc.VirtualInstancesApi.GetVirtualInstanceQueries(ctx, vID)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// ListCollectionMounts lists collection mounts for a particular virtual instance.
//
// REST API documentation https://rockset.com/docs/rest-api/#listcollectionmounts
func (rc *RockClient) ListCollectionMounts(ctx context.Context, vID string) ([]openapi.CollectionMount, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListCollectionMountsResponse

	q := rc.VirtualInstancesApi.ListCollectionMounts(ctx, vID)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// GetCollectionMount gets a mount on this virtual instance.
//
// REST API documentation https://rockset.com/docs/rest-api/#getcollectionmount
func (rc *RockClient) GetCollectionMount(ctx context.Context, vID,
	collectionPath string) (openapi.CollectionMount, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.CollectionMountResponse

	q := rc.VirtualInstancesApi.GetCollectionMount(ctx, vID, collectionPath)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.CollectionMount{}, err
	}

	return *resp.Data, nil
}

// MountCollections mounts collections on a virtual instance, where the collections are listed as
// workspace.collection.
//
// REST API documentation https://rockset.com/docs/rest-api/#mountcollection
func (rc *RockClient) MountCollections(ctx context.Context, vID string,
	collectionPaths []string) ([]openapi.CollectionMount, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.CreateCollectionMountsResponse

	q := rc.VirtualInstancesApi.MountCollection(ctx, vID)
	req := openapi.CreateCollectionMountRequest{
		CollectionPaths: collectionPaths,
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// UnmountCollection unmount a collection from a virtual instance.
//
// REST API documentation https://rockset.com/docs/rest-api/#unmountcollection
func (rc *RockClient) UnmountCollection(ctx context.Context, vID string,
	collectionPath string) (openapi.CollectionMount, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.CollectionMountResponse

	q := rc.VirtualInstancesApi.UnmountCollection(ctx, vID, collectionPath)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.CollectionMount{}, err
	}

	return *resp.Data, nil
}
