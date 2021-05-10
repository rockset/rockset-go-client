# FormatParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | Pointer to **bool** | source data is in json format | [optional] 
**Csv** | Pointer to [**CsvParams**](CsvParams.md) |  | [optional] 
**Xml** | Pointer to [**XmlParams**](XmlParams.md) |  | [optional] 

## Methods

### NewFormatParams

`func NewFormatParams() *FormatParams`

NewFormatParams instantiates a new FormatParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormatParamsWithDefaults

`func NewFormatParamsWithDefaults() *FormatParams`

NewFormatParamsWithDefaults instantiates a new FormatParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *FormatParams) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *FormatParams) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *FormatParams) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *FormatParams) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetCsv

`func (o *FormatParams) GetCsv() CsvParams`

GetCsv returns the Csv field if non-nil, zero value otherwise.

### GetCsvOk

`func (o *FormatParams) GetCsvOk() (*CsvParams, bool)`

GetCsvOk returns a tuple with the Csv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsv

`func (o *FormatParams) SetCsv(v CsvParams)`

SetCsv sets Csv field to given value.

### HasCsv

`func (o *FormatParams) HasCsv() bool`

HasCsv returns a boolean if a field has been set.

### GetXml

`func (o *FormatParams) GetXml() XmlParams`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *FormatParams) GetXmlOk() (*XmlParams, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *FormatParams) SetXml(v XmlParams)`

SetXml sets Xml field to given value.

### HasXml

`func (o *FormatParams) HasXml() bool`

HasXml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


