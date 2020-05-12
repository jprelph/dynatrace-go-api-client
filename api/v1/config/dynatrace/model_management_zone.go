/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ManagementZone The configuration of the management zone. It defines where the management zone should be applied.
type ManagementZone struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	// The ID of the management zone.
	Id string `json:"id,omitempty"`
	// The name of the management zone.
	Name string `json:"name"`
	// The list of rules to where the management zone applies. Each rule is evaluated independently of all other rules.
	Rules []ManagementZoneRule `json:"rules,omitempty"`
}
