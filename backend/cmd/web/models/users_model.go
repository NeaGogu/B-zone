package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersDBModel struct {
	DB *mongo.Database
}

// UsersModel struct for UsersModel
type UsersModel struct {
	//
	Id *int64 `json:"id,omitempty"`
	// unique per user
	Uuid *string `json:"uuid,omitempty"`
	// zones plots per user
	ZonesPlotIds *[]string `json:"user_zones_plot_ids,omitempty"`
}

// NewUsersModel instantiates a new UsersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersModel() *UsersModel {
	this := UsersModel{}
	return &this
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UsersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UsersModel) SetId(v int64) {
	o.Id = &v
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersModel) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UsersModel) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UsersModel) SetUuid(v string) {
	o.Uuid = &v
}

// GetZonesPlotsOk returns a tuple with the ZoneRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersModel) GetZonesPlotsOk() (*[]string, bool) {
	if o == nil || o.ZonesPlotIds == nil {
		return nil, false
	}
	return o.ZonesPlotIds, true
}

// HasZonesPlots returns a boolean if a field has been set.
func (o *UsersModel) HasZonesPlots() bool {
	if o != nil && o.ZonesPlotIds != nil {
		return true
	}

	return false
}

// SetZonesPlots SetUserZonesPlot gets a reference to the given []ZonesPlotModel and assigns it to the ZonesPlots field.
func (o *UsersModel) SetZonesPlots(v []string) {
	o.ZonesPlotIds = &v
}

func (o UsersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ZonesPlotIds != nil {
		toSerialize["user_zones_plot_ids"] = o.ZonesPlotIds
	}
	return json.Marshal(toSerialize)
}

type NullableUsersModel struct {
	value *UsersModel
	isSet bool
}

func (v *NullableUsersModel) Set(val *UsersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersModel(val *UsersModel) *NullableUsersModel {
	return &NullableUsersModel{value: val, isSet: true}
}

func (v NullableUsersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
