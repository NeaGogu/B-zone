# LinkFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | Pointer to **[]string** | Unique Link Identifier(s) | [optional] 
**ObjectType** | Pointer to **[]int32** | array of object type id(s) | [optional] 
**ObjectTypeName** | Pointer to **[]string** | Object type names | [optional] 
**ProviderId** | Pointer to **[]int32** | array of provider id(s) | [optional] 

## Methods

### NewLinkFiltersModel

`func NewLinkFiltersModel() *LinkFiltersModel`

NewLinkFiltersModel instantiates a new LinkFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkFiltersModelWithDefaults

`func NewLinkFiltersModelWithDefaults() *LinkFiltersModel`

NewLinkFiltersModelWithDefaults instantiates a new LinkFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *LinkFiltersModel) GetLinkId() []string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *LinkFiltersModel) GetLinkIdOk() (*[]string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *LinkFiltersModel) SetLinkId(v []string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *LinkFiltersModel) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetObjectType

`func (o *LinkFiltersModel) GetObjectType() []int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LinkFiltersModel) GetObjectTypeOk() (*[]int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LinkFiltersModel) SetObjectType(v []int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *LinkFiltersModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *LinkFiltersModel) GetObjectTypeName() []string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *LinkFiltersModel) GetObjectTypeNameOk() (*[]string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *LinkFiltersModel) SetObjectTypeName(v []string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *LinkFiltersModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetProviderId

`func (o *LinkFiltersModel) GetProviderId() []int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *LinkFiltersModel) GetProviderIdOk() (*[]int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *LinkFiltersModel) SetProviderId(v []int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *LinkFiltersModel) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


