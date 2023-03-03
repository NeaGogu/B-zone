# AssignmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID of this Assignment | [optional] [readonly] 
**PartyId** | Pointer to **int64** | Party ID | [optional] 
**BookingAccountId** | Pointer to **int64** | Booking account ID | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**TagIds** | Pointer to **[]int32** | Tag ids | [optional] 
**Activities** | Pointer to [**[]ActivityModel**](ActivityModel.md) |  | [optional] [readonly] 
**BookingAccount** | Pointer to [**PartyModel**](PartyModel.md) |  | [optional] 
**StatusId** | Pointer to **int64** | Status ID of this Assignment | [optional] 
**Nr** | Pointer to **string** | Non-Unique number of this Assignment | [optional] 
**PartyLink** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**AccountName** | Pointer to **string** | Account Name associated with this Assignment | [optional] 
**PartyName** | Pointer to **string** | Party Name associated with this Assignment | [optional] 
**Reference** | Pointer to **string** | Reference text of this Assignment | [optional] 
**Description** | Pointer to **string** | Description text of this Assignment | [optional] 
**Remarks** | Pointer to **string** | Remarks about this Assignment | [optional] 
**DateTimeFrom** | Pointer to **time.Time** | Earliest start time of all Activities is this Assignment | [optional] [readonly] 
**DateTimeTo** | Pointer to **time.Time** | Latest end time of all Activities is this Assignment | [optional] [readonly] 
**MultiDay** | Pointer to **bool** | Multi day type assignment | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Notes** | Pointer to [**[]NoteModel**](NoteModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**AssignmentCreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**AssignmentUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**AssignmentCreatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 
**AssignmentUpdatedBy** | Pointer to **int32** | updated_by user id | [optional] [readonly] 
**AssignmentCreatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**AssignmentUpdatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**AssignmentActive** | Pointer to **bool** | Assignment is active (&#x3D;true). Inactive assignments are not automatically considered in any of the application algorithms and will not be shown in the Bumbal Gui. | [optional] 
**AssignmentRemoved** | Pointer to **bool** | Assignment is removed (&#x3D;true). Removed assignments are not automatically considered in any of the application algorithms and will not be shown in the Bumbal Gui. Removed assignments are usually irrepairable. | [optional] 

## Methods

### NewAssignmentModel

`func NewAssignmentModel() *AssignmentModel`

NewAssignmentModel instantiates a new AssignmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentModelWithDefaults

`func NewAssignmentModelWithDefaults() *AssignmentModel`

NewAssignmentModelWithDefaults instantiates a new AssignmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssignmentModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignmentModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignmentModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AssignmentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartyId

`func (o *AssignmentModel) GetPartyId() int64`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *AssignmentModel) GetPartyIdOk() (*int64, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *AssignmentModel) SetPartyId(v int64)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *AssignmentModel) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetBookingAccountId

`func (o *AssignmentModel) GetBookingAccountId() int64`

GetBookingAccountId returns the BookingAccountId field if non-nil, zero value otherwise.

### GetBookingAccountIdOk

`func (o *AssignmentModel) GetBookingAccountIdOk() (*int64, bool)`

GetBookingAccountIdOk returns a tuple with the BookingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingAccountId

`func (o *AssignmentModel) SetBookingAccountId(v int64)`

SetBookingAccountId sets BookingAccountId field to given value.

### HasBookingAccountId

`func (o *AssignmentModel) HasBookingAccountId() bool`

HasBookingAccountId returns a boolean if a field has been set.

### GetTagNames

`func (o *AssignmentModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *AssignmentModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *AssignmentModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *AssignmentModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetTagIds

`func (o *AssignmentModel) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *AssignmentModel) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *AssignmentModel) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *AssignmentModel) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetActivities

`func (o *AssignmentModel) GetActivities() []ActivityModel`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *AssignmentModel) GetActivitiesOk() (*[]ActivityModel, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *AssignmentModel) SetActivities(v []ActivityModel)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *AssignmentModel) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetBookingAccount

`func (o *AssignmentModel) GetBookingAccount() PartyModel`

GetBookingAccount returns the BookingAccount field if non-nil, zero value otherwise.

### GetBookingAccountOk

`func (o *AssignmentModel) GetBookingAccountOk() (*PartyModel, bool)`

GetBookingAccountOk returns a tuple with the BookingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingAccount

`func (o *AssignmentModel) SetBookingAccount(v PartyModel)`

SetBookingAccount sets BookingAccount field to given value.

### HasBookingAccount

`func (o *AssignmentModel) HasBookingAccount() bool`

HasBookingAccount returns a boolean if a field has been set.

### GetStatusId

`func (o *AssignmentModel) GetStatusId() int64`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *AssignmentModel) GetStatusIdOk() (*int64, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *AssignmentModel) SetStatusId(v int64)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *AssignmentModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetNr

`func (o *AssignmentModel) GetNr() string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *AssignmentModel) GetNrOk() (*string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *AssignmentModel) SetNr(v string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *AssignmentModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetPartyLink

`func (o *AssignmentModel) GetPartyLink() []LinkModel`

GetPartyLink returns the PartyLink field if non-nil, zero value otherwise.

### GetPartyLinkOk

`func (o *AssignmentModel) GetPartyLinkOk() (*[]LinkModel, bool)`

GetPartyLinkOk returns a tuple with the PartyLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyLink

`func (o *AssignmentModel) SetPartyLink(v []LinkModel)`

SetPartyLink sets PartyLink field to given value.

### HasPartyLink

`func (o *AssignmentModel) HasPartyLink() bool`

HasPartyLink returns a boolean if a field has been set.

### GetAccountName

`func (o *AssignmentModel) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AssignmentModel) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AssignmentModel) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AssignmentModel) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetPartyName

`func (o *AssignmentModel) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *AssignmentModel) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *AssignmentModel) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *AssignmentModel) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.

### GetReference

`func (o *AssignmentModel) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AssignmentModel) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AssignmentModel) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AssignmentModel) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDescription

`func (o *AssignmentModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssignmentModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssignmentModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssignmentModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRemarks

`func (o *AssignmentModel) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *AssignmentModel) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *AssignmentModel) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *AssignmentModel) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *AssignmentModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *AssignmentModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *AssignmentModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *AssignmentModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *AssignmentModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *AssignmentModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *AssignmentModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *AssignmentModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetMultiDay

`func (o *AssignmentModel) GetMultiDay() bool`

GetMultiDay returns the MultiDay field if non-nil, zero value otherwise.

### GetMultiDayOk

`func (o *AssignmentModel) GetMultiDayOk() (*bool, bool)`

GetMultiDayOk returns a tuple with the MultiDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiDay

`func (o *AssignmentModel) SetMultiDay(v bool)`

SetMultiDay sets MultiDay field to given value.

### HasMultiDay

`func (o *AssignmentModel) HasMultiDay() bool`

HasMultiDay returns a boolean if a field has been set.

### GetLinks

`func (o *AssignmentModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AssignmentModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AssignmentModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AssignmentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *AssignmentModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AssignmentModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AssignmentModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AssignmentModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetNotes

`func (o *AssignmentModel) GetNotes() []NoteModel`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AssignmentModel) GetNotesOk() (*[]NoteModel, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AssignmentModel) SetNotes(v []NoteModel)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AssignmentModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFiles

`func (o *AssignmentModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AssignmentModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AssignmentModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AssignmentModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetAssignmentCreatedAt

`func (o *AssignmentModel) GetAssignmentCreatedAt() time.Time`

GetAssignmentCreatedAt returns the AssignmentCreatedAt field if non-nil, zero value otherwise.

### GetAssignmentCreatedAtOk

`func (o *AssignmentModel) GetAssignmentCreatedAtOk() (*time.Time, bool)`

GetAssignmentCreatedAtOk returns a tuple with the AssignmentCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentCreatedAt

`func (o *AssignmentModel) SetAssignmentCreatedAt(v time.Time)`

SetAssignmentCreatedAt sets AssignmentCreatedAt field to given value.

### HasAssignmentCreatedAt

`func (o *AssignmentModel) HasAssignmentCreatedAt() bool`

HasAssignmentCreatedAt returns a boolean if a field has been set.

### GetAssignmentUpdatedAt

`func (o *AssignmentModel) GetAssignmentUpdatedAt() time.Time`

GetAssignmentUpdatedAt returns the AssignmentUpdatedAt field if non-nil, zero value otherwise.

### GetAssignmentUpdatedAtOk

`func (o *AssignmentModel) GetAssignmentUpdatedAtOk() (*time.Time, bool)`

GetAssignmentUpdatedAtOk returns a tuple with the AssignmentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentUpdatedAt

`func (o *AssignmentModel) SetAssignmentUpdatedAt(v time.Time)`

SetAssignmentUpdatedAt sets AssignmentUpdatedAt field to given value.

### HasAssignmentUpdatedAt

`func (o *AssignmentModel) HasAssignmentUpdatedAt() bool`

HasAssignmentUpdatedAt returns a boolean if a field has been set.

### GetAssignmentCreatedBy

`func (o *AssignmentModel) GetAssignmentCreatedBy() int32`

GetAssignmentCreatedBy returns the AssignmentCreatedBy field if non-nil, zero value otherwise.

### GetAssignmentCreatedByOk

`func (o *AssignmentModel) GetAssignmentCreatedByOk() (*int32, bool)`

GetAssignmentCreatedByOk returns a tuple with the AssignmentCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentCreatedBy

`func (o *AssignmentModel) SetAssignmentCreatedBy(v int32)`

SetAssignmentCreatedBy sets AssignmentCreatedBy field to given value.

### HasAssignmentCreatedBy

`func (o *AssignmentModel) HasAssignmentCreatedBy() bool`

HasAssignmentCreatedBy returns a boolean if a field has been set.

### GetAssignmentUpdatedBy

`func (o *AssignmentModel) GetAssignmentUpdatedBy() int32`

GetAssignmentUpdatedBy returns the AssignmentUpdatedBy field if non-nil, zero value otherwise.

### GetAssignmentUpdatedByOk

`func (o *AssignmentModel) GetAssignmentUpdatedByOk() (*int32, bool)`

GetAssignmentUpdatedByOk returns a tuple with the AssignmentUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentUpdatedBy

`func (o *AssignmentModel) SetAssignmentUpdatedBy(v int32)`

SetAssignmentUpdatedBy sets AssignmentUpdatedBy field to given value.

### HasAssignmentUpdatedBy

`func (o *AssignmentModel) HasAssignmentUpdatedBy() bool`

HasAssignmentUpdatedBy returns a boolean if a field has been set.

### GetAssignmentCreatedByUser

`func (o *AssignmentModel) GetAssignmentCreatedByUser() UsersModel`

GetAssignmentCreatedByUser returns the AssignmentCreatedByUser field if non-nil, zero value otherwise.

### GetAssignmentCreatedByUserOk

`func (o *AssignmentModel) GetAssignmentCreatedByUserOk() (*UsersModel, bool)`

GetAssignmentCreatedByUserOk returns a tuple with the AssignmentCreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentCreatedByUser

`func (o *AssignmentModel) SetAssignmentCreatedByUser(v UsersModel)`

SetAssignmentCreatedByUser sets AssignmentCreatedByUser field to given value.

### HasAssignmentCreatedByUser

`func (o *AssignmentModel) HasAssignmentCreatedByUser() bool`

HasAssignmentCreatedByUser returns a boolean if a field has been set.

### GetAssignmentUpdatedByUser

`func (o *AssignmentModel) GetAssignmentUpdatedByUser() UsersModel`

GetAssignmentUpdatedByUser returns the AssignmentUpdatedByUser field if non-nil, zero value otherwise.

### GetAssignmentUpdatedByUserOk

`func (o *AssignmentModel) GetAssignmentUpdatedByUserOk() (*UsersModel, bool)`

GetAssignmentUpdatedByUserOk returns a tuple with the AssignmentUpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentUpdatedByUser

`func (o *AssignmentModel) SetAssignmentUpdatedByUser(v UsersModel)`

SetAssignmentUpdatedByUser sets AssignmentUpdatedByUser field to given value.

### HasAssignmentUpdatedByUser

`func (o *AssignmentModel) HasAssignmentUpdatedByUser() bool`

HasAssignmentUpdatedByUser returns a boolean if a field has been set.

### GetAssignmentActive

`func (o *AssignmentModel) GetAssignmentActive() bool`

GetAssignmentActive returns the AssignmentActive field if non-nil, zero value otherwise.

### GetAssignmentActiveOk

`func (o *AssignmentModel) GetAssignmentActiveOk() (*bool, bool)`

GetAssignmentActiveOk returns a tuple with the AssignmentActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentActive

`func (o *AssignmentModel) SetAssignmentActive(v bool)`

SetAssignmentActive sets AssignmentActive field to given value.

### HasAssignmentActive

`func (o *AssignmentModel) HasAssignmentActive() bool`

HasAssignmentActive returns a boolean if a field has been set.

### GetAssignmentRemoved

`func (o *AssignmentModel) GetAssignmentRemoved() bool`

GetAssignmentRemoved returns the AssignmentRemoved field if non-nil, zero value otherwise.

### GetAssignmentRemovedOk

`func (o *AssignmentModel) GetAssignmentRemovedOk() (*bool, bool)`

GetAssignmentRemovedOk returns a tuple with the AssignmentRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRemoved

`func (o *AssignmentModel) SetAssignmentRemoved(v bool)`

SetAssignmentRemoved sets AssignmentRemoved field to given value.

### HasAssignmentRemoved

`func (o *AssignmentModel) HasAssignmentRemoved() bool`

HasAssignmentRemoved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


