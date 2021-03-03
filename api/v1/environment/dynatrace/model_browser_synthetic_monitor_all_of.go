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

// BrowserSyntheticMonitorAllOf struct for BrowserSyntheticMonitorAllOf
type BrowserSyntheticMonitorAllOf struct {
	KeyPerformanceMetrics *KeyPerformanceMetrics `json:"keyPerformanceMetrics,omitempty"`
	// A list of events for this monitor
	Events *[]EventDto `json:"events,omitempty"`
}

// NewBrowserSyntheticMonitorAllOf instantiates a new BrowserSyntheticMonitorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserSyntheticMonitorAllOf() *BrowserSyntheticMonitorAllOf {
	this := BrowserSyntheticMonitorAllOf{}
	return &this
}

// NewBrowserSyntheticMonitorAllOfWithDefaults instantiates a new BrowserSyntheticMonitorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserSyntheticMonitorAllOfWithDefaults() *BrowserSyntheticMonitorAllOf {
	this := BrowserSyntheticMonitorAllOf{}
	return &this
}

// GetKeyPerformanceMetrics returns the KeyPerformanceMetrics field value if set, zero value otherwise.
func (o *BrowserSyntheticMonitorAllOf) GetKeyPerformanceMetrics() KeyPerformanceMetrics {
	if o == nil || o.KeyPerformanceMetrics == nil {
		var ret KeyPerformanceMetrics
		return ret
	}
	return *o.KeyPerformanceMetrics
}

// GetKeyPerformanceMetricsOk returns a tuple with the KeyPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserSyntheticMonitorAllOf) GetKeyPerformanceMetricsOk() (*KeyPerformanceMetrics, bool) {
	if o == nil || o.KeyPerformanceMetrics == nil {
		return nil, false
	}
	return o.KeyPerformanceMetrics, true
}

// HasKeyPerformanceMetrics returns a boolean if a field has been set.
func (o *BrowserSyntheticMonitorAllOf) HasKeyPerformanceMetrics() bool {
	if o != nil && o.KeyPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetKeyPerformanceMetrics gets a reference to the given KeyPerformanceMetrics and assigns it to the KeyPerformanceMetrics field.
func (o *BrowserSyntheticMonitorAllOf) SetKeyPerformanceMetrics(v KeyPerformanceMetrics) {
	o.KeyPerformanceMetrics = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *BrowserSyntheticMonitorAllOf) GetEvents() []EventDto {
	if o == nil || o.Events == nil {
		var ret []EventDto
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserSyntheticMonitorAllOf) GetEventsOk() (*[]EventDto, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *BrowserSyntheticMonitorAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []EventDto and assigns it to the Events field.
func (o *BrowserSyntheticMonitorAllOf) SetEvents(v []EventDto) {
	o.Events = &v
}

func (o BrowserSyntheticMonitorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyPerformanceMetrics != nil {
		toSerialize["keyPerformanceMetrics"] = o.KeyPerformanceMetrics
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableBrowserSyntheticMonitorAllOf struct {
	value *BrowserSyntheticMonitorAllOf
	isSet bool
}

func (v NullableBrowserSyntheticMonitorAllOf) Get() *BrowserSyntheticMonitorAllOf {
	return v.value
}

func (v *NullableBrowserSyntheticMonitorAllOf) Set(val *BrowserSyntheticMonitorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserSyntheticMonitorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserSyntheticMonitorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserSyntheticMonitorAllOf(val *BrowserSyntheticMonitorAllOf) *NullableBrowserSyntheticMonitorAllOf {
	return &NullableBrowserSyntheticMonitorAllOf{value: val, isSet: true}
}

func (v NullableBrowserSyntheticMonitorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserSyntheticMonitorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


