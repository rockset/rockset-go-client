package option

// VirtualInstanceOptions contains the optional settings for a virtual instance.
type VirtualInstanceOptions struct {
	MonitoringEnabled *bool
	Size              *string
}

type VirtualInstanceOption func(*VirtualInstanceOptions)

// WithVIMonitoring is used to optionally set the virtual instance monitoring.
func WithVIMonitoring(enabled bool) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.MonitoringEnabled = &enabled
	}
}

// WithVISize is used to optionally set the virtual instance size.
func WithVISize(size VirtualInstanceSize) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		s := size.String()
		o.Size = &s
	}
}

// TODO: once the openapi-generator generates a custom type for each enum the below two types can be replaced
// https://github.com/OpenAPITools/openapi-generator/issues/9567

type VirtualInstanceSize string

func (t VirtualInstanceSize) String() string {
	return string(t)
}

const (
	FreeSize     VirtualInstanceSize = "FREE"
	SharedSize   VirtualInstanceSize = "SHARED"
	SmallSize    VirtualInstanceSize = "SMALL"
	MediumSize   VirtualInstanceSize = "MEDIUM"
	LargeSize    VirtualInstanceSize = "LARGE"
	XLargeSize   VirtualInstanceSize = "XLARGE"
	XLarge2Size  VirtualInstanceSize = "XLARGE2"
	XLarge4Size  VirtualInstanceSize = "XLARGE4"
	XLarge8Size  VirtualInstanceSize = "XLARGE8"
	XLarge16Size VirtualInstanceSize = "XLARGE16"
)
