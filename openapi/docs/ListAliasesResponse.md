# ListAliasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Alias**](Alias.md) | List of all aliases. | [optional] 

## Methods

### NewListAliasesResponse

`func NewListAliasesResponse() *ListAliasesResponse`

NewListAliasesResponse instantiates a new ListAliasesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAliasesResponseWithDefaults

`func NewListAliasesResponseWithDefaults() *ListAliasesResponse`

NewListAliasesResponseWithDefaults instantiates a new ListAliasesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAliasesResponse) GetData() []Alias`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAliasesResponse) GetDataOk() (*[]Alias, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAliasesResponse) SetData(v []Alias)`

SetData sets Data field to given value.

### HasData

`func (o *ListAliasesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


