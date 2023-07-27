# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | Date that API key was created (ISO-8601 format). | [optional] 
**CreatedBy** | Pointer to **string** | Email of API key owner. | [optional] 
**CreatedByApikeyName** | Pointer to **string** | Name of the API key that was used to create this object if one was used. | [optional] 
**ExpiryTime** | Pointer to **string** | The expiration date of this API key. | [optional] 
**Key** | **string** | This field will only be populated with the full key when creating an API key. Otherwise, it will be an API key identifier of 6 characters. | 
**LastAccessTime** | Pointer to **string** | Date that API key was most recently used (ISO-8601 format). | [optional] 
**Name** | **string** | Name of the API key. | 
**Role** | Pointer to **string** | Role specifying access control. If not specified, API key will have access to all of the associated user&#39;s roles. | [optional] 
**State** | Pointer to **string** | Current state of this key. | [optional] 

## Methods

### NewApiKey

`func NewApiKey(key string, name string, ) *ApiKey`

NewApiKey instantiates a new ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ApiKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ApiKey) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApiKey) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApiKey) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ApiKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByApikeyName

`func (o *ApiKey) GetCreatedByApikeyName() string`

GetCreatedByApikeyName returns the CreatedByApikeyName field if non-nil, zero value otherwise.

### GetCreatedByApikeyNameOk

`func (o *ApiKey) GetCreatedByApikeyNameOk() (*string, bool)`

GetCreatedByApikeyNameOk returns a tuple with the CreatedByApikeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByApikeyName

`func (o *ApiKey) SetCreatedByApikeyName(v string)`

SetCreatedByApikeyName sets CreatedByApikeyName field to given value.

### HasCreatedByApikeyName

`func (o *ApiKey) HasCreatedByApikeyName() bool`

HasCreatedByApikeyName returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ApiKey) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ApiKey) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ApiKey) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ApiKey) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetKey

`func (o *ApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKey) SetKey(v string)`

SetKey sets Key field to given value.


### GetLastAccessTime

`func (o *ApiKey) GetLastAccessTime() string`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *ApiKey) GetLastAccessTimeOk() (*string, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *ApiKey) SetLastAccessTime(v string)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *ApiKey) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetName

`func (o *ApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKey) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *ApiKey) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApiKey) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApiKey) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApiKey) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetState

`func (o *ApiKey) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiKey) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiKey) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiKey) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


