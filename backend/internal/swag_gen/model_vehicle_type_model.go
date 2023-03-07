/*
 * Bumbal Client Api
 *
 * Bumbal API documentation
 *
 * API version: 2.0
 * Contact: info@bumbal.eu
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// VehicleTypeModel struct for VehicleTypeModel
type VehicleTypeModel struct {
	// Unique ID
	Id *int64 `json:"id,omitempty"`
	// The name of the VehicleType
	Name *string `json:"name,omitempty"`
	// Max Speed in km/h, Bikes (id 4) ignore max_speed
	MaxSpeed *int64 `json:"max_speed,omitempty"`
	// Speed Factor
	SpeedFactor *float64 `json:"speed_factor,omitempty"`
	// Duration Factor
	DurationFactor *float64 `json:"duration_factor,omitempty"`
	// Cost per meter
	CostPerMeter *float64 `json:"cost_per_meter,omitempty"`
	// Cost per route
	CostPerRoute *float64 `json:"cost_per_route,omitempty"`
	// Cost per driving minute
	CostPerDrivingMinute *float64 `json:"cost_per_driving_minute,omitempty"`
	// Cost per waiting minute
	CostPerWaitingMinute *float64 `json:"cost_per_waiting_minute,omitempty"`
	// Cost per service minute
	CostPerServiceMinute *float64 `json:"cost_per_service_minute,omitempty"`
	// created_at date time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// updated_at date time
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewVehicleTypeModel instantiates a new VehicleTypeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleTypeModel() *VehicleTypeModel {
	this := VehicleTypeModel{}
	return &this
}

// NewVehicleTypeModelWithDefaults instantiates a new VehicleTypeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleTypeModelWithDefaults() *VehicleTypeModel {
	this := VehicleTypeModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *VehicleTypeModel) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VehicleTypeModel) SetName(v string) {
	o.Name = &v
}

// GetMaxSpeed returns the MaxSpeed field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetMaxSpeed() int64 {
	if o == nil || o.MaxSpeed == nil {
		var ret int64
		return ret
	}
	return *o.MaxSpeed
}

// GetMaxSpeedOk returns a tuple with the MaxSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetMaxSpeedOk() (*int64, bool) {
	if o == nil || o.MaxSpeed == nil {
		return nil, false
	}
	return o.MaxSpeed, true
}

// HasMaxSpeed returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasMaxSpeed() bool {
	if o != nil && o.MaxSpeed != nil {
		return true
	}

	return false
}

// SetMaxSpeed gets a reference to the given int64 and assigns it to the MaxSpeed field.
func (o *VehicleTypeModel) SetMaxSpeed(v int64) {
	o.MaxSpeed = &v
}

// GetSpeedFactor returns the SpeedFactor field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetSpeedFactor() float64 {
	if o == nil || o.SpeedFactor == nil {
		var ret float64
		return ret
	}
	return *o.SpeedFactor
}

// GetSpeedFactorOk returns a tuple with the SpeedFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetSpeedFactorOk() (*float64, bool) {
	if o == nil || o.SpeedFactor == nil {
		return nil, false
	}
	return o.SpeedFactor, true
}

// HasSpeedFactor returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasSpeedFactor() bool {
	if o != nil && o.SpeedFactor != nil {
		return true
	}

	return false
}

// SetSpeedFactor gets a reference to the given float64 and assigns it to the SpeedFactor field.
func (o *VehicleTypeModel) SetSpeedFactor(v float64) {
	o.SpeedFactor = &v
}

// GetDurationFactor returns the DurationFactor field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetDurationFactor() float64 {
	if o == nil || o.DurationFactor == nil {
		var ret float64
		return ret
	}
	return *o.DurationFactor
}

// GetDurationFactorOk returns a tuple with the DurationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetDurationFactorOk() (*float64, bool) {
	if o == nil || o.DurationFactor == nil {
		return nil, false
	}
	return o.DurationFactor, true
}

// HasDurationFactor returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasDurationFactor() bool {
	if o != nil && o.DurationFactor != nil {
		return true
	}

	return false
}

// SetDurationFactor gets a reference to the given float64 and assigns it to the DurationFactor field.
func (o *VehicleTypeModel) SetDurationFactor(v float64) {
	o.DurationFactor = &v
}

// GetCostPerMeter returns the CostPerMeter field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetCostPerMeter() float64 {
	if o == nil || o.CostPerMeter == nil {
		var ret float64
		return ret
	}
	return *o.CostPerMeter
}

// GetCostPerMeterOk returns a tuple with the CostPerMeter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetCostPerMeterOk() (*float64, bool) {
	if o == nil || o.CostPerMeter == nil {
		return nil, false
	}
	return o.CostPerMeter, true
}

// HasCostPerMeter returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasCostPerMeter() bool {
	if o != nil && o.CostPerMeter != nil {
		return true
	}

	return false
}

// SetCostPerMeter gets a reference to the given float64 and assigns it to the CostPerMeter field.
func (o *VehicleTypeModel) SetCostPerMeter(v float64) {
	o.CostPerMeter = &v
}

// GetCostPerRoute returns the CostPerRoute field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetCostPerRoute() float64 {
	if o == nil || o.CostPerRoute == nil {
		var ret float64
		return ret
	}
	return *o.CostPerRoute
}

// GetCostPerRouteOk returns a tuple with the CostPerRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetCostPerRouteOk() (*float64, bool) {
	if o == nil || o.CostPerRoute == nil {
		return nil, false
	}
	return o.CostPerRoute, true
}

// HasCostPerRoute returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasCostPerRoute() bool {
	if o != nil && o.CostPerRoute != nil {
		return true
	}

	return false
}

// SetCostPerRoute gets a reference to the given float64 and assigns it to the CostPerRoute field.
func (o *VehicleTypeModel) SetCostPerRoute(v float64) {
	o.CostPerRoute = &v
}

// GetCostPerDrivingMinute returns the CostPerDrivingMinute field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetCostPerDrivingMinute() float64 {
	if o == nil || o.CostPerDrivingMinute == nil {
		var ret float64
		return ret
	}
	return *o.CostPerDrivingMinute
}

// GetCostPerDrivingMinuteOk returns a tuple with the CostPerDrivingMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetCostPerDrivingMinuteOk() (*float64, bool) {
	if o == nil || o.CostPerDrivingMinute == nil {
		return nil, false
	}
	return o.CostPerDrivingMinute, true
}

// HasCostPerDrivingMinute returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasCostPerDrivingMinute() bool {
	if o != nil && o.CostPerDrivingMinute != nil {
		return true
	}

	return false
}

// SetCostPerDrivingMinute gets a reference to the given float64 and assigns it to the CostPerDrivingMinute field.
func (o *VehicleTypeModel) SetCostPerDrivingMinute(v float64) {
	o.CostPerDrivingMinute = &v
}

// GetCostPerWaitingMinute returns the CostPerWaitingMinute field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetCostPerWaitingMinute() float64 {
	if o == nil || o.CostPerWaitingMinute == nil {
		var ret float64
		return ret
	}
	return *o.CostPerWaitingMinute
}

// GetCostPerWaitingMinuteOk returns a tuple with the CostPerWaitingMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetCostPerWaitingMinuteOk() (*float64, bool) {
	if o == nil || o.CostPerWaitingMinute == nil {
		return nil, false
	}
	return o.CostPerWaitingMinute, true
}

// HasCostPerWaitingMinute returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasCostPerWaitingMinute() bool {
	if o != nil && o.CostPerWaitingMinute != nil {
		return true
	}

	return false
}

// SetCostPerWaitingMinute gets a reference to the given float64 and assigns it to the CostPerWaitingMinute field.
func (o *VehicleTypeModel) SetCostPerWaitingMinute(v float64) {
	o.CostPerWaitingMinute = &v
}

// GetCostPerServiceMinute returns the CostPerServiceMinute field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetCostPerServiceMinute() float64 {
	if o == nil || o.CostPerServiceMinute == nil {
		var ret float64
		return ret
	}
	return *o.CostPerServiceMinute
}

// GetCostPerServiceMinuteOk returns a tuple with the CostPerServiceMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetCostPerServiceMinuteOk() (*float64, bool) {
	if o == nil || o.CostPerServiceMinute == nil {
		return nil, false
	}
	return o.CostPerServiceMinute, true
}

// HasCostPerServiceMinute returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasCostPerServiceMinute() bool {
	if o != nil && o.CostPerServiceMinute != nil {
		return true
	}

	return false
}

// SetCostPerServiceMinute gets a reference to the given float64 and assigns it to the CostPerServiceMinute field.
func (o *VehicleTypeModel) SetCostPerServiceMinute(v float64) {
	o.CostPerServiceMinute = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VehicleTypeModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VehicleTypeModel) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleTypeModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VehicleTypeModel) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VehicleTypeModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o VehicleTypeModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.MaxSpeed != nil {
		toSerialize["max_speed"] = o.MaxSpeed
	}
	if o.SpeedFactor != nil {
		toSerialize["speed_factor"] = o.SpeedFactor
	}
	if o.DurationFactor != nil {
		toSerialize["duration_factor"] = o.DurationFactor
	}
	if o.CostPerMeter != nil {
		toSerialize["cost_per_meter"] = o.CostPerMeter
	}
	if o.CostPerRoute != nil {
		toSerialize["cost_per_route"] = o.CostPerRoute
	}
	if o.CostPerDrivingMinute != nil {
		toSerialize["cost_per_driving_minute"] = o.CostPerDrivingMinute
	}
	if o.CostPerWaitingMinute != nil {
		toSerialize["cost_per_waiting_minute"] = o.CostPerWaitingMinute
	}
	if o.CostPerServiceMinute != nil {
		toSerialize["cost_per_service_minute"] = o.CostPerServiceMinute
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableVehicleTypeModel struct {
	value *VehicleTypeModel
	isSet bool
}

func (v NullableVehicleTypeModel) Get() *VehicleTypeModel {
	return v.value
}

func (v *NullableVehicleTypeModel) Set(val *VehicleTypeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleTypeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleTypeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleTypeModel(val *VehicleTypeModel) *NullableVehicleTypeModel {
	return &NullableVehicleTypeModel{value: val, isSet: true}
}

func (v NullableVehicleTypeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleTypeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

