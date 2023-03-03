# \CommunicationMessageTypeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommunicationMessageType**](CommunicationMessageTypeApi.md#CreateCommunicationMessageType) | **Post** /communication-message-type/set | Create a communicationmessagetype
[**DeleteCommunicationMessageType**](CommunicationMessageTypeApi.md#DeleteCommunicationMessageType) | **Delete** /communication-message-type/{communicationMessageTypeId} | Delete a communication message type



## CreateCommunicationMessageType

> ApiResponse CreateCommunicationMessageType(ctx).Body(body).Execute()

Create a communicationmessagetype



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
    body := *openapiclient.NewCommunicationMessageTypeModel() // CommunicationMessageTypeModel | CommunicationMessageType object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMessageTypeApi.CreateCommunicationMessageType(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMessageTypeApi.CreateCommunicationMessageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommunicationMessageType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMessageTypeApi.CreateCommunicationMessageType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommunicationMessageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommunicationMessageTypeModel**](CommunicationMessageTypeModel.md) | CommunicationMessageType object | 

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


## DeleteCommunicationMessageType

> ApiResponse DeleteCommunicationMessageType(ctx, communicationMessageTypeId).Execute()

Delete a communication message type



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
    communicationMessageTypeId := int64(789) // int64 | ID of the communication message type to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMessageTypeApi.DeleteCommunicationMessageType(context.Background(), communicationMessageTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMessageTypeApi.DeleteCommunicationMessageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCommunicationMessageType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMessageTypeApi.DeleteCommunicationMessageType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationMessageTypeId** | **int64** | ID of the communication message type to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommunicationMessageTypeRequest struct via the builder pattern


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

