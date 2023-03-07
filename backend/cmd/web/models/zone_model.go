package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
)

type ZoneDBModel struct {
	DB *mongo.Database
}

// ZoneModel struct for ZoneModel
type ZoneModel struct {
	// Unique Zone ID
	Id *int64 `json:"id,omitempty"`
	//Zones Ranges
	ZoneRanges *[]ZoneRangeModel `json:"zone_ranges,omitempty"`
	//Total fuel cost
	ZoneFuelCost *float64 `json:"zone_fuel_cost,omitempty"`
	//Total driving time
	ZoneDrivingTime *int64 `json:"zone_driving_time,omitempty"`
}

// NewZoneModel instantiates a new ZoneModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneModel() *ZoneModel {
	this := ZoneModel{}
	return &this
}

// NewZoneModelWithDefaults instantiates a new ZoneModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneModelWithDefaults() *ZoneModel {
	this := ZoneModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ZoneModel) SetId(v int64) {
	o.Id = &v
}

// GetZoneRanges returns the ZoneRanges field value if set, zero value otherwise.
func (o *ZoneModel) GetZoneRanges() []ZoneRangeModel {
	if o == nil || o.ZoneRanges == nil {
		var ret []ZoneRangeModel
		return ret
	}
	return *o.ZoneRanges
}

// GetZoneRangesOk returns a tuple with the ZoneRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModel) GetZoneRangesOk() (*[]ZoneRangeModel, bool) {
	if o == nil || o.ZoneRanges == nil {
		return nil, false
	}
	return o.ZoneRanges, true
}

// HasZoneRanges returns a boolean if a field has been set.
func (o *ZoneModel) HasZoneRanges() bool {
	if o != nil && o.ZoneRanges != nil {
		return true
	}

	return false
}

// SetZoneRanges gets a reference to the given []ZoneRangeModel and assigns it to the ZoneRanges field.
func (o *ZoneModel) SetZoneRanges(v []ZoneRangeModel) {
	o.ZoneRanges = &v
}

// GetZoneFuelCost returns the FuelCost field value if set, zero value otherwise.
func (o *ZoneModel) GetZoneFuelCost() float64 {
	if o == nil || o.ZoneFuelCost == nil {
		var ret float64
		return ret
	}
	return *o.ZoneFuelCost
}

// GetZoneFuelCostOk returns a tuple with the ZoneFuelCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModel) GetZoneFuelCostOk() (*float64, bool) {
	if o == nil || o.ZoneFuelCost == nil {
		return nil, false
	}
	return o.ZoneFuelCost, true
}

// HasZoneFuelCost returns a boolean if a field has been set.
func (o *ZoneModel) HasZoneFuelCost() bool {
	if o != nil && o.ZoneFuelCost != nil {
		return true
	}

	return false
}

// SetZoneFuelCost gets a reference to the given int32 and assigns it to the ZoneFuelCost field.
func (o *ZoneModel) SetZoneFuelCost(v float64) {
	o.ZoneFuelCost = &v
}

// GetZoneDrivingTime returns the ZoneDrivingTime field value if set, zero value otherwise.
func (o *ZoneModel) GetZoneDrivingTime() int64 {
	if o == nil || o.ZoneDrivingTime == nil {
		var ret int64
		return ret
	}
	return *o.ZoneDrivingTime
}

// GetZoneDrivingTimeOk returns a tuple with the ZoneZoneDrivingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModel) GetZoneDrivingTimeOk() (*int64, bool) {
	if o == nil || o.ZoneDrivingTime == nil {
		return nil, false
	}
	return o.ZoneDrivingTime, true
}

// HasZoneDrivingTime returns a boolean if a field has been set.
func (o *ZoneModel) HasZoneDrivingTime() bool {
	if o != nil && o.ZoneDrivingTime != nil {
		return true
	}

	return false
}

// SetZoneDrivingTime gets a reference to the given int32 and assigns it to the ZoneDrivingTime field.
func (o *ZoneModel) SetZoneDrivingTime(v int64) {
	o.ZoneDrivingTime = &v
}

func (o ZoneModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ZoneRanges != nil {
		toSerialize["zone_ranges"] = o.ZoneRanges
	}
	if o.ZoneFuelCost != nil {
		toSerialize["zone_fuel_cost"] = o.ZoneFuelCost
	}
	if o.ZoneDrivingTime != nil {
		toSerialize["zone_driving_time"] = o.ZoneDrivingTime
	}

	return json.Marshal(toSerialize)
}

type NullableZoneModel struct {
	value *ZoneModel
	isSet bool
}

func (v NullableZoneModel) Get() *ZoneModel {
	return v.value
}

func (v *NullableZoneModel) Set(val *ZoneModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneModel(val *ZoneModel) *NullableZoneModel {
	return &NullableZoneModel{value: val, isSet: true}
}

func (v NullableZoneModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
