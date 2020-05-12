/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// PagerDutyNotificationConfigAllOf struct for PagerDutyNotificationConfigAllOf
type PagerDutyNotificationConfigAllOf struct {
	// The name of the PagerDuty account.
	Account string `json:"account,omitempty"`
	// The API key to access PagerDuty.
	ServiceApiKey string `json:"serviceApiKey,omitempty"`
	// The name of the service.
	ServiceName string `json:"serviceName,omitempty"`
}
