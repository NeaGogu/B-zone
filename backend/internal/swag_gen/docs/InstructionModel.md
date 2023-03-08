# InstructionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Content** | Pointer to **string** | Instructie voor chauffeur | [optional] 

## Methods

### NewInstructionModel

`func NewInstructionModel() *InstructionModel`

NewInstructionModel instantiates a new InstructionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstructionModelWithDefaults

`func NewInstructionModelWithDefaults() *InstructionModel`

NewInstructionModelWithDefaults instantiates a new InstructionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstructionModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstructionModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstructionModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstructionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *InstructionModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InstructionModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InstructionModel) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *InstructionModel) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


