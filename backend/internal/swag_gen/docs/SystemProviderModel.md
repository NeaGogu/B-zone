# SystemProviderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | The name of the System Provider | [optional] 

## Methods

### NewSystemProviderModel

`func NewSystemProviderModel() *SystemProviderModel`

NewSystemProviderModel instantiates a new SystemProviderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProviderModelWithDefaults

`func NewSystemProviderModelWithDefaults() *SystemProviderModel`

NewSystemProviderModelWithDefaults instantiates a new SystemProviderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemProviderModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemProviderModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemProviderModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SystemProviderModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SystemProviderModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemProviderModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemProviderModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemProviderModel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


