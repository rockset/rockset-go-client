# MongoDbIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionUri** | **string** | MongoDB connection URI string | 

## Methods

### NewMongoDbIntegration

`func NewMongoDbIntegration(connectionUri string, ) *MongoDbIntegration`

NewMongoDbIntegration instantiates a new MongoDbIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDbIntegrationWithDefaults

`func NewMongoDbIntegrationWithDefaults() *MongoDbIntegration`

NewMongoDbIntegrationWithDefaults instantiates a new MongoDbIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionUri

`func (o *MongoDbIntegration) GetConnectionUri() string`

GetConnectionUri returns the ConnectionUri field if non-nil, zero value otherwise.

### GetConnectionUriOk

`func (o *MongoDbIntegration) GetConnectionUriOk() (*string, bool)`

GetConnectionUriOk returns a tuple with the ConnectionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUri

`func (o *MongoDbIntegration) SetConnectionUri(v string)`

SetConnectionUri sets ConnectionUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


