# TransactionTokenIsPaidResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **int32** | Is 1 on success (transaction id is found en transaction_type is online and amount is correct and is not yet paid), 0 when failed | [optional] 

## Methods

### NewTransactionTokenIsPaidResponse

`func NewTransactionTokenIsPaidResponse() *TransactionTokenIsPaidResponse`

NewTransactionTokenIsPaidResponse instantiates a new TransactionTokenIsPaidResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTokenIsPaidResponseWithDefaults

`func NewTransactionTokenIsPaidResponseWithDefaults() *TransactionTokenIsPaidResponse`

NewTransactionTokenIsPaidResponseWithDefaults instantiates a new TransactionTokenIsPaidResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TransactionTokenIsPaidResponse) GetSuccess() int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TransactionTokenIsPaidResponse) GetSuccessOk() (*int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TransactionTokenIsPaidResponse) SetSuccess(v int32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TransactionTokenIsPaidResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


