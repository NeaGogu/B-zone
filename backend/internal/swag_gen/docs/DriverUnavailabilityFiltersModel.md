# DriverUnavailabilityFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | DriverUnavailability id&#39;s | [optional] 
**UserId** | Pointer to **[]int32** | DriverUnavailability user_id&#39;s | [optional] 
**Reference** | Pointer to **[]string** | DriverUnavailability references | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**CreatedAtSince** | Pointer to **time.Time** | Show create since | [optional] 
**CreatedAtTill** | Pointer to **time.Time** | Show created till | [optional] 

## Methods

### NewDriverUnavailabilityFiltersModel

`func NewDriverUnavailabilityFiltersModel() *DriverUnavailabilityFiltersModel`

NewDriverUnavailabilityFiltersModel instantiates a new DriverUnavailabilityFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverUnavailabilityFiltersModelWithDefaults

`func NewDriverUnavailabilityFiltersModelWithDefaults() *DriverUnavailabilityFiltersModel`

NewDriverUnavailabilityFiltersModelWithDefaults instantiates a new DriverUnavailabilityFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriverUnavailabilityFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverUnavailabilityFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverUnavailabilityFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *DriverUnavailabilityFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *DriverUnavailabilityFiltersModel) GetUserId() []int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DriverUnavailabilityFiltersModel) GetUserIdOk() (*[]int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DriverUnavailabilityFiltersModel) SetUserId(v []int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DriverUnavailabilityFiltersModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetReference

`func (o *DriverUnavailabilityFiltersModel) GetReference() []string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *DriverUnavailabilityFiltersModel) GetReferenceOk() (*[]string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *DriverUnavailabilityFiltersModel) SetReference(v []string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *DriverUnavailabilityFiltersModel) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *DriverUnavailabilityFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *DriverUnavailabilityFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *DriverUnavailabilityFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *DriverUnavailabilityFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *DriverUnavailabilityFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *DriverUnavailabilityFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *DriverUnavailabilityFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *DriverUnavailabilityFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetCreatedAtSince

`func (o *DriverUnavailabilityFiltersModel) GetCreatedAtSince() time.Time`

GetCreatedAtSince returns the CreatedAtSince field if non-nil, zero value otherwise.

### GetCreatedAtSinceOk

`func (o *DriverUnavailabilityFiltersModel) GetCreatedAtSinceOk() (*time.Time, bool)`

GetCreatedAtSinceOk returns a tuple with the CreatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtSince

`func (o *DriverUnavailabilityFiltersModel) SetCreatedAtSince(v time.Time)`

SetCreatedAtSince sets CreatedAtSince field to given value.

### HasCreatedAtSince

`func (o *DriverUnavailabilityFiltersModel) HasCreatedAtSince() bool`

HasCreatedAtSince returns a boolean if a field has been set.

### GetCreatedAtTill

`func (o *DriverUnavailabilityFiltersModel) GetCreatedAtTill() time.Time`

GetCreatedAtTill returns the CreatedAtTill field if non-nil, zero value otherwise.

### GetCreatedAtTillOk

`func (o *DriverUnavailabilityFiltersModel) GetCreatedAtTillOk() (*time.Time, bool)`

GetCreatedAtTillOk returns a tuple with the CreatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTill

`func (o *DriverUnavailabilityFiltersModel) SetCreatedAtTill(v time.Time)`

SetCreatedAtTill sets CreatedAtTill field to given value.

### HasCreatedAtTill

`func (o *DriverUnavailabilityFiltersModel) HasCreatedAtTill() bool`

HasCreatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


