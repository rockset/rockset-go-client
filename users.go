package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
)

// https://docs.rockset.com/rest-api/#users

// CreateUser creates a new user.
//
// REST API documentation https://docs.rockset.com/rest-api/#createuser
func (rc *RockClient) CreateUser(ctx context.Context, email string, roles []string) (openapi.User, error) {
	var err error
	var resp *openapi.CreateUserResponse

	q := rc.UsersApi.CreateUser(ctx)
	req := openapi.NewCreateUserRequest(email, roles)

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.User{}, err
	}

	return resp.GetData(), nil
}

// DeleteUser deletes a user.
//
// REST API documentation https://docs.rockset.com/rest-api/#deleteuser
func (rc *RockClient) DeleteUser(ctx context.Context, email string) error {
	var err error

	q := rc.UsersApi.DeleteUser(ctx, email)

	err = rc.Retry(ctx, func() error {
		_, _, err = q.Execute()
		return err
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
	var user *openapi.User

	q := rc.UsersApi.GetCurrentUser(ctx)

	err = rc.Retry(ctx, func() error {
		user, _, err = q.Execute()
		return err
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
	var user *openapi.User

	q := rc.UsersApi.GetUser(ctx, email)

	err = rc.Retry(ctx, func() error {
		user, _, err = q.Execute()
		return err
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
	var resp *openapi.ListUsersResponse

	q := rc.UsersApi.ListUsers(ctx)

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
