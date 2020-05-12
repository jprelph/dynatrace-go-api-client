/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// CustomHostMetadataConditionKey Key for dynamic attributes with the `HOST_CUSTOM_METADATA_KEY` key format.
type CustomHostMetadataConditionKey struct {
	// The attribute to be used for comparision.
	Attribute string `json:"attribute"`
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `PROCESS_CUSTOM_METADATA_KEY` -> CustomProcessMetadataConditionKey  * `HOST_CUSTOM_METADATA_KEY` -> CustomHostMetadataConditionKey  * `PROCESS_PREDEFINED_METADATA_KEY` -> ProcessMetadataConditionKey  * `STRING` -> StringConditionKey  
	Type string `json:"type,omitempty"`
	DynamicKey CustomHostMetadataKey `json:"dynamicKey"`
}
