package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDBModel struct {
	DB *mongo.Database
}

// UsersModel struct for UsersModel
type UsersModel struct {
	//
	Id *int64 `json:"id,omitempty"`
	// unique per user
	Uuid *string `json:"uuid,omitempty"`
}

// NewUsersModel instantiates a new UsersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersModel() *UsersModel {
	this := UsersModel{}
	return &this
}

// NewUsersModelWithDefaults instantiates a new UsersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersModelWithDefaults() *UsersModel {
	this := UsersModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UsersModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
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

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UsersModel) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
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

func (o UsersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableUsersModel struct {
	value *UsersModel
	isSet bool
}

func (v NullableUsersModel) Get() *UsersModel {
	return v.value
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
