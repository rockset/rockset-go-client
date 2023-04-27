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

// ExecutePublicQueryLambdaRequest struct for ExecutePublicQueryLambdaRequest
type ExecutePublicQueryLambdaRequest struct {
	// Row limit to use if no limit specified in the SQL query text.
	DefaultRowLimit *int32 `json:"default_row_limit,omitempty"`
	// Whether to generate warnings.
	GenerateWarnings *bool `json:"generate_warnings,omitempty"`
	// List of named parameters.
	Parameters []QueryParameter `json:"parameters,omitempty"`
}

// NewExecutePublicQueryLambdaRequest instantiates a new ExecutePublicQueryLambdaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutePublicQueryLambdaRequest() *ExecutePublicQueryLambdaRequest {
	this := ExecutePublicQueryLambdaRequest{}
	return &this
}

// NewExecutePublicQueryLambdaRequestWithDefaults instantiates a new ExecutePublicQueryLambdaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutePublicQueryLambdaRequestWithDefaults() *ExecutePublicQueryLambdaRequest {
	this := ExecutePublicQueryLambdaRequest{}
	return &this
}

// GetDefaultRowLimit returns the DefaultRowLimit field value if set, zero value otherwise.
func (o *ExecutePublicQueryLambdaRequest) GetDefaultRowLimit() int32 {
	if o == nil || o.DefaultRowLimit == nil {
		var ret int32
		return ret
	}
	return *o.DefaultRowLimit
}

// GetDefaultRowLimitOk returns a tuple with the DefaultRowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutePublicQueryLambdaRequest) GetDefaultRowLimitOk() (*int32, bool) {
	if o == nil || o.DefaultRowLimit == nil {
		return nil, false
	}
	return o.DefaultRowLimit, true
}

// HasDefaultRowLimit returns a boolean if a field has been set.
func (o *ExecutePublicQueryLambdaRequest) HasDefaultRowLimit() bool {
	if o != nil && o.DefaultRowLimit != nil {
		return true
	}

	return false
}

// SetDefaultRowLimit gets a reference to the given int32 and assigns it to the DefaultRowLimit field.
func (o *ExecutePublicQueryLambdaRequest) SetDefaultRowLimit(v int32) {
	o.DefaultRowLimit = &v
}

// GetGenerateWarnings returns the GenerateWarnings field value if set, zero value otherwise.
func (o *ExecutePublicQueryLambdaRequest) GetGenerateWarnings() bool {
	if o == nil || o.GenerateWarnings == nil {
		var ret bool
		return ret
	}
	return *o.GenerateWarnings
}

// GetGenerateWarningsOk returns a tuple with the GenerateWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutePublicQueryLambdaRequest) GetGenerateWarningsOk() (*bool, bool) {
	if o == nil || o.GenerateWarnings == nil {
		return nil, false
	}
	return o.GenerateWarnings, true
}

// HasGenerateWarnings returns a boolean if a field has been set.
func (o *ExecutePublicQueryLambdaRequest) HasGenerateWarnings() bool {
	if o != nil && o.GenerateWarnings != nil {
		return true
	}

	return false
}

// SetGenerateWarnings gets a reference to the given bool and assigns it to the GenerateWarnings field.
func (o *ExecutePublicQueryLambdaRequest) SetGenerateWarnings(v bool) {
	o.GenerateWarnings = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ExecutePublicQueryLambdaRequest) GetParameters() []QueryParameter {
	if o == nil || o.Parameters == nil {
		var ret []QueryParameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutePublicQueryLambdaRequest) GetParametersOk() ([]QueryParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ExecutePublicQueryLambdaRequest) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []QueryParameter and assigns it to the Parameters field.
func (o *ExecutePublicQueryLambdaRequest) SetParameters(v []QueryParameter) {
	o.Parameters = v
}

func (o ExecutePublicQueryLambdaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultRowLimit != nil {
		toSerialize["default_row_limit"] = o.DefaultRowLimit
	}
	if o.GenerateWarnings != nil {
		toSerialize["generate_warnings"] = o.GenerateWarnings
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableExecutePublicQueryLambdaRequest struct {
	value *ExecutePublicQueryLambdaRequest
	isSet bool
}

func (v NullableExecutePublicQueryLambdaRequest) Get() *ExecutePublicQueryLambdaRequest {
	return v.value
}

func (v *NullableExecutePublicQueryLambdaRequest) Set(val *ExecutePublicQueryLambdaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutePublicQueryLambdaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutePublicQueryLambdaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutePublicQueryLambdaRequest(val *ExecutePublicQueryLambdaRequest) *NullableExecutePublicQueryLambdaRequest {
	return &NullableExecutePublicQueryLambdaRequest{value: val, isSet: true}
}

func (v NullableExecutePublicQueryLambdaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutePublicQueryLambdaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


