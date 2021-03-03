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

// TimeseriesQueryResultWrapper struct for TimeseriesQueryResultWrapper
type TimeseriesQueryResultWrapper struct {
	Result *TimeseriesDataPointQueryResult `json:"result,omitempty"`
}

// NewTimeseriesQueryResultWrapper instantiates a new TimeseriesQueryResultWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesQueryResultWrapper() *TimeseriesQueryResultWrapper {
	this := TimeseriesQueryResultWrapper{}
	return &this
}

// NewTimeseriesQueryResultWrapperWithDefaults instantiates a new TimeseriesQueryResultWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesQueryResultWrapperWithDefaults() *TimeseriesQueryResultWrapper {
	this := TimeseriesQueryResultWrapper{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TimeseriesQueryResultWrapper) GetResult() TimeseriesDataPointQueryResult {
	if o == nil || o.Result == nil {
		var ret TimeseriesDataPointQueryResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResultWrapper) GetResultOk() (*TimeseriesDataPointQueryResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TimeseriesQueryResultWrapper) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given TimeseriesDataPointQueryResult and assigns it to the Result field.
func (o *TimeseriesQueryResultWrapper) SetResult(v TimeseriesDataPointQueryResult) {
	o.Result = &v
}

func (o TimeseriesQueryResultWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableTimeseriesQueryResultWrapper struct {
	value *TimeseriesQueryResultWrapper
	isSet bool
}

func (v NullableTimeseriesQueryResultWrapper) Get() *TimeseriesQueryResultWrapper {
	return v.value
}

func (v *NullableTimeseriesQueryResultWrapper) Set(val *TimeseriesQueryResultWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesQueryResultWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesQueryResultWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesQueryResultWrapper(val *TimeseriesQueryResultWrapper) *NullableTimeseriesQueryResultWrapper {
	return &NullableTimeseriesQueryResultWrapper{value: val, isSet: true}
}

func (v NullableTimeseriesQueryResultWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesQueryResultWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


