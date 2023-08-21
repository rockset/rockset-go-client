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

// checks if the SourceKafka type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceKafka{}

// SourceKafka struct for SourceKafka
type SourceKafka struct {
	// The Kafka consumer group Id being used.
	ConsumerGroupId *string `json:"consumer_group_id,omitempty"`
	// The Kafka topic to be tailed.
	KafkaTopicName *string `json:"kafka_topic_name,omitempty"`
	// The offset reset policy.
	OffsetResetPolicy *string `json:"offset_reset_policy,omitempty"`
	Status *StatusKafka `json:"status,omitempty"`
	// Whether to use v3 integration.
	UseV3 *bool `json:"use_v3,omitempty"`
}

// NewSourceKafka instantiates a new SourceKafka object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceKafka() *SourceKafka {
	this := SourceKafka{}
	return &this
}

// NewSourceKafkaWithDefaults instantiates a new SourceKafka object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceKafkaWithDefaults() *SourceKafka {
	this := SourceKafka{}
	return &this
}

// GetConsumerGroupId returns the ConsumerGroupId field value if set, zero value otherwise.
func (o *SourceKafka) GetConsumerGroupId() string {
	if o == nil || IsNil(o.ConsumerGroupId) {
		var ret string
		return ret
	}
	return *o.ConsumerGroupId
}

// GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetConsumerGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerGroupId) {
		return nil, false
	}
	return o.ConsumerGroupId, true
}

// HasConsumerGroupId returns a boolean if a field has been set.
func (o *SourceKafka) HasConsumerGroupId() bool {
	if o != nil && !IsNil(o.ConsumerGroupId) {
		return true
	}

	return false
}

// SetConsumerGroupId gets a reference to the given string and assigns it to the ConsumerGroupId field.
func (o *SourceKafka) SetConsumerGroupId(v string) {
	o.ConsumerGroupId = &v
}

// GetKafkaTopicName returns the KafkaTopicName field value if set, zero value otherwise.
func (o *SourceKafka) GetKafkaTopicName() string {
	if o == nil || IsNil(o.KafkaTopicName) {
		var ret string
		return ret
	}
	return *o.KafkaTopicName
}

// GetKafkaTopicNameOk returns a tuple with the KafkaTopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetKafkaTopicNameOk() (*string, bool) {
	if o == nil || IsNil(o.KafkaTopicName) {
		return nil, false
	}
	return o.KafkaTopicName, true
}

// HasKafkaTopicName returns a boolean if a field has been set.
func (o *SourceKafka) HasKafkaTopicName() bool {
	if o != nil && !IsNil(o.KafkaTopicName) {
		return true
	}

	return false
}

// SetKafkaTopicName gets a reference to the given string and assigns it to the KafkaTopicName field.
func (o *SourceKafka) SetKafkaTopicName(v string) {
	o.KafkaTopicName = &v
}

// GetOffsetResetPolicy returns the OffsetResetPolicy field value if set, zero value otherwise.
func (o *SourceKafka) GetOffsetResetPolicy() string {
	if o == nil || IsNil(o.OffsetResetPolicy) {
		var ret string
		return ret
	}
	return *o.OffsetResetPolicy
}

// GetOffsetResetPolicyOk returns a tuple with the OffsetResetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetOffsetResetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.OffsetResetPolicy) {
		return nil, false
	}
	return o.OffsetResetPolicy, true
}

// HasOffsetResetPolicy returns a boolean if a field has been set.
func (o *SourceKafka) HasOffsetResetPolicy() bool {
	if o != nil && !IsNil(o.OffsetResetPolicy) {
		return true
	}

	return false
}

// SetOffsetResetPolicy gets a reference to the given string and assigns it to the OffsetResetPolicy field.
func (o *SourceKafka) SetOffsetResetPolicy(v string) {
	o.OffsetResetPolicy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SourceKafka) GetStatus() StatusKafka {
	if o == nil || IsNil(o.Status) {
		var ret StatusKafka
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetStatusOk() (*StatusKafka, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SourceKafka) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatusKafka and assigns it to the Status field.
func (o *SourceKafka) SetStatus(v StatusKafka) {
	o.Status = &v
}

// GetUseV3 returns the UseV3 field value if set, zero value otherwise.
func (o *SourceKafka) GetUseV3() bool {
	if o == nil || IsNil(o.UseV3) {
		var ret bool
		return ret
	}
	return *o.UseV3
}

// GetUseV3Ok returns a tuple with the UseV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetUseV3Ok() (*bool, bool) {
	if o == nil || IsNil(o.UseV3) {
		return nil, false
	}
	return o.UseV3, true
}

// HasUseV3 returns a boolean if a field has been set.
func (o *SourceKafka) HasUseV3() bool {
	if o != nil && !IsNil(o.UseV3) {
		return true
	}

	return false
}

// SetUseV3 gets a reference to the given bool and assigns it to the UseV3 field.
func (o *SourceKafka) SetUseV3(v bool) {
	o.UseV3 = &v
}

func (o SourceKafka) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceKafka) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumerGroupId) {
		toSerialize["consumer_group_id"] = o.ConsumerGroupId
	}
	if !IsNil(o.KafkaTopicName) {
		toSerialize["kafka_topic_name"] = o.KafkaTopicName
	}
	if !IsNil(o.OffsetResetPolicy) {
		toSerialize["offset_reset_policy"] = o.OffsetResetPolicy
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UseV3) {
		toSerialize["use_v3"] = o.UseV3
	}
	return toSerialize, nil
}

type NullableSourceKafka struct {
	value *SourceKafka
	isSet bool
}

func (v NullableSourceKafka) Get() *SourceKafka {
	return v.value
}

func (v *NullableSourceKafka) Set(val *SourceKafka) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceKafka) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceKafka) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceKafka(val *SourceKafka) *NullableSourceKafka {
	return &NullableSourceKafka{value: val, isSet: true}
}

func (v NullableSourceKafka) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceKafka) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


