# \AddressAppliedApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAddressApplied**](AddressAppliedApi.md#RetrieveAddressApplied) | **Get** /address-applied/{addressId} | Retrieve an Applied Address
[**UpdateAddressApplied**](AddressAppliedApi.md#UpdateAddressApplied) | **Put** /address-applied/{addressId} | Update a AddressApplied



## RetrieveAddressApplied

> AddressAppliedModel RetrieveAddressApplied(ctx, addressId).Execute()

Retrieve an Applied Address



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
    addressId := int64(789) // int64 | ID of the applied address to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressAppliedApi.RetrieveAddressApplied(context.Background(), addressId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAppliedApi.RetrieveAddressApplied``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAddressApplied`: AddressAppliedModel
    fmt.Fprintf(os.Stdout, "Response from `AddressAppliedApi.RetrieveAddressApplied`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int64** | ID of the applied address to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAddressAppliedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressAppliedModel**](AddressAppliedModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddressApplied

> ApiResponse UpdateAddressApplied(ctx, addressId).Body(body).Execute()

Update a AddressApplied



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
    addressId := int64(789) // int64 | ID of address to update
    body := *openapiclient.NewAddressAppliedModel() // AddressAppliedModel | AddressApplied object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressAppliedApi.UpdateAddressApplied(context.Background(), addressId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAppliedApi.UpdateAddressApplied``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAddressApplied`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressAppliedApi.UpdateAddressApplied`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int64** | ID of address to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddressAppliedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AddressAppliedModel**](AddressAppliedModel.md) | AddressApplied object that needs to be updated | 

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

