/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace Environment API v1. To read about use cases and examples, refer to the [help page](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// LogJobStatusResult The status the log analysis job.
type LogJobStatusResult struct {
	// The status of the log analysis job.
	LogAnalysisStatus *string `json:"logAnalysisStatus,omitempty"`
	// The timestamp of the last status change, in UTC milliseconds.
	StatusChangeTimestamp *int64 `json:"statusChangeTimestamp,omitempty"`
	// The cause of the job failure.    A successful job has the `NONE` value.
	LogHandlingError *string `json:"logHandlingError,omitempty"`
	// The number of analyzed log entries.
	RecordsTotal *int32 `json:"recordsTotal,omitempty"`
	// The map of the log entry sortable fields.
	SortableFields *[]string `json:"sortableFields,omitempty"`
	// The map of the log entry filterable fields.
	FilterableFields *[]string `json:"filterableFields,omitempty"`
}

// NewLogJobStatusResult instantiates a new LogJobStatusResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogJobStatusResult() *LogJobStatusResult {
	this := LogJobStatusResult{}
	return &this
}

// NewLogJobStatusResultWithDefaults instantiates a new LogJobStatusResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogJobStatusResultWithDefaults() *LogJobStatusResult {
	this := LogJobStatusResult{}
	return &this
}

// GetLogAnalysisStatus returns the LogAnalysisStatus field value if set, zero value otherwise.
func (o *LogJobStatusResult) GetLogAnalysisStatus() string {
	if o == nil || o.LogAnalysisStatus == nil {
		var ret string
		return ret
	}
	return *o.LogAnalysisStatus
}

// GetLogAnalysisStatusOk returns a tuple with the LogAnalysisStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobStatusResult) GetLogAnalysisStatusOk() (*string, bool) {
	if o == nil || o.LogAnalysisStatus == nil {
		return nil, false
	}
	return o.LogAnalysisStatus, true
}

// HasLogAnalysisStatus returns a boolean if a field has been set.
func (o *LogJobStatusResult) HasLogAnalysisStatus() bool {
	if o != nil && o.LogAnalysisStatus != nil {
		return true
	}

	return false
}

// SetLogAnalysisStatus gets a reference to the given string and assigns it to the LogAnalysisStatus field.
func (o *LogJobStatusResult) SetLogAnalysisStatus(v string) {
	o.LogAnalysisStatus = &v
}

// GetStatusChangeTimestamp returns the StatusChangeTimestamp field value if set, zero value otherwise.
func (o *LogJobStatusResult) GetStatusChangeTimestamp() int64 {
	if o == nil || o.StatusChangeTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.StatusChangeTimestamp
}

// GetStatusChangeTimestampOk returns a tuple with the StatusChangeTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobStatusResult) GetStatusChangeTimestampOk() (*int64, bool) {
	if o == nil || o.StatusChangeTimestamp == nil {
		return nil, false
	}
	return o.StatusChangeTimestamp, true
}

// HasStatusChangeTimestamp returns a boolean if a field has been set.
func (o *LogJobStatusResult) HasStatusChangeTimestamp() bool {
	if o != nil && o.StatusChangeTimestamp != nil {
		return true
	}

	return false
}

// SetStatusChangeTimestamp gets a reference to the given int64 and assigns it to the StatusChangeTimestamp field.
func (o *LogJobStatusResult) SetStatusChangeTimestamp(v int64) {
	o.StatusChangeTimestamp = &v
}

// GetLogHandlingError returns the LogHandlingError field value if set, zero value otherwise.
func (o *LogJobStatusResult) GetLogHandlingError() string {
	if o == nil || o.LogHandlingError == nil {
		var ret string
		return ret
	}
	return *o.LogHandlingError
}

// GetLogHandlingErrorOk returns a tuple with the LogHandlingError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobStatusResult) GetLogHandlingErrorOk() (*string, bool) {
	if o == nil || o.LogHandlingError == nil {
		return nil, false
	}
	return o.LogHandlingError, true
}

// HasLogHandlingError returns a boolean if a field has been set.
func (o *LogJobStatusResult) HasLogHandlingError() bool {
	if o != nil && o.LogHandlingError != nil {
		return true
	}

	return false
}

// SetLogHandlingError gets a reference to the given string and assigns it to the LogHandlingError field.
func (o *LogJobStatusResult) SetLogHandlingError(v string) {
	o.LogHandlingError = &v
}

// GetRecordsTotal returns the RecordsTotal field value if set, zero value otherwise.
func (o *LogJobStatusResult) GetRecordsTotal() int32 {
	if o == nil || o.RecordsTotal == nil {
		var ret int32
		return ret
	}
	return *o.RecordsTotal
}

// GetRecordsTotalOk returns a tuple with the RecordsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobStatusResult) GetRecordsTotalOk() (*int32, bool) {
	if o == nil || o.RecordsTotal == nil {
		return nil, false
	}
	return o.RecordsTotal, true
}

// HasRecordsTotal returns a boolean if a field has been set.
func (o *LogJobStatusResult) HasRecordsTotal() bool {
	if o != nil && o.RecordsTotal != nil {
		return true
	}

	return false
}

// SetRecordsTotal gets a reference to the given int32 and assigns it to the RecordsTotal field.
func (o *LogJobStatusResult) SetRecordsTotal(v int32) {
	o.RecordsTotal = &v
}

// GetSortableFields returns the SortableFields field value if set, zero value otherwise.
func (o *LogJobStatusResult) GetSortableFields() []string {
	if o == nil || o.SortableFields == nil {
		var ret []string
		return ret
	}
	return *o.SortableFields
}

// GetSortableFieldsOk returns a tuple with the SortableFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobStatusResult) GetSortableFieldsOk() (*[]string, bool) {
	if o == nil || o.SortableFields == nil {
		return nil, false
	}
	return o.SortableFields, true
}

// HasSortableFields returns a boolean if a field has been set.
func (o *LogJobStatusResult) HasSortableFields() bool {
	if o != nil && o.SortableFields != nil {
		return true
	}

	return false
}

// SetSortableFields gets a reference to the given []string and assigns it to the SortableFields field.
func (o *LogJobStatusResult) SetSortableFields(v []string) {
	o.SortableFields = &v
}

// GetFilterableFields returns the FilterableFields field value if set, zero value otherwise.
func (o *LogJobStatusResult) GetFilterableFields() []string {
	if o == nil || o.FilterableFields == nil {
		var ret []string
		return ret
	}
	return *o.FilterableFields
}

// GetFilterableFieldsOk returns a tuple with the FilterableFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobStatusResult) GetFilterableFieldsOk() (*[]string, bool) {
	if o == nil || o.FilterableFields == nil {
		return nil, false
	}
	return o.FilterableFields, true
}

// HasFilterableFields returns a boolean if a field has been set.
func (o *LogJobStatusResult) HasFilterableFields() bool {
	if o != nil && o.FilterableFields != nil {
		return true
	}

	return false
}

// SetFilterableFields gets a reference to the given []string and assigns it to the FilterableFields field.
func (o *LogJobStatusResult) SetFilterableFields(v []string) {
	o.FilterableFields = &v
}

func (o LogJobStatusResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogAnalysisStatus != nil {
		toSerialize["logAnalysisStatus"] = o.LogAnalysisStatus
	}
	if o.StatusChangeTimestamp != nil {
		toSerialize["statusChangeTimestamp"] = o.StatusChangeTimestamp
	}
	if o.LogHandlingError != nil {
		toSerialize["logHandlingError"] = o.LogHandlingError
	}
	if o.RecordsTotal != nil {
		toSerialize["recordsTotal"] = o.RecordsTotal
	}
	if o.SortableFields != nil {
		toSerialize["sortableFields"] = o.SortableFields
	}
	if o.FilterableFields != nil {
		toSerialize["filterableFields"] = o.FilterableFields
	}
	return json.Marshal(toSerialize)
}

type NullableLogJobStatusResult struct {
	value *LogJobStatusResult
	isSet bool
}

func (v NullableLogJobStatusResult) Get() *LogJobStatusResult {
	return v.value
}

func (v *NullableLogJobStatusResult) Set(val *LogJobStatusResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLogJobStatusResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLogJobStatusResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogJobStatusResult(val *LogJobStatusResult) *NullableLogJobStatusResult {
	return &NullableLogJobStatusResult{value: val, isSet: true}
}

func (v NullableLogJobStatusResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogJobStatusResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


