/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// AwsAuthenticationData A credentials for the AWS authentication.
type AwsAuthenticationData struct {
	// The type of the authentication: role-based or key-based.
	Type string `json:"type"`
	KeyBasedAuthentication KeyBasedAuthentication `json:"keyBasedAuthentication,omitempty"`
	RoleBasedAuthentication RoleBasedAuthentication `json:"roleBasedAuthentication,omitempty"`
}
