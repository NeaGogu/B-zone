# CommunicationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Saywhen** | Pointer to **bool** | Whether or not activity should be synced with Saywhen | [optional] 
**Bumbal** | Pointer to **bool** | Whether or not activity is handled by Bumbal Communication Server | [optional] 
**SendInvite** | Pointer to **bool** | Send invite | [optional] 
**SendReminder** | Pointer to **bool** | Send reminder | [optional] 
**SendPrefConfirmation** | Pointer to **bool** | Send confirmation on preferences received | [optional] 
**SendPlanned** | Pointer to **bool** | Send planned notice | [optional] 
**SendEta** | Pointer to **bool** | Send eta notice | [optional] 
**SendExecuted** | Pointer to **bool** | Send executed notice | [optional] 
**SendCancelled** | Pointer to **bool** | Send cancelled notice | [optional] 
**Email** | Pointer to **string** | Email for customer communication | [optional] 
**PhoneNr** | Pointer to **string** | Phone nr for customer communication | [optional] 
**SaywhenStatusId** | Pointer to **int32** | Saywhen StatusId of this Activity, 1:cancelled, 2:offered, 3:preffed, 4:confirmed, 5:accepted, 6:planned, 7:scheduled, 8:started, 9:completed | [optional] 
**SaywhenStatusName** | Pointer to **string** | Saywhen Status name | [optional] 

## Methods

### NewCommunicationModel

`func NewCommunicationModel() *CommunicationModel`

NewCommunicationModel instantiates a new CommunicationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationModelWithDefaults

`func NewCommunicationModelWithDefaults() *CommunicationModel`

NewCommunicationModelWithDefaults instantiates a new CommunicationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommunicationModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunicationModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunicationModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CommunicationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSaywhen

`func (o *CommunicationModel) GetSaywhen() bool`

GetSaywhen returns the Saywhen field if non-nil, zero value otherwise.

### GetSaywhenOk

`func (o *CommunicationModel) GetSaywhenOk() (*bool, bool)`

GetSaywhenOk returns a tuple with the Saywhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaywhen

`func (o *CommunicationModel) SetSaywhen(v bool)`

SetSaywhen sets Saywhen field to given value.

### HasSaywhen

`func (o *CommunicationModel) HasSaywhen() bool`

HasSaywhen returns a boolean if a field has been set.

### GetBumbal

`func (o *CommunicationModel) GetBumbal() bool`

GetBumbal returns the Bumbal field if non-nil, zero value otherwise.

### GetBumbalOk

`func (o *CommunicationModel) GetBumbalOk() (*bool, bool)`

GetBumbalOk returns a tuple with the Bumbal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBumbal

`func (o *CommunicationModel) SetBumbal(v bool)`

SetBumbal sets Bumbal field to given value.

### HasBumbal

`func (o *CommunicationModel) HasBumbal() bool`

HasBumbal returns a boolean if a field has been set.

### GetSendInvite

`func (o *CommunicationModel) GetSendInvite() bool`

GetSendInvite returns the SendInvite field if non-nil, zero value otherwise.

### GetSendInviteOk

`func (o *CommunicationModel) GetSendInviteOk() (*bool, bool)`

GetSendInviteOk returns a tuple with the SendInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvite

`func (o *CommunicationModel) SetSendInvite(v bool)`

SetSendInvite sets SendInvite field to given value.

### HasSendInvite

`func (o *CommunicationModel) HasSendInvite() bool`

HasSendInvite returns a boolean if a field has been set.

### GetSendReminder

`func (o *CommunicationModel) GetSendReminder() bool`

GetSendReminder returns the SendReminder field if non-nil, zero value otherwise.

### GetSendReminderOk

`func (o *CommunicationModel) GetSendReminderOk() (*bool, bool)`

GetSendReminderOk returns a tuple with the SendReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReminder

`func (o *CommunicationModel) SetSendReminder(v bool)`

SetSendReminder sets SendReminder field to given value.

### HasSendReminder

`func (o *CommunicationModel) HasSendReminder() bool`

HasSendReminder returns a boolean if a field has been set.

### GetSendPrefConfirmation

`func (o *CommunicationModel) GetSendPrefConfirmation() bool`

GetSendPrefConfirmation returns the SendPrefConfirmation field if non-nil, zero value otherwise.

### GetSendPrefConfirmationOk

`func (o *CommunicationModel) GetSendPrefConfirmationOk() (*bool, bool)`

GetSendPrefConfirmationOk returns a tuple with the SendPrefConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPrefConfirmation

`func (o *CommunicationModel) SetSendPrefConfirmation(v bool)`

SetSendPrefConfirmation sets SendPrefConfirmation field to given value.

### HasSendPrefConfirmation

`func (o *CommunicationModel) HasSendPrefConfirmation() bool`

HasSendPrefConfirmation returns a boolean if a field has been set.

### GetSendPlanned

`func (o *CommunicationModel) GetSendPlanned() bool`

GetSendPlanned returns the SendPlanned field if non-nil, zero value otherwise.

### GetSendPlannedOk

`func (o *CommunicationModel) GetSendPlannedOk() (*bool, bool)`

GetSendPlannedOk returns a tuple with the SendPlanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPlanned

`func (o *CommunicationModel) SetSendPlanned(v bool)`

SetSendPlanned sets SendPlanned field to given value.

### HasSendPlanned

`func (o *CommunicationModel) HasSendPlanned() bool`

HasSendPlanned returns a boolean if a field has been set.

### GetSendEta

`func (o *CommunicationModel) GetSendEta() bool`

GetSendEta returns the SendEta field if non-nil, zero value otherwise.

### GetSendEtaOk

`func (o *CommunicationModel) GetSendEtaOk() (*bool, bool)`

GetSendEtaOk returns a tuple with the SendEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEta

`func (o *CommunicationModel) SetSendEta(v bool)`

SetSendEta sets SendEta field to given value.

### HasSendEta

`func (o *CommunicationModel) HasSendEta() bool`

HasSendEta returns a boolean if a field has been set.

### GetSendExecuted

`func (o *CommunicationModel) GetSendExecuted() bool`

GetSendExecuted returns the SendExecuted field if non-nil, zero value otherwise.

### GetSendExecutedOk

`func (o *CommunicationModel) GetSendExecutedOk() (*bool, bool)`

GetSendExecutedOk returns a tuple with the SendExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendExecuted

`func (o *CommunicationModel) SetSendExecuted(v bool)`

SetSendExecuted sets SendExecuted field to given value.

### HasSendExecuted

`func (o *CommunicationModel) HasSendExecuted() bool`

HasSendExecuted returns a boolean if a field has been set.

### GetSendCancelled

`func (o *CommunicationModel) GetSendCancelled() bool`

GetSendCancelled returns the SendCancelled field if non-nil, zero value otherwise.

### GetSendCancelledOk

`func (o *CommunicationModel) GetSendCancelledOk() (*bool, bool)`

GetSendCancelledOk returns a tuple with the SendCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCancelled

`func (o *CommunicationModel) SetSendCancelled(v bool)`

SetSendCancelled sets SendCancelled field to given value.

### HasSendCancelled

`func (o *CommunicationModel) HasSendCancelled() bool`

HasSendCancelled returns a boolean if a field has been set.

### GetEmail

`func (o *CommunicationModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CommunicationModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CommunicationModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CommunicationModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNr

`func (o *CommunicationModel) GetPhoneNr() string`

GetPhoneNr returns the PhoneNr field if non-nil, zero value otherwise.

### GetPhoneNrOk

`func (o *CommunicationModel) GetPhoneNrOk() (*string, bool)`

GetPhoneNrOk returns a tuple with the PhoneNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNr

`func (o *CommunicationModel) SetPhoneNr(v string)`

SetPhoneNr sets PhoneNr field to given value.

### HasPhoneNr

`func (o *CommunicationModel) HasPhoneNr() bool`

HasPhoneNr returns a boolean if a field has been set.

### GetSaywhenStatusId

`func (o *CommunicationModel) GetSaywhenStatusId() int32`

GetSaywhenStatusId returns the SaywhenStatusId field if non-nil, zero value otherwise.

### GetSaywhenStatusIdOk

`func (o *CommunicationModel) GetSaywhenStatusIdOk() (*int32, bool)`

GetSaywhenStatusIdOk returns a tuple with the SaywhenStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaywhenStatusId

`func (o *CommunicationModel) SetSaywhenStatusId(v int32)`

SetSaywhenStatusId sets SaywhenStatusId field to given value.

### HasSaywhenStatusId

`func (o *CommunicationModel) HasSaywhenStatusId() bool`

HasSaywhenStatusId returns a boolean if a field has been set.

### GetSaywhenStatusName

`func (o *CommunicationModel) GetSaywhenStatusName() string`

GetSaywhenStatusName returns the SaywhenStatusName field if non-nil, zero value otherwise.

### GetSaywhenStatusNameOk

`func (o *CommunicationModel) GetSaywhenStatusNameOk() (*string, bool)`

GetSaywhenStatusNameOk returns a tuple with the SaywhenStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaywhenStatusName

`func (o *CommunicationModel) SetSaywhenStatusName(v string)`

SetSaywhenStatusName sets SaywhenStatusName field to given value.

### HasSaywhenStatusName

`func (o *CommunicationModel) HasSaywhenStatusName() bool`

HasSaywhenStatusName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


