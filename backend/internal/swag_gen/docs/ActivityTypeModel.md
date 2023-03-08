# ActivityTypeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Activity type ID. 1 (:pickup), 2 (:dropoff), 3 (:route_start), 4 (:route_end), 5 (:car_start), 6 (:car_end), 9 (:driver_start), 10 (:driver_end), 11 (:stop), 13 (:pause), 14 (:gas_refill), 15 (:maintenance_planned), 16 (:maintenance_unplanned), 17 (:user_other), 18 (:unplanned_stop), 19 (:wait), 20 (:breakdown), 21 (:sanitary), 26 (:maintenance), 27 (:car_wash), 28 (:depot), 29 (:combi), 30 (:interior_wash), 31 (:other), 32 (:bundled) | [optional] [readonly] 
**Name** | Pointer to **string** | Activity Type Name. breakdown (:20), bundled (:32), car_end (:6), car_start (:5), car_wash (:27), combi (:29), depot (:28), driver_end (:10), driver_start (:9), dropoff (:2), gas_refill (:14), interior_wash (:30), maintenance (:26), maintenance_planned (:15), maintenance_unplanned (:16), other (:31), pause (:13), pickup (:1), route_end (:4), route_start (:3), sanitary (:21), stop (:11), unplanned_stop (:18), user_other (:17), wait (:19) | [optional] [readonly] 
**Special** | Pointer to **bool** | special activities are ones which are not planned but simly added by the driver during the route execution | [optional] 
**AssignmentEntry** | Pointer to **bool** | assignment_entry activity types are used for order entry forms | [optional] 

## Methods

### NewActivityTypeModel

`func NewActivityTypeModel() *ActivityTypeModel`

NewActivityTypeModel instantiates a new ActivityTypeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTypeModelWithDefaults

`func NewActivityTypeModelWithDefaults() *ActivityTypeModel`

NewActivityTypeModelWithDefaults instantiates a new ActivityTypeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityTypeModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityTypeModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityTypeModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityTypeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivityTypeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityTypeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityTypeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActivityTypeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpecial

`func (o *ActivityTypeModel) GetSpecial() bool`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *ActivityTypeModel) GetSpecialOk() (*bool, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *ActivityTypeModel) SetSpecial(v bool)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *ActivityTypeModel) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.

### GetAssignmentEntry

`func (o *ActivityTypeModel) GetAssignmentEntry() bool`

GetAssignmentEntry returns the AssignmentEntry field if non-nil, zero value otherwise.

### GetAssignmentEntryOk

`func (o *ActivityTypeModel) GetAssignmentEntryOk() (*bool, bool)`

GetAssignmentEntryOk returns a tuple with the AssignmentEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentEntry

`func (o *ActivityTypeModel) SetAssignmentEntry(v bool)`

SetAssignmentEntry sets AssignmentEntry field to given value.

### HasAssignmentEntry

`func (o *ActivityTypeModel) HasAssignmentEntry() bool`

HasAssignmentEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


