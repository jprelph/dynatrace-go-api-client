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

// SyntheticPublicLocationUpdate The update of a public Synthetic location.
type SyntheticPublicLocationUpdate struct {
	SyntheticLocationUpdate
	// The status of the location:   * `ENABLED`: The location is displayed as active in the UI. You can assign monitors to the location.  * `DISABLED`: The location is displayed as inactive in the UI. You can't assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location.  * `HIDDEN`: The location is not displayed in the UI. You can't assign monitors to the location. You can only set location as `HIDDEN` when no monitor is assigned to it.
	Status string `json:"status"`
}

// NewSyntheticPublicLocationUpdate instantiates a new SyntheticPublicLocationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticPublicLocationUpdate(status string, ) *SyntheticPublicLocationUpdate {
	this := SyntheticPublicLocationUpdate{}
	this.Status = status
	return &this
}

// NewSyntheticPublicLocationUpdateWithDefaults instantiates a new SyntheticPublicLocationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticPublicLocationUpdateWithDefaults() *SyntheticPublicLocationUpdate {
	this := SyntheticPublicLocationUpdate{}
	return &this
}

// GetStatus returns the Status field value
func (o *SyntheticPublicLocationUpdate) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SyntheticPublicLocationUpdate) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SyntheticPublicLocationUpdate) SetStatus(v string) {
	o.Status = v
}

func (o SyntheticPublicLocationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSyntheticLocationUpdate, errSyntheticLocationUpdate := json.Marshal(o.SyntheticLocationUpdate)
	if errSyntheticLocationUpdate != nil {
		return []byte{}, errSyntheticLocationUpdate
	}
	errSyntheticLocationUpdate = json.Unmarshal([]byte(serializedSyntheticLocationUpdate), &toSerialize)
	if errSyntheticLocationUpdate != nil {
		return []byte{}, errSyntheticLocationUpdate
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticPublicLocationUpdate struct {
	value *SyntheticPublicLocationUpdate
	isSet bool
}

func (v NullableSyntheticPublicLocationUpdate) Get() *SyntheticPublicLocationUpdate {
	return v.value
}

func (v *NullableSyntheticPublicLocationUpdate) Set(val *SyntheticPublicLocationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticPublicLocationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticPublicLocationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticPublicLocationUpdate(val *SyntheticPublicLocationUpdate) *NullableSyntheticPublicLocationUpdate {
	return &NullableSyntheticPublicLocationUpdate{value: val, isSet: true}
}

func (v NullableSyntheticPublicLocationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticPublicLocationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


