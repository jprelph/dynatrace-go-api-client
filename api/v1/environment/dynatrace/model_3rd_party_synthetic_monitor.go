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

// Model3rdPartySyntheticMonitor The third-party synthetic monitor.
type Model3rdPartySyntheticMonitor struct {
	// The ID of the monitor.
	Id string `json:"id"`
	// The name of the monitor.
	Title string `json:"title"`
	// A description of the monitor.
	Description *string `json:"description,omitempty"`
	// The information on monitor setup, for example `browser`.
	TestSetup *string `json:"testSetup,omitempty"`
	// The timestamp of the monitor expiration, in UTC milliseconds.
	ExpirationTimestamp *int64 `json:"expirationTimestamp,omitempty"`
	// The URL to the results of monitor execution.
	DrilldownLink *string `json:"drilldownLink,omitempty"`
	// The URL to edit the monitor in the original UI.
	EditLink *string `json:"editLink,omitempty"`
	// The monitor is enabled (`true`) or disabled (`false`). Default is `true`.   If `true`, set the **deleted** parameter to `false`.
	Enabled *bool `json:"enabled,omitempty"`
	// The flag of the deleted monitor. Default is `false`.    If `true`, set the **enabled** parameter to `false`.
	Deleted *bool `json:"deleted,omitempty"`
	// Locations from which the synthetic monitor runs.
	Locations []SyntheticTestLocation `json:"locations"`
	// Steps of the third-party monitor.
	Steps *[]SyntheticTestStep `json:"steps,omitempty"`
	// The frequency of the monitor, in seconds. The monitor is repeated with the specified interval at the third-party source.   Dynatrace expects results of a monitor execution with the specified interval. If you report results to Dynatrace less often, adjust the **noDataTimeout** value accordingly.
	ScheduleIntervalInSeconds int32 `json:"scheduleIntervalInSeconds"`
	// The timeout of the monitor, in seconds. If no result is reported within this time, the availability state switches to unmonitored. Default is doubled frequency of the monitor.
	NoDataTimeout *int32 `json:"noDataTimeout,omitempty"`
}

// NewModel3rdPartySyntheticMonitor instantiates a new Model3rdPartySyntheticMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel3rdPartySyntheticMonitor(id string, title string, locations []SyntheticTestLocation, scheduleIntervalInSeconds int32, ) *Model3rdPartySyntheticMonitor {
	this := Model3rdPartySyntheticMonitor{}
	this.Id = id
	this.Title = title
	this.Locations = locations
	this.ScheduleIntervalInSeconds = scheduleIntervalInSeconds
	return &this
}

// NewModel3rdPartySyntheticMonitorWithDefaults instantiates a new Model3rdPartySyntheticMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel3rdPartySyntheticMonitorWithDefaults() *Model3rdPartySyntheticMonitor {
	this := Model3rdPartySyntheticMonitor{}
	return &this
}

// GetId returns the Id field value
func (o *Model3rdPartySyntheticMonitor) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Model3rdPartySyntheticMonitor) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *Model3rdPartySyntheticMonitor) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Model3rdPartySyntheticMonitor) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Model3rdPartySyntheticMonitor) SetDescription(v string) {
	o.Description = &v
}

// GetTestSetup returns the TestSetup field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetTestSetup() string {
	if o == nil || o.TestSetup == nil {
		var ret string
		return ret
	}
	return *o.TestSetup
}

// GetTestSetupOk returns a tuple with the TestSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetTestSetupOk() (*string, bool) {
	if o == nil || o.TestSetup == nil {
		return nil, false
	}
	return o.TestSetup, true
}

// HasTestSetup returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasTestSetup() bool {
	if o != nil && o.TestSetup != nil {
		return true
	}

	return false
}

// SetTestSetup gets a reference to the given string and assigns it to the TestSetup field.
func (o *Model3rdPartySyntheticMonitor) SetTestSetup(v string) {
	o.TestSetup = &v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetExpirationTimestamp() int64 {
	if o == nil || o.ExpirationTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetExpirationTimestampOk() (*int64, bool) {
	if o == nil || o.ExpirationTimestamp == nil {
		return nil, false
	}
	return o.ExpirationTimestamp, true
}

// HasExpirationTimestamp returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasExpirationTimestamp() bool {
	if o != nil && o.ExpirationTimestamp != nil {
		return true
	}

	return false
}

// SetExpirationTimestamp gets a reference to the given int64 and assigns it to the ExpirationTimestamp field.
func (o *Model3rdPartySyntheticMonitor) SetExpirationTimestamp(v int64) {
	o.ExpirationTimestamp = &v
}

// GetDrilldownLink returns the DrilldownLink field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetDrilldownLink() string {
	if o == nil || o.DrilldownLink == nil {
		var ret string
		return ret
	}
	return *o.DrilldownLink
}

// GetDrilldownLinkOk returns a tuple with the DrilldownLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetDrilldownLinkOk() (*string, bool) {
	if o == nil || o.DrilldownLink == nil {
		return nil, false
	}
	return o.DrilldownLink, true
}

// HasDrilldownLink returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasDrilldownLink() bool {
	if o != nil && o.DrilldownLink != nil {
		return true
	}

	return false
}

// SetDrilldownLink gets a reference to the given string and assigns it to the DrilldownLink field.
func (o *Model3rdPartySyntheticMonitor) SetDrilldownLink(v string) {
	o.DrilldownLink = &v
}

// GetEditLink returns the EditLink field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetEditLink() string {
	if o == nil || o.EditLink == nil {
		var ret string
		return ret
	}
	return *o.EditLink
}

// GetEditLinkOk returns a tuple with the EditLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetEditLinkOk() (*string, bool) {
	if o == nil || o.EditLink == nil {
		return nil, false
	}
	return o.EditLink, true
}

// HasEditLink returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasEditLink() bool {
	if o != nil && o.EditLink != nil {
		return true
	}

	return false
}

// SetEditLink gets a reference to the given string and assigns it to the EditLink field.
func (o *Model3rdPartySyntheticMonitor) SetEditLink(v string) {
	o.EditLink = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Model3rdPartySyntheticMonitor) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Model3rdPartySyntheticMonitor) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetLocations returns the Locations field value
func (o *Model3rdPartySyntheticMonitor) GetLocations() []SyntheticTestLocation {
	if o == nil  {
		var ret []SyntheticTestLocation
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetLocationsOk() (*[]SyntheticTestLocation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locations, true
}

// SetLocations sets field value
func (o *Model3rdPartySyntheticMonitor) SetLocations(v []SyntheticTestLocation) {
	o.Locations = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetSteps() []SyntheticTestStep {
	if o == nil || o.Steps == nil {
		var ret []SyntheticTestStep
		return ret
	}
	return *o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetStepsOk() (*[]SyntheticTestStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasSteps() bool {
	if o != nil && o.Steps != nil {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []SyntheticTestStep and assigns it to the Steps field.
func (o *Model3rdPartySyntheticMonitor) SetSteps(v []SyntheticTestStep) {
	o.Steps = &v
}

// GetScheduleIntervalInSeconds returns the ScheduleIntervalInSeconds field value
func (o *Model3rdPartySyntheticMonitor) GetScheduleIntervalInSeconds() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ScheduleIntervalInSeconds
}

// GetScheduleIntervalInSecondsOk returns a tuple with the ScheduleIntervalInSeconds field value
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetScheduleIntervalInSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScheduleIntervalInSeconds, true
}

// SetScheduleIntervalInSeconds sets field value
func (o *Model3rdPartySyntheticMonitor) SetScheduleIntervalInSeconds(v int32) {
	o.ScheduleIntervalInSeconds = v
}

// GetNoDataTimeout returns the NoDataTimeout field value if set, zero value otherwise.
func (o *Model3rdPartySyntheticMonitor) GetNoDataTimeout() int32 {
	if o == nil || o.NoDataTimeout == nil {
		var ret int32
		return ret
	}
	return *o.NoDataTimeout
}

// GetNoDataTimeoutOk returns a tuple with the NoDataTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model3rdPartySyntheticMonitor) GetNoDataTimeoutOk() (*int32, bool) {
	if o == nil || o.NoDataTimeout == nil {
		return nil, false
	}
	return o.NoDataTimeout, true
}

// HasNoDataTimeout returns a boolean if a field has been set.
func (o *Model3rdPartySyntheticMonitor) HasNoDataTimeout() bool {
	if o != nil && o.NoDataTimeout != nil {
		return true
	}

	return false
}

// SetNoDataTimeout gets a reference to the given int32 and assigns it to the NoDataTimeout field.
func (o *Model3rdPartySyntheticMonitor) SetNoDataTimeout(v int32) {
	o.NoDataTimeout = &v
}

func (o Model3rdPartySyntheticMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TestSetup != nil {
		toSerialize["testSetup"] = o.TestSetup
	}
	if o.ExpirationTimestamp != nil {
		toSerialize["expirationTimestamp"] = o.ExpirationTimestamp
	}
	if o.DrilldownLink != nil {
		toSerialize["drilldownLink"] = o.DrilldownLink
	}
	if o.EditLink != nil {
		toSerialize["editLink"] = o.EditLink
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if true {
		toSerialize["locations"] = o.Locations
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if true {
		toSerialize["scheduleIntervalInSeconds"] = o.ScheduleIntervalInSeconds
	}
	if o.NoDataTimeout != nil {
		toSerialize["noDataTimeout"] = o.NoDataTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableModel3rdPartySyntheticMonitor struct {
	value *Model3rdPartySyntheticMonitor
	isSet bool
}

func (v NullableModel3rdPartySyntheticMonitor) Get() *Model3rdPartySyntheticMonitor {
	return v.value
}

func (v *NullableModel3rdPartySyntheticMonitor) Set(val *Model3rdPartySyntheticMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableModel3rdPartySyntheticMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableModel3rdPartySyntheticMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel3rdPartySyntheticMonitor(val *Model3rdPartySyntheticMonitor) *NullableModel3rdPartySyntheticMonitor {
	return &NullableModel3rdPartySyntheticMonitor{value: val, isSet: true}
}

func (v NullableModel3rdPartySyntheticMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel3rdPartySyntheticMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


