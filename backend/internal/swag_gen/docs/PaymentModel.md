# PaymentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: Payment has been removed and is no longer visible in any bumbal interface | [optional] 
**ActivityId** | Pointer to **int32** | id of the activity this payment is bound to | [optional] 
**Amount** | Pointer to **int32** | amount of the payment in cents | [optional] 
**Title** | Pointer to **string** | title of the payment | [optional] 

## Methods

### NewPaymentModel

`func NewPaymentModel() *PaymentModel`

NewPaymentModel instantiates a new PaymentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentModelWithDefaults

`func NewPaymentModelWithDefaults() *PaymentModel`

NewPaymentModelWithDefaults instantiates a new PaymentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *PaymentModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PaymentModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActivityId

`func (o *PaymentModel) GetActivityId() int32`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *PaymentModel) GetActivityIdOk() (*int32, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *PaymentModel) SetActivityId(v int32)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *PaymentModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTitle

`func (o *PaymentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PaymentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


