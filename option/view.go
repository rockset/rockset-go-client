package option

// ViewOptions is used to hold optional view settings
type ViewOptions struct {
	// Description of the view
	Description *string
}

type ViewOption func(p *ViewOptions)

// WithViewDescription is used to optionally set a view description.
func WithViewDescription(desc string) ViewOption {
	return func(o *ViewOptions) {
		o.Description = &desc
	}
}

type ListViewOptions struct {
	Workspace string
}

type ListViewOption func(o *ListViewOptions)

// WithViewWorkspace limits ListViews to the specified workspace
func WithViewWorkspace(ws string) ListViewOption {
	return func(o *ListViewOptions) {
		o.Workspace = ws
	}
}
