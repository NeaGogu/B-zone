# DriverOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeRecordInfo** | Pointer to **bool** |  | [optional] 
**IncludeMetaData** | Pointer to **bool** |  | [optional] 
**IncludeTags** | Pointer to **bool** |  | [optional] 
**IncludeTagNames** | Pointer to **bool** |  | [optional] 
**IncludeTagIds** | Pointer to **bool** |  | [optional] 
**IncludeAddresses** | Pointer to **bool** |  | [optional] 
**IncludeLinks** | Pointer to **bool** |  | [optional] 
**IncludeNotes** | Pointer to **bool** |  | [optional] 
**IncludeZones** | Pointer to **bool** |  | [optional] 
**IncludeZoneNames** | Pointer to **bool** |  | [optional] 
**IncludeZoneIds** | Pointer to **bool** |  | [optional] 
**IncludeFiles** | Pointer to **bool** |  | [optional] 
**IncludeDriverRecordInfo** | Pointer to **bool** | Deprecated! Use include_record_info instead | [optional] 
**IncludeDriverMetaData** | Pointer to **bool** | Deprecated! Use include_meta_data instead | [optional] 
**IncludeDriverTags** | Pointer to **bool** | Deprecated! Use include_tags instead | [optional] 
**IncludeDriverTagTypeNames** | Pointer to **bool** | Deprecated! Use include_tag_names instead | [optional] 
**IncludeDriverLinks** | Pointer to **bool** | Deprecated! Use include_links instead | [optional] 
**IncludeDriverNotes** | Pointer to **bool** | Deprecated! Use include_notes instead | [optional] 
**IncludeDriverNoteTags** | Pointer to **bool** | Deprecated! Use include_notes + include_tags instead | [optional] 

## Methods

### NewDriverOptionsModel

`func NewDriverOptionsModel() *DriverOptionsModel`

NewDriverOptionsModel instantiates a new DriverOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverOptionsModelWithDefaults

`func NewDriverOptionsModelWithDefaults() *DriverOptionsModel`

NewDriverOptionsModelWithDefaults instantiates a new DriverOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeRecordInfo

`func (o *DriverOptionsModel) GetIncludeRecordInfo() bool`

GetIncludeRecordInfo returns the IncludeRecordInfo field if non-nil, zero value otherwise.

### GetIncludeRecordInfoOk

`func (o *DriverOptionsModel) GetIncludeRecordInfoOk() (*bool, bool)`

GetIncludeRecordInfoOk returns a tuple with the IncludeRecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordInfo

`func (o *DriverOptionsModel) SetIncludeRecordInfo(v bool)`

SetIncludeRecordInfo sets IncludeRecordInfo field to given value.

### HasIncludeRecordInfo

`func (o *DriverOptionsModel) HasIncludeRecordInfo() bool`

HasIncludeRecordInfo returns a boolean if a field has been set.

### GetIncludeMetaData

`func (o *DriverOptionsModel) GetIncludeMetaData() bool`

GetIncludeMetaData returns the IncludeMetaData field if non-nil, zero value otherwise.

### GetIncludeMetaDataOk

`func (o *DriverOptionsModel) GetIncludeMetaDataOk() (*bool, bool)`

GetIncludeMetaDataOk returns a tuple with the IncludeMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetaData

`func (o *DriverOptionsModel) SetIncludeMetaData(v bool)`

SetIncludeMetaData sets IncludeMetaData field to given value.

### HasIncludeMetaData

`func (o *DriverOptionsModel) HasIncludeMetaData() bool`

HasIncludeMetaData returns a boolean if a field has been set.

### GetIncludeTags

`func (o *DriverOptionsModel) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *DriverOptionsModel) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *DriverOptionsModel) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *DriverOptionsModel) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetIncludeTagNames

`func (o *DriverOptionsModel) GetIncludeTagNames() bool`

GetIncludeTagNames returns the IncludeTagNames field if non-nil, zero value otherwise.

### GetIncludeTagNamesOk

`func (o *DriverOptionsModel) GetIncludeTagNamesOk() (*bool, bool)`

GetIncludeTagNamesOk returns a tuple with the IncludeTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTagNames

`func (o *DriverOptionsModel) SetIncludeTagNames(v bool)`

SetIncludeTagNames sets IncludeTagNames field to given value.

### HasIncludeTagNames

`func (o *DriverOptionsModel) HasIncludeTagNames() bool`

HasIncludeTagNames returns a boolean if a field has been set.

### GetIncludeTagIds

`func (o *DriverOptionsModel) GetIncludeTagIds() bool`

GetIncludeTagIds returns the IncludeTagIds field if non-nil, zero value otherwise.

### GetIncludeTagIdsOk

`func (o *DriverOptionsModel) GetIncludeTagIdsOk() (*bool, bool)`

GetIncludeTagIdsOk returns a tuple with the IncludeTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTagIds

`func (o *DriverOptionsModel) SetIncludeTagIds(v bool)`

SetIncludeTagIds sets IncludeTagIds field to given value.

### HasIncludeTagIds

`func (o *DriverOptionsModel) HasIncludeTagIds() bool`

HasIncludeTagIds returns a boolean if a field has been set.

### GetIncludeAddresses

`func (o *DriverOptionsModel) GetIncludeAddresses() bool`

GetIncludeAddresses returns the IncludeAddresses field if non-nil, zero value otherwise.

### GetIncludeAddressesOk

`func (o *DriverOptionsModel) GetIncludeAddressesOk() (*bool, bool)`

GetIncludeAddressesOk returns a tuple with the IncludeAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAddresses

`func (o *DriverOptionsModel) SetIncludeAddresses(v bool)`

SetIncludeAddresses sets IncludeAddresses field to given value.

### HasIncludeAddresses

`func (o *DriverOptionsModel) HasIncludeAddresses() bool`

HasIncludeAddresses returns a boolean if a field has been set.

### GetIncludeLinks

`func (o *DriverOptionsModel) GetIncludeLinks() bool`

GetIncludeLinks returns the IncludeLinks field if non-nil, zero value otherwise.

### GetIncludeLinksOk

`func (o *DriverOptionsModel) GetIncludeLinksOk() (*bool, bool)`

GetIncludeLinksOk returns a tuple with the IncludeLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLinks

`func (o *DriverOptionsModel) SetIncludeLinks(v bool)`

SetIncludeLinks sets IncludeLinks field to given value.

### HasIncludeLinks

`func (o *DriverOptionsModel) HasIncludeLinks() bool`

HasIncludeLinks returns a boolean if a field has been set.

### GetIncludeNotes

`func (o *DriverOptionsModel) GetIncludeNotes() bool`

GetIncludeNotes returns the IncludeNotes field if non-nil, zero value otherwise.

### GetIncludeNotesOk

`func (o *DriverOptionsModel) GetIncludeNotesOk() (*bool, bool)`

GetIncludeNotesOk returns a tuple with the IncludeNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNotes

`func (o *DriverOptionsModel) SetIncludeNotes(v bool)`

SetIncludeNotes sets IncludeNotes field to given value.

### HasIncludeNotes

`func (o *DriverOptionsModel) HasIncludeNotes() bool`

HasIncludeNotes returns a boolean if a field has been set.

### GetIncludeZones

`func (o *DriverOptionsModel) GetIncludeZones() bool`

GetIncludeZones returns the IncludeZones field if non-nil, zero value otherwise.

### GetIncludeZonesOk

`func (o *DriverOptionsModel) GetIncludeZonesOk() (*bool, bool)`

GetIncludeZonesOk returns a tuple with the IncludeZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZones

`func (o *DriverOptionsModel) SetIncludeZones(v bool)`

SetIncludeZones sets IncludeZones field to given value.

### HasIncludeZones

`func (o *DriverOptionsModel) HasIncludeZones() bool`

HasIncludeZones returns a boolean if a field has been set.

### GetIncludeZoneNames

`func (o *DriverOptionsModel) GetIncludeZoneNames() bool`

GetIncludeZoneNames returns the IncludeZoneNames field if non-nil, zero value otherwise.

### GetIncludeZoneNamesOk

`func (o *DriverOptionsModel) GetIncludeZoneNamesOk() (*bool, bool)`

GetIncludeZoneNamesOk returns a tuple with the IncludeZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZoneNames

`func (o *DriverOptionsModel) SetIncludeZoneNames(v bool)`

SetIncludeZoneNames sets IncludeZoneNames field to given value.

### HasIncludeZoneNames

`func (o *DriverOptionsModel) HasIncludeZoneNames() bool`

HasIncludeZoneNames returns a boolean if a field has been set.

### GetIncludeZoneIds

`func (o *DriverOptionsModel) GetIncludeZoneIds() bool`

GetIncludeZoneIds returns the IncludeZoneIds field if non-nil, zero value otherwise.

### GetIncludeZoneIdsOk

`func (o *DriverOptionsModel) GetIncludeZoneIdsOk() (*bool, bool)`

GetIncludeZoneIdsOk returns a tuple with the IncludeZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZoneIds

`func (o *DriverOptionsModel) SetIncludeZoneIds(v bool)`

SetIncludeZoneIds sets IncludeZoneIds field to given value.

### HasIncludeZoneIds

`func (o *DriverOptionsModel) HasIncludeZoneIds() bool`

HasIncludeZoneIds returns a boolean if a field has been set.

### GetIncludeFiles

`func (o *DriverOptionsModel) GetIncludeFiles() bool`

GetIncludeFiles returns the IncludeFiles field if non-nil, zero value otherwise.

### GetIncludeFilesOk

`func (o *DriverOptionsModel) GetIncludeFilesOk() (*bool, bool)`

GetIncludeFilesOk returns a tuple with the IncludeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFiles

`func (o *DriverOptionsModel) SetIncludeFiles(v bool)`

SetIncludeFiles sets IncludeFiles field to given value.

### HasIncludeFiles

`func (o *DriverOptionsModel) HasIncludeFiles() bool`

HasIncludeFiles returns a boolean if a field has been set.

### GetIncludeDriverRecordInfo

`func (o *DriverOptionsModel) GetIncludeDriverRecordInfo() bool`

GetIncludeDriverRecordInfo returns the IncludeDriverRecordInfo field if non-nil, zero value otherwise.

### GetIncludeDriverRecordInfoOk

`func (o *DriverOptionsModel) GetIncludeDriverRecordInfoOk() (*bool, bool)`

GetIncludeDriverRecordInfoOk returns a tuple with the IncludeDriverRecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverRecordInfo

`func (o *DriverOptionsModel) SetIncludeDriverRecordInfo(v bool)`

SetIncludeDriverRecordInfo sets IncludeDriverRecordInfo field to given value.

### HasIncludeDriverRecordInfo

`func (o *DriverOptionsModel) HasIncludeDriverRecordInfo() bool`

HasIncludeDriverRecordInfo returns a boolean if a field has been set.

### GetIncludeDriverMetaData

`func (o *DriverOptionsModel) GetIncludeDriverMetaData() bool`

GetIncludeDriverMetaData returns the IncludeDriverMetaData field if non-nil, zero value otherwise.

### GetIncludeDriverMetaDataOk

`func (o *DriverOptionsModel) GetIncludeDriverMetaDataOk() (*bool, bool)`

GetIncludeDriverMetaDataOk returns a tuple with the IncludeDriverMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverMetaData

`func (o *DriverOptionsModel) SetIncludeDriverMetaData(v bool)`

SetIncludeDriverMetaData sets IncludeDriverMetaData field to given value.

### HasIncludeDriverMetaData

`func (o *DriverOptionsModel) HasIncludeDriverMetaData() bool`

HasIncludeDriverMetaData returns a boolean if a field has been set.

### GetIncludeDriverTags

`func (o *DriverOptionsModel) GetIncludeDriverTags() bool`

GetIncludeDriverTags returns the IncludeDriverTags field if non-nil, zero value otherwise.

### GetIncludeDriverTagsOk

`func (o *DriverOptionsModel) GetIncludeDriverTagsOk() (*bool, bool)`

GetIncludeDriverTagsOk returns a tuple with the IncludeDriverTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverTags

`func (o *DriverOptionsModel) SetIncludeDriverTags(v bool)`

SetIncludeDriverTags sets IncludeDriverTags field to given value.

### HasIncludeDriverTags

`func (o *DriverOptionsModel) HasIncludeDriverTags() bool`

HasIncludeDriverTags returns a boolean if a field has been set.

### GetIncludeDriverTagTypeNames

`func (o *DriverOptionsModel) GetIncludeDriverTagTypeNames() bool`

GetIncludeDriverTagTypeNames returns the IncludeDriverTagTypeNames field if non-nil, zero value otherwise.

### GetIncludeDriverTagTypeNamesOk

`func (o *DriverOptionsModel) GetIncludeDriverTagTypeNamesOk() (*bool, bool)`

GetIncludeDriverTagTypeNamesOk returns a tuple with the IncludeDriverTagTypeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverTagTypeNames

`func (o *DriverOptionsModel) SetIncludeDriverTagTypeNames(v bool)`

SetIncludeDriverTagTypeNames sets IncludeDriverTagTypeNames field to given value.

### HasIncludeDriverTagTypeNames

`func (o *DriverOptionsModel) HasIncludeDriverTagTypeNames() bool`

HasIncludeDriverTagTypeNames returns a boolean if a field has been set.

### GetIncludeDriverLinks

`func (o *DriverOptionsModel) GetIncludeDriverLinks() bool`

GetIncludeDriverLinks returns the IncludeDriverLinks field if non-nil, zero value otherwise.

### GetIncludeDriverLinksOk

`func (o *DriverOptionsModel) GetIncludeDriverLinksOk() (*bool, bool)`

GetIncludeDriverLinksOk returns a tuple with the IncludeDriverLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverLinks

`func (o *DriverOptionsModel) SetIncludeDriverLinks(v bool)`

SetIncludeDriverLinks sets IncludeDriverLinks field to given value.

### HasIncludeDriverLinks

`func (o *DriverOptionsModel) HasIncludeDriverLinks() bool`

HasIncludeDriverLinks returns a boolean if a field has been set.

### GetIncludeDriverNotes

`func (o *DriverOptionsModel) GetIncludeDriverNotes() bool`

GetIncludeDriverNotes returns the IncludeDriverNotes field if non-nil, zero value otherwise.

### GetIncludeDriverNotesOk

`func (o *DriverOptionsModel) GetIncludeDriverNotesOk() (*bool, bool)`

GetIncludeDriverNotesOk returns a tuple with the IncludeDriverNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverNotes

`func (o *DriverOptionsModel) SetIncludeDriverNotes(v bool)`

SetIncludeDriverNotes sets IncludeDriverNotes field to given value.

### HasIncludeDriverNotes

`func (o *DriverOptionsModel) HasIncludeDriverNotes() bool`

HasIncludeDriverNotes returns a boolean if a field has been set.

### GetIncludeDriverNoteTags

`func (o *DriverOptionsModel) GetIncludeDriverNoteTags() bool`

GetIncludeDriverNoteTags returns the IncludeDriverNoteTags field if non-nil, zero value otherwise.

### GetIncludeDriverNoteTagsOk

`func (o *DriverOptionsModel) GetIncludeDriverNoteTagsOk() (*bool, bool)`

GetIncludeDriverNoteTagsOk returns a tuple with the IncludeDriverNoteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverNoteTags

`func (o *DriverOptionsModel) SetIncludeDriverNoteTags(v bool)`

SetIncludeDriverNoteTags sets IncludeDriverNoteTags field to given value.

### HasIncludeDriverNoteTags

`func (o *DriverOptionsModel) HasIncludeDriverNoteTags() bool`

HasIncludeDriverNoteTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


