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

// PaginationInfo struct for PaginationInfo
type PaginationInfo struct {
	// Number of documents returned in this result set
	CurrentPageDocCount *int64 `json:"current_page_doc_count,omitempty"`
	// The doc offset that next_cursor starts at.
	NextCursorOffset *int64 `json:"next_cursor_offset,omitempty"`
	// Cursor used to retrieve the first set of documents.
	StartCursor *string `json:"start_cursor,omitempty"`
	// Cursor to use to get the list of documents
	NextCursor *string `json:"next_cursor,omitempty"`
}

// NewPaginationInfo instantiates a new PaginationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationInfo() *PaginationInfo {
	this := PaginationInfo{}
	return &this
}

// NewPaginationInfoWithDefaults instantiates a new PaginationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationInfoWithDefaults() *PaginationInfo {
	this := PaginationInfo{}
	return &this
}

// GetCurrentPageDocCount returns the CurrentPageDocCount field value if set, zero value otherwise.
func (o *PaginationInfo) GetCurrentPageDocCount() int64 {
	if o == nil || o.CurrentPageDocCount == nil {
		var ret int64
		return ret
	}
	return *o.CurrentPageDocCount
}

// GetCurrentPageDocCountOk returns a tuple with the CurrentPageDocCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetCurrentPageDocCountOk() (*int64, bool) {
	if o == nil || o.CurrentPageDocCount == nil {
		return nil, false
	}
	return o.CurrentPageDocCount, true
}

// HasCurrentPageDocCount returns a boolean if a field has been set.
func (o *PaginationInfo) HasCurrentPageDocCount() bool {
	if o != nil && o.CurrentPageDocCount != nil {
		return true
	}

	return false
}

// SetCurrentPageDocCount gets a reference to the given int64 and assigns it to the CurrentPageDocCount field.
func (o *PaginationInfo) SetCurrentPageDocCount(v int64) {
	o.CurrentPageDocCount = &v
}

// GetNextCursorOffset returns the NextCursorOffset field value if set, zero value otherwise.
func (o *PaginationInfo) GetNextCursorOffset() int64 {
	if o == nil || o.NextCursorOffset == nil {
		var ret int64
		return ret
	}
	return *o.NextCursorOffset
}

// GetNextCursorOffsetOk returns a tuple with the NextCursorOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetNextCursorOffsetOk() (*int64, bool) {
	if o == nil || o.NextCursorOffset == nil {
		return nil, false
	}
	return o.NextCursorOffset, true
}

// HasNextCursorOffset returns a boolean if a field has been set.
func (o *PaginationInfo) HasNextCursorOffset() bool {
	if o != nil && o.NextCursorOffset != nil {
		return true
	}

	return false
}

// SetNextCursorOffset gets a reference to the given int64 and assigns it to the NextCursorOffset field.
func (o *PaginationInfo) SetNextCursorOffset(v int64) {
	o.NextCursorOffset = &v
}

// GetStartCursor returns the StartCursor field value if set, zero value otherwise.
func (o *PaginationInfo) GetStartCursor() string {
	if o == nil || o.StartCursor == nil {
		var ret string
		return ret
	}
	return *o.StartCursor
}

// GetStartCursorOk returns a tuple with the StartCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetStartCursorOk() (*string, bool) {
	if o == nil || o.StartCursor == nil {
		return nil, false
	}
	return o.StartCursor, true
}

// HasStartCursor returns a boolean if a field has been set.
func (o *PaginationInfo) HasStartCursor() bool {
	if o != nil && o.StartCursor != nil {
		return true
	}

	return false
}

// SetStartCursor gets a reference to the given string and assigns it to the StartCursor field.
func (o *PaginationInfo) SetStartCursor(v string) {
	o.StartCursor = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *PaginationInfo) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *PaginationInfo) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *PaginationInfo) SetNextCursor(v string) {
	o.NextCursor = &v
}

func (o PaginationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPageDocCount != nil {
		toSerialize["current_page_doc_count"] = o.CurrentPageDocCount
	}
	if o.NextCursorOffset != nil {
		toSerialize["next_cursor_offset"] = o.NextCursorOffset
	}
	if o.StartCursor != nil {
		toSerialize["start_cursor"] = o.StartCursor
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	return json.Marshal(toSerialize)
}

type NullablePaginationInfo struct {
	value *PaginationInfo
	isSet bool
}

func (v NullablePaginationInfo) Get() *PaginationInfo {
	return v.value
}

func (v *NullablePaginationInfo) Set(val *PaginationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationInfo(val *PaginationInfo) *NullablePaginationInfo {
	return &NullablePaginationInfo{value: val, isSet: true}
}

func (v NullablePaginationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

