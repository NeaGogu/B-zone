# ActivityEtaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID of Activity | [optional] [readonly] 
**EtaDateTime** | Pointer to **time.Time** |  | [optional] 
**SequenceNr** | Pointer to **int32** | Position of activity on route | [optional] 

## Methods

### NewActivityEtaModel

`func NewActivityEtaModel() *ActivityEtaModel`

NewActivityEtaModel instantiates a new ActivityEtaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityEtaModelWithDefaults

`func NewActivityEtaModelWithDefaults() *ActivityEtaModel`

NewActivityEtaModelWithDefaults instantiates a new ActivityEtaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityEtaModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityEtaModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityEtaModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityEtaModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEtaDateTime

`func (o *ActivityEtaModel) GetEtaDateTime() time.Time`

GetEtaDateTime returns the EtaDateTime field if non-nil, zero value otherwise.

### GetEtaDateTimeOk

`func (o *ActivityEtaModel) GetEtaDateTimeOk() (*time.Time, bool)`

GetEtaDateTimeOk returns a tuple with the EtaDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtaDateTime

`func (o *ActivityEtaModel) SetEtaDateTime(v time.Time)`

SetEtaDateTime sets EtaDateTime field to given value.

### HasEtaDateTime

`func (o *ActivityEtaModel) HasEtaDateTime() bool`

HasEtaDateTime returns a boolean if a field has been set.

### GetSequenceNr

`func (o *ActivityEtaModel) GetSequenceNr() int32`

GetSequenceNr returns the SequenceNr field if non-nil, zero value otherwise.

### GetSequenceNrOk

`func (o *ActivityEtaModel) GetSequenceNrOk() (*int32, bool)`

GetSequenceNrOk returns a tuple with the SequenceNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNr

`func (o *ActivityEtaModel) SetSequenceNr(v int32)`

SetSequenceNr sets SequenceNr field to given value.

### HasSequenceNr

`func (o *ActivityEtaModel) HasSequenceNr() bool`

HasSequenceNr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


