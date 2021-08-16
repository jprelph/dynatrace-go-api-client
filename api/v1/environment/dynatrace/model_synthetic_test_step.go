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

// SyntheticTestStep The step of a synthetic monitor.
type SyntheticTestStep struct {
	// The ID of the step.
	Id int64 `json:"id"`
	// The name of the step, displayed in the UI.
	Title string `json:"title"`
}

// NewSyntheticTestStep instantiates a new SyntheticTestStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticTestStep(id int64, title string) *SyntheticTestStep {
	this := SyntheticTestStep{}
	this.Id = id
	this.Title = title
	return &this
}

// NewSyntheticTestStepWithDefaults instantiates a new SyntheticTestStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticTestStepWithDefaults() *SyntheticTestStep {
	this := SyntheticTestStep{}
	return &this
}

// GetId returns the Id field value
func (o *SyntheticTestStep) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SyntheticTestStep) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SyntheticTestStep) SetId(v int64) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *SyntheticTestStep) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SyntheticTestStep) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SyntheticTestStep) SetTitle(v string) {
	o.Title = v
}

func (o SyntheticTestStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticTestStep struct {
	value *SyntheticTestStep
	isSet bool
}

func (v NullableSyntheticTestStep) Get() *SyntheticTestStep {
	return v.value
}

func (v *NullableSyntheticTestStep) Set(val *SyntheticTestStep) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticTestStep) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticTestStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticTestStep(val *SyntheticTestStep) *NullableSyntheticTestStep {
	return &NullableSyntheticTestStep{value: val, isSet: true}
}

func (v NullableSyntheticTestStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticTestStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


