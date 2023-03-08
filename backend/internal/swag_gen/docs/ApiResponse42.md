# ApiResponse42

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing the code | [optional] 
**Type** | Pointer to **string** | Ready | [optional] 
**Code** | Pointer to **float32** |  | [optional] 
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiResponse42

`func NewApiResponse42() *ApiResponse42`

NewApiResponse42 instantiates a new ApiResponse42 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponse42WithDefaults

`func NewApiResponse42WithDefaults() *ApiResponse42`

NewApiResponse42WithDefaults instantiates a new ApiResponse42 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiResponse42) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiResponse42) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiResponse42) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiResponse42) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *ApiResponse42) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiResponse42) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiResponse42) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiResponse42) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *ApiResponse42) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiResponse42) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiResponse42) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiResponse42) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAdditionalData

`func (o *ApiResponse42) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ApiResponse42) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ApiResponse42) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ApiResponse42) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


