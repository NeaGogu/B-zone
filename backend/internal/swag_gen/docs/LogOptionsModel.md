# LogOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeMessage** | Pointer to **bool** | Default &#x3D; true | [optional] 
**IncludeTimestamp** | Pointer to **bool** | Default &#x3D; true | [optional] 
**IncludeFullMessage** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeOriginalMessage** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeClientIp** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeHttpMethod** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeApiEndpoint** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeLevelText** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeRequestId** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeChannel** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeLogType** | Pointer to **bool** | Default &#x3D; false | [optional] 
**IncludeUser** | Pointer to **bool** | Default &#x3D; false | [optional] 

## Methods

### NewLogOptionsModel

`func NewLogOptionsModel() *LogOptionsModel`

NewLogOptionsModel instantiates a new LogOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogOptionsModelWithDefaults

`func NewLogOptionsModelWithDefaults() *LogOptionsModel`

NewLogOptionsModelWithDefaults instantiates a new LogOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeMessage

`func (o *LogOptionsModel) GetIncludeMessage() bool`

GetIncludeMessage returns the IncludeMessage field if non-nil, zero value otherwise.

### GetIncludeMessageOk

`func (o *LogOptionsModel) GetIncludeMessageOk() (*bool, bool)`

GetIncludeMessageOk returns a tuple with the IncludeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMessage

`func (o *LogOptionsModel) SetIncludeMessage(v bool)`

SetIncludeMessage sets IncludeMessage field to given value.

### HasIncludeMessage

`func (o *LogOptionsModel) HasIncludeMessage() bool`

HasIncludeMessage returns a boolean if a field has been set.

### GetIncludeTimestamp

`func (o *LogOptionsModel) GetIncludeTimestamp() bool`

GetIncludeTimestamp returns the IncludeTimestamp field if non-nil, zero value otherwise.

### GetIncludeTimestampOk

`func (o *LogOptionsModel) GetIncludeTimestampOk() (*bool, bool)`

GetIncludeTimestampOk returns a tuple with the IncludeTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTimestamp

`func (o *LogOptionsModel) SetIncludeTimestamp(v bool)`

SetIncludeTimestamp sets IncludeTimestamp field to given value.

### HasIncludeTimestamp

`func (o *LogOptionsModel) HasIncludeTimestamp() bool`

HasIncludeTimestamp returns a boolean if a field has been set.

### GetIncludeFullMessage

`func (o *LogOptionsModel) GetIncludeFullMessage() bool`

GetIncludeFullMessage returns the IncludeFullMessage field if non-nil, zero value otherwise.

### GetIncludeFullMessageOk

`func (o *LogOptionsModel) GetIncludeFullMessageOk() (*bool, bool)`

GetIncludeFullMessageOk returns a tuple with the IncludeFullMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFullMessage

`func (o *LogOptionsModel) SetIncludeFullMessage(v bool)`

SetIncludeFullMessage sets IncludeFullMessage field to given value.

### HasIncludeFullMessage

`func (o *LogOptionsModel) HasIncludeFullMessage() bool`

HasIncludeFullMessage returns a boolean if a field has been set.

### GetIncludeOriginalMessage

`func (o *LogOptionsModel) GetIncludeOriginalMessage() bool`

GetIncludeOriginalMessage returns the IncludeOriginalMessage field if non-nil, zero value otherwise.

### GetIncludeOriginalMessageOk

`func (o *LogOptionsModel) GetIncludeOriginalMessageOk() (*bool, bool)`

GetIncludeOriginalMessageOk returns a tuple with the IncludeOriginalMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOriginalMessage

`func (o *LogOptionsModel) SetIncludeOriginalMessage(v bool)`

SetIncludeOriginalMessage sets IncludeOriginalMessage field to given value.

### HasIncludeOriginalMessage

`func (o *LogOptionsModel) HasIncludeOriginalMessage() bool`

HasIncludeOriginalMessage returns a boolean if a field has been set.

### GetIncludeClientIp

`func (o *LogOptionsModel) GetIncludeClientIp() bool`

GetIncludeClientIp returns the IncludeClientIp field if non-nil, zero value otherwise.

### GetIncludeClientIpOk

`func (o *LogOptionsModel) GetIncludeClientIpOk() (*bool, bool)`

GetIncludeClientIpOk returns a tuple with the IncludeClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeClientIp

`func (o *LogOptionsModel) SetIncludeClientIp(v bool)`

SetIncludeClientIp sets IncludeClientIp field to given value.

### HasIncludeClientIp

`func (o *LogOptionsModel) HasIncludeClientIp() bool`

HasIncludeClientIp returns a boolean if a field has been set.

### GetIncludeHttpMethod

`func (o *LogOptionsModel) GetIncludeHttpMethod() bool`

GetIncludeHttpMethod returns the IncludeHttpMethod field if non-nil, zero value otherwise.

### GetIncludeHttpMethodOk

`func (o *LogOptionsModel) GetIncludeHttpMethodOk() (*bool, bool)`

GetIncludeHttpMethodOk returns a tuple with the IncludeHttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHttpMethod

`func (o *LogOptionsModel) SetIncludeHttpMethod(v bool)`

SetIncludeHttpMethod sets IncludeHttpMethod field to given value.

### HasIncludeHttpMethod

`func (o *LogOptionsModel) HasIncludeHttpMethod() bool`

HasIncludeHttpMethod returns a boolean if a field has been set.

### GetIncludeApiEndpoint

`func (o *LogOptionsModel) GetIncludeApiEndpoint() bool`

GetIncludeApiEndpoint returns the IncludeApiEndpoint field if non-nil, zero value otherwise.

### GetIncludeApiEndpointOk

`func (o *LogOptionsModel) GetIncludeApiEndpointOk() (*bool, bool)`

GetIncludeApiEndpointOk returns a tuple with the IncludeApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeApiEndpoint

`func (o *LogOptionsModel) SetIncludeApiEndpoint(v bool)`

SetIncludeApiEndpoint sets IncludeApiEndpoint field to given value.

### HasIncludeApiEndpoint

`func (o *LogOptionsModel) HasIncludeApiEndpoint() bool`

HasIncludeApiEndpoint returns a boolean if a field has been set.

### GetIncludeLevelText

`func (o *LogOptionsModel) GetIncludeLevelText() bool`

GetIncludeLevelText returns the IncludeLevelText field if non-nil, zero value otherwise.

### GetIncludeLevelTextOk

`func (o *LogOptionsModel) GetIncludeLevelTextOk() (*bool, bool)`

GetIncludeLevelTextOk returns a tuple with the IncludeLevelText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLevelText

`func (o *LogOptionsModel) SetIncludeLevelText(v bool)`

SetIncludeLevelText sets IncludeLevelText field to given value.

### HasIncludeLevelText

`func (o *LogOptionsModel) HasIncludeLevelText() bool`

HasIncludeLevelText returns a boolean if a field has been set.

### GetIncludeRequestId

`func (o *LogOptionsModel) GetIncludeRequestId() bool`

GetIncludeRequestId returns the IncludeRequestId field if non-nil, zero value otherwise.

### GetIncludeRequestIdOk

`func (o *LogOptionsModel) GetIncludeRequestIdOk() (*bool, bool)`

GetIncludeRequestIdOk returns a tuple with the IncludeRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRequestId

`func (o *LogOptionsModel) SetIncludeRequestId(v bool)`

SetIncludeRequestId sets IncludeRequestId field to given value.

### HasIncludeRequestId

`func (o *LogOptionsModel) HasIncludeRequestId() bool`

HasIncludeRequestId returns a boolean if a field has been set.

### GetIncludeChannel

`func (o *LogOptionsModel) GetIncludeChannel() bool`

GetIncludeChannel returns the IncludeChannel field if non-nil, zero value otherwise.

### GetIncludeChannelOk

`func (o *LogOptionsModel) GetIncludeChannelOk() (*bool, bool)`

GetIncludeChannelOk returns a tuple with the IncludeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChannel

`func (o *LogOptionsModel) SetIncludeChannel(v bool)`

SetIncludeChannel sets IncludeChannel field to given value.

### HasIncludeChannel

`func (o *LogOptionsModel) HasIncludeChannel() bool`

HasIncludeChannel returns a boolean if a field has been set.

### GetIncludeLogType

`func (o *LogOptionsModel) GetIncludeLogType() bool`

GetIncludeLogType returns the IncludeLogType field if non-nil, zero value otherwise.

### GetIncludeLogTypeOk

`func (o *LogOptionsModel) GetIncludeLogTypeOk() (*bool, bool)`

GetIncludeLogTypeOk returns a tuple with the IncludeLogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLogType

`func (o *LogOptionsModel) SetIncludeLogType(v bool)`

SetIncludeLogType sets IncludeLogType field to given value.

### HasIncludeLogType

`func (o *LogOptionsModel) HasIncludeLogType() bool`

HasIncludeLogType returns a boolean if a field has been set.

### GetIncludeUser

`func (o *LogOptionsModel) GetIncludeUser() bool`

GetIncludeUser returns the IncludeUser field if non-nil, zero value otherwise.

### GetIncludeUserOk

`func (o *LogOptionsModel) GetIncludeUserOk() (*bool, bool)`

GetIncludeUserOk returns a tuple with the IncludeUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUser

`func (o *LogOptionsModel) SetIncludeUser(v bool)`

SetIncludeUser sets IncludeUser field to given value.

### HasIncludeUser

`func (o *LogOptionsModel) HasIncludeUser() bool`

HasIncludeUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


