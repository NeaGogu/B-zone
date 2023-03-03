# \RouteApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockRoutes**](RouteApi.md#BlockRoutes) | **Post** /route/block-routes | Block routes which satisfy set filters
[**CreateRoute**](RouteApi.md#CreateRoute) | **Post** /route | Add a new Route
[**DeleteRoute**](RouteApi.md#DeleteRoute) | **Delete** /route/{routeId} | Delete an Route
[**GetExecutableActivities**](RouteApi.md#GetExecutableActivities) | **Post** /route/get-executable-activities | Returns all activities in this route which hav enot been executed yet.
[**RetrieveListRoute**](RouteApi.md#RetrieveListRoute) | **Put** /route | Retrieve List of Routes
[**RetrieveRoute**](RouteApi.md#RetrieveRoute) | **Get** /route/{routeId} | Retrieve a Route
[**RouteStoreGeoLocations**](RouteApi.md#RouteStoreGeoLocations) | **Post** /route/store-geo-locations | Store tracked Geo Locations in bulk
[**SetRoute**](RouteApi.md#SetRoute) | **Post** /route/set | Set (create or update) an Route
[**UnblockRoutes**](RouteApi.md#UnblockRoutes) | **Post** /route/unblock-routes | Unblock routes which satisfy set filters
[**UpdateRoute**](RouteApi.md#UpdateRoute) | **Put** /route/{routeId} | Update a Route



## BlockRoutes

> ApiResponse BlockRoutes(ctx).Filters(filters).Execute()

Block routes which satisfy set filters



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
    filters := *openapiclient.NewRouteFiltersModel() // RouteFiltersModel | Request Filters

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.BlockRoutes(context.Background()).Filters(filters).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.BlockRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockRoutes`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.BlockRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**RouteFiltersModel**](RouteFiltersModel.md) | Request Filters | 

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


## CreateRoute

> ApiResponse CreateRoute(ctx).Body(body).Execute()

Add a new Route



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
    body := *openapiclient.NewRouteModel() // RouteModel | Route object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.CreateRoute(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.CreateRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoute`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.CreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RouteModel**](RouteModel.md) | Route object that needs to be created | 

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


## DeleteRoute

> ApiResponse DeleteRoute(ctx, routeId).CancelActivities(cancelActivities).Execute()

Delete an Route



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
    routeId := int64(789) // int64 | ID of route to update
    cancelActivities := true // bool | Cancel activities on Route

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.DeleteRoute(context.Background(), routeId).CancelActivities(cancelActivities).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.DeleteRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoute`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.DeleteRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int64** | ID of route to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelActivities** | **bool** | Cancel activities on Route | 

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


## GetExecutableActivities

> ApiResponse GetExecutableActivities(ctx).Arguments(arguments).Execute()

Returns all activities in this route which hav enot been executed yet.



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
    arguments := *openapiclient.NewGetExecutableActivitiesArguments() // GetExecutableActivitiesArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.GetExecutableActivities(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.GetExecutableActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecutableActivities`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.GetExecutableActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutableActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**GetExecutableActivitiesArguments**](GetExecutableActivitiesArguments.md) | Request Arguments | 

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


## RetrieveListRoute

> RouteListResponse RetrieveListRoute(ctx).Arguments(arguments).Execute()

Retrieve List of Routes



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
    arguments := *openapiclient.NewRouteRetrieveListArguments() // RouteRetrieveListArguments | Route RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.RetrieveListRoute(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.RetrieveListRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListRoute`: RouteListResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.RetrieveListRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**RouteRetrieveListArguments**](RouteRetrieveListArguments.md) | Route RetrieveList Arguments | 

### Return type

[**RouteListResponse**](RouteListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRoute

> RouteModel RetrieveRoute(ctx, routeId).IncludeAddressObject(includeAddressObject).IncludeRouteStatus(includeRouteStatus).IncludeRouteTags(includeRouteTags).IncludeDriverInfo(includeDriverInfo).IncludeEquipmentInfoCar(includeEquipmentInfoCar).IncludeGpsLocations(includeGpsLocations).IncludeActivityIds(includeActivityIds).IncludeLatestPosition(includeLatestPosition).Execute()

Retrieve a Route



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
    routeId := int64(789) // int64 | ID of route to retrieve
    includeAddressObject := true // bool | Include Address Objects (default to true)
    includeRouteStatus := true // bool | Include Status Name (default to true)
    includeRouteTags := true // bool | Include Tags (default to true)
    includeDriverInfo := true // bool | Include Driver info (default to true)
    includeEquipmentInfoCar := true // bool | Include Equipment info car (default to true)
    includeGpsLocations := true // bool | Include GPS locations (default to false)
    includeActivityIds := true // bool | Include Activity IDs (default to false)
    includeLatestPosition := true // bool | Include Latest Known GPS location (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.RetrieveRoute(context.Background(), routeId).IncludeAddressObject(includeAddressObject).IncludeRouteStatus(includeRouteStatus).IncludeRouteTags(includeRouteTags).IncludeDriverInfo(includeDriverInfo).IncludeEquipmentInfoCar(includeEquipmentInfoCar).IncludeGpsLocations(includeGpsLocations).IncludeActivityIds(includeActivityIds).IncludeLatestPosition(includeLatestPosition).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.RetrieveRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRoute`: RouteModel
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.RetrieveRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int64** | ID of route to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAddressObject** | **bool** | Include Address Objects | [default to true]
 **includeRouteStatus** | **bool** | Include Status Name | [default to true]
 **includeRouteTags** | **bool** | Include Tags | [default to true]
 **includeDriverInfo** | **bool** | Include Driver info | [default to true]
 **includeEquipmentInfoCar** | **bool** | Include Equipment info car | [default to true]
 **includeGpsLocations** | **bool** | Include GPS locations | [default to false]
 **includeActivityIds** | **bool** | Include Activity IDs | [default to false]
 **includeLatestPosition** | **bool** | Include Latest Known GPS location | [default to false]

### Return type

[**RouteModel**](RouteModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RouteStoreGeoLocations

> ApiResponse RouteStoreGeoLocations(ctx).Arguments(arguments).Execute()

Store tracked Geo Locations in bulk



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
    arguments := *openapiclient.NewRouteStoreGeoLocations() // RouteStoreGeoLocations | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.RouteStoreGeoLocations(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.RouteStoreGeoLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RouteStoreGeoLocations`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.RouteStoreGeoLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRouteStoreGeoLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**RouteStoreGeoLocations**](RouteStoreGeoLocations.md) | Request Arguments | 

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


## SetRoute

> ApiResponse SetRoute(ctx).Body(body).Execute()

Set (create or update) an Route



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
    body := *openapiclient.NewRouteModel() // RouteModel | Route object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.SetRoute(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.SetRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRoute`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.SetRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RouteModel**](RouteModel.md) | Route object | 

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


## UnblockRoutes

> ApiResponse UnblockRoutes(ctx).Filters(filters).Execute()

Unblock routes which satisfy set filters



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
    filters := *openapiclient.NewRouteFiltersModel() // RouteFiltersModel | Request Filters

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.UnblockRoutes(context.Background()).Filters(filters).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.UnblockRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnblockRoutes`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.UnblockRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnblockRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**RouteFiltersModel**](RouteFiltersModel.md) | Request Filters | 

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


## UpdateRoute

> ApiResponse UpdateRoute(ctx, routeId).Execute()

Update a Route



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
    routeId := int64(789) // int64 | ID of route to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RouteApi.UpdateRoute(context.Background(), routeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.UpdateRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoute`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int64** | ID of route to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


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

