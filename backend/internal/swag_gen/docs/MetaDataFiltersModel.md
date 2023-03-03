# MetaDataFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | MetaData id&#39;s | [optional] 
**Name** | Pointer to **[]string** | MetaData names | [optional] 
**ObjectTypeName** | Pointer to **[]string** | MetaData object names | [optional] 
**ObjectId** | Pointer to **[]int32** | MetaData object id&#39;s | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**CreatedAtSince** | Pointer to **time.Time** | Show create since | [optional] 
**CreatedAtTill** | Pointer to **time.Time** | Show created till | [optional] 

## Methods

### NewMetaDataFiltersModel

`func NewMetaDataFiltersModel() *MetaDataFiltersModel`

NewMetaDataFiltersModel instantiates a new MetaDataFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDataFiltersModelWithDefaults

`func NewMetaDataFiltersModelWithDefaults() *MetaDataFiltersModel`

NewMetaDataFiltersModelWithDefaults instantiates a new MetaDataFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetaDataFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetaDataFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetaDataFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetaDataFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetaDataFiltersModel) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaDataFiltersModel) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaDataFiltersModel) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaDataFiltersModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *MetaDataFiltersModel) GetObjectTypeName() []string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *MetaDataFiltersModel) GetObjectTypeNameOk() (*[]string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *MetaDataFiltersModel) SetObjectTypeName(v []string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *MetaDataFiltersModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetObjectId

`func (o *MetaDataFiltersModel) GetObjectId() []int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MetaDataFiltersModel) GetObjectIdOk() (*[]int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MetaDataFiltersModel) SetObjectId(v []int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *MetaDataFiltersModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *MetaDataFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *MetaDataFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *MetaDataFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *MetaDataFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *MetaDataFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *MetaDataFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *MetaDataFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *MetaDataFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetCreatedAtSince

`func (o *MetaDataFiltersModel) GetCreatedAtSince() time.Time`

GetCreatedAtSince returns the CreatedAtSince field if non-nil, zero value otherwise.

### GetCreatedAtSinceOk

`func (o *MetaDataFiltersModel) GetCreatedAtSinceOk() (*time.Time, bool)`

GetCreatedAtSinceOk returns a tuple with the CreatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtSince

`func (o *MetaDataFiltersModel) SetCreatedAtSince(v time.Time)`

SetCreatedAtSince sets CreatedAtSince field to given value.

### HasCreatedAtSince

`func (o *MetaDataFiltersModel) HasCreatedAtSince() bool`

HasCreatedAtSince returns a boolean if a field has been set.

### GetCreatedAtTill

`func (o *MetaDataFiltersModel) GetCreatedAtTill() time.Time`

GetCreatedAtTill returns the CreatedAtTill field if non-nil, zero value otherwise.

### GetCreatedAtTillOk

`func (o *MetaDataFiltersModel) GetCreatedAtTillOk() (*time.Time, bool)`

GetCreatedAtTillOk returns a tuple with the CreatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTill

`func (o *MetaDataFiltersModel) SetCreatedAtTill(v time.Time)`

SetCreatedAtTill sets CreatedAtTill field to given value.

### HasCreatedAtTill

`func (o *MetaDataFiltersModel) HasCreatedAtTill() bool`

HasCreatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


