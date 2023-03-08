# ActivityBulkUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotAddressId** | Pointer to **int32** | Depot Address ID | [optional] 
**Priority** | Pointer to **int32** | Priority level. 1 for highest priority, 3 for lowest priority. Default &#x3D; 2 | [optional] 

## Methods

### NewActivityBulkUpdateModel

`func NewActivityBulkUpdateModel() *ActivityBulkUpdateModel`

NewActivityBulkUpdateModel instantiates a new ActivityBulkUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityBulkUpdateModelWithDefaults

`func NewActivityBulkUpdateModelWithDefaults() *ActivityBulkUpdateModel`

NewActivityBulkUpdateModelWithDefaults instantiates a new ActivityBulkUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotAddressId

`func (o *ActivityBulkUpdateModel) GetDepotAddressId() int32`

GetDepotAddressId returns the DepotAddressId field if non-nil, zero value otherwise.

### GetDepotAddressIdOk

`func (o *ActivityBulkUpdateModel) GetDepotAddressIdOk() (*int32, bool)`

GetDepotAddressIdOk returns a tuple with the DepotAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotAddressId

`func (o *ActivityBulkUpdateModel) SetDepotAddressId(v int32)`

SetDepotAddressId sets DepotAddressId field to given value.

### HasDepotAddressId

`func (o *ActivityBulkUpdateModel) HasDepotAddressId() bool`

HasDepotAddressId returns a boolean if a field has been set.

### GetPriority

`func (o *ActivityBulkUpdateModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ActivityBulkUpdateModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ActivityBulkUpdateModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ActivityBulkUpdateModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


