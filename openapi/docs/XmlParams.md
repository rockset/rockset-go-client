# XmlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootTag** | Pointer to **string** | tag until which xml is ignored | [optional] 
**Encoding** | Pointer to **string** | encoding in which data source is encoded | [optional] 
**DocTag** | Pointer to **string** | tags with which documents are identified | [optional] 
**ValueTag** | Pointer to **string** | tag used for the value when there are attributes in the element having no child | [optional] 
**AttributePrefix** | Pointer to **string** | tag to differentiate between attributes and elements | [optional] 

## Methods

### NewXmlParams

`func NewXmlParams() *XmlParams`

NewXmlParams instantiates a new XmlParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXmlParamsWithDefaults

`func NewXmlParamsWithDefaults() *XmlParams`

NewXmlParamsWithDefaults instantiates a new XmlParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootTag

`func (o *XmlParams) GetRootTag() string`

GetRootTag returns the RootTag field if non-nil, zero value otherwise.

### GetRootTagOk

`func (o *XmlParams) GetRootTagOk() (*string, bool)`

GetRootTagOk returns a tuple with the RootTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTag

`func (o *XmlParams) SetRootTag(v string)`

SetRootTag sets RootTag field to given value.

### HasRootTag

`func (o *XmlParams) HasRootTag() bool`

HasRootTag returns a boolean if a field has been set.

### GetEncoding

`func (o *XmlParams) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *XmlParams) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *XmlParams) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *XmlParams) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetDocTag

`func (o *XmlParams) GetDocTag() string`

GetDocTag returns the DocTag field if non-nil, zero value otherwise.

### GetDocTagOk

`func (o *XmlParams) GetDocTagOk() (*string, bool)`

GetDocTagOk returns a tuple with the DocTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocTag

`func (o *XmlParams) SetDocTag(v string)`

SetDocTag sets DocTag field to given value.

### HasDocTag

`func (o *XmlParams) HasDocTag() bool`

HasDocTag returns a boolean if a field has been set.

### GetValueTag

`func (o *XmlParams) GetValueTag() string`

GetValueTag returns the ValueTag field if non-nil, zero value otherwise.

### GetValueTagOk

`func (o *XmlParams) GetValueTagOk() (*string, bool)`

GetValueTagOk returns a tuple with the ValueTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTag

`func (o *XmlParams) SetValueTag(v string)`

SetValueTag sets ValueTag field to given value.

### HasValueTag

`func (o *XmlParams) HasValueTag() bool`

HasValueTag returns a boolean if a field has been set.

### GetAttributePrefix

`func (o *XmlParams) GetAttributePrefix() string`

GetAttributePrefix returns the AttributePrefix field if non-nil, zero value otherwise.

### GetAttributePrefixOk

`func (o *XmlParams) GetAttributePrefixOk() (*string, bool)`

GetAttributePrefixOk returns a tuple with the AttributePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributePrefix

`func (o *XmlParams) SetAttributePrefix(v string)`

SetAttributePrefix sets AttributePrefix field to given value.

### HasAttributePrefix

`func (o *XmlParams) HasAttributePrefix() bool`

HasAttributePrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


