package option

import "time"

// VirtualInstanceOptions contains the optional settings for a virtual instance.
type VirtualInstanceOptions struct {
	Description           *string
	Size                  *string
	MountRefreshInterval  *time.Duration
	AutoSuspend           *time.Duration
	EnableRemountOnResume *bool
	Name                  *string
}

type VirtualInstanceOption func(*VirtualInstanceOptions)

// WithMountRefreshInterval is used to optionally set the mount refresh interval. Setting it to 0 means
// it will be a continuous refresh, but then WithContinuousMountRefresh should be used instead. Not specifying
// a mount refresh interval will make it never refresh, and should only be used for static collections.
func WithMountRefreshInterval(interval time.Duration) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.MountRefreshInterval = &interval
	}
}

// WithContinuousMountRefresh is used to set the mount refresh interval to a continuous refresh.
func WithContinuousMountRefresh() VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		var zero time.Duration
		o.MountRefreshInterval = &zero
	}
}

// WithNoMountRefresh disables mount refresh.
func WithNoMountRefresh() VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.MountRefreshInterval = nil
	}
}

// WithAutoSuspend duration without queries after which the virtual instance is suspended. Minimum time is 15 minutes.
func WithAutoSuspend(d time.Duration) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.AutoSuspend = &d
	}
}

// WithVirtualInstanceSize is used to optionally set the virtual instance size.
func WithVirtualInstanceSize(size VirtualInstanceSize) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		t := size.String()
		o.Size = &t
	}
}

// WithVirtualInstanceName is used to set a new name for the virtual instance.
func WithVirtualInstanceName(name string) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.Name = &name
	}
}

// WithVirtualInstanceDescription is used to set a description for the virtual instance.
func WithVirtualInstanceDescription(desc string) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.Description = &desc
	}
}

// TODO: once the openapi-generator generates a custom type for each enum the below two types can be replaced
// https://github.com/OpenAPITools/openapi-generator/issues/9567

type VirtualInstanceSize string

func (t VirtualInstanceSize) String() string {
	return string(t)
}

const (
	SizeFree     VirtualInstanceSize = "FREE"
	SizeNano     VirtualInstanceSize = "NANO"
	SizeShared   VirtualInstanceSize = "SHARED"
	SizeMilli    VirtualInstanceSize = "MILLI"
	SizeSmall    VirtualInstanceSize = "SMALL"
	SizeMedium   VirtualInstanceSize = "MEDIUM"
	SizeLarge    VirtualInstanceSize = "LARGE"
	SizeXLarge   VirtualInstanceSize = "XLARGE"
	SizeXLarge2  VirtualInstanceSize = "XLARGE2"
	SizeXLarge4  VirtualInstanceSize = "XLARGE4"
	SizeXLarge8  VirtualInstanceSize = "XLARGE8"
	SizeXLarge16 VirtualInstanceSize = "XLARGE16"
)
