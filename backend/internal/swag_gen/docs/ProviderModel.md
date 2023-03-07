# ProviderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | The name of the external Provider | [optional] 

## Methods

### NewProviderModel

`func NewProviderModel() *ProviderModel`

NewProviderModel instantiates a new ProviderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderModelWithDefaults

`func NewProviderModelWithDefaults() *ProviderModel`

NewProviderModelWithDefaults instantiates a new ProviderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProviderModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderModel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

