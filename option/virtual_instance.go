package option

// VirtualInstanceOptions contains the optional settings for a virtual instance.
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

// WithVIType is used to optionally set the virtual instance type.
func WithVIType(t VirtualInstanceType) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		s := t.String()
		o.Type = &s
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

type VirtualInstanceType string

func (t VirtualInstanceType) String() string {
	return string(t)
}

const (
	FreeType     VirtualInstanceType = "FREE"
	SharedType   VirtualInstanceType = "SHARED"
	SmallType    VirtualInstanceType = "SMALL"
	MediumType   VirtualInstanceType = "MEDIUM"
	LargeType    VirtualInstanceType = "LARGE"
	XLargeType   VirtualInstanceType = "XLARGE"
	XLarge2Type  VirtualInstanceType = "XLARGE2"
	XLarge4Type  VirtualInstanceType = "XLARGE4"
	XLarge8Type  VirtualInstanceType = "XLARGE8"
	XLarge16Type VirtualInstanceType = "XLARGE16"
)

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
