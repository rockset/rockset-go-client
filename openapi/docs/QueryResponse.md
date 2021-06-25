# QueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | Pointer to **string** | unique id for this query | [optional] 
**Collections** | Pointer to **[]string** | list of collections queried by the query | [optional] 
**Results** | Pointer to **[]map[string]interface{}** | list of objects returned by the query | [optional] 
**Stats** | Pointer to [**QueryResponseStats**](QueryResponseStats.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | warnings received from the query | [optional] 
**QueryErrors** | Pointer to [**[]QueryError**](QueryError.md) | errors encountered while streaming the query | [optional] 
**ColumnFields** | Pointer to [**[]QueryFieldType**](QueryFieldType.md) | meta information about each column in the result set | [optional] 

## Methods

### NewQueryResponse

`func NewQueryResponse() *QueryResponse`

NewQueryResponse instantiates a new QueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseWithDefaults

`func NewQueryResponseWithDefaults() *QueryResponse`

NewQueryResponseWithDefaults instantiates a new QueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *QueryResponse) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *QueryResponse) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *QueryResponse) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *QueryResponse) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetCollections

`func (o *QueryResponse) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *QueryResponse) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *QueryResponse) SetCollections(v []string)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *QueryResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetResults

`func (o *QueryResponse) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryResponse) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryResponse) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *QueryResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStats

`func (o *QueryResponse) GetStats() QueryResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *QueryResponse) GetStatsOk() (*QueryResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *QueryResponse) SetStats(v QueryResponseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *QueryResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetWarnings

`func (o *QueryResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *QueryResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *QueryResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *QueryResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetQueryErrors

`func (o *QueryResponse) GetQueryErrors() []QueryError`

GetQueryErrors returns the QueryErrors field if non-nil, zero value otherwise.

### GetQueryErrorsOk

`func (o *QueryResponse) GetQueryErrorsOk() (*[]QueryError, bool)`

GetQueryErrorsOk returns a tuple with the QueryErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryErrors

`func (o *QueryResponse) SetQueryErrors(v []QueryError)`

SetQueryErrors sets QueryErrors field to given value.

### HasQueryErrors

`func (o *QueryResponse) HasQueryErrors() bool`

HasQueryErrors returns a boolean if a field has been set.

### GetColumnFields

`func (o *QueryResponse) GetColumnFields() []QueryFieldType`

GetColumnFields returns the ColumnFields field if non-nil, zero value otherwise.

### GetColumnFieldsOk

`func (o *QueryResponse) GetColumnFieldsOk() (*[]QueryFieldType, bool)`

GetColumnFieldsOk returns a tuple with the ColumnFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnFields

`func (o *QueryResponse) SetColumnFields(v []QueryFieldType)`

SetColumnFields sets ColumnFields field to given value.

### HasColumnFields

`func (o *QueryResponse) HasColumnFields() bool`

HasColumnFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


