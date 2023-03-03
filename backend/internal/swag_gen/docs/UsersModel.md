# UsersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** | unique per user | [optional] [readonly] 
**RoleId** | Pointer to **int64** | id of the user role, 1: Guest, 2: Driver, 3: Planner, 4: Manager, 5: Admin | [optional] 
**RoleName** | Pointer to **string** | Role name | [optional] 
**PartyId** | Pointer to **int64** | Associated Party ID | [optional] [readonly] 
**PartyName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** | First name | [optional] 
**NamePrefix** | Pointer to **string** | Name prefix | [optional] 
**LastName** | Pointer to **string** | Last name | [optional] 
**FullName** | Pointer to **string** | Full name | [optional] [readonly] 
**Initials** | Pointer to **string** | user password (set only, no read) | [optional] 
**Email** | Pointer to **string** | user email (used for login) | [optional] 
**Password** | Pointer to **string** | user password (set only, no read) | [optional] 
**LangCode** | Pointer to **string** | lang code (nl &#x3D; default) | [optional] 
**AddressId** | Pointer to **int64** | id of the user address | [optional] 
**Addresses** | Pointer to [**[]AddressModel**](AddressModel.md) | user address (mostly interesting for drivers) | [optional] 
**PauseId** | Pointer to **int64** | id of the pause to be applied by default for new user routes | [optional] 
**SpeedFactor** | Pointer to **float64** | Speed Factor | [optional] 
**DurationFactor** | Pointer to **float64** | Duration Factor | [optional] 
**Activated** | Pointer to **bool** | Whether user is activated or not | [optional] 
**Removed** | Pointer to **bool** | Whether user is removed or not | [optional] 
**Active** | Pointer to **bool** | Whether user is still active or not | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**Driver** | Pointer to **bool** | The user can also be used as driver | [optional] 
**Zones** | Pointer to [**[]ZoneModel**](ZoneModel.md) |  | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**DriverUnavailabilities** | Pointer to [**[]DriverUnavailabilityModel**](DriverUnavailabilityModel.md) |  | [optional] 

## Methods

### NewUsersModel

`func NewUsersModel() *UsersModel`

NewUsersModel instantiates a new UsersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersModelWithDefaults

`func NewUsersModelWithDefaults() *UsersModel`

NewUsersModelWithDefaults instantiates a new UsersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UsersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UsersModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UsersModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UsersModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UsersModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetRoleId

`func (o *UsersModel) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UsersModel) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UsersModel) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UsersModel) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *UsersModel) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *UsersModel) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *UsersModel) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *UsersModel) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetPartyId

`func (o *UsersModel) GetPartyId() int64`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *UsersModel) GetPartyIdOk() (*int64, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *UsersModel) SetPartyId(v int64)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *UsersModel) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetPartyName

`func (o *UsersModel) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *UsersModel) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *UsersModel) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *UsersModel) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.

### GetFirstName

`func (o *UsersModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsersModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsersModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UsersModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetNamePrefix

`func (o *UsersModel) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *UsersModel) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *UsersModel) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *UsersModel) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### GetLastName

`func (o *UsersModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsersModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsersModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UsersModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFullName

`func (o *UsersModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UsersModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UsersModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UsersModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetInitials

`func (o *UsersModel) GetInitials() string`

GetInitials returns the Initials field if non-nil, zero value otherwise.

### GetInitialsOk

`func (o *UsersModel) GetInitialsOk() (*string, bool)`

GetInitialsOk returns a tuple with the Initials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitials

`func (o *UsersModel) SetInitials(v string)`

SetInitials sets Initials field to given value.

### HasInitials

`func (o *UsersModel) HasInitials() bool`

HasInitials returns a boolean if a field has been set.

### GetEmail

`func (o *UsersModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UsersModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *UsersModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UsersModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UsersModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UsersModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetLangCode

`func (o *UsersModel) GetLangCode() string`

GetLangCode returns the LangCode field if non-nil, zero value otherwise.

### GetLangCodeOk

`func (o *UsersModel) GetLangCodeOk() (*string, bool)`

GetLangCodeOk returns a tuple with the LangCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangCode

`func (o *UsersModel) SetLangCode(v string)`

SetLangCode sets LangCode field to given value.

### HasLangCode

`func (o *UsersModel) HasLangCode() bool`

HasLangCode returns a boolean if a field has been set.

### GetAddressId

`func (o *UsersModel) GetAddressId() int64`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *UsersModel) GetAddressIdOk() (*int64, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *UsersModel) SetAddressId(v int64)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *UsersModel) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetAddresses

`func (o *UsersModel) GetAddresses() []AddressModel`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *UsersModel) GetAddressesOk() (*[]AddressModel, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *UsersModel) SetAddresses(v []AddressModel)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *UsersModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetPauseId

`func (o *UsersModel) GetPauseId() int64`

GetPauseId returns the PauseId field if non-nil, zero value otherwise.

### GetPauseIdOk

`func (o *UsersModel) GetPauseIdOk() (*int64, bool)`

GetPauseIdOk returns a tuple with the PauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseId

`func (o *UsersModel) SetPauseId(v int64)`

SetPauseId sets PauseId field to given value.

### HasPauseId

`func (o *UsersModel) HasPauseId() bool`

HasPauseId returns a boolean if a field has been set.

### GetSpeedFactor

`func (o *UsersModel) GetSpeedFactor() float64`

GetSpeedFactor returns the SpeedFactor field if non-nil, zero value otherwise.

### GetSpeedFactorOk

`func (o *UsersModel) GetSpeedFactorOk() (*float64, bool)`

GetSpeedFactorOk returns a tuple with the SpeedFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedFactor

`func (o *UsersModel) SetSpeedFactor(v float64)`

SetSpeedFactor sets SpeedFactor field to given value.

### HasSpeedFactor

`func (o *UsersModel) HasSpeedFactor() bool`

HasSpeedFactor returns a boolean if a field has been set.

### GetDurationFactor

`func (o *UsersModel) GetDurationFactor() float64`

GetDurationFactor returns the DurationFactor field if non-nil, zero value otherwise.

### GetDurationFactorOk

`func (o *UsersModel) GetDurationFactorOk() (*float64, bool)`

GetDurationFactorOk returns a tuple with the DurationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationFactor

`func (o *UsersModel) SetDurationFactor(v float64)`

SetDurationFactor sets DurationFactor field to given value.

### HasDurationFactor

`func (o *UsersModel) HasDurationFactor() bool`

HasDurationFactor returns a boolean if a field has been set.

### GetActivated

`func (o *UsersModel) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *UsersModel) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *UsersModel) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *UsersModel) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetRemoved

`func (o *UsersModel) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *UsersModel) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *UsersModel) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *UsersModel) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetActive

`func (o *UsersModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UsersModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UsersModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UsersModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTags

`func (o *UsersModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UsersModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UsersModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UsersModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagNames

`func (o *UsersModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *UsersModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *UsersModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *UsersModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetDriver

`func (o *UsersModel) GetDriver() bool`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *UsersModel) GetDriverOk() (*bool, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *UsersModel) SetDriver(v bool)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *UsersModel) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetZones

`func (o *UsersModel) GetZones() []ZoneModel`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *UsersModel) GetZonesOk() (*[]ZoneModel, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *UsersModel) SetZones(v []ZoneModel)`

SetZones sets Zones field to given value.

### HasZones

`func (o *UsersModel) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZoneNames

`func (o *UsersModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *UsersModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *UsersModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *UsersModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetLinks

`func (o *UsersModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UsersModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UsersModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UsersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *UsersModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UsersModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UsersModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UsersModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetDriverUnavailabilities

`func (o *UsersModel) GetDriverUnavailabilities() []DriverUnavailabilityModel`

GetDriverUnavailabilities returns the DriverUnavailabilities field if non-nil, zero value otherwise.

### GetDriverUnavailabilitiesOk

`func (o *UsersModel) GetDriverUnavailabilitiesOk() (*[]DriverUnavailabilityModel, bool)`

GetDriverUnavailabilitiesOk returns a tuple with the DriverUnavailabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverUnavailabilities

`func (o *UsersModel) SetDriverUnavailabilities(v []DriverUnavailabilityModel)`

SetDriverUnavailabilities sets DriverUnavailabilities field to given value.

### HasDriverUnavailabilities

`func (o *UsersModel) HasDriverUnavailabilities() bool`

HasDriverUnavailabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


