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

// LogJobRecordsTopValuesRestResult The top values results of the log analysis job.
type LogJobRecordsTopValuesRestResult struct {
	// Log analysis parsing result top values
	ParsingFieldTopValues *[]ParsingFieldTopValue `json:"parsingFieldTopValues,omitempty"`
	// Log analysis parsing result top values count
	ValuesCount *int32 `json:"valuesCount,omitempty"`
}

// NewLogJobRecordsTopValuesRestResult instantiates a new LogJobRecordsTopValuesRestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogJobRecordsTopValuesRestResult() *LogJobRecordsTopValuesRestResult {
	this := LogJobRecordsTopValuesRestResult{}
	return &this
}

// NewLogJobRecordsTopValuesRestResultWithDefaults instantiates a new LogJobRecordsTopValuesRestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogJobRecordsTopValuesRestResultWithDefaults() *LogJobRecordsTopValuesRestResult {
	this := LogJobRecordsTopValuesRestResult{}
	return &this
}

// GetParsingFieldTopValues returns the ParsingFieldTopValues field value if set, zero value otherwise.
func (o *LogJobRecordsTopValuesRestResult) GetParsingFieldTopValues() []ParsingFieldTopValue {
	if o == nil || o.ParsingFieldTopValues == nil {
		var ret []ParsingFieldTopValue
		return ret
	}
	return *o.ParsingFieldTopValues
}

// GetParsingFieldTopValuesOk returns a tuple with the ParsingFieldTopValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobRecordsTopValuesRestResult) GetParsingFieldTopValuesOk() (*[]ParsingFieldTopValue, bool) {
	if o == nil || o.ParsingFieldTopValues == nil {
		return nil, false
	}
	return o.ParsingFieldTopValues, true
}

// HasParsingFieldTopValues returns a boolean if a field has been set.
func (o *LogJobRecordsTopValuesRestResult) HasParsingFieldTopValues() bool {
	if o != nil && o.ParsingFieldTopValues != nil {
		return true
	}

	return false
}

// SetParsingFieldTopValues gets a reference to the given []ParsingFieldTopValue and assigns it to the ParsingFieldTopValues field.
func (o *LogJobRecordsTopValuesRestResult) SetParsingFieldTopValues(v []ParsingFieldTopValue) {
	o.ParsingFieldTopValues = &v
}

// GetValuesCount returns the ValuesCount field value if set, zero value otherwise.
func (o *LogJobRecordsTopValuesRestResult) GetValuesCount() int32 {
	if o == nil || o.ValuesCount == nil {
		var ret int32
		return ret
	}
	return *o.ValuesCount
}

// GetValuesCountOk returns a tuple with the ValuesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogJobRecordsTopValuesRestResult) GetValuesCountOk() (*int32, bool) {
	if o == nil || o.ValuesCount == nil {
		return nil, false
	}
	return o.ValuesCount, true
}

// HasValuesCount returns a boolean if a field has been set.
func (o *LogJobRecordsTopValuesRestResult) HasValuesCount() bool {
	if o != nil && o.ValuesCount != nil {
		return true
	}

	return false
}

// SetValuesCount gets a reference to the given int32 and assigns it to the ValuesCount field.
func (o *LogJobRecordsTopValuesRestResult) SetValuesCount(v int32) {
	o.ValuesCount = &v
}

func (o LogJobRecordsTopValuesRestResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParsingFieldTopValues != nil {
		toSerialize["parsingFieldTopValues"] = o.ParsingFieldTopValues
	}
	if o.ValuesCount != nil {
		toSerialize["valuesCount"] = o.ValuesCount
	}
	return json.Marshal(toSerialize)
}

type NullableLogJobRecordsTopValuesRestResult struct {
	value *LogJobRecordsTopValuesRestResult
	isSet bool
}

func (v NullableLogJobRecordsTopValuesRestResult) Get() *LogJobRecordsTopValuesRestResult {
	return v.value
}

func (v *NullableLogJobRecordsTopValuesRestResult) Set(val *LogJobRecordsTopValuesRestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLogJobRecordsTopValuesRestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLogJobRecordsTopValuesRestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogJobRecordsTopValuesRestResult(val *LogJobRecordsTopValuesRestResult) *NullableLogJobRecordsTopValuesRestResult {
	return &NullableLogJobRecordsTopValuesRestResult{value: val, isSet: true}
}

func (v NullableLogJobRecordsTopValuesRestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogJobRecordsTopValuesRestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


