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

// EventEvidenceAllOf struct for EventEvidenceAllOf
type EventEvidenceAllOf struct {
	// The type of the event.
	EventType *string `json:"eventType,omitempty"`
	// The ID of the event.
	EventId *string `json:"eventId,omitempty"`
}

// NewEventEvidenceAllOf instantiates a new EventEvidenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventEvidenceAllOf() *EventEvidenceAllOf {
	this := EventEvidenceAllOf{}
	return &this
}

// NewEventEvidenceAllOfWithDefaults instantiates a new EventEvidenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventEvidenceAllOfWithDefaults() *EventEvidenceAllOf {
	this := EventEvidenceAllOf{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventEvidenceAllOf) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEvidenceAllOf) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventEvidenceAllOf) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *EventEvidenceAllOf) SetEventType(v string) {
	o.EventType = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EventEvidenceAllOf) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEvidenceAllOf) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EventEvidenceAllOf) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *EventEvidenceAllOf) SetEventId(v string) {
	o.EventId = &v
}

func (o EventEvidenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	return json.Marshal(toSerialize)
}

type NullableEventEvidenceAllOf struct {
	value *EventEvidenceAllOf
	isSet bool
}

func (v NullableEventEvidenceAllOf) Get() *EventEvidenceAllOf {
	return v.value
}

func (v *NullableEventEvidenceAllOf) Set(val *EventEvidenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventEvidenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventEvidenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventEvidenceAllOf(val *EventEvidenceAllOf) *NullableEventEvidenceAllOf {
	return &NullableEventEvidenceAllOf{value: val, isSet: true}
}

func (v NullableEventEvidenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventEvidenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


