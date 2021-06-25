package option

type IPAllowlistOptions struct {
	Description *string
}

type IPAllowlistOption func(*IPAllowlistOptions)

// WithIPAllowlistDescription is used to optionally set the IP allowlist description.
func WithIPAllowlistDescription(desc string) IPAllowlistOption {
	return func(o *IPAllowlistOptions) {
		o.Description = &desc
	}
}
