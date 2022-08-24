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

// StatusDynamoDb struct for StatusDynamoDb
type StatusDynamoDb struct {
	// DynamoDB scan start time.
	ScanStartTime *string `json:"scan_start_time,omitempty"`
	// DynamoDb scan end time.
	ScanEndTime *string `json:"scan_end_time,omitempty"`
	// Number of records inserted using scan.
	ScanRecordsProcessed *int64 `json:"scan_records_processed,omitempty"`
	// Number of records in DynamoDB table at time of scan.
	ScanTotalRecords *int64 `json:"scan_total_records,omitempty"`
	// State of current ingest for this table.
	State *string `json:"state,omitempty"`
	// ISO-8601 date when source was last processed.
	StreamLastProcessedAt *string `json:"stream_last_processed_at,omitempty"`
}

// NewStatusDynamoDb instantiates a new StatusDynamoDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusDynamoDb() *StatusDynamoDb {
	this := StatusDynamoDb{}
	return &this
}

// NewStatusDynamoDbWithDefaults instantiates a new StatusDynamoDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusDynamoDbWithDefaults() *StatusDynamoDb {
	this := StatusDynamoDb{}
	return &this
}

// GetScanStartTime returns the ScanStartTime field value if set, zero value otherwise.
func (o *StatusDynamoDb) GetScanStartTime() string {
	if o == nil || o.ScanStartTime == nil {
		var ret string
		return ret
	}
	return *o.ScanStartTime
}

// GetScanStartTimeOk returns a tuple with the ScanStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDynamoDb) GetScanStartTimeOk() (*string, bool) {
	if o == nil || o.ScanStartTime == nil {
		return nil, false
	}
	return o.ScanStartTime, true
}

// HasScanStartTime returns a boolean if a field has been set.
func (o *StatusDynamoDb) HasScanStartTime() bool {
	if o != nil && o.ScanStartTime != nil {
		return true
	}

	return false
}

// SetScanStartTime gets a reference to the given string and assigns it to the ScanStartTime field.
func (o *StatusDynamoDb) SetScanStartTime(v string) {
	o.ScanStartTime = &v
}

// GetScanEndTime returns the ScanEndTime field value if set, zero value otherwise.
func (o *StatusDynamoDb) GetScanEndTime() string {
	if o == nil || o.ScanEndTime == nil {
		var ret string
		return ret
	}
	return *o.ScanEndTime
}

// GetScanEndTimeOk returns a tuple with the ScanEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDynamoDb) GetScanEndTimeOk() (*string, bool) {
	if o == nil || o.ScanEndTime == nil {
		return nil, false
	}
	return o.ScanEndTime, true
}

// HasScanEndTime returns a boolean if a field has been set.
func (o *StatusDynamoDb) HasScanEndTime() bool {
	if o != nil && o.ScanEndTime != nil {
		return true
	}

	return false
}

// SetScanEndTime gets a reference to the given string and assigns it to the ScanEndTime field.
func (o *StatusDynamoDb) SetScanEndTime(v string) {
	o.ScanEndTime = &v
}

// GetScanRecordsProcessed returns the ScanRecordsProcessed field value if set, zero value otherwise.
func (o *StatusDynamoDb) GetScanRecordsProcessed() int64 {
	if o == nil || o.ScanRecordsProcessed == nil {
		var ret int64
		return ret
	}
	return *o.ScanRecordsProcessed
}

// GetScanRecordsProcessedOk returns a tuple with the ScanRecordsProcessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDynamoDb) GetScanRecordsProcessedOk() (*int64, bool) {
	if o == nil || o.ScanRecordsProcessed == nil {
		return nil, false
	}
	return o.ScanRecordsProcessed, true
}

// HasScanRecordsProcessed returns a boolean if a field has been set.
func (o *StatusDynamoDb) HasScanRecordsProcessed() bool {
	if o != nil && o.ScanRecordsProcessed != nil {
		return true
	}

	return false
}

// SetScanRecordsProcessed gets a reference to the given int64 and assigns it to the ScanRecordsProcessed field.
func (o *StatusDynamoDb) SetScanRecordsProcessed(v int64) {
	o.ScanRecordsProcessed = &v
}

// GetScanTotalRecords returns the ScanTotalRecords field value if set, zero value otherwise.
func (o *StatusDynamoDb) GetScanTotalRecords() int64 {
	if o == nil || o.ScanTotalRecords == nil {
		var ret int64
		return ret
	}
	return *o.ScanTotalRecords
}

// GetScanTotalRecordsOk returns a tuple with the ScanTotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDynamoDb) GetScanTotalRecordsOk() (*int64, bool) {
	if o == nil || o.ScanTotalRecords == nil {
		return nil, false
	}
	return o.ScanTotalRecords, true
}

// HasScanTotalRecords returns a boolean if a field has been set.
func (o *StatusDynamoDb) HasScanTotalRecords() bool {
	if o != nil && o.ScanTotalRecords != nil {
		return true
	}

	return false
}

// SetScanTotalRecords gets a reference to the given int64 and assigns it to the ScanTotalRecords field.
func (o *StatusDynamoDb) SetScanTotalRecords(v int64) {
	o.ScanTotalRecords = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StatusDynamoDb) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDynamoDb) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StatusDynamoDb) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StatusDynamoDb) SetState(v string) {
	o.State = &v
}

// GetStreamLastProcessedAt returns the StreamLastProcessedAt field value if set, zero value otherwise.
func (o *StatusDynamoDb) GetStreamLastProcessedAt() string {
	if o == nil || o.StreamLastProcessedAt == nil {
		var ret string
		return ret
	}
	return *o.StreamLastProcessedAt
}

// GetStreamLastProcessedAtOk returns a tuple with the StreamLastProcessedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDynamoDb) GetStreamLastProcessedAtOk() (*string, bool) {
	if o == nil || o.StreamLastProcessedAt == nil {
		return nil, false
	}
	return o.StreamLastProcessedAt, true
}

// HasStreamLastProcessedAt returns a boolean if a field has been set.
func (o *StatusDynamoDb) HasStreamLastProcessedAt() bool {
	if o != nil && o.StreamLastProcessedAt != nil {
		return true
	}

	return false
}

// SetStreamLastProcessedAt gets a reference to the given string and assigns it to the StreamLastProcessedAt field.
func (o *StatusDynamoDb) SetStreamLastProcessedAt(v string) {
	o.StreamLastProcessedAt = &v
}

func (o StatusDynamoDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScanStartTime != nil {
		toSerialize["scan_start_time"] = o.ScanStartTime
	}
	if o.ScanEndTime != nil {
		toSerialize["scan_end_time"] = o.ScanEndTime
	}
	if o.ScanRecordsProcessed != nil {
		toSerialize["scan_records_processed"] = o.ScanRecordsProcessed
	}
	if o.ScanTotalRecords != nil {
		toSerialize["scan_total_records"] = o.ScanTotalRecords
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StreamLastProcessedAt != nil {
		toSerialize["stream_last_processed_at"] = o.StreamLastProcessedAt
	}
	return json.Marshal(toSerialize)
}

type NullableStatusDynamoDb struct {
	value *StatusDynamoDb
	isSet bool
}

func (v NullableStatusDynamoDb) Get() *StatusDynamoDb {
	return v.value
}

func (v *NullableStatusDynamoDb) Set(val *StatusDynamoDb) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusDynamoDb) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusDynamoDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusDynamoDb(val *StatusDynamoDb) *NullableStatusDynamoDb {
	return &NullableStatusDynamoDb{value: val, isSet: true}
}

func (v NullableStatusDynamoDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusDynamoDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


