# TransactionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: Transaction has been removed and is no longer visible in any bumbal interface | [optional] 
**ActivityId** | Pointer to **int32** | activity id it belongs to | [optional] 
**Amount** | Pointer to **int32** | amount in cents, 42 euro is 4200 cents | [optional] 
**TransactionType** | Pointer to **int32** | Type of Transaction: 1 &#x3D; cash, 2 &#x3D; pin, 3 &#x3D; online | [optional] 
**Paid** | Pointer to **bool** | If paid &#x3D; 0: the transaction has not been fullfilled yet | [optional] 
**Failed** | Pointer to **bool** | if failed &#x3D; 1: the transaction has failed | [optional] 

## Methods

### NewTransactionModel

`func NewTransactionModel() *TransactionModel`

NewTransactionModel instantiates a new TransactionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionModelWithDefaults

`func NewTransactionModelWithDefaults() *TransactionModel`

NewTransactionModelWithDefaults instantiates a new TransactionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *TransactionModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TransactionModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TransactionModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TransactionModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActivityId

`func (o *TransactionModel) GetActivityId() int32`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *TransactionModel) GetActivityIdOk() (*int32, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *TransactionModel) SetActivityId(v int32)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *TransactionModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTransactionType

`func (o *TransactionModel) GetTransactionType() int32`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *TransactionModel) GetTransactionTypeOk() (*int32, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *TransactionModel) SetTransactionType(v int32)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *TransactionModel) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetPaid

`func (o *TransactionModel) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *TransactionModel) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *TransactionModel) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *TransactionModel) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetFailed

`func (o *TransactionModel) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *TransactionModel) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *TransactionModel) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *TransactionModel) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


