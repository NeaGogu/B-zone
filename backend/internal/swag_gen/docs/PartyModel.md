# PartyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**PartyTypeName** | Pointer to **string** | Type of this party | [optional] 
**PartyTypeId** | Pointer to **int32** | PartyTypeID of this party. 2 &#x3D; contractor, 3 &#x3D; booking | [optional] 
**Name1** | Pointer to **string** | Name 1 for party | [optional] 
**Name2** | Pointer to **string** | Name 2 for party | [optional] 
**Nr** | Pointer to **string** | Number of this party | [optional] 
**ContactPerson** | Pointer to **string** | Contact person for party | [optional] 
**Url** | Pointer to **string** | Url for party website | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**TagIds** | Pointer to **[]int32** | Tag ids | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Notes** | Pointer to [**[]NoteModel**](NoteModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**PartyCreatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**PartyUpdatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 

## Methods

### NewPartyModel

`func NewPartyModel() *PartyModel`

NewPartyModel instantiates a new PartyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyModelWithDefaults

`func NewPartyModelWithDefaults() *PartyModel`

NewPartyModelWithDefaults instantiates a new PartyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartyModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PartyModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartyTypeName

`func (o *PartyModel) GetPartyTypeName() string`

GetPartyTypeName returns the PartyTypeName field if non-nil, zero value otherwise.

### GetPartyTypeNameOk

`func (o *PartyModel) GetPartyTypeNameOk() (*string, bool)`

GetPartyTypeNameOk returns a tuple with the PartyTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyTypeName

`func (o *PartyModel) SetPartyTypeName(v string)`

SetPartyTypeName sets PartyTypeName field to given value.

### HasPartyTypeName

`func (o *PartyModel) HasPartyTypeName() bool`

HasPartyTypeName returns a boolean if a field has been set.

### GetPartyTypeId

`func (o *PartyModel) GetPartyTypeId() int32`

GetPartyTypeId returns the PartyTypeId field if non-nil, zero value otherwise.

### GetPartyTypeIdOk

`func (o *PartyModel) GetPartyTypeIdOk() (*int32, bool)`

GetPartyTypeIdOk returns a tuple with the PartyTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyTypeId

`func (o *PartyModel) SetPartyTypeId(v int32)`

SetPartyTypeId sets PartyTypeId field to given value.

### HasPartyTypeId

`func (o *PartyModel) HasPartyTypeId() bool`

HasPartyTypeId returns a boolean if a field has been set.

### GetName1

`func (o *PartyModel) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *PartyModel) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *PartyModel) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *PartyModel) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *PartyModel) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *PartyModel) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *PartyModel) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *PartyModel) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetNr

`func (o *PartyModel) GetNr() string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *PartyModel) GetNrOk() (*string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *PartyModel) SetNr(v string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *PartyModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetContactPerson

`func (o *PartyModel) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *PartyModel) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *PartyModel) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *PartyModel) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetUrl

`func (o *PartyModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PartyModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PartyModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PartyModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTags

`func (o *PartyModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PartyModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PartyModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PartyModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagNames

`func (o *PartyModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *PartyModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *PartyModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *PartyModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetTagIds

`func (o *PartyModel) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *PartyModel) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *PartyModel) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *PartyModel) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetLinks

`func (o *PartyModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PartyModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PartyModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PartyModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *PartyModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *PartyModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *PartyModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *PartyModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetNotes

`func (o *PartyModel) GetNotes() []NoteModel`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PartyModel) GetNotesOk() (*[]NoteModel, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PartyModel) SetNotes(v []NoteModel)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PartyModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFiles

`func (o *PartyModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *PartyModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *PartyModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *PartyModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PartyModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PartyModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PartyModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PartyModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PartyModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PartyModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PartyModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PartyModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPartyCreatedByUser

`func (o *PartyModel) GetPartyCreatedByUser() UsersModel`

GetPartyCreatedByUser returns the PartyCreatedByUser field if non-nil, zero value otherwise.

### GetPartyCreatedByUserOk

`func (o *PartyModel) GetPartyCreatedByUserOk() (*UsersModel, bool)`

GetPartyCreatedByUserOk returns a tuple with the PartyCreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyCreatedByUser

`func (o *PartyModel) SetPartyCreatedByUser(v UsersModel)`

SetPartyCreatedByUser sets PartyCreatedByUser field to given value.

### HasPartyCreatedByUser

`func (o *PartyModel) HasPartyCreatedByUser() bool`

HasPartyCreatedByUser returns a boolean if a field has been set.

### GetPartyUpdatedByUser

`func (o *PartyModel) GetPartyUpdatedByUser() UsersModel`

GetPartyUpdatedByUser returns the PartyUpdatedByUser field if non-nil, zero value otherwise.

### GetPartyUpdatedByUserOk

`func (o *PartyModel) GetPartyUpdatedByUserOk() (*UsersModel, bool)`

GetPartyUpdatedByUserOk returns a tuple with the PartyUpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyUpdatedByUser

`func (o *PartyModel) SetPartyUpdatedByUser(v UsersModel)`

SetPartyUpdatedByUser sets PartyUpdatedByUser field to given value.

### HasPartyUpdatedByUser

`func (o *PartyModel) HasPartyUpdatedByUser() bool`

HasPartyUpdatedByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


