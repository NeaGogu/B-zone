# TagTypeRetrieveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**AdditionalData** | Pointer to [**TagTypeListResponse**](TagTypeListResponse.md) |  | [optional] 

## Methods

### NewTagTypeRetrieveResponse

`func NewTagTypeRetrieveResponse() *TagTypeRetrieveResponse`

NewTagTypeRetrieveResponse instantiates a new TagTypeRetrieveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagTypeRetrieveResponseWithDefaults

`func NewTagTypeRetrieveResponseWithDefaults() *TagTypeRetrieveResponse`

NewTagTypeRetrieveResponseWithDefaults instantiates a new TagTypeRetrieveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TagTypeRetrieveResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagTypeRetrieveResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagTypeRetrieveResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *TagTypeRetrieveResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *TagTypeRetrieveResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagTypeRetrieveResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagTypeRetrieveResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TagTypeRetrieveResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *TagTypeRetrieveResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TagTypeRetrieveResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TagTypeRetrieveResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TagTypeRetrieveResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAdditionalData

`func (o *TagTypeRetrieveResponse) GetAdditionalData() TagTypeListResponse`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *TagTypeRetrieveResponse) GetAdditionalDataOk() (*TagTypeListResponse, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *TagTypeRetrieveResponse) SetAdditionalData(v TagTypeListResponse)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *TagTypeRetrieveResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


