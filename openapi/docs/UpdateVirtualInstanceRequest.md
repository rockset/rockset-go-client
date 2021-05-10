# UpdateVirtualInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSize** | Pointer to **string** | requested virtual instance size | [optional] 
**NewType** | Pointer to **string** |  | [optional] 
**MonitoringEnabled** | Pointer to **bool** |  | [optional] 

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

### GetNewType

`func (o *UpdateVirtualInstanceRequest) GetNewType() string`

GetNewType returns the NewType field if non-nil, zero value otherwise.

### GetNewTypeOk

`func (o *UpdateVirtualInstanceRequest) GetNewTypeOk() (*string, bool)`

GetNewTypeOk returns a tuple with the NewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewType

`func (o *UpdateVirtualInstanceRequest) SetNewType(v string)`

SetNewType sets NewType field to given value.

### HasNewType

`func (o *UpdateVirtualInstanceRequest) HasNewType() bool`

HasNewType returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


