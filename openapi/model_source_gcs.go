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

// SourceGcs struct for SourceGcs
type SourceGcs struct {
	// Name of GCS bucket you want to ingest from.
	Bucket *string `json:"bucket,omitempty"`
	ObjectBytesDownloaded *int64 `json:"object_bytes_downloaded,omitempty"`
	ObjectBytesTotal *int64 `json:"object_bytes_total,omitempty"`
	ObjectCountDownloaded *int64 `json:"object_count_downloaded,omitempty"`
	ObjectCountTotal *int64 `json:"object_count_total,omitempty"`
	// Glob-style pattern that selects keys to ingest. Only either prefix or pattern can be specified.
	Pattern *string `json:"pattern,omitempty"`
	// Prefix that selects keys to ingest.
	Prefix *string `json:"prefix,omitempty"`
}

// NewSourceGcs instantiates a new SourceGcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceGcs() *SourceGcs {
	this := SourceGcs{}
	return &this
}

// NewSourceGcsWithDefaults instantiates a new SourceGcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceGcsWithDefaults() *SourceGcs {
	this := SourceGcs{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *SourceGcs) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceGcs) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *SourceGcs) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *SourceGcs) SetBucket(v string) {
	o.Bucket = &v
}

// GetObjectBytesDownloaded returns the ObjectBytesDownloaded field value if set, zero value otherwise.
func (o *SourceGcs) GetObjectBytesDownloaded() int64 {
	if o == nil || o.ObjectBytesDownloaded == nil {
		var ret int64
		return ret
	}
	return *o.ObjectBytesDownloaded
}

// GetObjectBytesDownloadedOk returns a tuple with the ObjectBytesDownloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceGcs) GetObjectBytesDownloadedOk() (*int64, bool) {
	if o == nil || o.ObjectBytesDownloaded == nil {
		return nil, false
	}
	return o.ObjectBytesDownloaded, true
}

// HasObjectBytesDownloaded returns a boolean if a field has been set.
func (o *SourceGcs) HasObjectBytesDownloaded() bool {
	if o != nil && o.ObjectBytesDownloaded != nil {
		return true
	}

	return false
}

// SetObjectBytesDownloaded gets a reference to the given int64 and assigns it to the ObjectBytesDownloaded field.
func (o *SourceGcs) SetObjectBytesDownloaded(v int64) {
	o.ObjectBytesDownloaded = &v
}

// GetObjectBytesTotal returns the ObjectBytesTotal field value if set, zero value otherwise.
func (o *SourceGcs) GetObjectBytesTotal() int64 {
	if o == nil || o.ObjectBytesTotal == nil {
		var ret int64
		return ret
	}
	return *o.ObjectBytesTotal
}

// GetObjectBytesTotalOk returns a tuple with the ObjectBytesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceGcs) GetObjectBytesTotalOk() (*int64, bool) {
	if o == nil || o.ObjectBytesTotal == nil {
		return nil, false
	}
	return o.ObjectBytesTotal, true
}

// HasObjectBytesTotal returns a boolean if a field has been set.
func (o *SourceGcs) HasObjectBytesTotal() bool {
	if o != nil && o.ObjectBytesTotal != nil {
		return true
	}

	return false
}

// SetObjectBytesTotal gets a reference to the given int64 and assigns it to the ObjectBytesTotal field.
func (o *SourceGcs) SetObjectBytesTotal(v int64) {
	o.ObjectBytesTotal = &v
}

// GetObjectCountDownloaded returns the ObjectCountDownloaded field value if set, zero value otherwise.
func (o *SourceGcs) GetObjectCountDownloaded() int64 {
	if o == nil || o.ObjectCountDownloaded == nil {
		var ret int64
		return ret
	}
	return *o.ObjectCountDownloaded
}

// GetObjectCountDownloadedOk returns a tuple with the ObjectCountDownloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceGcs) GetObjectCountDownloadedOk() (*int64, bool) {
	if o == nil || o.ObjectCountDownloaded == nil {
		return nil, false
	}
	return o.ObjectCountDownloaded, true
}

// HasObjectCountDownloaded returns a boolean if a field has been set.
func (o *SourceGcs) HasObjectCountDownloaded() bool {
	if o != nil && o.ObjectCountDownloaded != nil {
		return true
	}

	return false
}

// SetObjectCountDownloaded gets a reference to the given int64 and assigns it to the ObjectCountDownloaded field.
func (o *SourceGcs) SetObjectCountDownloaded(v int64) {
	o.ObjectCountDownloaded = &v
}

// GetObjectCountTotal returns the ObjectCountTotal field value if set, zero value otherwise.
func (o *SourceGcs) GetObjectCountTotal() int64 {
	if o == nil || o.ObjectCountTotal == nil {
		var ret int64
		return ret
	}
	return *o.ObjectCountTotal
}

// GetObjectCountTotalOk returns a tuple with the ObjectCountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceGcs) GetObjectCountTotalOk() (*int64, bool) {
	if o == nil || o.ObjectCountTotal == nil {
		return nil, false
	}
	return o.ObjectCountTotal, true
}

// HasObjectCountTotal returns a boolean if a field has been set.
func (o *SourceGcs) HasObjectCountTotal() bool {
	if o != nil && o.ObjectCountTotal != nil {
		return true
	}

	return false
}

// SetObjectCountTotal gets a reference to the given int64 and assigns it to the ObjectCountTotal field.
func (o *SourceGcs) SetObjectCountTotal(v int64) {
	o.ObjectCountTotal = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SourceGcs) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceGcs) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SourceGcs) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SourceGcs) SetPattern(v string) {
	o.Pattern = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *SourceGcs) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceGcs) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *SourceGcs) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *SourceGcs) SetPrefix(v string) {
	o.Prefix = &v
}

func (o SourceGcs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.ObjectBytesDownloaded != nil {
		toSerialize["object_bytes_downloaded"] = o.ObjectBytesDownloaded
	}
	if o.ObjectBytesTotal != nil {
		toSerialize["object_bytes_total"] = o.ObjectBytesTotal
	}
	if o.ObjectCountDownloaded != nil {
		toSerialize["object_count_downloaded"] = o.ObjectCountDownloaded
	}
	if o.ObjectCountTotal != nil {
		toSerialize["object_count_total"] = o.ObjectCountTotal
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableSourceGcs struct {
	value *SourceGcs
	isSet bool
}

func (v NullableSourceGcs) Get() *SourceGcs {
	return v.value
}

func (v *NullableSourceGcs) Set(val *SourceGcs) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceGcs) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceGcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceGcs(val *SourceGcs) *NullableSourceGcs {
	return &NullableSourceGcs{value: val, isSet: true}
}

func (v NullableSourceGcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceGcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


