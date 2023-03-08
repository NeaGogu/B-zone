# NotificationFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**ObjectType** | Pointer to **[]int32** | Object type IDs | [optional] 
**ObjectTypeName** | Pointer to **[]string** | Object type names | [optional] 
**NotificationCategoryId** | Pointer to **[]int32** | Notification category id | [optional] 
**NotificationCategoryName** | Pointer to **[]string** | Notification category name | [optional] 
**Content** | Pointer to **[]string** | Notification content | [optional] 

## Methods

### NewNotificationFiltersModel

`func NewNotificationFiltersModel() *NotificationFiltersModel`

NewNotificationFiltersModel instantiates a new NotificationFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationFiltersModelWithDefaults

`func NewNotificationFiltersModelWithDefaults() *NotificationFiltersModel`

NewNotificationFiltersModelWithDefaults instantiates a new NotificationFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAtSince

`func (o *NotificationFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *NotificationFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *NotificationFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *NotificationFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *NotificationFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *NotificationFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *NotificationFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *NotificationFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetObjectType

`func (o *NotificationFiltersModel) GetObjectType() []int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationFiltersModel) GetObjectTypeOk() (*[]int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationFiltersModel) SetObjectType(v []int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *NotificationFiltersModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *NotificationFiltersModel) GetObjectTypeName() []string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *NotificationFiltersModel) GetObjectTypeNameOk() (*[]string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *NotificationFiltersModel) SetObjectTypeName(v []string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *NotificationFiltersModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetNotificationCategoryId

`func (o *NotificationFiltersModel) GetNotificationCategoryId() []int32`

GetNotificationCategoryId returns the NotificationCategoryId field if non-nil, zero value otherwise.

### GetNotificationCategoryIdOk

`func (o *NotificationFiltersModel) GetNotificationCategoryIdOk() (*[]int32, bool)`

GetNotificationCategoryIdOk returns a tuple with the NotificationCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCategoryId

`func (o *NotificationFiltersModel) SetNotificationCategoryId(v []int32)`

SetNotificationCategoryId sets NotificationCategoryId field to given value.

### HasNotificationCategoryId

`func (o *NotificationFiltersModel) HasNotificationCategoryId() bool`

HasNotificationCategoryId returns a boolean if a field has been set.

### GetNotificationCategoryName

`func (o *NotificationFiltersModel) GetNotificationCategoryName() []string`

GetNotificationCategoryName returns the NotificationCategoryName field if non-nil, zero value otherwise.

### GetNotificationCategoryNameOk

`func (o *NotificationFiltersModel) GetNotificationCategoryNameOk() (*[]string, bool)`

GetNotificationCategoryNameOk returns a tuple with the NotificationCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCategoryName

`func (o *NotificationFiltersModel) SetNotificationCategoryName(v []string)`

SetNotificationCategoryName sets NotificationCategoryName field to given value.

### HasNotificationCategoryName

`func (o *NotificationFiltersModel) HasNotificationCategoryName() bool`

HasNotificationCategoryName returns a boolean if a field has been set.

### GetContent

`func (o *NotificationFiltersModel) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NotificationFiltersModel) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NotificationFiltersModel) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *NotificationFiltersModel) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


