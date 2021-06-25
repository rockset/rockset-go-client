# SourceDynamoDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRegion** | Pointer to **string** | AWS region name of DynamoDB table, by default us-west-2 is used | [optional] 
**TableName** | **string** | name of DynamoDB table containing data | 
**Status** | Pointer to [**StatusDynamoDb**](StatusDynamoDb.md) |  | [optional] 
**Rcu** | Pointer to **int64** | Max RCU usage for scan | [optional] 

## Methods

### NewSourceDynamoDb

`func NewSourceDynamoDb(tableName string, ) *SourceDynamoDb`

NewSourceDynamoDb instantiates a new SourceDynamoDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceDynamoDbWithDefaults

`func NewSourceDynamoDbWithDefaults() *SourceDynamoDb`

NewSourceDynamoDbWithDefaults instantiates a new SourceDynamoDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRegion

`func (o *SourceDynamoDb) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *SourceDynamoDb) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *SourceDynamoDb) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *SourceDynamoDb) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetTableName

`func (o *SourceDynamoDb) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *SourceDynamoDb) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *SourceDynamoDb) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetStatus

`func (o *SourceDynamoDb) GetStatus() StatusDynamoDb`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceDynamoDb) GetStatusOk() (*StatusDynamoDb, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceDynamoDb) SetStatus(v StatusDynamoDb)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceDynamoDb) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRcu

`func (o *SourceDynamoDb) GetRcu() int64`

GetRcu returns the Rcu field if non-nil, zero value otherwise.

### GetRcuOk

`func (o *SourceDynamoDb) GetRcuOk() (*int64, bool)`

GetRcuOk returns a tuple with the Rcu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcu

`func (o *SourceDynamoDb) SetRcu(v int64)`

SetRcu sets Rcu field to given value.

### HasRcu

`func (o *SourceDynamoDb) HasRcu() bool`

HasRcu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


