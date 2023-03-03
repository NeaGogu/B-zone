# AddressModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AddressId** | Pointer to **int64** |  | [optional] 
**PartyId** | Pointer to **int32** |  | [optional] 
**PartyName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] [readonly] 
**Name1** | Pointer to **string** |  | [optional] 
**Name2** | Pointer to **string** |  | [optional] 
**Street1** | Pointer to **string** |  | [optional] 
**Street2** | Pointer to **string** |  | [optional] 
**FullAddressline** | Pointer to **string** |  | [optional] [readonly] 
**HouseNr** | Pointer to **string** |  | [optional] 
**HouseNrAddendum** | Pointer to **string** |  | [optional] 
**Zipcode** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**IsoCountry** | Pointer to **string** |  | [optional] 
**CountryName** | Pointer to **string** |  | [optional] 
**TimeFrom** | Pointer to **string** |  | [optional] 
**TimeTo** | Pointer to **string** |  | [optional] 
**OpeningHours** | Pointer to [**[]OpeningHoursRuleModel**](OpeningHoursRuleModel.md) |  | [optional] 
**Duration** | Pointer to **int32** | Default duration for activities on this address in minutes | [optional] 
**AddressTypeNames** | Pointer to **[]string** | Address Type names | [optional] 
**Emails** | Pointer to [**[]EmailModel**](EmailModel.md) |  | [optional] 
**PhoneNrs** | Pointer to [**[]PhoneNrModel**](PhoneNrModel.md) |  | [optional] 
**Latitude** | Pointer to **string** |  | [optional] 
**Longitude** | Pointer to **string** |  | [optional] 
**ContactPerson** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Notes** | Pointer to [**[]NoteModel**](NoteModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 

## Methods

### NewAddressModel

`func NewAddressModel() *AddressModel`

NewAddressModel instantiates a new AddressModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressModelWithDefaults

`func NewAddressModelWithDefaults() *AddressModel`

NewAddressModelWithDefaults instantiates a new AddressModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddressModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddressId

`func (o *AddressModel) GetAddressId() int64`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *AddressModel) GetAddressIdOk() (*int64, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *AddressModel) SetAddressId(v int64)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *AddressModel) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetPartyId

`func (o *AddressModel) GetPartyId() int32`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *AddressModel) GetPartyIdOk() (*int32, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *AddressModel) SetPartyId(v int32)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *AddressModel) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetPartyName

`func (o *AddressModel) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *AddressModel) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *AddressModel) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *AddressModel) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.

### GetCode

`func (o *AddressModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddressModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddressModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddressModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSummary

`func (o *AddressModel) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AddressModel) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AddressModel) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AddressModel) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetFullName

`func (o *AddressModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AddressModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AddressModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AddressModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetName1

`func (o *AddressModel) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *AddressModel) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *AddressModel) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *AddressModel) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *AddressModel) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *AddressModel) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *AddressModel) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *AddressModel) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetStreet1

`func (o *AddressModel) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *AddressModel) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *AddressModel) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.

### HasStreet1

`func (o *AddressModel) HasStreet1() bool`

HasStreet1 returns a boolean if a field has been set.

### GetStreet2

`func (o *AddressModel) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *AddressModel) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *AddressModel) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *AddressModel) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetFullAddressline

`func (o *AddressModel) GetFullAddressline() string`

GetFullAddressline returns the FullAddressline field if non-nil, zero value otherwise.

### GetFullAddresslineOk

`func (o *AddressModel) GetFullAddresslineOk() (*string, bool)`

GetFullAddresslineOk returns a tuple with the FullAddressline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddressline

`func (o *AddressModel) SetFullAddressline(v string)`

SetFullAddressline sets FullAddressline field to given value.

### HasFullAddressline

`func (o *AddressModel) HasFullAddressline() bool`

HasFullAddressline returns a boolean if a field has been set.

### GetHouseNr

`func (o *AddressModel) GetHouseNr() string`

GetHouseNr returns the HouseNr field if non-nil, zero value otherwise.

### GetHouseNrOk

`func (o *AddressModel) GetHouseNrOk() (*string, bool)`

GetHouseNrOk returns a tuple with the HouseNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNr

`func (o *AddressModel) SetHouseNr(v string)`

SetHouseNr sets HouseNr field to given value.

### HasHouseNr

`func (o *AddressModel) HasHouseNr() bool`

HasHouseNr returns a boolean if a field has been set.

### GetHouseNrAddendum

`func (o *AddressModel) GetHouseNrAddendum() string`

GetHouseNrAddendum returns the HouseNrAddendum field if non-nil, zero value otherwise.

### GetHouseNrAddendumOk

`func (o *AddressModel) GetHouseNrAddendumOk() (*string, bool)`

GetHouseNrAddendumOk returns a tuple with the HouseNrAddendum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNrAddendum

`func (o *AddressModel) SetHouseNrAddendum(v string)`

SetHouseNrAddendum sets HouseNrAddendum field to given value.

### HasHouseNrAddendum

`func (o *AddressModel) HasHouseNrAddendum() bool`

HasHouseNrAddendum returns a boolean if a field has been set.

### GetZipcode

`func (o *AddressModel) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *AddressModel) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *AddressModel) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *AddressModel) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### GetCity

`func (o *AddressModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *AddressModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIsoCountry

`func (o *AddressModel) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *AddressModel) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *AddressModel) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *AddressModel) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### GetCountryName

`func (o *AddressModel) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *AddressModel) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *AddressModel) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *AddressModel) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetTimeFrom

`func (o *AddressModel) GetTimeFrom() string`

GetTimeFrom returns the TimeFrom field if non-nil, zero value otherwise.

### GetTimeFromOk

`func (o *AddressModel) GetTimeFromOk() (*string, bool)`

GetTimeFromOk returns a tuple with the TimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrom

`func (o *AddressModel) SetTimeFrom(v string)`

SetTimeFrom sets TimeFrom field to given value.

### HasTimeFrom

`func (o *AddressModel) HasTimeFrom() bool`

HasTimeFrom returns a boolean if a field has been set.

### GetTimeTo

`func (o *AddressModel) GetTimeTo() string`

GetTimeTo returns the TimeTo field if non-nil, zero value otherwise.

### GetTimeToOk

`func (o *AddressModel) GetTimeToOk() (*string, bool)`

GetTimeToOk returns a tuple with the TimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTo

`func (o *AddressModel) SetTimeTo(v string)`

SetTimeTo sets TimeTo field to given value.

### HasTimeTo

`func (o *AddressModel) HasTimeTo() bool`

HasTimeTo returns a boolean if a field has been set.

### GetOpeningHours

`func (o *AddressModel) GetOpeningHours() []OpeningHoursRuleModel`

GetOpeningHours returns the OpeningHours field if non-nil, zero value otherwise.

### GetOpeningHoursOk

`func (o *AddressModel) GetOpeningHoursOk() (*[]OpeningHoursRuleModel, bool)`

GetOpeningHoursOk returns a tuple with the OpeningHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningHours

`func (o *AddressModel) SetOpeningHours(v []OpeningHoursRuleModel)`

SetOpeningHours sets OpeningHours field to given value.

### HasOpeningHours

`func (o *AddressModel) HasOpeningHours() bool`

HasOpeningHours returns a boolean if a field has been set.

### GetDuration

`func (o *AddressModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AddressModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AddressModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AddressModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAddressTypeNames

`func (o *AddressModel) GetAddressTypeNames() []string`

GetAddressTypeNames returns the AddressTypeNames field if non-nil, zero value otherwise.

### GetAddressTypeNamesOk

`func (o *AddressModel) GetAddressTypeNamesOk() (*[]string, bool)`

GetAddressTypeNamesOk returns a tuple with the AddressTypeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTypeNames

`func (o *AddressModel) SetAddressTypeNames(v []string)`

SetAddressTypeNames sets AddressTypeNames field to given value.

### HasAddressTypeNames

`func (o *AddressModel) HasAddressTypeNames() bool`

HasAddressTypeNames returns a boolean if a field has been set.

### GetEmails

`func (o *AddressModel) GetEmails() []EmailModel`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AddressModel) GetEmailsOk() (*[]EmailModel, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AddressModel) SetEmails(v []EmailModel)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AddressModel) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhoneNrs

`func (o *AddressModel) GetPhoneNrs() []PhoneNrModel`

GetPhoneNrs returns the PhoneNrs field if non-nil, zero value otherwise.

### GetPhoneNrsOk

`func (o *AddressModel) GetPhoneNrsOk() (*[]PhoneNrModel, bool)`

GetPhoneNrsOk returns a tuple with the PhoneNrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNrs

`func (o *AddressModel) SetPhoneNrs(v []PhoneNrModel)`

SetPhoneNrs sets PhoneNrs field to given value.

### HasPhoneNrs

`func (o *AddressModel) HasPhoneNrs() bool`

HasPhoneNrs returns a boolean if a field has been set.

### GetLatitude

`func (o *AddressModel) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *AddressModel) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *AddressModel) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *AddressModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *AddressModel) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *AddressModel) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *AddressModel) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *AddressModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetContactPerson

`func (o *AddressModel) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *AddressModel) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *AddressModel) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *AddressModel) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetUserId

`func (o *AddressModel) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AddressModel) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AddressModel) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AddressModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLinks

`func (o *AddressModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AddressModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AddressModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AddressModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *AddressModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AddressModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AddressModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AddressModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetNotes

`func (o *AddressModel) GetNotes() []NoteModel`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddressModel) GetNotesOk() (*[]NoteModel, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddressModel) SetNotes(v []NoteModel)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddressModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFiles

`func (o *AddressModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AddressModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AddressModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AddressModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetTagNames

`func (o *AddressModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *AddressModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *AddressModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *AddressModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetTags

`func (o *AddressModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddressModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddressModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddressModel) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


