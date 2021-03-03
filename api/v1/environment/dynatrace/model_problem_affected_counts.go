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

// ProblemAffectedCounts The number of affected entities per impact level.
type ProblemAffectedCounts struct {
	// The count of impacted entities for the `APPLICATION` impact level.
	APPLICATION *int32 `json:"APPLICATION,omitempty"`
	// The count of impacted entities for the `ENVIRONMENT` impact level.
	ENVIRONMENT *int32 `json:"ENVIRONMENT,omitempty"`
	// The count of impacted entities for the `INFRASTRUCTURE` impact level.
	INFRASTRUCTURE *int32 `json:"INFRASTRUCTURE,omitempty"`
	// The count of impacted entities for the `SERVICE` impact level.
	SERVICE *int32 `json:"SERVICE,omitempty"`
}

// NewProblemAffectedCounts instantiates a new ProblemAffectedCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemAffectedCounts() *ProblemAffectedCounts {
	this := ProblemAffectedCounts{}
	return &this
}

// NewProblemAffectedCountsWithDefaults instantiates a new ProblemAffectedCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemAffectedCountsWithDefaults() *ProblemAffectedCounts {
	this := ProblemAffectedCounts{}
	return &this
}

// GetAPPLICATION returns the APPLICATION field value if set, zero value otherwise.
func (o *ProblemAffectedCounts) GetAPPLICATION() int32 {
	if o == nil || o.APPLICATION == nil {
		var ret int32
		return ret
	}
	return *o.APPLICATION
}

// GetAPPLICATIONOk returns a tuple with the APPLICATION field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemAffectedCounts) GetAPPLICATIONOk() (*int32, bool) {
	if o == nil || o.APPLICATION == nil {
		return nil, false
	}
	return o.APPLICATION, true
}

// HasAPPLICATION returns a boolean if a field has been set.
func (o *ProblemAffectedCounts) HasAPPLICATION() bool {
	if o != nil && o.APPLICATION != nil {
		return true
	}

	return false
}

// SetAPPLICATION gets a reference to the given int32 and assigns it to the APPLICATION field.
func (o *ProblemAffectedCounts) SetAPPLICATION(v int32) {
	o.APPLICATION = &v
}

// GetENVIRONMENT returns the ENVIRONMENT field value if set, zero value otherwise.
func (o *ProblemAffectedCounts) GetENVIRONMENT() int32 {
	if o == nil || o.ENVIRONMENT == nil {
		var ret int32
		return ret
	}
	return *o.ENVIRONMENT
}

// GetENVIRONMENTOk returns a tuple with the ENVIRONMENT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemAffectedCounts) GetENVIRONMENTOk() (*int32, bool) {
	if o == nil || o.ENVIRONMENT == nil {
		return nil, false
	}
	return o.ENVIRONMENT, true
}

// HasENVIRONMENT returns a boolean if a field has been set.
func (o *ProblemAffectedCounts) HasENVIRONMENT() bool {
	if o != nil && o.ENVIRONMENT != nil {
		return true
	}

	return false
}

// SetENVIRONMENT gets a reference to the given int32 and assigns it to the ENVIRONMENT field.
func (o *ProblemAffectedCounts) SetENVIRONMENT(v int32) {
	o.ENVIRONMENT = &v
}

// GetINFRASTRUCTURE returns the INFRASTRUCTURE field value if set, zero value otherwise.
func (o *ProblemAffectedCounts) GetINFRASTRUCTURE() int32 {
	if o == nil || o.INFRASTRUCTURE == nil {
		var ret int32
		return ret
	}
	return *o.INFRASTRUCTURE
}

// GetINFRASTRUCTUREOk returns a tuple with the INFRASTRUCTURE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemAffectedCounts) GetINFRASTRUCTUREOk() (*int32, bool) {
	if o == nil || o.INFRASTRUCTURE == nil {
		return nil, false
	}
	return o.INFRASTRUCTURE, true
}

// HasINFRASTRUCTURE returns a boolean if a field has been set.
func (o *ProblemAffectedCounts) HasINFRASTRUCTURE() bool {
	if o != nil && o.INFRASTRUCTURE != nil {
		return true
	}

	return false
}

// SetINFRASTRUCTURE gets a reference to the given int32 and assigns it to the INFRASTRUCTURE field.
func (o *ProblemAffectedCounts) SetINFRASTRUCTURE(v int32) {
	o.INFRASTRUCTURE = &v
}

// GetSERVICE returns the SERVICE field value if set, zero value otherwise.
func (o *ProblemAffectedCounts) GetSERVICE() int32 {
	if o == nil || o.SERVICE == nil {
		var ret int32
		return ret
	}
	return *o.SERVICE
}

// GetSERVICEOk returns a tuple with the SERVICE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemAffectedCounts) GetSERVICEOk() (*int32, bool) {
	if o == nil || o.SERVICE == nil {
		return nil, false
	}
	return o.SERVICE, true
}

// HasSERVICE returns a boolean if a field has been set.
func (o *ProblemAffectedCounts) HasSERVICE() bool {
	if o != nil && o.SERVICE != nil {
		return true
	}

	return false
}

// SetSERVICE gets a reference to the given int32 and assigns it to the SERVICE field.
func (o *ProblemAffectedCounts) SetSERVICE(v int32) {
	o.SERVICE = &v
}

func (o ProblemAffectedCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.APPLICATION != nil {
		toSerialize["APPLICATION"] = o.APPLICATION
	}
	if o.ENVIRONMENT != nil {
		toSerialize["ENVIRONMENT"] = o.ENVIRONMENT
	}
	if o.INFRASTRUCTURE != nil {
		toSerialize["INFRASTRUCTURE"] = o.INFRASTRUCTURE
	}
	if o.SERVICE != nil {
		toSerialize["SERVICE"] = o.SERVICE
	}
	return json.Marshal(toSerialize)
}

type NullableProblemAffectedCounts struct {
	value *ProblemAffectedCounts
	isSet bool
}

func (v NullableProblemAffectedCounts) Get() *ProblemAffectedCounts {
	return v.value
}

func (v *NullableProblemAffectedCounts) Set(val *ProblemAffectedCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemAffectedCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemAffectedCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemAffectedCounts(val *ProblemAffectedCounts) *NullableProblemAffectedCounts {
	return &NullableProblemAffectedCounts{value: val, isSet: true}
}

func (v NullableProblemAffectedCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemAffectedCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


