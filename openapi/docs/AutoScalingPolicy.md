# AutoScalingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether auto scaling policy is enabled. | [optional] 
**MaxSize** | Pointer to **string** | Maximum size Rockset can auto scale the Virtual Instance to. This value should be one of the dedicated sizes greater than or same as the min_size and lower than or same as the current size. | [optional] 
**MinSize** | Pointer to **string** | Minimum size Rockset can auto scale the Virtual Instance to. This value should be one of the dedicated sizes lower than or same as the max_size and greater than or same as the current size. | [optional] 

## Methods

### NewAutoScalingPolicy

`func NewAutoScalingPolicy() *AutoScalingPolicy`

NewAutoScalingPolicy instantiates a new AutoScalingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingPolicyWithDefaults

`func NewAutoScalingPolicyWithDefaults() *AutoScalingPolicy`

NewAutoScalingPolicyWithDefaults instantiates a new AutoScalingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoScalingPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoScalingPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoScalingPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoScalingPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxSize

`func (o *AutoScalingPolicy) GetMaxSize() string`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *AutoScalingPolicy) GetMaxSizeOk() (*string, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *AutoScalingPolicy) SetMaxSize(v string)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *AutoScalingPolicy) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetMinSize

`func (o *AutoScalingPolicy) GetMinSize() string`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *AutoScalingPolicy) GetMinSizeOk() (*string, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *AutoScalingPolicy) SetMinSize(v string)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *AutoScalingPolicy) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


