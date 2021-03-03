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

// ApplicationToRelationships The list of incoming calls to the application.
type ApplicationToRelationships struct {
	// Represents relationship between a synthetic monitor and an application
	Monitors *[]string `json:"monitors,omitempty"`
}

// NewApplicationToRelationships instantiates a new ApplicationToRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationToRelationships() *ApplicationToRelationships {
	this := ApplicationToRelationships{}
	return &this
}

// NewApplicationToRelationshipsWithDefaults instantiates a new ApplicationToRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationToRelationshipsWithDefaults() *ApplicationToRelationships {
	this := ApplicationToRelationships{}
	return &this
}

// GetMonitors returns the Monitors field value if set, zero value otherwise.
func (o *ApplicationToRelationships) GetMonitors() []string {
	if o == nil || o.Monitors == nil {
		var ret []string
		return ret
	}
	return *o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationToRelationships) GetMonitorsOk() (*[]string, bool) {
	if o == nil || o.Monitors == nil {
		return nil, false
	}
	return o.Monitors, true
}

// HasMonitors returns a boolean if a field has been set.
func (o *ApplicationToRelationships) HasMonitors() bool {
	if o != nil && o.Monitors != nil {
		return true
	}

	return false
}

// SetMonitors gets a reference to the given []string and assigns it to the Monitors field.
func (o *ApplicationToRelationships) SetMonitors(v []string) {
	o.Monitors = &v
}

func (o ApplicationToRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Monitors != nil {
		toSerialize["monitors"] = o.Monitors
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationToRelationships struct {
	value *ApplicationToRelationships
	isSet bool
}

func (v NullableApplicationToRelationships) Get() *ApplicationToRelationships {
	return v.value
}

func (v *NullableApplicationToRelationships) Set(val *ApplicationToRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationToRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationToRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationToRelationships(val *ApplicationToRelationships) *NullableApplicationToRelationships {
	return &NullableApplicationToRelationships{value: val, isSet: true}
}

func (v NullableApplicationToRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationToRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


