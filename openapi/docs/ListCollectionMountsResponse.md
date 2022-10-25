# ListCollectionMountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CollectionMount**](CollectionMount.md) | List of all collection mounts. | [optional] 

## Methods

### NewListCollectionMountsResponse

`func NewListCollectionMountsResponse() *ListCollectionMountsResponse`

NewListCollectionMountsResponse instantiates a new ListCollectionMountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCollectionMountsResponseWithDefaults

`func NewListCollectionMountsResponseWithDefaults() *ListCollectionMountsResponse`

NewListCollectionMountsResponseWithDefaults instantiates a new ListCollectionMountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCollectionMountsResponse) GetData() []CollectionMount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCollectionMountsResponse) GetDataOk() (*[]CollectionMount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCollectionMountsResponse) SetData(v []CollectionMount)`

SetData sets Data field to given value.

### HasData

`func (o *ListCollectionMountsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


