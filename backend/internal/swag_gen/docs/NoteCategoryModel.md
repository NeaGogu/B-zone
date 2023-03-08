# NoteCategoryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | name of note category | [optional] 
**Sys** | Pointer to **bool** | system type categories can not be edited or deleted | [optional] 
**ShowInApp** | Pointer to **bool** | Whether the note category notes should be visible for teh drivers usign the bumbal mobile app | [optional] 

## Methods

### NewNoteCategoryModel

`func NewNoteCategoryModel() *NoteCategoryModel`

NewNoteCategoryModel instantiates a new NoteCategoryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteCategoryModelWithDefaults

`func NewNoteCategoryModelWithDefaults() *NoteCategoryModel`

NewNoteCategoryModelWithDefaults instantiates a new NoteCategoryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoteCategoryModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteCategoryModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteCategoryModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NoteCategoryModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NoteCategoryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NoteCategoryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NoteCategoryModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NoteCategoryModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSys

`func (o *NoteCategoryModel) GetSys() bool`

GetSys returns the Sys field if non-nil, zero value otherwise.

### GetSysOk

`func (o *NoteCategoryModel) GetSysOk() (*bool, bool)`

GetSysOk returns a tuple with the Sys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSys

`func (o *NoteCategoryModel) SetSys(v bool)`

SetSys sets Sys field to given value.

### HasSys

`func (o *NoteCategoryModel) HasSys() bool`

HasSys returns a boolean if a field has been set.

### GetShowInApp

`func (o *NoteCategoryModel) GetShowInApp() bool`

GetShowInApp returns the ShowInApp field if non-nil, zero value otherwise.

### GetShowInAppOk

`func (o *NoteCategoryModel) GetShowInAppOk() (*bool, bool)`

GetShowInAppOk returns a tuple with the ShowInApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInApp

`func (o *NoteCategoryModel) SetShowInApp(v bool)`

SetShowInApp sets ShowInApp field to given value.

### HasShowInApp

`func (o *NoteCategoryModel) HasShowInApp() bool`

HasShowInApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


