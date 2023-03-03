# CompartmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**Nr** | Pointer to **string** | Nr of compartment | [optional] 
**TrailerId** | Pointer to **int32** |  | [optional] 
**FilledCapacities** | Pointer to [**[]FilledCapacityModel**](FilledCapacityModel.md) |  | [optional] [readonly] 
**TrailerLink** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**CompartmentCreatedBy** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**CompartmentUpdatedBy** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**CompartmentCreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**CompartmentUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 

## Methods

### NewCompartmentModel

`func NewCompartmentModel() *CompartmentModel`

NewCompartmentModel instantiates a new CompartmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartmentModelWithDefaults

`func NewCompartmentModelWithDefaults() *CompartmentModel`

NewCompartmentModelWithDefaults instantiates a new CompartmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompartmentModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompartmentModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompartmentModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CompartmentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNr

`func (o *CompartmentModel) GetNr() string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *CompartmentModel) GetNrOk() (*string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *CompartmentModel) SetNr(v string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *CompartmentModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetTrailerId

`func (o *CompartmentModel) GetTrailerId() int32`

GetTrailerId returns the TrailerId field if non-nil, zero value otherwise.

### GetTrailerIdOk

`func (o *CompartmentModel) GetTrailerIdOk() (*int32, bool)`

GetTrailerIdOk returns a tuple with the TrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerId

`func (o *CompartmentModel) SetTrailerId(v int32)`

SetTrailerId sets TrailerId field to given value.

### HasTrailerId

`func (o *CompartmentModel) HasTrailerId() bool`

HasTrailerId returns a boolean if a field has been set.

### GetFilledCapacities

`func (o *CompartmentModel) GetFilledCapacities() []FilledCapacityModel`

GetFilledCapacities returns the FilledCapacities field if non-nil, zero value otherwise.

### GetFilledCapacitiesOk

`func (o *CompartmentModel) GetFilledCapacitiesOk() (*[]FilledCapacityModel, bool)`

GetFilledCapacitiesOk returns a tuple with the FilledCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledCapacities

`func (o *CompartmentModel) SetFilledCapacities(v []FilledCapacityModel)`

SetFilledCapacities sets FilledCapacities field to given value.

### HasFilledCapacities

`func (o *CompartmentModel) HasFilledCapacities() bool`

HasFilledCapacities returns a boolean if a field has been set.

### GetTrailerLink

`func (o *CompartmentModel) GetTrailerLink() []LinkModel`

GetTrailerLink returns the TrailerLink field if non-nil, zero value otherwise.

### GetTrailerLinkOk

`func (o *CompartmentModel) GetTrailerLinkOk() (*[]LinkModel, bool)`

GetTrailerLinkOk returns a tuple with the TrailerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLink

`func (o *CompartmentModel) SetTrailerLink(v []LinkModel)`

SetTrailerLink sets TrailerLink field to given value.

### HasTrailerLink

`func (o *CompartmentModel) HasTrailerLink() bool`

HasTrailerLink returns a boolean if a field has been set.

### GetLinks

`func (o *CompartmentModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CompartmentModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CompartmentModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CompartmentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *CompartmentModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *CompartmentModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *CompartmentModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *CompartmentModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetCompartmentCreatedBy

`func (o *CompartmentModel) GetCompartmentCreatedBy() UsersModel`

GetCompartmentCreatedBy returns the CompartmentCreatedBy field if non-nil, zero value otherwise.

### GetCompartmentCreatedByOk

`func (o *CompartmentModel) GetCompartmentCreatedByOk() (*UsersModel, bool)`

GetCompartmentCreatedByOk returns a tuple with the CompartmentCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentCreatedBy

`func (o *CompartmentModel) SetCompartmentCreatedBy(v UsersModel)`

SetCompartmentCreatedBy sets CompartmentCreatedBy field to given value.

### HasCompartmentCreatedBy

`func (o *CompartmentModel) HasCompartmentCreatedBy() bool`

HasCompartmentCreatedBy returns a boolean if a field has been set.

### GetCompartmentUpdatedBy

`func (o *CompartmentModel) GetCompartmentUpdatedBy() UsersModel`

GetCompartmentUpdatedBy returns the CompartmentUpdatedBy field if non-nil, zero value otherwise.

### GetCompartmentUpdatedByOk

`func (o *CompartmentModel) GetCompartmentUpdatedByOk() (*UsersModel, bool)`

GetCompartmentUpdatedByOk returns a tuple with the CompartmentUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentUpdatedBy

`func (o *CompartmentModel) SetCompartmentUpdatedBy(v UsersModel)`

SetCompartmentUpdatedBy sets CompartmentUpdatedBy field to given value.

### HasCompartmentUpdatedBy

`func (o *CompartmentModel) HasCompartmentUpdatedBy() bool`

HasCompartmentUpdatedBy returns a boolean if a field has been set.

### GetCompartmentCreatedAt

`func (o *CompartmentModel) GetCompartmentCreatedAt() time.Time`

GetCompartmentCreatedAt returns the CompartmentCreatedAt field if non-nil, zero value otherwise.

### GetCompartmentCreatedAtOk

`func (o *CompartmentModel) GetCompartmentCreatedAtOk() (*time.Time, bool)`

GetCompartmentCreatedAtOk returns a tuple with the CompartmentCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentCreatedAt

`func (o *CompartmentModel) SetCompartmentCreatedAt(v time.Time)`

SetCompartmentCreatedAt sets CompartmentCreatedAt field to given value.

### HasCompartmentCreatedAt

`func (o *CompartmentModel) HasCompartmentCreatedAt() bool`

HasCompartmentCreatedAt returns a boolean if a field has been set.

### GetCompartmentUpdatedAt

`func (o *CompartmentModel) GetCompartmentUpdatedAt() time.Time`

GetCompartmentUpdatedAt returns the CompartmentUpdatedAt field if non-nil, zero value otherwise.

### GetCompartmentUpdatedAtOk

`func (o *CompartmentModel) GetCompartmentUpdatedAtOk() (*time.Time, bool)`

GetCompartmentUpdatedAtOk returns a tuple with the CompartmentUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentUpdatedAt

`func (o *CompartmentModel) SetCompartmentUpdatedAt(v time.Time)`

SetCompartmentUpdatedAt sets CompartmentUpdatedAt field to given value.

### HasCompartmentUpdatedAt

`func (o *CompartmentModel) HasCompartmentUpdatedAt() bool`

HasCompartmentUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


