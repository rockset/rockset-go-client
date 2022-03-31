package option

type APIKeyOptions struct {
	User *string
}

type APIKeyOption func(*APIKeyOptions)

// ForUser is used to scope API Key commands to a particular user.
func ForUser(username string) APIKeyOption {
	return func(o *APIKeyOptions) {
		o.User = &username
	}
}

type APIKeyRoleOptions struct {
	Role *string
}

type APIKeyRoleOption func(*APIKeyRoleOptions)

func WithRole(role string) APIKeyRoleOption {
	return func(o *APIKeyRoleOptions) {
		o.Role = &role
	}
}
