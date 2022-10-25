package option

import "github.com/rockset/rockset-go-client/openapi"

// UserOptions is used to hold optional user settings
type UserOptions struct {
	openapi.User
	FirstName *string
	LastName  *string
	Roles     []openapi.Role
}

type UserOption func(*UserOptions)

// WithUserFirstName is used to optionally set a first name for the user.
func WithUserFirstName(firstName string) UserOption {
	return func(o *UserOptions) {
		o.FirstName = &firstName
	}
}

// WithUserLastName is used to optionally set a last name for the user.
func WithUserLastName(lastName string) UserOption {
	return func(o *UserOptions) {
		o.LastName = &lastName
	}
}
