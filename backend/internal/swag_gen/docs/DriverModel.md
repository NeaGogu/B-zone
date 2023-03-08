# DriverModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**PauseId** | Pointer to **int32** | id of pause scheme to apply | [optional] 
**Uuid** | Pointer to **string** | unique per user | [optional] [readonly] 
**Addresses** | Pointer to [**[]AddressModel**](AddressModel.md) | user address (mostly interesting for drivers) | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**TagIds** | Pointer to **[]int32** | Tag type ids | [optional] 
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**NamePrefix** | Pointer to **string** | Name Prefix | [optional] 
**FullName** | Pointer to **string** | Full name | [optional] [readonly] 
**Initials** | Pointer to **string** | Full name | [optional] [readonly] 
**Email** | Pointer to **string** | Email | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**Zones** | Pointer to [**[]ZoneModel**](ZoneModel.md) |  | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**ZoneIds** | Pointer to **[]int32** | Zone ids | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Notes** | Pointer to [**[]NoteModel**](NoteModel.md) |  | [optional] 
**Removed** | Pointer to **bool** | Whether user is removed or not | [optional] 
**Active** | Pointer to **bool** | Whether user is still active or not | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**DriverCreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**DriverUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**DriverCreatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 
**DriverUpdatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 
**UpdatedByName** | Pointer to **string** | Driver updated by user full name | [optional] 

## Methods

### NewDriverModel

`func NewDriverModel() *DriverModel`

NewDriverModel instantiates a new DriverModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverModelWithDefaults

`func NewDriverModelWithDefaults() *DriverModel`

NewDriverModelWithDefaults instantiates a new DriverModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriverModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DriverModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPauseId

`func (o *DriverModel) GetPauseId() int32`

GetPauseId returns the PauseId field if non-nil, zero value otherwise.

### GetPauseIdOk

`func (o *DriverModel) GetPauseIdOk() (*int32, bool)`

GetPauseIdOk returns a tuple with the PauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseId

`func (o *DriverModel) SetPauseId(v int32)`

SetPauseId sets PauseId field to given value.

### HasPauseId

`func (o *DriverModel) HasPauseId() bool`

HasPauseId returns a boolean if a field has been set.

### GetUuid

`func (o *DriverModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DriverModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DriverModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DriverModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAddresses

`func (o *DriverModel) GetAddresses() []AddressModel`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *DriverModel) GetAddressesOk() (*[]AddressModel, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *DriverModel) SetAddresses(v []AddressModel)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *DriverModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTagNames

`func (o *DriverModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *DriverModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *DriverModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *DriverModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetTagIds

`func (o *DriverModel) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *DriverModel) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *DriverModel) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *DriverModel) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetFirstName

`func (o *DriverModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DriverModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DriverModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DriverModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *DriverModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DriverModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DriverModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DriverModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetNamePrefix

`func (o *DriverModel) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *DriverModel) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *DriverModel) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *DriverModel) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### GetFullName

`func (o *DriverModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DriverModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DriverModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DriverModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetInitials

`func (o *DriverModel) GetInitials() string`

GetInitials returns the Initials field if non-nil, zero value otherwise.

### GetInitialsOk

`func (o *DriverModel) GetInitialsOk() (*string, bool)`

GetInitialsOk returns a tuple with the Initials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitials

`func (o *DriverModel) SetInitials(v string)`

SetInitials sets Initials field to given value.

### HasInitials

`func (o *DriverModel) HasInitials() bool`

HasInitials returns a boolean if a field has been set.

### GetEmail

`func (o *DriverModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DriverModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DriverModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DriverModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTags

`func (o *DriverModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DriverModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DriverModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DriverModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetZones

`func (o *DriverModel) GetZones() []ZoneModel`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *DriverModel) GetZonesOk() (*[]ZoneModel, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *DriverModel) SetZones(v []ZoneModel)`

SetZones sets Zones field to given value.

### HasZones

`func (o *DriverModel) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZoneNames

`func (o *DriverModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *DriverModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *DriverModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *DriverModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetZoneIds

`func (o *DriverModel) GetZoneIds() []int32`

GetZoneIds returns the ZoneIds field if non-nil, zero value otherwise.

### GetZoneIdsOk

`func (o *DriverModel) GetZoneIdsOk() (*[]int32, bool)`

GetZoneIdsOk returns a tuple with the ZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIds

`func (o *DriverModel) SetZoneIds(v []int32)`

SetZoneIds sets ZoneIds field to given value.

### HasZoneIds

`func (o *DriverModel) HasZoneIds() bool`

HasZoneIds returns a boolean if a field has been set.

### GetLinks

`func (o *DriverModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DriverModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DriverModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DriverModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetNotes

`func (o *DriverModel) GetNotes() []NoteModel`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DriverModel) GetNotesOk() (*[]NoteModel, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DriverModel) SetNotes(v []NoteModel)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DriverModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRemoved

`func (o *DriverModel) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *DriverModel) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *DriverModel) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *DriverModel) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetActive

`func (o *DriverModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DriverModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DriverModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DriverModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMetaData

`func (o *DriverModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *DriverModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *DriverModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *DriverModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetFiles

`func (o *DriverModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DriverModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DriverModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *DriverModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetDriverCreatedAt

`func (o *DriverModel) GetDriverCreatedAt() time.Time`

GetDriverCreatedAt returns the DriverCreatedAt field if non-nil, zero value otherwise.

### GetDriverCreatedAtOk

`func (o *DriverModel) GetDriverCreatedAtOk() (*time.Time, bool)`

GetDriverCreatedAtOk returns a tuple with the DriverCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverCreatedAt

`func (o *DriverModel) SetDriverCreatedAt(v time.Time)`

SetDriverCreatedAt sets DriverCreatedAt field to given value.

### HasDriverCreatedAt

`func (o *DriverModel) HasDriverCreatedAt() bool`

HasDriverCreatedAt returns a boolean if a field has been set.

### GetDriverUpdatedAt

`func (o *DriverModel) GetDriverUpdatedAt() time.Time`

GetDriverUpdatedAt returns the DriverUpdatedAt field if non-nil, zero value otherwise.

### GetDriverUpdatedAtOk

`func (o *DriverModel) GetDriverUpdatedAtOk() (*time.Time, bool)`

GetDriverUpdatedAtOk returns a tuple with the DriverUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverUpdatedAt

`func (o *DriverModel) SetDriverUpdatedAt(v time.Time)`

SetDriverUpdatedAt sets DriverUpdatedAt field to given value.

### HasDriverUpdatedAt

`func (o *DriverModel) HasDriverUpdatedAt() bool`

HasDriverUpdatedAt returns a boolean if a field has been set.

### GetDriverCreatedBy

`func (o *DriverModel) GetDriverCreatedBy() int32`

GetDriverCreatedBy returns the DriverCreatedBy field if non-nil, zero value otherwise.

### GetDriverCreatedByOk

`func (o *DriverModel) GetDriverCreatedByOk() (*int32, bool)`

GetDriverCreatedByOk returns a tuple with the DriverCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverCreatedBy

`func (o *DriverModel) SetDriverCreatedBy(v int32)`

SetDriverCreatedBy sets DriverCreatedBy field to given value.

### HasDriverCreatedBy

`func (o *DriverModel) HasDriverCreatedBy() bool`

HasDriverCreatedBy returns a boolean if a field has been set.

### GetDriverUpdatedBy

`func (o *DriverModel) GetDriverUpdatedBy() int32`

GetDriverUpdatedBy returns the DriverUpdatedBy field if non-nil, zero value otherwise.

### GetDriverUpdatedByOk

`func (o *DriverModel) GetDriverUpdatedByOk() (*int32, bool)`

GetDriverUpdatedByOk returns a tuple with the DriverUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverUpdatedBy

`func (o *DriverModel) SetDriverUpdatedBy(v int32)`

SetDriverUpdatedBy sets DriverUpdatedBy field to given value.

### HasDriverUpdatedBy

`func (o *DriverModel) HasDriverUpdatedBy() bool`

HasDriverUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *DriverModel) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *DriverModel) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *DriverModel) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *DriverModel) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


