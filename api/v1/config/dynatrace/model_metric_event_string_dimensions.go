/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// MetricEventStringDimensions A filter for the metrics string dimensions.
type MetricEventStringDimensions struct {
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `ENTITY` -> MetricEventEntityDimensions  * `STRING` -> MetricEventStringDimensions  
	FilterType string `json:"filterType"`
	// The dimensions name. Sending this has no effect while creating a configuration, as only the *index* of the dimension is used here -> dimension names might change, indexes not.
	Name string `json:"name,omitempty"`
	// The dimensions index on the metric.
	Index int32 `json:"index"`
	TextFilter MetricEventTextFilterMetricEventDimensionsFilterOperatorDto `json:"textFilter"`
}
