# AddTagToObjectFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id of the object where the tag should be added to | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) | The links to find the object where the tag should be added to | [optional] 
**ObjectType** | Pointer to **int64** | The object type id of the object where the tag should be added to | [optional] 
**ObjectTypeName** | Pointer to **string** | The object type name of the object where the tag should be added to | [optional] 

## Methods

### NewAddTagToObjectFiltersModel

`func NewAddTagToObjectFiltersModel() *AddTagToObjectFiltersModel`

NewAddTagToObjectFiltersModel instantiates a new AddTagToObjectFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTagToObjectFiltersModelWithDefaults

`func NewAddTagToObjectFiltersModelWithDefaults() *AddTagToObjectFiltersModel`

NewAddTagToObjectFiltersModelWithDefaults instantiates a new AddTagToObjectFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddTagToObjectFiltersModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddTagToObjectFiltersModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddTagToObjectFiltersModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddTagToObjectFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *AddTagToObjectFiltersModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AddTagToObjectFiltersModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AddTagToObjectFiltersModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AddTagToObjectFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetObjectType

`func (o *AddTagToObjectFiltersModel) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AddTagToObjectFiltersModel) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AddTagToObjectFiltersModel) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AddTagToObjectFiltersModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *AddTagToObjectFiltersModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *AddTagToObjectFiltersModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *AddTagToObjectFiltersModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *AddTagToObjectFiltersModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


