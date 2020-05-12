/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// RemoveCreditCardNumbersTransformation The transformation of the `REMOVE_CREDIT_CARDS` type.   The transformation automatically detects and removes credit card numbers. No additional parameters needed.
type RemoveCreditCardNumbersTransformation struct {
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `BEFORE` -> BeforeTransformation  * `AFTER` -> AfterTransformation  * `BETWEEN` -> BetweenTransformation  * `REPLACE_BETWEEN` -> ReplaceBetweenTransformation  * `REMOVE_NUMBERS` -> RemoveNumbersTransformation  * `REMOVE_CREDIT_CARDS` -> RemoveCreditCardNumbersTransformation  * `REMOVE_IBANS` -> RemoveIBANsTransformation  * `REMOVE_IPS` -> RemoveIPsTransformation  * `SPLIT_SELECT` -> SplitSelectTransformation  * `TAKE_SEGMENTS` -> TakeSegmentsTransformation  
	Type string `json:"type"`
}
