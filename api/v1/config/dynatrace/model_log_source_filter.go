/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// LogSourceFilter Configuration of the log filter for a custom log metric.   If several criteria are specified, the AND logic applies.
type LogSourceFilter struct {
	// A list of filtering criteria for log path.   If several criteria are specified, the OR logic applies.
	PathDefinitions []PathDefinition `json:"pathDefinitions,omitempty"`
	// A list of Dynatrace entities, where the log can originate from. Specify Dynatrace entity IDs here.    Allowed types of entities are `PROCESS_GROUP` and `PROCESS_GROUP_INSTANCE`. You can't mix entity types.   If several entities are specified, the OR logic applies.   This field is mutually exclusive with the **osTypes** field.
	SourceEntities []string `json:"sourceEntities,omitempty"`
	// A list of hosts, where the log can originate from. Specify Dynatrace entity IDs here.   If several hosts are specified, the OR logic applies.
	HostFilters []string `json:"hostFilters,omitempty"`
	// A list of operating system types, where the log can originate from.   If set, **only OS logs** are included in the result.   If several OS are specified, the OR logic applies.   This field is mutually exclusive with the **sourceEntities** field.
	OsTypes []string `json:"osTypes,omitempty"`
}
