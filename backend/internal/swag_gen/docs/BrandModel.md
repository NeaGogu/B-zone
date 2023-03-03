# BrandModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Name of brand | [optional] [readonly] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**BrandColours** | Pointer to [**[]BrandColourModel**](BrandColourModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**PortalSettings** | Pointer to [**[]PortalSettingModel**](PortalSettingModel.md) |  | [optional] 
**BrandCreatedBy** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**BrandUpdatedBy** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**BrandCreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**BrandUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 

## Methods

### NewBrandModel

`func NewBrandModel() *BrandModel`

NewBrandModel instantiates a new BrandModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandModelWithDefaults

`func NewBrandModelWithDefaults() *BrandModel`

NewBrandModelWithDefaults instantiates a new BrandModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BrandModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrandModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrandModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BrandModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BrandModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrandModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrandModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BrandModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *BrandModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BrandModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BrandModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BrandModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLinks

`func (o *BrandModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrandModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrandModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrandModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *BrandModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *BrandModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *BrandModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *BrandModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetBrandColours

`func (o *BrandModel) GetBrandColours() []BrandColourModel`

GetBrandColours returns the BrandColours field if non-nil, zero value otherwise.

### GetBrandColoursOk

`func (o *BrandModel) GetBrandColoursOk() (*[]BrandColourModel, bool)`

GetBrandColoursOk returns a tuple with the BrandColours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandColours

`func (o *BrandModel) SetBrandColours(v []BrandColourModel)`

SetBrandColours sets BrandColours field to given value.

### HasBrandColours

`func (o *BrandModel) HasBrandColours() bool`

HasBrandColours returns a boolean if a field has been set.

### GetFiles

`func (o *BrandModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *BrandModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *BrandModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *BrandModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPortalSettings

`func (o *BrandModel) GetPortalSettings() []PortalSettingModel`

GetPortalSettings returns the PortalSettings field if non-nil, zero value otherwise.

### GetPortalSettingsOk

`func (o *BrandModel) GetPortalSettingsOk() (*[]PortalSettingModel, bool)`

GetPortalSettingsOk returns a tuple with the PortalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalSettings

`func (o *BrandModel) SetPortalSettings(v []PortalSettingModel)`

SetPortalSettings sets PortalSettings field to given value.

### HasPortalSettings

`func (o *BrandModel) HasPortalSettings() bool`

HasPortalSettings returns a boolean if a field has been set.

### GetBrandCreatedBy

`func (o *BrandModel) GetBrandCreatedBy() UsersModel`

GetBrandCreatedBy returns the BrandCreatedBy field if non-nil, zero value otherwise.

### GetBrandCreatedByOk

`func (o *BrandModel) GetBrandCreatedByOk() (*UsersModel, bool)`

GetBrandCreatedByOk returns a tuple with the BrandCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandCreatedBy

`func (o *BrandModel) SetBrandCreatedBy(v UsersModel)`

SetBrandCreatedBy sets BrandCreatedBy field to given value.

### HasBrandCreatedBy

`func (o *BrandModel) HasBrandCreatedBy() bool`

HasBrandCreatedBy returns a boolean if a field has been set.

### GetBrandUpdatedBy

`func (o *BrandModel) GetBrandUpdatedBy() UsersModel`

GetBrandUpdatedBy returns the BrandUpdatedBy field if non-nil, zero value otherwise.

### GetBrandUpdatedByOk

`func (o *BrandModel) GetBrandUpdatedByOk() (*UsersModel, bool)`

GetBrandUpdatedByOk returns a tuple with the BrandUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandUpdatedBy

`func (o *BrandModel) SetBrandUpdatedBy(v UsersModel)`

SetBrandUpdatedBy sets BrandUpdatedBy field to given value.

### HasBrandUpdatedBy

`func (o *BrandModel) HasBrandUpdatedBy() bool`

HasBrandUpdatedBy returns a boolean if a field has been set.

### GetBrandCreatedAt

`func (o *BrandModel) GetBrandCreatedAt() time.Time`

GetBrandCreatedAt returns the BrandCreatedAt field if non-nil, zero value otherwise.

### GetBrandCreatedAtOk

`func (o *BrandModel) GetBrandCreatedAtOk() (*time.Time, bool)`

GetBrandCreatedAtOk returns a tuple with the BrandCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandCreatedAt

`func (o *BrandModel) SetBrandCreatedAt(v time.Time)`

SetBrandCreatedAt sets BrandCreatedAt field to given value.

### HasBrandCreatedAt

`func (o *BrandModel) HasBrandCreatedAt() bool`

HasBrandCreatedAt returns a boolean if a field has been set.

### GetBrandUpdatedAt

`func (o *BrandModel) GetBrandUpdatedAt() time.Time`

GetBrandUpdatedAt returns the BrandUpdatedAt field if non-nil, zero value otherwise.

### GetBrandUpdatedAtOk

`func (o *BrandModel) GetBrandUpdatedAtOk() (*time.Time, bool)`

GetBrandUpdatedAtOk returns a tuple with the BrandUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandUpdatedAt

`func (o *BrandModel) SetBrandUpdatedAt(v time.Time)`

SetBrandUpdatedAt sets BrandUpdatedAt field to given value.

### HasBrandUpdatedAt

`func (o *BrandModel) HasBrandUpdatedAt() bool`

HasBrandUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


