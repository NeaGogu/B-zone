# TransactionFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Transaction id&#39;s | [optional] 
**ActivityId** | Pointer to **[]int32** | Activity id&#39;s where the transaction should belong to | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 

## Methods

### NewTransactionFiltersModel

`func NewTransactionFiltersModel() *TransactionFiltersModel`

NewTransactionFiltersModel instantiates a new TransactionFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFiltersModelWithDefaults

`func NewTransactionFiltersModelWithDefaults() *TransactionFiltersModel`

NewTransactionFiltersModelWithDefaults instantiates a new TransactionFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityId

`func (o *TransactionFiltersModel) GetActivityId() []int32`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *TransactionFiltersModel) GetActivityIdOk() (*[]int32, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *TransactionFiltersModel) SetActivityId(v []int32)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *TransactionFiltersModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetSearchText

`func (o *TransactionFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *TransactionFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *TransactionFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *TransactionFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *TransactionFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *TransactionFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *TransactionFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *TransactionFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *TransactionFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *TransactionFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *TransactionFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *TransactionFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


