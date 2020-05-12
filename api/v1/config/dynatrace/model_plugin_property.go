/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// PluginProperty A property of a plugin.
type PluginProperty struct {
	// The key of the property.
	Key string `json:"key,omitempty"`
	// The type of the property.
	Type string `json:"type,omitempty"`
	// The default value of the property.
	DefaultValue string `json:"defaultValue,omitempty"`
	// The list of possible values of the property.    If such a list is defined, only values from this list can be assigned to the property.
	DropdownValues []string `json:"dropdownValues,omitempty"`
}
