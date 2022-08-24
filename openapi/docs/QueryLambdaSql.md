# QueryLambdaSql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | SQL text. | 
**DefaultParameters** | Pointer to [**[]QueryParameter**](QueryParameter.md) | Default parameters for this Query Lambda. | [optional] 

## Methods

### NewQueryLambdaSql

`func NewQueryLambdaSql(query string, ) *QueryLambdaSql`

NewQueryLambdaSql instantiates a new QueryLambdaSql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLambdaSqlWithDefaults

`func NewQueryLambdaSqlWithDefaults() *QueryLambdaSql`

NewQueryLambdaSqlWithDefaults instantiates a new QueryLambdaSql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *QueryLambdaSql) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryLambdaSql) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryLambdaSql) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetDefaultParameters

`func (o *QueryLambdaSql) GetDefaultParameters() []QueryParameter`

GetDefaultParameters returns the DefaultParameters field if non-nil, zero value otherwise.

### GetDefaultParametersOk

`func (o *QueryLambdaSql) GetDefaultParametersOk() (*[]QueryParameter, bool)`

GetDefaultParametersOk returns a tuple with the DefaultParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameters

`func (o *QueryLambdaSql) SetDefaultParameters(v []QueryParameter)`

SetDefaultParameters sets DefaultParameters field to given value.

### HasDefaultParameters

`func (o *QueryLambdaSql) HasDefaultParameters() bool`

HasDefaultParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


