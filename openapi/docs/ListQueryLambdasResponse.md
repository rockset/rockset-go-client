# ListQueryLambdasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]QueryLambda**](QueryLambda.md) | List of all Query Lambdas. | [optional] 

## Methods

### NewListQueryLambdasResponse

`func NewListQueryLambdasResponse() *ListQueryLambdasResponse`

NewListQueryLambdasResponse instantiates a new ListQueryLambdasResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListQueryLambdasResponseWithDefaults

`func NewListQueryLambdasResponseWithDefaults() *ListQueryLambdasResponse`

NewListQueryLambdasResponseWithDefaults instantiates a new ListQueryLambdasResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListQueryLambdasResponse) GetData() []QueryLambda`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListQueryLambdasResponse) GetDataOk() (*[]QueryLambda, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListQueryLambdasResponse) SetData(v []QueryLambda)`

SetData sets Data field to given value.

### HasData

`func (o *ListQueryLambdasResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


