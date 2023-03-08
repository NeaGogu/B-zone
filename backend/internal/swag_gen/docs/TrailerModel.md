# TrailerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**Name** | Pointer to **string** | name | [optional] 
**RegistrationNr** | Pointer to **string** | registration_nr | [optional] 
**Info** | Pointer to **string** | name | [optional] 
**AppliedCapacities** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Capacities** | Pointer to [**[]CapacityModel**](CapacityModel.md) |  | [optional] 
**Compartments** | Pointer to [**[]CompartmentModel**](CompartmentModel.md) |  | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**CreatedBy** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedBy** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**CreatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**UpdatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**UpdatedByName** | Pointer to **string** | Trailer updated by user full name | [optional] 

## Methods

### NewTrailerModel

`func NewTrailerModel() *TrailerModel`

NewTrailerModel instantiates a new TrailerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerModelWithDefaults

`func NewTrailerModelWithDefaults() *TrailerModel`

NewTrailerModelWithDefaults instantiates a new TrailerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrailerModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrailerModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrailerModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TrailerModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TrailerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrailerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrailerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrailerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationNr

`func (o *TrailerModel) GetRegistrationNr() string`

GetRegistrationNr returns the RegistrationNr field if non-nil, zero value otherwise.

### GetRegistrationNrOk

`func (o *TrailerModel) GetRegistrationNrOk() (*string, bool)`

GetRegistrationNrOk returns a tuple with the RegistrationNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNr

`func (o *TrailerModel) SetRegistrationNr(v string)`

SetRegistrationNr sets RegistrationNr field to given value.

### HasRegistrationNr

`func (o *TrailerModel) HasRegistrationNr() bool`

HasRegistrationNr returns a boolean if a field has been set.

### GetInfo

`func (o *TrailerModel) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TrailerModel) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TrailerModel) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TrailerModel) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *TrailerModel) GetAppliedCapacities() map[string]interface{}`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *TrailerModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *TrailerModel) SetAppliedCapacities(v map[string]interface{})`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *TrailerModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *TrailerModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *TrailerModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *TrailerModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *TrailerModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetCompartments

`func (o *TrailerModel) GetCompartments() []CompartmentModel`

GetCompartments returns the Compartments field if non-nil, zero value otherwise.

### GetCompartmentsOk

`func (o *TrailerModel) GetCompartmentsOk() (*[]CompartmentModel, bool)`

GetCompartmentsOk returns a tuple with the Compartments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartments

`func (o *TrailerModel) SetCompartments(v []CompartmentModel)`

SetCompartments sets Compartments field to given value.

### HasCompartments

`func (o *TrailerModel) HasCompartments() bool`

HasCompartments returns a boolean if a field has been set.

### GetTags

`func (o *TrailerModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TrailerModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TrailerModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TrailerModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagNames

`func (o *TrailerModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *TrailerModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *TrailerModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *TrailerModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetLinks

`func (o *TrailerModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TrailerModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TrailerModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TrailerModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *TrailerModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *TrailerModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *TrailerModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *TrailerModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetFiles

`func (o *TrailerModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TrailerModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TrailerModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TrailerModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TrailerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrailerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrailerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrailerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TrailerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TrailerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TrailerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TrailerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TrailerModel) GetCreatedBy() time.Time`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TrailerModel) GetCreatedByOk() (*time.Time, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TrailerModel) SetCreatedBy(v time.Time)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TrailerModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *TrailerModel) GetUpdatedBy() time.Time`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *TrailerModel) GetUpdatedByOk() (*time.Time, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *TrailerModel) SetUpdatedBy(v time.Time)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *TrailerModel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *TrailerModel) GetCreatedByUser() UsersModel`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *TrailerModel) GetCreatedByUserOk() (*UsersModel, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *TrailerModel) SetCreatedByUser(v UsersModel)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *TrailerModel) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *TrailerModel) GetUpdatedByUser() UsersModel`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *TrailerModel) GetUpdatedByUserOk() (*UsersModel, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *TrailerModel) SetUpdatedByUser(v UsersModel)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *TrailerModel) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *TrailerModel) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *TrailerModel) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *TrailerModel) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *TrailerModel) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


