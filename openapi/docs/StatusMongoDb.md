# StatusMongoDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanStartTime** | Pointer to **string** | MongoDB scan start time. | [optional] 
**ScanEndTime** | Pointer to **string** | MongoDB scan end time. | [optional] 
**ScanRecordsProcessed** | Pointer to **int64** | Number of records inserted using scan. | [optional] 
**ScanTotalRecords** | Pointer to **int64** | Number of records in MongoDB table at time of scan. | [optional] 
**State** | Pointer to **string** | State of current ingest for this table. | [optional] 
**StreamLastInsertProcessedAt** | Pointer to **string** | ISO-8601 date when new insert from source was last processed. | [optional] 
**StreamLastUpdateProcessedAt** | Pointer to **string** | ISO-8601 date when update from source was last processed. | [optional] 
**StreamLastDeleteProcessedAt** | Pointer to **string** | ISO-8601 date when delete from source was last processed. | [optional] 
**StreamRecordsInserted** | Pointer to **int64** | Number of new records inserted using stream. | [optional] 
**StreamRecordsUpdated** | Pointer to **int64** | Number of new records updated using stream. | [optional] 
**StreamRecordsDeleted** | Pointer to **int64** | Number of new records deleted using stream. | [optional] 

## Methods

### NewStatusMongoDb

`func NewStatusMongoDb() *StatusMongoDb`

NewStatusMongoDb instantiates a new StatusMongoDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusMongoDbWithDefaults

`func NewStatusMongoDbWithDefaults() *StatusMongoDb`

NewStatusMongoDbWithDefaults instantiates a new StatusMongoDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanStartTime

`func (o *StatusMongoDb) GetScanStartTime() string`

GetScanStartTime returns the ScanStartTime field if non-nil, zero value otherwise.

### GetScanStartTimeOk

`func (o *StatusMongoDb) GetScanStartTimeOk() (*string, bool)`

GetScanStartTimeOk returns a tuple with the ScanStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStartTime

`func (o *StatusMongoDb) SetScanStartTime(v string)`

SetScanStartTime sets ScanStartTime field to given value.

### HasScanStartTime

`func (o *StatusMongoDb) HasScanStartTime() bool`

HasScanStartTime returns a boolean if a field has been set.

### GetScanEndTime

`func (o *StatusMongoDb) GetScanEndTime() string`

GetScanEndTime returns the ScanEndTime field if non-nil, zero value otherwise.

### GetScanEndTimeOk

`func (o *StatusMongoDb) GetScanEndTimeOk() (*string, bool)`

GetScanEndTimeOk returns a tuple with the ScanEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanEndTime

`func (o *StatusMongoDb) SetScanEndTime(v string)`

SetScanEndTime sets ScanEndTime field to given value.

### HasScanEndTime

`func (o *StatusMongoDb) HasScanEndTime() bool`

HasScanEndTime returns a boolean if a field has been set.

### GetScanRecordsProcessed

`func (o *StatusMongoDb) GetScanRecordsProcessed() int64`

GetScanRecordsProcessed returns the ScanRecordsProcessed field if non-nil, zero value otherwise.

### GetScanRecordsProcessedOk

`func (o *StatusMongoDb) GetScanRecordsProcessedOk() (*int64, bool)`

GetScanRecordsProcessedOk returns a tuple with the ScanRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanRecordsProcessed

`func (o *StatusMongoDb) SetScanRecordsProcessed(v int64)`

SetScanRecordsProcessed sets ScanRecordsProcessed field to given value.

### HasScanRecordsProcessed

`func (o *StatusMongoDb) HasScanRecordsProcessed() bool`

HasScanRecordsProcessed returns a boolean if a field has been set.

### GetScanTotalRecords

`func (o *StatusMongoDb) GetScanTotalRecords() int64`

GetScanTotalRecords returns the ScanTotalRecords field if non-nil, zero value otherwise.

### GetScanTotalRecordsOk

`func (o *StatusMongoDb) GetScanTotalRecordsOk() (*int64, bool)`

GetScanTotalRecordsOk returns a tuple with the ScanTotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTotalRecords

`func (o *StatusMongoDb) SetScanTotalRecords(v int64)`

SetScanTotalRecords sets ScanTotalRecords field to given value.

### HasScanTotalRecords

`func (o *StatusMongoDb) HasScanTotalRecords() bool`

HasScanTotalRecords returns a boolean if a field has been set.

### GetState

`func (o *StatusMongoDb) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StatusMongoDb) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StatusMongoDb) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StatusMongoDb) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreamLastInsertProcessedAt

`func (o *StatusMongoDb) GetStreamLastInsertProcessedAt() string`

GetStreamLastInsertProcessedAt returns the StreamLastInsertProcessedAt field if non-nil, zero value otherwise.

### GetStreamLastInsertProcessedAtOk

`func (o *StatusMongoDb) GetStreamLastInsertProcessedAtOk() (*string, bool)`

GetStreamLastInsertProcessedAtOk returns a tuple with the StreamLastInsertProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamLastInsertProcessedAt

`func (o *StatusMongoDb) SetStreamLastInsertProcessedAt(v string)`

SetStreamLastInsertProcessedAt sets StreamLastInsertProcessedAt field to given value.

### HasStreamLastInsertProcessedAt

`func (o *StatusMongoDb) HasStreamLastInsertProcessedAt() bool`

HasStreamLastInsertProcessedAt returns a boolean if a field has been set.

### GetStreamLastUpdateProcessedAt

`func (o *StatusMongoDb) GetStreamLastUpdateProcessedAt() string`

GetStreamLastUpdateProcessedAt returns the StreamLastUpdateProcessedAt field if non-nil, zero value otherwise.

### GetStreamLastUpdateProcessedAtOk

`func (o *StatusMongoDb) GetStreamLastUpdateProcessedAtOk() (*string, bool)`

GetStreamLastUpdateProcessedAtOk returns a tuple with the StreamLastUpdateProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamLastUpdateProcessedAt

`func (o *StatusMongoDb) SetStreamLastUpdateProcessedAt(v string)`

SetStreamLastUpdateProcessedAt sets StreamLastUpdateProcessedAt field to given value.

### HasStreamLastUpdateProcessedAt

`func (o *StatusMongoDb) HasStreamLastUpdateProcessedAt() bool`

HasStreamLastUpdateProcessedAt returns a boolean if a field has been set.

### GetStreamLastDeleteProcessedAt

`func (o *StatusMongoDb) GetStreamLastDeleteProcessedAt() string`

GetStreamLastDeleteProcessedAt returns the StreamLastDeleteProcessedAt field if non-nil, zero value otherwise.

### GetStreamLastDeleteProcessedAtOk

`func (o *StatusMongoDb) GetStreamLastDeleteProcessedAtOk() (*string, bool)`

GetStreamLastDeleteProcessedAtOk returns a tuple with the StreamLastDeleteProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamLastDeleteProcessedAt

`func (o *StatusMongoDb) SetStreamLastDeleteProcessedAt(v string)`

SetStreamLastDeleteProcessedAt sets StreamLastDeleteProcessedAt field to given value.

### HasStreamLastDeleteProcessedAt

`func (o *StatusMongoDb) HasStreamLastDeleteProcessedAt() bool`

HasStreamLastDeleteProcessedAt returns a boolean if a field has been set.

### GetStreamRecordsInserted

`func (o *StatusMongoDb) GetStreamRecordsInserted() int64`

GetStreamRecordsInserted returns the StreamRecordsInserted field if non-nil, zero value otherwise.

### GetStreamRecordsInsertedOk

`func (o *StatusMongoDb) GetStreamRecordsInsertedOk() (*int64, bool)`

GetStreamRecordsInsertedOk returns a tuple with the StreamRecordsInserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamRecordsInserted

`func (o *StatusMongoDb) SetStreamRecordsInserted(v int64)`

SetStreamRecordsInserted sets StreamRecordsInserted field to given value.

### HasStreamRecordsInserted

`func (o *StatusMongoDb) HasStreamRecordsInserted() bool`

HasStreamRecordsInserted returns a boolean if a field has been set.

### GetStreamRecordsUpdated

`func (o *StatusMongoDb) GetStreamRecordsUpdated() int64`

GetStreamRecordsUpdated returns the StreamRecordsUpdated field if non-nil, zero value otherwise.

### GetStreamRecordsUpdatedOk

`func (o *StatusMongoDb) GetStreamRecordsUpdatedOk() (*int64, bool)`

GetStreamRecordsUpdatedOk returns a tuple with the StreamRecordsUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamRecordsUpdated

`func (o *StatusMongoDb) SetStreamRecordsUpdated(v int64)`

SetStreamRecordsUpdated sets StreamRecordsUpdated field to given value.

### HasStreamRecordsUpdated

`func (o *StatusMongoDb) HasStreamRecordsUpdated() bool`

HasStreamRecordsUpdated returns a boolean if a field has been set.

### GetStreamRecordsDeleted

`func (o *StatusMongoDb) GetStreamRecordsDeleted() int64`

GetStreamRecordsDeleted returns the StreamRecordsDeleted field if non-nil, zero value otherwise.

### GetStreamRecordsDeletedOk

`func (o *StatusMongoDb) GetStreamRecordsDeletedOk() (*int64, bool)`

GetStreamRecordsDeletedOk returns a tuple with the StreamRecordsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamRecordsDeleted

`func (o *StatusMongoDb) SetStreamRecordsDeleted(v int64)`

SetStreamRecordsDeleted sets StreamRecordsDeleted field to given value.

### HasStreamRecordsDeleted

`func (o *StatusMongoDb) HasStreamRecordsDeleted() bool`

HasStreamRecordsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


