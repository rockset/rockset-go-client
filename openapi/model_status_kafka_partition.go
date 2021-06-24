/*
 * REST API
 *
 * Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatusKafkaPartition struct for StatusKafkaPartition
type StatusKafkaPartition struct {
	// The number of this partition
	PartitionNumber *int32 `json:"partition_number,omitempty"`
	// Number of documents consumed by this partition
	NumDocumentsProcessed *int64 `json:"num_documents_processed,omitempty"`
}

// NewStatusKafkaPartition instantiates a new StatusKafkaPartition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusKafkaPartition() *StatusKafkaPartition {
	this := StatusKafkaPartition{}
	return &this
}

// NewStatusKafkaPartitionWithDefaults instantiates a new StatusKafkaPartition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusKafkaPartitionWithDefaults() *StatusKafkaPartition {
	this := StatusKafkaPartition{}
	return &this
}

// GetPartitionNumber returns the PartitionNumber field value if set, zero value otherwise.
func (o *StatusKafkaPartition) GetPartitionNumber() int32 {
	if o == nil || o.PartitionNumber == nil {
		var ret int32
		return ret
	}
	return *o.PartitionNumber
}

// GetPartitionNumberOk returns a tuple with the PartitionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusKafkaPartition) GetPartitionNumberOk() (*int32, bool) {
	if o == nil || o.PartitionNumber == nil {
		return nil, false
	}
	return o.PartitionNumber, true
}

// HasPartitionNumber returns a boolean if a field has been set.
func (o *StatusKafkaPartition) HasPartitionNumber() bool {
	if o != nil && o.PartitionNumber != nil {
		return true
	}

	return false
}

// SetPartitionNumber gets a reference to the given int32 and assigns it to the PartitionNumber field.
func (o *StatusKafkaPartition) SetPartitionNumber(v int32) {
	o.PartitionNumber = &v
}

// GetNumDocumentsProcessed returns the NumDocumentsProcessed field value if set, zero value otherwise.
func (o *StatusKafkaPartition) GetNumDocumentsProcessed() int64 {
	if o == nil || o.NumDocumentsProcessed == nil {
		var ret int64
		return ret
	}
	return *o.NumDocumentsProcessed
}

// GetNumDocumentsProcessedOk returns a tuple with the NumDocumentsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusKafkaPartition) GetNumDocumentsProcessedOk() (*int64, bool) {
	if o == nil || o.NumDocumentsProcessed == nil {
		return nil, false
	}
	return o.NumDocumentsProcessed, true
}

// HasNumDocumentsProcessed returns a boolean if a field has been set.
func (o *StatusKafkaPartition) HasNumDocumentsProcessed() bool {
	if o != nil && o.NumDocumentsProcessed != nil {
		return true
	}

	return false
}

// SetNumDocumentsProcessed gets a reference to the given int64 and assigns it to the NumDocumentsProcessed field.
func (o *StatusKafkaPartition) SetNumDocumentsProcessed(v int64) {
	o.NumDocumentsProcessed = &v
}

func (o StatusKafkaPartition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PartitionNumber != nil {
		toSerialize["partition_number"] = o.PartitionNumber
	}
	if o.NumDocumentsProcessed != nil {
		toSerialize["num_documents_processed"] = o.NumDocumentsProcessed
	}
	return json.Marshal(toSerialize)
}

type NullableStatusKafkaPartition struct {
	value *StatusKafkaPartition
	isSet bool
}

func (v NullableStatusKafkaPartition) Get() *StatusKafkaPartition {
	return v.value
}

func (v *NullableStatusKafkaPartition) Set(val *StatusKafkaPartition) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusKafkaPartition) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusKafkaPartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusKafkaPartition(val *StatusKafkaPartition) *NullableStatusKafkaPartition {
	return &NullableStatusKafkaPartition{value: val, isSet: true}
}

func (v NullableStatusKafkaPartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusKafkaPartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

