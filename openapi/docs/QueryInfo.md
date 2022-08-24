# QueryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | Pointer to **string** | Unique Query ID. | [optional] 
**Status** | Pointer to **string** | Status of the query. | [optional] 
**ExecutedBy** | Pointer to **string** | User ID who executed the query. | [optional] 
**SubmittedAt** | Pointer to **string** | Time (UTC) the query request was first received and queued for execution. | [optional] 
**ExpiresAt** | Pointer to **string** | Time (UTC) that query results expire. Only populated if &#x60;status&#x60; is &#x60;COMPLETE&#x60;. | [optional] 
**Stats** | Pointer to [**Stats**](Stats.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**LastOffset** | Pointer to **string** | The log offset that query results were written to in the destination collection. Only populated for INSERT INTO queries. | [optional] 
**QueryErrors** | Pointer to [**[]QueryError**](QueryError.md) | Errors encountered while executing the query. | [optional] 

## Methods

### NewQueryInfo

`func NewQueryInfo() *QueryInfo`

NewQueryInfo instantiates a new QueryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInfoWithDefaults

`func NewQueryInfoWithDefaults() *QueryInfo`

NewQueryInfoWithDefaults instantiates a new QueryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *QueryInfo) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *QueryInfo) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *QueryInfo) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *QueryInfo) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetStatus

`func (o *QueryInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExecutedBy

`func (o *QueryInfo) GetExecutedBy() string`

GetExecutedBy returns the ExecutedBy field if non-nil, zero value otherwise.

### GetExecutedByOk

`func (o *QueryInfo) GetExecutedByOk() (*string, bool)`

GetExecutedByOk returns a tuple with the ExecutedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedBy

`func (o *QueryInfo) SetExecutedBy(v string)`

SetExecutedBy sets ExecutedBy field to given value.

### HasExecutedBy

`func (o *QueryInfo) HasExecutedBy() bool`

HasExecutedBy returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *QueryInfo) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *QueryInfo) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *QueryInfo) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *QueryInfo) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *QueryInfo) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *QueryInfo) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *QueryInfo) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *QueryInfo) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStats

`func (o *QueryInfo) GetStats() Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *QueryInfo) GetStatsOk() (*Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *QueryInfo) SetStats(v Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *QueryInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPagination

`func (o *QueryInfo) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *QueryInfo) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *QueryInfo) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *QueryInfo) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLastOffset

`func (o *QueryInfo) GetLastOffset() string`

GetLastOffset returns the LastOffset field if non-nil, zero value otherwise.

### GetLastOffsetOk

`func (o *QueryInfo) GetLastOffsetOk() (*string, bool)`

GetLastOffsetOk returns a tuple with the LastOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOffset

`func (o *QueryInfo) SetLastOffset(v string)`

SetLastOffset sets LastOffset field to given value.

### HasLastOffset

`func (o *QueryInfo) HasLastOffset() bool`

HasLastOffset returns a boolean if a field has been set.

### GetQueryErrors

`func (o *QueryInfo) GetQueryErrors() []QueryError`

GetQueryErrors returns the QueryErrors field if non-nil, zero value otherwise.

### GetQueryErrorsOk

`func (o *QueryInfo) GetQueryErrorsOk() (*[]QueryError, bool)`

GetQueryErrorsOk returns a tuple with the QueryErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryErrors

`func (o *QueryInfo) SetQueryErrors(v []QueryError)`

SetQueryErrors sets QueryErrors field to given value.

### HasQueryErrors

`func (o *QueryInfo) HasQueryErrors() bool`

HasQueryErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


