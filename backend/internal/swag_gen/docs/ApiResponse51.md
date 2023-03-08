# ApiResponse51

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing the error | [optional] 
**Code** | Pointer to **float32** | Error code | [optional] 

## Methods

### NewApiResponse51

`func NewApiResponse51() *ApiResponse51`

NewApiResponse51 instantiates a new ApiResponse51 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponse51WithDefaults

`func NewApiResponse51WithDefaults() *ApiResponse51`

NewApiResponse51WithDefaults instantiates a new ApiResponse51 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiResponse51) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiResponse51) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiResponse51) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiResponse51) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *ApiResponse51) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiResponse51) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiResponse51) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiResponse51) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


