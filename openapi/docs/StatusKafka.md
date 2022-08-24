# StatusKafka

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | State of the Kafka source. | [optional] 
**LastConsumedTime** | Pointer to **string** | Time at which the last document was consumed from Kafka. | [optional] 
**NumDocumentsProcessed** | Pointer to **int64** | Number of documents consumed by this Kafka topic. | [optional] 
**KafkaPartitions** | Pointer to [**[]StatusKafkaPartition**](StatusKafkaPartition.md) | Status info per partition. | [optional] 

## Methods

### NewStatusKafka

`func NewStatusKafka() *StatusKafka`

NewStatusKafka instantiates a new StatusKafka object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusKafkaWithDefaults

`func NewStatusKafkaWithDefaults() *StatusKafka`

NewStatusKafkaWithDefaults instantiates a new StatusKafka object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *StatusKafka) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StatusKafka) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StatusKafka) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StatusKafka) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLastConsumedTime

`func (o *StatusKafka) GetLastConsumedTime() string`

GetLastConsumedTime returns the LastConsumedTime field if non-nil, zero value otherwise.

### GetLastConsumedTimeOk

`func (o *StatusKafka) GetLastConsumedTimeOk() (*string, bool)`

GetLastConsumedTimeOk returns a tuple with the LastConsumedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedTime

`func (o *StatusKafka) SetLastConsumedTime(v string)`

SetLastConsumedTime sets LastConsumedTime field to given value.

### HasLastConsumedTime

`func (o *StatusKafka) HasLastConsumedTime() bool`

HasLastConsumedTime returns a boolean if a field has been set.

### GetNumDocumentsProcessed

`func (o *StatusKafka) GetNumDocumentsProcessed() int64`

GetNumDocumentsProcessed returns the NumDocumentsProcessed field if non-nil, zero value otherwise.

### GetNumDocumentsProcessedOk

`func (o *StatusKafka) GetNumDocumentsProcessedOk() (*int64, bool)`

GetNumDocumentsProcessedOk returns a tuple with the NumDocumentsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDocumentsProcessed

`func (o *StatusKafka) SetNumDocumentsProcessed(v int64)`

SetNumDocumentsProcessed sets NumDocumentsProcessed field to given value.

### HasNumDocumentsProcessed

`func (o *StatusKafka) HasNumDocumentsProcessed() bool`

HasNumDocumentsProcessed returns a boolean if a field has been set.

### GetKafkaPartitions

`func (o *StatusKafka) GetKafkaPartitions() []StatusKafkaPartition`

GetKafkaPartitions returns the KafkaPartitions field if non-nil, zero value otherwise.

### GetKafkaPartitionsOk

`func (o *StatusKafka) GetKafkaPartitionsOk() (*[]StatusKafkaPartition, bool)`

GetKafkaPartitionsOk returns a tuple with the KafkaPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaPartitions

`func (o *StatusKafka) SetKafkaPartitions(v []StatusKafkaPartition)`

SetKafkaPartitions sets KafkaPartitions field to given value.

### HasKafkaPartitions

`func (o *StatusKafka) HasKafkaPartitions() bool`

HasKafkaPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


