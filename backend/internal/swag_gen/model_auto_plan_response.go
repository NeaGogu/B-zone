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

// AutoPlanResponse struct for AutoPlanResponse
type AutoPlanResponse struct {
	// token for the auto plan job
	Token *string `json:"token,omitempty"`
	// current status for request
	Status *string `json:"status,omitempty"`
	// 
	AffectedActivities *[]ActivityModel `json:"affected_activities,omitempty"`
	// 
	UnavailableTimewindows *[]TimeSlotModel `json:"unavailable_timewindows,omitempty"`
	// 
	LatestAnalyzedDate *string `json:"latest_analyzed_date,omitempty"`
}

// NewAutoPlanResponse instantiates a new AutoPlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoPlanResponse() *AutoPlanResponse {
	this := AutoPlanResponse{}
	return &this
}

// NewAutoPlanResponseWithDefaults instantiates a new AutoPlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoPlanResponseWithDefaults() *AutoPlanResponse {
	this := AutoPlanResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AutoPlanResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPlanResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AutoPlanResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AutoPlanResponse) SetToken(v string) {
	o.Token = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AutoPlanResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPlanResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AutoPlanResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AutoPlanResponse) SetStatus(v string) {
	o.Status = &v
}

// GetAffectedActivities returns the AffectedActivities field value if set, zero value otherwise.
func (o *AutoPlanResponse) GetAffectedActivities() []ActivityModel {
	if o == nil || o.AffectedActivities == nil {
		var ret []ActivityModel
		return ret
	}
	return *o.AffectedActivities
}

// GetAffectedActivitiesOk returns a tuple with the AffectedActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPlanResponse) GetAffectedActivitiesOk() (*[]ActivityModel, bool) {
	if o == nil || o.AffectedActivities == nil {
		return nil, false
	}
	return o.AffectedActivities, true
}

// HasAffectedActivities returns a boolean if a field has been set.
func (o *AutoPlanResponse) HasAffectedActivities() bool {
	if o != nil && o.AffectedActivities != nil {
		return true
	}

	return false
}

// SetAffectedActivities gets a reference to the given []ActivityModel and assigns it to the AffectedActivities field.
func (o *AutoPlanResponse) SetAffectedActivities(v []ActivityModel) {
	o.AffectedActivities = &v
}

// GetUnavailableTimewindows returns the UnavailableTimewindows field value if set, zero value otherwise.
func (o *AutoPlanResponse) GetUnavailableTimewindows() []TimeSlotModel {
	if o == nil || o.UnavailableTimewindows == nil {
		var ret []TimeSlotModel
		return ret
	}
	return *o.UnavailableTimewindows
}

// GetUnavailableTimewindowsOk returns a tuple with the UnavailableTimewindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPlanResponse) GetUnavailableTimewindowsOk() (*[]TimeSlotModel, bool) {
	if o == nil || o.UnavailableTimewindows == nil {
		return nil, false
	}
	return o.UnavailableTimewindows, true
}

// HasUnavailableTimewindows returns a boolean if a field has been set.
func (o *AutoPlanResponse) HasUnavailableTimewindows() bool {
	if o != nil && o.UnavailableTimewindows != nil {
		return true
	}

	return false
}

// SetUnavailableTimewindows gets a reference to the given []TimeSlotModel and assigns it to the UnavailableTimewindows field.
func (o *AutoPlanResponse) SetUnavailableTimewindows(v []TimeSlotModel) {
	o.UnavailableTimewindows = &v
}

// GetLatestAnalyzedDate returns the LatestAnalyzedDate field value if set, zero value otherwise.
func (o *AutoPlanResponse) GetLatestAnalyzedDate() string {
	if o == nil || o.LatestAnalyzedDate == nil {
		var ret string
		return ret
	}
	return *o.LatestAnalyzedDate
}

// GetLatestAnalyzedDateOk returns a tuple with the LatestAnalyzedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPlanResponse) GetLatestAnalyzedDateOk() (*string, bool) {
	if o == nil || o.LatestAnalyzedDate == nil {
		return nil, false
	}
	return o.LatestAnalyzedDate, true
}

// HasLatestAnalyzedDate returns a boolean if a field has been set.
func (o *AutoPlanResponse) HasLatestAnalyzedDate() bool {
	if o != nil && o.LatestAnalyzedDate != nil {
		return true
	}

	return false
}

// SetLatestAnalyzedDate gets a reference to the given string and assigns it to the LatestAnalyzedDate field.
func (o *AutoPlanResponse) SetLatestAnalyzedDate(v string) {
	o.LatestAnalyzedDate = &v
}

func (o AutoPlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.AffectedActivities != nil {
		toSerialize["affected_activities"] = o.AffectedActivities
	}
	if o.UnavailableTimewindows != nil {
		toSerialize["unavailable_timewindows"] = o.UnavailableTimewindows
	}
	if o.LatestAnalyzedDate != nil {
		toSerialize["latest_analyzed_date"] = o.LatestAnalyzedDate
	}
	return json.Marshal(toSerialize)
}

type NullableAutoPlanResponse struct {
	value *AutoPlanResponse
	isSet bool
}

func (v NullableAutoPlanResponse) Get() *AutoPlanResponse {
	return v.value
}

func (v *NullableAutoPlanResponse) Set(val *AutoPlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoPlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoPlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoPlanResponse(val *AutoPlanResponse) *NullableAutoPlanResponse {
	return &NullableAutoPlanResponse{value: val, isSet: true}
}

func (v NullableAutoPlanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoPlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

