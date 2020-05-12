/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// NumberComparisonInfo Comparison for `NUMBER` attributes.
type NumberComparisonInfo struct {
	// Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Comparison string `json:"comparison"`
	// The value to compare to.
	Value float32 `json:"value,omitempty"`
	// Reverse the comparison **operator**. For example, it turns **equals** into **does not equal**.
	Negate bool `json:"negate"`
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `STRING` -> StringComparisonInfo  * `NUMBER` -> NumberComparisonInfo  * `BOOLEAN` -> BooleanComparisonInfo  * `HTTP_METHOD` -> HttpMethodComparisonInfo  * `STRING_REQUEST_ATTRIBUTE` -> StringRequestAttributeComparisonInfo  * `NUMBER_REQUEST_ATTRIBUTE` -> NumberRequestAttributeComparisonInfo  * `ZOS_CALL_TYPE` -> ZosComparisonInfo  * `IIB_INPUT_NODE_TYPE` -> IIBInputNodeTypeComparisonInfo  * `ESB_INPUT_NODE_TYPE` -> ESBInputNodeTypeComparisonInfo  * `FAILED_STATE` -> FailedStateComparisonInfo  * `FLAW_STATE` -> FlawStateComparisonInfo  * `FAILURE_REASON` -> FailureReasonComparisonInfo  * `HTTP_STATUS_CLASS` -> HttpStatusClassComparisonInfo  * `TAG` -> TagComparisonInfo  * `FAST_STRING` -> FastStringComparisonInfo  * `SERVICE_TYPE` -> ServiceTypeComparisonInfo  
	Type string `json:"type"`
}
