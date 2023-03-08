# VehicleTypeUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** | Ready | [optional] 
**Message** | Pointer to **string** | Message describing the code | [optional] 
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVehicleTypeUpdateResponse

`func NewVehicleTypeUpdateResponse() *VehicleTypeUpdateResponse`

NewVehicleTypeUpdateResponse instantiates a new VehicleTypeUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleTypeUpdateResponseWithDefaults

`func NewVehicleTypeUpdateResponseWithDefaults() *VehicleTypeUpdateResponse`

NewVehicleTypeUpdateResponseWithDefaults instantiates a new VehicleTypeUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VehicleTypeUpdateResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VehicleTypeUpdateResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VehicleTypeUpdateResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *VehicleTypeUpdateResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *VehicleTypeUpdateResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VehicleTypeUpdateResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VehicleTypeUpdateResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VehicleTypeUpdateResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *VehicleTypeUpdateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VehicleTypeUpdateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VehicleTypeUpdateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VehicleTypeUpdateResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAdditionalData

`func (o *VehicleTypeUpdateResponse) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *VehicleTypeUpdateResponse) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *VehicleTypeUpdateResponse) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *VehicleTypeUpdateResponse) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


