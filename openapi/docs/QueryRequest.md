# QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | [**QueryRequestSql**](QueryRequestSql.md) |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


