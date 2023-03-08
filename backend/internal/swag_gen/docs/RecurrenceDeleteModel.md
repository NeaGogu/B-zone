# RecurrenceDeleteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**RemoveObjects** | Pointer to **bool** | Remove recurrence objects | [optional] 

## Methods

### NewRecurrenceDeleteModel

`func NewRecurrenceDeleteModel() *RecurrenceDeleteModel`

NewRecurrenceDeleteModel instantiates a new RecurrenceDeleteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceDeleteModelWithDefaults

`func NewRecurrenceDeleteModelWithDefaults() *RecurrenceDeleteModel`

NewRecurrenceDeleteModelWithDefaults instantiates a new RecurrenceDeleteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecurrenceDeleteModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurrenceDeleteModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurrenceDeleteModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RecurrenceDeleteModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoveObjects

`func (o *RecurrenceDeleteModel) GetRemoveObjects() bool`

GetRemoveObjects returns the RemoveObjects field if non-nil, zero value otherwise.

### GetRemoveObjectsOk

`func (o *RecurrenceDeleteModel) GetRemoveObjectsOk() (*bool, bool)`

GetRemoveObjectsOk returns a tuple with the RemoveObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveObjects

`func (o *RecurrenceDeleteModel) SetRemoveObjects(v bool)`

SetRemoveObjects sets RemoveObjects field to given value.

### HasRemoveObjects

`func (o *RecurrenceDeleteModel) HasRemoveObjects() bool`

HasRemoveObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


