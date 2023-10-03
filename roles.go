package rockset

import (
	"context"
	"net/http"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

const (
	ReadOnlyRole = "read-only"
	MemberRole   = "member"
	AdminRole    = "admin"
)

// CreateRole creates a new role
//
// REST API documentation https://docs.rockset.com/rest-api/#createrole
func (rc *RockClient) CreateRole(ctx context.Context, roleName string,
	options ...option.RoleOption) (openapi.Role, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.RoleResponse

	createReq := rc.CustomRolesApi.CreateRole(ctx)
	b := openapi.NewCreateRoleRequest()
	b.RoleName = &roleName
	b.Privileges = []openapi.Privilege{}

	var opts option.RoleOptions
	for _, o := range options {
		o(&opts)
	}

	if opts.Description != nil {
		b.Description = opts.Description
	}
	if len(opts.Privileges) > 0 {
		b.Privileges = opts.Privileges
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = createReq.Body(*b).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.Role{}, err
	}

	return resp.GetData(), nil
}

// UpdateRole updates a role.
//
// REST API documentation https://docs.rockset.com/rest-api/#updaterole
func (rc *RockClient) UpdateRole(ctx context.Context, roleName string,
	options ...option.RoleOption) (openapi.Role, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.RoleResponse

	createReq := rc.CustomRolesApi.UpdateRole(ctx, roleName)
	b := openapi.NewUpdateRoleRequest()

	var opts option.RoleOptions
	for _, o := range options {
		o(&opts)
	}

	if opts.Description != nil {
		b.Description = opts.Description
	}
	if len(opts.Privileges) > 0 {
		b.Privileges = opts.Privileges
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = createReq.Body(*b).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.Role{}, err
	}

	return resp.GetData(), nil
}

// DeleteRole deletes a role.
//
// REST API documentation https://docs.rockset.com/rest-api/#deleterole
func (rc *RockClient) DeleteRole(ctx context.Context, roleName string) error {
	var err error
	var httpResp *http.Response

	getReq := rc.CustomRolesApi.DeleteRole(ctx, roleName)

	err = rc.Retry(ctx, func() error {
		_, httpResp, err = getReq.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return err
	}

	return nil
}

// GetRole retrieve a role.
//
// REST API documentation https://docs.rockset.com/rest-api/#getrole
func (rc *RockClient) GetRole(ctx context.Context, roleName string) (openapi.Role, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.RoleResponse

	getReq := rc.CustomRolesApi.GetRole(ctx, roleName)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = getReq.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.Role{}, err
	}

	return resp.GetData(), nil
}

// ListRoles list all roles.
//
// REST API documentation https://docs.rockset.com/rest-api/#listroles
func (rc *RockClient) ListRoles(ctx context.Context) ([]openapi.Role, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListRolesResponse

	getReq := rc.CustomRolesApi.ListRoles(ctx)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = getReq.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
