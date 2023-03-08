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

// NotificationCategoryListResponse struct for NotificationCategoryListResponse
type NotificationCategoryListResponse struct {
	// 
	Items *[]NotificationCategoryModel `json:"items,omitempty"`
	// Count of total items with filters in place
	CountFiltered *int32 `json:"count_filtered,omitempty"`
	// Count of total items without filters in place
	CountUnfiltered *int32 `json:"count_unfiltered,omitempty"`
	// Count of items with limit in place
	CountLimited *int32 `json:"count_limited,omitempty"`
}

// NewNotificationCategoryListResponse instantiates a new NotificationCategoryListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationCategoryListResponse() *NotificationCategoryListResponse {
	this := NotificationCategoryListResponse{}
	return &this
}

// NewNotificationCategoryListResponseWithDefaults instantiates a new NotificationCategoryListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationCategoryListResponseWithDefaults() *NotificationCategoryListResponse {
	this := NotificationCategoryListResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *NotificationCategoryListResponse) GetItems() []NotificationCategoryModel {
	if o == nil || o.Items == nil {
		var ret []NotificationCategoryModel
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationCategoryListResponse) GetItemsOk() (*[]NotificationCategoryModel, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *NotificationCategoryListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []NotificationCategoryModel and assigns it to the Items field.
func (o *NotificationCategoryListResponse) SetItems(v []NotificationCategoryModel) {
	o.Items = &v
}

// GetCountFiltered returns the CountFiltered field value if set, zero value otherwise.
func (o *NotificationCategoryListResponse) GetCountFiltered() int32 {
	if o == nil || o.CountFiltered == nil {
		var ret int32
		return ret
	}
	return *o.CountFiltered
}

// GetCountFilteredOk returns a tuple with the CountFiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationCategoryListResponse) GetCountFilteredOk() (*int32, bool) {
	if o == nil || o.CountFiltered == nil {
		return nil, false
	}
	return o.CountFiltered, true
}

// HasCountFiltered returns a boolean if a field has been set.
func (o *NotificationCategoryListResponse) HasCountFiltered() bool {
	if o != nil && o.CountFiltered != nil {
		return true
	}

	return false
}

// SetCountFiltered gets a reference to the given int32 and assigns it to the CountFiltered field.
func (o *NotificationCategoryListResponse) SetCountFiltered(v int32) {
	o.CountFiltered = &v
}

// GetCountUnfiltered returns the CountUnfiltered field value if set, zero value otherwise.
func (o *NotificationCategoryListResponse) GetCountUnfiltered() int32 {
	if o == nil || o.CountUnfiltered == nil {
		var ret int32
		return ret
	}
	return *o.CountUnfiltered
}

// GetCountUnfilteredOk returns a tuple with the CountUnfiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationCategoryListResponse) GetCountUnfilteredOk() (*int32, bool) {
	if o == nil || o.CountUnfiltered == nil {
		return nil, false
	}
	return o.CountUnfiltered, true
}

// HasCountUnfiltered returns a boolean if a field has been set.
func (o *NotificationCategoryListResponse) HasCountUnfiltered() bool {
	if o != nil && o.CountUnfiltered != nil {
		return true
	}

	return false
}

// SetCountUnfiltered gets a reference to the given int32 and assigns it to the CountUnfiltered field.
func (o *NotificationCategoryListResponse) SetCountUnfiltered(v int32) {
	o.CountUnfiltered = &v
}

// GetCountLimited returns the CountLimited field value if set, zero value otherwise.
func (o *NotificationCategoryListResponse) GetCountLimited() int32 {
	if o == nil || o.CountLimited == nil {
		var ret int32
		return ret
	}
	return *o.CountLimited
}

// GetCountLimitedOk returns a tuple with the CountLimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationCategoryListResponse) GetCountLimitedOk() (*int32, bool) {
	if o == nil || o.CountLimited == nil {
		return nil, false
	}
	return o.CountLimited, true
}

// HasCountLimited returns a boolean if a field has been set.
func (o *NotificationCategoryListResponse) HasCountLimited() bool {
	if o != nil && o.CountLimited != nil {
		return true
	}

	return false
}

// SetCountLimited gets a reference to the given int32 and assigns it to the CountLimited field.
func (o *NotificationCategoryListResponse) SetCountLimited(v int32) {
	o.CountLimited = &v
}

func (o NotificationCategoryListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.CountFiltered != nil {
		toSerialize["count_filtered"] = o.CountFiltered
	}
	if o.CountUnfiltered != nil {
		toSerialize["count_unfiltered"] = o.CountUnfiltered
	}
	if o.CountLimited != nil {
		toSerialize["count_limited"] = o.CountLimited
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationCategoryListResponse struct {
	value *NotificationCategoryListResponse
	isSet bool
}

func (v NullableNotificationCategoryListResponse) Get() *NotificationCategoryListResponse {
	return v.value
}

func (v *NullableNotificationCategoryListResponse) Set(val *NotificationCategoryListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationCategoryListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationCategoryListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationCategoryListResponse(val *NotificationCategoryListResponse) *NullableNotificationCategoryListResponse {
	return &NullableNotificationCategoryListResponse{value: val, isSet: true}
}

func (v NullableNotificationCategoryListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationCategoryListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


