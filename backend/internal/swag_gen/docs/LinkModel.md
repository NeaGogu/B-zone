# LinkModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**LinkId** | Pointer to **string** | The external Link ID | [optional] 
**ObjectType** | Pointer to **int64** | The object type id which this link is connected to | [optional] 
**ObjectId** | Pointer to **int64** | The object ID which this link is connected to | [optional] 
**ProviderId** | Pointer to **int64** | The Provider ID which this link is connected to | [optional] 
**ProviderReference** | Pointer to **string** | A Provider reference which this link is connected to | [optional] 
**ProviderName** | Pointer to **string** | The name of the external Provider | [optional] 

## Methods

### NewLinkModel

`func NewLinkModel() *LinkModel`

NewLinkModel instantiates a new LinkModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkModelWithDefaults

`func NewLinkModelWithDefaults() *LinkModel`

NewLinkModelWithDefaults instantiates a new LinkModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LinkModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinkId

`func (o *LinkModel) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *LinkModel) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *LinkModel) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *LinkModel) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetObjectType

`func (o *LinkModel) GetObjectType() int64`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LinkModel) GetObjectTypeOk() (*int64, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LinkModel) SetObjectType(v int64)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *LinkModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *LinkModel) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *LinkModel) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *LinkModel) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *LinkModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetProviderId

`func (o *LinkModel) GetProviderId() int64`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *LinkModel) GetProviderIdOk() (*int64, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *LinkModel) SetProviderId(v int64)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *LinkModel) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderReference

`func (o *LinkModel) GetProviderReference() string`

GetProviderReference returns the ProviderReference field if non-nil, zero value otherwise.

### GetProviderReferenceOk

`func (o *LinkModel) GetProviderReferenceOk() (*string, bool)`

GetProviderReferenceOk returns a tuple with the ProviderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderReference

`func (o *LinkModel) SetProviderReference(v string)`

SetProviderReference sets ProviderReference field to given value.

### HasProviderReference

`func (o *LinkModel) HasProviderReference() bool`

HasProviderReference returns a boolean if a field has been set.

### GetProviderName

`func (o *LinkModel) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *LinkModel) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *LinkModel) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *LinkModel) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


