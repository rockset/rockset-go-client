/*
REST API

Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateVirtualInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVirtualInstanceRequest{}

// UpdateVirtualInstanceRequest struct for UpdateVirtualInstanceRequest
type UpdateVirtualInstanceRequest struct {
	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`
	// Whether Query VI auto-suspend should be enabled for this Virtual Instance.
	AutoSuspendEnabled *bool `json:"auto_suspend_enabled,omitempty"`
	// Number of seconds without queries after which the Query VI is suspended
	AutoSuspendSeconds *int32 `json:"auto_suspend_seconds,omitempty"`
	// New virtual instance description.
	Description *string `json:"description,omitempty"`
	// When a Virtual Instance is resumed, it will remount all collections that were mounted when the Virtual Instance was suspended.
	EnableRemountOnResume *bool `json:"enable_remount_on_resume,omitempty"`
	// Virtual Instance Class. Use `MO_IL` for Memory Optimized and `GP_IL` for General Purpose instance class.
	InstanceClass *string `json:"instance_class,omitempty"`
	// DEPRECATED. Use `mount_type` instead. Number of seconds between data refreshes for mounts on this Virtual Instance. The only valid values are 0 and null. 0 means the data will be refreshed continuously and null means the data will never refresh.
	MountRefreshIntervalSeconds *int32 `json:"mount_refresh_interval_seconds,omitempty"`
	// The mount type of collections that this Virtual Instance will query. Live mounted collections stay up-to-date with the underlying collection in real-time. Static mounted collections do not stay up-to-date. See https://docs.rockset.com/documentation/docs/using-virtual-instances#virtual-instance-configuration
	MountType *string `json:"mount_type,omitempty"`
	// New virtual instance name.
	Name *string `json:"name,omitempty"`
	// Requested virtual instance size.
	NewSize *string `json:"new_size,omitempty"`
}

// NewUpdateVirtualInstanceRequest instantiates a new UpdateVirtualInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVirtualInstanceRequest() *UpdateVirtualInstanceRequest {
	this := UpdateVirtualInstanceRequest{}
	return &this
}

// NewUpdateVirtualInstanceRequestWithDefaults instantiates a new UpdateVirtualInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVirtualInstanceRequestWithDefaults() *UpdateVirtualInstanceRequest {
	this := UpdateVirtualInstanceRequest{}
	return &this
}

// GetAutoScalingPolicy returns the AutoScalingPolicy field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetAutoScalingPolicy() AutoScalingPolicy {
	if o == nil || IsNil(o.AutoScalingPolicy) {
		var ret AutoScalingPolicy
		return ret
	}
	return *o.AutoScalingPolicy
}

// GetAutoScalingPolicyOk returns a tuple with the AutoScalingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetAutoScalingPolicyOk() (*AutoScalingPolicy, bool) {
	if o == nil || IsNil(o.AutoScalingPolicy) {
		return nil, false
	}
	return o.AutoScalingPolicy, true
}

// HasAutoScalingPolicy returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasAutoScalingPolicy() bool {
	if o != nil && !IsNil(o.AutoScalingPolicy) {
		return true
	}

	return false
}

// SetAutoScalingPolicy gets a reference to the given AutoScalingPolicy and assigns it to the AutoScalingPolicy field.
func (o *UpdateVirtualInstanceRequest) SetAutoScalingPolicy(v AutoScalingPolicy) {
	o.AutoScalingPolicy = &v
}

// GetAutoSuspendEnabled returns the AutoSuspendEnabled field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetAutoSuspendEnabled() bool {
	if o == nil || IsNil(o.AutoSuspendEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoSuspendEnabled
}

// GetAutoSuspendEnabledOk returns a tuple with the AutoSuspendEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetAutoSuspendEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoSuspendEnabled) {
		return nil, false
	}
	return o.AutoSuspendEnabled, true
}

// HasAutoSuspendEnabled returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasAutoSuspendEnabled() bool {
	if o != nil && !IsNil(o.AutoSuspendEnabled) {
		return true
	}

	return false
}

// SetAutoSuspendEnabled gets a reference to the given bool and assigns it to the AutoSuspendEnabled field.
func (o *UpdateVirtualInstanceRequest) SetAutoSuspendEnabled(v bool) {
	o.AutoSuspendEnabled = &v
}

// GetAutoSuspendSeconds returns the AutoSuspendSeconds field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetAutoSuspendSeconds() int32 {
	if o == nil || IsNil(o.AutoSuspendSeconds) {
		var ret int32
		return ret
	}
	return *o.AutoSuspendSeconds
}

// GetAutoSuspendSecondsOk returns a tuple with the AutoSuspendSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetAutoSuspendSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoSuspendSeconds) {
		return nil, false
	}
	return o.AutoSuspendSeconds, true
}

// HasAutoSuspendSeconds returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasAutoSuspendSeconds() bool {
	if o != nil && !IsNil(o.AutoSuspendSeconds) {
		return true
	}

	return false
}

// SetAutoSuspendSeconds gets a reference to the given int32 and assigns it to the AutoSuspendSeconds field.
func (o *UpdateVirtualInstanceRequest) SetAutoSuspendSeconds(v int32) {
	o.AutoSuspendSeconds = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVirtualInstanceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnableRemountOnResume returns the EnableRemountOnResume field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetEnableRemountOnResume() bool {
	if o == nil || IsNil(o.EnableRemountOnResume) {
		var ret bool
		return ret
	}
	return *o.EnableRemountOnResume
}

// GetEnableRemountOnResumeOk returns a tuple with the EnableRemountOnResume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetEnableRemountOnResumeOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRemountOnResume) {
		return nil, false
	}
	return o.EnableRemountOnResume, true
}

// HasEnableRemountOnResume returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasEnableRemountOnResume() bool {
	if o != nil && !IsNil(o.EnableRemountOnResume) {
		return true
	}

	return false
}

// SetEnableRemountOnResume gets a reference to the given bool and assigns it to the EnableRemountOnResume field.
func (o *UpdateVirtualInstanceRequest) SetEnableRemountOnResume(v bool) {
	o.EnableRemountOnResume = &v
}

// GetInstanceClass returns the InstanceClass field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetInstanceClass() string {
	if o == nil || IsNil(o.InstanceClass) {
		var ret string
		return ret
	}
	return *o.InstanceClass
}

// GetInstanceClassOk returns a tuple with the InstanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetInstanceClassOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceClass) {
		return nil, false
	}
	return o.InstanceClass, true
}

// HasInstanceClass returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasInstanceClass() bool {
	if o != nil && !IsNil(o.InstanceClass) {
		return true
	}

	return false
}

// SetInstanceClass gets a reference to the given string and assigns it to the InstanceClass field.
func (o *UpdateVirtualInstanceRequest) SetInstanceClass(v string) {
	o.InstanceClass = &v
}

// GetMountRefreshIntervalSeconds returns the MountRefreshIntervalSeconds field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetMountRefreshIntervalSeconds() int32 {
	if o == nil || IsNil(o.MountRefreshIntervalSeconds) {
		var ret int32
		return ret
	}
	return *o.MountRefreshIntervalSeconds
}

// GetMountRefreshIntervalSecondsOk returns a tuple with the MountRefreshIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetMountRefreshIntervalSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MountRefreshIntervalSeconds) {
		return nil, false
	}
	return o.MountRefreshIntervalSeconds, true
}

// HasMountRefreshIntervalSeconds returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasMountRefreshIntervalSeconds() bool {
	if o != nil && !IsNil(o.MountRefreshIntervalSeconds) {
		return true
	}

	return false
}

// SetMountRefreshIntervalSeconds gets a reference to the given int32 and assigns it to the MountRefreshIntervalSeconds field.
func (o *UpdateVirtualInstanceRequest) SetMountRefreshIntervalSeconds(v int32) {
	o.MountRefreshIntervalSeconds = &v
}

// GetMountType returns the MountType field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetMountType() string {
	if o == nil || IsNil(o.MountType) {
		var ret string
		return ret
	}
	return *o.MountType
}

// GetMountTypeOk returns a tuple with the MountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetMountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MountType) {
		return nil, false
	}
	return o.MountType, true
}

// HasMountType returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasMountType() bool {
	if o != nil && !IsNil(o.MountType) {
		return true
	}

	return false
}

// SetMountType gets a reference to the given string and assigns it to the MountType field.
func (o *UpdateVirtualInstanceRequest) SetMountType(v string) {
	o.MountType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVirtualInstanceRequest) SetName(v string) {
	o.Name = &v
}

// GetNewSize returns the NewSize field value if set, zero value otherwise.
func (o *UpdateVirtualInstanceRequest) GetNewSize() string {
	if o == nil || IsNil(o.NewSize) {
		var ret string
		return ret
	}
	return *o.NewSize
}

// GetNewSizeOk returns a tuple with the NewSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVirtualInstanceRequest) GetNewSizeOk() (*string, bool) {
	if o == nil || IsNil(o.NewSize) {
		return nil, false
	}
	return o.NewSize, true
}

// HasNewSize returns a boolean if a field has been set.
func (o *UpdateVirtualInstanceRequest) HasNewSize() bool {
	if o != nil && !IsNil(o.NewSize) {
		return true
	}

	return false
}

// SetNewSize gets a reference to the given string and assigns it to the NewSize field.
func (o *UpdateVirtualInstanceRequest) SetNewSize(v string) {
	o.NewSize = &v
}

func (o UpdateVirtualInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVirtualInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoScalingPolicy) {
		toSerialize["auto_scaling_policy"] = o.AutoScalingPolicy
	}
	if !IsNil(o.AutoSuspendEnabled) {
		toSerialize["auto_suspend_enabled"] = o.AutoSuspendEnabled
	}
	if !IsNil(o.AutoSuspendSeconds) {
		toSerialize["auto_suspend_seconds"] = o.AutoSuspendSeconds
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EnableRemountOnResume) {
		toSerialize["enable_remount_on_resume"] = o.EnableRemountOnResume
	}
	if !IsNil(o.InstanceClass) {
		toSerialize["instance_class"] = o.InstanceClass
	}
	if !IsNil(o.MountRefreshIntervalSeconds) {
		toSerialize["mount_refresh_interval_seconds"] = o.MountRefreshIntervalSeconds
	}
	if !IsNil(o.MountType) {
		toSerialize["mount_type"] = o.MountType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NewSize) {
		toSerialize["new_size"] = o.NewSize
	}
	return toSerialize, nil
}

type NullableUpdateVirtualInstanceRequest struct {
	value *UpdateVirtualInstanceRequest
	isSet bool
}

func (v NullableUpdateVirtualInstanceRequest) Get() *UpdateVirtualInstanceRequest {
	return v.value
}

func (v *NullableUpdateVirtualInstanceRequest) Set(val *UpdateVirtualInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVirtualInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVirtualInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVirtualInstanceRequest(val *UpdateVirtualInstanceRequest) *NullableUpdateVirtualInstanceRequest {
	return &NullableUpdateVirtualInstanceRequest{value: val, isSet: true}
}

func (v NullableUpdateVirtualInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVirtualInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


