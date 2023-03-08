# PaymentFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Payment id&#39;s | [optional] 
**ActivityId** | Pointer to **[]int32** | Activity id&#39;s where the payment should belong to | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 

## Methods

### NewPaymentFiltersModel

`func NewPaymentFiltersModel() *PaymentFiltersModel`

NewPaymentFiltersModel instantiates a new PaymentFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentFiltersModelWithDefaults

`func NewPaymentFiltersModelWithDefaults() *PaymentFiltersModel`

NewPaymentFiltersModelWithDefaults instantiates a new PaymentFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityId

`func (o *PaymentFiltersModel) GetActivityId() []int32`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *PaymentFiltersModel) GetActivityIdOk() (*[]int32, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *PaymentFiltersModel) SetActivityId(v []int32)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *PaymentFiltersModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetSearchText

`func (o *PaymentFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *PaymentFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *PaymentFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *PaymentFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *PaymentFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *PaymentFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *PaymentFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *PaymentFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *PaymentFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *PaymentFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *PaymentFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *PaymentFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


