# UsersFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **[]string** | unique per user | [optional] [readonly] 
**PauseId** | Pointer to **[]int32** | ids of pause schemes applied to user with teh role driver | [optional] 
**RoleId** | Pointer to **[]int32** | ids of the user roles, 1: Guest, 2: Driver, 3: Planner, 4: Manager, 5: Admin | [optional] 
**PartyId** | Pointer to **[]int32** | Associated Party IDs | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**System** | Pointer to **[]bool** | System users | [optional] 
**Activated** | Pointer to **[]bool** | Activated users | [optional] 

## Methods

### NewUsersFiltersModel

`func NewUsersFiltersModel() *UsersFiltersModel`

NewUsersFiltersModel instantiates a new UsersFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersFiltersModelWithDefaults

`func NewUsersFiltersModelWithDefaults() *UsersFiltersModel`

NewUsersFiltersModelWithDefaults instantiates a new UsersFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *UsersFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UsersFiltersModel) GetUuid() []string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UsersFiltersModel) GetUuidOk() (*[]string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UsersFiltersModel) SetUuid(v []string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UsersFiltersModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPauseId

`func (o *UsersFiltersModel) GetPauseId() []int32`

GetPauseId returns the PauseId field if non-nil, zero value otherwise.

### GetPauseIdOk

`func (o *UsersFiltersModel) GetPauseIdOk() (*[]int32, bool)`

GetPauseIdOk returns a tuple with the PauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseId

`func (o *UsersFiltersModel) SetPauseId(v []int32)`

SetPauseId sets PauseId field to given value.

### HasPauseId

`func (o *UsersFiltersModel) HasPauseId() bool`

HasPauseId returns a boolean if a field has been set.

### GetRoleId

`func (o *UsersFiltersModel) GetRoleId() []int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UsersFiltersModel) GetRoleIdOk() (*[]int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UsersFiltersModel) SetRoleId(v []int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UsersFiltersModel) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetPartyId

`func (o *UsersFiltersModel) GetPartyId() []int32`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *UsersFiltersModel) GetPartyIdOk() (*[]int32, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *UsersFiltersModel) SetPartyId(v []int32)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *UsersFiltersModel) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetTagNames

`func (o *UsersFiltersModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *UsersFiltersModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *UsersFiltersModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *UsersFiltersModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetZoneNames

`func (o *UsersFiltersModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *UsersFiltersModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *UsersFiltersModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *UsersFiltersModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetSystem

`func (o *UsersFiltersModel) GetSystem() []bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *UsersFiltersModel) GetSystemOk() (*[]bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *UsersFiltersModel) SetSystem(v []bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *UsersFiltersModel) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetActivated

`func (o *UsersFiltersModel) GetActivated() []bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *UsersFiltersModel) GetActivatedOk() (*[]bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *UsersFiltersModel) SetActivated(v []bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *UsersFiltersModel) HasActivated() bool`

HasActivated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


