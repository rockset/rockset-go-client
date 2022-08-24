# CsvParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstLineAsColumnNames** | Pointer to **bool** | If the first line in every object specifies the column names. | [optional] 
**Separator** | Pointer to **string** | A single character that is the column separator. | [optional] 
**Encoding** | Pointer to **string** | One of: UTF-8, ISO_8859_1, UTF-16. | [optional] 
**ColumnNames** | Pointer to **[]string** | Names of columns. | [optional] 
**ColumnTypes** | Pointer to **[]string** | Names of columns. | [optional] 
**QuoteChar** | Pointer to **string** | character within which a cell value is enclosed,null character if no such character, default is &#39;\&quot;&#39; | [optional] 
**EscapeChar** | Pointer to **string** | escape character removes any special meaning from the following character,default is &#39;\\&#39; | [optional] 

## Methods

### NewCsvParams

`func NewCsvParams() *CsvParams`

NewCsvParams instantiates a new CsvParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsvParamsWithDefaults

`func NewCsvParamsWithDefaults() *CsvParams`

NewCsvParamsWithDefaults instantiates a new CsvParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstLineAsColumnNames

`func (o *CsvParams) GetFirstLineAsColumnNames() bool`

GetFirstLineAsColumnNames returns the FirstLineAsColumnNames field if non-nil, zero value otherwise.

### GetFirstLineAsColumnNamesOk

`func (o *CsvParams) GetFirstLineAsColumnNamesOk() (*bool, bool)`

GetFirstLineAsColumnNamesOk returns a tuple with the FirstLineAsColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLineAsColumnNames

`func (o *CsvParams) SetFirstLineAsColumnNames(v bool)`

SetFirstLineAsColumnNames sets FirstLineAsColumnNames field to given value.

### HasFirstLineAsColumnNames

`func (o *CsvParams) HasFirstLineAsColumnNames() bool`

HasFirstLineAsColumnNames returns a boolean if a field has been set.

### GetSeparator

`func (o *CsvParams) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *CsvParams) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *CsvParams) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *CsvParams) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetEncoding

`func (o *CsvParams) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *CsvParams) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *CsvParams) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *CsvParams) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetColumnNames

`func (o *CsvParams) GetColumnNames() []string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *CsvParams) GetColumnNamesOk() (*[]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *CsvParams) SetColumnNames(v []string)`

SetColumnNames sets ColumnNames field to given value.

### HasColumnNames

`func (o *CsvParams) HasColumnNames() bool`

HasColumnNames returns a boolean if a field has been set.

### GetColumnTypes

`func (o *CsvParams) GetColumnTypes() []string`

GetColumnTypes returns the ColumnTypes field if non-nil, zero value otherwise.

### GetColumnTypesOk

`func (o *CsvParams) GetColumnTypesOk() (*[]string, bool)`

GetColumnTypesOk returns a tuple with the ColumnTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnTypes

`func (o *CsvParams) SetColumnTypes(v []string)`

SetColumnTypes sets ColumnTypes field to given value.

### HasColumnTypes

`func (o *CsvParams) HasColumnTypes() bool`

HasColumnTypes returns a boolean if a field has been set.

### GetQuoteChar

`func (o *CsvParams) GetQuoteChar() string`

GetQuoteChar returns the QuoteChar field if non-nil, zero value otherwise.

### GetQuoteCharOk

`func (o *CsvParams) GetQuoteCharOk() (*string, bool)`

GetQuoteCharOk returns a tuple with the QuoteChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteChar

`func (o *CsvParams) SetQuoteChar(v string)`

SetQuoteChar sets QuoteChar field to given value.

### HasQuoteChar

`func (o *CsvParams) HasQuoteChar() bool`

HasQuoteChar returns a boolean if a field has been set.

### GetEscapeChar

`func (o *CsvParams) GetEscapeChar() string`

GetEscapeChar returns the EscapeChar field if non-nil, zero value otherwise.

### GetEscapeCharOk

`func (o *CsvParams) GetEscapeCharOk() (*string, bool)`

GetEscapeCharOk returns a tuple with the EscapeChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscapeChar

`func (o *CsvParams) SetEscapeChar(v string)`

SetEscapeChar sets EscapeChar field to given value.

### HasEscapeChar

`func (o *CsvParams) HasEscapeChar() bool`

HasEscapeChar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


