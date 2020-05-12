/*
 * Dynatrace Configuration API
 *
 * Documentation of the Dynatrace Configuration API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/config-api) to read about use-cases and examples.  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dynatrace
// PaasTypeComparison Comparison for `PAAS_TYPE` attributes.
type PaasTypeComparison struct {
	// Operator of the comparison.   You can reverse it by setting **negate** to `true`.   Example: EQUALS, EXISTS   Example operators might not necessarily be available for the selected type
	Operator string `json:"operator"`
	// The value to compare to.
	Value string `json:"value,omitempty"`
	// Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**.
	Negate bool `json:"negate"`
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `STRING` -> StringComparison  * `INDEXED_NAME` -> IndexedNameComparison  * `INDEXED_STRING` -> IndexedStringComparison  * `INTEGER` -> IntegerComparison  * `SERVICE_TYPE` -> ServiceTypeComparison  * `PAAS_TYPE` -> PaasTypeComparison  * `CLOUD_TYPE` -> CloudTypeComparison  * `AZURE_SKU` -> AzureSkuComparision  * `AZURE_COMPUTE_MODE` -> AzureComputeModeComparison  * `ENTITY_ID` -> EntityIdComparison  * `SIMPLE_TECH` -> SimpleTechComparison  * `SIMPLE_HOST_TECH` -> SimpleHostTechComparison  * `SERVICE_TOPOLOGY` -> ServiceTopologyComparison  * `DATABASE_TOPOLOGY` -> DatabaseTopologyComparison  * `OS_TYPE` -> OsTypeComparison  * `HYPERVISOR_TYPE` -> HypervisorTypeComparision  * `IP_ADDRESS` -> IpAddressComparison  * `OS_ARCHITECTURE` -> OsArchitectureComparison  * `BITNESS` -> BitnessComparision  * `APPLICATION_TYPE` -> ApplicationTypeComparison  * `MOBILE_PLATFORM` -> MobilePlatformComparison  * `CUSTOM_APPLICATION_TYPE` -> CustomApplicationTypeComparison  * `DCRUM_DECODER_TYPE` -> DcrumDecoderComparison  * `SYNTHETIC_ENGINE_TYPE` -> SyntheticEngineTypeComparison  * `TAG` -> TagComparison  * `INDEXED_TAG` -> IndexedTagComparison  
	Type string `json:"type"`
}
