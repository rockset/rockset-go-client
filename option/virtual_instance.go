package option

import "time"

type VirtualInstanceState string

func (v VirtualInstanceState) String() string { return string(v) }

type MountState string

func (m MountState) String() string { return string(m) }

const (
	VirtualInstanceInitializing           VirtualInstanceState = "INITIALIZING"
	VirtualInstanceProvisioningResources  VirtualInstanceState = "PROVISIONING_RESOURCES"
	VirtualInstanceRebalancingCollections VirtualInstanceState = "REBALANCING_COLLECTIONS"
	VirtualInstanceActive                 VirtualInstanceState = "ACTIVE"
	VirtualInstanceSuspending             VirtualInstanceState = "SUSPENDING"
	VirtualInstanceSuspended              VirtualInstanceState = "SUSPENDED"
	VirtualInstanceResuming               VirtualInstanceState = "RESUMING"
	VirtualInstanceDeleted                VirtualInstanceState = "DELETED"

	MountCreating             MountState = "CREATING"
	MountActive               MountState = "ACTIVE"
	MountRefreshing           MountState = "REFRESHING"
	MountExpired              MountState = "EXPIRED"
	MountDeleting             MountState = "DELETING"
	MountSwitchingRefreshType MountState = "SWITCHING_REFRESH_TYPE"
	MountSuspended            MountState = "SUSPENDED"
	MountSuspending           MountState = "SUSPENDING"
)

// VirtualInstanceOptions contains the optional settings for a virtual instance.
type VirtualInstanceOptions struct {
	Description           *string
	Size                  *string
	AutoSuspend           *time.Duration
	EnableRemountOnResume *bool
	Name                  *string
}

type VirtualInstanceOption func(*VirtualInstanceOptions)

// WithRemountOnResume sets remount on Virtual Instance resume.
func WithRemountOnResume(remount bool) VirtualInstanceOption {
	return func(o *VirtualInstanceOptions) {
		o.EnableRemountOnResume = &remount
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
