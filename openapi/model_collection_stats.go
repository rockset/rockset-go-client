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

// CollectionStats struct for CollectionStats
type CollectionStats struct {
	// Number of documents in the collection.
	DocCount *int64 `json:"doc_count,omitempty"`
	// Number of documents purged from the collection.
	PurgedDocCount *int64 `json:"purged_doc_count,omitempty"`
	// Number between 0 and 1 that indicates progress of collection creation.
	FillProgress *float64 `json:"fill_progress,omitempty"`
	// Milliseconds since Unix epoch Jan 1, 1970.
	LastQueriedMs *int64 `json:"last_queried_ms,omitempty"`
	// Milliseconds since Unix epoch Jan 1, 1970.
	LastUpdatedMs *int64 `json:"last_updated_ms,omitempty"`
	// Total collection size in bytes.
	TotalSize *int64 `json:"total_size,omitempty"`
	// Total collection index size in bytes.
	TotalIndexSize *int64 `json:"total_index_size,omitempty"`
	// Total collection row index size in bytes.
	RowIndexSize *int64 `json:"row_index_size,omitempty"`
	// Total collection column index size in bytes.
	ColumnIndexSize *int64 `json:"column_index_size,omitempty"`
	// Total collection inverted index size in bytes.
	InvertedIndexSize *int64 `json:"inverted_index_size,omitempty"`
	// Total collection range index size in bytes.
	RangeIndexSize *int64 `json:"range_index_size,omitempty"`
	// Total size of bytes purged in bytes.
	PurgedDocSize *int64 `json:"purged_doc_size,omitempty"`
	// Total number of bytes inserted into the collection.
	BytesInserted *int64 `json:"bytes_inserted,omitempty"`
	// Total number of bytes overwritten in writing into the collection.
	BytesOverwritten *int64 `json:"bytes_overwritten,omitempty"`
	// Total number of bytes inserted into the collection during bulk.
	BulkBytesInserted *int64 `json:"bulk_bytes_inserted,omitempty"`
	// Total number of bytes overwritten in writing into the collection during bulk.
	BulkBytesOverwritten *int64 `json:"bulk_bytes_overwritten,omitempty"`
}

// NewCollectionStats instantiates a new CollectionStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionStats() *CollectionStats {
	this := CollectionStats{}
	return &this
}

// NewCollectionStatsWithDefaults instantiates a new CollectionStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionStatsWithDefaults() *CollectionStats {
	this := CollectionStats{}
	return &this
}

// GetDocCount returns the DocCount field value if set, zero value otherwise.
func (o *CollectionStats) GetDocCount() int64 {
	if o == nil || o.DocCount == nil {
		var ret int64
		return ret
	}
	return *o.DocCount
}

// GetDocCountOk returns a tuple with the DocCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetDocCountOk() (*int64, bool) {
	if o == nil || o.DocCount == nil {
		return nil, false
	}
	return o.DocCount, true
}

// HasDocCount returns a boolean if a field has been set.
func (o *CollectionStats) HasDocCount() bool {
	if o != nil && o.DocCount != nil {
		return true
	}

	return false
}

// SetDocCount gets a reference to the given int64 and assigns it to the DocCount field.
func (o *CollectionStats) SetDocCount(v int64) {
	o.DocCount = &v
}

// GetPurgedDocCount returns the PurgedDocCount field value if set, zero value otherwise.
func (o *CollectionStats) GetPurgedDocCount() int64 {
	if o == nil || o.PurgedDocCount == nil {
		var ret int64
		return ret
	}
	return *o.PurgedDocCount
}

// GetPurgedDocCountOk returns a tuple with the PurgedDocCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetPurgedDocCountOk() (*int64, bool) {
	if o == nil || o.PurgedDocCount == nil {
		return nil, false
	}
	return o.PurgedDocCount, true
}

// HasPurgedDocCount returns a boolean if a field has been set.
func (o *CollectionStats) HasPurgedDocCount() bool {
	if o != nil && o.PurgedDocCount != nil {
		return true
	}

	return false
}

// SetPurgedDocCount gets a reference to the given int64 and assigns it to the PurgedDocCount field.
func (o *CollectionStats) SetPurgedDocCount(v int64) {
	o.PurgedDocCount = &v
}

// GetFillProgress returns the FillProgress field value if set, zero value otherwise.
func (o *CollectionStats) GetFillProgress() float64 {
	if o == nil || o.FillProgress == nil {
		var ret float64
		return ret
	}
	return *o.FillProgress
}

// GetFillProgressOk returns a tuple with the FillProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetFillProgressOk() (*float64, bool) {
	if o == nil || o.FillProgress == nil {
		return nil, false
	}
	return o.FillProgress, true
}

// HasFillProgress returns a boolean if a field has been set.
func (o *CollectionStats) HasFillProgress() bool {
	if o != nil && o.FillProgress != nil {
		return true
	}

	return false
}

// SetFillProgress gets a reference to the given float64 and assigns it to the FillProgress field.
func (o *CollectionStats) SetFillProgress(v float64) {
	o.FillProgress = &v
}

// GetLastQueriedMs returns the LastQueriedMs field value if set, zero value otherwise.
func (o *CollectionStats) GetLastQueriedMs() int64 {
	if o == nil || o.LastQueriedMs == nil {
		var ret int64
		return ret
	}
	return *o.LastQueriedMs
}

// GetLastQueriedMsOk returns a tuple with the LastQueriedMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetLastQueriedMsOk() (*int64, bool) {
	if o == nil || o.LastQueriedMs == nil {
		return nil, false
	}
	return o.LastQueriedMs, true
}

// HasLastQueriedMs returns a boolean if a field has been set.
func (o *CollectionStats) HasLastQueriedMs() bool {
	if o != nil && o.LastQueriedMs != nil {
		return true
	}

	return false
}

// SetLastQueriedMs gets a reference to the given int64 and assigns it to the LastQueriedMs field.
func (o *CollectionStats) SetLastQueriedMs(v int64) {
	o.LastQueriedMs = &v
}

// GetLastUpdatedMs returns the LastUpdatedMs field value if set, zero value otherwise.
func (o *CollectionStats) GetLastUpdatedMs() int64 {
	if o == nil || o.LastUpdatedMs == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedMs
}

// GetLastUpdatedMsOk returns a tuple with the LastUpdatedMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetLastUpdatedMsOk() (*int64, bool) {
	if o == nil || o.LastUpdatedMs == nil {
		return nil, false
	}
	return o.LastUpdatedMs, true
}

// HasLastUpdatedMs returns a boolean if a field has been set.
func (o *CollectionStats) HasLastUpdatedMs() bool {
	if o != nil && o.LastUpdatedMs != nil {
		return true
	}

	return false
}

// SetLastUpdatedMs gets a reference to the given int64 and assigns it to the LastUpdatedMs field.
func (o *CollectionStats) SetLastUpdatedMs(v int64) {
	o.LastUpdatedMs = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *CollectionStats) GetTotalSize() int64 {
	if o == nil || o.TotalSize == nil {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetTotalSizeOk() (*int64, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *CollectionStats) HasTotalSize() bool {
	if o != nil && o.TotalSize != nil {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *CollectionStats) SetTotalSize(v int64) {
	o.TotalSize = &v
}

// GetTotalIndexSize returns the TotalIndexSize field value if set, zero value otherwise.
func (o *CollectionStats) GetTotalIndexSize() int64 {
	if o == nil || o.TotalIndexSize == nil {
		var ret int64
		return ret
	}
	return *o.TotalIndexSize
}

// GetTotalIndexSizeOk returns a tuple with the TotalIndexSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetTotalIndexSizeOk() (*int64, bool) {
	if o == nil || o.TotalIndexSize == nil {
		return nil, false
	}
	return o.TotalIndexSize, true
}

// HasTotalIndexSize returns a boolean if a field has been set.
func (o *CollectionStats) HasTotalIndexSize() bool {
	if o != nil && o.TotalIndexSize != nil {
		return true
	}

	return false
}

// SetTotalIndexSize gets a reference to the given int64 and assigns it to the TotalIndexSize field.
func (o *CollectionStats) SetTotalIndexSize(v int64) {
	o.TotalIndexSize = &v
}

// GetRowIndexSize returns the RowIndexSize field value if set, zero value otherwise.
func (o *CollectionStats) GetRowIndexSize() int64 {
	if o == nil || o.RowIndexSize == nil {
		var ret int64
		return ret
	}
	return *o.RowIndexSize
}

// GetRowIndexSizeOk returns a tuple with the RowIndexSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetRowIndexSizeOk() (*int64, bool) {
	if o == nil || o.RowIndexSize == nil {
		return nil, false
	}
	return o.RowIndexSize, true
}

// HasRowIndexSize returns a boolean if a field has been set.
func (o *CollectionStats) HasRowIndexSize() bool {
	if o != nil && o.RowIndexSize != nil {
		return true
	}

	return false
}

// SetRowIndexSize gets a reference to the given int64 and assigns it to the RowIndexSize field.
func (o *CollectionStats) SetRowIndexSize(v int64) {
	o.RowIndexSize = &v
}

// GetColumnIndexSize returns the ColumnIndexSize field value if set, zero value otherwise.
func (o *CollectionStats) GetColumnIndexSize() int64 {
	if o == nil || o.ColumnIndexSize == nil {
		var ret int64
		return ret
	}
	return *o.ColumnIndexSize
}

// GetColumnIndexSizeOk returns a tuple with the ColumnIndexSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetColumnIndexSizeOk() (*int64, bool) {
	if o == nil || o.ColumnIndexSize == nil {
		return nil, false
	}
	return o.ColumnIndexSize, true
}

// HasColumnIndexSize returns a boolean if a field has been set.
func (o *CollectionStats) HasColumnIndexSize() bool {
	if o != nil && o.ColumnIndexSize != nil {
		return true
	}

	return false
}

// SetColumnIndexSize gets a reference to the given int64 and assigns it to the ColumnIndexSize field.
func (o *CollectionStats) SetColumnIndexSize(v int64) {
	o.ColumnIndexSize = &v
}

// GetInvertedIndexSize returns the InvertedIndexSize field value if set, zero value otherwise.
func (o *CollectionStats) GetInvertedIndexSize() int64 {
	if o == nil || o.InvertedIndexSize == nil {
		var ret int64
		return ret
	}
	return *o.InvertedIndexSize
}

// GetInvertedIndexSizeOk returns a tuple with the InvertedIndexSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetInvertedIndexSizeOk() (*int64, bool) {
	if o == nil || o.InvertedIndexSize == nil {
		return nil, false
	}
	return o.InvertedIndexSize, true
}

// HasInvertedIndexSize returns a boolean if a field has been set.
func (o *CollectionStats) HasInvertedIndexSize() bool {
	if o != nil && o.InvertedIndexSize != nil {
		return true
	}

	return false
}

// SetInvertedIndexSize gets a reference to the given int64 and assigns it to the InvertedIndexSize field.
func (o *CollectionStats) SetInvertedIndexSize(v int64) {
	o.InvertedIndexSize = &v
}

// GetRangeIndexSize returns the RangeIndexSize field value if set, zero value otherwise.
func (o *CollectionStats) GetRangeIndexSize() int64 {
	if o == nil || o.RangeIndexSize == nil {
		var ret int64
		return ret
	}
	return *o.RangeIndexSize
}

// GetRangeIndexSizeOk returns a tuple with the RangeIndexSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetRangeIndexSizeOk() (*int64, bool) {
	if o == nil || o.RangeIndexSize == nil {
		return nil, false
	}
	return o.RangeIndexSize, true
}

// HasRangeIndexSize returns a boolean if a field has been set.
func (o *CollectionStats) HasRangeIndexSize() bool {
	if o != nil && o.RangeIndexSize != nil {
		return true
	}

	return false
}

// SetRangeIndexSize gets a reference to the given int64 and assigns it to the RangeIndexSize field.
func (o *CollectionStats) SetRangeIndexSize(v int64) {
	o.RangeIndexSize = &v
}

// GetPurgedDocSize returns the PurgedDocSize field value if set, zero value otherwise.
func (o *CollectionStats) GetPurgedDocSize() int64 {
	if o == nil || o.PurgedDocSize == nil {
		var ret int64
		return ret
	}
	return *o.PurgedDocSize
}

// GetPurgedDocSizeOk returns a tuple with the PurgedDocSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetPurgedDocSizeOk() (*int64, bool) {
	if o == nil || o.PurgedDocSize == nil {
		return nil, false
	}
	return o.PurgedDocSize, true
}

// HasPurgedDocSize returns a boolean if a field has been set.
func (o *CollectionStats) HasPurgedDocSize() bool {
	if o != nil && o.PurgedDocSize != nil {
		return true
	}

	return false
}

// SetPurgedDocSize gets a reference to the given int64 and assigns it to the PurgedDocSize field.
func (o *CollectionStats) SetPurgedDocSize(v int64) {
	o.PurgedDocSize = &v
}

// GetBytesInserted returns the BytesInserted field value if set, zero value otherwise.
func (o *CollectionStats) GetBytesInserted() int64 {
	if o == nil || o.BytesInserted == nil {
		var ret int64
		return ret
	}
	return *o.BytesInserted
}

// GetBytesInsertedOk returns a tuple with the BytesInserted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetBytesInsertedOk() (*int64, bool) {
	if o == nil || o.BytesInserted == nil {
		return nil, false
	}
	return o.BytesInserted, true
}

// HasBytesInserted returns a boolean if a field has been set.
func (o *CollectionStats) HasBytesInserted() bool {
	if o != nil && o.BytesInserted != nil {
		return true
	}

	return false
}

// SetBytesInserted gets a reference to the given int64 and assigns it to the BytesInserted field.
func (o *CollectionStats) SetBytesInserted(v int64) {
	o.BytesInserted = &v
}

// GetBytesOverwritten returns the BytesOverwritten field value if set, zero value otherwise.
func (o *CollectionStats) GetBytesOverwritten() int64 {
	if o == nil || o.BytesOverwritten == nil {
		var ret int64
		return ret
	}
	return *o.BytesOverwritten
}

// GetBytesOverwrittenOk returns a tuple with the BytesOverwritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetBytesOverwrittenOk() (*int64, bool) {
	if o == nil || o.BytesOverwritten == nil {
		return nil, false
	}
	return o.BytesOverwritten, true
}

// HasBytesOverwritten returns a boolean if a field has been set.
func (o *CollectionStats) HasBytesOverwritten() bool {
	if o != nil && o.BytesOverwritten != nil {
		return true
	}

	return false
}

// SetBytesOverwritten gets a reference to the given int64 and assigns it to the BytesOverwritten field.
func (o *CollectionStats) SetBytesOverwritten(v int64) {
	o.BytesOverwritten = &v
}

// GetBulkBytesInserted returns the BulkBytesInserted field value if set, zero value otherwise.
func (o *CollectionStats) GetBulkBytesInserted() int64 {
	if o == nil || o.BulkBytesInserted == nil {
		var ret int64
		return ret
	}
	return *o.BulkBytesInserted
}

// GetBulkBytesInsertedOk returns a tuple with the BulkBytesInserted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetBulkBytesInsertedOk() (*int64, bool) {
	if o == nil || o.BulkBytesInserted == nil {
		return nil, false
	}
	return o.BulkBytesInserted, true
}

// HasBulkBytesInserted returns a boolean if a field has been set.
func (o *CollectionStats) HasBulkBytesInserted() bool {
	if o != nil && o.BulkBytesInserted != nil {
		return true
	}

	return false
}

// SetBulkBytesInserted gets a reference to the given int64 and assigns it to the BulkBytesInserted field.
func (o *CollectionStats) SetBulkBytesInserted(v int64) {
	o.BulkBytesInserted = &v
}

// GetBulkBytesOverwritten returns the BulkBytesOverwritten field value if set, zero value otherwise.
func (o *CollectionStats) GetBulkBytesOverwritten() int64 {
	if o == nil || o.BulkBytesOverwritten == nil {
		var ret int64
		return ret
	}
	return *o.BulkBytesOverwritten
}

// GetBulkBytesOverwrittenOk returns a tuple with the BulkBytesOverwritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionStats) GetBulkBytesOverwrittenOk() (*int64, bool) {
	if o == nil || o.BulkBytesOverwritten == nil {
		return nil, false
	}
	return o.BulkBytesOverwritten, true
}

// HasBulkBytesOverwritten returns a boolean if a field has been set.
func (o *CollectionStats) HasBulkBytesOverwritten() bool {
	if o != nil && o.BulkBytesOverwritten != nil {
		return true
	}

	return false
}

// SetBulkBytesOverwritten gets a reference to the given int64 and assigns it to the BulkBytesOverwritten field.
func (o *CollectionStats) SetBulkBytesOverwritten(v int64) {
	o.BulkBytesOverwritten = &v
}

func (o CollectionStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocCount != nil {
		toSerialize["doc_count"] = o.DocCount
	}
	if o.PurgedDocCount != nil {
		toSerialize["purged_doc_count"] = o.PurgedDocCount
	}
	if o.FillProgress != nil {
		toSerialize["fill_progress"] = o.FillProgress
	}
	if o.LastQueriedMs != nil {
		toSerialize["last_queried_ms"] = o.LastQueriedMs
	}
	if o.LastUpdatedMs != nil {
		toSerialize["last_updated_ms"] = o.LastUpdatedMs
	}
	if o.TotalSize != nil {
		toSerialize["total_size"] = o.TotalSize
	}
	if o.TotalIndexSize != nil {
		toSerialize["total_index_size"] = o.TotalIndexSize
	}
	if o.RowIndexSize != nil {
		toSerialize["row_index_size"] = o.RowIndexSize
	}
	if o.ColumnIndexSize != nil {
		toSerialize["column_index_size"] = o.ColumnIndexSize
	}
	if o.InvertedIndexSize != nil {
		toSerialize["inverted_index_size"] = o.InvertedIndexSize
	}
	if o.RangeIndexSize != nil {
		toSerialize["range_index_size"] = o.RangeIndexSize
	}
	if o.PurgedDocSize != nil {
		toSerialize["purged_doc_size"] = o.PurgedDocSize
	}
	if o.BytesInserted != nil {
		toSerialize["bytes_inserted"] = o.BytesInserted
	}
	if o.BytesOverwritten != nil {
		toSerialize["bytes_overwritten"] = o.BytesOverwritten
	}
	if o.BulkBytesInserted != nil {
		toSerialize["bulk_bytes_inserted"] = o.BulkBytesInserted
	}
	if o.BulkBytesOverwritten != nil {
		toSerialize["bulk_bytes_overwritten"] = o.BulkBytesOverwritten
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionStats struct {
	value *CollectionStats
	isSet bool
}

func (v NullableCollectionStats) Get() *CollectionStats {
	return v.value
}

func (v *NullableCollectionStats) Set(val *CollectionStats) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionStats) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionStats(val *CollectionStats) *NullableCollectionStats {
	return &NullableCollectionStats{value: val, isSet: true}
}

func (v NullableCollectionStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


