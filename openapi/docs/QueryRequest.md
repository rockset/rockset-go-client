# QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | [**QueryRequestSql**](QueryRequestSql.md) |  | 
**AsyncOptions** | Pointer to [**AsyncQueryOptions**](AsyncQueryOptions.md) |  | [optional] 

## Methods

### NewQueryRequest

`func NewQueryRequest(sql QueryRequestSql, ) *QueryRequest`

NewQueryRequest instantiates a new QueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestWithDefaults

`func NewQueryRequestWithDefaults() *QueryRequest`

NewQueryRequestWithDefaults instantiates a new QueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSql

`func (o *QueryRequest) GetSql() QueryRequestSql`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *QueryRequest) GetSqlOk() (*QueryRequestSql, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *QueryRequest) SetSql(v QueryRequestSql)`

SetSql sets Sql field to given value.


### GetAsyncOptions

`func (o *QueryRequest) GetAsyncOptions() AsyncQueryOptions`

GetAsyncOptions returns the AsyncOptions field if non-nil, zero value otherwise.

### GetAsyncOptionsOk

`func (o *QueryRequest) GetAsyncOptionsOk() (*AsyncQueryOptions, bool)`

GetAsyncOptionsOk returns a tuple with the AsyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncOptions

`func (o *QueryRequest) SetAsyncOptions(v AsyncQueryOptions)`

SetAsyncOptions sets AsyncOptions field to given value.

### HasAsyncOptions

`func (o *QueryRequest) HasAsyncOptions() bool`

HasAsyncOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


