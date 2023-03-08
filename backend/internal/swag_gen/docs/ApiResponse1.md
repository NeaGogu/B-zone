# ApiResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing the code | [optional] 
**Type** | Pointer to **string** | Ready | [optional] 
**Code** | Pointer to **float32** |  | [optional] 
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiResponse1

`func NewApiResponse1() *ApiResponse1`

NewApiResponse1 instantiates a new ApiResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponse1WithDefaults

`func NewApiResponse1WithDefaults() *ApiResponse1`

NewApiResponse1WithDefaults instantiates a new ApiResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiResponse1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiResponse1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiResponse1) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiResponse1) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *ApiResponse1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiResponse1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiResponse1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiResponse1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *ApiResponse1) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiResponse1) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiResponse1) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiResponse1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAdditionalData

`func (o *ApiResponse1) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ApiResponse1) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ApiResponse1) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ApiResponse1) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


