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

// SyntheticMonitorUpdate The synthetic monitor update. This is a base object, the exact type depends on the value of the `type` field.
type SyntheticMonitorUpdate struct {
	// The frequency of the monitor, in minutes.    You can use one of the following values: `5`, `10`, `15`, `30`, and `60`.
	FrequencyMin int32 `json:"frequencyMin"`
	AnomalyDetection *AnomalyDetection `json:"anomalyDetection,omitempty"`
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `BROWSER` -> BrowserSyntheticMonitorUpdate  * `HTTP` -> HttpSyntheticMonitorUpdate  
	Type string `json:"type"`
	// The name of the monitor.
	Name string `json:"name"`
	// A list of locations from which the monitor is executed.    To specify a location, use its entity ID.
	Locations []string `json:"locations"`
	// The monitor is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// The script of a [browser](https://dt-url.net/9c103rda) or HTTP monitor.
	Script map[string]interface{} `json:"script"`
	// A set of tags assigned to the monitor.    You can specify only the value of the tag here and the `CONTEXTLESS` context and source 'USER' will be added automatically. But preferred option is usage of TagWithSourceDto model.
	Tags []TagWithSourceInfo `json:"tags"`
	// A set of manually assigned applications.
	ManuallyAssignedApps []string `json:"manuallyAssignedApps"`
}

// NewSyntheticMonitorUpdate instantiates a new SyntheticMonitorUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticMonitorUpdate(frequencyMin int32, type_ string, name string, locations []string, enabled bool, script map[string]interface{}, tags []TagWithSourceInfo, manuallyAssignedApps []string, ) *SyntheticMonitorUpdate {
	this := SyntheticMonitorUpdate{}
	this.FrequencyMin = frequencyMin
	this.Type = type_
	this.Name = name
	this.Locations = locations
	this.Enabled = enabled
	this.Script = script
	this.Tags = tags
	this.ManuallyAssignedApps = manuallyAssignedApps
	return &this
}

// NewSyntheticMonitorUpdateWithDefaults instantiates a new SyntheticMonitorUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticMonitorUpdateWithDefaults() *SyntheticMonitorUpdate {
	this := SyntheticMonitorUpdate{}
	return &this
}

// GetFrequencyMin returns the FrequencyMin field value
func (o *SyntheticMonitorUpdate) GetFrequencyMin() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FrequencyMin
}

// GetFrequencyMinOk returns a tuple with the FrequencyMin field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetFrequencyMinOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FrequencyMin, true
}

// SetFrequencyMin sets field value
func (o *SyntheticMonitorUpdate) SetFrequencyMin(v int32) {
	o.FrequencyMin = v
}

// GetAnomalyDetection returns the AnomalyDetection field value if set, zero value otherwise.
func (o *SyntheticMonitorUpdate) GetAnomalyDetection() AnomalyDetection {
	if o == nil || o.AnomalyDetection == nil {
		var ret AnomalyDetection
		return ret
	}
	return *o.AnomalyDetection
}

// GetAnomalyDetectionOk returns a tuple with the AnomalyDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetAnomalyDetectionOk() (*AnomalyDetection, bool) {
	if o == nil || o.AnomalyDetection == nil {
		return nil, false
	}
	return o.AnomalyDetection, true
}

// HasAnomalyDetection returns a boolean if a field has been set.
func (o *SyntheticMonitorUpdate) HasAnomalyDetection() bool {
	if o != nil && o.AnomalyDetection != nil {
		return true
	}

	return false
}

// SetAnomalyDetection gets a reference to the given AnomalyDetection and assigns it to the AnomalyDetection field.
func (o *SyntheticMonitorUpdate) SetAnomalyDetection(v AnomalyDetection) {
	o.AnomalyDetection = &v
}

// GetType returns the Type field value
func (o *SyntheticMonitorUpdate) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticMonitorUpdate) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *SyntheticMonitorUpdate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SyntheticMonitorUpdate) SetName(v string) {
	o.Name = v
}

// GetLocations returns the Locations field value
func (o *SyntheticMonitorUpdate) GetLocations() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetLocationsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locations, true
}

// SetLocations sets field value
func (o *SyntheticMonitorUpdate) SetLocations(v []string) {
	o.Locations = v
}

// GetEnabled returns the Enabled field value
func (o *SyntheticMonitorUpdate) GetEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SyntheticMonitorUpdate) SetEnabled(v bool) {
	o.Enabled = v
}

// GetScript returns the Script field value
func (o *SyntheticMonitorUpdate) GetScript() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetScriptOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *SyntheticMonitorUpdate) SetScript(v map[string]interface{}) {
	o.Script = v
}

// GetTags returns the Tags field value
func (o *SyntheticMonitorUpdate) GetTags() []TagWithSourceInfo {
	if o == nil  {
		var ret []TagWithSourceInfo
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetTagsOk() (*[]TagWithSourceInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *SyntheticMonitorUpdate) SetTags(v []TagWithSourceInfo) {
	o.Tags = v
}

// GetManuallyAssignedApps returns the ManuallyAssignedApps field value
func (o *SyntheticMonitorUpdate) GetManuallyAssignedApps() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.ManuallyAssignedApps
}

// GetManuallyAssignedAppsOk returns a tuple with the ManuallyAssignedApps field value
// and a boolean to check if the value has been set.
func (o *SyntheticMonitorUpdate) GetManuallyAssignedAppsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ManuallyAssignedApps, true
}

// SetManuallyAssignedApps sets field value
func (o *SyntheticMonitorUpdate) SetManuallyAssignedApps(v []string) {
	o.ManuallyAssignedApps = v
}

func (o SyntheticMonitorUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["frequencyMin"] = o.FrequencyMin
	}
	if o.AnomalyDetection != nil {
		toSerialize["anomalyDetection"] = o.AnomalyDetection
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["locations"] = o.Locations
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["script"] = o.Script
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["manuallyAssignedApps"] = o.ManuallyAssignedApps
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticMonitorUpdate struct {
	value *SyntheticMonitorUpdate
	isSet bool
}

func (v NullableSyntheticMonitorUpdate) Get() *SyntheticMonitorUpdate {
	return v.value
}

func (v *NullableSyntheticMonitorUpdate) Set(val *SyntheticMonitorUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticMonitorUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticMonitorUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticMonitorUpdate(val *SyntheticMonitorUpdate) *NullableSyntheticMonitorUpdate {
	return &NullableSyntheticMonitorUpdate{value: val, isSet: true}
}

func (v NullableSyntheticMonitorUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticMonitorUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


