# UpdateScheduledLambdaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apikey** | Pointer to **string** | The apikey to use when triggering execution of the associated query lambda. | [optional] 
**ResumePermanentError** | Pointer to **bool** | Boolean flag to allow a scheduled query lambda to resume execution after being suspended due to execution failure. This flag will be unset after scheduled lambda execution. | [optional] 
**TotalTimesToExecute** | Pointer to **int64** | The number of times to execute this scheduled query lambda. | [optional] 
**WebhookAuthHeader** | Pointer to **string** | The value to use as the authorization header when hitting the webhook. | [optional] 
**WebhookPayload** | Pointer to **string** | The payload that should be sent to the webhook. JSON format. | [optional] 
**WebhookUrl** | Pointer to **string** | The URL of the webhook that should be triggered after this scheduled query lambda completes. | [optional] 

## Methods

### NewUpdateScheduledLambdaRequest

`func NewUpdateScheduledLambdaRequest() *UpdateScheduledLambdaRequest`

NewUpdateScheduledLambdaRequest instantiates a new UpdateScheduledLambdaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateScheduledLambdaRequestWithDefaults

`func NewUpdateScheduledLambdaRequestWithDefaults() *UpdateScheduledLambdaRequest`

NewUpdateScheduledLambdaRequestWithDefaults instantiates a new UpdateScheduledLambdaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApikey

`func (o *UpdateScheduledLambdaRequest) GetApikey() string`

GetApikey returns the Apikey field if non-nil, zero value otherwise.

### GetApikeyOk

`func (o *UpdateScheduledLambdaRequest) GetApikeyOk() (*string, bool)`

GetApikeyOk returns a tuple with the Apikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApikey

`func (o *UpdateScheduledLambdaRequest) SetApikey(v string)`

SetApikey sets Apikey field to given value.

### HasApikey

`func (o *UpdateScheduledLambdaRequest) HasApikey() bool`

HasApikey returns a boolean if a field has been set.

### GetResumePermanentError

`func (o *UpdateScheduledLambdaRequest) GetResumePermanentError() bool`

GetResumePermanentError returns the ResumePermanentError field if non-nil, zero value otherwise.

### GetResumePermanentErrorOk

`func (o *UpdateScheduledLambdaRequest) GetResumePermanentErrorOk() (*bool, bool)`

GetResumePermanentErrorOk returns a tuple with the ResumePermanentError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePermanentError

`func (o *UpdateScheduledLambdaRequest) SetResumePermanentError(v bool)`

SetResumePermanentError sets ResumePermanentError field to given value.

### HasResumePermanentError

`func (o *UpdateScheduledLambdaRequest) HasResumePermanentError() bool`

HasResumePermanentError returns a boolean if a field has been set.

### GetTotalTimesToExecute

`func (o *UpdateScheduledLambdaRequest) GetTotalTimesToExecute() int64`

GetTotalTimesToExecute returns the TotalTimesToExecute field if non-nil, zero value otherwise.

### GetTotalTimesToExecuteOk

`func (o *UpdateScheduledLambdaRequest) GetTotalTimesToExecuteOk() (*int64, bool)`

GetTotalTimesToExecuteOk returns a tuple with the TotalTimesToExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimesToExecute

`func (o *UpdateScheduledLambdaRequest) SetTotalTimesToExecute(v int64)`

SetTotalTimesToExecute sets TotalTimesToExecute field to given value.

### HasTotalTimesToExecute

`func (o *UpdateScheduledLambdaRequest) HasTotalTimesToExecute() bool`

HasTotalTimesToExecute returns a boolean if a field has been set.

### GetWebhookAuthHeader

`func (o *UpdateScheduledLambdaRequest) GetWebhookAuthHeader() string`

GetWebhookAuthHeader returns the WebhookAuthHeader field if non-nil, zero value otherwise.

### GetWebhookAuthHeaderOk

`func (o *UpdateScheduledLambdaRequest) GetWebhookAuthHeaderOk() (*string, bool)`

GetWebhookAuthHeaderOk returns a tuple with the WebhookAuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookAuthHeader

`func (o *UpdateScheduledLambdaRequest) SetWebhookAuthHeader(v string)`

SetWebhookAuthHeader sets WebhookAuthHeader field to given value.

### HasWebhookAuthHeader

`func (o *UpdateScheduledLambdaRequest) HasWebhookAuthHeader() bool`

HasWebhookAuthHeader returns a boolean if a field has been set.

### GetWebhookPayload

`func (o *UpdateScheduledLambdaRequest) GetWebhookPayload() string`

GetWebhookPayload returns the WebhookPayload field if non-nil, zero value otherwise.

### GetWebhookPayloadOk

`func (o *UpdateScheduledLambdaRequest) GetWebhookPayloadOk() (*string, bool)`

GetWebhookPayloadOk returns a tuple with the WebhookPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookPayload

`func (o *UpdateScheduledLambdaRequest) SetWebhookPayload(v string)`

SetWebhookPayload sets WebhookPayload field to given value.

### HasWebhookPayload

`func (o *UpdateScheduledLambdaRequest) HasWebhookPayload() bool`

HasWebhookPayload returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *UpdateScheduledLambdaRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateScheduledLambdaRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateScheduledLambdaRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdateScheduledLambdaRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


