# ConfigModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to **string** |  | [optional] 
**KeyRing** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigModel

`func NewConfigModel() *ConfigModel`

NewConfigModel instantiates a new ConfigModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigModelWithDefaults

`func NewConfigModelWithDefaults() *ConfigModel`

NewConfigModelWithDefaults instantiates a new ConfigModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *ConfigModel) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigModel) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigModel) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ConfigModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetKeyRing

`func (o *ConfigModel) GetKeyRing() string`

GetKeyRing returns the KeyRing field if non-nil, zero value otherwise.

### GetKeyRingOk

`func (o *ConfigModel) GetKeyRingOk() (*string, bool)`

GetKeyRingOk returns a tuple with the KeyRing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRing

`func (o *ConfigModel) SetKeyRing(v string)`

SetKeyRing sets KeyRing field to given value.

### HasKeyRing

`func (o *ConfigModel) HasKeyRing() bool`

HasKeyRing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


