# \UnsuccessfulReasonApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUnsuccessfulReason**](UnsuccessfulReasonApi.md#CreateUnsuccessfulReason) | **Post** /unsuccessful-reason | Add a new UnsuccessfulReason
[**DeleteUnsuccessfulReason**](UnsuccessfulReasonApi.md#DeleteUnsuccessfulReason) | **Delete** /unsuccessful-reason/{unsuccessful-reasonId} | Delete a UnsuccessfulReason entry
[**RetrieveListUnsuccessfulReason**](UnsuccessfulReasonApi.md#RetrieveListUnsuccessfulReason) | **Put** /unsuccessful-reason | Retrieve List of UnsuccessfulReason
[**RetrieveUnsuccessfulReason**](UnsuccessfulReasonApi.md#RetrieveUnsuccessfulReason) | **Get** /unsuccessful-reason/{unsuccessful-reasonId} | Retrieve a UnsuccessfulReason
[**UpdateUnsuccessfulReason**](UnsuccessfulReasonApi.md#UpdateUnsuccessfulReason) | **Put** /unsuccessful-reason/{unsuccessful-reasonId} | Update a specific UnsuccessfulReason object



## CreateUnsuccessfulReason

> ApiResponse55 CreateUnsuccessfulReason(ctx).Body(body).Execute()

Add a new UnsuccessfulReason



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
    body := *openapiclient.NewUnsuccessfulReasonModel() // UnsuccessfulReasonModel | UnsuccessfulReason object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnsuccessfulReasonApi.CreateUnsuccessfulReason(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsuccessfulReasonApi.CreateUnsuccessfulReason``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUnsuccessfulReason`: ApiResponse55
    fmt.Fprintf(os.Stdout, "Response from `UnsuccessfulReasonApi.CreateUnsuccessfulReason`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUnsuccessfulReasonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UnsuccessfulReasonModel**](UnsuccessfulReasonModel.md) | UnsuccessfulReason object that needs to be created | 

### Return type

[**ApiResponse55**](ApiResponse_55.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUnsuccessfulReason

> ApiResponse53 DeleteUnsuccessfulReason(ctx, unsuccessfulReasonId).Execute()

Delete a UnsuccessfulReason entry



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
    unsuccessfulReasonId := int64(789) // int64 | ID of UnsuccessfulReason to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnsuccessfulReasonApi.DeleteUnsuccessfulReason(context.Background(), unsuccessfulReasonId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsuccessfulReasonApi.DeleteUnsuccessfulReason``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUnsuccessfulReason`: ApiResponse53
    fmt.Fprintf(os.Stdout, "Response from `UnsuccessfulReasonApi.DeleteUnsuccessfulReason`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unsuccessfulReasonId** | **int64** | ID of UnsuccessfulReason to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUnsuccessfulReasonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse53**](ApiResponse_53.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListUnsuccessfulReason

> UnsuccessfulReasonListResponse RetrieveListUnsuccessfulReason(ctx).Arguments(arguments).Execute()

Retrieve List of UnsuccessfulReason



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
    arguments := *openapiclient.NewUnsuccessfulReasonRetrieveListArguments() // UnsuccessfulReasonRetrieveListArguments | UnsuccessfulReason RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnsuccessfulReasonApi.RetrieveListUnsuccessfulReason(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsuccessfulReasonApi.RetrieveListUnsuccessfulReason``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListUnsuccessfulReason`: UnsuccessfulReasonListResponse
    fmt.Fprintf(os.Stdout, "Response from `UnsuccessfulReasonApi.RetrieveListUnsuccessfulReason`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListUnsuccessfulReasonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**UnsuccessfulReasonRetrieveListArguments**](UnsuccessfulReasonRetrieveListArguments.md) | UnsuccessfulReason RetrieveList Arguments | 

### Return type

[**UnsuccessfulReasonListResponse**](UnsuccessfulReasonListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUnsuccessfulReason

> UnsuccessfulReasonModel RetrieveUnsuccessfulReason(ctx, unsuccessfulReasonId).IncludeObjectTypeName(includeObjectTypeName).IncludeRecordInfo(includeRecordInfo).Execute()

Retrieve a UnsuccessfulReason



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
    unsuccessfulReasonId := int64(789) // int64 | ID of UnsuccessfulReason to retrieve
    includeObjectTypeName := true // bool | Show the name of the object type (optional) (default to false)
    includeRecordInfo := true // bool | Show the record info (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnsuccessfulReasonApi.RetrieveUnsuccessfulReason(context.Background(), unsuccessfulReasonId).IncludeObjectTypeName(includeObjectTypeName).IncludeRecordInfo(includeRecordInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsuccessfulReasonApi.RetrieveUnsuccessfulReason``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUnsuccessfulReason`: UnsuccessfulReasonModel
    fmt.Fprintf(os.Stdout, "Response from `UnsuccessfulReasonApi.RetrieveUnsuccessfulReason`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unsuccessfulReasonId** | **int64** | ID of UnsuccessfulReason to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUnsuccessfulReasonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeObjectTypeName** | **bool** | Show the name of the object type | [default to false]
 **includeRecordInfo** | **bool** | Show the record info | [default to false]

### Return type

[**UnsuccessfulReasonModel**](UnsuccessfulReasonModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUnsuccessfulReason

> ApiResponse52 UpdateUnsuccessfulReason(ctx, unsuccessfulReasonId).Body(body).Execute()

Update a specific UnsuccessfulReason object



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
    unsuccessfulReasonId := int64(789) // int64 | ID of the UnsuccessfulReason object to update
    body := *openapiclient.NewUnsuccessfulReasonModel() // UnsuccessfulReasonModel | UnsuccessfulReason object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UnsuccessfulReasonApi.UpdateUnsuccessfulReason(context.Background(), unsuccessfulReasonId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsuccessfulReasonApi.UpdateUnsuccessfulReason``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUnsuccessfulReason`: ApiResponse52
    fmt.Fprintf(os.Stdout, "Response from `UnsuccessfulReasonApi.UpdateUnsuccessfulReason`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unsuccessfulReasonId** | **int64** | ID of the UnsuccessfulReason object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUnsuccessfulReasonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UnsuccessfulReasonModel**](UnsuccessfulReasonModel.md) | UnsuccessfulReason object that needs to be updated | 

### Return type

[**ApiResponse52**](ApiResponse_52.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

