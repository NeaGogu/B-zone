# PaymentCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** | Ready | [optional] 
**Message** | Pointer to **string** | Message describing the code | [optional] 
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPaymentCreateResponse

`func NewPaymentCreateResponse() *PaymentCreateResponse`

NewPaymentCreateResponse instantiates a new PaymentCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCreateResponseWithDefaults

`func NewPaymentCreateResponseWithDefaults() *PaymentCreateResponse`

NewPaymentCreateResponseWithDefaults instantiates a new PaymentCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PaymentCreateResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentCreateResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentCreateResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *PaymentCreateResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *PaymentCreateResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentCreateResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentCreateResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentCreateResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentCreateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentCreateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentCreateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaymentCreateResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAdditionalData

`func (o *PaymentCreateResponse) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentCreateResponse) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentCreateResponse) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentCreateResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


