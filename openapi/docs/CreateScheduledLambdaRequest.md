# CreateScheduledLambdaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apikey** | **string** | The apikey to use when triggering execution of the associated query lambda. | 
**CronString** | **string** | The UNIX-formatted cron string for this scheduled query lambda. | 
**QlName** | **string** | The name of the QL to use for scheduled execution. | 
**Tag** | Pointer to **string** | The QL tag to use for scheduled execution. One of either the QL tag or version must be specified | [optional] 
**TotalTimesToExecute** | Pointer to **int64** | The number of times to execute this scheduled query lambda. Once this scheduled query lambda has been executed this many times, it will no longer be executed. | [optional] 
**Version** | Pointer to **string** | The version of the QL to use for scheduled execution. One of either the QL version or tag must be specified. | [optional] 
**WebhookAuthHeader** | Pointer to **string** | The value to use as the authorization header when hitting the webhook. | [optional] 
**WebhookPayload** | Pointer to **string** | The payload that should be sent to the webhook. JSON format. | [optional] 
**WebhookUrl** | Pointer to **string** | The URL of the webhook that should be triggered after this scheduled query lambda completes. | [optional] 

## Methods

### NewCreateScheduledLambdaRequest

`func NewCreateScheduledLambdaRequest(apikey string, cronString string, qlName string, ) *CreateScheduledLambdaRequest`

NewCreateScheduledLambdaRequest instantiates a new CreateScheduledLambdaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScheduledLambdaRequestWithDefaults

`func NewCreateScheduledLambdaRequestWithDefaults() *CreateScheduledLambdaRequest`

NewCreateScheduledLambdaRequestWithDefaults instantiates a new CreateScheduledLambdaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApikey

`func (o *CreateScheduledLambdaRequest) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *CreateScheduledLambdaRequest) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *CreateScheduledLambdaRequest) SetApikey(v string)`

SetApikey sets Apikey field to given value.


### GetCronString

`func (o *CreateScheduledLambdaRequest) GetCronString() string`

GetCronString returns the CronString field if non-nil, zero value otherwise.

### GetCronStringOk

`func (o *CreateScheduledLambdaRequest) GetCronStringOk() (*string, bool)`

GetCronStringOk returns a tuple with the CronString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronString

`func (o *CreateScheduledLambdaRequest) SetCronString(v string)`

SetCronString sets CronString field to given value.


### GetQlName

`func (o *CreateScheduledLambdaRequest) GetQlName() string`

GetQlName returns the QlName field if non-nil, zero value otherwise.

### GetQlNameOk

`func (o *CreateScheduledLambdaRequest) GetQlNameOk() (*string, bool)`

GetQlNameOk returns a tuple with the QlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQlName

`func (o *CreateScheduledLambdaRequest) SetQlName(v string)`

SetQlName sets QlName field to given value.


### GetTag

`func (o *CreateScheduledLambdaRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateScheduledLambdaRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateScheduledLambdaRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateScheduledLambdaRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTotalTimesToExecute

`func (o *CreateScheduledLambdaRequest) GetTotalTimesToExecute() int64`

GetTotalTimesToExecute returns the TotalTimesToExecute field if non-nil, zero value otherwise.

### GetTotalTimesToExecuteOk

`func (o *CreateScheduledLambdaRequest) GetTotalTimesToExecuteOk() (*int64, bool)`

GetTotalTimesToExecuteOk returns a tuple with the TotalTimesToExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimesToExecute

`func (o *CreateScheduledLambdaRequest) SetTotalTimesToExecute(v int64)`

SetTotalTimesToExecute sets TotalTimesToExecute field to given value.

### HasTotalTimesToExecute

`func (o *CreateScheduledLambdaRequest) HasTotalTimesToExecute() bool`

HasTotalTimesToExecute returns a boolean if a field has been set.

### GetVersion

`func (o *CreateScheduledLambdaRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateScheduledLambdaRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateScheduledLambdaRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateScheduledLambdaRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWebhookAuthHeader

`func (o *CreateScheduledLambdaRequest) GetWebhookAuthHeader() string`

GetWebhookAuthHeader returns the WebhookAuthHeader field if non-nil, zero value otherwise.

### GetWebhookAuthHeaderOk

`func (o *CreateScheduledLambdaRequest) GetWebhookAuthHeaderOk() (*string, bool)`

GetWebhookAuthHeaderOk returns a tuple with the WebhookAuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookAuthHeader

`func (o *CreateScheduledLambdaRequest) SetWebhookAuthHeader(v string)`

SetWebhookAuthHeader sets WebhookAuthHeader field to given value.

### HasWebhookAuthHeader

`func (o *CreateScheduledLambdaRequest) HasWebhookAuthHeader() bool`

HasWebhookAuthHeader returns a boolean if a field has been set.

### GetWebhookPayload

`func (o *CreateScheduledLambdaRequest) GetWebhookPayload() string`

GetWebhookPayload returns the WebhookPayload field if non-nil, zero value otherwise.

### GetWebhookPayloadOk

`func (o *CreateScheduledLambdaRequest) GetWebhookPayloadOk() (*string, bool)`

GetWebhookPayloadOk returns a tuple with the WebhookPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookPayload

`func (o *CreateScheduledLambdaRequest) SetWebhookPayload(v string)`

SetWebhookPayload sets WebhookPayload field to given value.

### HasWebhookPayload

`func (o *CreateScheduledLambdaRequest) HasWebhookPayload() bool`

HasWebhookPayload returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateScheduledLambdaRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateScheduledLambdaRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateScheduledLambdaRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateScheduledLambdaRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


