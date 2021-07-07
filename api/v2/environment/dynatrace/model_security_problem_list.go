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

// SecurityProblemList A list of security problems.
type SecurityProblemList struct {
	// The total number of entries in the result.
	TotalCount int64 `json:"totalCount"`
	// The number of entries per page.
	PageSize *int32 `json:"pageSize,omitempty"`
	// The cursor for the next page of results. Has the value of `null` on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result.
	NextPageKey *string `json:"nextPageKey,omitempty"`
	// A list of security problems.
	SecurityProblems *[]SecurityProblem `json:"securityProblems,omitempty"`
}

// NewSecurityProblemList instantiates a new SecurityProblemList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityProblemList(totalCount int64) *SecurityProblemList {
	this := SecurityProblemList{}
	this.TotalCount = totalCount
	return &this
}

// NewSecurityProblemListWithDefaults instantiates a new SecurityProblemList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityProblemListWithDefaults() *SecurityProblemList {
	this := SecurityProblemList{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *SecurityProblemList) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SecurityProblemList) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SecurityProblemList) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *SecurityProblemList) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblemList) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *SecurityProblemList) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *SecurityProblemList) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetNextPageKey returns the NextPageKey field value if set, zero value otherwise.
func (o *SecurityProblemList) GetNextPageKey() string {
	if o == nil || o.NextPageKey == nil {
		var ret string
		return ret
	}
	return *o.NextPageKey
}

// GetNextPageKeyOk returns a tuple with the NextPageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblemList) GetNextPageKeyOk() (*string, bool) {
	if o == nil || o.NextPageKey == nil {
		return nil, false
	}
	return o.NextPageKey, true
}

// HasNextPageKey returns a boolean if a field has been set.
func (o *SecurityProblemList) HasNextPageKey() bool {
	if o != nil && o.NextPageKey != nil {
		return true
	}

	return false
}

// SetNextPageKey gets a reference to the given string and assigns it to the NextPageKey field.
func (o *SecurityProblemList) SetNextPageKey(v string) {
	o.NextPageKey = &v
}

// GetSecurityProblems returns the SecurityProblems field value if set, zero value otherwise.
func (o *SecurityProblemList) GetSecurityProblems() []SecurityProblem {
	if o == nil || o.SecurityProblems == nil {
		var ret []SecurityProblem
		return ret
	}
	return *o.SecurityProblems
}

// GetSecurityProblemsOk returns a tuple with the SecurityProblems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblemList) GetSecurityProblemsOk() (*[]SecurityProblem, bool) {
	if o == nil || o.SecurityProblems == nil {
		return nil, false
	}
	return o.SecurityProblems, true
}

// HasSecurityProblems returns a boolean if a field has been set.
func (o *SecurityProblemList) HasSecurityProblems() bool {
	if o != nil && o.SecurityProblems != nil {
		return true
	}

	return false
}

// SetSecurityProblems gets a reference to the given []SecurityProblem and assigns it to the SecurityProblems field.
func (o *SecurityProblemList) SetSecurityProblems(v []SecurityProblem) {
	o.SecurityProblems = &v
}

func (o SecurityProblemList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.NextPageKey != nil {
		toSerialize["nextPageKey"] = o.NextPageKey
	}
	if o.SecurityProblems != nil {
		toSerialize["securityProblems"] = o.SecurityProblems
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityProblemList struct {
	value *SecurityProblemList
	isSet bool
}

func (v NullableSecurityProblemList) Get() *SecurityProblemList {
	return v.value
}

func (v *NullableSecurityProblemList) Set(val *SecurityProblemList) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityProblemList) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityProblemList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityProblemList(val *SecurityProblemList) *NullableSecurityProblemList {
	return &NullableSecurityProblemList{value: val, isSet: true}
}

func (v NullableSecurityProblemList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityProblemList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


