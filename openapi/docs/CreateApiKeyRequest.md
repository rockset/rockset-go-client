# CreateApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | descriptive label | 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateApiKeyRequest

`func NewCreateApiKeyRequest(name string, ) *CreateApiKeyRequest`

NewCreateApiKeyRequest instantiates a new CreateApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyRequestWithDefaults

`func NewCreateApiKeyRequestWithDefaults() *CreateApiKeyRequest`

NewCreateApiKeyRequestWithDefaults instantiates a new CreateApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApiKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApiKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApiKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *CreateApiKeyRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateApiKeyRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateApiKeyRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateApiKeyRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


