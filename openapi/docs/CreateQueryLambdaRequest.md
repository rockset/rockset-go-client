# CreateQueryLambdaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Query Lambda name. | 
**Description** | Pointer to **string** | Optional description. | [optional] 
**Sql** | [**QueryLambdaSql**](QueryLambdaSql.md) |  | 

## Methods

### NewCreateQueryLambdaRequest

`func NewCreateQueryLambdaRequest(name string, sql QueryLambdaSql, ) *CreateQueryLambdaRequest`

NewCreateQueryLambdaRequest instantiates a new CreateQueryLambdaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQueryLambdaRequestWithDefaults

`func NewCreateQueryLambdaRequestWithDefaults() *CreateQueryLambdaRequest`

NewCreateQueryLambdaRequestWithDefaults instantiates a new CreateQueryLambdaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateQueryLambdaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateQueryLambdaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateQueryLambdaRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateQueryLambdaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateQueryLambdaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateQueryLambdaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateQueryLambdaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSql

`func (o *CreateQueryLambdaRequest) GetSql() QueryLambdaSql`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *CreateQueryLambdaRequest) GetSqlOk() (*QueryLambdaSql, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *CreateQueryLambdaRequest) SetSql(v QueryLambdaSql)`

SetSql sets Sql field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


