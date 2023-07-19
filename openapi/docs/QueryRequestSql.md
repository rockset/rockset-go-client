# QueryRequestSql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRowLimit** | Pointer to **int32** | Row limit to use. Limits specified in the query text will override this default. | [optional] 
**GenerateWarnings** | Pointer to **bool** | Flag to enable warnings. Warnings can help debug query issues but negatively affect performance. | [optional] 
**InitialPaginateResponseDocCount** | Pointer to **int32** | [DEPRECATED] Use &#x60;max_initial_results&#x60; instead. Number of documents to return in addition to paginating for this query call. Only relevant if &#x60;paginate&#x60; flag is also set. | [optional] 
**Parameters** | Pointer to [**[]QueryParameter**](QueryParameter.md) | List of named parameters. | [optional] 
**Query** | **string** | SQL query string. | 

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

### GetInitialPaginateResponseDocCount

`func (o *QueryRequestSql) GetInitialPaginateResponseDocCount() int32`

GetInitialPaginateResponseDocCount returns the InitialPaginateResponseDocCount field if non-nil, zero value otherwise.

### GetInitialPaginateResponseDocCountOk

`func (o *QueryRequestSql) GetInitialPaginateResponseDocCountOk() (*int32, bool)`

GetInitialPaginateResponseDocCountOk returns a tuple with the InitialPaginateResponseDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPaginateResponseDocCount

`func (o *QueryRequestSql) SetInitialPaginateResponseDocCount(v int32)`

SetInitialPaginateResponseDocCount sets InitialPaginateResponseDocCount field to given value.

### HasInitialPaginateResponseDocCount

`func (o *QueryRequestSql) HasInitialPaginateResponseDocCount() bool`

HasInitialPaginateResponseDocCount returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


