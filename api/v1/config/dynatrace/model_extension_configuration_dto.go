/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ExtensionConfigurationDto struct for ExtensionConfigurationDto
type ExtensionConfigurationDto struct {
	// The ID of the extension.
	ExtensionId string `json:"extensionId,omitempty"`
	// The extension is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// Allows to skip current configuration and use global one.
	UseGlobal bool `json:"useGlobal"`
	// The list of extension parameters.    Each parameter is a key-value pair.
	Properties map[string]string `json:"properties,omitempty"`
	// The ID of the host on which the extension runs.
	HostId string `json:"hostId,omitempty"`
	ActiveGate EntityShortRepresentation `json:"activeGate,omitempty"`
	// The ID of the endpoint.
	EndpointId string `json:"endpointId,omitempty"`
	// The name of the endpoint, displayed in Dynatrace.
	EndpointName string `json:"endpointName,omitempty"`
}
