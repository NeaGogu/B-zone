# ApiResponse45

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing the code | [optional] 
**Type** | Pointer to **string** | Ready | [optional] 
**Code** | Pointer to **float32** |  | [optional] 
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiResponse45

`func NewApiResponse45() *ApiResponse45`

NewApiResponse45 instantiates a new ApiResponse45 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponse45WithDefaults

`func NewApiResponse45WithDefaults() *ApiResponse45`

NewApiResponse45WithDefaults instantiates a new ApiResponse45 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiResponse45) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiResponse45) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiResponse45) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiResponse45) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *ApiResponse45) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiResponse45) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiResponse45) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiResponse45) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *ApiResponse45) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiResponse45) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiResponse45) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiResponse45) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAdditionalData

`func (o *ApiResponse45) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ApiResponse45) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ApiResponse45) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ApiResponse45) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


