# BrandFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**Name** | Pointer to **[]string** | Brand name | [optional] 
**PortalName** | Pointer to **[]string** | When getting the portal settings for a brand: the portal name | [optional] 
**Links** | Pointer to **[]map[string]interface{}** | Activity Link ids | [optional] 

## Methods

### NewBrandFiltersModel

`func NewBrandFiltersModel() *BrandFiltersModel`

NewBrandFiltersModel instantiates a new BrandFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandFiltersModelWithDefaults

`func NewBrandFiltersModelWithDefaults() *BrandFiltersModel`

NewBrandFiltersModelWithDefaults instantiates a new BrandFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAtSince

`func (o *BrandFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *BrandFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *BrandFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *BrandFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *BrandFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *BrandFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *BrandFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *BrandFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetName

`func (o *BrandFiltersModel) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrandFiltersModel) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrandFiltersModel) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *BrandFiltersModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortalName

`func (o *BrandFiltersModel) GetPortalName() []string`

GetPortalName returns the PortalName field if non-nil, zero value otherwise.

### GetPortalNameOk

`func (o *BrandFiltersModel) GetPortalNameOk() (*[]string, bool)`

GetPortalNameOk returns a tuple with the PortalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalName

`func (o *BrandFiltersModel) SetPortalName(v []string)`

SetPortalName sets PortalName field to given value.

### HasPortalName

`func (o *BrandFiltersModel) HasPortalName() bool`

HasPortalName returns a boolean if a field has been set.

### GetLinks

`func (o *BrandFiltersModel) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrandFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrandFiltersModel) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrandFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


