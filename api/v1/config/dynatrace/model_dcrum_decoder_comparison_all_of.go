/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// DcrumDecoderComparisonAllOf struct for DcrumDecoderComparisonAllOf
type DcrumDecoderComparisonAllOf struct {
	// Operator of the comparison.   You can reverse it by setting **negate** to `true`.   Example: EQUALS, EXISTS   Example operators might not necessarily be available for the selected type
	Operator string `json:"operator,omitempty"`
	// The value to compare to.
	Value string `json:"value,omitempty"`
}
