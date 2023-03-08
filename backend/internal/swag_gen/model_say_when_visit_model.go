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
)

// SayWhenVisitModel struct for SayWhenVisitModel
type SayWhenVisitModel struct {
	// SayWhen Visit ID
	Id *string `json:"id,omitempty"`
	// 
	CurrentState *string `json:"current_state,omitempty"`
	// 
	Status *string `json:"status,omitempty"`
	// offered_by_owner
	OfferedByOwner *string `json:"offered_by_owner,omitempty"`
	// preferences_received
	PreferencesReceived *string `json:"preferences_received,omitempty"`
	// planned_by_partner
	PlannedByPartner *string `json:"planned_by_partner,omitempty"`
	// completed_by_partner
	CompletedByPartner *string `json:"completed_by_partner,omitempty"`
}

// NewSayWhenVisitModel instantiates a new SayWhenVisitModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSayWhenVisitModel() *SayWhenVisitModel {
	this := SayWhenVisitModel{}
	return &this
}

// NewSayWhenVisitModelWithDefaults instantiates a new SayWhenVisitModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSayWhenVisitModelWithDefaults() *SayWhenVisitModel {
	this := SayWhenVisitModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SayWhenVisitModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenVisitModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SayWhenVisitModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SayWhenVisitModel) SetId(v string) {
	o.Id = &v
}

// GetCurrentState returns the CurrentState field value if set, zero value otherwise.
func (o *SayWhenVisitModel) GetCurrentState() string {
	if o == nil || o.CurrentState == nil {
		var ret string
		return ret
	}
	return *o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenVisitModel) GetCurrentStateOk() (*string, bool) {
	if o == nil || o.CurrentState == nil {
		return nil, false
	}
	return o.CurrentState, true
}

// HasCurrentState returns a boolean if a field has been set.
func (o *SayWhenVisitModel) HasCurrentState() bool {
	if o != nil && o.CurrentState != nil {
		return true
	}

	return false
}

// SetCurrentState gets a reference to the given string and assigns it to the CurrentState field.
func (o *SayWhenVisitModel) SetCurrentState(v string) {
	o.CurrentState = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SayWhenVisitModel) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenVisitModel) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SayWhenVisitModel) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SayWhenVisitModel) SetStatus(v string) {
	o.Status = &v
}

// GetOfferedByOwner returns the OfferedByOwner field value if set, zero value otherwise.
func (o *SayWhenVisitModel) GetOfferedByOwner() string {
	if o == nil || o.OfferedByOwner == nil {
		var ret string
		return ret
	}
	return *o.OfferedByOwner
}

// GetOfferedByOwnerOk returns a tuple with the OfferedByOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenVisitModel) GetOfferedByOwnerOk() (*string, bool) {
	if o == nil || o.OfferedByOwner == nil {
		return nil, false
	}
	return o.OfferedByOwner, true
}

// HasOfferedByOwner returns a boolean if a field has been set.
func (o *SayWhenVisitModel) HasOfferedByOwner() bool {
	if o != nil && o.OfferedByOwner != nil {
		return true
	}

	return false
}

// SetOfferedByOwner gets a reference to the given string and assigns it to the OfferedByOwner field.
func (o *SayWhenVisitModel) SetOfferedByOwner(v string) {
	o.OfferedByOwner = &v
}

// GetPreferencesReceived returns the PreferencesReceived field value if set, zero value otherwise.
func (o *SayWhenVisitModel) GetPreferencesReceived() string {
	if o == nil || o.PreferencesReceived == nil {
		var ret string
		return ret
	}
	return *o.PreferencesReceived
}

// GetPreferencesReceivedOk returns a tuple with the PreferencesReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenVisitModel) GetPreferencesReceivedOk() (*string, bool) {
	if o == nil || o.PreferencesReceived == nil {
		return nil, false
	}
	return o.PreferencesReceived, true
}

// HasPreferencesReceived returns a boolean if a field has been set.
func (o *SayWhenVisitModel) HasPreferencesReceived() bool {
	if o != nil && o.PreferencesReceived != nil {
		return true
	}

	return false
}

// SetPreferencesReceived gets a reference to the given string and assigns it to the PreferencesReceived field.
func (o *SayWhenVisitModel) SetPreferencesReceived(v string) {
	o.PreferencesReceived = &v
}

// GetPlannedByPartner returns the PlannedByPartner field value if set, zero value otherwise.
func (o *SayWhenVisitModel) GetPlannedByPartner() string {
	if o == nil || o.PlannedByPartner == nil {
		var ret string
		return ret
	}
	return *o.PlannedByPartner
}

// GetPlannedByPartnerOk returns a tuple with the PlannedByPartner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenVisitModel) GetPlannedByPartnerOk() (*string, bool) {
	if o == nil || o.PlannedByPartner == nil {
		return nil, false
	}
	return o.PlannedByPartner, true
}

// HasPlannedByPartner returns a boolean if a field has been set.
func (o *SayWhenVisitModel) HasPlannedByPartner() bool {
	if o != nil && o.PlannedByPartner != nil {
		return true
	}

	return false
}

// SetPlannedByPartner gets a reference to the given string and assigns it to the PlannedByPartner field.
func (o *SayWhenVisitModel) SetPlannedByPartner(v string) {
	o.PlannedByPartner = &v
}

// GetCompletedByPartner returns the CompletedByPartner field value if set, zero value otherwise.
func (o *SayWhenVisitModel) GetCompletedByPartner() string {
	if o == nil || o.CompletedByPartner == nil {
		var ret string
		return ret
	}
	return *o.CompletedByPartner
}

// GetCompletedByPartnerOk returns a tuple with the CompletedByPartner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenVisitModel) GetCompletedByPartnerOk() (*string, bool) {
	if o == nil || o.CompletedByPartner == nil {
		return nil, false
	}
	return o.CompletedByPartner, true
}

// HasCompletedByPartner returns a boolean if a field has been set.
func (o *SayWhenVisitModel) HasCompletedByPartner() bool {
	if o != nil && o.CompletedByPartner != nil {
		return true
	}

	return false
}

// SetCompletedByPartner gets a reference to the given string and assigns it to the CompletedByPartner field.
func (o *SayWhenVisitModel) SetCompletedByPartner(v string) {
	o.CompletedByPartner = &v
}

func (o SayWhenVisitModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CurrentState != nil {
		toSerialize["current_state"] = o.CurrentState
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.OfferedByOwner != nil {
		toSerialize["offered_by_owner"] = o.OfferedByOwner
	}
	if o.PreferencesReceived != nil {
		toSerialize["preferences_received"] = o.PreferencesReceived
	}
	if o.PlannedByPartner != nil {
		toSerialize["planned_by_partner"] = o.PlannedByPartner
	}
	if o.CompletedByPartner != nil {
		toSerialize["completed_by_partner"] = o.CompletedByPartner
	}
	return json.Marshal(toSerialize)
}

type NullableSayWhenVisitModel struct {
	value *SayWhenVisitModel
	isSet bool
}

func (v NullableSayWhenVisitModel) Get() *SayWhenVisitModel {
	return v.value
}

func (v *NullableSayWhenVisitModel) Set(val *SayWhenVisitModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSayWhenVisitModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSayWhenVisitModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSayWhenVisitModel(val *SayWhenVisitModel) *NullableSayWhenVisitModel {
	return &NullableSayWhenVisitModel{value: val, isSet: true}
}

func (v NullableSayWhenVisitModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSayWhenVisitModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


