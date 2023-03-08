# ApiResponse403

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Error code | [optional] 
**Message** | Pointer to **string** | Message describing the error | [optional] 

## Methods

### NewApiResponse403

`func NewApiResponse403() *ApiResponse403`

NewApiResponse403 instantiates a new ApiResponse403 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponse403WithDefaults

`func NewApiResponse403WithDefaults() *ApiResponse403`

NewApiResponse403WithDefaults instantiates a new ApiResponse403 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiResponse403) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiResponse403) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiResponse403) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiResponse403) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ApiResponse403) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiResponse403) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiResponse403) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiResponse403) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


