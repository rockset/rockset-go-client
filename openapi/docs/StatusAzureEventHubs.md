# StatusAzureEventHubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastConsumedTime** | Pointer to **string** | Time at which the last document was consumed. | [optional] 
**NumDocumentsProcessed** | Pointer to **int64** | Number of documents consumed. | [optional] 
**Partitions** | Pointer to [**[]StatusAzureEventHubsPartition**](StatusAzureEventHubsPartition.md) | Status info per partition. | [optional] 
**State** | Pointer to **string** | State of the source. | [optional] 

## Methods

### NewStatusAzureEventHubs

`func NewStatusAzureEventHubs() *StatusAzureEventHubs`

NewStatusAzureEventHubs instantiates a new StatusAzureEventHubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusAzureEventHubsWithDefaults

`func NewStatusAzureEventHubsWithDefaults() *StatusAzureEventHubs`

NewStatusAzureEventHubsWithDefaults instantiates a new StatusAzureEventHubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastConsumedTime

`func (o *StatusAzureEventHubs) GetLastConsumedTime() string`

GetLastConsumedTime returns the LastConsumedTime field if non-nil, zero value otherwise.

### GetLastConsumedTimeOk

`func (o *StatusAzureEventHubs) GetLastConsumedTimeOk() (*string, bool)`

GetLastConsumedTimeOk returns a tuple with the LastConsumedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedTime

`func (o *StatusAzureEventHubs) SetLastConsumedTime(v string)`

SetLastConsumedTime sets LastConsumedTime field to given value.

### HasLastConsumedTime

`func (o *StatusAzureEventHubs) HasLastConsumedTime() bool`

HasLastConsumedTime returns a boolean if a field has been set.

### GetNumDocumentsProcessed

`func (o *StatusAzureEventHubs) GetNumDocumentsProcessed() int64`

GetNumDocumentsProcessed returns the NumDocumentsProcessed field if non-nil, zero value otherwise.

### GetNumDocumentsProcessedOk

`func (o *StatusAzureEventHubs) GetNumDocumentsProcessedOk() (*int64, bool)`

GetNumDocumentsProcessedOk returns a tuple with the NumDocumentsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDocumentsProcessed

`func (o *StatusAzureEventHubs) SetNumDocumentsProcessed(v int64)`

SetNumDocumentsProcessed sets NumDocumentsProcessed field to given value.

### HasNumDocumentsProcessed

`func (o *StatusAzureEventHubs) HasNumDocumentsProcessed() bool`

HasNumDocumentsProcessed returns a boolean if a field has been set.

### GetPartitions

`func (o *StatusAzureEventHubs) GetPartitions() []StatusAzureEventHubsPartition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *StatusAzureEventHubs) GetPartitionsOk() (*[]StatusAzureEventHubsPartition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *StatusAzureEventHubs) SetPartitions(v []StatusAzureEventHubsPartition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *StatusAzureEventHubs) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetState

`func (o *StatusAzureEventHubs) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StatusAzureEventHubs) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StatusAzureEventHubs) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StatusAzureEventHubs) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


