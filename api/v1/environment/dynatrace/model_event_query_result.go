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

// EventQueryResult Events of the environment.
type EventQueryResult struct {
	// The cursor for the next 150 events, fitting the specified criteria.    Set this value for the **cursor** query parameter. Without it you'll get the first 150 events again.   You don't have to specify any additional parameters, because the cursor already contains all of them.
	NextCursor *string `json:"nextCursor,omitempty"`
	// Start of the query timeframe.
	From *int64 `json:"from,omitempty"`
	// End of the query timeframe.
	To *int64 `json:"to,omitempty"`
	// The total amount of events, fitting the specified criteria.
	TotalEventCount *int64 `json:"totalEventCount,omitempty"`
	// The list of events.
	Events *[]EventRestEntry `json:"events,omitempty"`
}

// NewEventQueryResult instantiates a new EventQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventQueryResult() *EventQueryResult {
	this := EventQueryResult{}
	return &this
}

// NewEventQueryResultWithDefaults instantiates a new EventQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventQueryResultWithDefaults() *EventQueryResult {
	this := EventQueryResult{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *EventQueryResult) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventQueryResult) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *EventQueryResult) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *EventQueryResult) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EventQueryResult) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventQueryResult) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EventQueryResult) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *EventQueryResult) SetFrom(v int64) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EventQueryResult) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventQueryResult) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EventQueryResult) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *EventQueryResult) SetTo(v int64) {
	o.To = &v
}

// GetTotalEventCount returns the TotalEventCount field value if set, zero value otherwise.
func (o *EventQueryResult) GetTotalEventCount() int64 {
	if o == nil || o.TotalEventCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalEventCount
}

// GetTotalEventCountOk returns a tuple with the TotalEventCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventQueryResult) GetTotalEventCountOk() (*int64, bool) {
	if o == nil || o.TotalEventCount == nil {
		return nil, false
	}
	return o.TotalEventCount, true
}

// HasTotalEventCount returns a boolean if a field has been set.
func (o *EventQueryResult) HasTotalEventCount() bool {
	if o != nil && o.TotalEventCount != nil {
		return true
	}

	return false
}

// SetTotalEventCount gets a reference to the given int64 and assigns it to the TotalEventCount field.
func (o *EventQueryResult) SetTotalEventCount(v int64) {
	o.TotalEventCount = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *EventQueryResult) GetEvents() []EventRestEntry {
	if o == nil || o.Events == nil {
		var ret []EventRestEntry
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventQueryResult) GetEventsOk() (*[]EventRestEntry, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *EventQueryResult) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []EventRestEntry and assigns it to the Events field.
func (o *EventQueryResult) SetEvents(v []EventRestEntry) {
	o.Events = &v
}

func (o EventQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextCursor != nil {
		toSerialize["nextCursor"] = o.NextCursor
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.TotalEventCount != nil {
		toSerialize["totalEventCount"] = o.TotalEventCount
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableEventQueryResult struct {
	value *EventQueryResult
	isSet bool
}

func (v NullableEventQueryResult) Get() *EventQueryResult {
	return v.value
}

func (v *NullableEventQueryResult) Set(val *EventQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEventQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEventQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventQueryResult(val *EventQueryResult) *NullableEventQueryResult {
	return &NullableEventQueryResult{value: val, isSet: true}
}

func (v NullableEventQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


