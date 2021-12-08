# ListApiKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ApiKey**](ApiKey.md) | List of API key objects. | [optional] 

## Methods

### NewListApiKeysResponse

`func NewListApiKeysResponse() *ListApiKeysResponse`

NewListApiKeysResponse instantiates a new ListApiKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiKeysResponseWithDefaults

`func NewListApiKeysResponseWithDefaults() *ListApiKeysResponse`

NewListApiKeysResponseWithDefaults instantiates a new ListApiKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListApiKeysResponse) GetData() []ApiKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListApiKeysResponse) GetDataOk() (*[]ApiKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListApiKeysResponse) SetData(v []ApiKey)`

SetData sets Data field to given value.

### HasData

`func (o *ListApiKeysResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


