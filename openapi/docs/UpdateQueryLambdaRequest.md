# UpdateQueryLambdaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional description. | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Sql** | Pointer to [**QueryLambdaSql**](QueryLambdaSql.md) |  | [optional] 

## Methods

### NewUpdateQueryLambdaRequest

`func NewUpdateQueryLambdaRequest() *UpdateQueryLambdaRequest`

NewUpdateQueryLambdaRequest instantiates a new UpdateQueryLambdaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateQueryLambdaRequestWithDefaults

`func NewUpdateQueryLambdaRequestWithDefaults() *UpdateQueryLambdaRequest`

NewUpdateQueryLambdaRequestWithDefaults instantiates a new UpdateQueryLambdaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateQueryLambdaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateQueryLambdaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateQueryLambdaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateQueryLambdaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *UpdateQueryLambdaRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *UpdateQueryLambdaRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *UpdateQueryLambdaRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *UpdateQueryLambdaRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetSql

`func (o *UpdateQueryLambdaRequest) GetSql() QueryLambdaSql`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *UpdateQueryLambdaRequest) GetSqlOk() (*QueryLambdaSql, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *UpdateQueryLambdaRequest) SetSql(v QueryLambdaSql)`

SetSql sets Sql field to given value.

### HasSql

`func (o *UpdateQueryLambdaRequest) HasSql() bool`

HasSql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


