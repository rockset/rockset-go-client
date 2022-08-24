# QueryLambdaTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | Pointer to **string** | Name of Query Lambda tag. | [optional] 
**Version** | Pointer to [**QueryLambdaVersion**](QueryLambdaVersion.md) |  | [optional] 

## Methods

### NewQueryLambdaTag

`func NewQueryLambdaTag() *QueryLambdaTag`

NewQueryLambdaTag instantiates a new QueryLambdaTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLambdaTagWithDefaults

`func NewQueryLambdaTagWithDefaults() *QueryLambdaTag`

NewQueryLambdaTagWithDefaults instantiates a new QueryLambdaTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagName

`func (o *QueryLambdaTag) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *QueryLambdaTag) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *QueryLambdaTag) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *QueryLambdaTag) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetVersion

`func (o *QueryLambdaTag) GetVersion() QueryLambdaVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryLambdaTag) GetVersionOk() (*QueryLambdaVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryLambdaTag) SetVersion(v QueryLambdaVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *QueryLambdaTag) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


