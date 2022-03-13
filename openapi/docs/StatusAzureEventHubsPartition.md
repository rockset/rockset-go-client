# StatusAzureEventHubsPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartitionNumber** | Pointer to **int32** | The number of this partition | [optional] 
**PartitionOffset** | Pointer to **int64** | Latest offset of partition | [optional] 
**OffsetLag** | Pointer to **int64** | Per partition lag for offset | [optional] 

## Methods

### NewStatusAzureEventHubsPartition

`func NewStatusAzureEventHubsPartition() *StatusAzureEventHubsPartition`

NewStatusAzureEventHubsPartition instantiates a new StatusAzureEventHubsPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusAzureEventHubsPartitionWithDefaults

`func NewStatusAzureEventHubsPartitionWithDefaults() *StatusAzureEventHubsPartition`

NewStatusAzureEventHubsPartitionWithDefaults instantiates a new StatusAzureEventHubsPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitionNumber

`func (o *StatusAzureEventHubsPartition) GetPartitionNumber() int32`

GetPartitionNumber returns the PartitionNumber field if non-nil, zero value otherwise.

### GetPartitionNumberOk

`func (o *StatusAzureEventHubsPartition) GetPartitionNumberOk() (*int32, bool)`

GetPartitionNumberOk returns a tuple with the PartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionNumber

`func (o *StatusAzureEventHubsPartition) SetPartitionNumber(v int32)`

SetPartitionNumber sets PartitionNumber field to given value.

### HasPartitionNumber

`func (o *StatusAzureEventHubsPartition) HasPartitionNumber() bool`

HasPartitionNumber returns a boolean if a field has been set.

### GetPartitionOffset

`func (o *StatusAzureEventHubsPartition) GetPartitionOffset() int64`

GetPartitionOffset returns the PartitionOffset field if non-nil, zero value otherwise.

### GetPartitionOffsetOk

`func (o *StatusAzureEventHubsPartition) GetPartitionOffsetOk() (*int64, bool)`

GetPartitionOffsetOk returns a tuple with the PartitionOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionOffset

`func (o *StatusAzureEventHubsPartition) SetPartitionOffset(v int64)`

SetPartitionOffset sets PartitionOffset field to given value.

### HasPartitionOffset

`func (o *StatusAzureEventHubsPartition) HasPartitionOffset() bool`

HasPartitionOffset returns a boolean if a field has been set.

### GetOffsetLag

`func (o *StatusAzureEventHubsPartition) GetOffsetLag() int64`

GetOffsetLag returns the OffsetLag field if non-nil, zero value otherwise.

### GetOffsetLagOk

`func (o *StatusAzureEventHubsPartition) GetOffsetLagOk() (*int64, bool)`

GetOffsetLagOk returns a tuple with the OffsetLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetLag

`func (o *StatusAzureEventHubsPartition) SetOffsetLag(v int64)`

SetOffsetLag sets OffsetLag field to given value.

### HasOffsetLag

`func (o *StatusAzureEventHubsPartition) HasOffsetLag() bool`

HasOffsetLag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


