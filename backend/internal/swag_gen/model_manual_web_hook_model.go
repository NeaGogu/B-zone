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

// ManualWebHookModel struct for ManualWebHookModel
type ManualWebHookModel struct {
	// Unique Identifier
	ObjectId *int64 `json:"object_id,omitempty"`
	// Name of this Web Hook
	WebHookName *string `json:"web_hook_name,omitempty"`
	// extra payload to be sent when the webhook is triggered
	ExtraPayload *[]PayloadItem `json:"extra_payload,omitempty"`
}

// NewManualWebHookModel instantiates a new ManualWebHookModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualWebHookModel() *ManualWebHookModel {
	this := ManualWebHookModel{}
	return &this
}

// NewManualWebHookModelWithDefaults instantiates a new ManualWebHookModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualWebHookModelWithDefaults() *ManualWebHookModel {
	this := ManualWebHookModel{}
	return &this
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *ManualWebHookModel) GetObjectId() int64 {
	if o == nil || o.ObjectId == nil {
		var ret int64
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualWebHookModel) GetObjectIdOk() (*int64, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *ManualWebHookModel) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given int64 and assigns it to the ObjectId field.
func (o *ManualWebHookModel) SetObjectId(v int64) {
	o.ObjectId = &v
}

// GetWebHookName returns the WebHookName field value if set, zero value otherwise.
func (o *ManualWebHookModel) GetWebHookName() string {
	if o == nil || o.WebHookName == nil {
		var ret string
		return ret
	}
	return *o.WebHookName
}

// GetWebHookNameOk returns a tuple with the WebHookName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualWebHookModel) GetWebHookNameOk() (*string, bool) {
	if o == nil || o.WebHookName == nil {
		return nil, false
	}
	return o.WebHookName, true
}

// HasWebHookName returns a boolean if a field has been set.
func (o *ManualWebHookModel) HasWebHookName() bool {
	if o != nil && o.WebHookName != nil {
		return true
	}

	return false
}

// SetWebHookName gets a reference to the given string and assigns it to the WebHookName field.
func (o *ManualWebHookModel) SetWebHookName(v string) {
	o.WebHookName = &v
}

// GetExtraPayload returns the ExtraPayload field value if set, zero value otherwise.
func (o *ManualWebHookModel) GetExtraPayload() []PayloadItem {
	if o == nil || o.ExtraPayload == nil {
		var ret []PayloadItem
		return ret
	}
	return *o.ExtraPayload
}

// GetExtraPayloadOk returns a tuple with the ExtraPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualWebHookModel) GetExtraPayloadOk() (*[]PayloadItem, bool) {
	if o == nil || o.ExtraPayload == nil {
		return nil, false
	}
	return o.ExtraPayload, true
}

// HasExtraPayload returns a boolean if a field has been set.
func (o *ManualWebHookModel) HasExtraPayload() bool {
	if o != nil && o.ExtraPayload != nil {
		return true
	}

	return false
}

// SetExtraPayload gets a reference to the given []PayloadItem and assigns it to the ExtraPayload field.
func (o *ManualWebHookModel) SetExtraPayload(v []PayloadItem) {
	o.ExtraPayload = &v
}

func (o ManualWebHookModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectId != nil {
		toSerialize["object_id"] = o.ObjectId
	}
	if o.WebHookName != nil {
		toSerialize["web_hook_name"] = o.WebHookName
	}
	if o.ExtraPayload != nil {
		toSerialize["extra_payload"] = o.ExtraPayload
	}
	return json.Marshal(toSerialize)
}

type NullableManualWebHookModel struct {
	value *ManualWebHookModel
	isSet bool
}

func (v NullableManualWebHookModel) Get() *ManualWebHookModel {
	return v.value
}

func (v *NullableManualWebHookModel) Set(val *ManualWebHookModel) {
	v.value = val
	v.isSet = true
}

func (v NullableManualWebHookModel) IsSet() bool {
	return v.isSet
}

func (v *NullableManualWebHookModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualWebHookModel(val *ManualWebHookModel) *NullableManualWebHookModel {
	return &NullableManualWebHookModel{value: val, isSet: true}
}

func (v NullableManualWebHookModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualWebHookModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

