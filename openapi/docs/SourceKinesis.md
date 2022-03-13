# SourceKinesis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRegion** | Pointer to **string** | AWS region name of Kinesis stream, by default us-west-2 is used | [optional] 
**StreamName** | **string** | name of kinesis stream | 
**DmsPrimaryKey** | Pointer to **[]string** | set of fields that correspond to a DMS primary key | [optional] 
**OffsetResetPolicy** | Pointer to **string** | For non-DMS streams, Rockset can tail from the earliest end or latest end of kinesis source. | [optional] 

## Methods

### NewSourceKinesis

`func NewSourceKinesis(streamName string, ) *SourceKinesis`

NewSourceKinesis instantiates a new SourceKinesis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceKinesisWithDefaults

`func NewSourceKinesisWithDefaults() *SourceKinesis`

NewSourceKinesisWithDefaults instantiates a new SourceKinesis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRegion

`func (o *SourceKinesis) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *SourceKinesis) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *SourceKinesis) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *SourceKinesis) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetStreamName

`func (o *SourceKinesis) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *SourceKinesis) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *SourceKinesis) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.


### GetDmsPrimaryKey

`func (o *SourceKinesis) GetDmsPrimaryKey() []string`

GetDmsPrimaryKey returns the DmsPrimaryKey field if non-nil, zero value otherwise.

### GetDmsPrimaryKeyOk

`func (o *SourceKinesis) GetDmsPrimaryKeyOk() (*[]string, bool)`

GetDmsPrimaryKeyOk returns a tuple with the DmsPrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmsPrimaryKey

`func (o *SourceKinesis) SetDmsPrimaryKey(v []string)`

SetDmsPrimaryKey sets DmsPrimaryKey field to given value.

### HasDmsPrimaryKey

`func (o *SourceKinesis) HasDmsPrimaryKey() bool`

HasDmsPrimaryKey returns a boolean if a field has been set.

### GetOffsetResetPolicy

`func (o *SourceKinesis) GetOffsetResetPolicy() string`

GetOffsetResetPolicy returns the OffsetResetPolicy field if non-nil, zero value otherwise.

### GetOffsetResetPolicyOk

`func (o *SourceKinesis) GetOffsetResetPolicyOk() (*string, bool)`

GetOffsetResetPolicyOk returns a tuple with the OffsetResetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetResetPolicy

`func (o *SourceKinesis) SetOffsetResetPolicy(v string)`

SetOffsetResetPolicy sets OffsetResetPolicy field to given value.

### HasOffsetResetPolicy

`func (o *SourceKinesis) HasOffsetResetPolicy() bool`

HasOffsetResetPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


