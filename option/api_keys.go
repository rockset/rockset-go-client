package option

type KeyState string

func (k KeyState) String() string { return string(k) }

const (
	KeyActive    KeyState = "ACTIVE"
	KeySuspended KeyState = "SUSPENDED"
)

type APIKeyOptions struct {
	User  *string
	State *KeyState
}

type APIKeyOption func(*APIKeyOptions)

// ForUser is used to scope API Key commands to a particular user.
func ForUser(username string) APIKeyOption {
	return func(o *APIKeyOptions) {
		o.User = &username
	}
}

// State is used to set the key state to either KeyActive or KeySuspended.
func State(state KeyState) APIKeyOption {
	return func(o *APIKeyOptions) {
		o.State = &state
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
