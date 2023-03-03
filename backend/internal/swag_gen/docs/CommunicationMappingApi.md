# \CommunicationMappingApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommunicationMapping**](CommunicationMappingApi.md#CreateCommunicationMapping) | **Post** /communication-mapping/set | Create or update a communication mapping
[**DeleteCommunicationMapping**](CommunicationMappingApi.md#DeleteCommunicationMapping) | **Delete** /communication-mapping/{communicationMappingId} | Delete a communication mapping
[**RetrieveCommunicationMapping**](CommunicationMappingApi.md#RetrieveCommunicationMapping) | **Get** /communication-mapping/{communicationMappingId} | Retrieve a communication mapping
[**RetrieveListCommunicationMapping**](CommunicationMappingApi.md#RetrieveListCommunicationMapping) | **Put** /communication-mapping | Retrieve list of communication mappings



## CreateCommunicationMapping

> ApiResponse CreateCommunicationMapping(ctx).Body(body).Execute()

Create or update a communication mapping



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
    body := *openapiclient.NewCommunicationMappingModel() // CommunicationMappingModel | Mapping object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMappingApi.CreateCommunicationMapping(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMappingApi.CreateCommunicationMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommunicationMapping`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMappingApi.CreateCommunicationMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommunicationMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommunicationMappingModel**](CommunicationMappingModel.md) | Mapping object | 

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


## DeleteCommunicationMapping

> ApiResponse DeleteCommunicationMapping(ctx, communicationMappingId).Execute()

Delete a communication mapping



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
    communicationMappingId := int64(789) // int64 | ID of the communication mapping to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMappingApi.DeleteCommunicationMapping(context.Background(), communicationMappingId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMappingApi.DeleteCommunicationMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCommunicationMapping`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMappingApi.DeleteCommunicationMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationMappingId** | **int64** | ID of the communication mapping to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommunicationMappingRequest struct via the builder pattern


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


## RetrieveCommunicationMapping

> CommunicationMappingModel RetrieveCommunicationMapping(ctx, communicationMappingId).Execute()

Retrieve a communication mapping



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
    communicationMappingId := int64(789) // int64 | ID of communication mapping to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMappingApi.RetrieveCommunicationMapping(context.Background(), communicationMappingId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMappingApi.RetrieveCommunicationMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCommunicationMapping`: CommunicationMappingModel
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMappingApi.RetrieveCommunicationMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationMappingId** | **int64** | ID of communication mapping to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCommunicationMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommunicationMappingModel**](CommunicationMappingModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListCommunicationMapping

> CommunicationMappingListResponse RetrieveListCommunicationMapping(ctx).Arguments(arguments).Execute()

Retrieve list of communication mappings



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
    arguments := *openapiclient.NewCommunicationMappingRetrieveListArguments() // CommunicationMappingRetrieveListArguments | Communication Mapping RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationMappingApi.RetrieveListCommunicationMapping(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMappingApi.RetrieveListCommunicationMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListCommunicationMapping`: CommunicationMappingListResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMappingApi.RetrieveListCommunicationMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListCommunicationMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**CommunicationMappingRetrieveListArguments**](CommunicationMappingRetrieveListArguments.md) | Communication Mapping RetrieveList Arguments | 

### Return type

[**CommunicationMappingListResponse**](CommunicationMappingListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

