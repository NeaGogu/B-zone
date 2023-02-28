# \BlockedDateApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBlockedDate**](BlockedDateApi.md#DeleteBlockedDate) | **Delete** /blocked-date/{blockedDateId} | Delete a blocked date
[**RetrieveListBlockedDate**](BlockedDateApi.md#RetrieveListBlockedDate) | **Put** /blocked-date | Retrieve List of blocked dates
[**SetBlockedDate**](BlockedDateApi.md#SetBlockedDate) | **Post** /blocked-date/set | Set (create or update) a blocked date



## DeleteBlockedDate

> ApiResponse DeleteBlockedDate(ctx, blockedDateId).Execute()

Delete a blocked date



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
    blockedDateId := int64(789) // int64 | ID of blocked date to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockedDateApi.DeleteBlockedDate(context.Background(), blockedDateId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockedDateApi.DeleteBlockedDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlockedDate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `BlockedDateApi.DeleteBlockedDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockedDateId** | **int64** | ID of blocked date to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockedDateRequest struct via the builder pattern


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


## RetrieveListBlockedDate

> BlockedDateListResponse RetrieveListBlockedDate(ctx).Arguments(arguments).Execute()

Retrieve List of blocked dates



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
    arguments := *openapiclient.NewBlockedDateRetrieveListArguments() // BlockedDateRetrieveListArguments | Blocked date RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockedDateApi.RetrieveListBlockedDate(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockedDateApi.RetrieveListBlockedDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListBlockedDate`: BlockedDateListResponse
    fmt.Fprintf(os.Stdout, "Response from `BlockedDateApi.RetrieveListBlockedDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListBlockedDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**BlockedDateRetrieveListArguments**](BlockedDateRetrieveListArguments.md) | Blocked date RetrieveList Arguments | 

### Return type

[**BlockedDateListResponse**](BlockedDateListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBlockedDate

> ApiResponse SetBlockedDate(ctx).Body(body).Execute()

Set (create or update) a blocked date



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
    body := *openapiclient.NewBlockedDateModel() // BlockedDateModel | Blocked date object object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockedDateApi.SetBlockedDate(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockedDateApi.SetBlockedDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetBlockedDate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `BlockedDateApi.SetBlockedDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBlockedDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BlockedDateModel**](BlockedDateModel.md) | Blocked date object object | 

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

