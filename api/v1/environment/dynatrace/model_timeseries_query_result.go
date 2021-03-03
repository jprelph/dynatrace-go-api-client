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

// TimeseriesQueryResult The configuration of a metric with all its parameters and, optionally, data points.
type TimeseriesQueryResult struct {
	// The ID of the metric.
	TimeseriesId *string `json:"timeseriesId,omitempty"`
	// The name of the metric in the user interface.
	DisplayName *string `json:"displayName,omitempty"`
	// The fine metric division, for example process group and process ID for some process-related metric.
	Dimensions *[]string `json:"dimensions,omitempty"`
	// The list of allowed aggregations for this metric.
	AggregationTypes *[]string `json:"aggregationTypes,omitempty"`
	// The unit of the metric.
	Unit *string `json:"unit,omitempty"`
	// The feature, where the metric originates.
	Filter *string `json:"filter,omitempty"`
	// The feature, where the metric originates.
	DetailedSource *string `json:"detailedSource,omitempty"`
	// The ID of the plugin, where the metric originates.
	PluginId *string `json:"pluginId,omitempty"`
	// Technology type definition. Used to group metrics under a logical technology name.
	Types *[]string `json:"types,omitempty"`
	DataResult *TimeseriesDataPointQueryResult `json:"dataResult,omitempty"`
}

// NewTimeseriesQueryResult instantiates a new TimeseriesQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesQueryResult() *TimeseriesQueryResult {
	this := TimeseriesQueryResult{}
	return &this
}

// NewTimeseriesQueryResultWithDefaults instantiates a new TimeseriesQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesQueryResultWithDefaults() *TimeseriesQueryResult {
	this := TimeseriesQueryResult{}
	return &this
}

// GetTimeseriesId returns the TimeseriesId field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetTimeseriesId() string {
	if o == nil || o.TimeseriesId == nil {
		var ret string
		return ret
	}
	return *o.TimeseriesId
}

// GetTimeseriesIdOk returns a tuple with the TimeseriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetTimeseriesIdOk() (*string, bool) {
	if o == nil || o.TimeseriesId == nil {
		return nil, false
	}
	return o.TimeseriesId, true
}

// HasTimeseriesId returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasTimeseriesId() bool {
	if o != nil && o.TimeseriesId != nil {
		return true
	}

	return false
}

// SetTimeseriesId gets a reference to the given string and assigns it to the TimeseriesId field.
func (o *TimeseriesQueryResult) SetTimeseriesId(v string) {
	o.TimeseriesId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TimeseriesQueryResult) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetDimensions() []string {
	if o == nil || o.Dimensions == nil {
		var ret []string
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetDimensionsOk() (*[]string, bool) {
	if o == nil || o.Dimensions == nil {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasDimensions() bool {
	if o != nil && o.Dimensions != nil {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []string and assigns it to the Dimensions field.
func (o *TimeseriesQueryResult) SetDimensions(v []string) {
	o.Dimensions = &v
}

// GetAggregationTypes returns the AggregationTypes field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetAggregationTypes() []string {
	if o == nil || o.AggregationTypes == nil {
		var ret []string
		return ret
	}
	return *o.AggregationTypes
}

// GetAggregationTypesOk returns a tuple with the AggregationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetAggregationTypesOk() (*[]string, bool) {
	if o == nil || o.AggregationTypes == nil {
		return nil, false
	}
	return o.AggregationTypes, true
}

// HasAggregationTypes returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasAggregationTypes() bool {
	if o != nil && o.AggregationTypes != nil {
		return true
	}

	return false
}

// SetAggregationTypes gets a reference to the given []string and assigns it to the AggregationTypes field.
func (o *TimeseriesQueryResult) SetAggregationTypes(v []string) {
	o.AggregationTypes = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *TimeseriesQueryResult) SetUnit(v string) {
	o.Unit = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *TimeseriesQueryResult) SetFilter(v string) {
	o.Filter = &v
}

// GetDetailedSource returns the DetailedSource field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetDetailedSource() string {
	if o == nil || o.DetailedSource == nil {
		var ret string
		return ret
	}
	return *o.DetailedSource
}

// GetDetailedSourceOk returns a tuple with the DetailedSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetDetailedSourceOk() (*string, bool) {
	if o == nil || o.DetailedSource == nil {
		return nil, false
	}
	return o.DetailedSource, true
}

// HasDetailedSource returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasDetailedSource() bool {
	if o != nil && o.DetailedSource != nil {
		return true
	}

	return false
}

// SetDetailedSource gets a reference to the given string and assigns it to the DetailedSource field.
func (o *TimeseriesQueryResult) SetDetailedSource(v string) {
	o.DetailedSource = &v
}

// GetPluginId returns the PluginId field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetPluginId() string {
	if o == nil || o.PluginId == nil {
		var ret string
		return ret
	}
	return *o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetPluginIdOk() (*string, bool) {
	if o == nil || o.PluginId == nil {
		return nil, false
	}
	return o.PluginId, true
}

// HasPluginId returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasPluginId() bool {
	if o != nil && o.PluginId != nil {
		return true
	}

	return false
}

// SetPluginId gets a reference to the given string and assigns it to the PluginId field.
func (o *TimeseriesQueryResult) SetPluginId(v string) {
	o.PluginId = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetTypesOk() (*[]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *TimeseriesQueryResult) SetTypes(v []string) {
	o.Types = &v
}

// GetDataResult returns the DataResult field value if set, zero value otherwise.
func (o *TimeseriesQueryResult) GetDataResult() TimeseriesDataPointQueryResult {
	if o == nil || o.DataResult == nil {
		var ret TimeseriesDataPointQueryResult
		return ret
	}
	return *o.DataResult
}

// GetDataResultOk returns a tuple with the DataResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesQueryResult) GetDataResultOk() (*TimeseriesDataPointQueryResult, bool) {
	if o == nil || o.DataResult == nil {
		return nil, false
	}
	return o.DataResult, true
}

// HasDataResult returns a boolean if a field has been set.
func (o *TimeseriesQueryResult) HasDataResult() bool {
	if o != nil && o.DataResult != nil {
		return true
	}

	return false
}

// SetDataResult gets a reference to the given TimeseriesDataPointQueryResult and assigns it to the DataResult field.
func (o *TimeseriesQueryResult) SetDataResult(v TimeseriesDataPointQueryResult) {
	o.DataResult = &v
}

func (o TimeseriesQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TimeseriesId != nil {
		toSerialize["timeseriesId"] = o.TimeseriesId
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Dimensions != nil {
		toSerialize["dimensions"] = o.Dimensions
	}
	if o.AggregationTypes != nil {
		toSerialize["aggregationTypes"] = o.AggregationTypes
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.DetailedSource != nil {
		toSerialize["detailedSource"] = o.DetailedSource
	}
	if o.PluginId != nil {
		toSerialize["pluginId"] = o.PluginId
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.DataResult != nil {
		toSerialize["dataResult"] = o.DataResult
	}
	return json.Marshal(toSerialize)
}

type NullableTimeseriesQueryResult struct {
	value *TimeseriesQueryResult
	isSet bool
}

func (v NullableTimeseriesQueryResult) Get() *TimeseriesQueryResult {
	return v.value
}

func (v *NullableTimeseriesQueryResult) Set(val *TimeseriesQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesQueryResult(val *TimeseriesQueryResult) *NullableTimeseriesQueryResult {
	return &NullableTimeseriesQueryResult{value: val, isSet: true}
}

func (v NullableTimeseriesQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


