/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// KubernetesCredentials Configuration for specific Kubernetes credentials.
type KubernetesCredentials struct {
	Metadata ConfigurationMetadata `json:"metadata,omitempty"`
	// The ID of the given credentials configuration.
	Id string `json:"id,omitempty"`
	// The monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.   If not set on creation, the `true` value is used.   If the field is omitted during an update, the old value remains unaffected.
	Active bool `json:"active,omitempty"`
	// The status of the configured endpoint.  ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing. UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing. DISABLED: The credentials have been disabled by the user. FASTCHECK_AUTH_ERROR: The credentials are invalid. FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid. FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached. FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid. FASTCHECK_AUTH_LOCKED: The credentials seem to be locked. UNKNOWN: An unknown error occured. 
	EndpointStatus string `json:"endpointStatus,omitempty"`
	// The detailed status info of the configured endpoint.
	EndpointStatusInfo string `json:"endpointStatusInfo,omitempty"`
	// The name of the Kubernetes credentials configuration.   Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.
	Label string `json:"label"`
	// The URL of the Kubernetes API server.   It must be unique within a Dynatrace environment.   The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed.
	EndpointUrl string `json:"endpointUrl,omitempty"`
	// The service account bearer token for the Kubernetes API server.   Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as `null`.   If the field is omitted during an update, the old value remains unaffected.
	AuthToken string `json:"authToken,omitempty"`
	// The monitoring of events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. Event monitoring depends on the active state of this configuration to be true.   If not set on creation, the `true` value is used.   If the field is omitted during an update, the old value remains unaffected.
	EventsIntegrationEnabled bool `json:"eventsIntegrationEnabled,omitempty"`
	// Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.
	WorkloadIntegrationEnabled bool `json:"workloadIntegrationEnabled,omitempty"`
	// Kubernetes event filters based on field-selectors. If set to `null` on creation, no events field selectors are subscribed. If set to `null` on update, no change of stored events field selectors is applied. Set an empty list to clear all events field selectors.
	EventsFieldSelectors []KubernetesEventPattern `json:"eventsFieldSelectors,omitempty"`
	// The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.   If not set on creation, the `true` value is used.   If the field is omitted during an update, the old value remains unaffected.
	CertificateCheckEnabled bool `json:"certificateCheckEnabled,omitempty"`
}
