/*
REST API

Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateScheduledLambdaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateScheduledLambdaRequest{}

// UpdateScheduledLambdaRequest struct for UpdateScheduledLambdaRequest
type UpdateScheduledLambdaRequest struct {
	// The apikey to use when triggering execution of the associated query lambda.
	Apikey *string `json:"apikey,omitempty"`
	// Boolean flag to allow a scheduled query lambda to resume execution after being suspended due to execution failure. This flag will be unset after scheduled lambda execution.
	ResumePermanentError *bool `json:"resume_permanent_error,omitempty"`
	// The number of times to execute this scheduled query lambda.
	TotalTimesToExecute *int64 `json:"total_times_to_execute,omitempty"`
	// The value to use as the authorization header when hitting the webhook.
	WebhookAuthHeader *string `json:"webhook_auth_header,omitempty"`
	// The payload that should be sent to the webhook. JSON format.
	WebhookPayload *string `json:"webhook_payload,omitempty"`
	// The URL of the webhook that should be triggered after this scheduled query lambda completes.
	WebhookUrl *string `json:"webhook_url,omitempty"`
}

// NewUpdateScheduledLambdaRequest instantiates a new UpdateScheduledLambdaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateScheduledLambdaRequest() *UpdateScheduledLambdaRequest {
	this := UpdateScheduledLambdaRequest{}
	return &this
}

// NewUpdateScheduledLambdaRequestWithDefaults instantiates a new UpdateScheduledLambdaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateScheduledLambdaRequestWithDefaults() *UpdateScheduledLambdaRequest {
	this := UpdateScheduledLambdaRequest{}
	return &this
}

// GetApikey returns the Apikey field value if set, zero value otherwise.
func (o *UpdateScheduledLambdaRequest) GetApikey() string {
	if o == nil || IsNil(o.Apikey) {
		var ret string
		return ret
	}
	return *o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledLambdaRequest) GetApikeyOk() (*string, bool) {
	if o == nil || IsNil(o.Apikey) {
		return nil, false
	}
	return o.Apikey, true
}

// HasApikey returns a boolean if a field has been set.
func (o *UpdateScheduledLambdaRequest) HasApikey() bool {
	if o != nil && !IsNil(o.Apikey) {
		return true
	}

	return false
}

// SetApikey gets a reference to the given string and assigns it to the Apikey field.
func (o *UpdateScheduledLambdaRequest) SetApikey(v string) {
	o.Apikey = &v
}

// GetResumePermanentError returns the ResumePermanentError field value if set, zero value otherwise.
func (o *UpdateScheduledLambdaRequest) GetResumePermanentError() bool {
	if o == nil || IsNil(o.ResumePermanentError) {
		var ret bool
		return ret
	}
	return *o.ResumePermanentError
}

// GetResumePermanentErrorOk returns a tuple with the ResumePermanentError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledLambdaRequest) GetResumePermanentErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.ResumePermanentError) {
		return nil, false
	}
	return o.ResumePermanentError, true
}

// HasResumePermanentError returns a boolean if a field has been set.
func (o *UpdateScheduledLambdaRequest) HasResumePermanentError() bool {
	if o != nil && !IsNil(o.ResumePermanentError) {
		return true
	}

	return false
}

// SetResumePermanentError gets a reference to the given bool and assigns it to the ResumePermanentError field.
func (o *UpdateScheduledLambdaRequest) SetResumePermanentError(v bool) {
	o.ResumePermanentError = &v
}

// GetTotalTimesToExecute returns the TotalTimesToExecute field value if set, zero value otherwise.
func (o *UpdateScheduledLambdaRequest) GetTotalTimesToExecute() int64 {
	if o == nil || IsNil(o.TotalTimesToExecute) {
		var ret int64
		return ret
	}
	return *o.TotalTimesToExecute
}

// GetTotalTimesToExecuteOk returns a tuple with the TotalTimesToExecute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledLambdaRequest) GetTotalTimesToExecuteOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalTimesToExecute) {
		return nil, false
	}
	return o.TotalTimesToExecute, true
}

// HasTotalTimesToExecute returns a boolean if a field has been set.
func (o *UpdateScheduledLambdaRequest) HasTotalTimesToExecute() bool {
	if o != nil && !IsNil(o.TotalTimesToExecute) {
		return true
	}

	return false
}

// SetTotalTimesToExecute gets a reference to the given int64 and assigns it to the TotalTimesToExecute field.
func (o *UpdateScheduledLambdaRequest) SetTotalTimesToExecute(v int64) {
	o.TotalTimesToExecute = &v
}

// GetWebhookAuthHeader returns the WebhookAuthHeader field value if set, zero value otherwise.
func (o *UpdateScheduledLambdaRequest) GetWebhookAuthHeader() string {
	if o == nil || IsNil(o.WebhookAuthHeader) {
		var ret string
		return ret
	}
	return *o.WebhookAuthHeader
}

// GetWebhookAuthHeaderOk returns a tuple with the WebhookAuthHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledLambdaRequest) GetWebhookAuthHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookAuthHeader) {
		return nil, false
	}
	return o.WebhookAuthHeader, true
}

// HasWebhookAuthHeader returns a boolean if a field has been set.
func (o *UpdateScheduledLambdaRequest) HasWebhookAuthHeader() bool {
	if o != nil && !IsNil(o.WebhookAuthHeader) {
		return true
	}

	return false
}

// SetWebhookAuthHeader gets a reference to the given string and assigns it to the WebhookAuthHeader field.
func (o *UpdateScheduledLambdaRequest) SetWebhookAuthHeader(v string) {
	o.WebhookAuthHeader = &v
}

// GetWebhookPayload returns the WebhookPayload field value if set, zero value otherwise.
func (o *UpdateScheduledLambdaRequest) GetWebhookPayload() string {
	if o == nil || IsNil(o.WebhookPayload) {
		var ret string
		return ret
	}
	return *o.WebhookPayload
}

// GetWebhookPayloadOk returns a tuple with the WebhookPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledLambdaRequest) GetWebhookPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookPayload) {
		return nil, false
	}
	return o.WebhookPayload, true
}

// HasWebhookPayload returns a boolean if a field has been set.
func (o *UpdateScheduledLambdaRequest) HasWebhookPayload() bool {
	if o != nil && !IsNil(o.WebhookPayload) {
		return true
	}

	return false
}

// SetWebhookPayload gets a reference to the given string and assigns it to the WebhookPayload field.
func (o *UpdateScheduledLambdaRequest) SetWebhookPayload(v string) {
	o.WebhookPayload = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *UpdateScheduledLambdaRequest) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScheduledLambdaRequest) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *UpdateScheduledLambdaRequest) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *UpdateScheduledLambdaRequest) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o UpdateScheduledLambdaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateScheduledLambdaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apikey) {
		toSerialize["apikey"] = o.Apikey
	}
	if !IsNil(o.ResumePermanentError) {
		toSerialize["resume_permanent_error"] = o.ResumePermanentError
	}
	if !IsNil(o.TotalTimesToExecute) {
		toSerialize["total_times_to_execute"] = o.TotalTimesToExecute
	}
	if !IsNil(o.WebhookAuthHeader) {
		toSerialize["webhook_auth_header"] = o.WebhookAuthHeader
	}
	if !IsNil(o.WebhookPayload) {
		toSerialize["webhook_payload"] = o.WebhookPayload
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	return toSerialize, nil
}

type NullableUpdateScheduledLambdaRequest struct {
	value *UpdateScheduledLambdaRequest
	isSet bool
}

func (v NullableUpdateScheduledLambdaRequest) Get() *UpdateScheduledLambdaRequest {
	return v.value
}

func (v *NullableUpdateScheduledLambdaRequest) Set(val *UpdateScheduledLambdaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateScheduledLambdaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateScheduledLambdaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateScheduledLambdaRequest(val *UpdateScheduledLambdaRequest) *NullableUpdateScheduledLambdaRequest {
	return &NullableUpdateScheduledLambdaRequest{value: val, isSet: true}
}

func (v NullableUpdateScheduledLambdaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateScheduledLambdaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

