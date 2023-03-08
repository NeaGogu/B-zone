# \CommunicationApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveActivityCommunication**](CommunicationApi.md#RetrieveActivityCommunication) | **Post** /communication/retrieve-activity | Retrieve Activity
[**TriggerMessageCommunication**](CommunicationApi.md#TriggerMessageCommunication) | **Post** /communication/trigger-message | Trigger Message to Communication



## RetrieveActivityCommunication

> ApiResponse RetrieveActivityCommunication(ctx).ActivityId(activityId).Execute()

Retrieve Activity



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
    activityId := int32(56) // int32 | ActivityId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationApi.RetrieveActivityCommunication(context.Background()).ActivityId(activityId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationApi.RetrieveActivityCommunication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveActivityCommunication`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationApi.RetrieveActivityCommunication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveActivityCommunicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activityId** | **int32** | ActivityId | 

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


## TriggerMessageCommunication

> ApiResponse TriggerMessageCommunication(ctx).Body(body).Execute()

Trigger Message to Communication



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
    resp, r, err := api_client.CommunicationApi.TriggerMessageCommunication(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationApi.TriggerMessageCommunication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerMessageCommunication`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationApi.TriggerMessageCommunication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerMessageCommunicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommunicationTriggerMessageModel**](CommunicationTriggerMessageModel.md) | communication trigger message object | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

