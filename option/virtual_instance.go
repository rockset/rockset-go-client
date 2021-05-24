package option

type VirtualInstanceOptions struct {
	MonitoringEnabled *bool
	Type              *string
	Size              *string
}

type VirtualInstanceOption func(*VirtualInstanceOptions)

// WithVIMonitoring is used to optionally set the virtual instance monitoring.
func WithVIMonitoring(enabled bool) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.MonitoringEnabled = &enabled
	}
}

// WithVIType is used to optionally set the virtual instance type
func WithVIType(t string) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.Type = &t
	}
}

// WithVISize is used to optionally set the virtual instance type
func WithVISize(size string) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.Size = &size
	}
}
