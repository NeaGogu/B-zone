# UserNotificationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**UserId** | Pointer to **int64** | User Identifier | [optional] [readonly] 
**NotificationId** | Pointer to **int64** | Notification Identifier | [optional] [readonly] 
**Unread** | Pointer to **int64** | Unread status | [optional] [readonly] 
**Pinned** | Pointer to **int64** | Pinned status | [optional] [readonly] 
**Archived** | Pointer to **int64** | Archived status | [optional] [readonly] 
**ArchiveDate** | Pointer to **time.Time** | Archive date | [optional] [readonly] 
**NotificationCreatedAt** | Pointer to **time.Time** | Creation date | [optional] [readonly] 
**ObjectId** | Pointer to **int64** | Object Identifier | [optional] [readonly] 
**ObjectType** | Pointer to **int64** | Object type ID | [optional] 
**NotificationUpdatedAt** | Pointer to **time.Time** | Modification date | [optional] [readonly] 
**ObjectTypeName** | Pointer to **string** | Object type name | [optional] 

## Methods

### NewUserNotificationModel

`func NewUserNotificationModel() *UserNotificationModel`

NewUserNotificationModel instantiates a new UserNotificationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNotificationModelWithDefaults

`func NewUserNotificationModelWithDefaults() *UserNotificationModel`

NewUserNotificationModelWithDefaults instantiates a new UserNotificationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserNotificationModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserNotificationModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserNotificationModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserNotificationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *UserNotificationModel) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserNotificationModel) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserNotificationModel) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserNotificationModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNotificationId

`func (o *UserNotificationModel) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *UserNotificationModel) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *UserNotificationModel) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *UserNotificationModel) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetUnread

`func (o *UserNotificationModel) GetUnread() int64`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *UserNotificationModel) GetUnreadOk() (*int64, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *UserNotificationModel) SetUnread(v int64)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *UserNotificationModel) HasUnread() bool`

HasUnread returns a boolean if a field has been set.

### GetPinned

`func (o *UserNotificationModel) GetPinned() int64`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *UserNotificationModel) GetPinnedOk() (*int64, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *UserNotificationModel) SetPinned(v int64)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *UserNotificationModel) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetArchived

`func (o *UserNotificationModel) GetArchived() int64`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UserNotificationModel) GetArchivedOk() (*int64, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UserNotificationModel) SetArchived(v int64)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UserNotificationModel) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchiveDate

`func (o *UserNotificationModel) GetArchiveDate() time.Time`

GetArchiveDate returns the ArchiveDate field if non-nil, zero value otherwise.

### GetArchiveDateOk

`func (o *UserNotificationModel) GetArchiveDateOk() (*time.Time, bool)`

GetArchiveDateOk returns a tuple with the ArchiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveDate

`func (o *UserNotificationModel) SetArchiveDate(v time.Time)`

SetArchiveDate sets ArchiveDate field to given value.

### HasArchiveDate

`func (o *UserNotificationModel) HasArchiveDate() bool`

HasArchiveDate returns a boolean if a field has been set.

### GetNotificationCreatedAt

`func (o *UserNotificationModel) GetNotificationCreatedAt() time.Time`

GetNotificationCreatedAt returns the NotificationCreatedAt field if non-nil, zero value otherwise.

### GetNotificationCreatedAtOk

`func (o *UserNotificationModel) GetNotificationCreatedAtOk() (*time.Time, bool)`

GetNotificationCreatedAtOk returns a tuple with the NotificationCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCreatedAt

`func (o *UserNotificationModel) SetNotificationCreatedAt(v time.Time)`

SetNotificationCreatedAt sets NotificationCreatedAt field to given value.

### HasNotificationCreatedAt

`func (o *UserNotificationModel) HasNotificationCreatedAt() bool`

HasNotificationCreatedAt returns a boolean if a field has been set.

### GetObjectId

`func (o *UserNotificationModel) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *UserNotificationModel) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *UserNotificationModel) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *UserNotificationModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *UserNotificationModel) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UserNotificationModel) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UserNotificationModel) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *UserNotificationModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetNotificationUpdatedAt

`func (o *UserNotificationModel) GetNotificationUpdatedAt() time.Time`

GetNotificationUpdatedAt returns the NotificationUpdatedAt field if non-nil, zero value otherwise.

### GetNotificationUpdatedAtOk

`func (o *UserNotificationModel) GetNotificationUpdatedAtOk() (*time.Time, bool)`

GetNotificationUpdatedAtOk returns a tuple with the NotificationUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUpdatedAt

`func (o *UserNotificationModel) SetNotificationUpdatedAt(v time.Time)`

SetNotificationUpdatedAt sets NotificationUpdatedAt field to given value.

### HasNotificationUpdatedAt

`func (o *UserNotificationModel) HasNotificationUpdatedAt() bool`

HasNotificationUpdatedAt returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *UserNotificationModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *UserNotificationModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *UserNotificationModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *UserNotificationModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


