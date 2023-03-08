# \ProviderApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProvider**](ProviderApi.md#CreateProvider) | **Post** /provider | Add a new Provider
[**RetrieveListProvider**](ProviderApi.md#RetrieveListProvider) | **Put** /provider | Retrieve List of Providers
[**RetrieveProvider**](ProviderApi.md#RetrieveProvider) | **Get** /provider/{providerId} | Retrieve a Provider
[**UpdateProvider**](ProviderApi.md#UpdateProvider) | **Put** /provider/{providerId} | Update a specific provider object



## CreateProvider

> ApiResponse22 CreateProvider(ctx).Body(body).Execute()

Add a new Provider



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
    body := *openapiclient.NewProviderModel() // ProviderModel | Provider object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderApi.CreateProvider(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.CreateProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProvider`: ApiResponse22
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.CreateProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProviderModel**](ProviderModel.md) | Provider object that needs to be created | 

### Return type

[**ApiResponse22**](ApiResponse_22.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListProvider

> ProviderListResponse RetrieveListProvider(ctx).Arguments(arguments).Execute()

Retrieve List of Providers



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
    arguments := *openapiclient.NewProviderRetrieveListArguments() // ProviderRetrieveListArguments | Provider RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderApi.RetrieveListProvider(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.RetrieveListProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListProvider`: ProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.RetrieveListProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**ProviderRetrieveListArguments**](ProviderRetrieveListArguments.md) | Provider RetrieveList Arguments | 

### Return type

[**ProviderListResponse**](ProviderListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveProvider

> ProviderModel RetrieveProvider(ctx, providerId).Execute()

Retrieve a Provider



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
    providerId := int64(789) // int64 | ID of provider to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderApi.RetrieveProvider(context.Background(), providerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.RetrieveProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveProvider`: ProviderModel
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.RetrieveProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int64** | ID of provider to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderModel**](ProviderModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvider

> ApiResponse21 UpdateProvider(ctx, providerId).Body(body).Execute()

Update a specific provider object



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
    providerId := int64(789) // int64 | ID of the provider object to update
    body := *openapiclient.NewProviderModel() // ProviderModel | Provider object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProviderApi.UpdateProvider(context.Background(), providerId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.UpdateProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProvider`: ApiResponse21
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.UpdateProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int64** | ID of the provider object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProviderModel**](ProviderModel.md) | Provider object that needs to be updated | 

### Return type

[**ApiResponse21**](ApiResponse_21.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

