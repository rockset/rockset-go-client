# UpdateIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureBlobStorage** | Pointer to [**AzureBlobStorageIntegration**](AzureBlobStorageIntegration.md) |  | [optional] 
**AzureEventHubs** | Pointer to [**AzureEventHubsIntegration**](AzureEventHubsIntegration.md) |  | [optional] 
**AzureServiceBus** | Pointer to [**AzureServiceBusIntegration**](AzureServiceBusIntegration.md) |  | [optional] 
**Description** | Pointer to **string** | Longer explanation for the integration. | [optional] 
**Dynamodb** | Pointer to [**DynamodbIntegration**](DynamodbIntegration.md) |  | [optional] 
**Gcs** | Pointer to [**GcsIntegration**](GcsIntegration.md) |  | [optional] 
**Kafka** | Pointer to [**KafkaIntegration**](KafkaIntegration.md) |  | [optional] 
**Kinesis** | Pointer to [**KinesisIntegration**](KinesisIntegration.md) |  | [optional] 
**Mongodb** | Pointer to [**MongoDbIntegration**](MongoDbIntegration.md) |  | [optional] 
**S3** | Pointer to [**S3Integration**](S3Integration.md) |  | [optional] 
**Snowflake** | Pointer to [**SnowflakeIntegration**](SnowflakeIntegration.md) |  | [optional] 

## Methods

### NewUpdateIntegrationRequest

`func NewUpdateIntegrationRequest() *UpdateIntegrationRequest`

NewUpdateIntegrationRequest instantiates a new UpdateIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationRequestWithDefaults

`func NewUpdateIntegrationRequestWithDefaults() *UpdateIntegrationRequest`

NewUpdateIntegrationRequestWithDefaults instantiates a new UpdateIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureBlobStorage

`func (o *UpdateIntegrationRequest) GetAzureBlobStorage() AzureBlobStorageIntegration`

GetAzureBlobStorage returns the AzureBlobStorage field if non-nil, zero value otherwise.

### GetAzureBlobStorageOk

`func (o *UpdateIntegrationRequest) GetAzureBlobStorageOk() (*AzureBlobStorageIntegration, bool)`

GetAzureBlobStorageOk returns a tuple with the AzureBlobStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBlobStorage

`func (o *UpdateIntegrationRequest) SetAzureBlobStorage(v AzureBlobStorageIntegration)`

SetAzureBlobStorage sets AzureBlobStorage field to given value.

### HasAzureBlobStorage

`func (o *UpdateIntegrationRequest) HasAzureBlobStorage() bool`

HasAzureBlobStorage returns a boolean if a field has been set.

### GetAzureEventHubs

`func (o *UpdateIntegrationRequest) GetAzureEventHubs() AzureEventHubsIntegration`

GetAzureEventHubs returns the AzureEventHubs field if non-nil, zero value otherwise.

### GetAzureEventHubsOk

`func (o *UpdateIntegrationRequest) GetAzureEventHubsOk() (*AzureEventHubsIntegration, bool)`

GetAzureEventHubsOk returns a tuple with the AzureEventHubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEventHubs

`func (o *UpdateIntegrationRequest) SetAzureEventHubs(v AzureEventHubsIntegration)`

SetAzureEventHubs sets AzureEventHubs field to given value.

### HasAzureEventHubs

`func (o *UpdateIntegrationRequest) HasAzureEventHubs() bool`

HasAzureEventHubs returns a boolean if a field has been set.

### GetAzureServiceBus

`func (o *UpdateIntegrationRequest) GetAzureServiceBus() AzureServiceBusIntegration`

GetAzureServiceBus returns the AzureServiceBus field if non-nil, zero value otherwise.

### GetAzureServiceBusOk

`func (o *UpdateIntegrationRequest) GetAzureServiceBusOk() (*AzureServiceBusIntegration, bool)`

GetAzureServiceBusOk returns a tuple with the AzureServiceBus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureServiceBus

`func (o *UpdateIntegrationRequest) SetAzureServiceBus(v AzureServiceBusIntegration)`

SetAzureServiceBus sets AzureServiceBus field to given value.

### HasAzureServiceBus

`func (o *UpdateIntegrationRequest) HasAzureServiceBus() bool`

HasAzureServiceBus returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIntegrationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIntegrationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIntegrationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIntegrationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDynamodb

`func (o *UpdateIntegrationRequest) GetDynamodb() DynamodbIntegration`

GetDynamodb returns the Dynamodb field if non-nil, zero value otherwise.

### GetDynamodbOk

`func (o *UpdateIntegrationRequest) GetDynamodbOk() (*DynamodbIntegration, bool)`

GetDynamodbOk returns a tuple with the Dynamodb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamodb

`func (o *UpdateIntegrationRequest) SetDynamodb(v DynamodbIntegration)`

SetDynamodb sets Dynamodb field to given value.

### HasDynamodb

`func (o *UpdateIntegrationRequest) HasDynamodb() bool`

HasDynamodb returns a boolean if a field has been set.

### GetGcs

`func (o *UpdateIntegrationRequest) GetGcs() GcsIntegration`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *UpdateIntegrationRequest) GetGcsOk() (*GcsIntegration, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *UpdateIntegrationRequest) SetGcs(v GcsIntegration)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *UpdateIntegrationRequest) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetKafka

`func (o *UpdateIntegrationRequest) GetKafka() KafkaIntegration`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *UpdateIntegrationRequest) GetKafkaOk() (*KafkaIntegration, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *UpdateIntegrationRequest) SetKafka(v KafkaIntegration)`

SetKafka sets Kafka field to given value.

### HasKafka

`func (o *UpdateIntegrationRequest) HasKafka() bool`

HasKafka returns a boolean if a field has been set.

### GetKinesis

`func (o *UpdateIntegrationRequest) GetKinesis() KinesisIntegration`

GetKinesis returns the Kinesis field if non-nil, zero value otherwise.

### GetKinesisOk

`func (o *UpdateIntegrationRequest) GetKinesisOk() (*KinesisIntegration, bool)`

GetKinesisOk returns a tuple with the Kinesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinesis

`func (o *UpdateIntegrationRequest) SetKinesis(v KinesisIntegration)`

SetKinesis sets Kinesis field to given value.

### HasKinesis

`func (o *UpdateIntegrationRequest) HasKinesis() bool`

HasKinesis returns a boolean if a field has been set.

### GetMongodb

`func (o *UpdateIntegrationRequest) GetMongodb() MongoDbIntegration`

GetMongodb returns the Mongodb field if non-nil, zero value otherwise.

### GetMongodbOk

`func (o *UpdateIntegrationRequest) GetMongodbOk() (*MongoDbIntegration, bool)`

GetMongodbOk returns a tuple with the Mongodb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodb

`func (o *UpdateIntegrationRequest) SetMongodb(v MongoDbIntegration)`

SetMongodb sets Mongodb field to given value.

### HasMongodb

`func (o *UpdateIntegrationRequest) HasMongodb() bool`

HasMongodb returns a boolean if a field has been set.

### GetS3

`func (o *UpdateIntegrationRequest) GetS3() S3Integration`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *UpdateIntegrationRequest) GetS3Ok() (*S3Integration, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *UpdateIntegrationRequest) SetS3(v S3Integration)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *UpdateIntegrationRequest) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetSnowflake

`func (o *UpdateIntegrationRequest) GetSnowflake() SnowflakeIntegration`

GetSnowflake returns the Snowflake field if non-nil, zero value otherwise.

### GetSnowflakeOk

`func (o *UpdateIntegrationRequest) GetSnowflakeOk() (*SnowflakeIntegration, bool)`

GetSnowflakeOk returns a tuple with the Snowflake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflake

`func (o *UpdateIntegrationRequest) SetSnowflake(v SnowflakeIntegration)`

SetSnowflake sets Snowflake field to given value.

### HasSnowflake

`func (o *UpdateIntegrationRequest) HasSnowflake() bool`

HasSnowflake returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


