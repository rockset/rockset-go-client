# Stats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElapsedTimeMs** | Pointer to **int64** | Total execution time (including time queued) of the query, in milliseconds. | [optional] 
**ResultSetBytesSize** | Pointer to **int64** | Number of bytes in the query result set. Only populated if &#x60;status&#x60; is &#x60;COMPLETE&#x60;. Not populated for INSERT INTO queries. | [optional] 
**ResultSetDocumentCount** | Pointer to **int64** | Number of documents returned by the query. Only populated if &#x60;status&#x60; is &#x60;COMPLETE&#x60;. | [optional] 
**ThrottledTimeMs** | Pointer to **int64** | Time query spent queued, in milliseconds. | [optional] 

## Methods

### NewStats

`func NewStats() *Stats`

NewStats instantiates a new Stats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsWithDefaults

`func NewStatsWithDefaults() *Stats`

NewStatsWithDefaults instantiates a new Stats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsedTimeMs

`func (o *Stats) GetElapsedTimeMs() int64`

GetElapsedTimeMs returns the ElapsedTimeMs field if non-nil, zero value otherwise.

### GetElapsedTimeMsOk

`func (o *Stats) GetElapsedTimeMsOk() (*int64, bool)`

GetElapsedTimeMsOk returns a tuple with the ElapsedTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTimeMs

`func (o *Stats) SetElapsedTimeMs(v int64)`

SetElapsedTimeMs sets ElapsedTimeMs field to given value.

### HasElapsedTimeMs

`func (o *Stats) HasElapsedTimeMs() bool`

HasElapsedTimeMs returns a boolean if a field has been set.

### GetResultSetBytesSize

`func (o *Stats) GetResultSetBytesSize() int64`

GetResultSetBytesSize returns the ResultSetBytesSize field if non-nil, zero value otherwise.

### GetResultSetBytesSizeOk

`func (o *Stats) GetResultSetBytesSizeOk() (*int64, bool)`

GetResultSetBytesSizeOk returns a tuple with the ResultSetBytesSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSetBytesSize

`func (o *Stats) SetResultSetBytesSize(v int64)`

SetResultSetBytesSize sets ResultSetBytesSize field to given value.

### HasResultSetBytesSize

`func (o *Stats) HasResultSetBytesSize() bool`

HasResultSetBytesSize returns a boolean if a field has been set.

### GetResultSetDocumentCount

`func (o *Stats) GetResultSetDocumentCount() int64`

GetResultSetDocumentCount returns the ResultSetDocumentCount field if non-nil, zero value otherwise.

### GetResultSetDocumentCountOk

`func (o *Stats) GetResultSetDocumentCountOk() (*int64, bool)`

GetResultSetDocumentCountOk returns a tuple with the ResultSetDocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSetDocumentCount

`func (o *Stats) SetResultSetDocumentCount(v int64)`

SetResultSetDocumentCount sets ResultSetDocumentCount field to given value.

### HasResultSetDocumentCount

`func (o *Stats) HasResultSetDocumentCount() bool`

HasResultSetDocumentCount returns a boolean if a field has been set.

### GetThrottledTimeMs

`func (o *Stats) GetThrottledTimeMs() int64`

GetThrottledTimeMs returns the ThrottledTimeMs field if non-nil, zero value otherwise.

### GetThrottledTimeMsOk

`func (o *Stats) GetThrottledTimeMsOk() (*int64, bool)`

GetThrottledTimeMsOk returns a tuple with the ThrottledTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottledTimeMs

`func (o *Stats) SetThrottledTimeMs(v int64)`

SetThrottledTimeMs sets ThrottledTimeMs field to given value.

### HasThrottledTimeMs

`func (o *Stats) HasThrottledTimeMs() bool`

HasThrottledTimeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


