# UpdateCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Updated text describing the collection. | [optional] 
**FieldMappingQuery** | Pointer to [**FieldMappingQuery**](FieldMappingQuery.md) |  | [optional] 

## Methods

### NewUpdateCollectionRequest

`func NewUpdateCollectionRequest() *UpdateCollectionRequest`

NewUpdateCollectionRequest instantiates a new UpdateCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCollectionRequestWithDefaults

`func NewUpdateCollectionRequestWithDefaults() *UpdateCollectionRequest`

NewUpdateCollectionRequestWithDefaults instantiates a new UpdateCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateCollectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCollectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCollectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCollectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldMappingQuery

`func (o *UpdateCollectionRequest) GetFieldMappingQuery() FieldMappingQuery`

GetFieldMappingQuery returns the FieldMappingQuery field if non-nil, zero value otherwise.

### GetFieldMappingQueryOk

`func (o *UpdateCollectionRequest) GetFieldMappingQueryOk() (*FieldMappingQuery, bool)`

GetFieldMappingQueryOk returns a tuple with the FieldMappingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappingQuery

`func (o *UpdateCollectionRequest) SetFieldMappingQuery(v FieldMappingQuery)`

SetFieldMappingQuery sets FieldMappingQuery field to given value.

### HasFieldMappingQuery

`func (o *UpdateCollectionRequest) HasFieldMappingQuery() bool`

HasFieldMappingQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


