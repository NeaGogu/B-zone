# NoteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**ObjectType** | Pointer to **int64** | Object type ID | [optional] 
**ObjectTypeName** | Pointer to **string** | Object type name | [optional] 
**ObjectId** | Pointer to **int32** | Object ID | [optional] 
**ObjectLink** | Pointer to [**LinkModel**](LinkModel.md) |  | [optional] 
**NoteCategoryId** | Pointer to **int32** | Note category id | [optional] 
**NoteCategoryName** | Pointer to **string** | Note category name | [optional] 
**VisibleForDriver** | Pointer to **bool** | Notition is visible in driver App | [optional] 
**Title** | Pointer to **string** | Note title | [optional] 
**UpdatedByName** | Pointer to **string** | Note updated by user full name | [optional] 
**Content** | Pointer to **string** | Note content | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: note has been removed and is no longer visible in any bumbal interface | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**CreatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 
**UpdatedBy** | Pointer to **int32** | updated_by user id | [optional] [readonly] 

## Methods

### NewNoteModel

`func NewNoteModel() *NoteModel`

NewNoteModel instantiates a new NoteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteModelWithDefaults

`func NewNoteModelWithDefaults() *NoteModel`

NewNoteModelWithDefaults instantiates a new NoteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoteModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NoteModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *NoteModel) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NoteModel) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NoteModel) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *NoteModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *NoteModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *NoteModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *NoteModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *NoteModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetObjectId

`func (o *NoteModel) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *NoteModel) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *NoteModel) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *NoteModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectLink

`func (o *NoteModel) GetObjectLink() LinkModel`

GetObjectLink returns the ObjectLink field if non-nil, zero value otherwise.

### GetObjectLinkOk

`func (o *NoteModel) GetObjectLinkOk() (*LinkModel, bool)`

GetObjectLinkOk returns a tuple with the ObjectLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectLink

`func (o *NoteModel) SetObjectLink(v LinkModel)`

SetObjectLink sets ObjectLink field to given value.

### HasObjectLink

`func (o *NoteModel) HasObjectLink() bool`

HasObjectLink returns a boolean if a field has been set.

### GetNoteCategoryId

`func (o *NoteModel) GetNoteCategoryId() int32`

GetNoteCategoryId returns the NoteCategoryId field if non-nil, zero value otherwise.

### GetNoteCategoryIdOk

`func (o *NoteModel) GetNoteCategoryIdOk() (*int32, bool)`

GetNoteCategoryIdOk returns a tuple with the NoteCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteCategoryId

`func (o *NoteModel) SetNoteCategoryId(v int32)`

SetNoteCategoryId sets NoteCategoryId field to given value.

### HasNoteCategoryId

`func (o *NoteModel) HasNoteCategoryId() bool`

HasNoteCategoryId returns a boolean if a field has been set.

### GetNoteCategoryName

`func (o *NoteModel) GetNoteCategoryName() string`

GetNoteCategoryName returns the NoteCategoryName field if non-nil, zero value otherwise.

### GetNoteCategoryNameOk

`func (o *NoteModel) GetNoteCategoryNameOk() (*string, bool)`

GetNoteCategoryNameOk returns a tuple with the NoteCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteCategoryName

`func (o *NoteModel) SetNoteCategoryName(v string)`

SetNoteCategoryName sets NoteCategoryName field to given value.

### HasNoteCategoryName

`func (o *NoteModel) HasNoteCategoryName() bool`

HasNoteCategoryName returns a boolean if a field has been set.

### GetVisibleForDriver

`func (o *NoteModel) GetVisibleForDriver() bool`

GetVisibleForDriver returns the VisibleForDriver field if non-nil, zero value otherwise.

### GetVisibleForDriverOk

`func (o *NoteModel) GetVisibleForDriverOk() (*bool, bool)`

GetVisibleForDriverOk returns a tuple with the VisibleForDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleForDriver

`func (o *NoteModel) SetVisibleForDriver(v bool)`

SetVisibleForDriver sets VisibleForDriver field to given value.

### HasVisibleForDriver

`func (o *NoteModel) HasVisibleForDriver() bool`

HasVisibleForDriver returns a boolean if a field has been set.

### GetTitle

`func (o *NoteModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NoteModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NoteModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NoteModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *NoteModel) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *NoteModel) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *NoteModel) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *NoteModel) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.

### GetContent

`func (o *NoteModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteModel) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *NoteModel) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFiles

`func (o *NoteModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *NoteModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *NoteModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *NoteModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLinks

`func (o *NoteModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NoteModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NoteModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NoteModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *NoteModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *NoteModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *NoteModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *NoteModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetActive

`func (o *NoteModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NoteModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NoteModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NoteModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NoteModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NoteModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NoteModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NoteModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NoteModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NoteModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NoteModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NoteModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *NoteModel) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NoteModel) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *NoteModel) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *NoteModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *NoteModel) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *NoteModel) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *NoteModel) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *NoteModel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


