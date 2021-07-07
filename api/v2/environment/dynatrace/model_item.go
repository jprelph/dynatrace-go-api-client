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

// Item An item of a collection property.
type Item struct {
	// The type referenced by the item's value.
	ReferencedType *string `json:"referencedType,omitempty"`
	// An extended description and/or links to documentation.
	Documentation *string `json:"documentation,omitempty"`
	// The subtype of the item's value.
	SubType *string `json:"subType,omitempty"`
	// The type of the item's value.
	Type *map[string]interface{} `json:"type,omitempty"`
	// The display name of the item.
	DisplayName *string `json:"displayName,omitempty"`
	// A short description of the item.
	Description *string `json:"description,omitempty"`
	// A list of constraints limiting the values to be accepted.
	Constraints *[]Constraint `json:"constraints,omitempty"`
}

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem() *Item {
	this := Item{}
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetReferencedType returns the ReferencedType field value if set, zero value otherwise.
func (o *Item) GetReferencedType() string {
	if o == nil || o.ReferencedType == nil {
		var ret string
		return ret
	}
	return *o.ReferencedType
}

// GetReferencedTypeOk returns a tuple with the ReferencedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetReferencedTypeOk() (*string, bool) {
	if o == nil || o.ReferencedType == nil {
		return nil, false
	}
	return o.ReferencedType, true
}

// HasReferencedType returns a boolean if a field has been set.
func (o *Item) HasReferencedType() bool {
	if o != nil && o.ReferencedType != nil {
		return true
	}

	return false
}

// SetReferencedType gets a reference to the given string and assigns it to the ReferencedType field.
func (o *Item) SetReferencedType(v string) {
	o.ReferencedType = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *Item) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *Item) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *Item) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *Item) GetSubType() string {
	if o == nil || o.SubType == nil {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetSubTypeOk() (*string, bool) {
	if o == nil || o.SubType == nil {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *Item) HasSubType() bool {
	if o != nil && o.SubType != nil {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *Item) SetSubType(v string) {
	o.SubType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Item) GetType() map[string]interface{} {
	if o == nil || o.Type == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetTypeOk() (*map[string]interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Item) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *Item) SetType(v map[string]interface{}) {
	o.Type = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Item) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Item) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Item) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Item) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Item) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Item) SetDescription(v string) {
	o.Description = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *Item) GetConstraints() []Constraint {
	if o == nil || o.Constraints == nil {
		var ret []Constraint
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetConstraintsOk() (*[]Constraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *Item) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []Constraint and assigns it to the Constraints field.
func (o *Item) SetConstraints(v []Constraint) {
	o.Constraints = &v
}

func (o Item) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferencedType != nil {
		toSerialize["referencedType"] = o.ReferencedType
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	if o.SubType != nil {
		toSerialize["subType"] = o.SubType
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	return json.Marshal(toSerialize)
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


