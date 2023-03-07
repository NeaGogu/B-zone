package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ZonesPlotDBModel struct {
	DB *mongo.Database
}

// ZonesPlotModel struct for ZonesPLotModel
type ZonesPlotModel struct {
	// Unique Plot ID
	Id *int64 `json:"id,omitempty"`
	// Plot Name
	Name *string `json:"name,omitempty"`
	//Zones in the plot
	Zones *[]ZoneModel `json:"zones_plot,omitempty"`
	//created_at date time
	ZonesPlotCreatedAt *time.Time `json:"zones_plot_created_at,omitempty"`
	//saved_at date time
	ZonesPlotSavedAt *time.Time `json:"zones_plot_saved_at,omitempty"`
}

// NewZonesPlotModel instantiates a new ZonesPlotModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZonesPlotModel() *ZonesPlotModel {
	this := ZonesPlotModel{}
	return &this
}

// NewZonesPlotModelWithDefaults instantiates a new ZonesPlotModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZonesPlotModelWithDefaults() *ZonesPlotModel {
	this := ZonesPlotModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZonesPlotModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesPlotModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZonesPlotModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ZonesPlotModel) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZonesPlotModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesPlotModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZonesPlotModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZonesPlotModel) SetName(v string) {
	o.Name = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *ZonesPlotModel) GetZones() []ZoneModel {
	if o == nil || o.Zones == nil {
		var ret []ZoneModel
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the ZoneRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesPlotModel) GetZonesOk() (*[]ZoneModel, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *ZonesPlotModel) HasZones() bool {
	if o != nil && o.Zones != nil {
		return true
	}

	return false
}

// SetZones gets a reference to the given []ZonesPlotModel and assigns it to the ZoneRanges field.
func (o *ZonesPlotModel) SetZones(v []ZoneModel) {
	o.Zones = &v
}

// GetZonesPlotCreatedAt returns the ZoneCreatedAt field value if set, zero value otherwise.
func (o *ZonesPlotModel) GetZonesPlotCreatedAt() time.Time {
	if o == nil || o.ZonesPlotCreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ZonesPlotCreatedAt
}

// GetZonesPlotCreatedAtOk returns a tuple with the ZoneCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesPlotModel) GetZonesPlotCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.ZonesPlotCreatedAt == nil {
		return nil, false
	}
	return o.ZonesPlotCreatedAt, true
}

// HasZonesPlotCreatedAt returns a boolean if a field has been set.
func (o *ZonesPlotModel) HasZonesPlotCreatedAt() bool {
	if o != nil && o.ZonesPlotCreatedAt != nil {
		return true
	}

	return false
}

// SetZonesPlotCreatedAt gets a reference to the given time.Time and assigns it to the ZoneCreatedAt field.
func (o *ZonesPlotModel) SetZonesPlotCreatedAt(v time.Time) {
	o.ZonesPlotCreatedAt = &v
}

// GetZonesPlotSavedAt returns the ZoneCreatedAt field value if set, zero value otherwise.
func (o *ZonesPlotModel) GetZonesPlotSavedAt() time.Time {
	if o == nil || o.ZonesPlotSavedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ZonesPlotSavedAt
}

// GetZonesPlotSavedAtOk returns a tuple with the ZoneCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZonesPlotModel) GetZonesPlotSavedAtOk() (*time.Time, bool) {
	if o == nil || o.ZonesPlotSavedAt == nil {
		return nil, false
	}
	return o.ZonesPlotSavedAt, true
}

// HasZonesPlotSavedAt returns a boolean if a field has been set.
func (o *ZonesPlotModel) HasZonesPlotSavedAt() bool {
	if o != nil && o.ZonesPlotSavedAt != nil {
		return true
	}

	return false
}

// SetZonesPlotSavedAt gets a reference to the given time.Time and assigns it to the ZoneCreatedAt field.
func (o *ZonesPlotModel) SetZonesPlotSavedAt(v time.Time) {
	o.ZonesPlotSavedAt = &v
}

func (o ZonesPlotModel) MarshalJSON() ([]byte, error) {

	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Zones != nil {
		toSerialize["zones_plot"] = o.Zones
	}
	if o.ZonesPlotCreatedAt != nil {
		toSerialize["zones_plot_created_at"] = o.ZonesPlotCreatedAt
	}
	if o.ZonesPlotSavedAt != nil {
		toSerialize["zones_plot_saved_at"] = o.ZonesPlotSavedAt
	}
	return json.Marshal(toSerialize)
}

type NullableZonesPlotModel struct {
	value *ZonesPlotModel
	isSet bool
}

func (v NullableZonesPlotModel) Get() *ZonesPlotModel {
	return v.value
}

func (v *NullableZonesPlotModel) Set(val *ZonesPlotModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZonesPlotModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZonesPlotModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZonesPlotModel(val *ZonesPlotModel) *NullableZonesPlotModel {
	return &NullableZonesPlotModel{value: val, isSet: true}
}

func (v NullableZonesPlotModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZonesPlotModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
