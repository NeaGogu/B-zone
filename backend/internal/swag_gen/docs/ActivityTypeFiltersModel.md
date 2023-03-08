# ActivityTypeFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | ActivityType Unique Identifier(s) | [optional] 
**Name** | Pointer to **[]string** | ActivityType names | [optional] 
**Special** | Pointer to **[]bool** | special activities are ones which are not planned but simly added by the driver during the route execution | [optional] 
**AssignmentEntry** | Pointer to **[]bool** | assignment_entry activity types are used for order entry forms | [optional] 

## Methods

### NewActivityTypeFiltersModel

`func NewActivityTypeFiltersModel() *ActivityTypeFiltersModel`

NewActivityTypeFiltersModel instantiates a new ActivityTypeFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTypeFiltersModelWithDefaults

`func NewActivityTypeFiltersModelWithDefaults() *ActivityTypeFiltersModel`

NewActivityTypeFiltersModelWithDefaults instantiates a new ActivityTypeFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityTypeFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityTypeFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityTypeFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityTypeFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivityTypeFiltersModel) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityTypeFiltersModel) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityTypeFiltersModel) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ActivityTypeFiltersModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpecial

`func (o *ActivityTypeFiltersModel) GetSpecial() []bool`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *ActivityTypeFiltersModel) GetSpecialOk() (*[]bool, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *ActivityTypeFiltersModel) SetSpecial(v []bool)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *ActivityTypeFiltersModel) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.

### GetAssignmentEntry

`func (o *ActivityTypeFiltersModel) GetAssignmentEntry() []bool`

GetAssignmentEntry returns the AssignmentEntry field if non-nil, zero value otherwise.

### GetAssignmentEntryOk

`func (o *ActivityTypeFiltersModel) GetAssignmentEntryOk() (*[]bool, bool)`

GetAssignmentEntryOk returns a tuple with the AssignmentEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentEntry

`func (o *ActivityTypeFiltersModel) SetAssignmentEntry(v []bool)`

SetAssignmentEntry sets AssignmentEntry field to given value.

### HasAssignmentEntry

`func (o *ActivityTypeFiltersModel) HasAssignmentEntry() bool`

HasAssignmentEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


