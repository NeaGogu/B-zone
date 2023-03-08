# FileModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**FileTypeId** | Pointer to **int64** |  | [optional] 
**FileTypeName** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**ObjectType** | Pointer to **int64** | Object type ID | [optional] 
**ObjectTypeName** | Pointer to **string** | Object type name | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Base64** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**CreatedBy** | Pointer to **int64** |  | [optional] 
**UpdatedBy** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at | [optional] [readonly] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 

## Methods

### NewFileModel

`func NewFileModel() *FileModel`

NewFileModel instantiates a new FileModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileModelWithDefaults

`func NewFileModelWithDefaults() *FileModel`

NewFileModelWithDefaults instantiates a new FileModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FileModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *FileModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FileModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FileModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FileModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetFileTypeId

`func (o *FileModel) GetFileTypeId() int64`

GetFileTypeId returns the FileTypeId field if non-nil, zero value otherwise.

### GetFileTypeIdOk

`func (o *FileModel) GetFileTypeIdOk() (*int64, bool)`

GetFileTypeIdOk returns a tuple with the FileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTypeId

`func (o *FileModel) SetFileTypeId(v int64)`

SetFileTypeId sets FileTypeId field to given value.

### HasFileTypeId

`func (o *FileModel) HasFileTypeId() bool`

HasFileTypeId returns a boolean if a field has been set.

### GetFileTypeName

`func (o *FileModel) GetFileTypeName() string`

GetFileTypeName returns the FileTypeName field if non-nil, zero value otherwise.

### GetFileTypeNameOk

`func (o *FileModel) GetFileTypeNameOk() (*string, bool)`

GetFileTypeNameOk returns a tuple with the FileTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTypeName

`func (o *FileModel) SetFileTypeName(v string)`

SetFileTypeName sets FileTypeName field to given value.

### HasFileTypeName

`func (o *FileModel) HasFileTypeName() bool`

HasFileTypeName returns a boolean if a field has been set.

### GetObjectId

`func (o *FileModel) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *FileModel) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *FileModel) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *FileModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *FileModel) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FileModel) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FileModel) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *FileModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *FileModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *FileModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *FileModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *FileModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetReference

`func (o *FileModel) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *FileModel) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *FileModel) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *FileModel) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetLocation

`func (o *FileModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FileModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FileModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FileModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBase64

`func (o *FileModel) GetBase64() string`

GetBase64 returns the Base64 field if non-nil, zero value otherwise.

### GetBase64Ok

`func (o *FileModel) GetBase64Ok() (*string, bool)`

GetBase64Ok returns a tuple with the Base64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64

`func (o *FileModel) SetBase64(v string)`

SetBase64 sets Base64 field to given value.

### HasBase64

`func (o *FileModel) HasBase64() bool`

HasBase64 returns a boolean if a field has been set.

### GetMetaData

`func (o *FileModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *FileModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *FileModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *FileModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FileModel) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FileModel) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FileModel) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FileModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FileModel) GetUpdatedBy() int64`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FileModel) GetUpdatedByOk() (*int64, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FileModel) SetUpdatedBy(v int64)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FileModel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FileModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FileModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FileModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FileModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *FileModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FileModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FileModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FileModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


