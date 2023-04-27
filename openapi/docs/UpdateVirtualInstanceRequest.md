# UpdateVirtualInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSuspendEnabled** | Pointer to **bool** | Whether auto-suspend should be enabled for this Virtual Instance. | [optional] 
**AutoSuspendSeconds** | Pointer to **int32** | Number of seconds without queries after which the VI is suspended | [optional] 
**Description** | Pointer to **string** | New virtual instance description. | [optional] 
**MonitoringEnabled** | Pointer to **bool** |  | [optional] 
**MountRefreshIntervalSeconds** | Pointer to **int32** | Number of seconds between data refreshes for mounts on this Virtual Instance | [optional] 
**Name** | Pointer to **string** | New virtual instance name. | [optional] 
**NewSize** | Pointer to **string** | Requested virtual instance size. | [optional] 

## Methods

### NewUpdateVirtualInstanceRequest

`func NewUpdateVirtualInstanceRequest() *UpdateVirtualInstanceRequest`

NewUpdateVirtualInstanceRequest instantiates a new UpdateVirtualInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVirtualInstanceRequestWithDefaults

`func NewUpdateVirtualInstanceRequestWithDefaults() *UpdateVirtualInstanceRequest`

NewUpdateVirtualInstanceRequestWithDefaults instantiates a new UpdateVirtualInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSuspendEnabled

`func (o *UpdateVirtualInstanceRequest) GetAutoSuspendEnabled() bool`

GetAutoSuspendEnabled returns the AutoSuspendEnabled field if non-nil, zero value otherwise.

### GetAutoSuspendEnabledOk

`func (o *UpdateVirtualInstanceRequest) GetAutoSuspendEnabledOk() (*bool, bool)`

GetAutoSuspendEnabledOk returns a tuple with the AutoSuspendEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSuspendEnabled

`func (o *UpdateVirtualInstanceRequest) SetAutoSuspendEnabled(v bool)`

SetAutoSuspendEnabled sets AutoSuspendEnabled field to given value.

### HasAutoSuspendEnabled

`func (o *UpdateVirtualInstanceRequest) HasAutoSuspendEnabled() bool`

HasAutoSuspendEnabled returns a boolean if a field has been set.

### GetAutoSuspendSeconds

`func (o *UpdateVirtualInstanceRequest) GetAutoSuspendSeconds() int32`

GetAutoSuspendSeconds returns the AutoSuspendSeconds field if non-nil, zero value otherwise.

### GetAutoSuspendSecondsOk

`func (o *UpdateVirtualInstanceRequest) GetAutoSuspendSecondsOk() (*int32, bool)`

GetAutoSuspendSecondsOk returns a tuple with the AutoSuspendSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSuspendSeconds

`func (o *UpdateVirtualInstanceRequest) SetAutoSuspendSeconds(v int32)`

SetAutoSuspendSeconds sets AutoSuspendSeconds field to given value.

### HasAutoSuspendSeconds

`func (o *UpdateVirtualInstanceRequest) HasAutoSuspendSeconds() bool`

HasAutoSuspendSeconds returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateVirtualInstanceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVirtualInstanceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVirtualInstanceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVirtualInstanceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *UpdateVirtualInstanceRequest) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *UpdateVirtualInstanceRequest) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *UpdateVirtualInstanceRequest) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *UpdateVirtualInstanceRequest) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetMountRefreshIntervalSeconds

`func (o *UpdateVirtualInstanceRequest) GetMountRefreshIntervalSeconds() int32`

GetMountRefreshIntervalSeconds returns the MountRefreshIntervalSeconds field if non-nil, zero value otherwise.

### GetMountRefreshIntervalSecondsOk

`func (o *UpdateVirtualInstanceRequest) GetMountRefreshIntervalSecondsOk() (*int32, bool)`

GetMountRefreshIntervalSecondsOk returns a tuple with the MountRefreshIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountRefreshIntervalSeconds

`func (o *UpdateVirtualInstanceRequest) SetMountRefreshIntervalSeconds(v int32)`

SetMountRefreshIntervalSeconds sets MountRefreshIntervalSeconds field to given value.

### HasMountRefreshIntervalSeconds

`func (o *UpdateVirtualInstanceRequest) HasMountRefreshIntervalSeconds() bool`

HasMountRefreshIntervalSeconds returns a boolean if a field has been set.

### GetName

`func (o *UpdateVirtualInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVirtualInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVirtualInstanceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVirtualInstanceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewSize

`func (o *UpdateVirtualInstanceRequest) GetNewSize() string`

GetNewSize returns the NewSize field if non-nil, zero value otherwise.

### GetNewSizeOk

`func (o *UpdateVirtualInstanceRequest) GetNewSizeOk() (*string, bool)`

GetNewSizeOk returns a tuple with the NewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSize

`func (o *UpdateVirtualInstanceRequest) SetNewSize(v string)`

SetNewSize sets NewSize field to given value.

### HasNewSize

`func (o *UpdateVirtualInstanceRequest) HasNewSize() bool`

HasNewSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


