# \SystemProviderApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveListSystemProvider**](SystemProviderApi.md#RetrieveListSystemProvider) | **Put** /system-provider | Retrieve List of System Providers
[**RetrieveSystemProvider**](SystemProviderApi.md#RetrieveSystemProvider) | **Get** /system-provider/{providerId} | Retrieve a System Provider



## RetrieveListSystemProvider

> SystemProviderListResponse RetrieveListSystemProvider(ctx).Arguments(arguments).Execute()

Retrieve List of System Providers



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
    arguments := *openapiclient.NewSystemProviderRetrieveListArguments() // SystemProviderRetrieveListArguments | System Provider RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemProviderApi.RetrieveListSystemProvider(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemProviderApi.RetrieveListSystemProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListSystemProvider`: SystemProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemProviderApi.RetrieveListSystemProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListSystemProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**SystemProviderRetrieveListArguments**](SystemProviderRetrieveListArguments.md) | System Provider RetrieveList Arguments | 

### Return type

[**SystemProviderListResponse**](SystemProviderListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSystemProvider

> SystemProviderModel RetrieveSystemProvider(ctx, providerId).Execute()

Retrieve a System Provider



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
    providerId := int64(789) // int64 | ID of System Provider to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemProviderApi.RetrieveSystemProvider(context.Background(), providerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemProviderApi.RetrieveSystemProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSystemProvider`: SystemProviderModel
    fmt.Fprintf(os.Stdout, "Response from `SystemProviderApi.RetrieveSystemProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int64** | ID of System Provider to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSystemProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemProviderModel**](SystemProviderModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

