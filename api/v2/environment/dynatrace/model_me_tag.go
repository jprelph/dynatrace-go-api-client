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

// METag The tag of a monitored entity.
type METag struct {
	// The string representation of the tag.
	StringRepresentation *string `json:"stringRepresentation,omitempty"`
	// The value of the tag.
	Value *string `json:"value,omitempty"`
	// The key of the tag.
	Key *string `json:"key,omitempty"`
	// The origin of the tag, such as AWS or Cloud Foundry.    Custom tags use the `CONTEXTLESS` value.
	Context *string `json:"context,omitempty"`
}

// NewMETag instantiates a new METag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMETag() *METag {
	this := METag{}
	return &this
}

// NewMETagWithDefaults instantiates a new METag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMETagWithDefaults() *METag {
	this := METag{}
	return &this
}

// GetStringRepresentation returns the StringRepresentation field value if set, zero value otherwise.
func (o *METag) GetStringRepresentation() string {
	if o == nil || o.StringRepresentation == nil {
		var ret string
		return ret
	}
	return *o.StringRepresentation
}

// GetStringRepresentationOk returns a tuple with the StringRepresentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *METag) GetStringRepresentationOk() (*string, bool) {
	if o == nil || o.StringRepresentation == nil {
		return nil, false
	}
	return o.StringRepresentation, true
}

// HasStringRepresentation returns a boolean if a field has been set.
func (o *METag) HasStringRepresentation() bool {
	if o != nil && o.StringRepresentation != nil {
		return true
	}

	return false
}

// SetStringRepresentation gets a reference to the given string and assigns it to the StringRepresentation field.
func (o *METag) SetStringRepresentation(v string) {
	o.StringRepresentation = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *METag) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *METag) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *METag) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *METag) SetValue(v string) {
	o.Value = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *METag) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *METag) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *METag) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *METag) SetKey(v string) {
	o.Key = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *METag) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *METag) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *METag) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *METag) SetContext(v string) {
	o.Context = &v
}

func (o METag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StringRepresentation != nil {
		toSerialize["stringRepresentation"] = o.StringRepresentation
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableMETag struct {
	value *METag
	isSet bool
}

func (v NullableMETag) Get() *METag {
	return v.value
}

func (v *NullableMETag) Set(val *METag) {
	v.value = val
	v.isSet = true
}

func (v NullableMETag) IsSet() bool {
	return v.isSet
}

func (v *NullableMETag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMETag(val *METag) *NullableMETag {
	return &NullableMETag{value: val, isSet: true}
}

func (v NullableMETag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMETag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


