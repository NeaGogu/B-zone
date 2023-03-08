# LogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalMessage** | Pointer to **string** | The log message. | [optional] 
**Message** | Pointer to **string** | The log message. Can be decorated with extra data | [optional] 
**FullMessage** | Pointer to **map[string]interface{}** | Any additional data you want to store about this log message | [optional] 
**Channel** | Pointer to **string** | Groups log messages | [optional] 
**LevelText** | Pointer to **string** | Log level | [optional] 
**HttpMethod** | Pointer to **string** | HTTP method | [optional] 
**ApiEndpoint** | Pointer to **string** | Path of API endpoint | [optional] 
**ClientIp** | Pointer to **string** | Client IP | [optional] 
**RequestId** | Pointer to **string** | Identifier to relate different log messages | [optional] 
**LogType** | Pointer to **string** | Type of log message | [optional] 
**Timestamp** | Pointer to **string** | Datetime of log message | [optional] [readonly] 
**User** | Pointer to **map[string]interface{}** | User details | [optional] 

## Methods

### NewLogModel

`func NewLogModel() *LogModel`

NewLogModel instantiates a new LogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogModelWithDefaults

`func NewLogModelWithDefaults() *LogModel`

NewLogModelWithDefaults instantiates a new LogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalMessage

`func (o *LogModel) GetOriginalMessage() string`

GetOriginalMessage returns the OriginalMessage field if non-nil, zero value otherwise.

### GetOriginalMessageOk

`func (o *LogModel) GetOriginalMessageOk() (*string, bool)`

GetOriginalMessageOk returns a tuple with the OriginalMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMessage

`func (o *LogModel) SetOriginalMessage(v string)`

SetOriginalMessage sets OriginalMessage field to given value.

### HasOriginalMessage

`func (o *LogModel) HasOriginalMessage() bool`

HasOriginalMessage returns a boolean if a field has been set.

### GetMessage

`func (o *LogModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFullMessage

`func (o *LogModel) GetFullMessage() map[string]interface{}`

GetFullMessage returns the FullMessage field if non-nil, zero value otherwise.

### GetFullMessageOk

`func (o *LogModel) GetFullMessageOk() (*map[string]interface{}, bool)`

GetFullMessageOk returns a tuple with the FullMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullMessage

`func (o *LogModel) SetFullMessage(v map[string]interface{})`

SetFullMessage sets FullMessage field to given value.

### HasFullMessage

`func (o *LogModel) HasFullMessage() bool`

HasFullMessage returns a boolean if a field has been set.

### GetChannel

`func (o *LogModel) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *LogModel) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *LogModel) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *LogModel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetLevelText

`func (o *LogModel) GetLevelText() string`

GetLevelText returns the LevelText field if non-nil, zero value otherwise.

### GetLevelTextOk

`func (o *LogModel) GetLevelTextOk() (*string, bool)`

GetLevelTextOk returns a tuple with the LevelText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelText

`func (o *LogModel) SetLevelText(v string)`

SetLevelText sets LevelText field to given value.

### HasLevelText

`func (o *LogModel) HasLevelText() bool`

HasLevelText returns a boolean if a field has been set.

### GetHttpMethod

`func (o *LogModel) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *LogModel) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *LogModel) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *LogModel) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetApiEndpoint

`func (o *LogModel) GetApiEndpoint() string`

GetApiEndpoint returns the ApiEndpoint field if non-nil, zero value otherwise.

### GetApiEndpointOk

`func (o *LogModel) GetApiEndpointOk() (*string, bool)`

GetApiEndpointOk returns a tuple with the ApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpoint

`func (o *LogModel) SetApiEndpoint(v string)`

SetApiEndpoint sets ApiEndpoint field to given value.

### HasApiEndpoint

`func (o *LogModel) HasApiEndpoint() bool`

HasApiEndpoint returns a boolean if a field has been set.

### GetClientIp

`func (o *LogModel) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *LogModel) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *LogModel) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *LogModel) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetRequestId

`func (o *LogModel) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *LogModel) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *LogModel) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *LogModel) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetLogType

`func (o *LogModel) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *LogModel) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *LogModel) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *LogModel) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetTimestamp

`func (o *LogModel) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogModel) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogModel) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogModel) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUser

`func (o *LogModel) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LogModel) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LogModel) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *LogModel) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


