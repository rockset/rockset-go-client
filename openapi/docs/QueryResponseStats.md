# QueryResponseStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElapsedTimeMs** | Pointer to **int64** | query time in milliseconds | [optional] 
**ThrottledTimeMicros** | Pointer to **int64** | time query was throttled by admission control | [optional] 

## Methods

### NewQueryResponseStats

`func NewQueryResponseStats() *QueryResponseStats`

NewQueryResponseStats instantiates a new QueryResponseStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseStatsWithDefaults

`func NewQueryResponseStatsWithDefaults() *QueryResponseStats`

NewQueryResponseStatsWithDefaults instantiates a new QueryResponseStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsedTimeMs

`func (o *QueryResponseStats) GetElapsedTimeMs() int64`

GetElapsedTimeMs returns the ElapsedTimeMs field if non-nil, zero value otherwise.

### GetElapsedTimeMsOk

`func (o *QueryResponseStats) GetElapsedTimeMsOk() (*int64, bool)`

GetElapsedTimeMsOk returns a tuple with the ElapsedTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTimeMs

`func (o *QueryResponseStats) SetElapsedTimeMs(v int64)`

SetElapsedTimeMs sets ElapsedTimeMs field to given value.

### HasElapsedTimeMs

`func (o *QueryResponseStats) HasElapsedTimeMs() bool`

HasElapsedTimeMs returns a boolean if a field has been set.

### GetThrottledTimeMicros

`func (o *QueryResponseStats) GetThrottledTimeMicros() int64`

GetThrottledTimeMicros returns the ThrottledTimeMicros field if non-nil, zero value otherwise.

### GetThrottledTimeMicrosOk

`func (o *QueryResponseStats) GetThrottledTimeMicrosOk() (*int64, bool)`

GetThrottledTimeMicrosOk returns a tuple with the ThrottledTimeMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottledTimeMicros

`func (o *QueryResponseStats) SetThrottledTimeMicros(v int64)`

SetThrottledTimeMicros sets ThrottledTimeMicros field to given value.

### HasThrottledTimeMicros

`func (o *QueryResponseStats) HasThrottledTimeMicros() bool`

HasThrottledTimeMicros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


