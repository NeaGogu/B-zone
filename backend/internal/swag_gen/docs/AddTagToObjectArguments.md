# AddTagToObjectArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AddTagToObjectDataModel**](AddTagToObjectDataModel.md) |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**Filters** | Pointer to [**AddTagToObjectFiltersModel**](AddTagToObjectFiltersModel.md) |  | [optional] 

## Methods

### NewAddTagToObjectArguments

`func NewAddTagToObjectArguments() *AddTagToObjectArguments`

NewAddTagToObjectArguments instantiates a new AddTagToObjectArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTagToObjectArgumentsWithDefaults

`func NewAddTagToObjectArgumentsWithDefaults() *AddTagToObjectArguments`

NewAddTagToObjectArgumentsWithDefaults instantiates a new AddTagToObjectArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddTagToObjectArguments) GetData() AddTagToObjectDataModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddTagToObjectArguments) GetDataOk() (*AddTagToObjectDataModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddTagToObjectArguments) SetData(v AddTagToObjectDataModel)`

SetData sets Data field to given value.

### HasData

`func (o *AddTagToObjectArguments) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOptions

`func (o *AddTagToObjectArguments) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AddTagToObjectArguments) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AddTagToObjectArguments) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AddTagToObjectArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *AddTagToObjectArguments) GetFilters() AddTagToObjectFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AddTagToObjectArguments) GetFiltersOk() (*AddTagToObjectFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AddTagToObjectArguments) SetFilters(v AddTagToObjectFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AddTagToObjectArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


