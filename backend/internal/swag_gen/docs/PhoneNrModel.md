# PhoneNrModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**AddressId** | Pointer to **int64** | Address ID asociated with this email address | [optional] 
**CountryCode** | Pointer to **string** | Phone number Country Code | [optional] 
**PhoneNrTypeName** | Pointer to **string** | Phone number type name | [optional] 
**Nr** | Pointer to **string** | Phone number | [optional] 
**Description** | Pointer to **string** | Phone number description | [optional] 
**Primary** | Pointer to **bool** | primary phone number | [optional] 

## Methods

### NewPhoneNrModel

`func NewPhoneNrModel() *PhoneNrModel`

NewPhoneNrModel instantiates a new PhoneNrModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNrModelWithDefaults

`func NewPhoneNrModelWithDefaults() *PhoneNrModel`

NewPhoneNrModelWithDefaults instantiates a new PhoneNrModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhoneNrModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneNrModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneNrModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneNrModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddressId

`func (o *PhoneNrModel) GetAddressId() int64`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *PhoneNrModel) GetAddressIdOk() (*int64, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *PhoneNrModel) SetAddressId(v int64)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *PhoneNrModel) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetCountryCode

`func (o *PhoneNrModel) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PhoneNrModel) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PhoneNrModel) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PhoneNrModel) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNrTypeName

`func (o *PhoneNrModel) GetPhoneNrTypeName() string`

GetPhoneNrTypeName returns the PhoneNrTypeName field if non-nil, zero value otherwise.

### GetPhoneNrTypeNameOk

`func (o *PhoneNrModel) GetPhoneNrTypeNameOk() (*string, bool)`

GetPhoneNrTypeNameOk returns a tuple with the PhoneNrTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNrTypeName

`func (o *PhoneNrModel) SetPhoneNrTypeName(v string)`

SetPhoneNrTypeName sets PhoneNrTypeName field to given value.

### HasPhoneNrTypeName

`func (o *PhoneNrModel) HasPhoneNrTypeName() bool`

HasPhoneNrTypeName returns a boolean if a field has been set.

### GetNr

`func (o *PhoneNrModel) GetNr() string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *PhoneNrModel) GetNrOk() (*string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *PhoneNrModel) SetNr(v string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *PhoneNrModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetDescription

`func (o *PhoneNrModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PhoneNrModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PhoneNrModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PhoneNrModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrimary

`func (o *PhoneNrModel) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PhoneNrModel) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PhoneNrModel) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *PhoneNrModel) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


