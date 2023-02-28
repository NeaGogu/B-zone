# AddressAppliedModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**PartyId** | Pointer to **int32** |  | [optional] 
**AddressId** | Pointer to **int32** |  | [optional] 
**AddressLinks** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] [readonly] 
**Name1** | Pointer to **string** |  | [optional] 
**Name2** | Pointer to **string** |  | [optional] 
**Street1** | Pointer to **string** |  | [optional] 
**Street2** | Pointer to **string** |  | [optional] 
**HouseNr** | Pointer to **string** |  | [optional] 
**HouseNrAddendum** | Pointer to **string** |  | [optional] 
**Zipcode** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**IsoCountry** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]EmailModel**](EmailModel.md) |  | [optional] 
**PhoneNrs** | Pointer to [**[]PhoneNrModel**](PhoneNrModel.md) |  | [optional] 
**ContactPerson** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **string** |  | [optional] 
**Longitude** | Pointer to **string** |  | [optional] 

## Methods

### NewAddressAppliedModel

`func NewAddressAppliedModel() *AddressAppliedModel`

NewAddressAppliedModel instantiates a new AddressAppliedModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressAppliedModelWithDefaults

`func NewAddressAppliedModelWithDefaults() *AddressAppliedModel`

NewAddressAppliedModelWithDefaults instantiates a new AddressAppliedModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressAppliedModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressAppliedModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressAppliedModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddressAppliedModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartyId

`func (o *AddressAppliedModel) GetPartyId() int32`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *AddressAppliedModel) GetPartyIdOk() (*int32, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *AddressAppliedModel) SetPartyId(v int32)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *AddressAppliedModel) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetAddressId

`func (o *AddressAppliedModel) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *AddressAppliedModel) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *AddressAppliedModel) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *AddressAppliedModel) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetAddressLinks

`func (o *AddressAppliedModel) GetAddressLinks() []LinkModel`

GetAddressLinks returns the AddressLinks field if non-nil, zero value otherwise.

### GetAddressLinksOk

`func (o *AddressAppliedModel) GetAddressLinksOk() (*[]LinkModel, bool)`

GetAddressLinksOk returns a tuple with the AddressLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLinks

`func (o *AddressAppliedModel) SetAddressLinks(v []LinkModel)`

SetAddressLinks sets AddressLinks field to given value.

### HasAddressLinks

`func (o *AddressAppliedModel) HasAddressLinks() bool`

HasAddressLinks returns a boolean if a field has been set.

### GetSummary

`func (o *AddressAppliedModel) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AddressAppliedModel) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AddressAppliedModel) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AddressAppliedModel) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetFullName

`func (o *AddressAppliedModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AddressAppliedModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AddressAppliedModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AddressAppliedModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetName1

`func (o *AddressAppliedModel) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *AddressAppliedModel) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *AddressAppliedModel) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *AddressAppliedModel) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *AddressAppliedModel) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *AddressAppliedModel) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *AddressAppliedModel) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *AddressAppliedModel) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetStreet1

`func (o *AddressAppliedModel) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *AddressAppliedModel) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *AddressAppliedModel) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.

### HasStreet1

`func (o *AddressAppliedModel) HasStreet1() bool`

HasStreet1 returns a boolean if a field has been set.

### GetStreet2

`func (o *AddressAppliedModel) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *AddressAppliedModel) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *AddressAppliedModel) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *AddressAppliedModel) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetHouseNr

`func (o *AddressAppliedModel) GetHouseNr() string`

GetHouseNr returns the HouseNr field if non-nil, zero value otherwise.

### GetHouseNrOk

`func (o *AddressAppliedModel) GetHouseNrOk() (*string, bool)`

GetHouseNrOk returns a tuple with the HouseNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNr

`func (o *AddressAppliedModel) SetHouseNr(v string)`

SetHouseNr sets HouseNr field to given value.

### HasHouseNr

`func (o *AddressAppliedModel) HasHouseNr() bool`

HasHouseNr returns a boolean if a field has been set.

### GetHouseNrAddendum

`func (o *AddressAppliedModel) GetHouseNrAddendum() string`

GetHouseNrAddendum returns the HouseNrAddendum field if non-nil, zero value otherwise.

### GetHouseNrAddendumOk

`func (o *AddressAppliedModel) GetHouseNrAddendumOk() (*string, bool)`

GetHouseNrAddendumOk returns a tuple with the HouseNrAddendum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNrAddendum

`func (o *AddressAppliedModel) SetHouseNrAddendum(v string)`

SetHouseNrAddendum sets HouseNrAddendum field to given value.

### HasHouseNrAddendum

`func (o *AddressAppliedModel) HasHouseNrAddendum() bool`

HasHouseNrAddendum returns a boolean if a field has been set.

### GetZipcode

`func (o *AddressAppliedModel) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *AddressAppliedModel) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *AddressAppliedModel) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *AddressAppliedModel) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### GetCity

`func (o *AddressAppliedModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressAppliedModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressAppliedModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressAppliedModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *AddressAppliedModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressAppliedModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressAppliedModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressAppliedModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIsoCountry

`func (o *AddressAppliedModel) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *AddressAppliedModel) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *AddressAppliedModel) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *AddressAppliedModel) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### GetEmails

`func (o *AddressAppliedModel) GetEmails() []EmailModel`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AddressAppliedModel) GetEmailsOk() (*[]EmailModel, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AddressAppliedModel) SetEmails(v []EmailModel)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AddressAppliedModel) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhoneNrs

`func (o *AddressAppliedModel) GetPhoneNrs() []PhoneNrModel`

GetPhoneNrs returns the PhoneNrs field if non-nil, zero value otherwise.

### GetPhoneNrsOk

`func (o *AddressAppliedModel) GetPhoneNrsOk() (*[]PhoneNrModel, bool)`

GetPhoneNrsOk returns a tuple with the PhoneNrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNrs

`func (o *AddressAppliedModel) SetPhoneNrs(v []PhoneNrModel)`

SetPhoneNrs sets PhoneNrs field to given value.

### HasPhoneNrs

`func (o *AddressAppliedModel) HasPhoneNrs() bool`

HasPhoneNrs returns a boolean if a field has been set.

### GetContactPerson

`func (o *AddressAppliedModel) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *AddressAppliedModel) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *AddressAppliedModel) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *AddressAppliedModel) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetLatitude

`func (o *AddressAppliedModel) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *AddressAppliedModel) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *AddressAppliedModel) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *AddressAppliedModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *AddressAppliedModel) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *AddressAppliedModel) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *AddressAppliedModel) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *AddressAppliedModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


