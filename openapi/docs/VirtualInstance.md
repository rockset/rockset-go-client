# VirtualInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | virtual instance state | [optional] 
**CurrentType** | Pointer to **string** |  | [optional] 
**DesiredType** | Pointer to **string** |  | [optional] 
**CurrentSize** | Pointer to **string** | virtual instance current size | [optional] [readonly] 
**DesiredSize** | Pointer to **string** | virtual instance desired size | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | ISO-8601 date of when virtual instance size was last updated | [optional] 
**EstimatedSwitchDurationMinutes** | Pointer to **int64** | estimated duration in minutes of last virtual instance size update | [optional] 
**MonitoringEnabled** | Pointer to **bool** |  | [optional] 
**DefaultPodCount** | Pointer to **int32** |  | [optional] 
**ScaledPodCount** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** | unique identifier for virtual instance | [optional] 

## Methods

### NewVirtualInstance

`func NewVirtualInstance() *VirtualInstance`

NewVirtualInstance instantiates a new VirtualInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualInstanceWithDefaults

`func NewVirtualInstanceWithDefaults() *VirtualInstance`

NewVirtualInstanceWithDefaults instantiates a new VirtualInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VirtualInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VirtualInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCurrentType

`func (o *VirtualInstance) GetCurrentType() string`

GetCurrentType returns the CurrentType field if non-nil, zero value otherwise.

### GetCurrentTypeOk

`func (o *VirtualInstance) GetCurrentTypeOk() (*string, bool)`

GetCurrentTypeOk returns a tuple with the CurrentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentType

`func (o *VirtualInstance) SetCurrentType(v string)`

SetCurrentType sets CurrentType field to given value.

### HasCurrentType

`func (o *VirtualInstance) HasCurrentType() bool`

HasCurrentType returns a boolean if a field has been set.

### GetDesiredType

`func (o *VirtualInstance) GetDesiredType() string`

GetDesiredType returns the DesiredType field if non-nil, zero value otherwise.

### GetDesiredTypeOk

`func (o *VirtualInstance) GetDesiredTypeOk() (*string, bool)`

GetDesiredTypeOk returns a tuple with the DesiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredType

`func (o *VirtualInstance) SetDesiredType(v string)`

SetDesiredType sets DesiredType field to given value.

### HasDesiredType

`func (o *VirtualInstance) HasDesiredType() bool`

HasDesiredType returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VirtualInstance) GetCurrentSize() string`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VirtualInstance) GetCurrentSizeOk() (*string, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VirtualInstance) SetCurrentSize(v string)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VirtualInstance) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetDesiredSize

`func (o *VirtualInstance) GetDesiredSize() string`

GetDesiredSize returns the DesiredSize field if non-nil, zero value otherwise.

### GetDesiredSizeOk

`func (o *VirtualInstance) GetDesiredSizeOk() (*string, bool)`

GetDesiredSizeOk returns a tuple with the DesiredSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredSize

`func (o *VirtualInstance) SetDesiredSize(v string)`

SetDesiredSize sets DesiredSize field to given value.

### HasDesiredSize

`func (o *VirtualInstance) HasDesiredSize() bool`

HasDesiredSize returns a boolean if a field has been set.

### GetLastUpdated

`func (o *VirtualInstance) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VirtualInstance) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VirtualInstance) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VirtualInstance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEstimatedSwitchDurationMinutes

`func (o *VirtualInstance) GetEstimatedSwitchDurationMinutes() int64`

GetEstimatedSwitchDurationMinutes returns the EstimatedSwitchDurationMinutes field if non-nil, zero value otherwise.

### GetEstimatedSwitchDurationMinutesOk

`func (o *VirtualInstance) GetEstimatedSwitchDurationMinutesOk() (*int64, bool)`

GetEstimatedSwitchDurationMinutesOk returns a tuple with the EstimatedSwitchDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSwitchDurationMinutes

`func (o *VirtualInstance) SetEstimatedSwitchDurationMinutes(v int64)`

SetEstimatedSwitchDurationMinutes sets EstimatedSwitchDurationMinutes field to given value.

### HasEstimatedSwitchDurationMinutes

`func (o *VirtualInstance) HasEstimatedSwitchDurationMinutes() bool`

HasEstimatedSwitchDurationMinutes returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *VirtualInstance) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *VirtualInstance) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *VirtualInstance) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *VirtualInstance) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetDefaultPodCount

`func (o *VirtualInstance) GetDefaultPodCount() int32`

GetDefaultPodCount returns the DefaultPodCount field if non-nil, zero value otherwise.

### GetDefaultPodCountOk

`func (o *VirtualInstance) GetDefaultPodCountOk() (*int32, bool)`

GetDefaultPodCountOk returns a tuple with the DefaultPodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPodCount

`func (o *VirtualInstance) SetDefaultPodCount(v int32)`

SetDefaultPodCount sets DefaultPodCount field to given value.

### HasDefaultPodCount

`func (o *VirtualInstance) HasDefaultPodCount() bool`

HasDefaultPodCount returns a boolean if a field has been set.

### GetScaledPodCount

`func (o *VirtualInstance) GetScaledPodCount() int32`

GetScaledPodCount returns the ScaledPodCount field if non-nil, zero value otherwise.

### GetScaledPodCountOk

`func (o *VirtualInstance) GetScaledPodCountOk() (*int32, bool)`

GetScaledPodCountOk returns a tuple with the ScaledPodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaledPodCount

`func (o *VirtualInstance) SetScaledPodCount(v int32)`

SetScaledPodCount sets ScaledPodCount field to given value.

### HasScaledPodCount

`func (o *VirtualInstance) HasScaledPodCount() bool`

HasScaledPodCount returns a boolean if a field has been set.

### GetId

`func (o *VirtualInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualInstance) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


