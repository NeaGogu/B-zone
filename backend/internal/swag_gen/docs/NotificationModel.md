# NotificationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**ObjectType** | Pointer to **int64** | Object type ID | [optional] 
**Title** | Pointer to **string** | Notification title | [optional] 
**ObjectTypeName** | Pointer to **string** | Object type name | [optional] 
**ObjectId** | Pointer to **int32** | Object ID | [optional] 
**NotificationCategoryId** | Pointer to **int32** | Notification category id | [optional] 
**NotificationCategoryName** | Pointer to **string** | Notification category name | [optional] 
**UpdatedByName** | Pointer to **string** | Notification updated by user full name | [optional] 
**Content** | Pointer to **string** | Notification content | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: note has been removed and is no longer visible in any bumbal interface | [optional] 
**NotificationCreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**NotificationUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**NotificationCreatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 
**NotificationUpdatedBy** | Pointer to **int32** | updated_by user id | [optional] [readonly] 
**Users** | Pointer to [**[]UserNotificationModel**](UserNotificationModel.md) |  | [optional] 
**RoleNames** | Pointer to **[]string** | Roles to enable notification for, works only on create, ignored on update | [optional] 

## Methods

### NewNotificationModel

`func NewNotificationModel() *NotificationModel`

NewNotificationModel instantiates a new NotificationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationModelWithDefaults

`func NewNotificationModelWithDefaults() *NotificationModel`

NewNotificationModelWithDefaults instantiates a new NotificationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *NotificationModel) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationModel) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationModel) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *NotificationModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetTitle

`func (o *NotificationModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NotificationModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *NotificationModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *NotificationModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *NotificationModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *NotificationModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetObjectId

`func (o *NotificationModel) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *NotificationModel) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *NotificationModel) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *NotificationModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetNotificationCategoryId

`func (o *NotificationModel) GetNotificationCategoryId() int32`

GetNotificationCategoryId returns the NotificationCategoryId field if non-nil, zero value otherwise.

### GetNotificationCategoryIdOk

`func (o *NotificationModel) GetNotificationCategoryIdOk() (*int32, bool)`

GetNotificationCategoryIdOk returns a tuple with the NotificationCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCategoryId

`func (o *NotificationModel) SetNotificationCategoryId(v int32)`

SetNotificationCategoryId sets NotificationCategoryId field to given value.

### HasNotificationCategoryId

`func (o *NotificationModel) HasNotificationCategoryId() bool`

HasNotificationCategoryId returns a boolean if a field has been set.

### GetNotificationCategoryName

`func (o *NotificationModel) GetNotificationCategoryName() string`

GetNotificationCategoryName returns the NotificationCategoryName field if non-nil, zero value otherwise.

### GetNotificationCategoryNameOk

`func (o *NotificationModel) GetNotificationCategoryNameOk() (*string, bool)`

GetNotificationCategoryNameOk returns a tuple with the NotificationCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCategoryName

`func (o *NotificationModel) SetNotificationCategoryName(v string)`

SetNotificationCategoryName sets NotificationCategoryName field to given value.

### HasNotificationCategoryName

`func (o *NotificationModel) HasNotificationCategoryName() bool`

HasNotificationCategoryName returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *NotificationModel) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *NotificationModel) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *NotificationModel) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *NotificationModel) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.

### GetContent

`func (o *NotificationModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NotificationModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NotificationModel) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *NotificationModel) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetActive

`func (o *NotificationModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NotificationModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NotificationModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NotificationModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotificationCreatedAt

`func (o *NotificationModel) GetNotificationCreatedAt() time.Time`

GetNotificationCreatedAt returns the NotificationCreatedAt field if non-nil, zero value otherwise.

### GetNotificationCreatedAtOk

`func (o *NotificationModel) GetNotificationCreatedAtOk() (*time.Time, bool)`

GetNotificationCreatedAtOk returns a tuple with the NotificationCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCreatedAt

`func (o *NotificationModel) SetNotificationCreatedAt(v time.Time)`

SetNotificationCreatedAt sets NotificationCreatedAt field to given value.

### HasNotificationCreatedAt

`func (o *NotificationModel) HasNotificationCreatedAt() bool`

HasNotificationCreatedAt returns a boolean if a field has been set.

### GetNotificationUpdatedAt

`func (o *NotificationModel) GetNotificationUpdatedAt() time.Time`

GetNotificationUpdatedAt returns the NotificationUpdatedAt field if non-nil, zero value otherwise.

### GetNotificationUpdatedAtOk

`func (o *NotificationModel) GetNotificationUpdatedAtOk() (*time.Time, bool)`

GetNotificationUpdatedAtOk returns a tuple with the NotificationUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUpdatedAt

`func (o *NotificationModel) SetNotificationUpdatedAt(v time.Time)`

SetNotificationUpdatedAt sets NotificationUpdatedAt field to given value.

### HasNotificationUpdatedAt

`func (o *NotificationModel) HasNotificationUpdatedAt() bool`

HasNotificationUpdatedAt returns a boolean if a field has been set.

### GetNotificationCreatedBy

`func (o *NotificationModel) GetNotificationCreatedBy() int32`

GetNotificationCreatedBy returns the NotificationCreatedBy field if non-nil, zero value otherwise.

### GetNotificationCreatedByOk

`func (o *NotificationModel) GetNotificationCreatedByOk() (*int32, bool)`

GetNotificationCreatedByOk returns a tuple with the NotificationCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCreatedBy

`func (o *NotificationModel) SetNotificationCreatedBy(v int32)`

SetNotificationCreatedBy sets NotificationCreatedBy field to given value.

### HasNotificationCreatedBy

`func (o *NotificationModel) HasNotificationCreatedBy() bool`

HasNotificationCreatedBy returns a boolean if a field has been set.

### GetNotificationUpdatedBy

`func (o *NotificationModel) GetNotificationUpdatedBy() int32`

GetNotificationUpdatedBy returns the NotificationUpdatedBy field if non-nil, zero value otherwise.

### GetNotificationUpdatedByOk

`func (o *NotificationModel) GetNotificationUpdatedByOk() (*int32, bool)`

GetNotificationUpdatedByOk returns a tuple with the NotificationUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUpdatedBy

`func (o *NotificationModel) SetNotificationUpdatedBy(v int32)`

SetNotificationUpdatedBy sets NotificationUpdatedBy field to given value.

### HasNotificationUpdatedBy

`func (o *NotificationModel) HasNotificationUpdatedBy() bool`

HasNotificationUpdatedBy returns a boolean if a field has been set.

### GetUsers

`func (o *NotificationModel) GetUsers() []UserNotificationModel`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *NotificationModel) GetUsersOk() (*[]UserNotificationModel, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *NotificationModel) SetUsers(v []UserNotificationModel)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *NotificationModel) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetRoleNames

`func (o *NotificationModel) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *NotificationModel) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *NotificationModel) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *NotificationModel) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


