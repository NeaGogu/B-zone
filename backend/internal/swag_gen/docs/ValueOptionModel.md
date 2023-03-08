# ValueOptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the value option (just as a discription) | [optional] 
**Value** | Pointer to **string** | The value of the value option (should be posted as value for the setting) | [optional] 

## Methods

### NewValueOptionModel

`func NewValueOptionModel() *ValueOptionModel`

NewValueOptionModel instantiates a new ValueOptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueOptionModelWithDefaults

`func NewValueOptionModelWithDefaults() *ValueOptionModel`

NewValueOptionModelWithDefaults instantiates a new ValueOptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValueOptionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValueOptionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValueOptionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ValueOptionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ValueOptionModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValueOptionModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValueOptionModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ValueOptionModel) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


