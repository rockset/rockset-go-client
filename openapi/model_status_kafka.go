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

// StatusKafka struct for StatusKafka
type StatusKafka struct {
	// State of the Kafka source
	State *string `json:"state,omitempty"`
	// Time at which the last document was consumed from Kafka
	LastConsumedTime *string `json:"last_consumed_time,omitempty"`
	// Number of documents consumed by this Kafka topic
	NumDocumentsProcessed *int64 `json:"num_documents_processed,omitempty"`
	// Status info per partition
	KafkaPartitions []StatusKafkaPartition `json:"kafka_partitions,omitempty"`
}

// NewStatusKafka instantiates a new StatusKafka object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusKafka() *StatusKafka {
	this := StatusKafka{}
	return &this
}

// NewStatusKafkaWithDefaults instantiates a new StatusKafka object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusKafkaWithDefaults() *StatusKafka {
	this := StatusKafka{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StatusKafka) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusKafka) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StatusKafka) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StatusKafka) SetState(v string) {
	o.State = &v
}

// GetLastConsumedTime returns the LastConsumedTime field value if set, zero value otherwise.
func (o *StatusKafka) GetLastConsumedTime() string {
	if o == nil || o.LastConsumedTime == nil {
		var ret string
		return ret
	}
	return *o.LastConsumedTime
}

// GetLastConsumedTimeOk returns a tuple with the LastConsumedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusKafka) GetLastConsumedTimeOk() (*string, bool) {
	if o == nil || o.LastConsumedTime == nil {
		return nil, false
	}
	return o.LastConsumedTime, true
}

// HasLastConsumedTime returns a boolean if a field has been set.
func (o *StatusKafka) HasLastConsumedTime() bool {
	if o != nil && o.LastConsumedTime != nil {
		return true
	}

	return false
}

// SetLastConsumedTime gets a reference to the given string and assigns it to the LastConsumedTime field.
func (o *StatusKafka) SetLastConsumedTime(v string) {
	o.LastConsumedTime = &v
}

// GetNumDocumentsProcessed returns the NumDocumentsProcessed field value if set, zero value otherwise.
func (o *StatusKafka) GetNumDocumentsProcessed() int64 {
	if o == nil || o.NumDocumentsProcessed == nil {
		var ret int64
		return ret
	}
	return *o.NumDocumentsProcessed
}

// GetNumDocumentsProcessedOk returns a tuple with the NumDocumentsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusKafka) GetNumDocumentsProcessedOk() (*int64, bool) {
	if o == nil || o.NumDocumentsProcessed == nil {
		return nil, false
	}
	return o.NumDocumentsProcessed, true
}

// HasNumDocumentsProcessed returns a boolean if a field has been set.
func (o *StatusKafka) HasNumDocumentsProcessed() bool {
	if o != nil && o.NumDocumentsProcessed != nil {
		return true
	}

	return false
}

// SetNumDocumentsProcessed gets a reference to the given int64 and assigns it to the NumDocumentsProcessed field.
func (o *StatusKafka) SetNumDocumentsProcessed(v int64) {
	o.NumDocumentsProcessed = &v
}

// GetKafkaPartitions returns the KafkaPartitions field value if set, zero value otherwise.
func (o *StatusKafka) GetKafkaPartitions() []StatusKafkaPartition {
	if o == nil || o.KafkaPartitions == nil {
		var ret []StatusKafkaPartition
		return ret
	}
	return o.KafkaPartitions
}

// GetKafkaPartitionsOk returns a tuple with the KafkaPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusKafka) GetKafkaPartitionsOk() ([]StatusKafkaPartition, bool) {
	if o == nil || o.KafkaPartitions == nil {
		return nil, false
	}
	return o.KafkaPartitions, true
}

// HasKafkaPartitions returns a boolean if a field has been set.
func (o *StatusKafka) HasKafkaPartitions() bool {
	if o != nil && o.KafkaPartitions != nil {
		return true
	}

	return false
}

// SetKafkaPartitions gets a reference to the given []StatusKafkaPartition and assigns it to the KafkaPartitions field.
func (o *StatusKafka) SetKafkaPartitions(v []StatusKafkaPartition) {
	o.KafkaPartitions = v
}

func (o StatusKafka) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.LastConsumedTime != nil {
		toSerialize["last_consumed_time"] = o.LastConsumedTime
	}
	if o.NumDocumentsProcessed != nil {
		toSerialize["num_documents_processed"] = o.NumDocumentsProcessed
	}
	if o.KafkaPartitions != nil {
		toSerialize["kafka_partitions"] = o.KafkaPartitions
	}
	return json.Marshal(toSerialize)
}

type NullableStatusKafka struct {
	value *StatusKafka
	isSet bool
}

func (v NullableStatusKafka) Get() *StatusKafka {
	return v.value
}

func (v *NullableStatusKafka) Set(val *StatusKafka) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusKafka) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusKafka) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusKafka(val *StatusKafka) *NullableStatusKafka {
	return &NullableStatusKafka{value: val, isSet: true}
}

func (v NullableStatusKafka) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusKafka) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


