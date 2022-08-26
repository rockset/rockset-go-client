package option

type KeyState string

const (
	KeyActive    KeyState = "ACTIVE"
	KeySuspended KeyState = "SUSPENDED"
)

type APIKeyOptions struct {
	User   *string
	State  *KeyState
	Reveal bool
}

type APIKeyOption func(*APIKeyOptions)

// ForUser is used to scope API Key commands to a particular user.
func ForUser(username string) APIKeyOption {
	return func(o *APIKeyOptions) {
		o.User = &username
	}
}

// RevealKey is used to retrieve the full API key.
func RevealKey() APIKeyOption {
	return func(o *APIKeyOptions) {
		o.Reveal = true
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
