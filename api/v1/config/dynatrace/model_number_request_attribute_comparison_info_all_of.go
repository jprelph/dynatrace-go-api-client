/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// NumberRequestAttributeComparisonInfoAllOf struct for NumberRequestAttributeComparisonInfoAllOf
type NumberRequestAttributeComparisonInfoAllOf struct {
	// Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Comparison string `json:"comparison,omitempty"`
	// The value to compare to.
	Value float32 `json:"value,omitempty"`
	RequestAttribute string `json:"requestAttribute,omitempty"`
	// If `true`, the request attribute is matched on child service calls.     Default is `false`.
	MatchOnChildCalls bool `json:"matchOnChildCalls,omitempty"`
}
