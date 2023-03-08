# \ServiceWindowsSchemeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceWindowsScheme**](ServiceWindowsSchemeApi.md#DeleteServiceWindowsScheme) | **Delete** /service-windows-scheme/{serviceWindowsSchemeId} | Delete a service windows scheme
[**RetrieveListServiceWindowsScheme**](ServiceWindowsSchemeApi.md#RetrieveListServiceWindowsScheme) | **Put** /service-windows-scheme | Retrieve a list of service windows schemes
[**RetrieveServiceWindowsScheme**](ServiceWindowsSchemeApi.md#RetrieveServiceWindowsScheme) | **Get** /service-windows-scheme/{serviceWindowsSchemeId} | Retrieve a single service windows scheme
[**SetServiceWindowsScheme**](ServiceWindowsSchemeApi.md#SetServiceWindowsScheme) | **Post** /service-windows-scheme/set | Add or update a service windows scheme



## DeleteServiceWindowsScheme

> ApiResponse DeleteServiceWindowsScheme(ctx, serviceWindowsSchemeId).Execute()

Delete a service windows scheme



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
    serviceWindowsSchemeId := int64(789) // int64 | ID of Service Windows Scheme to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceWindowsSchemeApi.DeleteServiceWindowsScheme(context.Background(), serviceWindowsSchemeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceWindowsSchemeApi.DeleteServiceWindowsScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServiceWindowsScheme`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceWindowsSchemeApi.DeleteServiceWindowsScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceWindowsSchemeId** | **int64** | ID of Service Windows Scheme to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceWindowsSchemeRequest struct via the builder pattern


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


## RetrieveListServiceWindowsScheme

> ServiceWindowsSchemeListResponse RetrieveListServiceWindowsScheme(ctx).Arguments(arguments).Execute()

Retrieve a list of service windows schemes



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
    arguments := *openapiclient.NewServiceWindowsSchemeRetrieveListArguments() // ServiceWindowsSchemeRetrieveListArguments | Service Windows Scheme RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceWindowsSchemeApi.RetrieveListServiceWindowsScheme(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceWindowsSchemeApi.RetrieveListServiceWindowsScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListServiceWindowsScheme`: ServiceWindowsSchemeListResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceWindowsSchemeApi.RetrieveListServiceWindowsScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListServiceWindowsSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**ServiceWindowsSchemeRetrieveListArguments**](ServiceWindowsSchemeRetrieveListArguments.md) | Service Windows Scheme RetrieveList Arguments | 

### Return type

[**ServiceWindowsSchemeListResponse**](ServiceWindowsSchemeListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveServiceWindowsScheme

> ServiceWindowsSchemeModel RetrieveServiceWindowsScheme(ctx, serviceWindowsSchemeId).Execute()

Retrieve a single service windows scheme



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
    serviceWindowsSchemeId := int64(789) // int64 | ID of Service windows scheme to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceWindowsSchemeApi.RetrieveServiceWindowsScheme(context.Background(), serviceWindowsSchemeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceWindowsSchemeApi.RetrieveServiceWindowsScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveServiceWindowsScheme`: ServiceWindowsSchemeModel
    fmt.Fprintf(os.Stdout, "Response from `ServiceWindowsSchemeApi.RetrieveServiceWindowsScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceWindowsSchemeId** | **int64** | ID of Service windows scheme to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveServiceWindowsSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceWindowsSchemeModel**](ServiceWindowsSchemeModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetServiceWindowsScheme

> ApiResponse SetServiceWindowsScheme(ctx).Body(body).Execute()

Add or update a service windows scheme



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
    body := *openapiclient.NewServiceWindowsSchemeModel() // ServiceWindowsSchemeModel | Service windows scheme object that needs to be created or updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceWindowsSchemeApi.SetServiceWindowsScheme(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceWindowsSchemeApi.SetServiceWindowsScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetServiceWindowsScheme`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceWindowsSchemeApi.SetServiceWindowsScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetServiceWindowsSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceWindowsSchemeModel**](ServiceWindowsSchemeModel.md) | Service windows scheme object that needs to be created or updated | 

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

