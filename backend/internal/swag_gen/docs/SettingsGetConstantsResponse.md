# SettingsGetConstantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IrregularityCategory** | Pointer to **map[string]string** | Irregularity Categories | [optional] 
**IrregularityCause** | Pointer to **map[string]string** | Irregularity Causes | [optional] 
**InstructionType** | Pointer to **map[string]string** | Instruction Types | [optional] 
**ReferenceType** | Pointer to **map[string]string** | Reference Types | [optional] 
**RelationType** | Pointer to **map[string]string** | Relation Types | [optional] 
**WeekDay** | Pointer to **map[string]string** | Names of the days in the week | [optional] 
**ActionType** | Pointer to **map[string]string** | Action Types | [optional] 
**ActivityType** | Pointer to **map[string]string** | Activity Types | [optional] 
**SpecialActivity** | Pointer to **map[string]string** | Special Activities | [optional] 
**AeActivity** | Pointer to **map[string]string** | EA Activity | [optional] 
**AddressType** | Pointer to **map[string]string** | Address Types | [optional] 
**AeAddressType** | Pointer to **map[string]string** | EA Address Types | [optional] 
**CommunicationDeliveryMethod** | Pointer to **map[string]string** | Communication Delivery Methods | [optional] 
**CommunicationMessageType** | Pointer to **map[string]string** | Communication Message Types | [optional] 
**Counter** | Pointer to **map[string]string** | Counters | [optional] 
**EquipmentType** | Pointer to **map[string]string** | Equipment Types | [optional] 
**FileType** | Pointer to **map[string]string** | File Types | [optional] 
**Incoterm** | Pointer to **map[string]string** | Incoterms | [optional] 
**LogType** | Pointer to **map[string]string** | Log Types | [optional] 
**ObjectType** | Pointer to **map[string]string** | Object Types | [optional] 
**OptimisationObjective** | Pointer to **map[string]string** | Optimisation Objectives | [optional] 
**PartyType** | Pointer to **map[string]string** | Party Types | [optional] 
**PhoneNrType** | Pointer to **map[string]string** | Phone NR Types | [optional] 
**RecurrencePeriod** | Pointer to **map[string]string** | Recurrence Periods | [optional] 
**RecurrenceType** | Pointer to **map[string]string** | Recurrence Types | [optional] 
**Role** | Pointer to **map[string]string** | Roles | [optional] 
**SaywhenStatus** | Pointer to **map[string]string** | SayWhen Statuses | [optional] 
**Settings** | Pointer to **map[string]string** | Settings | [optional] 
**SettingsGroup** | Pointer to **map[string]string** | Settings Groups | [optional] 
**StatusType** | Pointer to **map[string]string** | Status Types | [optional] 
**RouteStatus** | Pointer to **map[string]string** | Route Statuses | [optional] 
**ActivityStatus** | Pointer to **map[string]string** | Activity Statuses | [optional] 
**ShipmentStatus** | Pointer to **map[string]string** | Shipment Statuses | [optional] 
**PackageLineStatus** | Pointer to **map[string]string** | Package Line Statuses | [optional] 
**AssignmentStatus** | Pointer to **map[string]string** | Assignment Statuses | [optional] 
**MessageStatus** | Pointer to **map[string]string** | Message Statuses | [optional] 
**UserStatus** | Pointer to **map[string]string** | User Statuses | [optional] 
**JobStatus** | Pointer to **map[string]string** | Job Statuses | [optional] 
**ApiRequestStatus** | Pointer to **map[string]string** | API Request Statuses | [optional] 
**SystemProvider** | Pointer to **map[string]string** | System Providers | [optional] 
**Provider** | Pointer to **map[string]string** | Providers | [optional] 
**TimeSlotType** | Pointer to **map[string]string** | Time Slot Types | [optional] 
**TransactionType** | Pointer to **map[string]string** | Transaction Types | [optional] 
**Uom** | Pointer to **map[string]string** | UOM | [optional] 
**UomGroup** | Pointer to **map[string]string** | UOM Groups | [optional] 
**VehicleType** | Pointer to **map[string]string** | Vehicle Types | [optional] 
**PackageType** | Pointer to **map[string]string** | Package Types | [optional] 

## Methods

### NewSettingsGetConstantsResponse

`func NewSettingsGetConstantsResponse() *SettingsGetConstantsResponse`

NewSettingsGetConstantsResponse instantiates a new SettingsGetConstantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsGetConstantsResponseWithDefaults

`func NewSettingsGetConstantsResponseWithDefaults() *SettingsGetConstantsResponse`

NewSettingsGetConstantsResponseWithDefaults instantiates a new SettingsGetConstantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIrregularityCategory

`func (o *SettingsGetConstantsResponse) GetIrregularityCategory() map[string]string`

GetIrregularityCategory returns the IrregularityCategory field if non-nil, zero value otherwise.

### GetIrregularityCategoryOk

`func (o *SettingsGetConstantsResponse) GetIrregularityCategoryOk() (*map[string]string, bool)`

GetIrregularityCategoryOk returns a tuple with the IrregularityCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrregularityCategory

`func (o *SettingsGetConstantsResponse) SetIrregularityCategory(v map[string]string)`

SetIrregularityCategory sets IrregularityCategory field to given value.

### HasIrregularityCategory

`func (o *SettingsGetConstantsResponse) HasIrregularityCategory() bool`

HasIrregularityCategory returns a boolean if a field has been set.

### GetIrregularityCause

`func (o *SettingsGetConstantsResponse) GetIrregularityCause() map[string]string`

GetIrregularityCause returns the IrregularityCause field if non-nil, zero value otherwise.

### GetIrregularityCauseOk

`func (o *SettingsGetConstantsResponse) GetIrregularityCauseOk() (*map[string]string, bool)`

GetIrregularityCauseOk returns a tuple with the IrregularityCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrregularityCause

`func (o *SettingsGetConstantsResponse) SetIrregularityCause(v map[string]string)`

SetIrregularityCause sets IrregularityCause field to given value.

### HasIrregularityCause

`func (o *SettingsGetConstantsResponse) HasIrregularityCause() bool`

HasIrregularityCause returns a boolean if a field has been set.

### GetInstructionType

`func (o *SettingsGetConstantsResponse) GetInstructionType() map[string]string`

GetInstructionType returns the InstructionType field if non-nil, zero value otherwise.

### GetInstructionTypeOk

`func (o *SettingsGetConstantsResponse) GetInstructionTypeOk() (*map[string]string, bool)`

GetInstructionTypeOk returns a tuple with the InstructionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionType

`func (o *SettingsGetConstantsResponse) SetInstructionType(v map[string]string)`

SetInstructionType sets InstructionType field to given value.

### HasInstructionType

`func (o *SettingsGetConstantsResponse) HasInstructionType() bool`

HasInstructionType returns a boolean if a field has been set.

### GetReferenceType

`func (o *SettingsGetConstantsResponse) GetReferenceType() map[string]string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *SettingsGetConstantsResponse) GetReferenceTypeOk() (*map[string]string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *SettingsGetConstantsResponse) SetReferenceType(v map[string]string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *SettingsGetConstantsResponse) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetRelationType

`func (o *SettingsGetConstantsResponse) GetRelationType() map[string]string`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *SettingsGetConstantsResponse) GetRelationTypeOk() (*map[string]string, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *SettingsGetConstantsResponse) SetRelationType(v map[string]string)`

SetRelationType sets RelationType field to given value.

### HasRelationType

`func (o *SettingsGetConstantsResponse) HasRelationType() bool`

HasRelationType returns a boolean if a field has been set.

### GetWeekDay

`func (o *SettingsGetConstantsResponse) GetWeekDay() map[string]string`

GetWeekDay returns the WeekDay field if non-nil, zero value otherwise.

### GetWeekDayOk

`func (o *SettingsGetConstantsResponse) GetWeekDayOk() (*map[string]string, bool)`

GetWeekDayOk returns a tuple with the WeekDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDay

`func (o *SettingsGetConstantsResponse) SetWeekDay(v map[string]string)`

SetWeekDay sets WeekDay field to given value.

### HasWeekDay

`func (o *SettingsGetConstantsResponse) HasWeekDay() bool`

HasWeekDay returns a boolean if a field has been set.

### GetActionType

`func (o *SettingsGetConstantsResponse) GetActionType() map[string]string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *SettingsGetConstantsResponse) GetActionTypeOk() (*map[string]string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *SettingsGetConstantsResponse) SetActionType(v map[string]string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *SettingsGetConstantsResponse) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActivityType

`func (o *SettingsGetConstantsResponse) GetActivityType() map[string]string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *SettingsGetConstantsResponse) GetActivityTypeOk() (*map[string]string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *SettingsGetConstantsResponse) SetActivityType(v map[string]string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *SettingsGetConstantsResponse) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetSpecialActivity

`func (o *SettingsGetConstantsResponse) GetSpecialActivity() map[string]string`

GetSpecialActivity returns the SpecialActivity field if non-nil, zero value otherwise.

### GetSpecialActivityOk

`func (o *SettingsGetConstantsResponse) GetSpecialActivityOk() (*map[string]string, bool)`

GetSpecialActivityOk returns a tuple with the SpecialActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialActivity

`func (o *SettingsGetConstantsResponse) SetSpecialActivity(v map[string]string)`

SetSpecialActivity sets SpecialActivity field to given value.

### HasSpecialActivity

`func (o *SettingsGetConstantsResponse) HasSpecialActivity() bool`

HasSpecialActivity returns a boolean if a field has been set.

### GetAeActivity

`func (o *SettingsGetConstantsResponse) GetAeActivity() map[string]string`

GetAeActivity returns the AeActivity field if non-nil, zero value otherwise.

### GetAeActivityOk

`func (o *SettingsGetConstantsResponse) GetAeActivityOk() (*map[string]string, bool)`

GetAeActivityOk returns a tuple with the AeActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeActivity

`func (o *SettingsGetConstantsResponse) SetAeActivity(v map[string]string)`

SetAeActivity sets AeActivity field to given value.

### HasAeActivity

`func (o *SettingsGetConstantsResponse) HasAeActivity() bool`

HasAeActivity returns a boolean if a field has been set.

### GetAddressType

`func (o *SettingsGetConstantsResponse) GetAddressType() map[string]string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *SettingsGetConstantsResponse) GetAddressTypeOk() (*map[string]string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *SettingsGetConstantsResponse) SetAddressType(v map[string]string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *SettingsGetConstantsResponse) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetAeAddressType

`func (o *SettingsGetConstantsResponse) GetAeAddressType() map[string]string`

GetAeAddressType returns the AeAddressType field if non-nil, zero value otherwise.

### GetAeAddressTypeOk

`func (o *SettingsGetConstantsResponse) GetAeAddressTypeOk() (*map[string]string, bool)`

GetAeAddressTypeOk returns a tuple with the AeAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeAddressType

`func (o *SettingsGetConstantsResponse) SetAeAddressType(v map[string]string)`

SetAeAddressType sets AeAddressType field to given value.

### HasAeAddressType

`func (o *SettingsGetConstantsResponse) HasAeAddressType() bool`

HasAeAddressType returns a boolean if a field has been set.

### GetCommunicationDeliveryMethod

`func (o *SettingsGetConstantsResponse) GetCommunicationDeliveryMethod() map[string]string`

GetCommunicationDeliveryMethod returns the CommunicationDeliveryMethod field if non-nil, zero value otherwise.

### GetCommunicationDeliveryMethodOk

`func (o *SettingsGetConstantsResponse) GetCommunicationDeliveryMethodOk() (*map[string]string, bool)`

GetCommunicationDeliveryMethodOk returns a tuple with the CommunicationDeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDeliveryMethod

`func (o *SettingsGetConstantsResponse) SetCommunicationDeliveryMethod(v map[string]string)`

SetCommunicationDeliveryMethod sets CommunicationDeliveryMethod field to given value.

### HasCommunicationDeliveryMethod

`func (o *SettingsGetConstantsResponse) HasCommunicationDeliveryMethod() bool`

HasCommunicationDeliveryMethod returns a boolean if a field has been set.

### GetCommunicationMessageType

`func (o *SettingsGetConstantsResponse) GetCommunicationMessageType() map[string]string`

GetCommunicationMessageType returns the CommunicationMessageType field if non-nil, zero value otherwise.

### GetCommunicationMessageTypeOk

`func (o *SettingsGetConstantsResponse) GetCommunicationMessageTypeOk() (*map[string]string, bool)`

GetCommunicationMessageTypeOk returns a tuple with the CommunicationMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationMessageType

`func (o *SettingsGetConstantsResponse) SetCommunicationMessageType(v map[string]string)`

SetCommunicationMessageType sets CommunicationMessageType field to given value.

### HasCommunicationMessageType

`func (o *SettingsGetConstantsResponse) HasCommunicationMessageType() bool`

HasCommunicationMessageType returns a boolean if a field has been set.

### GetCounter

`func (o *SettingsGetConstantsResponse) GetCounter() map[string]string`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *SettingsGetConstantsResponse) GetCounterOk() (*map[string]string, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *SettingsGetConstantsResponse) SetCounter(v map[string]string)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *SettingsGetConstantsResponse) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetEquipmentType

`func (o *SettingsGetConstantsResponse) GetEquipmentType() map[string]string`

GetEquipmentType returns the EquipmentType field if non-nil, zero value otherwise.

### GetEquipmentTypeOk

`func (o *SettingsGetConstantsResponse) GetEquipmentTypeOk() (*map[string]string, bool)`

GetEquipmentTypeOk returns a tuple with the EquipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentType

`func (o *SettingsGetConstantsResponse) SetEquipmentType(v map[string]string)`

SetEquipmentType sets EquipmentType field to given value.

### HasEquipmentType

`func (o *SettingsGetConstantsResponse) HasEquipmentType() bool`

HasEquipmentType returns a boolean if a field has been set.

### GetFileType

`func (o *SettingsGetConstantsResponse) GetFileType() map[string]string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *SettingsGetConstantsResponse) GetFileTypeOk() (*map[string]string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *SettingsGetConstantsResponse) SetFileType(v map[string]string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *SettingsGetConstantsResponse) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetIncoterm

`func (o *SettingsGetConstantsResponse) GetIncoterm() map[string]string`

GetIncoterm returns the Incoterm field if non-nil, zero value otherwise.

### GetIncotermOk

`func (o *SettingsGetConstantsResponse) GetIncotermOk() (*map[string]string, bool)`

GetIncotermOk returns a tuple with the Incoterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoterm

`func (o *SettingsGetConstantsResponse) SetIncoterm(v map[string]string)`

SetIncoterm sets Incoterm field to given value.

### HasIncoterm

`func (o *SettingsGetConstantsResponse) HasIncoterm() bool`

HasIncoterm returns a boolean if a field has been set.

### GetLogType

`func (o *SettingsGetConstantsResponse) GetLogType() map[string]string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *SettingsGetConstantsResponse) GetLogTypeOk() (*map[string]string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *SettingsGetConstantsResponse) SetLogType(v map[string]string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *SettingsGetConstantsResponse) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetObjectType

`func (o *SettingsGetConstantsResponse) GetObjectType() map[string]string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SettingsGetConstantsResponse) GetObjectTypeOk() (*map[string]string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SettingsGetConstantsResponse) SetObjectType(v map[string]string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *SettingsGetConstantsResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOptimisationObjective

`func (o *SettingsGetConstantsResponse) GetOptimisationObjective() map[string]string`

GetOptimisationObjective returns the OptimisationObjective field if non-nil, zero value otherwise.

### GetOptimisationObjectiveOk

`func (o *SettingsGetConstantsResponse) GetOptimisationObjectiveOk() (*map[string]string, bool)`

GetOptimisationObjectiveOk returns a tuple with the OptimisationObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimisationObjective

`func (o *SettingsGetConstantsResponse) SetOptimisationObjective(v map[string]string)`

SetOptimisationObjective sets OptimisationObjective field to given value.

### HasOptimisationObjective

`func (o *SettingsGetConstantsResponse) HasOptimisationObjective() bool`

HasOptimisationObjective returns a boolean if a field has been set.

### GetPartyType

`func (o *SettingsGetConstantsResponse) GetPartyType() map[string]string`

GetPartyType returns the PartyType field if non-nil, zero value otherwise.

### GetPartyTypeOk

`func (o *SettingsGetConstantsResponse) GetPartyTypeOk() (*map[string]string, bool)`

GetPartyTypeOk returns a tuple with the PartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyType

`func (o *SettingsGetConstantsResponse) SetPartyType(v map[string]string)`

SetPartyType sets PartyType field to given value.

### HasPartyType

`func (o *SettingsGetConstantsResponse) HasPartyType() bool`

HasPartyType returns a boolean if a field has been set.

### GetPhoneNrType

`func (o *SettingsGetConstantsResponse) GetPhoneNrType() map[string]string`

GetPhoneNrType returns the PhoneNrType field if non-nil, zero value otherwise.

### GetPhoneNrTypeOk

`func (o *SettingsGetConstantsResponse) GetPhoneNrTypeOk() (*map[string]string, bool)`

GetPhoneNrTypeOk returns a tuple with the PhoneNrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNrType

`func (o *SettingsGetConstantsResponse) SetPhoneNrType(v map[string]string)`

SetPhoneNrType sets PhoneNrType field to given value.

### HasPhoneNrType

`func (o *SettingsGetConstantsResponse) HasPhoneNrType() bool`

HasPhoneNrType returns a boolean if a field has been set.

### GetRecurrencePeriod

`func (o *SettingsGetConstantsResponse) GetRecurrencePeriod() map[string]string`

GetRecurrencePeriod returns the RecurrencePeriod field if non-nil, zero value otherwise.

### GetRecurrencePeriodOk

`func (o *SettingsGetConstantsResponse) GetRecurrencePeriodOk() (*map[string]string, bool)`

GetRecurrencePeriodOk returns a tuple with the RecurrencePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrencePeriod

`func (o *SettingsGetConstantsResponse) SetRecurrencePeriod(v map[string]string)`

SetRecurrencePeriod sets RecurrencePeriod field to given value.

### HasRecurrencePeriod

`func (o *SettingsGetConstantsResponse) HasRecurrencePeriod() bool`

HasRecurrencePeriod returns a boolean if a field has been set.

### GetRecurrenceType

`func (o *SettingsGetConstantsResponse) GetRecurrenceType() map[string]string`

GetRecurrenceType returns the RecurrenceType field if non-nil, zero value otherwise.

### GetRecurrenceTypeOk

`func (o *SettingsGetConstantsResponse) GetRecurrenceTypeOk() (*map[string]string, bool)`

GetRecurrenceTypeOk returns a tuple with the RecurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceType

`func (o *SettingsGetConstantsResponse) SetRecurrenceType(v map[string]string)`

SetRecurrenceType sets RecurrenceType field to given value.

### HasRecurrenceType

`func (o *SettingsGetConstantsResponse) HasRecurrenceType() bool`

HasRecurrenceType returns a boolean if a field has been set.

### GetRole

`func (o *SettingsGetConstantsResponse) GetRole() map[string]string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SettingsGetConstantsResponse) GetRoleOk() (*map[string]string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SettingsGetConstantsResponse) SetRole(v map[string]string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SettingsGetConstantsResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSaywhenStatus

`func (o *SettingsGetConstantsResponse) GetSaywhenStatus() map[string]string`

GetSaywhenStatus returns the SaywhenStatus field if non-nil, zero value otherwise.

### GetSaywhenStatusOk

`func (o *SettingsGetConstantsResponse) GetSaywhenStatusOk() (*map[string]string, bool)`

GetSaywhenStatusOk returns a tuple with the SaywhenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaywhenStatus

`func (o *SettingsGetConstantsResponse) SetSaywhenStatus(v map[string]string)`

SetSaywhenStatus sets SaywhenStatus field to given value.

### HasSaywhenStatus

`func (o *SettingsGetConstantsResponse) HasSaywhenStatus() bool`

HasSaywhenStatus returns a boolean if a field has been set.

### GetSettings

`func (o *SettingsGetConstantsResponse) GetSettings() map[string]string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SettingsGetConstantsResponse) GetSettingsOk() (*map[string]string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SettingsGetConstantsResponse) SetSettings(v map[string]string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SettingsGetConstantsResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSettingsGroup

`func (o *SettingsGetConstantsResponse) GetSettingsGroup() map[string]string`

GetSettingsGroup returns the SettingsGroup field if non-nil, zero value otherwise.

### GetSettingsGroupOk

`func (o *SettingsGetConstantsResponse) GetSettingsGroupOk() (*map[string]string, bool)`

GetSettingsGroupOk returns a tuple with the SettingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsGroup

`func (o *SettingsGetConstantsResponse) SetSettingsGroup(v map[string]string)`

SetSettingsGroup sets SettingsGroup field to given value.

### HasSettingsGroup

`func (o *SettingsGetConstantsResponse) HasSettingsGroup() bool`

HasSettingsGroup returns a boolean if a field has been set.

### GetStatusType

`func (o *SettingsGetConstantsResponse) GetStatusType() map[string]string`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *SettingsGetConstantsResponse) GetStatusTypeOk() (*map[string]string, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *SettingsGetConstantsResponse) SetStatusType(v map[string]string)`

SetStatusType sets StatusType field to given value.

### HasStatusType

`func (o *SettingsGetConstantsResponse) HasStatusType() bool`

HasStatusType returns a boolean if a field has been set.

### GetRouteStatus

`func (o *SettingsGetConstantsResponse) GetRouteStatus() map[string]string`

GetRouteStatus returns the RouteStatus field if non-nil, zero value otherwise.

### GetRouteStatusOk

`func (o *SettingsGetConstantsResponse) GetRouteStatusOk() (*map[string]string, bool)`

GetRouteStatusOk returns a tuple with the RouteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteStatus

`func (o *SettingsGetConstantsResponse) SetRouteStatus(v map[string]string)`

SetRouteStatus sets RouteStatus field to given value.

### HasRouteStatus

`func (o *SettingsGetConstantsResponse) HasRouteStatus() bool`

HasRouteStatus returns a boolean if a field has been set.

### GetActivityStatus

`func (o *SettingsGetConstantsResponse) GetActivityStatus() map[string]string`

GetActivityStatus returns the ActivityStatus field if non-nil, zero value otherwise.

### GetActivityStatusOk

`func (o *SettingsGetConstantsResponse) GetActivityStatusOk() (*map[string]string, bool)`

GetActivityStatusOk returns a tuple with the ActivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatus

`func (o *SettingsGetConstantsResponse) SetActivityStatus(v map[string]string)`

SetActivityStatus sets ActivityStatus field to given value.

### HasActivityStatus

`func (o *SettingsGetConstantsResponse) HasActivityStatus() bool`

HasActivityStatus returns a boolean if a field has been set.

### GetShipmentStatus

`func (o *SettingsGetConstantsResponse) GetShipmentStatus() map[string]string`

GetShipmentStatus returns the ShipmentStatus field if non-nil, zero value otherwise.

### GetShipmentStatusOk

`func (o *SettingsGetConstantsResponse) GetShipmentStatusOk() (*map[string]string, bool)`

GetShipmentStatusOk returns a tuple with the ShipmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentStatus

`func (o *SettingsGetConstantsResponse) SetShipmentStatus(v map[string]string)`

SetShipmentStatus sets ShipmentStatus field to given value.

### HasShipmentStatus

`func (o *SettingsGetConstantsResponse) HasShipmentStatus() bool`

HasShipmentStatus returns a boolean if a field has been set.

### GetPackageLineStatus

`func (o *SettingsGetConstantsResponse) GetPackageLineStatus() map[string]string`

GetPackageLineStatus returns the PackageLineStatus field if non-nil, zero value otherwise.

### GetPackageLineStatusOk

`func (o *SettingsGetConstantsResponse) GetPackageLineStatusOk() (*map[string]string, bool)`

GetPackageLineStatusOk returns a tuple with the PackageLineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLineStatus

`func (o *SettingsGetConstantsResponse) SetPackageLineStatus(v map[string]string)`

SetPackageLineStatus sets PackageLineStatus field to given value.

### HasPackageLineStatus

`func (o *SettingsGetConstantsResponse) HasPackageLineStatus() bool`

HasPackageLineStatus returns a boolean if a field has been set.

### GetAssignmentStatus

`func (o *SettingsGetConstantsResponse) GetAssignmentStatus() map[string]string`

GetAssignmentStatus returns the AssignmentStatus field if non-nil, zero value otherwise.

### GetAssignmentStatusOk

`func (o *SettingsGetConstantsResponse) GetAssignmentStatusOk() (*map[string]string, bool)`

GetAssignmentStatusOk returns a tuple with the AssignmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentStatus

`func (o *SettingsGetConstantsResponse) SetAssignmentStatus(v map[string]string)`

SetAssignmentStatus sets AssignmentStatus field to given value.

### HasAssignmentStatus

`func (o *SettingsGetConstantsResponse) HasAssignmentStatus() bool`

HasAssignmentStatus returns a boolean if a field has been set.

### GetMessageStatus

`func (o *SettingsGetConstantsResponse) GetMessageStatus() map[string]string`

GetMessageStatus returns the MessageStatus field if non-nil, zero value otherwise.

### GetMessageStatusOk

`func (o *SettingsGetConstantsResponse) GetMessageStatusOk() (*map[string]string, bool)`

GetMessageStatusOk returns a tuple with the MessageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageStatus

`func (o *SettingsGetConstantsResponse) SetMessageStatus(v map[string]string)`

SetMessageStatus sets MessageStatus field to given value.

### HasMessageStatus

`func (o *SettingsGetConstantsResponse) HasMessageStatus() bool`

HasMessageStatus returns a boolean if a field has been set.

### GetUserStatus

`func (o *SettingsGetConstantsResponse) GetUserStatus() map[string]string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *SettingsGetConstantsResponse) GetUserStatusOk() (*map[string]string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *SettingsGetConstantsResponse) SetUserStatus(v map[string]string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *SettingsGetConstantsResponse) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetJobStatus

`func (o *SettingsGetConstantsResponse) GetJobStatus() map[string]string`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *SettingsGetConstantsResponse) GetJobStatusOk() (*map[string]string, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *SettingsGetConstantsResponse) SetJobStatus(v map[string]string)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *SettingsGetConstantsResponse) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetApiRequestStatus

`func (o *SettingsGetConstantsResponse) GetApiRequestStatus() map[string]string`

GetApiRequestStatus returns the ApiRequestStatus field if non-nil, zero value otherwise.

### GetApiRequestStatusOk

`func (o *SettingsGetConstantsResponse) GetApiRequestStatusOk() (*map[string]string, bool)`

GetApiRequestStatusOk returns a tuple with the ApiRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRequestStatus

`func (o *SettingsGetConstantsResponse) SetApiRequestStatus(v map[string]string)`

SetApiRequestStatus sets ApiRequestStatus field to given value.

### HasApiRequestStatus

`func (o *SettingsGetConstantsResponse) HasApiRequestStatus() bool`

HasApiRequestStatus returns a boolean if a field has been set.

### GetSystemProvider

`func (o *SettingsGetConstantsResponse) GetSystemProvider() map[string]string`

GetSystemProvider returns the SystemProvider field if non-nil, zero value otherwise.

### GetSystemProviderOk

`func (o *SettingsGetConstantsResponse) GetSystemProviderOk() (*map[string]string, bool)`

GetSystemProviderOk returns a tuple with the SystemProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemProvider

`func (o *SettingsGetConstantsResponse) SetSystemProvider(v map[string]string)`

SetSystemProvider sets SystemProvider field to given value.

### HasSystemProvider

`func (o *SettingsGetConstantsResponse) HasSystemProvider() bool`

HasSystemProvider returns a boolean if a field has been set.

### GetProvider

`func (o *SettingsGetConstantsResponse) GetProvider() map[string]string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SettingsGetConstantsResponse) GetProviderOk() (*map[string]string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SettingsGetConstantsResponse) SetProvider(v map[string]string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SettingsGetConstantsResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetTimeSlotType

`func (o *SettingsGetConstantsResponse) GetTimeSlotType() map[string]string`

GetTimeSlotType returns the TimeSlotType field if non-nil, zero value otherwise.

### GetTimeSlotTypeOk

`func (o *SettingsGetConstantsResponse) GetTimeSlotTypeOk() (*map[string]string, bool)`

GetTimeSlotTypeOk returns a tuple with the TimeSlotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotType

`func (o *SettingsGetConstantsResponse) SetTimeSlotType(v map[string]string)`

SetTimeSlotType sets TimeSlotType field to given value.

### HasTimeSlotType

`func (o *SettingsGetConstantsResponse) HasTimeSlotType() bool`

HasTimeSlotType returns a boolean if a field has been set.

### GetTransactionType

`func (o *SettingsGetConstantsResponse) GetTransactionType() map[string]string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *SettingsGetConstantsResponse) GetTransactionTypeOk() (*map[string]string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *SettingsGetConstantsResponse) SetTransactionType(v map[string]string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *SettingsGetConstantsResponse) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetUom

`func (o *SettingsGetConstantsResponse) GetUom() map[string]string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *SettingsGetConstantsResponse) GetUomOk() (*map[string]string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *SettingsGetConstantsResponse) SetUom(v map[string]string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *SettingsGetConstantsResponse) HasUom() bool`

HasUom returns a boolean if a field has been set.

### GetUomGroup

`func (o *SettingsGetConstantsResponse) GetUomGroup() map[string]string`

GetUomGroup returns the UomGroup field if non-nil, zero value otherwise.

### GetUomGroupOk

`func (o *SettingsGetConstantsResponse) GetUomGroupOk() (*map[string]string, bool)`

GetUomGroupOk returns a tuple with the UomGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomGroup

`func (o *SettingsGetConstantsResponse) SetUomGroup(v map[string]string)`

SetUomGroup sets UomGroup field to given value.

### HasUomGroup

`func (o *SettingsGetConstantsResponse) HasUomGroup() bool`

HasUomGroup returns a boolean if a field has been set.

### GetVehicleType

`func (o *SettingsGetConstantsResponse) GetVehicleType() map[string]string`

GetVehicleType returns the VehicleType field if non-nil, zero value otherwise.

### GetVehicleTypeOk

`func (o *SettingsGetConstantsResponse) GetVehicleTypeOk() (*map[string]string, bool)`

GetVehicleTypeOk returns a tuple with the VehicleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleType

`func (o *SettingsGetConstantsResponse) SetVehicleType(v map[string]string)`

SetVehicleType sets VehicleType field to given value.

### HasVehicleType

`func (o *SettingsGetConstantsResponse) HasVehicleType() bool`

HasVehicleType returns a boolean if a field has been set.

### GetPackageType

`func (o *SettingsGetConstantsResponse) GetPackageType() map[string]string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *SettingsGetConstantsResponse) GetPackageTypeOk() (*map[string]string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *SettingsGetConstantsResponse) SetPackageType(v map[string]string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *SettingsGetConstantsResponse) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


