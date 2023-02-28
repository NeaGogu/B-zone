# BrandColourModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**BrandId** | Pointer to **int64** | id for brand where this colour belongs to | [optional] [readonly] 
**Name** | Pointer to **string** | name | [optional] 
**Hex** | Pointer to **string** | hex | [optional] 
**R** | Pointer to **int32** | red value for rgba | [optional] 
**G** | Pointer to **int32** | green value for rgba | [optional] 
**B** | Pointer to **int32** | blue value for rgba | [optional] 
**A** | Pointer to **float64** | alpha value (opacity/transparency) for rgba | [optional] 

## Methods

### NewBrandColourModel

`func NewBrandColourModel() *BrandColourModel`

NewBrandColourModel instantiates a new BrandColourModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandColourModelWithDefaults

`func NewBrandColourModelWithDefaults() *BrandColourModel`

NewBrandColourModelWithDefaults instantiates a new BrandColourModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BrandColourModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrandColourModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrandColourModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BrandColourModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBrandId

`func (o *BrandColourModel) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *BrandColourModel) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *BrandColourModel) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *BrandColourModel) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetName

`func (o *BrandColourModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrandColourModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrandColourModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BrandColourModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHex

`func (o *BrandColourModel) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *BrandColourModel) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *BrandColourModel) SetHex(v string)`

SetHex sets Hex field to given value.

### HasHex

`func (o *BrandColourModel) HasHex() bool`

HasHex returns a boolean if a field has been set.

### GetR

`func (o *BrandColourModel) GetR() int32`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *BrandColourModel) GetROk() (*int32, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *BrandColourModel) SetR(v int32)`

SetR sets R field to given value.

### HasR

`func (o *BrandColourModel) HasR() bool`

HasR returns a boolean if a field has been set.

### GetG

`func (o *BrandColourModel) GetG() int32`

GetG returns the G field if non-nil, zero value otherwise.

### GetGOk

`func (o *BrandColourModel) GetGOk() (*int32, bool)`

GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetG

`func (o *BrandColourModel) SetG(v int32)`

SetG sets G field to given value.

### HasG

`func (o *BrandColourModel) HasG() bool`

HasG returns a boolean if a field has been set.

### GetB

`func (o *BrandColourModel) GetB() int32`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *BrandColourModel) GetBOk() (*int32, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *BrandColourModel) SetB(v int32)`

SetB sets B field to given value.

### HasB

`func (o *BrandColourModel) HasB() bool`

HasB returns a boolean if a field has been set.

### GetA

`func (o *BrandColourModel) GetA() float64`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *BrandColourModel) GetAOk() (*float64, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *BrandColourModel) SetA(v float64)`

SetA sets A field to given value.

### HasA

`func (o *BrandColourModel) HasA() bool`

HasA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


