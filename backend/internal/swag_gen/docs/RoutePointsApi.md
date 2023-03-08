# \RoutePointsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveListRoutePoints**](RoutePointsApi.md#RetrieveListRoutePoints) | **Put** /route-points | Find Route Points for multiple routes by route ID
[**RetrieveRoutePoints**](RoutePointsApi.md#RetrieveRoutePoints) | **Get** /route-points/{routeId} | Find Route Points by route ID



## RetrieveListRoutePoints

> []RoutePointsModel RetrieveListRoutePoints(ctx).Arguments(arguments).Execute()

Find Route Points for multiple routes by route ID



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
    arguments := *openapiclient.NewRoutePointsRetrieveListArguments() // RoutePointsRetrieveListArguments | Route Points RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutePointsApi.RetrieveListRoutePoints(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutePointsApi.RetrieveListRoutePoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListRoutePoints`: []RoutePointsModel
    fmt.Fprintf(os.Stdout, "Response from `RoutePointsApi.RetrieveListRoutePoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListRoutePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**RoutePointsRetrieveListArguments**](RoutePointsRetrieveListArguments.md) | Route Points RetrieveList Arguments | 

### Return type

[**[]RoutePointsModel**](RoutePointsModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRoutePoints

> RoutePointsModel RetrieveRoutePoints(ctx, routeId).Execute()

Find Route Points by route ID



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
    routeId := int64(789) // int64 | ID of Route for which to return the Route Points

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutePointsApi.RetrieveRoutePoints(context.Background(), routeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutePointsApi.RetrieveRoutePoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRoutePoints`: RoutePointsModel
    fmt.Fprintf(os.Stdout, "Response from `RoutePointsApi.RetrieveRoutePoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int64** | ID of Route for which to return the Route Points | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRoutePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoutePointsModel**](RoutePointsModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

