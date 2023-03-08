# CredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**ReturnJwt** | Pointer to **bool** | Set to true for requesting a jwt token | [optional] 

## Methods

### NewCredentialsModel

`func NewCredentialsModel() *CredentialsModel`

NewCredentialsModel instantiates a new CredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsModelWithDefaults

`func NewCredentialsModelWithDefaults() *CredentialsModel`

NewCredentialsModelWithDefaults instantiates a new CredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CredentialsModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CredentialsModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CredentialsModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CredentialsModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *CredentialsModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialsModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialsModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CredentialsModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetReturnJwt

`func (o *CredentialsModel) GetReturnJwt() bool`

GetReturnJwt returns the ReturnJwt field if non-nil, zero value otherwise.

### GetReturnJwtOk

`func (o *CredentialsModel) GetReturnJwtOk() (*bool, bool)`

GetReturnJwtOk returns a tuple with the ReturnJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnJwt

`func (o *CredentialsModel) SetReturnJwt(v bool)`

SetReturnJwt sets ReturnJwt field to given value.

### HasReturnJwt

`func (o *CredentialsModel) HasReturnJwt() bool`

HasReturnJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


