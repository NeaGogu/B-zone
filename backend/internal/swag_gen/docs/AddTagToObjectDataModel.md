# AddTagToObjectDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id of the tag which should be added to the object | [optional] 
**Name** | Pointer to **string** | The name of the tag which should be added to the object | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) | link to the tag which should be added to the object | [optional] 

## Methods

### NewAddTagToObjectDataModel

`func NewAddTagToObjectDataModel() *AddTagToObjectDataModel`

NewAddTagToObjectDataModel instantiates a new AddTagToObjectDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTagToObjectDataModelWithDefaults

`func NewAddTagToObjectDataModelWithDefaults() *AddTagToObjectDataModel`

NewAddTagToObjectDataModelWithDefaults instantiates a new AddTagToObjectDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddTagToObjectDataModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddTagToObjectDataModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddTagToObjectDataModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddTagToObjectDataModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddTagToObjectDataModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTagToObjectDataModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTagToObjectDataModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddTagToObjectDataModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLinks

`func (o *AddTagToObjectDataModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AddTagToObjectDataModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AddTagToObjectDataModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AddTagToObjectDataModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


