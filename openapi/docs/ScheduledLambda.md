# ScheduledLambda

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronString** | Pointer to **string** | The UNIX-formatted cron string for this scheduled query lambda. | [optional] 
**ExecutionCount** | Pointer to **int64** | The number of times this scheduled QL has been executed. | [optional] 
**LastCompletionDate** | Pointer to **string** | The last time this scheduled query lambda completed successfully. | [optional] 
**LastQueryId** | Pointer to **string** | The ID of the query that was triggered by this scheduled lambda&#39;s last run. | [optional] 
**NextExecutionDate** | Pointer to **string** | The next time this scheduled query lambda will be executed. | [optional] 
**QlName** | Pointer to **string** | The name of the associated query lambda. | [optional] 
**QueryExecutionStatus** | Pointer to [**ExecutionStatus**](ExecutionStatus.md) |  | [optional] 
**ResumePermanentError** | Pointer to **bool** | Boolean flag to allow a scheduled query lambda to resume execution after being suspended due to execution failure. This flag will be unset after scheduled lambda execution. | [optional] 
**Rrn** | Pointer to **string** | Scheduled Lambda mapping RRN. | [optional] 
**Tag** | Pointer to **string** | The query lambda tag. | [optional] 
**TotalTimesToExecute** | Pointer to **int64** | The number of times to execute this scheduled query lambda. Once this scheduled query lambda has been executed this many times, it will no longer be executed. | [optional] 
**Version** | Pointer to **string** | The version of the associated query lambda. | [optional] 
**WebhookExecutionStatus** | Pointer to [**ExecutionStatus**](ExecutionStatus.md) |  | [optional] 
**WebhookPayload** | Pointer to **string** | The payload that should be sent to the webhook. | [optional] 
**WebhookUrl** | Pointer to **string** | The URL of the webhook that should be triggered after this scheduled query lambda completes. | [optional] 
**Workspace** | Pointer to **string** | Workspace of the associated query lambda. | [optional] 

## Methods

### NewScheduledLambda

`func NewScheduledLambda() *ScheduledLambda`

NewScheduledLambda instantiates a new ScheduledLambda object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledLambdaWithDefaults

`func NewScheduledLambdaWithDefaults() *ScheduledLambda`

NewScheduledLambdaWithDefaults instantiates a new ScheduledLambda object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronString

`func (o *ScheduledLambda) GetCronString() string`

GetCronString returns the CronString field if non-nil, zero value otherwise.

### GetCronStringOk

`func (o *ScheduledLambda) GetCronStringOk() (*string, bool)`

GetCronStringOk returns a tuple with the CronString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronString

`func (o *ScheduledLambda) SetCronString(v string)`

SetCronString sets CronString field to given value.

### HasCronString

`func (o *ScheduledLambda) HasCronString() bool`

HasCronString returns a boolean if a field has been set.

### GetExecutionCount

`func (o *ScheduledLambda) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *ScheduledLambda) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *ScheduledLambda) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *ScheduledLambda) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.

### GetLastCompletionDate

`func (o *ScheduledLambda) GetLastCompletionDate() string`

GetLastCompletionDate returns the LastCompletionDate field if non-nil, zero value otherwise.

### GetLastCompletionDateOk

`func (o *ScheduledLambda) GetLastCompletionDateOk() (*string, bool)`

GetLastCompletionDateOk returns a tuple with the LastCompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletionDate

`func (o *ScheduledLambda) SetLastCompletionDate(v string)`

SetLastCompletionDate sets LastCompletionDate field to given value.

### HasLastCompletionDate

`func (o *ScheduledLambda) HasLastCompletionDate() bool`

HasLastCompletionDate returns a boolean if a field has been set.

### GetLastQueryId

`func (o *ScheduledLambda) GetLastQueryId() string`

GetLastQueryId returns the LastQueryId field if non-nil, zero value otherwise.

### GetLastQueryIdOk

`func (o *ScheduledLambda) GetLastQueryIdOk() (*string, bool)`

GetLastQueryIdOk returns a tuple with the LastQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueryId

`func (o *ScheduledLambda) SetLastQueryId(v string)`

SetLastQueryId sets LastQueryId field to given value.

### HasLastQueryId

`func (o *ScheduledLambda) HasLastQueryId() bool`

HasLastQueryId returns a boolean if a field has been set.

### GetNextExecutionDate

`func (o *ScheduledLambda) GetNextExecutionDate() string`

GetNextExecutionDate returns the NextExecutionDate field if non-nil, zero value otherwise.

### GetNextExecutionDateOk

`func (o *ScheduledLambda) GetNextExecutionDateOk() (*string, bool)`

GetNextExecutionDateOk returns a tuple with the NextExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecutionDate

`func (o *ScheduledLambda) SetNextExecutionDate(v string)`

SetNextExecutionDate sets NextExecutionDate field to given value.

### HasNextExecutionDate

`func (o *ScheduledLambda) HasNextExecutionDate() bool`

HasNextExecutionDate returns a boolean if a field has been set.

### GetQlName

`func (o *ScheduledLambda) GetQlName() string`

GetQlName returns the QlName field if non-nil, zero value otherwise.

### GetQlNameOk

`func (o *ScheduledLambda) GetQlNameOk() (*string, bool)`

GetQlNameOk returns a tuple with the QlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQlName

`func (o *ScheduledLambda) SetQlName(v string)`

SetQlName sets QlName field to given value.

### HasQlName

`func (o *ScheduledLambda) HasQlName() bool`

HasQlName returns a boolean if a field has been set.

### GetQueryExecutionStatus

`func (o *ScheduledLambda) GetQueryExecutionStatus() ExecutionStatus`

GetQueryExecutionStatus returns the QueryExecutionStatus field if non-nil, zero value otherwise.

### GetQueryExecutionStatusOk

`func (o *ScheduledLambda) GetQueryExecutionStatusOk() (*ExecutionStatus, bool)`

GetQueryExecutionStatusOk returns a tuple with the QueryExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryExecutionStatus

`func (o *ScheduledLambda) SetQueryExecutionStatus(v ExecutionStatus)`

SetQueryExecutionStatus sets QueryExecutionStatus field to given value.

### HasQueryExecutionStatus

`func (o *ScheduledLambda) HasQueryExecutionStatus() bool`

HasQueryExecutionStatus returns a boolean if a field has been set.

### GetResumePermanentError

`func (o *ScheduledLambda) GetResumePermanentError() bool`

GetResumePermanentError returns the ResumePermanentError field if non-nil, zero value otherwise.

### GetResumePermanentErrorOk

`func (o *ScheduledLambda) GetResumePermanentErrorOk() (*bool, bool)`

GetResumePermanentErrorOk returns a tuple with the ResumePermanentError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePermanentError

`func (o *ScheduledLambda) SetResumePermanentError(v bool)`

SetResumePermanentError sets ResumePermanentError field to given value.

### HasResumePermanentError

`func (o *ScheduledLambda) HasResumePermanentError() bool`

HasResumePermanentError returns a boolean if a field has been set.

### GetRrn

`func (o *ScheduledLambda) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *ScheduledLambda) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *ScheduledLambda) SetRrn(v string)`

SetRrn sets Rrn field to given value.

### HasRrn

`func (o *ScheduledLambda) HasRrn() bool`

HasRrn returns a boolean if a field has been set.

### GetTag

`func (o *ScheduledLambda) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ScheduledLambda) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ScheduledLambda) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ScheduledLambda) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTotalTimesToExecute

`func (o *ScheduledLambda) GetTotalTimesToExecute() int64`

GetTotalTimesToExecute returns the TotalTimesToExecute field if non-nil, zero value otherwise.

### GetTotalTimesToExecuteOk

`func (o *ScheduledLambda) GetTotalTimesToExecuteOk() (*int64, bool)`

GetTotalTimesToExecuteOk returns a tuple with the TotalTimesToExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimesToExecute

`func (o *ScheduledLambda) SetTotalTimesToExecute(v int64)`

SetTotalTimesToExecute sets TotalTimesToExecute field to given value.

### HasTotalTimesToExecute

`func (o *ScheduledLambda) HasTotalTimesToExecute() bool`

HasTotalTimesToExecute returns a boolean if a field has been set.

### GetVersion

`func (o *ScheduledLambda) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScheduledLambda) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScheduledLambda) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ScheduledLambda) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWebhookExecutionStatus

`func (o *ScheduledLambda) GetWebhookExecutionStatus() ExecutionStatus`

GetWebhookExecutionStatus returns the WebhookExecutionStatus field if non-nil, zero value otherwise.

### GetWebhookExecutionStatusOk

`func (o *ScheduledLambda) GetWebhookExecutionStatusOk() (*ExecutionStatus, bool)`

GetWebhookExecutionStatusOk returns a tuple with the WebhookExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookExecutionStatus

`func (o *ScheduledLambda) SetWebhookExecutionStatus(v ExecutionStatus)`

SetWebhookExecutionStatus sets WebhookExecutionStatus field to given value.

### HasWebhookExecutionStatus

`func (o *ScheduledLambda) HasWebhookExecutionStatus() bool`

HasWebhookExecutionStatus returns a boolean if a field has been set.

### GetWebhookPayload

`func (o *ScheduledLambda) GetWebhookPayload() string`

GetWebhookPayload returns the WebhookPayload field if non-nil, zero value otherwise.

### GetWebhookPayloadOk

`func (o *ScheduledLambda) GetWebhookPayloadOk() (*string, bool)`

GetWebhookPayloadOk returns a tuple with the WebhookPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookPayload

`func (o *ScheduledLambda) SetWebhookPayload(v string)`

SetWebhookPayload sets WebhookPayload field to given value.

### HasWebhookPayload

`func (o *ScheduledLambda) HasWebhookPayload() bool`

HasWebhookPayload returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ScheduledLambda) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ScheduledLambda) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ScheduledLambda) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ScheduledLambda) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWorkspace

`func (o *ScheduledLambda) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *ScheduledLambda) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *ScheduledLambda) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *ScheduledLambda) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


