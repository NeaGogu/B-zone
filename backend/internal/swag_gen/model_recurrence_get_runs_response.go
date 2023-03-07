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

// RecurrenceGetRunsResponse struct for RecurrenceGetRunsResponse
type RecurrenceGetRunsResponse struct {
	// success
	Success *bool `json:"success,omitempty"`
	// data
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewRecurrenceGetRunsResponse instantiates a new RecurrenceGetRunsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurrenceGetRunsResponse() *RecurrenceGetRunsResponse {
	this := RecurrenceGetRunsResponse{}
	return &this
}

// NewRecurrenceGetRunsResponseWithDefaults instantiates a new RecurrenceGetRunsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurrenceGetRunsResponseWithDefaults() *RecurrenceGetRunsResponse {
	this := RecurrenceGetRunsResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RecurrenceGetRunsResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceGetRunsResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RecurrenceGetRunsResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RecurrenceGetRunsResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RecurrenceGetRunsResponse) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceGetRunsResponse) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RecurrenceGetRunsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *RecurrenceGetRunsResponse) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o RecurrenceGetRunsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRecurrenceGetRunsResponse struct {
	value *RecurrenceGetRunsResponse
	isSet bool
}

func (v NullableRecurrenceGetRunsResponse) Get() *RecurrenceGetRunsResponse {
	return v.value
}

func (v *NullableRecurrenceGetRunsResponse) Set(val *RecurrenceGetRunsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurrenceGetRunsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurrenceGetRunsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurrenceGetRunsResponse(val *RecurrenceGetRunsResponse) *NullableRecurrenceGetRunsResponse {
	return &NullableRecurrenceGetRunsResponse{value: val, isSet: true}
}

func (v NullableRecurrenceGetRunsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurrenceGetRunsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

