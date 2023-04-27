# ErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **int32** | Column where the error happened (if applicable). | [optional] 
**ErrorId** | Pointer to **string** | ID of the error. | [optional] 
**Line** | Pointer to **int32** | Line where the error happened (if applicable). | [optional] 
**Message** | Pointer to **string** | Descriptive message about the error. | [optional] 
**QueryId** | Pointer to **string** | ID of the query (if applicable). | [optional] 
**TraceId** | Pointer to **string** | Internal trace ID to help with debugging. | [optional] 
**Type** | Pointer to **string** | Category of the error. | [optional] 

## Methods

### NewErrorModel

`func NewErrorModel() *ErrorModel`

NewErrorModel instantiates a new ErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorModelWithDefaults

`func NewErrorModelWithDefaults() *ErrorModel`

NewErrorModelWithDefaults instantiates a new ErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *ErrorModel) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *ErrorModel) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *ErrorModel) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *ErrorModel) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetErrorId

`func (o *ErrorModel) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ErrorModel) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ErrorModel) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *ErrorModel) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetLine

`func (o *ErrorModel) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ErrorModel) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ErrorModel) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *ErrorModel) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetQueryId

`func (o *ErrorModel) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *ErrorModel) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *ErrorModel) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *ErrorModel) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetTraceId

`func (o *ErrorModel) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ErrorModel) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ErrorModel) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ErrorModel) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *ErrorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


