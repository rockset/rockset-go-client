# SourceRedshift

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | **string** | name of the database in Redshift Cluster | 
**Schema** | **string** | schema which contains the Redshift table | 
**TableName** | **string** | name of Redshift table containing data | 
**IncrementalField** | Pointer to **string** | field in Redshift source table to monitor for updates | [optional] 

## Methods

### NewSourceRedshift

`func NewSourceRedshift(database string, schema string, tableName string, ) *SourceRedshift`

NewSourceRedshift instantiates a new SourceRedshift object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceRedshiftWithDefaults

`func NewSourceRedshiftWithDefaults() *SourceRedshift`

NewSourceRedshiftWithDefaults instantiates a new SourceRedshift object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *SourceRedshift) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *SourceRedshift) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *SourceRedshift) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetSchema

`func (o *SourceRedshift) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SourceRedshift) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SourceRedshift) SetSchema(v string)`

SetSchema sets Schema field to given value.


### GetTableName

`func (o *SourceRedshift) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *SourceRedshift) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *SourceRedshift) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetIncrementalField

`func (o *SourceRedshift) GetIncrementalField() string`

GetIncrementalField returns the IncrementalField field if non-nil, zero value otherwise.

### GetIncrementalFieldOk

`func (o *SourceRedshift) GetIncrementalFieldOk() (*string, bool)`

GetIncrementalFieldOk returns a tuple with the IncrementalField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalField

`func (o *SourceRedshift) SetIncrementalField(v string)`

SetIncrementalField sets IncrementalField field to given value.

### HasIncrementalField

`func (o *SourceRedshift) HasIncrementalField() bool`

HasIncrementalField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


