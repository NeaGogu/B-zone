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

// TransactionTokenHasFailedResponse struct for TransactionTokenHasFailedResponse
type TransactionTokenHasFailedResponse struct {
	// Is 1 on success (transaction id is found en transaction_type is online and amount is correct and is not yet paid), 0 when failed
	Success *int32 `json:"success,omitempty"`
}

// NewTransactionTokenHasFailedResponse instantiates a new TransactionTokenHasFailedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionTokenHasFailedResponse() *TransactionTokenHasFailedResponse {
	this := TransactionTokenHasFailedResponse{}
	return &this
}

// NewTransactionTokenHasFailedResponseWithDefaults instantiates a new TransactionTokenHasFailedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionTokenHasFailedResponseWithDefaults() *TransactionTokenHasFailedResponse {
	this := TransactionTokenHasFailedResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *TransactionTokenHasFailedResponse) GetSuccess() int32 {
	if o == nil || o.Success == nil {
		var ret int32
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionTokenHasFailedResponse) GetSuccessOk() (*int32, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *TransactionTokenHasFailedResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given int32 and assigns it to the Success field.
func (o *TransactionTokenHasFailedResponse) SetSuccess(v int32) {
	o.Success = &v
}

func (o TransactionTokenHasFailedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionTokenHasFailedResponse struct {
	value *TransactionTokenHasFailedResponse
	isSet bool
}

func (v NullableTransactionTokenHasFailedResponse) Get() *TransactionTokenHasFailedResponse {
	return v.value
}

func (v *NullableTransactionTokenHasFailedResponse) Set(val *TransactionTokenHasFailedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTokenHasFailedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTokenHasFailedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTokenHasFailedResponse(val *TransactionTokenHasFailedResponse) *NullableTransactionTokenHasFailedResponse {
	return &NullableTransactionTokenHasFailedResponse{value: val, isSet: true}
}

func (v NullableTransactionTokenHasFailedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTokenHasFailedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


