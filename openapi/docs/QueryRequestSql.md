# QueryRequestSql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**[]QueryParameter**](QueryParameter.md) | list of named parameters | [optional] 
**Query** | **string** | SQL query as a string | 
**DefaultRowLimit** | Pointer to **int32** | Row limit to use if no limit specified in the query | [optional] 
**GenerateWarnings** | Pointer to **bool** | Whether to generate warnings | [optional] 
**ProfilingEnabled** | Pointer to **bool** | Whether to generate a performance profile for this query | [optional] 

## Methods

### NewQueryRequestSql

`func NewQueryRequestSql(query string, ) *QueryRequestSql`

NewQueryRequestSql instantiates a new QueryRequestSql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestSqlWithDefaults

`func NewQueryRequestSqlWithDefaults() *QueryRequestSql`

NewQueryRequestSqlWithDefaults instantiates a new QueryRequestSql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *QueryRequestSql) GetParameters() []QueryParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *QueryRequestSql) GetParametersOk() (*[]QueryParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *QueryRequestSql) SetParameters(v []QueryParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *QueryRequestSql) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetQuery

`func (o *QueryRequestSql) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryRequestSql) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryRequestSql) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetDefaultRowLimit

`func (o *QueryRequestSql) GetDefaultRowLimit() int32`

GetDefaultRowLimit returns the DefaultRowLimit field if non-nil, zero value otherwise.

### GetDefaultRowLimitOk

`func (o *QueryRequestSql) GetDefaultRowLimitOk() (*int32, bool)`

GetDefaultRowLimitOk returns a tuple with the DefaultRowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRowLimit

`func (o *QueryRequestSql) SetDefaultRowLimit(v int32)`

SetDefaultRowLimit sets DefaultRowLimit field to given value.

### HasDefaultRowLimit

`func (o *QueryRequestSql) HasDefaultRowLimit() bool`

HasDefaultRowLimit returns a boolean if a field has been set.

### GetGenerateWarnings

`func (o *QueryRequestSql) GetGenerateWarnings() bool`

GetGenerateWarnings returns the GenerateWarnings field if non-nil, zero value otherwise.

### GetGenerateWarningsOk

`func (o *QueryRequestSql) GetGenerateWarningsOk() (*bool, bool)`

GetGenerateWarningsOk returns a tuple with the GenerateWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateWarnings

`func (o *QueryRequestSql) SetGenerateWarnings(v bool)`

SetGenerateWarnings sets GenerateWarnings field to given value.

### HasGenerateWarnings

`func (o *QueryRequestSql) HasGenerateWarnings() bool`

HasGenerateWarnings returns a boolean if a field has been set.

### GetProfilingEnabled

`func (o *QueryRequestSql) GetProfilingEnabled() bool`

GetProfilingEnabled returns the ProfilingEnabled field if non-nil, zero value otherwise.

### GetProfilingEnabledOk

`func (o *QueryRequestSql) GetProfilingEnabledOk() (*bool, bool)`

GetProfilingEnabledOk returns a tuple with the ProfilingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilingEnabled

`func (o *QueryRequestSql) SetProfilingEnabled(v bool)`

SetProfilingEnabled sets ProfilingEnabled field to given value.

### HasProfilingEnabled

`func (o *QueryRequestSql) HasProfilingEnabled() bool`

HasProfilingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


