# CreateQueryLambdaTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | **string** | name of Query Lambda tag | 
**Version** | **string** | hash identifying a Query Lambda tag | 

## Methods

### NewCreateQueryLambdaTagRequest

`func NewCreateQueryLambdaTagRequest(tagName string, version string, ) *CreateQueryLambdaTagRequest`

NewCreateQueryLambdaTagRequest instantiates a new CreateQueryLambdaTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQueryLambdaTagRequestWithDefaults

`func NewCreateQueryLambdaTagRequestWithDefaults() *CreateQueryLambdaTagRequest`

NewCreateQueryLambdaTagRequestWithDefaults instantiates a new CreateQueryLambdaTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagName

`func (o *CreateQueryLambdaTagRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *CreateQueryLambdaTagRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *CreateQueryLambdaTagRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetVersion

`func (o *CreateQueryLambdaTagRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateQueryLambdaTagRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateQueryLambdaTagRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


