# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | descriptive label and unique identifier | 
**Description** | Pointer to **string** | longer explanation for the integration | [optional] 
**Collections** | Pointer to [**[]Collection**](Collection.md) | list of collections that use the integration | [optional] 
**CreatedBy** | **string** | email of user who created the integration | 
**CreatedAt** | Pointer to **string** | ISO-8601 date | [optional] 
**S3** | Pointer to [**S3Integration**](S3Integration.md) |  | [optional] 
**Kinesis** | Pointer to [**KinesisIntegration**](KinesisIntegration.md) |  | [optional] 
**Dynamodb** | Pointer to [**DynamodbIntegration**](DynamodbIntegration.md) |  | [optional] 
**Redshift** | Pointer to [**RedshiftIntegration**](RedshiftIntegration.md) |  | [optional] 
**Gcs** | Pointer to [**GcsIntegration**](GcsIntegration.md) |  | [optional] 
**Segment** | Pointer to [**SegmentIntegration**](SegmentIntegration.md) |  | [optional] 
**Kafka** | Pointer to [**KafkaIntegration**](KafkaIntegration.md) |  | [optional] 
**Mongodb** | Pointer to [**MongoDbIntegration**](MongoDbIntegration.md) |  | [optional] 

## Methods

### NewIntegration

`func NewIntegration(name string, createdBy string, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Integration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Integration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Integration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Integration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCollections

`func (o *Integration) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *Integration) GetCollectionsOk() (*[]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *Integration) SetCollections(v []Collection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *Integration) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Integration) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Integration) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Integration) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Integration) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Integration) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Integration) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Integration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetS3

`func (o *Integration) GetS3() S3Integration`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *Integration) GetS3Ok() (*S3Integration, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *Integration) SetS3(v S3Integration)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *Integration) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetKinesis

`func (o *Integration) GetKinesis() KinesisIntegration`

GetKinesis returns the Kinesis field if non-nil, zero value otherwise.

### GetKinesisOk

`func (o *Integration) GetKinesisOk() (*KinesisIntegration, bool)`

GetKinesisOk returns a tuple with the Kinesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinesis

`func (o *Integration) SetKinesis(v KinesisIntegration)`

SetKinesis sets Kinesis field to given value.

### HasKinesis

`func (o *Integration) HasKinesis() bool`

HasKinesis returns a boolean if a field has been set.

### GetDynamodb

`func (o *Integration) GetDynamodb() DynamodbIntegration`

GetDynamodb returns the Dynamodb field if non-nil, zero value otherwise.

### GetDynamodbOk

`func (o *Integration) GetDynamodbOk() (*DynamodbIntegration, bool)`

GetDynamodbOk returns a tuple with the Dynamodb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamodb

`func (o *Integration) SetDynamodb(v DynamodbIntegration)`

SetDynamodb sets Dynamodb field to given value.

### HasDynamodb

`func (o *Integration) HasDynamodb() bool`

HasDynamodb returns a boolean if a field has been set.

### GetRedshift

`func (o *Integration) GetRedshift() RedshiftIntegration`

GetRedshift returns the Redshift field if non-nil, zero value otherwise.

### GetRedshiftOk

`func (o *Integration) GetRedshiftOk() (*RedshiftIntegration, bool)`

GetRedshiftOk returns a tuple with the Redshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshift

`func (o *Integration) SetRedshift(v RedshiftIntegration)`

SetRedshift sets Redshift field to given value.

### HasRedshift

`func (o *Integration) HasRedshift() bool`

HasRedshift returns a boolean if a field has been set.

### GetGcs

`func (o *Integration) GetGcs() GcsIntegration`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *Integration) GetGcsOk() (*GcsIntegration, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *Integration) SetGcs(v GcsIntegration)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *Integration) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetSegment

`func (o *Integration) GetSegment() SegmentIntegration`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *Integration) GetSegmentOk() (*SegmentIntegration, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *Integration) SetSegment(v SegmentIntegration)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *Integration) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetKafka

`func (o *Integration) GetKafka() KafkaIntegration`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *Integration) GetKafkaOk() (*KafkaIntegration, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *Integration) SetKafka(v KafkaIntegration)`

SetKafka sets Kafka field to given value.

### HasKafka

`func (o *Integration) HasKafka() bool`

HasKafka returns a boolean if a field has been set.

### GetMongodb

`func (o *Integration) GetMongodb() MongoDbIntegration`

GetMongodb returns the Mongodb field if non-nil, zero value otherwise.

### GetMongodbOk

`func (o *Integration) GetMongodbOk() (*MongoDbIntegration, bool)`

GetMongodbOk returns a tuple with the Mongodb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodb

`func (o *Integration) SetMongodb(v MongoDbIntegration)`

SetMongodb sets Mongodb field to given value.

### HasMongodb

`func (o *Integration) HasMongodb() bool`

HasMongodb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


