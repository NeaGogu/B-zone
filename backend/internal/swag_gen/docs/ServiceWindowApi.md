# \ServiceWindowApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveListServiceWindow**](ServiceWindowApi.md#RetrieveListServiceWindow) | **Put** /service-window | Retrieve a list of service windows
[**RetrieveServiceWindow**](ServiceWindowApi.md#RetrieveServiceWindow) | **Get** /service-window/{serviceWindowId} | Retrieve a single service window
[**SetServiceWindow**](ServiceWindowApi.md#SetServiceWindow) | **Post** /service-window/set | Add or update a service window



## RetrieveListServiceWindow

> ServiceWindowListResponse RetrieveListServiceWindow(ctx).Arguments(arguments).Execute()

Retrieve a list of service windows



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
    arguments := *openapiclient.NewServiceWindowRetrieveListArguments() // ServiceWindowRetrieveListArguments | Service Window RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceWindowApi.RetrieveListServiceWindow(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceWindowApi.RetrieveListServiceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListServiceWindow`: ServiceWindowListResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceWindowApi.RetrieveListServiceWindow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListServiceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**ServiceWindowRetrieveListArguments**](ServiceWindowRetrieveListArguments.md) | Service Window RetrieveList Arguments | 

### Return type

[**ServiceWindowListResponse**](ServiceWindowListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveServiceWindow

> ServiceWindowModel RetrieveServiceWindow(ctx, serviceWindowId).Execute()

Retrieve a single service window



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
    serviceWindowId := int64(789) // int64 | ID of Service windows scheme to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceWindowApi.RetrieveServiceWindow(context.Background(), serviceWindowId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceWindowApi.RetrieveServiceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveServiceWindow`: ServiceWindowModel
    fmt.Fprintf(os.Stdout, "Response from `ServiceWindowApi.RetrieveServiceWindow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceWindowId** | **int64** | ID of Service windows scheme to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveServiceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceWindowModel**](ServiceWindowModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetServiceWindow

> ApiResponse SetServiceWindow(ctx).Body(body).Execute()

Add or update a service window



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
    body := *openapiclient.NewServiceWindowModel() // ServiceWindowModel | Service window object that needs to be created or updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceWindowApi.SetServiceWindow(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceWindowApi.SetServiceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetServiceWindow`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceWindowApi.SetServiceWindow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetServiceWindowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceWindowModel**](ServiceWindowModel.md) | Service window object that needs to be created or updated | 

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

