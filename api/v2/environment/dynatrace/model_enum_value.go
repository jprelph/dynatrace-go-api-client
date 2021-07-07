/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// EnumValue An allowed value for an enum property.
type EnumValue struct {
	// The name of the value in an existing Java enum class.
	EnumInstance *string `json:"enumInstance,omitempty"`
	// The icon of the value.
	Icon *string `json:"icon,omitempty"`
	// The allowed value of the enum.
	Value *map[string]interface{} `json:"value,omitempty"`
	// The display name of the value.
	DisplayName *string `json:"displayName,omitempty"`
	// A short description of the value.
	Description *string `json:"description,omitempty"`
}

// NewEnumValue instantiates a new EnumValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumValue() *EnumValue {
	this := EnumValue{}
	return &this
}

// NewEnumValueWithDefaults instantiates a new EnumValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumValueWithDefaults() *EnumValue {
	this := EnumValue{}
	return &this
}

// GetEnumInstance returns the EnumInstance field value if set, zero value otherwise.
func (o *EnumValue) GetEnumInstance() string {
	if o == nil || o.EnumInstance == nil {
		var ret string
		return ret
	}
	return *o.EnumInstance
}

// GetEnumInstanceOk returns a tuple with the EnumInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumValue) GetEnumInstanceOk() (*string, bool) {
	if o == nil || o.EnumInstance == nil {
		return nil, false
	}
	return o.EnumInstance, true
}

// HasEnumInstance returns a boolean if a field has been set.
func (o *EnumValue) HasEnumInstance() bool {
	if o != nil && o.EnumInstance != nil {
		return true
	}

	return false
}

// SetEnumInstance gets a reference to the given string and assigns it to the EnumInstance field.
func (o *EnumValue) SetEnumInstance(v string) {
	o.EnumInstance = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *EnumValue) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumValue) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *EnumValue) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *EnumValue) SetIcon(v string) {
	o.Icon = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EnumValue) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumValue) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EnumValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *EnumValue) SetValue(v map[string]interface{}) {
	o.Value = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EnumValue) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumValue) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EnumValue) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EnumValue) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnumValue) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumValue) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnumValue) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnumValue) SetDescription(v string) {
	o.Description = &v
}

func (o EnumValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnumInstance != nil {
		toSerialize["enumInstance"] = o.EnumInstance
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableEnumValue struct {
	value *EnumValue
	isSet bool
}

func (v NullableEnumValue) Get() *EnumValue {
	return v.value
}

func (v *NullableEnumValue) Set(val *EnumValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumValue(val *EnumValue) *NullableEnumValue {
	return &NullableEnumValue{value: val, isSet: true}
}

func (v NullableEnumValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


