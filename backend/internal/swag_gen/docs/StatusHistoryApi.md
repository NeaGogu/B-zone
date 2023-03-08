# \StatusHistoryApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveStatusHistoryActivity**](StatusHistoryApi.md#RetrieveStatusHistoryActivity) | **Post** /status-history/activity | Retrieve the status history of an activity
[**RetrieveStatusHistoryRoute**](StatusHistoryApi.md#RetrieveStatusHistoryRoute) | **Post** /status-history/route | Retrieve the status history of a route



## RetrieveStatusHistoryActivity

> StatusHistoryListModel RetrieveStatusHistoryActivity(ctx).Arguments(arguments).Execute()

Retrieve the status history of an activity



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
    arguments := *openapiclient.NewStatusHistoryRequestModel() // StatusHistoryRequestModel | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatusHistoryApi.RetrieveStatusHistoryActivity(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusHistoryApi.RetrieveStatusHistoryActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveStatusHistoryActivity`: StatusHistoryListModel
    fmt.Fprintf(os.Stdout, "Response from `StatusHistoryApi.RetrieveStatusHistoryActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveStatusHistoryActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**StatusHistoryRequestModel**](StatusHistoryRequestModel.md) | Request Arguments | 

### Return type

[**StatusHistoryListModel**](StatusHistoryListModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveStatusHistoryRoute

> StatusHistoryListModel RetrieveStatusHistoryRoute(ctx).Arguments(arguments).Execute()

Retrieve the status history of a route



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
    arguments := *openapiclient.NewStatusHistoryRequestModel() // StatusHistoryRequestModel | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatusHistoryApi.RetrieveStatusHistoryRoute(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusHistoryApi.RetrieveStatusHistoryRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveStatusHistoryRoute`: StatusHistoryListModel
    fmt.Fprintf(os.Stdout, "Response from `StatusHistoryApi.RetrieveStatusHistoryRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveStatusHistoryRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**StatusHistoryRequestModel**](StatusHistoryRequestModel.md) | Request Arguments | 

### Return type

[**StatusHistoryListModel**](StatusHistoryListModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

