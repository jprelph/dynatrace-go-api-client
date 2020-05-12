/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// GuestCpuLimitThresholds Custom thresholds for guest CPU limit detection. If not set then the automatic mode is used.    **All** conditions must be fulfilled to trigger an alert.
type GuestCpuLimitThresholds struct {
	// Hypervisor CPU usage is higher than *X*% in 3 out of 5 samples.
	HostCpuUsageMinPercentage int32 `json:"hostCpuUsageMinPercentage"`
	// VM CPU usage (VM CPU Usage Mhz / VM CPU limit in Mhz) is higher than *X*% in 3 out of 5 samples.
	VmCpuUsageMaxPercentage int32 `json:"vmCpuUsageMaxPercentage"`
	// VM CPU ready is higher than *X*% occurred in 3 out of 5 samples.
	VmCpuReadyMaxPercentage int32 `json:"vmCpuReadyMaxPercentage"`
}
