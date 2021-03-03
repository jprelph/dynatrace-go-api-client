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

// HostsListPage A list of hosts with OneAgent deployment information for each host.
type HostsListPage struct {
	// The progress of the environment search, in percent.
	PercentageOfEnvironmentSearched *float64 `json:"percentageOfEnvironmentSearched,omitempty"`
	// The cursor for the next page of results.    Has the value of `null` on the last page.   There might be another page of results even if the current page is empty.
	NextPageKey *string `json:"nextPageKey,omitempty"`
	// A list of hosts with OneAgent deployment information for each host.
	Hosts *[]HostAgentInfo `json:"hosts,omitempty"`
}

// NewHostsListPage instantiates a new HostsListPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostsListPage() *HostsListPage {
	this := HostsListPage{}
	return &this
}

// NewHostsListPageWithDefaults instantiates a new HostsListPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostsListPageWithDefaults() *HostsListPage {
	this := HostsListPage{}
	return &this
}

// GetPercentageOfEnvironmentSearched returns the PercentageOfEnvironmentSearched field value if set, zero value otherwise.
func (o *HostsListPage) GetPercentageOfEnvironmentSearched() float64 {
	if o == nil || o.PercentageOfEnvironmentSearched == nil {
		var ret float64
		return ret
	}
	return *o.PercentageOfEnvironmentSearched
}

// GetPercentageOfEnvironmentSearchedOk returns a tuple with the PercentageOfEnvironmentSearched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsListPage) GetPercentageOfEnvironmentSearchedOk() (*float64, bool) {
	if o == nil || o.PercentageOfEnvironmentSearched == nil {
		return nil, false
	}
	return o.PercentageOfEnvironmentSearched, true
}

// HasPercentageOfEnvironmentSearched returns a boolean if a field has been set.
func (o *HostsListPage) HasPercentageOfEnvironmentSearched() bool {
	if o != nil && o.PercentageOfEnvironmentSearched != nil {
		return true
	}

	return false
}

// SetPercentageOfEnvironmentSearched gets a reference to the given float64 and assigns it to the PercentageOfEnvironmentSearched field.
func (o *HostsListPage) SetPercentageOfEnvironmentSearched(v float64) {
	o.PercentageOfEnvironmentSearched = &v
}

// GetNextPageKey returns the NextPageKey field value if set, zero value otherwise.
func (o *HostsListPage) GetNextPageKey() string {
	if o == nil || o.NextPageKey == nil {
		var ret string
		return ret
	}
	return *o.NextPageKey
}

// GetNextPageKeyOk returns a tuple with the NextPageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsListPage) GetNextPageKeyOk() (*string, bool) {
	if o == nil || o.NextPageKey == nil {
		return nil, false
	}
	return o.NextPageKey, true
}

// HasNextPageKey returns a boolean if a field has been set.
func (o *HostsListPage) HasNextPageKey() bool {
	if o != nil && o.NextPageKey != nil {
		return true
	}

	return false
}

// SetNextPageKey gets a reference to the given string and assigns it to the NextPageKey field.
func (o *HostsListPage) SetNextPageKey(v string) {
	o.NextPageKey = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *HostsListPage) GetHosts() []HostAgentInfo {
	if o == nil || o.Hosts == nil {
		var ret []HostAgentInfo
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsListPage) GetHostsOk() (*[]HostAgentInfo, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *HostsListPage) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []HostAgentInfo and assigns it to the Hosts field.
func (o *HostsListPage) SetHosts(v []HostAgentInfo) {
	o.Hosts = &v
}

func (o HostsListPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PercentageOfEnvironmentSearched != nil {
		toSerialize["percentageOfEnvironmentSearched"] = o.PercentageOfEnvironmentSearched
	}
	if o.NextPageKey != nil {
		toSerialize["nextPageKey"] = o.NextPageKey
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	return json.Marshal(toSerialize)
}

type NullableHostsListPage struct {
	value *HostsListPage
	isSet bool
}

func (v NullableHostsListPage) Get() *HostsListPage {
	return v.value
}

func (v *NullableHostsListPage) Set(val *HostsListPage) {
	v.value = val
	v.isSet = true
}

func (v NullableHostsListPage) IsSet() bool {
	return v.isSet
}

func (v *NullableHostsListPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostsListPage(val *HostsListPage) *NullableHostsListPage {
	return &NullableHostsListPage{value: val, isSet: true}
}

func (v NullableHostsListPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostsListPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


