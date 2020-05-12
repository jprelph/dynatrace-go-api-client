/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// EntityIdAlertingScope A scope filter for a monitored entity identifier.
type EntityIdAlertingScope struct {
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `ENTITY_ID` -> EntityIdAlertingScope  * `MANAGEMENT_ZONE` -> ManagementZoneAlertingScope  * `TAG` -> TagFilterAlertingScope  * `NAME` -> NameAlertingScope  * `CUSTOM_DEVICE_GROUP_NAME` -> CustomDeviceGroupNameAlertingScope  * `HOST_GROUP_NAME` -> HostGroupNameAlertingScope  * `HOST_NAME` -> HostNameAlertingScope  * `PROCESS_GROUP_ID` -> ProcessGroupIdAlertingScope  * `PROCESS_GROUP_NAME` -> ProcessGroupNameAlertingScope  
	FilterType string `json:"filterType"`
	// The monitored entities id to match on.
	EntityId string `json:"entityId"`
}
