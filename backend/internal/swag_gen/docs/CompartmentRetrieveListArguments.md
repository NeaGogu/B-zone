# CompartmentRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**CompartmentOptionsModel**](CompartmentOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**CompartmentFiltersModel**](CompartmentFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**AsList** | Pointer to **bool** |  | [optional] 

## Methods

### NewCompartmentRetrieveListArguments

`func NewCompartmentRetrieveListArguments() *CompartmentRetrieveListArguments`

NewCompartmentRetrieveListArguments instantiates a new CompartmentRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartmentRetrieveListArgumentsWithDefaults

`func NewCompartmentRetrieveListArgumentsWithDefaults() *CompartmentRetrieveListArguments`

NewCompartmentRetrieveListArgumentsWithDefaults instantiates a new CompartmentRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *CompartmentRetrieveListArguments) GetOptions() CompartmentOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CompartmentRetrieveListArguments) GetOptionsOk() (*CompartmentOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CompartmentRetrieveListArguments) SetOptions(v CompartmentOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CompartmentRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *CompartmentRetrieveListArguments) GetFilters() CompartmentFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CompartmentRetrieveListArguments) GetFiltersOk() (*CompartmentFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CompartmentRetrieveListArguments) SetFilters(v CompartmentFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CompartmentRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *CompartmentRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CompartmentRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CompartmentRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CompartmentRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CompartmentRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CompartmentRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CompartmentRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CompartmentRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *CompartmentRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *CompartmentRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *CompartmentRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *CompartmentRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *CompartmentRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *CompartmentRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *CompartmentRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *CompartmentRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *CompartmentRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *CompartmentRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *CompartmentRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *CompartmentRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetAsList

`func (o *CompartmentRetrieveListArguments) GetAsList() bool`

GetAsList returns the AsList field if non-nil, zero value otherwise.

### GetAsListOk

`func (o *CompartmentRetrieveListArguments) GetAsListOk() (*bool, bool)`

GetAsListOk returns a tuple with the AsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsList

`func (o *CompartmentRetrieveListArguments) SetAsList(v bool)`

SetAsList sets AsList field to given value.

### HasAsList

`func (o *CompartmentRetrieveListArguments) HasAsList() bool`

HasAsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

