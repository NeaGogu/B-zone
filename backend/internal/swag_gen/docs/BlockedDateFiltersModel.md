# BlockedDateFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Unique Identifier | [optional] 
**DateTimeFrom** | Pointer to **time.Time** | DateTime From | [optional] 
**DateTimeTo** | Pointer to **time.Time** | DateTime To | [optional] 

## Methods

### NewBlockedDateFiltersModel

`func NewBlockedDateFiltersModel() *BlockedDateFiltersModel`

NewBlockedDateFiltersModel instantiates a new BlockedDateFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockedDateFiltersModelWithDefaults

`func NewBlockedDateFiltersModelWithDefaults() *BlockedDateFiltersModel`

NewBlockedDateFiltersModelWithDefaults instantiates a new BlockedDateFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlockedDateFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockedDateFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockedDateFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *BlockedDateFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *BlockedDateFiltersModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *BlockedDateFiltersModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *BlockedDateFiltersModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *BlockedDateFiltersModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *BlockedDateFiltersModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *BlockedDateFiltersModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *BlockedDateFiltersModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *BlockedDateFiltersModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


