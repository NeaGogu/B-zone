# DriverUnavailabilityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**UserId** | Pointer to **string** | The user_id of the DriverUnavailability | [optional] 
**Reference** | Pointer to **string** | The reference of the DriverUnavailability | [optional] 
**DateTimeTo** | Pointer to **time.Time** | date_time_to date time | [optional] 
**DateTimeFrom** | Pointer to **time.Time** | date_time_from date time | [optional] 
**UserLink** | Pointer to [**LinkModel**](LinkModel.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: Driver Unavailability has been removed and is no longer visible in any bumbal interface | [optional] 

## Methods

### NewDriverUnavailabilityModel

`func NewDriverUnavailabilityModel() *DriverUnavailabilityModel`

NewDriverUnavailabilityModel instantiates a new DriverUnavailabilityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverUnavailabilityModelWithDefaults

`func NewDriverUnavailabilityModelWithDefaults() *DriverUnavailabilityModel`

NewDriverUnavailabilityModelWithDefaults instantiates a new DriverUnavailabilityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriverUnavailabilityModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverUnavailabilityModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverUnavailabilityModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DriverUnavailabilityModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *DriverUnavailabilityModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DriverUnavailabilityModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DriverUnavailabilityModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DriverUnavailabilityModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetReference

`func (o *DriverUnavailabilityModel) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *DriverUnavailabilityModel) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *DriverUnavailabilityModel) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *DriverUnavailabilityModel) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *DriverUnavailabilityModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *DriverUnavailabilityModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *DriverUnavailabilityModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *DriverUnavailabilityModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *DriverUnavailabilityModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *DriverUnavailabilityModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *DriverUnavailabilityModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *DriverUnavailabilityModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetUserLink

`func (o *DriverUnavailabilityModel) GetUserLink() LinkModel`

GetUserLink returns the UserLink field if non-nil, zero value otherwise.

### GetUserLinkOk

`func (o *DriverUnavailabilityModel) GetUserLinkOk() (*LinkModel, bool)`

GetUserLinkOk returns a tuple with the UserLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLink

`func (o *DriverUnavailabilityModel) SetUserLink(v LinkModel)`

SetUserLink sets UserLink field to given value.

### HasUserLink

`func (o *DriverUnavailabilityModel) HasUserLink() bool`

HasUserLink returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DriverUnavailabilityModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DriverUnavailabilityModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DriverUnavailabilityModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DriverUnavailabilityModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DriverUnavailabilityModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DriverUnavailabilityModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DriverUnavailabilityModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DriverUnavailabilityModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *DriverUnavailabilityModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DriverUnavailabilityModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DriverUnavailabilityModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DriverUnavailabilityModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActive

`func (o *DriverUnavailabilityModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DriverUnavailabilityModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DriverUnavailabilityModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DriverUnavailabilityModel) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


