# UserNotificationFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | UserNotification id&#39;s | [optional] 
**UserId** | Pointer to **[]int32** | User id&#39;s | [optional] 
**NotificationId** | Pointer to **[]int32** | Notification id&#39;s | [optional] 
**Unread** | Pointer to **int32** | Unread status | [optional] 
**Pinned** | Pointer to **int32** | Pinned status | [optional] 
**Archived** | Pointer to **int32** | Archived status | [optional] 
**NotificationCategoryId** | Pointer to **[]int32** | Notification category id | [optional] 
**NotificationCategoryName** | Pointer to **[]string** | Notification category name | [optional] 
**ObjectType** | Pointer to **[]int32** | Object type ID | [optional] 
**ObjectTypeName** | Pointer to **[]string** | Object type name | [optional] 

## Methods

### NewUserNotificationFiltersModel

`func NewUserNotificationFiltersModel() *UserNotificationFiltersModel`

NewUserNotificationFiltersModel instantiates a new UserNotificationFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNotificationFiltersModelWithDefaults

`func NewUserNotificationFiltersModelWithDefaults() *UserNotificationFiltersModel`

NewUserNotificationFiltersModelWithDefaults instantiates a new UserNotificationFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserNotificationFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserNotificationFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserNotificationFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserNotificationFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *UserNotificationFiltersModel) GetUserId() []int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserNotificationFiltersModel) GetUserIdOk() (*[]int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserNotificationFiltersModel) SetUserId(v []int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserNotificationFiltersModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetNotificationId

`func (o *UserNotificationFiltersModel) GetNotificationId() []int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *UserNotificationFiltersModel) GetNotificationIdOk() (*[]int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *UserNotificationFiltersModel) SetNotificationId(v []int32)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *UserNotificationFiltersModel) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetUnread

`func (o *UserNotificationFiltersModel) GetUnread() int32`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *UserNotificationFiltersModel) GetUnreadOk() (*int32, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *UserNotificationFiltersModel) SetUnread(v int32)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *UserNotificationFiltersModel) HasUnread() bool`

HasUnread returns a boolean if a field has been set.

### GetPinned

`func (o *UserNotificationFiltersModel) GetPinned() int32`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *UserNotificationFiltersModel) GetPinnedOk() (*int32, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *UserNotificationFiltersModel) SetPinned(v int32)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *UserNotificationFiltersModel) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetArchived

`func (o *UserNotificationFiltersModel) GetArchived() int32`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UserNotificationFiltersModel) GetArchivedOk() (*int32, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UserNotificationFiltersModel) SetArchived(v int32)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UserNotificationFiltersModel) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetNotificationCategoryId

`func (o *UserNotificationFiltersModel) GetNotificationCategoryId() []int32`

GetNotificationCategoryId returns the NotificationCategoryId field if non-nil, zero value otherwise.

### GetNotificationCategoryIdOk

`func (o *UserNotificationFiltersModel) GetNotificationCategoryIdOk() (*[]int32, bool)`

GetNotificationCategoryIdOk returns a tuple with the NotificationCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCategoryId

`func (o *UserNotificationFiltersModel) SetNotificationCategoryId(v []int32)`

SetNotificationCategoryId sets NotificationCategoryId field to given value.

### HasNotificationCategoryId

`func (o *UserNotificationFiltersModel) HasNotificationCategoryId() bool`

HasNotificationCategoryId returns a boolean if a field has been set.

### GetNotificationCategoryName

`func (o *UserNotificationFiltersModel) GetNotificationCategoryName() []string`

GetNotificationCategoryName returns the NotificationCategoryName field if non-nil, zero value otherwise.

### GetNotificationCategoryNameOk

`func (o *UserNotificationFiltersModel) GetNotificationCategoryNameOk() (*[]string, bool)`

GetNotificationCategoryNameOk returns a tuple with the NotificationCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCategoryName

`func (o *UserNotificationFiltersModel) SetNotificationCategoryName(v []string)`

SetNotificationCategoryName sets NotificationCategoryName field to given value.

### HasNotificationCategoryName

`func (o *UserNotificationFiltersModel) HasNotificationCategoryName() bool`

HasNotificationCategoryName returns a boolean if a field has been set.

### GetObjectType

`func (o *UserNotificationFiltersModel) GetObjectType() []int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UserNotificationFiltersModel) GetObjectTypeOk() (*[]int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UserNotificationFiltersModel) SetObjectType(v []int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *UserNotificationFiltersModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *UserNotificationFiltersModel) GetObjectTypeName() []string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *UserNotificationFiltersModel) GetObjectTypeNameOk() (*[]string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *UserNotificationFiltersModel) SetObjectTypeName(v []string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *UserNotificationFiltersModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


