# StatusKafkaPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OffsetLag** | Pointer to **int64** | Per partition lag for offset. | [optional] 
**PartitionNumber** | Pointer to **int32** | The number of this partition. | [optional] 
**PartitionOffset** | Pointer to **int64** | Latest offset of partition. | [optional] 

## Methods

### NewStatusKafkaPartition

`func NewStatusKafkaPartition() *StatusKafkaPartition`

NewStatusKafkaPartition instantiates a new StatusKafkaPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusKafkaPartitionWithDefaults

`func NewStatusKafkaPartitionWithDefaults() *StatusKafkaPartition`

NewStatusKafkaPartitionWithDefaults instantiates a new StatusKafkaPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffsetLag

`func (o *StatusKafkaPartition) GetOffsetLag() int64`

GetOffsetLag returns the OffsetLag field if non-nil, zero value otherwise.

### GetOffsetLagOk

`func (o *StatusKafkaPartition) GetOffsetLagOk() (*int64, bool)`

GetOffsetLagOk returns a tuple with the OffsetLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetLag

`func (o *StatusKafkaPartition) SetOffsetLag(v int64)`

SetOffsetLag sets OffsetLag field to given value.

### HasOffsetLag

`func (o *StatusKafkaPartition) HasOffsetLag() bool`

HasOffsetLag returns a boolean if a field has been set.

### GetPartitionNumber

`func (o *StatusKafkaPartition) GetPartitionNumber() int32`

GetPartitionNumber returns the PartitionNumber field if non-nil, zero value otherwise.

### GetPartitionNumberOk

`func (o *StatusKafkaPartition) GetPartitionNumberOk() (*int32, bool)`

GetPartitionNumberOk returns a tuple with the PartitionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionNumber

`func (o *StatusKafkaPartition) SetPartitionNumber(v int32)`

SetPartitionNumber sets PartitionNumber field to given value.

### HasPartitionNumber

`func (o *StatusKafkaPartition) HasPartitionNumber() bool`

HasPartitionNumber returns a boolean if a field has been set.

### GetPartitionOffset

`func (o *StatusKafkaPartition) GetPartitionOffset() int64`

GetPartitionOffset returns the PartitionOffset field if non-nil, zero value otherwise.

### GetPartitionOffsetOk

`func (o *StatusKafkaPartition) GetPartitionOffsetOk() (*int64, bool)`

GetPartitionOffsetOk returns a tuple with the PartitionOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionOffset

`func (o *StatusKafkaPartition) SetPartitionOffset(v int64)`

SetPartitionOffset sets PartitionOffset field to given value.

### HasPartitionOffset

`func (o *StatusKafkaPartition) HasPartitionOffset() bool`

HasPartitionOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


