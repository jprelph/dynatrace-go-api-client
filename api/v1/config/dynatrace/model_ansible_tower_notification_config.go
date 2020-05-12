/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// AnsibleTowerNotificationConfig Configuration of the Ansible Tower notification.
type AnsibleTowerNotificationConfig struct {
	// The ID of the notification configuration.
	Id string `json:"id,omitempty"`
	// The name of the notification configuration.
	Name string `json:"name"`
	// The ID of the associated alerting profile.
	AlertingProfile string `json:"alertingProfile"`
	// The configuration is enabled (`true`) or disabled (`false`).
	Active bool `json:"active"`
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `EMAIL` -> EmailNotificationConfig  * `PAGER_DUTY` -> PagerDutyNotificationConfig  * `WEBHOOK` -> WebHookNotificationConfig  * `SLACK` -> SlackNotificationConfig  * `HIPCHAT` -> HipChatNotificationConfig  * `VICTOROPS` -> VictorOpsNotificationConfig  * `SERVICE_NOW` -> ServiceNowNotificationConfig  * `XMATTERS` -> XMattersNotificationConfig  * `ANSIBLETOWER` -> AnsibleTowerNotificationConfig  * `OPS_GENIE` -> OpsGenieNotificationConfig  * `JIRA` -> JiraNotificationConfig  * `TRELLO` -> TrelloNotificationConfig  
	Type string `json:"type"`
	// The URL of the target Ansible Tower job template.
	JobTemplateURL string `json:"jobTemplateURL"`
	// Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
	AcceptAnyCertificate bool `json:"acceptAnyCertificate"`
	// The username of the Ansible Tower account.
	Username string `json:"username"`
	// The password for the Ansible Tower account.
	Password string `json:"password,omitempty"`
	// The ID of the target Ansible Tower job template.
	JobTemplateID int32 `json:"jobTemplateID"`
	// The custom message of the notification.    This message will be displayed in the extra variables **Message** field of your job template.   You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`. If the problem has been merged into another problem, it has the `MERGED` value.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.  
	CustomMessage string `json:"customMessage"`
}
