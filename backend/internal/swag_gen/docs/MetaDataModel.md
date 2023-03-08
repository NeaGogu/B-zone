# MetaDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | The name of the MetaData field | [optional] 
**Value** | Pointer to **string** | The value of the MetaData field | [optional] 
**ObjectId** | Pointer to **int64** | Object ID | [optional] 
**ObjectTypeName** | Pointer to **string** | Object type name | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 

## Methods

### NewMetaDataModel

`func NewMetaDataModel() *MetaDataModel`

NewMetaDataModel instantiates a new MetaDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDataModelWithDefaults

`func NewMetaDataModelWithDefaults() *MetaDataModel`

NewMetaDataModelWithDefaults instantiates a new MetaDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetaDataModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetaDataModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetaDataModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MetaDataModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetaDataModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaDataModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaDataModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaDataModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *MetaDataModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetaDataModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetaDataModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetaDataModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetObjectId

`func (o *MetaDataModel) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MetaDataModel) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MetaDataModel) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *MetaDataModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *MetaDataModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *MetaDataModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *MetaDataModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *MetaDataModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MetaDataModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetaDataModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetaDataModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MetaDataModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MetaDataModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MetaDataModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MetaDataModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MetaDataModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


