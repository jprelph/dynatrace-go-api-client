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

// ClusterVersion struct for ClusterVersion
type ClusterVersion struct {
	// The version of the Dynatrace server.
	Version *string `json:"version,omitempty"`
}

// NewClusterVersion instantiates a new ClusterVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterVersion() *ClusterVersion {
	this := ClusterVersion{}
	return &this
}

// NewClusterVersionWithDefaults instantiates a new ClusterVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterVersionWithDefaults() *ClusterVersion {
	this := ClusterVersion{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterVersion) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterVersion) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterVersion) SetVersion(v string) {
	o.Version = &v
}

func (o ClusterVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableClusterVersion struct {
	value *ClusterVersion
	isSet bool
}

func (v NullableClusterVersion) Get() *ClusterVersion {
	return v.value
}

func (v *NullableClusterVersion) Set(val *ClusterVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterVersion(val *ClusterVersion) *NullableClusterVersion {
	return &NullableClusterVersion{value: val, isSet: true}
}

func (v NullableClusterVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


