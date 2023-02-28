# \CommunicationMessageApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommunicationMessage**](CommunicationMessageApi.md#CreateCommunicationMessage) | **Post** /communication-message/set | Create a communicationmessage
[**DeleteCommunicationMessage**](CommunicationMessageApi.md#DeleteCommunicationMessage) | **Delete** /communication-message/{communicationMessageId} | Delete a communication message
[**RetrieveCommunicationMessage**](CommunicationMessageApi.md#RetrieveCommunicationMessage) | **Get** /communication-message/{communicationMessageId} | Retrieve a communicationmessage
[**RetrieveListCommunicationMessages**](CommunicationMessageApi.md#RetrieveListCommunicationMessages) | **Put** /communication-message | Retrieve List of communication messages



## CreateCommunicationMessage

> ApiResponse CreateCommunicationMessage(ctx).Body(body).Execute()

Create a communicationmessage



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
    body := *openapiclient.NewCommunicationMessageModel() // CommunicationMessageModel | CommunicationMessage object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMessageApi.CreateCommunicationMessage(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMessageApi.CreateCommunicationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommunicationMessage`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMessageApi.CreateCommunicationMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommunicationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommunicationMessageModel**](CommunicationMessageModel.md) | CommunicationMessage object | 

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


## DeleteCommunicationMessage

> ApiResponse DeleteCommunicationMessage(ctx, communicationMessageId).Execute()

Delete a communication message



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
    communicationMessageId := int64(789) // int64 | ID of the communication message to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMessageApi.DeleteCommunicationMessage(context.Background(), communicationMessageId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMessageApi.DeleteCommunicationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCommunicationMessage`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMessageApi.DeleteCommunicationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationMessageId** | **int64** | ID of the communication message to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommunicationMessageRequest struct via the builder pattern


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


## RetrieveCommunicationMessage

> CommunicationMessageModel RetrieveCommunicationMessage(ctx, communicationMessageId).Execute()

Retrieve a communicationmessage



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
    communicationMessageId := int64(789) // int64 | ID of template to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMessageApi.RetrieveCommunicationMessage(context.Background(), communicationMessageId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMessageApi.RetrieveCommunicationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCommunicationMessage`: CommunicationMessageModel
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMessageApi.RetrieveCommunicationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationMessageId** | **int64** | ID of template to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCommunicationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommunicationMessageModel**](CommunicationMessageModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListCommunicationMessages

> CommunicationMessageListResponse RetrieveListCommunicationMessages(ctx).Arguments(arguments).Execute()

Retrieve List of communication messages



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
    arguments := *openapiclient.NewCommunicationMessageRetrieveListArguments() // CommunicationMessageRetrieveListArguments | Template RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMessageApi.RetrieveListCommunicationMessages(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMessageApi.RetrieveListCommunicationMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListCommunicationMessages`: CommunicationMessageListResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMessageApi.RetrieveListCommunicationMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListCommunicationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**CommunicationMessageRetrieveListArguments**](CommunicationMessageRetrieveListArguments.md) | Template RetrieveList Arguments | 

### Return type

[**CommunicationMessageListResponse**](CommunicationMessageListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

