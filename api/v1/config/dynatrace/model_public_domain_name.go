/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// PublicDomainName The contribution to the service ID calculation from the domain name where the web request has been detected.    You have two mutually exclusive options:  * Override the detected value with a specified static value. Specify the new value in the **valueOverride** field.  * Dynamically transform the detected value. Specify the transformation parameters in the **transformations** field.
type PublicDomainName struct {
	// Transformations to be applied to the detected value.
	Transformations []TransformationBase `json:"transformations,omitempty"`
	// The value to be used instead of the detected value.
	ValueOverride string `json:"valueOverride,omitempty"`
	// Use (`true`) or don't use (`false`) the detected host name as base for transformation.    Not applicable if the override is specified.
	CopyFromHostName bool `json:"copyFromHostName,omitempty"`
}
