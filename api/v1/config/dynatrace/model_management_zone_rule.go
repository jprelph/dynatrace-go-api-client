/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ManagementZoneRule A rule defines when to apply a management zone. Each rule is evaluated independently of all other rules.
type ManagementZoneRule struct {
	// The type of Dynatrace entities the management zone can be applied to.
	Type string `json:"type"`
	// The evaluation rule is enabled(`true`) or disabled(`false`).
	Enabled bool `json:"enabled"`
	// How to apply the management zone to underlying entities:   `SERVICE_TO_HOST_LIKE`: Apply to underlying hosts of matching services.  `SERVICE_TO_PROCESS_GROUP_LIKE`: Apply to underlying process groups of matching services.   `PROCESS_GROUP_TO_HOST`: Apply to underlying hosts of matching process groups.  `PROCESS_GROUP_TO_SERVICE`: Apply to all services provided by matching process groups.   `HOST_TO_PROCESS_GROUP_INSTANCE`: Apply to processes running on matching hosts.   `CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE`: Apply to custom devices in matching custom device groups.
	PropagationTypes []string `json:"propagationTypes,omitempty"`
	// Matching rules of the management zone. The management zone applies only if all conditions are fulfilled.
	Conditions []EntityRuleEngineCondition `json:"conditions"`
}
