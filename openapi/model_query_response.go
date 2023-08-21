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

// checks if the QueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryResponse{}

// QueryResponse struct for QueryResponse
type QueryResponse struct {
	// List of collections referenced in the query.
	Collections []string `json:"collections,omitempty"`
	// Meta information about each column in the result set. Not populated in `SELECT *` queries.
	ColumnFields []QueryFieldType `json:"column_fields,omitempty"`
	// If this was a write query, this is the log offset the query was written to.
	LastOffset *string `json:"last_offset,omitempty"`
	Pagination *PaginationInfo `json:"pagination,omitempty"`
	// Errors encountered while executing the query.
	QueryErrors []QueryError `json:"query_errors,omitempty"`
	// Unique ID for this query.
	QueryId *string `json:"query_id,omitempty"`
	// The full path of the executed query lambda. Includes version information.
	QueryLambdaPath *string `json:"query_lambda_path,omitempty"`
	// Results from the query.
	Results []map[string]interface{} `json:"results,omitempty"`
	// Number of results generated by the query.
	ResultsTotalDocCount *int64 `json:"results_total_doc_count,omitempty"`
	Stats *QueryResponseStats `json:"stats,omitempty"`
	// Status of query execution. Possible values: `QUEUED`, `RUNNING`, `COMPLETED`, `ERROR`.
	Status *string `json:"status,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// NewQueryResponse instantiates a new QueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryResponse() *QueryResponse {
	this := QueryResponse{}
	return &this
}

// NewQueryResponseWithDefaults instantiates a new QueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryResponseWithDefaults() *QueryResponse {
	this := QueryResponse{}
	return &this
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *QueryResponse) GetCollections() []string {
	if o == nil || IsNil(o.Collections) {
		var ret []string
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetCollectionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Collections) {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *QueryResponse) HasCollections() bool {
	if o != nil && !IsNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []string and assigns it to the Collections field.
func (o *QueryResponse) SetCollections(v []string) {
	o.Collections = v
}

// GetColumnFields returns the ColumnFields field value if set, zero value otherwise.
func (o *QueryResponse) GetColumnFields() []QueryFieldType {
	if o == nil || IsNil(o.ColumnFields) {
		var ret []QueryFieldType
		return ret
	}
	return o.ColumnFields
}

// GetColumnFieldsOk returns a tuple with the ColumnFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetColumnFieldsOk() ([]QueryFieldType, bool) {
	if o == nil || IsNil(o.ColumnFields) {
		return nil, false
	}
	return o.ColumnFields, true
}

// HasColumnFields returns a boolean if a field has been set.
func (o *QueryResponse) HasColumnFields() bool {
	if o != nil && !IsNil(o.ColumnFields) {
		return true
	}

	return false
}

// SetColumnFields gets a reference to the given []QueryFieldType and assigns it to the ColumnFields field.
func (o *QueryResponse) SetColumnFields(v []QueryFieldType) {
	o.ColumnFields = v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise.
func (o *QueryResponse) GetLastOffset() string {
	if o == nil || IsNil(o.LastOffset) {
		var ret string
		return ret
	}
	return *o.LastOffset
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetLastOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.LastOffset) {
		return nil, false
	}
	return o.LastOffset, true
}

// HasLastOffset returns a boolean if a field has been set.
func (o *QueryResponse) HasLastOffset() bool {
	if o != nil && !IsNil(o.LastOffset) {
		return true
	}

	return false
}

// SetLastOffset gets a reference to the given string and assigns it to the LastOffset field.
func (o *QueryResponse) SetLastOffset(v string) {
	o.LastOffset = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *QueryResponse) GetPagination() PaginationInfo {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationInfo
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetPaginationOk() (*PaginationInfo, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *QueryResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationInfo and assigns it to the Pagination field.
func (o *QueryResponse) SetPagination(v PaginationInfo) {
	o.Pagination = &v
}

// GetQueryErrors returns the QueryErrors field value if set, zero value otherwise.
func (o *QueryResponse) GetQueryErrors() []QueryError {
	if o == nil || IsNil(o.QueryErrors) {
		var ret []QueryError
		return ret
	}
	return o.QueryErrors
}

// GetQueryErrorsOk returns a tuple with the QueryErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetQueryErrorsOk() ([]QueryError, bool) {
	if o == nil || IsNil(o.QueryErrors) {
		return nil, false
	}
	return o.QueryErrors, true
}

// HasQueryErrors returns a boolean if a field has been set.
func (o *QueryResponse) HasQueryErrors() bool {
	if o != nil && !IsNil(o.QueryErrors) {
		return true
	}

	return false
}

// SetQueryErrors gets a reference to the given []QueryError and assigns it to the QueryErrors field.
func (o *QueryResponse) SetQueryErrors(v []QueryError) {
	o.QueryErrors = v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *QueryResponse) GetQueryId() string {
	if o == nil || IsNil(o.QueryId) {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetQueryIdOk() (*string, bool) {
	if o == nil || IsNil(o.QueryId) {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *QueryResponse) HasQueryId() bool {
	if o != nil && !IsNil(o.QueryId) {
		return true
	}

	return false
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *QueryResponse) SetQueryId(v string) {
	o.QueryId = &v
}

// GetQueryLambdaPath returns the QueryLambdaPath field value if set, zero value otherwise.
func (o *QueryResponse) GetQueryLambdaPath() string {
	if o == nil || IsNil(o.QueryLambdaPath) {
		var ret string
		return ret
	}
	return *o.QueryLambdaPath
}

// GetQueryLambdaPathOk returns a tuple with the QueryLambdaPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetQueryLambdaPathOk() (*string, bool) {
	if o == nil || IsNil(o.QueryLambdaPath) {
		return nil, false
	}
	return o.QueryLambdaPath, true
}

// HasQueryLambdaPath returns a boolean if a field has been set.
func (o *QueryResponse) HasQueryLambdaPath() bool {
	if o != nil && !IsNil(o.QueryLambdaPath) {
		return true
	}

	return false
}

// SetQueryLambdaPath gets a reference to the given string and assigns it to the QueryLambdaPath field.
func (o *QueryResponse) SetQueryLambdaPath(v string) {
	o.QueryLambdaPath = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *QueryResponse) GetResults() []map[string]interface{} {
	if o == nil || IsNil(o.Results) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetResultsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *QueryResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []map[string]interface{} and assigns it to the Results field.
func (o *QueryResponse) SetResults(v []map[string]interface{}) {
	o.Results = v
}

// GetResultsTotalDocCount returns the ResultsTotalDocCount field value if set, zero value otherwise.
func (o *QueryResponse) GetResultsTotalDocCount() int64 {
	if o == nil || IsNil(o.ResultsTotalDocCount) {
		var ret int64
		return ret
	}
	return *o.ResultsTotalDocCount
}

// GetResultsTotalDocCountOk returns a tuple with the ResultsTotalDocCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetResultsTotalDocCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ResultsTotalDocCount) {
		return nil, false
	}
	return o.ResultsTotalDocCount, true
}

// HasResultsTotalDocCount returns a boolean if a field has been set.
func (o *QueryResponse) HasResultsTotalDocCount() bool {
	if o != nil && !IsNil(o.ResultsTotalDocCount) {
		return true
	}

	return false
}

// SetResultsTotalDocCount gets a reference to the given int64 and assigns it to the ResultsTotalDocCount field.
func (o *QueryResponse) SetResultsTotalDocCount(v int64) {
	o.ResultsTotalDocCount = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *QueryResponse) GetStats() QueryResponseStats {
	if o == nil || IsNil(o.Stats) {
		var ret QueryResponseStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetStatsOk() (*QueryResponseStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *QueryResponse) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given QueryResponseStats and assigns it to the Stats field.
func (o *QueryResponse) SetStats(v QueryResponseStats) {
	o.Stats = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryResponse) SetStatus(v string) {
	o.Status = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *QueryResponse) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponse) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *QueryResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *QueryResponse) SetWarnings(v []string) {
	o.Warnings = v
}

func (o QueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collections) {
		toSerialize["collections"] = o.Collections
	}
	if !IsNil(o.ColumnFields) {
		toSerialize["column_fields"] = o.ColumnFields
	}
	if !IsNil(o.LastOffset) {
		toSerialize["last_offset"] = o.LastOffset
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.QueryErrors) {
		toSerialize["query_errors"] = o.QueryErrors
	}
	if !IsNil(o.QueryId) {
		toSerialize["query_id"] = o.QueryId
	}
	if !IsNil(o.QueryLambdaPath) {
		toSerialize["query_lambda_path"] = o.QueryLambdaPath
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.ResultsTotalDocCount) {
		toSerialize["results_total_doc_count"] = o.ResultsTotalDocCount
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableQueryResponse struct {
	value *QueryResponse
	isSet bool
}

func (v NullableQueryResponse) Get() *QueryResponse {
	return v.value
}

func (v *NullableQueryResponse) Set(val *QueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryResponse(val *QueryResponse) *NullableQueryResponse {
	return &NullableQueryResponse{value: val, isSet: true}
}

func (v NullableQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


