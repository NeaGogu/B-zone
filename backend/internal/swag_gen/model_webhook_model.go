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

// WebhookModel struct for WebhookModel
type WebhookModel struct {
	// objectId
	ObjectId int32 `json:"objectId"`
	// Name of this Web Hook
	WebHookName string `json:"webHookName"`
	// extra payload to be sent when the webhook is triggered
	ExtraPayload *[]PayloadItem `json:"extraPayload,omitempty"`
}

// NewWebhookModel instantiates a new WebhookModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookModel(objectId int32, webHookName string, ) *WebhookModel {
	this := WebhookModel{}
	this.ObjectId = objectId
	this.WebHookName = webHookName
	return &this
}

// NewWebhookModelWithDefaults instantiates a new WebhookModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookModelWithDefaults() *WebhookModel {
	this := WebhookModel{}
	return &this
}

// GetObjectId returns the ObjectId field value
func (o *WebhookModel) GetObjectId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *WebhookModel) GetObjectIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *WebhookModel) SetObjectId(v int32) {
	o.ObjectId = v
}

// GetWebHookName returns the WebHookName field value
func (o *WebhookModel) GetWebHookName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.WebHookName
}

// GetWebHookNameOk returns a tuple with the WebHookName field value
// and a boolean to check if the value has been set.
func (o *WebhookModel) GetWebHookNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebHookName, true
}

// SetWebHookName sets field value
func (o *WebhookModel) SetWebHookName(v string) {
	o.WebHookName = v
}

// GetExtraPayload returns the ExtraPayload field value if set, zero value otherwise.
func (o *WebhookModel) GetExtraPayload() []PayloadItem {
	if o == nil || o.ExtraPayload == nil {
		var ret []PayloadItem
		return ret
	}
	return *o.ExtraPayload
}

// GetExtraPayloadOk returns a tuple with the ExtraPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookModel) GetExtraPayloadOk() (*[]PayloadItem, bool) {
	if o == nil || o.ExtraPayload == nil {
		return nil, false
	}
	return o.ExtraPayload, true
}

// HasExtraPayload returns a boolean if a field has been set.
func (o *WebhookModel) HasExtraPayload() bool {
	if o != nil && o.ExtraPayload != nil {
		return true
	}

	return false
}

// SetExtraPayload gets a reference to the given []PayloadItem and assigns it to the ExtraPayload field.
func (o *WebhookModel) SetExtraPayload(v []PayloadItem) {
	o.ExtraPayload = &v
}

func (o WebhookModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objectId"] = o.ObjectId
	}
	if true {
		toSerialize["webHookName"] = o.WebHookName
	}
	if o.ExtraPayload != nil {
		toSerialize["extraPayload"] = o.ExtraPayload
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookModel struct {
	value *WebhookModel
	isSet bool
}

func (v NullableWebhookModel) Get() *WebhookModel {
	return v.value
}

func (v *NullableWebhookModel) Set(val *WebhookModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookModel(val *WebhookModel) *NullableWebhookModel {
	return &NullableWebhookModel{value: val, isSet: true}
}

func (v NullableWebhookModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

