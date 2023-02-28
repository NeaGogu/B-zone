# TransactionTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **int32** | Is 1 on success (transaction id is found en transaction_type is online), 0 when failed | [optional] 
**Token** | Pointer to **string** | Token on success | [optional] 

## Methods

### NewTransactionTokenResponse

`func NewTransactionTokenResponse() *TransactionTokenResponse`

NewTransactionTokenResponse instantiates a new TransactionTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTokenResponseWithDefaults

`func NewTransactionTokenResponseWithDefaults() *TransactionTokenResponse`

NewTransactionTokenResponseWithDefaults instantiates a new TransactionTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TransactionTokenResponse) GetSuccess() int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TransactionTokenResponse) GetSuccessOk() (*int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TransactionTokenResponse) SetSuccess(v int32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TransactionTokenResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetToken

`func (o *TransactionTokenResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TransactionTokenResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TransactionTokenResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TransactionTokenResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


