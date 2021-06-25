package option

type AliasOptions struct {
	Description *string
}

type AliasOption func(*AliasOptions)

// WithAliasDescription is used to optionally set the description for an alias.
func WithAliasDescription(name string) AliasOption {
	return func(o *AliasOptions) {
		o.Description = &name
	}
}

type ListAliasesOptions struct {
	Workspace string
}

type ListAliasesOption func(*ListAliasesOptions)

// WithAliasWorkspace is used to scope a listing of aliases to a workspace.
func WithAliasWorkspace(name string) ListAliasesOption {
	return func(o *ListAliasesOptions) {
		o.Workspace = name
	}
}
