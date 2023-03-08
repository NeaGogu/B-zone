# DriverFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Driver id&#39;s | [optional] 
**PauseId** | Pointer to **[]int32** | ids of pause schemes applied to drivers | [optional] 
**Links** | Pointer to **[]map[string]interface{}** | Driver Link ids | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 

## Methods

### NewDriverFiltersModel

`func NewDriverFiltersModel() *DriverFiltersModel`

NewDriverFiltersModel instantiates a new DriverFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverFiltersModelWithDefaults

`func NewDriverFiltersModelWithDefaults() *DriverFiltersModel`

NewDriverFiltersModelWithDefaults instantiates a new DriverFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriverFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *DriverFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPauseId

`func (o *DriverFiltersModel) GetPauseId() []int32`

GetPauseId returns the PauseId field if non-nil, zero value otherwise.

### GetPauseIdOk

`func (o *DriverFiltersModel) GetPauseIdOk() (*[]int32, bool)`

GetPauseIdOk returns a tuple with the PauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseId

`func (o *DriverFiltersModel) SetPauseId(v []int32)`

SetPauseId sets PauseId field to given value.

### HasPauseId

`func (o *DriverFiltersModel) HasPauseId() bool`

HasPauseId returns a boolean if a field has been set.

### GetLinks

`func (o *DriverFiltersModel) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DriverFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DriverFiltersModel) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DriverFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTagNames

`func (o *DriverFiltersModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *DriverFiltersModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *DriverFiltersModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *DriverFiltersModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *DriverFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *DriverFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *DriverFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *DriverFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *DriverFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *DriverFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *DriverFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *DriverFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


