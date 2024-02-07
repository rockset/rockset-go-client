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

// checks if the VirtualInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualInstance{}

// VirtualInstance struct for VirtualInstance
type VirtualInstance struct {
	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`
	// Number of seconds without queries after which the VI is suspended
	AutoSuspendSeconds *int32 `json:"auto_suspend_seconds,omitempty"`
	// ISO-8601 date of when virtual instance was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Creator of requested virtual instance.
	CreatedBy *string `json:"created_by,omitempty"`
	// Virtual Instance Class. `MO_IL` represents Memory Optimized and `GP_IL` represents General Purpose instance class.
	CurrentInstanceClass *string `json:"current_instance_class,omitempty"`
	// Virtual instance current size.
	CurrentSize     *string `json:"current_size,omitempty"`
	DefaultPodCount *int32  `json:"default_pod_count,omitempty"`
	DefaultVi       *bool   `json:"default_vi,omitempty"`
	// Virtual instance description.
	Description *string `json:"description,omitempty"`
	// Virtual Instance Class.
	DesiredInstanceClass *string `json:"desired_instance_class,omitempty"`
	// Virtual instance desired size.
	DesiredSize *string `json:"desired_size,omitempty"`
	// When a Virtual Instance is resumed, it will remount all collections that were mounted when the Virtual Instance was suspended.
	EnableRemountOnResume *bool `json:"enable_remount_on_resume,omitempty"`
	// Unique identifier for virtual instance.
	Id                *string `json:"id,omitempty"`
	MonitoringEnabled *bool   `json:"monitoring_enabled,omitempty"`
	// DEPRECATED. Number of seconds between data refreshes for mounts on this Virtual Instance
	MountRefreshIntervalSeconds *int32 `json:"mount_refresh_interval_seconds,omitempty"`
	// The mount type of collections that this Virtual Instance will query. Live mounted collections stay up-to-date with the underlying collection in real-time. Static mounted collections do not stay up-to-date. See https://docs.rockset.com/documentation/docs/virtual-instances#virtual-instance-configuration
	MountType *string `json:"mount_type,omitempty"`
	// Virtual instance name.
	Name string `json:"name"`
	// ISO-8601 date of when virtual instance was created.
	ResumedAt *string `json:"resumed_at,omitempty"`
	// Virtual Instance RRN.
	Rrn            *string `json:"rrn,omitempty"`
	ScaledPodCount *int32  `json:"scaled_pod_count,omitempty"`
	// Virtual instance state.
	State *string               `json:"state,omitempty"`
	Stats *VirtualInstanceStats `json:"stats,omitempty"`
}

// NewVirtualInstance instantiates a new VirtualInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualInstance(name string) *VirtualInstance {
	this := VirtualInstance{}
	this.Name = name
	return &this
}

// NewVirtualInstanceWithDefaults instantiates a new VirtualInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualInstanceWithDefaults() *VirtualInstance {
	this := VirtualInstance{}
	return &this
}

// GetAutoScalingPolicy returns the AutoScalingPolicy field value if set, zero value otherwise.
func (o *VirtualInstance) GetAutoScalingPolicy() AutoScalingPolicy {
	if o == nil || IsNil(o.AutoScalingPolicy) {
		var ret AutoScalingPolicy
		return ret
	}
	return *o.AutoScalingPolicy
}

// GetAutoScalingPolicyOk returns a tuple with the AutoScalingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetAutoScalingPolicyOk() (*AutoScalingPolicy, bool) {
	if o == nil || IsNil(o.AutoScalingPolicy) {
		return nil, false
	}
	return o.AutoScalingPolicy, true
}

// HasAutoScalingPolicy returns a boolean if a field has been set.
func (o *VirtualInstance) HasAutoScalingPolicy() bool {
	if o != nil && !IsNil(o.AutoScalingPolicy) {
		return true
	}

	return false
}

// SetAutoScalingPolicy gets a reference to the given AutoScalingPolicy and assigns it to the AutoScalingPolicy field.
func (o *VirtualInstance) SetAutoScalingPolicy(v AutoScalingPolicy) {
	o.AutoScalingPolicy = &v
}

// GetAutoSuspendSeconds returns the AutoSuspendSeconds field value if set, zero value otherwise.
func (o *VirtualInstance) GetAutoSuspendSeconds() int32 {
	if o == nil || IsNil(o.AutoSuspendSeconds) {
		var ret int32
		return ret
	}
	return *o.AutoSuspendSeconds
}

// GetAutoSuspendSecondsOk returns a tuple with the AutoSuspendSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetAutoSuspendSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoSuspendSeconds) {
		return nil, false
	}
	return o.AutoSuspendSeconds, true
}

// HasAutoSuspendSeconds returns a boolean if a field has been set.
func (o *VirtualInstance) HasAutoSuspendSeconds() bool {
	if o != nil && !IsNil(o.AutoSuspendSeconds) {
		return true
	}

	return false
}

// SetAutoSuspendSeconds gets a reference to the given int32 and assigns it to the AutoSuspendSeconds field.
func (o *VirtualInstance) SetAutoSuspendSeconds(v int32) {
	o.AutoSuspendSeconds = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VirtualInstance) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VirtualInstance) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *VirtualInstance) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *VirtualInstance) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *VirtualInstance) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *VirtualInstance) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCurrentInstanceClass returns the CurrentInstanceClass field value if set, zero value otherwise.
func (o *VirtualInstance) GetCurrentInstanceClass() string {
	if o == nil || IsNil(o.CurrentInstanceClass) {
		var ret string
		return ret
	}
	return *o.CurrentInstanceClass
}

// GetCurrentInstanceClassOk returns a tuple with the CurrentInstanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetCurrentInstanceClassOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentInstanceClass) {
		return nil, false
	}
	return o.CurrentInstanceClass, true
}

// HasCurrentInstanceClass returns a boolean if a field has been set.
func (o *VirtualInstance) HasCurrentInstanceClass() bool {
	if o != nil && !IsNil(o.CurrentInstanceClass) {
		return true
	}

	return false
}

// SetCurrentInstanceClass gets a reference to the given string and assigns it to the CurrentInstanceClass field.
func (o *VirtualInstance) SetCurrentInstanceClass(v string) {
	o.CurrentInstanceClass = &v
}

// GetCurrentSize returns the CurrentSize field value if set, zero value otherwise.
func (o *VirtualInstance) GetCurrentSize() string {
	if o == nil || IsNil(o.CurrentSize) {
		var ret string
		return ret
	}
	return *o.CurrentSize
}

// GetCurrentSizeOk returns a tuple with the CurrentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetCurrentSizeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentSize) {
		return nil, false
	}
	return o.CurrentSize, true
}

// HasCurrentSize returns a boolean if a field has been set.
func (o *VirtualInstance) HasCurrentSize() bool {
	if o != nil && !IsNil(o.CurrentSize) {
		return true
	}

	return false
}

// SetCurrentSize gets a reference to the given string and assigns it to the CurrentSize field.
func (o *VirtualInstance) SetCurrentSize(v string) {
	o.CurrentSize = &v
}

// GetDefaultPodCount returns the DefaultPodCount field value if set, zero value otherwise.
func (o *VirtualInstance) GetDefaultPodCount() int32 {
	if o == nil || IsNil(o.DefaultPodCount) {
		var ret int32
		return ret
	}
	return *o.DefaultPodCount
}

// GetDefaultPodCountOk returns a tuple with the DefaultPodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetDefaultPodCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultPodCount) {
		return nil, false
	}
	return o.DefaultPodCount, true
}

// HasDefaultPodCount returns a boolean if a field has been set.
func (o *VirtualInstance) HasDefaultPodCount() bool {
	if o != nil && !IsNil(o.DefaultPodCount) {
		return true
	}

	return false
}

// SetDefaultPodCount gets a reference to the given int32 and assigns it to the DefaultPodCount field.
func (o *VirtualInstance) SetDefaultPodCount(v int32) {
	o.DefaultPodCount = &v
}

// GetDefaultVi returns the DefaultVi field value if set, zero value otherwise.
func (o *VirtualInstance) GetDefaultVi() bool {
	if o == nil || IsNil(o.DefaultVi) {
		var ret bool
		return ret
	}
	return *o.DefaultVi
}

// GetDefaultViOk returns a tuple with the DefaultVi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetDefaultViOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultVi) {
		return nil, false
	}
	return o.DefaultVi, true
}

// HasDefaultVi returns a boolean if a field has been set.
func (o *VirtualInstance) HasDefaultVi() bool {
	if o != nil && !IsNil(o.DefaultVi) {
		return true
	}

	return false
}

// SetDefaultVi gets a reference to the given bool and assigns it to the DefaultVi field.
func (o *VirtualInstance) SetDefaultVi(v bool) {
	o.DefaultVi = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualInstance) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualInstance) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualInstance) SetDescription(v string) {
	o.Description = &v
}

// GetDesiredInstanceClass returns the DesiredInstanceClass field value if set, zero value otherwise.
func (o *VirtualInstance) GetDesiredInstanceClass() string {
	if o == nil || IsNil(o.DesiredInstanceClass) {
		var ret string
		return ret
	}
	return *o.DesiredInstanceClass
}

// GetDesiredInstanceClassOk returns a tuple with the DesiredInstanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetDesiredInstanceClassOk() (*string, bool) {
	if o == nil || IsNil(o.DesiredInstanceClass) {
		return nil, false
	}
	return o.DesiredInstanceClass, true
}

// HasDesiredInstanceClass returns a boolean if a field has been set.
func (o *VirtualInstance) HasDesiredInstanceClass() bool {
	if o != nil && !IsNil(o.DesiredInstanceClass) {
		return true
	}

	return false
}

// SetDesiredInstanceClass gets a reference to the given string and assigns it to the DesiredInstanceClass field.
func (o *VirtualInstance) SetDesiredInstanceClass(v string) {
	o.DesiredInstanceClass = &v
}

// GetDesiredSize returns the DesiredSize field value if set, zero value otherwise.
func (o *VirtualInstance) GetDesiredSize() string {
	if o == nil || IsNil(o.DesiredSize) {
		var ret string
		return ret
	}
	return *o.DesiredSize
}

// GetDesiredSizeOk returns a tuple with the DesiredSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetDesiredSizeOk() (*string, bool) {
	if o == nil || IsNil(o.DesiredSize) {
		return nil, false
	}
	return o.DesiredSize, true
}

// HasDesiredSize returns a boolean if a field has been set.
func (o *VirtualInstance) HasDesiredSize() bool {
	if o != nil && !IsNil(o.DesiredSize) {
		return true
	}

	return false
}

// SetDesiredSize gets a reference to the given string and assigns it to the DesiredSize field.
func (o *VirtualInstance) SetDesiredSize(v string) {
	o.DesiredSize = &v
}

// GetEnableRemountOnResume returns the EnableRemountOnResume field value if set, zero value otherwise.
func (o *VirtualInstance) GetEnableRemountOnResume() bool {
	if o == nil || IsNil(o.EnableRemountOnResume) {
		var ret bool
		return ret
	}
	return *o.EnableRemountOnResume
}

// GetEnableRemountOnResumeOk returns a tuple with the EnableRemountOnResume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetEnableRemountOnResumeOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRemountOnResume) {
		return nil, false
	}
	return o.EnableRemountOnResume, true
}

// HasEnableRemountOnResume returns a boolean if a field has been set.
func (o *VirtualInstance) HasEnableRemountOnResume() bool {
	if o != nil && !IsNil(o.EnableRemountOnResume) {
		return true
	}

	return false
}

// SetEnableRemountOnResume gets a reference to the given bool and assigns it to the EnableRemountOnResume field.
func (o *VirtualInstance) SetEnableRemountOnResume(v bool) {
	o.EnableRemountOnResume = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VirtualInstance) SetId(v string) {
	o.Id = &v
}

// GetMonitoringEnabled returns the MonitoringEnabled field value if set, zero value otherwise.
func (o *VirtualInstance) GetMonitoringEnabled() bool {
	if o == nil || IsNil(o.MonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.MonitoringEnabled
}

// GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MonitoringEnabled) {
		return nil, false
	}
	return o.MonitoringEnabled, true
}

// HasMonitoringEnabled returns a boolean if a field has been set.
func (o *VirtualInstance) HasMonitoringEnabled() bool {
	if o != nil && !IsNil(o.MonitoringEnabled) {
		return true
	}

	return false
}

// SetMonitoringEnabled gets a reference to the given bool and assigns it to the MonitoringEnabled field.
func (o *VirtualInstance) SetMonitoringEnabled(v bool) {
	o.MonitoringEnabled = &v
}

// GetMountRefreshIntervalSeconds returns the MountRefreshIntervalSeconds field value if set, zero value otherwise.
func (o *VirtualInstance) GetMountRefreshIntervalSeconds() int32 {
	if o == nil || IsNil(o.MountRefreshIntervalSeconds) {
		var ret int32
		return ret
	}
	return *o.MountRefreshIntervalSeconds
}

// GetMountRefreshIntervalSecondsOk returns a tuple with the MountRefreshIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetMountRefreshIntervalSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MountRefreshIntervalSeconds) {
		return nil, false
	}
	return o.MountRefreshIntervalSeconds, true
}

// HasMountRefreshIntervalSeconds returns a boolean if a field has been set.
func (o *VirtualInstance) HasMountRefreshIntervalSeconds() bool {
	if o != nil && !IsNil(o.MountRefreshIntervalSeconds) {
		return true
	}

	return false
}

// SetMountRefreshIntervalSeconds gets a reference to the given int32 and assigns it to the MountRefreshIntervalSeconds field.
func (o *VirtualInstance) SetMountRefreshIntervalSeconds(v int32) {
	o.MountRefreshIntervalSeconds = &v
}

// GetMountType returns the MountType field value if set, zero value otherwise.
func (o *VirtualInstance) GetMountType() string {
	if o == nil || IsNil(o.MountType) {
		var ret string
		return ret
	}
	return *o.MountType
}

// GetMountTypeOk returns a tuple with the MountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetMountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MountType) {
		return nil, false
	}
	return o.MountType, true
}

// HasMountType returns a boolean if a field has been set.
func (o *VirtualInstance) HasMountType() bool {
	if o != nil && !IsNil(o.MountType) {
		return true
	}

	return false
}

// SetMountType gets a reference to the given string and assigns it to the MountType field.
func (o *VirtualInstance) SetMountType(v string) {
	o.MountType = &v
}

// GetName returns the Name field value
func (o *VirtualInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualInstance) SetName(v string) {
	o.Name = v
}

// GetResumedAt returns the ResumedAt field value if set, zero value otherwise.
func (o *VirtualInstance) GetResumedAt() string {
	if o == nil || IsNil(o.ResumedAt) {
		var ret string
		return ret
	}
	return *o.ResumedAt
}

// GetResumedAtOk returns a tuple with the ResumedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetResumedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ResumedAt) {
		return nil, false
	}
	return o.ResumedAt, true
}

// HasResumedAt returns a boolean if a field has been set.
func (o *VirtualInstance) HasResumedAt() bool {
	if o != nil && !IsNil(o.ResumedAt) {
		return true
	}

	return false
}

// SetResumedAt gets a reference to the given string and assigns it to the ResumedAt field.
func (o *VirtualInstance) SetResumedAt(v string) {
	o.ResumedAt = &v
}

// GetRrn returns the Rrn field value if set, zero value otherwise.
func (o *VirtualInstance) GetRrn() string {
	if o == nil || IsNil(o.Rrn) {
		var ret string
		return ret
	}
	return *o.Rrn
}

// GetRrnOk returns a tuple with the Rrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetRrnOk() (*string, bool) {
	if o == nil || IsNil(o.Rrn) {
		return nil, false
	}
	return o.Rrn, true
}

// HasRrn returns a boolean if a field has been set.
func (o *VirtualInstance) HasRrn() bool {
	if o != nil && !IsNil(o.Rrn) {
		return true
	}

	return false
}

// SetRrn gets a reference to the given string and assigns it to the Rrn field.
func (o *VirtualInstance) SetRrn(v string) {
	o.Rrn = &v
}

// GetScaledPodCount returns the ScaledPodCount field value if set, zero value otherwise.
func (o *VirtualInstance) GetScaledPodCount() int32 {
	if o == nil || IsNil(o.ScaledPodCount) {
		var ret int32
		return ret
	}
	return *o.ScaledPodCount
}

// GetScaledPodCountOk returns a tuple with the ScaledPodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetScaledPodCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ScaledPodCount) {
		return nil, false
	}
	return o.ScaledPodCount, true
}

// HasScaledPodCount returns a boolean if a field has been set.
func (o *VirtualInstance) HasScaledPodCount() bool {
	if o != nil && !IsNil(o.ScaledPodCount) {
		return true
	}

	return false
}

// SetScaledPodCount gets a reference to the given int32 and assigns it to the ScaledPodCount field.
func (o *VirtualInstance) SetScaledPodCount(v int32) {
	o.ScaledPodCount = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VirtualInstance) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VirtualInstance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VirtualInstance) SetState(v string) {
	o.State = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *VirtualInstance) GetStats() VirtualInstanceStats {
	if o == nil || IsNil(o.Stats) {
		var ret VirtualInstanceStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstance) GetStatsOk() (*VirtualInstanceStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *VirtualInstance) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given VirtualInstanceStats and assigns it to the Stats field.
func (o *VirtualInstance) SetStats(v VirtualInstanceStats) {
	o.Stats = &v
}

func (o VirtualInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoScalingPolicy) {
		toSerialize["auto_scaling_policy"] = o.AutoScalingPolicy
	}
	if !IsNil(o.AutoSuspendSeconds) {
		toSerialize["auto_suspend_seconds"] = o.AutoSuspendSeconds
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.CurrentInstanceClass) {
		toSerialize["current_instance_class"] = o.CurrentInstanceClass
	}
	// skip: current_size is readOnly
	if !IsNil(o.DefaultPodCount) {
		toSerialize["default_pod_count"] = o.DefaultPodCount
	}
	if !IsNil(o.DefaultVi) {
		toSerialize["default_vi"] = o.DefaultVi
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DesiredInstanceClass) {
		toSerialize["desired_instance_class"] = o.DesiredInstanceClass
	}
	// skip: desired_size is readOnly
	if !IsNil(o.EnableRemountOnResume) {
		toSerialize["enable_remount_on_resume"] = o.EnableRemountOnResume
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MonitoringEnabled) {
		toSerialize["monitoring_enabled"] = o.MonitoringEnabled
	}
	if !IsNil(o.MountRefreshIntervalSeconds) {
		toSerialize["mount_refresh_interval_seconds"] = o.MountRefreshIntervalSeconds
	}
	if !IsNil(o.MountType) {
		toSerialize["mount_type"] = o.MountType
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ResumedAt) {
		toSerialize["resumed_at"] = o.ResumedAt
	}
	if !IsNil(o.Rrn) {
		toSerialize["rrn"] = o.Rrn
	}
	if !IsNil(o.ScaledPodCount) {
		toSerialize["scaled_pod_count"] = o.ScaledPodCount
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	return toSerialize, nil
}

type NullableVirtualInstance struct {
	value *VirtualInstance
	isSet bool
}

func (v NullableVirtualInstance) Get() *VirtualInstance {
	return v.value
}

func (v *NullableVirtualInstance) Set(val *VirtualInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualInstance(val *VirtualInstance) *NullableVirtualInstance {
	return &NullableVirtualInstance{value: val, isSet: true}
}

func (v NullableVirtualInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
