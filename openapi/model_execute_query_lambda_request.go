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

// ExecuteQueryLambdaRequest struct for ExecuteQueryLambdaRequest
type ExecuteQueryLambdaRequest struct {
	// list of named parameters
	Parameters *[]QueryParameter `json:"parameters,omitempty"`
	// Row limit to use if no limit specified in the SQL query text
	DefaultRowLimit *int32 `json:"default_row_limit,omitempty"`
	// Whether to generate warnings
	GenerateWarnings *bool `json:"generate_warnings,omitempty"`
	// Flag to paginate and store the results of this query for later / sequential retrieval.
	Paginate *bool `json:"paginate,omitempty"`
	// Number of documents to return in addition to paginating for this query call. Only relevant if `paginate` flag is also set.
	InitialPaginateResponseDocCount *int32 `json:"initial_paginate_response_doc_count,omitempty"`
}

// NewExecuteQueryLambdaRequest instantiates a new ExecuteQueryLambdaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteQueryLambdaRequest() *ExecuteQueryLambdaRequest {
	this := ExecuteQueryLambdaRequest{}
	return &this
}

// NewExecuteQueryLambdaRequestWithDefaults instantiates a new ExecuteQueryLambdaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteQueryLambdaRequestWithDefaults() *ExecuteQueryLambdaRequest {
	this := ExecuteQueryLambdaRequest{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ExecuteQueryLambdaRequest) GetParameters() []QueryParameter {
	if o == nil || o.Parameters == nil {
		var ret []QueryParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteQueryLambdaRequest) GetParametersOk() (*[]QueryParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ExecuteQueryLambdaRequest) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []QueryParameter and assigns it to the Parameters field.
func (o *ExecuteQueryLambdaRequest) SetParameters(v []QueryParameter) {
	o.Parameters = &v
}

// GetDefaultRowLimit returns the DefaultRowLimit field value if set, zero value otherwise.
func (o *ExecuteQueryLambdaRequest) GetDefaultRowLimit() int32 {
	if o == nil || o.DefaultRowLimit == nil {
		var ret int32
		return ret
	}
	return *o.DefaultRowLimit
}

// GetDefaultRowLimitOk returns a tuple with the DefaultRowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteQueryLambdaRequest) GetDefaultRowLimitOk() (*int32, bool) {
	if o == nil || o.DefaultRowLimit == nil {
		return nil, false
	}
	return o.DefaultRowLimit, true
}

// HasDefaultRowLimit returns a boolean if a field has been set.
func (o *ExecuteQueryLambdaRequest) HasDefaultRowLimit() bool {
	if o != nil && o.DefaultRowLimit != nil {
		return true
	}

	return false
}

// SetDefaultRowLimit gets a reference to the given int32 and assigns it to the DefaultRowLimit field.
func (o *ExecuteQueryLambdaRequest) SetDefaultRowLimit(v int32) {
	o.DefaultRowLimit = &v
}

// GetGenerateWarnings returns the GenerateWarnings field value if set, zero value otherwise.
func (o *ExecuteQueryLambdaRequest) GetGenerateWarnings() bool {
	if o == nil || o.GenerateWarnings == nil {
		var ret bool
		return ret
	}
	return *o.GenerateWarnings
}

// GetGenerateWarningsOk returns a tuple with the GenerateWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteQueryLambdaRequest) GetGenerateWarningsOk() (*bool, bool) {
	if o == nil || o.GenerateWarnings == nil {
		return nil, false
	}
	return o.GenerateWarnings, true
}

// HasGenerateWarnings returns a boolean if a field has been set.
func (o *ExecuteQueryLambdaRequest) HasGenerateWarnings() bool {
	if o != nil && o.GenerateWarnings != nil {
		return true
	}

	return false
}

// SetGenerateWarnings gets a reference to the given bool and assigns it to the GenerateWarnings field.
func (o *ExecuteQueryLambdaRequest) SetGenerateWarnings(v bool) {
	o.GenerateWarnings = &v
}

// GetPaginate returns the Paginate field value if set, zero value otherwise.
func (o *ExecuteQueryLambdaRequest) GetPaginate() bool {
	if o == nil || o.Paginate == nil {
		var ret bool
		return ret
	}
	return *o.Paginate
}

// GetPaginateOk returns a tuple with the Paginate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteQueryLambdaRequest) GetPaginateOk() (*bool, bool) {
	if o == nil || o.Paginate == nil {
		return nil, false
	}
	return o.Paginate, true
}

// HasPaginate returns a boolean if a field has been set.
func (o *ExecuteQueryLambdaRequest) HasPaginate() bool {
	if o != nil && o.Paginate != nil {
		return true
	}

	return false
}

// SetPaginate gets a reference to the given bool and assigns it to the Paginate field.
func (o *ExecuteQueryLambdaRequest) SetPaginate(v bool) {
	o.Paginate = &v
}

// GetInitialPaginateResponseDocCount returns the InitialPaginateResponseDocCount field value if set, zero value otherwise.
func (o *ExecuteQueryLambdaRequest) GetInitialPaginateResponseDocCount() int32 {
	if o == nil || o.InitialPaginateResponseDocCount == nil {
		var ret int32
		return ret
	}
	return *o.InitialPaginateResponseDocCount
}

// GetInitialPaginateResponseDocCountOk returns a tuple with the InitialPaginateResponseDocCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteQueryLambdaRequest) GetInitialPaginateResponseDocCountOk() (*int32, bool) {
	if o == nil || o.InitialPaginateResponseDocCount == nil {
		return nil, false
	}
	return o.InitialPaginateResponseDocCount, true
}

// HasInitialPaginateResponseDocCount returns a boolean if a field has been set.
func (o *ExecuteQueryLambdaRequest) HasInitialPaginateResponseDocCount() bool {
	if o != nil && o.InitialPaginateResponseDocCount != nil {
		return true
	}

	return false
}

// SetInitialPaginateResponseDocCount gets a reference to the given int32 and assigns it to the InitialPaginateResponseDocCount field.
func (o *ExecuteQueryLambdaRequest) SetInitialPaginateResponseDocCount(v int32) {
	o.InitialPaginateResponseDocCount = &v
}

func (o ExecuteQueryLambdaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.DefaultRowLimit != nil {
		toSerialize["default_row_limit"] = o.DefaultRowLimit
	}
	if o.GenerateWarnings != nil {
		toSerialize["generate_warnings"] = o.GenerateWarnings
	}
	if o.Paginate != nil {
		toSerialize["paginate"] = o.Paginate
	}
	if o.InitialPaginateResponseDocCount != nil {
		toSerialize["initial_paginate_response_doc_count"] = o.InitialPaginateResponseDocCount
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteQueryLambdaRequest struct {
	value *ExecuteQueryLambdaRequest
	isSet bool
}

func (v NullableExecuteQueryLambdaRequest) Get() *ExecuteQueryLambdaRequest {
	return v.value
}

func (v *NullableExecuteQueryLambdaRequest) Set(val *ExecuteQueryLambdaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteQueryLambdaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteQueryLambdaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteQueryLambdaRequest(val *ExecuteQueryLambdaRequest) *NullableExecuteQueryLambdaRequest {
	return &NullableExecuteQueryLambdaRequest{value: val, isSet: true}
}

func (v NullableExecuteQueryLambdaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteQueryLambdaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


