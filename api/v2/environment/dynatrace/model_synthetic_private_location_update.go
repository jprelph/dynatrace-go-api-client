/*
 * Dynatrace Environment API
 *
 *  Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dynatrace

import (
	"encoding/json"
)

// SyntheticPrivateLocationUpdate Configuration of a private synthetic location
type SyntheticPrivateLocationUpdate struct {
	SyntheticLocationUpdate
	// A list of synthetic nodes belonging to the location.    You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call.
	Nodes []string `json:"nodes"`
	// The name of the location.
	Name string `json:"name"`
	// The country code of the location.    Use the alpha-2 code of the [ISO 3166-2 standard](https://dt-url.net/iso3166-2), (for example, `AT` for Austria or `PL` for Poland).
	CountryCode *string `json:"countryCode,omitempty"`
	// The region code of the location.    For the [USA](https://dt-url.net/iso3166us) or [Canada](https://dt-url.net/iso3166ca) use ISO 3166-2 state codes (without `US-` or `CA-` prefix), for example, `VA` for Virginia or `OR` for Oregon.    For the rest of the world use [FIPS 10-4 codes](https://dt-url.net/fipscodes).
	RegionCode *string `json:"regionCode,omitempty"`
	// The city of the location.
	City *string `json:"city,omitempty"`
	// The latitude of the location in `DDD.dddd` format.
	Latitude float64 `json:"latitude"`
	// The longitude of the location in `DDD.dddd` format.
	Longitude float64 `json:"longitude"`
	// The status of the location:   * `ENABLED`: The location is displayed as active in the UI. You can assign monitors to the location.  * `DISABLED`: The location is displayed as inactive in the UI. You can't assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location.  * `HIDDEN`: The location is not displayed in the UI. You can't assign monitors to the location. You can only set location as `HIDDEN` when no monitor is assigned to it.
	Status *string `json:"status,omitempty"`
	// The alerting of location outage is enabled (`true`) or disabled (`false`).
	AvailabilityLocationOutage *bool `json:"availabilityLocationOutage,omitempty"`
	// The alerting of node outage is enabled (`true`) or disabled (`false`).    If enabled, the outage of *any* node in the location triggers an alert.
	AvailabilityNodeOutage *bool `json:"availabilityNodeOutage,omitempty"`
	// Alert if the location or node outage lasts longer than *X* minutes.    Only applicable when **availabilityLocationOutage** or **availabilityNodeOutage** is set to `true`.
	LocationNodeOutageDelayInMinutes *int32 `json:"locationNodeOutageDelayInMinutes,omitempty"`
	// The notifications of location and node outage is enabled (`true`) or disabled (`false`).
	AvailabilityNotificationsEnabled *bool `json:"availabilityNotificationsEnabled,omitempty"`
}

// NewSyntheticPrivateLocationUpdate instantiates a new SyntheticPrivateLocationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticPrivateLocationUpdate(nodes []string, name string, latitude float64, longitude float64) *SyntheticPrivateLocationUpdate {
	this := SyntheticPrivateLocationUpdate{}
	this.Nodes = nodes
	this.Name = name
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewSyntheticPrivateLocationUpdateWithDefaults instantiates a new SyntheticPrivateLocationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticPrivateLocationUpdateWithDefaults() *SyntheticPrivateLocationUpdate {
	this := SyntheticPrivateLocationUpdate{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *SyntheticPrivateLocationUpdate) GetNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetNodesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *SyntheticPrivateLocationUpdate) SetNodes(v []string) {
	o.Nodes = v
}

// GetName returns the Name field value
func (o *SyntheticPrivateLocationUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SyntheticPrivateLocationUpdate) SetName(v string) {
	o.Name = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *SyntheticPrivateLocationUpdate) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *SyntheticPrivateLocationUpdate) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *SyntheticPrivateLocationUpdate) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *SyntheticPrivateLocationUpdate) GetRegionCode() string {
	if o == nil || o.RegionCode == nil {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetRegionCodeOk() (*string, bool) {
	if o == nil || o.RegionCode == nil {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *SyntheticPrivateLocationUpdate) HasRegionCode() bool {
	if o != nil && o.RegionCode != nil {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *SyntheticPrivateLocationUpdate) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SyntheticPrivateLocationUpdate) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SyntheticPrivateLocationUpdate) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SyntheticPrivateLocationUpdate) SetCity(v string) {
	o.City = &v
}

// GetLatitude returns the Latitude field value
func (o *SyntheticPrivateLocationUpdate) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetLatitudeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *SyntheticPrivateLocationUpdate) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *SyntheticPrivateLocationUpdate) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetLongitudeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *SyntheticPrivateLocationUpdate) SetLongitude(v float64) {
	o.Longitude = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticPrivateLocationUpdate) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticPrivateLocationUpdate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticPrivateLocationUpdate) SetStatus(v string) {
	o.Status = &v
}

// GetAvailabilityLocationOutage returns the AvailabilityLocationOutage field value if set, zero value otherwise.
func (o *SyntheticPrivateLocationUpdate) GetAvailabilityLocationOutage() bool {
	if o == nil || o.AvailabilityLocationOutage == nil {
		var ret bool
		return ret
	}
	return *o.AvailabilityLocationOutage
}

// GetAvailabilityLocationOutageOk returns a tuple with the AvailabilityLocationOutage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetAvailabilityLocationOutageOk() (*bool, bool) {
	if o == nil || o.AvailabilityLocationOutage == nil {
		return nil, false
	}
	return o.AvailabilityLocationOutage, true
}

// HasAvailabilityLocationOutage returns a boolean if a field has been set.
func (o *SyntheticPrivateLocationUpdate) HasAvailabilityLocationOutage() bool {
	if o != nil && o.AvailabilityLocationOutage != nil {
		return true
	}

	return false
}

// SetAvailabilityLocationOutage gets a reference to the given bool and assigns it to the AvailabilityLocationOutage field.
func (o *SyntheticPrivateLocationUpdate) SetAvailabilityLocationOutage(v bool) {
	o.AvailabilityLocationOutage = &v
}

// GetAvailabilityNodeOutage returns the AvailabilityNodeOutage field value if set, zero value otherwise.
func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNodeOutage() bool {
	if o == nil || o.AvailabilityNodeOutage == nil {
		var ret bool
		return ret
	}
	return *o.AvailabilityNodeOutage
}

// GetAvailabilityNodeOutageOk returns a tuple with the AvailabilityNodeOutage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNodeOutageOk() (*bool, bool) {
	if o == nil || o.AvailabilityNodeOutage == nil {
		return nil, false
	}
	return o.AvailabilityNodeOutage, true
}

// HasAvailabilityNodeOutage returns a boolean if a field has been set.
func (o *SyntheticPrivateLocationUpdate) HasAvailabilityNodeOutage() bool {
	if o != nil && o.AvailabilityNodeOutage != nil {
		return true
	}

	return false
}

// SetAvailabilityNodeOutage gets a reference to the given bool and assigns it to the AvailabilityNodeOutage field.
func (o *SyntheticPrivateLocationUpdate) SetAvailabilityNodeOutage(v bool) {
	o.AvailabilityNodeOutage = &v
}

// GetLocationNodeOutageDelayInMinutes returns the LocationNodeOutageDelayInMinutes field value if set, zero value otherwise.
func (o *SyntheticPrivateLocationUpdate) GetLocationNodeOutageDelayInMinutes() int32 {
	if o == nil || o.LocationNodeOutageDelayInMinutes == nil {
		var ret int32
		return ret
	}
	return *o.LocationNodeOutageDelayInMinutes
}

// GetLocationNodeOutageDelayInMinutesOk returns a tuple with the LocationNodeOutageDelayInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetLocationNodeOutageDelayInMinutesOk() (*int32, bool) {
	if o == nil || o.LocationNodeOutageDelayInMinutes == nil {
		return nil, false
	}
	return o.LocationNodeOutageDelayInMinutes, true
}

// HasLocationNodeOutageDelayInMinutes returns a boolean if a field has been set.
func (o *SyntheticPrivateLocationUpdate) HasLocationNodeOutageDelayInMinutes() bool {
	if o != nil && o.LocationNodeOutageDelayInMinutes != nil {
		return true
	}

	return false
}

// SetLocationNodeOutageDelayInMinutes gets a reference to the given int32 and assigns it to the LocationNodeOutageDelayInMinutes field.
func (o *SyntheticPrivateLocationUpdate) SetLocationNodeOutageDelayInMinutes(v int32) {
	o.LocationNodeOutageDelayInMinutes = &v
}

// GetAvailabilityNotificationsEnabled returns the AvailabilityNotificationsEnabled field value if set, zero value otherwise.
func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNotificationsEnabled() bool {
	if o == nil || o.AvailabilityNotificationsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AvailabilityNotificationsEnabled
}

// GetAvailabilityNotificationsEnabledOk returns a tuple with the AvailabilityNotificationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNotificationsEnabledOk() (*bool, bool) {
	if o == nil || o.AvailabilityNotificationsEnabled == nil {
		return nil, false
	}
	return o.AvailabilityNotificationsEnabled, true
}

// HasAvailabilityNotificationsEnabled returns a boolean if a field has been set.
func (o *SyntheticPrivateLocationUpdate) HasAvailabilityNotificationsEnabled() bool {
	if o != nil && o.AvailabilityNotificationsEnabled != nil {
		return true
	}

	return false
}

// SetAvailabilityNotificationsEnabled gets a reference to the given bool and assigns it to the AvailabilityNotificationsEnabled field.
func (o *SyntheticPrivateLocationUpdate) SetAvailabilityNotificationsEnabled(v bool) {
	o.AvailabilityNotificationsEnabled = &v
}

func (o SyntheticPrivateLocationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSyntheticLocationUpdate, errSyntheticLocationUpdate := json.Marshal(o.SyntheticLocationUpdate)
	if errSyntheticLocationUpdate != nil {
		return []byte{}, errSyntheticLocationUpdate
	}
	errSyntheticLocationUpdate = json.Unmarshal([]byte(serializedSyntheticLocationUpdate), &toSerialize)
	if errSyntheticLocationUpdate != nil {
		return []byte{}, errSyntheticLocationUpdate
	}
	if true {
		toSerialize["nodes"] = o.Nodes
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.RegionCode != nil {
		toSerialize["regionCode"] = o.RegionCode
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["latitude"] = o.Latitude
	}
	if true {
		toSerialize["longitude"] = o.Longitude
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.AvailabilityLocationOutage != nil {
		toSerialize["availabilityLocationOutage"] = o.AvailabilityLocationOutage
	}
	if o.AvailabilityNodeOutage != nil {
		toSerialize["availabilityNodeOutage"] = o.AvailabilityNodeOutage
	}
	if o.LocationNodeOutageDelayInMinutes != nil {
		toSerialize["locationNodeOutageDelayInMinutes"] = o.LocationNodeOutageDelayInMinutes
	}
	if o.AvailabilityNotificationsEnabled != nil {
		toSerialize["availabilityNotificationsEnabled"] = o.AvailabilityNotificationsEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticPrivateLocationUpdate struct {
	value *SyntheticPrivateLocationUpdate
	isSet bool
}

func (v NullableSyntheticPrivateLocationUpdate) Get() *SyntheticPrivateLocationUpdate {
	return v.value
}

func (v *NullableSyntheticPrivateLocationUpdate) Set(val *SyntheticPrivateLocationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticPrivateLocationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticPrivateLocationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticPrivateLocationUpdate(val *SyntheticPrivateLocationUpdate) *NullableSyntheticPrivateLocationUpdate {
	return &NullableSyntheticPrivateLocationUpdate{value: val, isSet: true}
}

func (v NullableSyntheticPrivateLocationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticPrivateLocationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


