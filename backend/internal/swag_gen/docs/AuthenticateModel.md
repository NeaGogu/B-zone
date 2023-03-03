# AuthenticateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | your access token. If request was made with &#39;return_jwt&#39;, then this token will be a JWT token. Use this token in a header Authorization:Bearer &lt;jwt goes here&gt;. This token currently has an expiration time of 8 hours, but will also be invalidated when signing out | [optional] [readonly] 
**User** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 

## Methods

### NewAuthenticateModel

`func NewAuthenticateModel() *AuthenticateModel`

NewAuthenticateModel instantiates a new AuthenticateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticateModelWithDefaults

`func NewAuthenticateModelWithDefaults() *AuthenticateModel`

NewAuthenticateModelWithDefaults instantiates a new AuthenticateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AuthenticateModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthenticateModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthenticateModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthenticateModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUser

`func (o *AuthenticateModel) GetUser() UsersModel`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthenticateModel) GetUserOk() (*UsersModel, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthenticateModel) SetUser(v UsersModel)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthenticateModel) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


