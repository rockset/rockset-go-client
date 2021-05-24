package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// https://docs.rockset.com/rest-api/#virtual-instances

// GetVirtualInstance gets a virtual instance by the uuid.
//
// REST API documentation https://docs.rockset.com/rest-api/#getvirtualinstance
func (rc *RockClient) GetVirtualInstance(ctx context.Context, vID string) (openapi.VirtualInstance, error) {
	var err error
	var resp openapi.GetVirtualInstanceResponse

	q := rc.VirtualInstancesApi.GetVirtualInstance(ctx, vID)

	err = rc.Retry(ctx, func() (bool, error) {
		resp, _, err = q.Execute()
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}
			return false, re
		}

		return false, nil
	})

	if err != nil {
		return openapi.VirtualInstance{}, err
	}

	return *resp.Data, nil
}

// ListVirtualInstances gets a virtual instance by the uuid.
//
// REST API documentation https://docs.rockset.com/rest-api/#listvirtualinstances
func (rc *RockClient) ListVirtualInstances(ctx context.Context) ([]openapi.VirtualInstance, error) {
	var err error
	var resp openapi.ListVirtualInstancesResponse

	q := rc.VirtualInstancesApi.ListVirtualInstances(ctx)

	err = rc.Retry(ctx, func() (bool, error) {
		resp, _, err = q.Execute()
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}
			return false, re
		}

		return false, err
	})

	if err != nil {
		return nil, err
	}

	return *resp.Data, nil
}

// UpdateVirtualInstance updates the properties of a virtual instance.
//
// REST API documentation https://docs.rockset.com/rest-api/#setvirtualinstance
func (rc *RockClient) UpdateVirtualInstance(ctx context.Context, vID string, options ...option.VirtualInstanceOption) (openapi.VirtualInstance, error) {
	var err error
	var resp openapi.UpdateVirtualInstanceResponse

	opts := option.VirtualInstanceOptions{}
	for _, o := range options {
		o(&opts)
	}

	q := rc.VirtualInstancesApi.SetVirtualInstance(ctx, vID)
	req := openapi.NewUpdateVirtualInstanceRequest()

	if opts.Size != nil {
		req.NewSize = opts.Size
	}
	if opts.Type != nil {
		req.NewType = opts.Type
	}
	if opts.MonitoringEnabled != nil {
		req.MonitoringEnabled = opts.MonitoringEnabled
	}

	err = rc.Retry(ctx, func() (bool, error) {
		resp, _, err = q.Execute()
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}
			return false, re
		}

		return false, err
	})

	if err != nil {
		return openapi.VirtualInstance{}, err
	}

	return *resp.Data, nil
}
