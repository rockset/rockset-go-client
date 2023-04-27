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

// ErrorModel Describes details about an error
type ErrorModel struct {
	// Column where the error happened (if applicable).
	Column *int32 `json:"column,omitempty"`
	// ID of the error.
	ErrorId *string `json:"error_id,omitempty"`
	// Line where the error happened (if applicable).
	Line *int32 `json:"line,omitempty"`
	// Descriptive message about the error.
	Message *string `json:"message,omitempty"`
	// ID of the query (if applicable).
	QueryId *string `json:"query_id,omitempty"`
	// Internal trace ID to help with debugging.
	TraceId *string `json:"trace_id,omitempty"`
	// Category of the error.
	Type *string `json:"type,omitempty"`
}

// NewErrorModel instantiates a new ErrorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorModel() *ErrorModel {
	this := ErrorModel{}
	return &this
}

// NewErrorModelWithDefaults instantiates a new ErrorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorModelWithDefaults() *ErrorModel {
	this := ErrorModel{}
	return &this
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *ErrorModel) GetColumn() int32 {
	if o == nil || o.Column == nil {
		var ret int32
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetColumnOk() (*int32, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *ErrorModel) HasColumn() bool {
	if o != nil && o.Column != nil {
		return true
	}

	return false
}

// SetColumn gets a reference to the given int32 and assigns it to the Column field.
func (o *ErrorModel) SetColumn(v int32) {
	o.Column = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *ErrorModel) GetErrorId() string {
	if o == nil || o.ErrorId == nil {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetErrorIdOk() (*string, bool) {
	if o == nil || o.ErrorId == nil {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *ErrorModel) HasErrorId() bool {
	if o != nil && o.ErrorId != nil {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *ErrorModel) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *ErrorModel) GetLine() int32 {
	if o == nil || o.Line == nil {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetLineOk() (*int32, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *ErrorModel) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *ErrorModel) SetLine(v int32) {
	o.Line = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorModel) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorModel) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorModel) SetMessage(v string) {
	o.Message = &v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *ErrorModel) GetQueryId() string {
	if o == nil || o.QueryId == nil {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetQueryIdOk() (*string, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *ErrorModel) HasQueryId() bool {
	if o != nil && o.QueryId != nil {
		return true
	}

	return false
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *ErrorModel) SetQueryId(v string) {
	o.QueryId = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *ErrorModel) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *ErrorModel) HasTraceId() bool {
	if o != nil && o.TraceId != nil {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *ErrorModel) SetTraceId(v string) {
	o.TraceId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorModel) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorModel) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorModel) SetType(v string) {
	o.Type = &v
}

func (o ErrorModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	if o.ErrorId != nil {
		toSerialize["error_id"] = o.ErrorId
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.QueryId != nil {
		toSerialize["query_id"] = o.QueryId
	}
	if o.TraceId != nil {
		toSerialize["trace_id"] = o.TraceId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableErrorModel struct {
	value *ErrorModel
	isSet bool
}

func (v NullableErrorModel) Get() *ErrorModel {
	return v.value
}

func (v *NullableErrorModel) Set(val *ErrorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorModel(val *ErrorModel) *NullableErrorModel {
	return &NullableErrorModel{value: val, isSet: true}
}

func (v NullableErrorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


