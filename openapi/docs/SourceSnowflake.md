# SourceSnowflake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | **string** | Name of the snowflake database. | 
**Schema** | **string** | Name of the snowflake database schema. | 
**TableName** | **string** | Name of the snowflake table. | 
**Warehouse** | Pointer to **string** | Name of the data warehouse to be used. | [optional] 
**Status** | Pointer to [**StatusSnowflake**](StatusSnowflake.md) |  | [optional] 

## Methods

### NewSourceSnowflake

`func NewSourceSnowflake(database string, schema string, tableName string, ) *SourceSnowflake`

NewSourceSnowflake instantiates a new SourceSnowflake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceSnowflakeWithDefaults

`func NewSourceSnowflakeWithDefaults() *SourceSnowflake`

NewSourceSnowflakeWithDefaults instantiates a new SourceSnowflake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *SourceSnowflake) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *SourceSnowflake) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *SourceSnowflake) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetSchema

`func (o *SourceSnowflake) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SourceSnowflake) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SourceSnowflake) SetSchema(v string)`

SetSchema sets Schema field to given value.


### GetTableName

`func (o *SourceSnowflake) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *SourceSnowflake) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *SourceSnowflake) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetWarehouse

`func (o *SourceSnowflake) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *SourceSnowflake) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *SourceSnowflake) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *SourceSnowflake) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetStatus

`func (o *SourceSnowflake) GetStatus() StatusSnowflake`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceSnowflake) GetStatusOk() (*StatusSnowflake, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceSnowflake) SetStatus(v StatusSnowflake)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceSnowflake) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


