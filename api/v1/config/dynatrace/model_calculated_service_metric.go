/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// CalculatedServiceMetric Descriptor of a calculated service metric.
type CalculatedServiceMetric struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	// The key of the calculated service metric.
	TsmMetricKey string `json:"tsmMetricKey"`
	// The displayed name of the metric.
	Name string `json:"name"`
	// The metric is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	MetricDefinition CalculatedMetricDefinition `json:"metricDefinition"`
	// The unit of the metric.
	Unit string `json:"unit"`
	// The display name of the metric's unit.    Only applicable when the **unit** parameter is set to `UNSPECIFIED`.
	UnitDisplayName string `json:"unitDisplayName,omitempty"`
	// Restricts the metric usage to the specified service.    This field is mutually exclusive with the **managementZones** field.
	EntityId string `json:"entityId,omitempty"`
	// Restricts the metric usage to specified management zones.    This field is mutually exclusive with the **entityId** field.
	ManagementZones []string `json:"managementZones,omitempty"`
	// The set of conditions for the metric usage.    **All** the specified conditions must be fulfilled to use the metric.
	Conditions []Condition `json:"conditions,omitempty"`
	DimensionDefinition DimensionDefinition `json:"dimensionDefinition,omitempty"`
}
