package option

// WorkspaceOptions is used to hold optional workspace settings
type WorkspaceOptions struct {
	// Description of the workspace
	Description *string
}

type WorkspaceOption func(p *WorkspaceOptions)

// WithWorkspaceDescription is used to optionally set a workspace description.
func WithWorkspaceDescription(desc string) WorkspaceOption {
	return func(o *WorkspaceOptions) {
		o.Description = &desc
	}
}
