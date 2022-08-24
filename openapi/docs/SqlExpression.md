# SqlExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | Pointer to **string** | The name of a sql function. | [optional] 

## Methods

### NewSqlExpression

`func NewSqlExpression() *SqlExpression`

NewSqlExpression instantiates a new SqlExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlExpressionWithDefaults

`func NewSqlExpressionWithDefaults() *SqlExpression`

NewSqlExpressionWithDefaults instantiates a new SqlExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSql

`func (o *SqlExpression) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqlExpression) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqlExpression) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *SqlExpression) HasSql() bool`

HasSql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


