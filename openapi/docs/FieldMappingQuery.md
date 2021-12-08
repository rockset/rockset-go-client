# FieldMappingQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | Pointer to **string** | SELECT * EXCEPT (name), SHA256(name) AS name_anon FROM _input | [optional] 

## Methods

### NewFieldMappingQuery

`func NewFieldMappingQuery() *FieldMappingQuery`

NewFieldMappingQuery instantiates a new FieldMappingQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMappingQueryWithDefaults

`func NewFieldMappingQueryWithDefaults() *FieldMappingQuery`

NewFieldMappingQueryWithDefaults instantiates a new FieldMappingQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSql

`func (o *FieldMappingQuery) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *FieldMappingQuery) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *FieldMappingQuery) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *FieldMappingQuery) HasSql() bool`

HasSql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


