# NoteCategoryFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Bumbal note category id&#39;s | [optional] 
**Name** | Pointer to **[]string** | Bumbal note category names | [optional] 
**ShowInApp** | Pointer to **[]bool** | whether notes of this category should be visible for the driver in the mobile app | [optional] 
**Sys** | Pointer to **[]bool** | Bumbal note categories which can not be edited or removed | [optional] 

## Methods

### NewNoteCategoryFiltersModel

`func NewNoteCategoryFiltersModel() *NoteCategoryFiltersModel`

NewNoteCategoryFiltersModel instantiates a new NoteCategoryFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteCategoryFiltersModelWithDefaults

`func NewNoteCategoryFiltersModelWithDefaults() *NoteCategoryFiltersModel`

NewNoteCategoryFiltersModelWithDefaults instantiates a new NoteCategoryFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoteCategoryFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteCategoryFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteCategoryFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *NoteCategoryFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NoteCategoryFiltersModel) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NoteCategoryFiltersModel) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NoteCategoryFiltersModel) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *NoteCategoryFiltersModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShowInApp

`func (o *NoteCategoryFiltersModel) GetShowInApp() []bool`

GetShowInApp returns the ShowInApp field if non-nil, zero value otherwise.

### GetShowInAppOk

`func (o *NoteCategoryFiltersModel) GetShowInAppOk() (*[]bool, bool)`

GetShowInAppOk returns a tuple with the ShowInApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInApp

`func (o *NoteCategoryFiltersModel) SetShowInApp(v []bool)`

SetShowInApp sets ShowInApp field to given value.

### HasShowInApp

`func (o *NoteCategoryFiltersModel) HasShowInApp() bool`

HasShowInApp returns a boolean if a field has been set.

### GetSys

`func (o *NoteCategoryFiltersModel) GetSys() []bool`

GetSys returns the Sys field if non-nil, zero value otherwise.

### GetSysOk

`func (o *NoteCategoryFiltersModel) GetSysOk() (*[]bool, bool)`

GetSysOk returns a tuple with the Sys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSys

`func (o *NoteCategoryFiltersModel) SetSys(v []bool)`

SetSys sets Sys field to given value.

### HasSys

`func (o *NoteCategoryFiltersModel) HasSys() bool`

HasSys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


