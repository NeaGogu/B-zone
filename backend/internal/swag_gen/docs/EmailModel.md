# EmailModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**AddressId** | Pointer to **int64** | Address ID asociated with this email address | [optional] 
**Email** | Pointer to **string** | Email address | [optional] 
**Description** | Pointer to **string** | Email address description | [optional] 
**Primary** | Pointer to **bool** | primary Email address | [optional] 

## Methods

### NewEmailModel

`func NewEmailModel() *EmailModel`

NewEmailModel instantiates a new EmailModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailModelWithDefaults

`func NewEmailModelWithDefaults() *EmailModel`

NewEmailModelWithDefaults instantiates a new EmailModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EmailModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddressId

`func (o *EmailModel) GetAddressId() int64`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *EmailModel) GetAddressIdOk() (*int64, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *EmailModel) SetAddressId(v int64)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *EmailModel) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetEmail

`func (o *EmailModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *EmailModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrimary

`func (o *EmailModel) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *EmailModel) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *EmailModel) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *EmailModel) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


