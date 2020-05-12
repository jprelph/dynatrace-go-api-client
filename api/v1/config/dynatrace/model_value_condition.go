/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ValueCondition CICS SDK node name condition for which the value is captured.   This  is required if the **source** is: `CICS_SDK`.   Not applicable in other cases.
type ValueCondition struct {
	// Operator comparing the extracted value to the comparison value.
	Operator string `json:"operator"`
	// Negate the comparison.
	Negate bool `json:"negate"`
	// The value to compare to.
	Value string `json:"value"`
}
