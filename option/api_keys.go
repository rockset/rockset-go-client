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
