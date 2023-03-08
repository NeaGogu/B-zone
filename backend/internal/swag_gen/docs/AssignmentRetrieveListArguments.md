# AssignmentRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**AssignmentOptionsModel**](AssignmentOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**AssignmentFiltersModel**](AssignmentFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column. Option include_party_name has to be set for party_name as the sorting column. Option include_booking_account has to be set for account_name as the sorting column. | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**AsList** | Pointer to **bool** |  | [optional] 
**CountOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewAssignmentRetrieveListArguments

`func NewAssignmentRetrieveListArguments() *AssignmentRetrieveListArguments`

NewAssignmentRetrieveListArguments instantiates a new AssignmentRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentRetrieveListArgumentsWithDefaults

`func NewAssignmentRetrieveListArgumentsWithDefaults() *AssignmentRetrieveListArguments`

NewAssignmentRetrieveListArgumentsWithDefaults instantiates a new AssignmentRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *AssignmentRetrieveListArguments) GetOptions() AssignmentOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AssignmentRetrieveListArguments) GetOptionsOk() (*AssignmentOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AssignmentRetrieveListArguments) SetOptions(v AssignmentOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AssignmentRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *AssignmentRetrieveListArguments) GetFilters() AssignmentFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AssignmentRetrieveListArguments) GetFiltersOk() (*AssignmentFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AssignmentRetrieveListArguments) SetFilters(v AssignmentFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AssignmentRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *AssignmentRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AssignmentRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AssignmentRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AssignmentRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *AssignmentRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AssignmentRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AssignmentRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AssignmentRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSearchText

`func (o *AssignmentRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *AssignmentRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *AssignmentRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *AssignmentRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetSortingColumn

`func (o *AssignmentRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *AssignmentRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *AssignmentRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *AssignmentRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *AssignmentRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *AssignmentRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *AssignmentRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *AssignmentRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetAsList

`func (o *AssignmentRetrieveListArguments) GetAsList() bool`

GetAsList returns the AsList field if non-nil, zero value otherwise.

### GetAsListOk

`func (o *AssignmentRetrieveListArguments) GetAsListOk() (*bool, bool)`

GetAsListOk returns a tuple with the AsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsList

`func (o *AssignmentRetrieveListArguments) SetAsList(v bool)`

SetAsList sets AsList field to given value.

### HasAsList

`func (o *AssignmentRetrieveListArguments) HasAsList() bool`

HasAsList returns a boolean if a field has been set.

### GetCountOnly

`func (o *AssignmentRetrieveListArguments) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *AssignmentRetrieveListArguments) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *AssignmentRetrieveListArguments) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *AssignmentRetrieveListArguments) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


