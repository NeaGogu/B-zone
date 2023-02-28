# ActivityBulkUpdateArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**ActivityFiltersModel**](ActivityFiltersModel.md) |  | [optional] 
**UpdateInfo** | Pointer to [**ActivityBulkUpdateModel**](ActivityBulkUpdateModel.md) |  | [optional] 

## Methods

### NewActivityBulkUpdateArguments

`func NewActivityBulkUpdateArguments() *ActivityBulkUpdateArguments`

NewActivityBulkUpdateArguments instantiates a new ActivityBulkUpdateArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityBulkUpdateArgumentsWithDefaults

`func NewActivityBulkUpdateArgumentsWithDefaults() *ActivityBulkUpdateArguments`

NewActivityBulkUpdateArgumentsWithDefaults instantiates a new ActivityBulkUpdateArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ActivityBulkUpdateArguments) GetFilters() ActivityFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ActivityBulkUpdateArguments) GetFiltersOk() (*ActivityFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ActivityBulkUpdateArguments) SetFilters(v ActivityFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ActivityBulkUpdateArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUpdateInfo

`func (o *ActivityBulkUpdateArguments) GetUpdateInfo() ActivityBulkUpdateModel`

GetUpdateInfo returns the UpdateInfo field if non-nil, zero value otherwise.

### GetUpdateInfoOk

`func (o *ActivityBulkUpdateArguments) GetUpdateInfoOk() (*ActivityBulkUpdateModel, bool)`

GetUpdateInfoOk returns a tuple with the UpdateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInfo

`func (o *ActivityBulkUpdateArguments) SetUpdateInfo(v ActivityBulkUpdateModel)`

SetUpdateInfo sets UpdateInfo field to given value.

### HasUpdateInfo

`func (o *ActivityBulkUpdateArguments) HasUpdateInfo() bool`

HasUpdateInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


