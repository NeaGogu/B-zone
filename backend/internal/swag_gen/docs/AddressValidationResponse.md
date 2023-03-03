# AddressValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** | Indicates if a location was found for the address and the minimum_certainty requirement was met. | [optional] 
**Certainty** | Pointer to **int32** | Certainty that latitude/longitude fields in address are accurate expressed as an integer between 0 and 100. When the valid field is false, certainty will always be set to 0 | [optional] 
**Address** | Pointer to [**AddressModel**](AddressModel.md) |  | [optional] 

## Methods

### NewAddressValidationResponse

`func NewAddressValidationResponse() *AddressValidationResponse`

NewAddressValidationResponse instantiates a new AddressValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressValidationResponseWithDefaults

`func NewAddressValidationResponseWithDefaults() *AddressValidationResponse`

NewAddressValidationResponseWithDefaults instantiates a new AddressValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *AddressValidationResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AddressValidationResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AddressValidationResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AddressValidationResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetCertainty

`func (o *AddressValidationResponse) GetCertainty() int32`

GetCertainty returns the Certainty field if non-nil, zero value otherwise.

### GetCertaintyOk

`func (o *AddressValidationResponse) GetCertaintyOk() (*int32, bool)`

GetCertaintyOk returns a tuple with the Certainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertainty

`func (o *AddressValidationResponse) SetCertainty(v int32)`

SetCertainty sets Certainty field to given value.

### HasCertainty

`func (o *AddressValidationResponse) HasCertainty() bool`

HasCertainty returns a boolean if a field has been set.

### GetAddress

`func (o *AddressValidationResponse) GetAddress() AddressModel`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressValidationResponse) GetAddressOk() (*AddressModel, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressValidationResponse) SetAddress(v AddressModel)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressValidationResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


