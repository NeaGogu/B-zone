# \TrackTraceApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackTraceCalculateETA**](TrackTraceApi.md#TrackTraceCalculateETA) | **Get** /track-and-trace/calculate-eta/{activityId} | Calculate ETA for Activity
[**TrackTraceCalculateRoutesETA**](TrackTraceApi.md#TrackTraceCalculateRoutesETA) | **Put** /track-and-trace/calculate-routes-eta | Calculate ETA for Activities on Routes



## TrackTraceCalculateETA

> ApiResponse TrackTraceCalculateETA(ctx, activityId).Execute()

Calculate ETA for Activity



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
    activityId := int64(789) // int64 | ID of the activity to calculate ETA for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrackTraceApi.TrackTraceCalculateETA(context.Background(), activityId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackTraceApi.TrackTraceCalculateETA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrackTraceCalculateETA`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TrackTraceApi.TrackTraceCalculateETA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityId** | **int64** | ID of the activity to calculate ETA for | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackTraceCalculateETARequest struct via the builder pattern


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


## TrackTraceCalculateRoutesETA

> RoutesEtaResponse TrackTraceCalculateRoutesETA(ctx).Arguments(arguments).Execute()

Calculate ETA for Activities on Routes



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
    arguments := *openapiclient.NewRoutesEtaArguments() // RoutesEtaArguments | Routes ETA Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrackTraceApi.TrackTraceCalculateRoutesETA(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackTraceApi.TrackTraceCalculateRoutesETA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrackTraceCalculateRoutesETA`: RoutesEtaResponse
    fmt.Fprintf(os.Stdout, "Response from `TrackTraceApi.TrackTraceCalculateRoutesETA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackTraceCalculateRoutesETARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**RoutesEtaArguments**](RoutesEtaArguments.md) | Routes ETA Arguments | 

### Return type

[**RoutesEtaResponse**](RoutesEtaResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

