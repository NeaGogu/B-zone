# NoteFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**ObjectType** | Pointer to **[]int32** | Object type IDs | [optional] 
**ObjectTypeName** | Pointer to **[]string** | Object type names | [optional] 
**NoteCategoryId** | Pointer to **[]int32** | Note category id | [optional] 
**NoteCategoryName** | Pointer to **[]string** | Note category name | [optional] 
**ObjectId** | Pointer to **[]int32** | Object ID | [optional] 
**Title** | Pointer to **[]string** | Note titles | [optional] 
**Content** | Pointer to **[]string** | Note content | [optional] 

## Methods

### NewNoteFiltersModel

`func NewNoteFiltersModel() *NoteFiltersModel`

NewNoteFiltersModel instantiates a new NoteFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteFiltersModelWithDefaults

`func NewNoteFiltersModelWithDefaults() *NoteFiltersModel`

NewNoteFiltersModelWithDefaults instantiates a new NoteFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAtSince

`func (o *NoteFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *NoteFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *NoteFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *NoteFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *NoteFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *NoteFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *NoteFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *NoteFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetObjectType

`func (o *NoteFiltersModel) GetObjectType() []int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NoteFiltersModel) GetObjectTypeOk() (*[]int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NoteFiltersModel) SetObjectType(v []int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *NoteFiltersModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *NoteFiltersModel) GetObjectTypeName() []string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *NoteFiltersModel) GetObjectTypeNameOk() (*[]string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *NoteFiltersModel) SetObjectTypeName(v []string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *NoteFiltersModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetNoteCategoryId

`func (o *NoteFiltersModel) GetNoteCategoryId() []int32`

GetNoteCategoryId returns the NoteCategoryId field if non-nil, zero value otherwise.

### GetNoteCategoryIdOk

`func (o *NoteFiltersModel) GetNoteCategoryIdOk() (*[]int32, bool)`

GetNoteCategoryIdOk returns a tuple with the NoteCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteCategoryId

`func (o *NoteFiltersModel) SetNoteCategoryId(v []int32)`

SetNoteCategoryId sets NoteCategoryId field to given value.

### HasNoteCategoryId

`func (o *NoteFiltersModel) HasNoteCategoryId() bool`

HasNoteCategoryId returns a boolean if a field has been set.

### GetNoteCategoryName

`func (o *NoteFiltersModel) GetNoteCategoryName() []string`

GetNoteCategoryName returns the NoteCategoryName field if non-nil, zero value otherwise.

### GetNoteCategoryNameOk

`func (o *NoteFiltersModel) GetNoteCategoryNameOk() (*[]string, bool)`

GetNoteCategoryNameOk returns a tuple with the NoteCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteCategoryName

`func (o *NoteFiltersModel) SetNoteCategoryName(v []string)`

SetNoteCategoryName sets NoteCategoryName field to given value.

### HasNoteCategoryName

`func (o *NoteFiltersModel) HasNoteCategoryName() bool`

HasNoteCategoryName returns a boolean if a field has been set.

### GetObjectId

`func (o *NoteFiltersModel) GetObjectId() []int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *NoteFiltersModel) GetObjectIdOk() (*[]int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *NoteFiltersModel) SetObjectId(v []int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *NoteFiltersModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetTitle

`func (o *NoteFiltersModel) GetTitle() []string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NoteFiltersModel) GetTitleOk() (*[]string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NoteFiltersModel) SetTitle(v []string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NoteFiltersModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *NoteFiltersModel) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteFiltersModel) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteFiltersModel) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *NoteFiltersModel) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


