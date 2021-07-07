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

// SecurityProblemEvent The event of a security problem.
type SecurityProblemEvent struct {
	// The timestamp when the event occurred.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// The reason of the event creation.
	Reason *string `json:"reason,omitempty"`
	RiskAssessmentSnapshot *RiskAssessmentSnapshot `json:"riskAssessmentSnapshot,omitempty"`
	MuteState *MuteState `json:"muteState,omitempty"`
}

// NewSecurityProblemEvent instantiates a new SecurityProblemEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityProblemEvent() *SecurityProblemEvent {
	this := SecurityProblemEvent{}
	return &this
}

// NewSecurityProblemEventWithDefaults instantiates a new SecurityProblemEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityProblemEventWithDefaults() *SecurityProblemEvent {
	this := SecurityProblemEvent{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SecurityProblemEvent) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblemEvent) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SecurityProblemEvent) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *SecurityProblemEvent) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SecurityProblemEvent) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblemEvent) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SecurityProblemEvent) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SecurityProblemEvent) SetReason(v string) {
	o.Reason = &v
}

// GetRiskAssessmentSnapshot returns the RiskAssessmentSnapshot field value if set, zero value otherwise.
func (o *SecurityProblemEvent) GetRiskAssessmentSnapshot() RiskAssessmentSnapshot {
	if o == nil || o.RiskAssessmentSnapshot == nil {
		var ret RiskAssessmentSnapshot
		return ret
	}
	return *o.RiskAssessmentSnapshot
}

// GetRiskAssessmentSnapshotOk returns a tuple with the RiskAssessmentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblemEvent) GetRiskAssessmentSnapshotOk() (*RiskAssessmentSnapshot, bool) {
	if o == nil || o.RiskAssessmentSnapshot == nil {
		return nil, false
	}
	return o.RiskAssessmentSnapshot, true
}

// HasRiskAssessmentSnapshot returns a boolean if a field has been set.
func (o *SecurityProblemEvent) HasRiskAssessmentSnapshot() bool {
	if o != nil && o.RiskAssessmentSnapshot != nil {
		return true
	}

	return false
}

// SetRiskAssessmentSnapshot gets a reference to the given RiskAssessmentSnapshot and assigns it to the RiskAssessmentSnapshot field.
func (o *SecurityProblemEvent) SetRiskAssessmentSnapshot(v RiskAssessmentSnapshot) {
	o.RiskAssessmentSnapshot = &v
}

// GetMuteState returns the MuteState field value if set, zero value otherwise.
func (o *SecurityProblemEvent) GetMuteState() MuteState {
	if o == nil || o.MuteState == nil {
		var ret MuteState
		return ret
	}
	return *o.MuteState
}

// GetMuteStateOk returns a tuple with the MuteState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblemEvent) GetMuteStateOk() (*MuteState, bool) {
	if o == nil || o.MuteState == nil {
		return nil, false
	}
	return o.MuteState, true
}

// HasMuteState returns a boolean if a field has been set.
func (o *SecurityProblemEvent) HasMuteState() bool {
	if o != nil && o.MuteState != nil {
		return true
	}

	return false
}

// SetMuteState gets a reference to the given MuteState and assigns it to the MuteState field.
func (o *SecurityProblemEvent) SetMuteState(v MuteState) {
	o.MuteState = &v
}

func (o SecurityProblemEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.RiskAssessmentSnapshot != nil {
		toSerialize["riskAssessmentSnapshot"] = o.RiskAssessmentSnapshot
	}
	if o.MuteState != nil {
		toSerialize["muteState"] = o.MuteState
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityProblemEvent struct {
	value *SecurityProblemEvent
	isSet bool
}

func (v NullableSecurityProblemEvent) Get() *SecurityProblemEvent {
	return v.value
}

func (v *NullableSecurityProblemEvent) Set(val *SecurityProblemEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityProblemEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityProblemEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityProblemEvent(val *SecurityProblemEvent) *NullableSecurityProblemEvent {
	return &NullableSecurityProblemEvent{value: val, isSet: true}
}

func (v NullableSecurityProblemEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityProblemEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


