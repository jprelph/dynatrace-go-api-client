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

// SchemaType A list of definitions of types.    A type is a complex property that contains its own set of subproperties.
type SchemaType struct {
	// The pattern for the summary (for example, \"Alert after *X* minutes.\") of the configuration in the UI.
	SummaryPattern *string `json:"summaryPattern,omitempty"`
	// A short description of the version.
	VersionInfo *string `json:"versionInfo,omitempty"`
	// Definition of properties that can be persisted.
	Properties *map[string]PropertyDefinition `json:"properties,omitempty"`
	// The version of the type.
	Version *string `json:"version,omitempty"`
	// A list of constraints limiting the values to be accepted.
	Constraints *[]ComplexConstraint `json:"constraints,omitempty"`
	// An extended description and/or links to documentation.
	Documentation *string `json:"documentation,omitempty"`
	// The display name of the property.
	DisplayName *string `json:"displayName,omitempty"`
	// A short description of the property.
	Description *string `json:"description,omitempty"`
}

// NewSchemaType instantiates a new SchemaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaType() *SchemaType {
	this := SchemaType{}
	return &this
}

// NewSchemaTypeWithDefaults instantiates a new SchemaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaTypeWithDefaults() *SchemaType {
	this := SchemaType{}
	return &this
}

// GetSummaryPattern returns the SummaryPattern field value if set, zero value otherwise.
func (o *SchemaType) GetSummaryPattern() string {
	if o == nil || o.SummaryPattern == nil {
		var ret string
		return ret
	}
	return *o.SummaryPattern
}

// GetSummaryPatternOk returns a tuple with the SummaryPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetSummaryPatternOk() (*string, bool) {
	if o == nil || o.SummaryPattern == nil {
		return nil, false
	}
	return o.SummaryPattern, true
}

// HasSummaryPattern returns a boolean if a field has been set.
func (o *SchemaType) HasSummaryPattern() bool {
	if o != nil && o.SummaryPattern != nil {
		return true
	}

	return false
}

// SetSummaryPattern gets a reference to the given string and assigns it to the SummaryPattern field.
func (o *SchemaType) SetSummaryPattern(v string) {
	o.SummaryPattern = &v
}

// GetVersionInfo returns the VersionInfo field value if set, zero value otherwise.
func (o *SchemaType) GetVersionInfo() string {
	if o == nil || o.VersionInfo == nil {
		var ret string
		return ret
	}
	return *o.VersionInfo
}

// GetVersionInfoOk returns a tuple with the VersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetVersionInfoOk() (*string, bool) {
	if o == nil || o.VersionInfo == nil {
		return nil, false
	}
	return o.VersionInfo, true
}

// HasVersionInfo returns a boolean if a field has been set.
func (o *SchemaType) HasVersionInfo() bool {
	if o != nil && o.VersionInfo != nil {
		return true
	}

	return false
}

// SetVersionInfo gets a reference to the given string and assigns it to the VersionInfo field.
func (o *SchemaType) SetVersionInfo(v string) {
	o.VersionInfo = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *SchemaType) GetProperties() map[string]PropertyDefinition {
	if o == nil || o.Properties == nil {
		var ret map[string]PropertyDefinition
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetPropertiesOk() (*map[string]PropertyDefinition, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *SchemaType) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]PropertyDefinition and assigns it to the Properties field.
func (o *SchemaType) SetProperties(v map[string]PropertyDefinition) {
	o.Properties = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SchemaType) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SchemaType) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SchemaType) SetVersion(v string) {
	o.Version = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *SchemaType) GetConstraints() []ComplexConstraint {
	if o == nil || o.Constraints == nil {
		var ret []ComplexConstraint
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetConstraintsOk() (*[]ComplexConstraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *SchemaType) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []ComplexConstraint and assigns it to the Constraints field.
func (o *SchemaType) SetConstraints(v []ComplexConstraint) {
	o.Constraints = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *SchemaType) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *SchemaType) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *SchemaType) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SchemaType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SchemaType) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SchemaType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemaType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemaType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemaType) SetDescription(v string) {
	o.Description = &v
}

func (o SchemaType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SummaryPattern != nil {
		toSerialize["summaryPattern"] = o.SummaryPattern
	}
	if o.VersionInfo != nil {
		toSerialize["versionInfo"] = o.VersionInfo
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaType struct {
	value *SchemaType
	isSet bool
}

func (v NullableSchemaType) Get() *SchemaType {
	return v.value
}

func (v *NullableSchemaType) Set(val *SchemaType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaType(val *SchemaType) *NullableSchemaType {
	return &NullableSchemaType{value: val, isSet: true}
}

func (v NullableSchemaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


