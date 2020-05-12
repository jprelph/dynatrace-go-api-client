/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// JavaScriptInjectionRules Rules for javascript injection
type JavaScriptInjectionRules struct {
	// The enable or disable rule of the java script injection.
	Enabled bool `json:"enabled"`
	// The url operator of the java script injection.
	UrlOperator string `json:"urlOperator"`
	// The url pattern of the java script injection.
	UrlPattern string `json:"urlPattern,omitempty"`
	// The url rule of the java script injection.
	Rule string `json:"rule"`
	// The html pattern of the java script injection.
	HtmlPattern string `json:"htmlPattern,omitempty"`
}
