# TagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**ObjectType** | Pointer to **int64** | Object type ID | [optional] 
**TagTypeId** | Pointer to **int64** | Tag type ID | [optional] 
**TagTypeName** | Pointer to **string** | Tag type Name | [optional] 
**TagName** | Pointer to **string** | Tag Name (same as tag_type_name, but tag_type_name will become deprecated in the future) | [optional] 

## Methods

### NewTagModel

`func NewTagModel() *TagModel`

NewTagModel instantiates a new TagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagModelWithDefaults

`func NewTagModelWithDefaults() *TagModel`

NewTagModelWithDefaults instantiates a new TagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TagModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *TagModel) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TagModel) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TagModel) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *TagModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetTagTypeId

`func (o *TagModel) GetTagTypeId() int64`

GetTagTypeId returns the TagTypeId field if non-nil, zero value otherwise.

### GetTagTypeIdOk

`func (o *TagModel) GetTagTypeIdOk() (*int64, bool)`

GetTagTypeIdOk returns a tuple with the TagTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypeId

`func (o *TagModel) SetTagTypeId(v int64)`

SetTagTypeId sets TagTypeId field to given value.

### HasTagTypeId

`func (o *TagModel) HasTagTypeId() bool`

HasTagTypeId returns a boolean if a field has been set.

### GetTagTypeName

`func (o *TagModel) GetTagTypeName() string`

GetTagTypeName returns the TagTypeName field if non-nil, zero value otherwise.

### GetTagTypeNameOk

`func (o *TagModel) GetTagTypeNameOk() (*string, bool)`

GetTagTypeNameOk returns a tuple with the TagTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypeName

`func (o *TagModel) SetTagTypeName(v string)`

SetTagTypeName sets TagTypeName field to given value.

### HasTagTypeName

`func (o *TagModel) HasTagTypeName() bool`

HasTagTypeName returns a boolean if a field has been set.

### GetTagName

`func (o *TagModel) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *TagModel) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *TagModel) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *TagModel) HasTagName() bool`

HasTagName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


