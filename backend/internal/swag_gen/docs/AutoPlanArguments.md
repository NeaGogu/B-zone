# AutoPlanArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AutoPlanDataModel**](AutoPlanDataModel.md) |  | [optional] 
**Options** | Pointer to [**AutoPlanOptionsModel**](AutoPlanOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**AutoPlanFiltersModel**](AutoPlanFiltersModel.md) |  | [optional] 

## Methods

### NewAutoPlanArguments

`func NewAutoPlanArguments() *AutoPlanArguments`

NewAutoPlanArguments instantiates a new AutoPlanArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoPlanArgumentsWithDefaults

`func NewAutoPlanArgumentsWithDefaults() *AutoPlanArguments`

NewAutoPlanArgumentsWithDefaults instantiates a new AutoPlanArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutoPlanArguments) GetData() AutoPlanDataModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutoPlanArguments) GetDataOk() (*AutoPlanDataModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutoPlanArguments) SetData(v AutoPlanDataModel)`

SetData sets Data field to given value.

### HasData

`func (o *AutoPlanArguments) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOptions

`func (o *AutoPlanArguments) GetOptions() AutoPlanOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AutoPlanArguments) GetOptionsOk() (*AutoPlanOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AutoPlanArguments) SetOptions(v AutoPlanOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AutoPlanArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *AutoPlanArguments) GetFilters() AutoPlanFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AutoPlanArguments) GetFiltersOk() (*AutoPlanFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AutoPlanArguments) SetFilters(v AutoPlanFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AutoPlanArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


