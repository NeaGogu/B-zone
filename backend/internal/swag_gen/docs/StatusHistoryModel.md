# StatusHistoryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**OldStatusId** | Pointer to **int64** | Service windows scheme ID | [optional] 
**NewStatusId** | Pointer to **int64** | Service windows scheme ID | [optional] 
**ObjectType** | Pointer to **int64** | Service windows scheme ID | [optional] 
**ObjectId** | Pointer to **int64** | Service windows scheme ID | [optional] 
**RecordSnapshot** | Pointer to **string** | Service windows scheme ID | [optional] 
**StatusHistoryCreatedBy** | Pointer to **string** | Service windows scheme ID | [optional] 
**StatusHistoryCreatedAt** | Pointer to **string** | Service windows scheme ID | [optional] 
**UserFirstName** | Pointer to **string** | user first name | [optional] 
**UserLastName** | Pointer to **string** | user last name | [optional] 
**UserNamePrefix** | Pointer to **string** | User name prefix | [optional] 
**UserFullName** | Pointer to **string** | User full name | [optional] 
**UserEmail** | Pointer to **string** | User email | [optional] 
**UserUuid** | Pointer to **string** | user uuid | [optional] 
**OldStatusName** | Pointer to **string** | old status name | [optional] 
**NewStatusName** | Pointer to **string** | new status name | [optional] 

## Methods

### NewStatusHistoryModel

`func NewStatusHistoryModel() *StatusHistoryModel`

NewStatusHistoryModel instantiates a new StatusHistoryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusHistoryModelWithDefaults

`func NewStatusHistoryModelWithDefaults() *StatusHistoryModel`

NewStatusHistoryModelWithDefaults instantiates a new StatusHistoryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StatusHistoryModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatusHistoryModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatusHistoryModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StatusHistoryModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOldStatusId

`func (o *StatusHistoryModel) GetOldStatusId() int64`

GetOldStatusId returns the OldStatusId field if non-nil, zero value otherwise.

### GetOldStatusIdOk

`func (o *StatusHistoryModel) GetOldStatusIdOk() (*int64, bool)`

GetOldStatusIdOk returns a tuple with the OldStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldStatusId

`func (o *StatusHistoryModel) SetOldStatusId(v int64)`

SetOldStatusId sets OldStatusId field to given value.

### HasOldStatusId

`func (o *StatusHistoryModel) HasOldStatusId() bool`

HasOldStatusId returns a boolean if a field has been set.

### GetNewStatusId

`func (o *StatusHistoryModel) GetNewStatusId() int64`

GetNewStatusId returns the NewStatusId field if non-nil, zero value otherwise.

### GetNewStatusIdOk

`func (o *StatusHistoryModel) GetNewStatusIdOk() (*int64, bool)`

GetNewStatusIdOk returns a tuple with the NewStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatusId

`func (o *StatusHistoryModel) SetNewStatusId(v int64)`

SetNewStatusId sets NewStatusId field to given value.

### HasNewStatusId

`func (o *StatusHistoryModel) HasNewStatusId() bool`

HasNewStatusId returns a boolean if a field has been set.

### GetObjectType

`func (o *StatusHistoryModel) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StatusHistoryModel) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StatusHistoryModel) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *StatusHistoryModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *StatusHistoryModel) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *StatusHistoryModel) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *StatusHistoryModel) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *StatusHistoryModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetRecordSnapshot

`func (o *StatusHistoryModel) GetRecordSnapshot() string`

GetRecordSnapshot returns the RecordSnapshot field if non-nil, zero value otherwise.

### GetRecordSnapshotOk

`func (o *StatusHistoryModel) GetRecordSnapshotOk() (*string, bool)`

GetRecordSnapshotOk returns a tuple with the RecordSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordSnapshot

`func (o *StatusHistoryModel) SetRecordSnapshot(v string)`

SetRecordSnapshot sets RecordSnapshot field to given value.

### HasRecordSnapshot

`func (o *StatusHistoryModel) HasRecordSnapshot() bool`

HasRecordSnapshot returns a boolean if a field has been set.

### GetStatusHistoryCreatedBy

`func (o *StatusHistoryModel) GetStatusHistoryCreatedBy() string`

GetStatusHistoryCreatedBy returns the StatusHistoryCreatedBy field if non-nil, zero value otherwise.

### GetStatusHistoryCreatedByOk

`func (o *StatusHistoryModel) GetStatusHistoryCreatedByOk() (*string, bool)`

GetStatusHistoryCreatedByOk returns a tuple with the StatusHistoryCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistoryCreatedBy

`func (o *StatusHistoryModel) SetStatusHistoryCreatedBy(v string)`

SetStatusHistoryCreatedBy sets StatusHistoryCreatedBy field to given value.

### HasStatusHistoryCreatedBy

`func (o *StatusHistoryModel) HasStatusHistoryCreatedBy() bool`

HasStatusHistoryCreatedBy returns a boolean if a field has been set.

### GetStatusHistoryCreatedAt

`func (o *StatusHistoryModel) GetStatusHistoryCreatedAt() string`

GetStatusHistoryCreatedAt returns the StatusHistoryCreatedAt field if non-nil, zero value otherwise.

### GetStatusHistoryCreatedAtOk

`func (o *StatusHistoryModel) GetStatusHistoryCreatedAtOk() (*string, bool)`

GetStatusHistoryCreatedAtOk returns a tuple with the StatusHistoryCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistoryCreatedAt

`func (o *StatusHistoryModel) SetStatusHistoryCreatedAt(v string)`

SetStatusHistoryCreatedAt sets StatusHistoryCreatedAt field to given value.

### HasStatusHistoryCreatedAt

`func (o *StatusHistoryModel) HasStatusHistoryCreatedAt() bool`

HasStatusHistoryCreatedAt returns a boolean if a field has been set.

### GetUserFirstName

`func (o *StatusHistoryModel) GetUserFirstName() string`

GetUserFirstName returns the UserFirstName field if non-nil, zero value otherwise.

### GetUserFirstNameOk

`func (o *StatusHistoryModel) GetUserFirstNameOk() (*string, bool)`

GetUserFirstNameOk returns a tuple with the UserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFirstName

`func (o *StatusHistoryModel) SetUserFirstName(v string)`

SetUserFirstName sets UserFirstName field to given value.

### HasUserFirstName

`func (o *StatusHistoryModel) HasUserFirstName() bool`

HasUserFirstName returns a boolean if a field has been set.

### GetUserLastName

`func (o *StatusHistoryModel) GetUserLastName() string`

GetUserLastName returns the UserLastName field if non-nil, zero value otherwise.

### GetUserLastNameOk

`func (o *StatusHistoryModel) GetUserLastNameOk() (*string, bool)`

GetUserLastNameOk returns a tuple with the UserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLastName

`func (o *StatusHistoryModel) SetUserLastName(v string)`

SetUserLastName sets UserLastName field to given value.

### HasUserLastName

`func (o *StatusHistoryModel) HasUserLastName() bool`

HasUserLastName returns a boolean if a field has been set.

### GetUserNamePrefix

`func (o *StatusHistoryModel) GetUserNamePrefix() string`

GetUserNamePrefix returns the UserNamePrefix field if non-nil, zero value otherwise.

### GetUserNamePrefixOk

`func (o *StatusHistoryModel) GetUserNamePrefixOk() (*string, bool)`

GetUserNamePrefixOk returns a tuple with the UserNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNamePrefix

`func (o *StatusHistoryModel) SetUserNamePrefix(v string)`

SetUserNamePrefix sets UserNamePrefix field to given value.

### HasUserNamePrefix

`func (o *StatusHistoryModel) HasUserNamePrefix() bool`

HasUserNamePrefix returns a boolean if a field has been set.

### GetUserFullName

`func (o *StatusHistoryModel) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *StatusHistoryModel) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *StatusHistoryModel) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.

### HasUserFullName

`func (o *StatusHistoryModel) HasUserFullName() bool`

HasUserFullName returns a boolean if a field has been set.

### GetUserEmail

`func (o *StatusHistoryModel) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *StatusHistoryModel) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *StatusHistoryModel) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *StatusHistoryModel) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserUuid

`func (o *StatusHistoryModel) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *StatusHistoryModel) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *StatusHistoryModel) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *StatusHistoryModel) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

### GetOldStatusName

`func (o *StatusHistoryModel) GetOldStatusName() string`

GetOldStatusName returns the OldStatusName field if non-nil, zero value otherwise.

### GetOldStatusNameOk

`func (o *StatusHistoryModel) GetOldStatusNameOk() (*string, bool)`

GetOldStatusNameOk returns a tuple with the OldStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldStatusName

`func (o *StatusHistoryModel) SetOldStatusName(v string)`

SetOldStatusName sets OldStatusName field to given value.

### HasOldStatusName

`func (o *StatusHistoryModel) HasOldStatusName() bool`

HasOldStatusName returns a boolean if a field has been set.

### GetNewStatusName

`func (o *StatusHistoryModel) GetNewStatusName() string`

GetNewStatusName returns the NewStatusName field if non-nil, zero value otherwise.

### GetNewStatusNameOk

`func (o *StatusHistoryModel) GetNewStatusNameOk() (*string, bool)`

GetNewStatusNameOk returns a tuple with the NewStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatusName

`func (o *StatusHistoryModel) SetNewStatusName(v string)`

SetNewStatusName sets NewStatusName field to given value.

### HasNewStatusName

`func (o *StatusHistoryModel) HasNewStatusName() bool`

HasNewStatusName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


