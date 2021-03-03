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

// SyntheticPublicLocationUpdateAllOf struct for SyntheticPublicLocationUpdateAllOf
type SyntheticPublicLocationUpdateAllOf struct {
	// The status of the location:   * `ENABLED`: The location is displayed as active in the UI. You can assign monitors to the location.  * `DISABLED`: The location is displayed as inactive in the UI. You can't assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location.  * `HIDDEN`: The location is not displayed in the UI. You can't assign monitors to the location. You can only set location as `HIDDEN` when no monitor is assigned to it.
	Status *string `json:"status,omitempty"`
}

// NewSyntheticPublicLocationUpdateAllOf instantiates a new SyntheticPublicLocationUpdateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticPublicLocationUpdateAllOf() *SyntheticPublicLocationUpdateAllOf {
	this := SyntheticPublicLocationUpdateAllOf{}
	return &this
}

// NewSyntheticPublicLocationUpdateAllOfWithDefaults instantiates a new SyntheticPublicLocationUpdateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticPublicLocationUpdateAllOfWithDefaults() *SyntheticPublicLocationUpdateAllOf {
	this := SyntheticPublicLocationUpdateAllOf{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticPublicLocationUpdateAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPublicLocationUpdateAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticPublicLocationUpdateAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticPublicLocationUpdateAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o SyntheticPublicLocationUpdateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticPublicLocationUpdateAllOf struct {
	value *SyntheticPublicLocationUpdateAllOf
	isSet bool
}

func (v NullableSyntheticPublicLocationUpdateAllOf) Get() *SyntheticPublicLocationUpdateAllOf {
	return v.value
}

func (v *NullableSyntheticPublicLocationUpdateAllOf) Set(val *SyntheticPublicLocationUpdateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticPublicLocationUpdateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticPublicLocationUpdateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticPublicLocationUpdateAllOf(val *SyntheticPublicLocationUpdateAllOf) *NullableSyntheticPublicLocationUpdateAllOf {
	return &NullableSyntheticPublicLocationUpdateAllOf{value: val, isSet: true}
}

func (v NullableSyntheticPublicLocationUpdateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticPublicLocationUpdateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


