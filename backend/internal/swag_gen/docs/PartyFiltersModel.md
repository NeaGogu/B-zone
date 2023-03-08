# PartyFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Unique Identifier | [optional] 
**PartyTypeName** | Pointer to **[]string** | Type of this party | [optional] 
**PartyTypeId** | Pointer to **[]int32** | PartyTypeID of this party. 2 &#x3D; contractor, 3 &#x3D; booking | [optional] 
**Name1** | Pointer to **[]string** | Name 1 for party | [optional] 
**Name2** | Pointer to **[]string** | Name 2 for party | [optional] 
**Nr** | Pointer to **[]string** | Number of this party | [optional] 
**ContactPerson** | Pointer to **[]string** | Contact person for party | [optional] 
**Url** | Pointer to **[]string** | Url for party website | [optional] 
**Links** | Pointer to **[]map[string]interface{}** | Activity Link ids | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 

## Methods

### NewPartyFiltersModel

`func NewPartyFiltersModel() *PartyFiltersModel`

NewPartyFiltersModel instantiates a new PartyFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyFiltersModelWithDefaults

`func NewPartyFiltersModelWithDefaults() *PartyFiltersModel`

NewPartyFiltersModelWithDefaults instantiates a new PartyFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartyFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *PartyFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartyTypeName

`func (o *PartyFiltersModel) GetPartyTypeName() []string`

GetPartyTypeName returns the PartyTypeName field if non-nil, zero value otherwise.

### GetPartyTypeNameOk

`func (o *PartyFiltersModel) GetPartyTypeNameOk() (*[]string, bool)`

GetPartyTypeNameOk returns a tuple with the PartyTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyTypeName

`func (o *PartyFiltersModel) SetPartyTypeName(v []string)`

SetPartyTypeName sets PartyTypeName field to given value.

### HasPartyTypeName

`func (o *PartyFiltersModel) HasPartyTypeName() bool`

HasPartyTypeName returns a boolean if a field has been set.

### GetPartyTypeId

`func (o *PartyFiltersModel) GetPartyTypeId() []int32`

GetPartyTypeId returns the PartyTypeId field if non-nil, zero value otherwise.

### GetPartyTypeIdOk

`func (o *PartyFiltersModel) GetPartyTypeIdOk() (*[]int32, bool)`

GetPartyTypeIdOk returns a tuple with the PartyTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyTypeId

`func (o *PartyFiltersModel) SetPartyTypeId(v []int32)`

SetPartyTypeId sets PartyTypeId field to given value.

### HasPartyTypeId

`func (o *PartyFiltersModel) HasPartyTypeId() bool`

HasPartyTypeId returns a boolean if a field has been set.

### GetName1

`func (o *PartyFiltersModel) GetName1() []string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *PartyFiltersModel) GetName1Ok() (*[]string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *PartyFiltersModel) SetName1(v []string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *PartyFiltersModel) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *PartyFiltersModel) GetName2() []string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *PartyFiltersModel) GetName2Ok() (*[]string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *PartyFiltersModel) SetName2(v []string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *PartyFiltersModel) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetNr

`func (o *PartyFiltersModel) GetNr() []string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *PartyFiltersModel) GetNrOk() (*[]string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *PartyFiltersModel) SetNr(v []string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *PartyFiltersModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetContactPerson

`func (o *PartyFiltersModel) GetContactPerson() []string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *PartyFiltersModel) GetContactPersonOk() (*[]string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *PartyFiltersModel) SetContactPerson(v []string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *PartyFiltersModel) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetUrl

`func (o *PartyFiltersModel) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PartyFiltersModel) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PartyFiltersModel) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PartyFiltersModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLinks

`func (o *PartyFiltersModel) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PartyFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PartyFiltersModel) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PartyFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *PartyFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *PartyFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *PartyFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *PartyFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *PartyFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *PartyFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *PartyFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *PartyFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


