/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace

// Schedule The schedule of the maintenance window.
type Schedule struct {
	// The type of the schedule recurrence.
	RecurrenceType string      `json:"recurrenceType"`
	Recurrence     *Recurrence `json:"recurrence,omitempty"`
	// The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
	Start string `json:"start"`
	// The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
	End string `json:"end"`
	// The time zone of the start and end time. Default time zone is UTC.   You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`).
	ZoneId string `json:"zoneId"`
}
