/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// ResponseTimeDegradationAutodetectionConfig Parameters of the response time degradation auto-detection. Required if the **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.   Violation of **any** criterion triggers an alert.
type ResponseTimeDegradationAutodetectionConfig struct {
	// Alert if the response time degrades by more than *X* milliseconds.
	ResponseTimeDegradationMilliseconds int32 `json:"responseTimeDegradationMilliseconds"`
	// Alert if the response time degrades by more than *X* %.
	ResponseTimeDegradationPercent int32 `json:"responseTimeDegradationPercent"`
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds.
	SlowestResponseTimeDegradationMilliseconds int32 `json:"slowestResponseTimeDegradationMilliseconds"`
	// Alert if the response time of the slowest 10% degrades by more than *X* %.
	SlowestResponseTimeDegradationPercent int32 `json:"slowestResponseTimeDegradationPercent"`
	// Minimal service load to detect response time degradation.    Response time degradation of services with smaller load won't trigger alerts.
	LoadThreshold string `json:"loadThreshold"`
}
