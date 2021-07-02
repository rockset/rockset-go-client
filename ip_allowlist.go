package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// Convenience methods for IP allowlist operations
// https://docs.rockset.com/rest-api/#ip-allowlist

// CreateIPAllowlistNetworkPolicy creates a new IP allow list policy.
//
// REST API documentation https://docs.rockset.com/rest-api/#createallowlistip
func (rc *RockClient) CreateIPAllowlistNetworkPolicy(ctx context.Context, name, ipAddress string,
	options ...option.IPAllowlistOption) (openapi.IpAllowlist, error) {
	var err error
	var resp openapi.CreateIpAllowlistResponse

	opts := option.IPAllowlistOptions{}
	for _, o := range options {
		o(&opts)
	}

	req := rc.IPAllowlistApi.CreateAllowlistIp(ctx)

	b := openapi.NewCreateIpAllowlistRequest(name, ipAddress)
	if opts.Description != nil {
		b.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = req.Body(*b).Execute()
		return err
	})

	if err != nil {
		return openapi.IpAllowlist{}, err
	}

	return resp.GetData(), nil
}

// DeleteIPAllowlistNetworkPolicy deletes an IP allow list policy by name.
//
// REST API documentation https://docs.rockset.com/rest-api/#deleteallowlistip
func (rc *RockClient) DeleteIPAllowlistNetworkPolicy(ctx context.Context, name string) error {
	var err error

	req := rc.IPAllowlistApi.DeleteAllowlistIp(ctx, name)

	err = rc.Retry(ctx, func() error {
		_, _, err = req.Execute()
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

// GetIPAllowlistNetworkPolicy gets an IP allowlost policy by name.
//
// REST API documentation https://docs.rockset.com/rest-api/#getallowlistip
func (rc *RockClient) GetIPAllowlistNetworkPolicy(ctx context.Context, name string) (openapi.IpAllowlist, error) {
	var err error
	var resp openapi.GetIpAllowlistResponse

	req := rc.IPAllowlistApi.GetAllowlistIp(ctx, name)

	err = rc.Retry(ctx, func() error {
		resp, _, err = req.Execute()
		return err
	})

	if err != nil {
		return openapi.IpAllowlist{}, err
	}

	return resp.GetData(), nil
}

// ListIPAllowlistEntries lists all IP allowlists policies.
//
// REST API documentation https://docs.rockset.com/rest-api/#listallowlistips
func (rc *RockClient) ListIPAllowlistEntries(ctx context.Context) ([]openapi.IpAllowlist, error) {
	var err error
	var resp openapi.ListIpAllowlistsResponse

	req := rc.IPAllowlistApi.ListAllowlistIps(ctx)

	err = rc.Retry(ctx, func() error {
		resp, _, err = req.Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
