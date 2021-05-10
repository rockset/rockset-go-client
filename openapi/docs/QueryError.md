# QueryError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of error | [optional] 
**Message** | Pointer to **string** | A message associated with the error, containing more information about it | [optional] 
**StatusCode** | Pointer to **int32** | The HTTP status code associated with this error, had it been sent as the response status code | [optional] 

## Methods

### NewQueryError

`func NewQueryError() *QueryError`

NewQueryError instantiates a new QueryError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryErrorWithDefaults

`func NewQueryErrorWithDefaults() *QueryError`

NewQueryErrorWithDefaults instantiates a new QueryError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *QueryError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QueryError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QueryError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QueryError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *QueryError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *QueryError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *QueryError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *QueryError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


