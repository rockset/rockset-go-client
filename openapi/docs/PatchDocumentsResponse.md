# PatchDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DocumentStatus**](DocumentStatus.md) |  | 
**LastOffset** | Pointer to **string** | A string representing the collection offset after completing the patch. | [optional] 

## Methods

### NewPatchDocumentsResponse

`func NewPatchDocumentsResponse(data []DocumentStatus, ) *PatchDocumentsResponse`

NewPatchDocumentsResponse instantiates a new PatchDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchDocumentsResponseWithDefaults

`func NewPatchDocumentsResponseWithDefaults() *PatchDocumentsResponse`

NewPatchDocumentsResponseWithDefaults instantiates a new PatchDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PatchDocumentsResponse) GetData() []DocumentStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatchDocumentsResponse) GetDataOk() (*[]DocumentStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatchDocumentsResponse) SetData(v []DocumentStatus)`

SetData sets Data field to given value.


### GetLastOffset

`func (o *PatchDocumentsResponse) GetLastOffset() string`

GetLastOffset returns the LastOffset field if non-nil, zero value otherwise.

### GetLastOffsetOk

`func (o *PatchDocumentsResponse) GetLastOffsetOk() (*string, bool)`

GetLastOffsetOk returns a tuple with the LastOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOffset

`func (o *PatchDocumentsResponse) SetLastOffset(v string)`

SetLastOffset sets LastOffset field to given value.

### HasLastOffset

`func (o *PatchDocumentsResponse) HasLastOffset() bool`

HasLastOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


