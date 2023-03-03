# \PortalCommunicationApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveActivityByIdCommunication**](PortalCommunicationApi.md#RetrieveActivityByIdCommunication) | **Get** /portal-communication/retrieve-activity-by-id/{activityId} | Retrieve an activity by ID
[**TriggerMessageCommunicationByMessageType**](PortalCommunicationApi.md#TriggerMessageCommunicationByMessageType) | **Post** /portal-communication/trigger-message-by-message-type | Trigger Message to Communication by Message Type



## RetrieveActivityByIdCommunication

> ApiResponse RetrieveActivityByIdCommunication(ctx, activityId).Execute()

Retrieve an activity by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    activityId := int64(789) // int64 | ID of activity to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortalCommunicationApi.RetrieveActivityByIdCommunication(context.Background(), activityId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PortalCommunicationApi.RetrieveActivityByIdCommunication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveActivityByIdCommunication`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PortalCommunicationApi.RetrieveActivityByIdCommunication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityId** | **int64** | ID of activity to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveActivityByIdCommunicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerMessageCommunicationByMessageType

> ApiResponse TriggerMessageCommunicationByMessageType(ctx).Body(body).Execute()

Trigger Message to Communication by Message Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCommunicationTriggerMessageModel() // CommunicationTriggerMessageModel | communication trigger message object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortalCommunicationApi.TriggerMessageCommunicationByMessageType(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PortalCommunicationApi.TriggerMessageCommunicationByMessageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerMessageCommunicationByMessageType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PortalCommunicationApi.TriggerMessageCommunicationByMessageType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerMessageCommunicationByMessageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommunicationTriggerMessageModel**](CommunicationTriggerMessageModel.md) | communication trigger message object | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml, multipart/form-data
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

