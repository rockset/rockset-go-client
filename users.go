package rockset

import (
	"context"
	"net/http"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// https://docs.rockset.com/rest-api/#users

// CreateUser creates a new user.
//
// REST API documentation https://docs.rockset.com/rest-api/#createuser
func (rc *RockClient) CreateUser(ctx context.Context, email string, roles []string) (openapi.User, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.CreateUserResponse

	q := rc.UsersApi.CreateUser(ctx)
	req := openapi.NewCreateUserRequest(email, roles)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.User{}, err
	}

	return resp.GetData(), nil
}

// UpdateUser updates as existing user. Note that the user first and last name aren't visible for privacy reasons,
// until the user has accepted the invite, i.e. is in the ACTIVE state.
//
// REST API documentation https://docs.rockset.com/rest-api/#updateuser
func (rc *RockClient) UpdateUser(ctx context.Context, email string, roles []string,
	options ...option.UserOption) (openapi.User, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.User

	q := rc.UsersApi.UpdateUser(ctx, email)
	req := openapi.NewUpdateUserRequest()
	req.Roles = roles

	var opts option.UserOptions
	for _, opt := range options {
		opt(&opts)
	}

	if opts.FirstName != nil {
		req.FirstName = opts.FirstName
	}

	if opts.LastName != nil {
		req.LastName = opts.LastName
	}

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.User{}, err
	}

	return *resp, nil
}

// DeleteUser deletes a user.
//
// REST API documentation https://docs.rockset.com/rest-api/#deleteuser
func (rc *RockClient) DeleteUser(ctx context.Context, email string) error {
	var err error
	var httpResp *http.Response

	q := rc.UsersApi.DeleteUser(ctx, email)

	err = rc.Retry(ctx, func() error {
		_, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return err
	}

	return nil
}

// GetCurrentUser gets the current user.
//
// REST API documentation https://docs.rockset.com/rest-api/#getcurrentuser
func (rc *RockClient) GetCurrentUser(ctx context.Context) (openapi.User, error) {
	var err error
	var httpResp *http.Response
	var user *openapi.User

	q := rc.UsersApi.GetCurrentUser(ctx)

	err = rc.Retry(ctx, func() error {
		user, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.User{}, err
	}

	return *user, nil
}

// GetUser gets a user.
//
// REST API documentation https://docs.rockset.com/rest-api/#getuser
func (rc *RockClient) GetUser(ctx context.Context, email string) (openapi.User, error) {
	var err error
	var httpResp *http.Response
	var user *openapi.User

	q := rc.UsersApi.GetUser(ctx, email)

	err = rc.Retry(ctx, func() error {
		user, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.User{}, err
	}

	return *user, nil
}

// ListUsers lists all users.
//
// REST API documentation https://docs.rockset.com/rest-api/#listusers
func (rc *RockClient) ListUsers(ctx context.Context) ([]openapi.User, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.ListUsersResponse

	q := rc.UsersApi.ListUsers(ctx)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
