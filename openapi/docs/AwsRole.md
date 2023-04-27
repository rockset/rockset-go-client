# AwsRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsExternalId** | Pointer to **string** | External id used for integration. | [optional] 
**AwsRoleArn** | **string** | ARN of rockset-role created in your account. | 

## Methods

### NewAwsRole

`func NewAwsRole(awsRoleArn string, ) *AwsRole`

NewAwsRole instantiates a new AwsRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRoleWithDefaults

`func NewAwsRoleWithDefaults() *AwsRole`

NewAwsRoleWithDefaults instantiates a new AwsRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsExternalId

`func (o *AwsRole) GetAwsExternalId() string`

GetAwsExternalId returns the AwsExternalId field if non-nil, zero value otherwise.

### GetAwsExternalIdOk

`func (o *AwsRole) GetAwsExternalIdOk() (*string, bool)`

GetAwsExternalIdOk returns a tuple with the AwsExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalId

`func (o *AwsRole) SetAwsExternalId(v string)`

SetAwsExternalId sets AwsExternalId field to given value.

### HasAwsExternalId

`func (o *AwsRole) HasAwsExternalId() bool`

HasAwsExternalId returns a boolean if a field has been set.

### GetAwsRoleArn

`func (o *AwsRole) GetAwsRoleArn() string`

GetAwsRoleArn returns the AwsRoleArn field if non-nil, zero value otherwise.

### GetAwsRoleArnOk

`func (o *AwsRole) GetAwsRoleArnOk() (*string, bool)`

GetAwsRoleArnOk returns a tuple with the AwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleArn

`func (o *AwsRole) SetAwsRoleArn(v string)`

SetAwsRoleArn sets AwsRoleArn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


