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

// PartyFiltersModel struct for PartyFiltersModel
type PartyFiltersModel struct {
	// Unique Identifier
	Id *[]int32 `json:"id,omitempty"`
	// Type of this party
	PartyTypeName *[]string `json:"party_type_name,omitempty"`
	// PartyTypeID of this party. 2 = contractor, 3 = booking
	PartyTypeId *[]int32 `json:"party_type_id,omitempty"`
	// Name 1 for party
	Name1 *[]string `json:"name_1,omitempty"`
	// Name 2 for party
	Name2 *[]string `json:"name_2,omitempty"`
	// Number of this party
	Nr *[]string `json:"nr,omitempty"`
	// Contact person for party
	ContactPerson *[]string `json:"contact_person,omitempty"`
	// Url for party website
	Url *[]string `json:"url,omitempty"`
	// Activity Link ids
	Links *[]map[string]interface{} `json:"links,omitempty"`
	// Show updated since
	UpdatedAtSince *time.Time `json:"updated_at_since,omitempty"`
	// Show updated till
	UpdatedAtTill *time.Time `json:"updated_at_till,omitempty"`
}

// NewPartyFiltersModel instantiates a new PartyFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyFiltersModel() *PartyFiltersModel {
	this := PartyFiltersModel{}
	return &this
}

// NewPartyFiltersModelWithDefaults instantiates a new PartyFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyFiltersModelWithDefaults() *PartyFiltersModel {
	this := PartyFiltersModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetId() []int32 {
	if o == nil || o.Id == nil {
		var ret []int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetIdOk() (*[]int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *PartyFiltersModel) SetId(v []int32) {
	o.Id = &v
}

// GetPartyTypeName returns the PartyTypeName field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetPartyTypeName() []string {
	if o == nil || o.PartyTypeName == nil {
		var ret []string
		return ret
	}
	return *o.PartyTypeName
}

// GetPartyTypeNameOk returns a tuple with the PartyTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetPartyTypeNameOk() (*[]string, bool) {
	if o == nil || o.PartyTypeName == nil {
		return nil, false
	}
	return o.PartyTypeName, true
}

// HasPartyTypeName returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasPartyTypeName() bool {
	if o != nil && o.PartyTypeName != nil {
		return true
	}

	return false
}

// SetPartyTypeName gets a reference to the given []string and assigns it to the PartyTypeName field.
func (o *PartyFiltersModel) SetPartyTypeName(v []string) {
	o.PartyTypeName = &v
}

// GetPartyTypeId returns the PartyTypeId field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetPartyTypeId() []int32 {
	if o == nil || o.PartyTypeId == nil {
		var ret []int32
		return ret
	}
	return *o.PartyTypeId
}

// GetPartyTypeIdOk returns a tuple with the PartyTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetPartyTypeIdOk() (*[]int32, bool) {
	if o == nil || o.PartyTypeId == nil {
		return nil, false
	}
	return o.PartyTypeId, true
}

// HasPartyTypeId returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasPartyTypeId() bool {
	if o != nil && o.PartyTypeId != nil {
		return true
	}

	return false
}

// SetPartyTypeId gets a reference to the given []int32 and assigns it to the PartyTypeId field.
func (o *PartyFiltersModel) SetPartyTypeId(v []int32) {
	o.PartyTypeId = &v
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetName1() []string {
	if o == nil || o.Name1 == nil {
		var ret []string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetName1Ok() (*[]string, bool) {
	if o == nil || o.Name1 == nil {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasName1() bool {
	if o != nil && o.Name1 != nil {
		return true
	}

	return false
}

// SetName1 gets a reference to the given []string and assigns it to the Name1 field.
func (o *PartyFiltersModel) SetName1(v []string) {
	o.Name1 = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetName2() []string {
	if o == nil || o.Name2 == nil {
		var ret []string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetName2Ok() (*[]string, bool) {
	if o == nil || o.Name2 == nil {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasName2() bool {
	if o != nil && o.Name2 != nil {
		return true
	}

	return false
}

// SetName2 gets a reference to the given []string and assigns it to the Name2 field.
func (o *PartyFiltersModel) SetName2(v []string) {
	o.Name2 = &v
}

// GetNr returns the Nr field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetNr() []string {
	if o == nil || o.Nr == nil {
		var ret []string
		return ret
	}
	return *o.Nr
}

// GetNrOk returns a tuple with the Nr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetNrOk() (*[]string, bool) {
	if o == nil || o.Nr == nil {
		return nil, false
	}
	return o.Nr, true
}

// HasNr returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasNr() bool {
	if o != nil && o.Nr != nil {
		return true
	}

	return false
}

// SetNr gets a reference to the given []string and assigns it to the Nr field.
func (o *PartyFiltersModel) SetNr(v []string) {
	o.Nr = &v
}

// GetContactPerson returns the ContactPerson field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetContactPerson() []string {
	if o == nil || o.ContactPerson == nil {
		var ret []string
		return ret
	}
	return *o.ContactPerson
}

// GetContactPersonOk returns a tuple with the ContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetContactPersonOk() (*[]string, bool) {
	if o == nil || o.ContactPerson == nil {
		return nil, false
	}
	return o.ContactPerson, true
}

// HasContactPerson returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasContactPerson() bool {
	if o != nil && o.ContactPerson != nil {
		return true
	}

	return false
}

// SetContactPerson gets a reference to the given []string and assigns it to the ContactPerson field.
func (o *PartyFiltersModel) SetContactPerson(v []string) {
	o.ContactPerson = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetUrl() []string {
	if o == nil || o.Url == nil {
		var ret []string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetUrlOk() (*[]string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given []string and assigns it to the Url field.
func (o *PartyFiltersModel) SetUrl(v []string) {
	o.Url = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetLinks() []map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []map[string]interface{} and assigns it to the Links field.
func (o *PartyFiltersModel) SetLinks(v []map[string]interface{}) {
	o.Links = &v
}

// GetUpdatedAtSince returns the UpdatedAtSince field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetUpdatedAtSince() time.Time {
	if o == nil || o.UpdatedAtSince == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtSince
}

// GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtSince == nil {
		return nil, false
	}
	return o.UpdatedAtSince, true
}

// HasUpdatedAtSince returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasUpdatedAtSince() bool {
	if o != nil && o.UpdatedAtSince != nil {
		return true
	}

	return false
}

// SetUpdatedAtSince gets a reference to the given time.Time and assigns it to the UpdatedAtSince field.
func (o *PartyFiltersModel) SetUpdatedAtSince(v time.Time) {
	o.UpdatedAtSince = &v
}

// GetUpdatedAtTill returns the UpdatedAtTill field value if set, zero value otherwise.
func (o *PartyFiltersModel) GetUpdatedAtTill() time.Time {
	if o == nil || o.UpdatedAtTill == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtTill
}

// GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtTill == nil {
		return nil, false
	}
	return o.UpdatedAtTill, true
}

// HasUpdatedAtTill returns a boolean if a field has been set.
func (o *PartyFiltersModel) HasUpdatedAtTill() bool {
	if o != nil && o.UpdatedAtTill != nil {
		return true
	}

	return false
}

// SetUpdatedAtTill gets a reference to the given time.Time and assigns it to the UpdatedAtTill field.
func (o *PartyFiltersModel) SetUpdatedAtTill(v time.Time) {
	o.UpdatedAtTill = &v
}

func (o PartyFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PartyTypeName != nil {
		toSerialize["party_type_name"] = o.PartyTypeName
	}
	if o.PartyTypeId != nil {
		toSerialize["party_type_id"] = o.PartyTypeId
	}
	if o.Name1 != nil {
		toSerialize["name_1"] = o.Name1
	}
	if o.Name2 != nil {
		toSerialize["name_2"] = o.Name2
	}
	if o.Nr != nil {
		toSerialize["nr"] = o.Nr
	}
	if o.ContactPerson != nil {
		toSerialize["contact_person"] = o.ContactPerson
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.UpdatedAtSince != nil {
		toSerialize["updated_at_since"] = o.UpdatedAtSince
	}
	if o.UpdatedAtTill != nil {
		toSerialize["updated_at_till"] = o.UpdatedAtTill
	}
	return json.Marshal(toSerialize)
}

type NullablePartyFiltersModel struct {
	value *PartyFiltersModel
	isSet bool
}

func (v NullablePartyFiltersModel) Get() *PartyFiltersModel {
	return v.value
}

func (v *NullablePartyFiltersModel) Set(val *PartyFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyFiltersModel(val *PartyFiltersModel) *NullablePartyFiltersModel {
	return &NullablePartyFiltersModel{value: val, isSet: true}
}

func (v NullablePartyFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


