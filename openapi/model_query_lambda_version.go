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

// QueryLambdaVersion struct for QueryLambdaVersion
type QueryLambdaVersion struct {
	// workspace of this Query Lambda
	Workspace *string `json:"workspace,omitempty"`
	// user that created this Query Lambda
	CreatedBy *string `json:"created_by,omitempty"`
	// ISO-8601 date of when Query Lambda was created
	CreatedAt *string `json:"created_at,omitempty"`
	// Query Lambda name
	Name *string `json:"name,omitempty"`
	// Query Lambda version
	Version *string `json:"version,omitempty"`
	// optional description
	Description *string `json:"description,omitempty"`
	Sql *QueryLambdaSql `json:"sql,omitempty"`
	// collections queried by underlying SQL query
	Collections []string `json:"collections,omitempty"`
	// status of this Query Lambda
	State *string `json:"state,omitempty"`
	Stats *QueryLambdaStats `json:"stats,omitempty"`
}

// NewQueryLambdaVersion instantiates a new QueryLambdaVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLambdaVersion() *QueryLambdaVersion {
	this := QueryLambdaVersion{}
	return &this
}

// NewQueryLambdaVersionWithDefaults instantiates a new QueryLambdaVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLambdaVersionWithDefaults() *QueryLambdaVersion {
	this := QueryLambdaVersion{}
	return &this
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetWorkspace() string {
	if o == nil || o.Workspace == nil {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetWorkspaceOk() (*string, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *QueryLambdaVersion) SetWorkspace(v string) {
	o.Workspace = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *QueryLambdaVersion) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *QueryLambdaVersion) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueryLambdaVersion) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *QueryLambdaVersion) SetVersion(v string) {
	o.Version = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QueryLambdaVersion) SetDescription(v string) {
	o.Description = &v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetSql() QueryLambdaSql {
	if o == nil || o.Sql == nil {
		var ret QueryLambdaSql
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetSqlOk() (*QueryLambdaSql, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasSql() bool {
	if o != nil && o.Sql != nil {
		return true
	}

	return false
}

// SetSql gets a reference to the given QueryLambdaSql and assigns it to the Sql field.
func (o *QueryLambdaVersion) SetSql(v QueryLambdaSql) {
	o.Sql = &v
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetCollections() []string {
	if o == nil || o.Collections == nil {
		var ret []string
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetCollectionsOk() ([]string, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []string and assigns it to the Collections field.
func (o *QueryLambdaVersion) SetCollections(v []string) {
	o.Collections = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *QueryLambdaVersion) SetState(v string) {
	o.State = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *QueryLambdaVersion) GetStats() QueryLambdaStats {
	if o == nil || o.Stats == nil {
		var ret QueryLambdaStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaVersion) GetStatsOk() (*QueryLambdaStats, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *QueryLambdaVersion) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given QueryLambdaStats and assigns it to the Stats field.
func (o *QueryLambdaVersion) SetStats(v QueryLambdaStats) {
	o.Stats = &v
}

func (o QueryLambdaVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.Collections != nil {
		toSerialize["collections"] = o.Collections
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}
	return json.Marshal(toSerialize)
}

type NullableQueryLambdaVersion struct {
	value *QueryLambdaVersion
	isSet bool
}

func (v NullableQueryLambdaVersion) Get() *QueryLambdaVersion {
	return v.value
}

func (v *NullableQueryLambdaVersion) Set(val *QueryLambdaVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLambdaVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLambdaVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLambdaVersion(val *QueryLambdaVersion) *NullableQueryLambdaVersion {
	return &NullableQueryLambdaVersion{value: val, isSet: true}
}

func (v NullableQueryLambdaVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLambdaVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


