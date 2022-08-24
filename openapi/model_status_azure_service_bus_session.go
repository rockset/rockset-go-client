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

// StatusAzureServiceBusSession struct for StatusAzureServiceBusSession
type StatusAzureServiceBusSession struct {
	// The last processed sequence number within this session.
	SequenceNumber *int64 `json:"sequence_number,omitempty"`
	// Most recent ISO-8601 date when a message from this session was processed.
	LastProcessed *string `json:"last_processed,omitempty"`
}

// NewStatusAzureServiceBusSession instantiates a new StatusAzureServiceBusSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusAzureServiceBusSession() *StatusAzureServiceBusSession {
	this := StatusAzureServiceBusSession{}
	return &this
}

// NewStatusAzureServiceBusSessionWithDefaults instantiates a new StatusAzureServiceBusSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusAzureServiceBusSessionWithDefaults() *StatusAzureServiceBusSession {
	this := StatusAzureServiceBusSession{}
	return &this
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *StatusAzureServiceBusSession) GetSequenceNumber() int64 {
	if o == nil || o.SequenceNumber == nil {
		var ret int64
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusAzureServiceBusSession) GetSequenceNumberOk() (*int64, bool) {
	if o == nil || o.SequenceNumber == nil {
		return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *StatusAzureServiceBusSession) HasSequenceNumber() bool {
	if o != nil && o.SequenceNumber != nil {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given int64 and assigns it to the SequenceNumber field.
func (o *StatusAzureServiceBusSession) SetSequenceNumber(v int64) {
	o.SequenceNumber = &v
}

// GetLastProcessed returns the LastProcessed field value if set, zero value otherwise.
func (o *StatusAzureServiceBusSession) GetLastProcessed() string {
	if o == nil || o.LastProcessed == nil {
		var ret string
		return ret
	}
	return *o.LastProcessed
}

// GetLastProcessedOk returns a tuple with the LastProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusAzureServiceBusSession) GetLastProcessedOk() (*string, bool) {
	if o == nil || o.LastProcessed == nil {
		return nil, false
	}
	return o.LastProcessed, true
}

// HasLastProcessed returns a boolean if a field has been set.
func (o *StatusAzureServiceBusSession) HasLastProcessed() bool {
	if o != nil && o.LastProcessed != nil {
		return true
	}

	return false
}

// SetLastProcessed gets a reference to the given string and assigns it to the LastProcessed field.
func (o *StatusAzureServiceBusSession) SetLastProcessed(v string) {
	o.LastProcessed = &v
}

func (o StatusAzureServiceBusSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SequenceNumber != nil {
		toSerialize["sequence_number"] = o.SequenceNumber
	}
	if o.LastProcessed != nil {
		toSerialize["last_processed"] = o.LastProcessed
	}
	return json.Marshal(toSerialize)
}

type NullableStatusAzureServiceBusSession struct {
	value *StatusAzureServiceBusSession
	isSet bool
}

func (v NullableStatusAzureServiceBusSession) Get() *StatusAzureServiceBusSession {
	return v.value
}

func (v *NullableStatusAzureServiceBusSession) Set(val *StatusAzureServiceBusSession) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusAzureServiceBusSession) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusAzureServiceBusSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusAzureServiceBusSession(val *StatusAzureServiceBusSession) *NullableStatusAzureServiceBusSession {
	return &NullableStatusAzureServiceBusSession{value: val, isSet: true}
}

func (v NullableStatusAzureServiceBusSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusAzureServiceBusSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


