# SourceMongoDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** | MongoDB database name containing this collection | 
**CollectionName** | **string** | MongoDB collection name | 
**Status** | Pointer to [**StatusMongoDb**](StatusMongoDb.md) |  | [optional] 

## Methods

### NewSourceMongoDb

`func NewSourceMongoDb(databaseName string, collectionName string, ) *SourceMongoDb`

NewSourceMongoDb instantiates a new SourceMongoDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceMongoDbWithDefaults

`func NewSourceMongoDbWithDefaults() *SourceMongoDb`

NewSourceMongoDbWithDefaults instantiates a new SourceMongoDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *SourceMongoDb) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *SourceMongoDb) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *SourceMongoDb) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetCollectionName

`func (o *SourceMongoDb) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *SourceMongoDb) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *SourceMongoDb) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetStatus

`func (o *SourceMongoDb) GetStatus() StatusMongoDb`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceMongoDb) GetStatusOk() (*StatusMongoDb, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceMongoDb) SetStatus(v StatusMongoDb)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceMongoDb) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


