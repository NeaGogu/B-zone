# \ZoneApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateZone**](ZoneApi.md#CreateZone) | **Post** /zone | Add a new Zone
[**DeleteZone**](ZoneApi.md#DeleteZone) | **Delete** /zone/{zoneId} | Delete a Zone
[**RetrieveListZone**](ZoneApi.md#RetrieveListZone) | **Put** /zone | Retrieve List of Zone
[**RetrieveZone**](ZoneApi.md#RetrieveZone) | **Get** /zone/{zoneId} | Retrieve a Zone
[**SetZone**](ZoneApi.md#SetZone) | **Post** /zone/set | Set (create or update) a Zone
[**UpdateZone**](ZoneApi.md#UpdateZone) | **Put** /zone/{zoneId} | Update a Zone



## CreateZone

> ApiResponse CreateZone(ctx).Body(body).Execute()

Add a new Zone



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
    body := *openapiclient.NewZoneModel() // ZoneModel | Zone object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.CreateZone(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.CreateZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZone`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.CreateZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ZoneModel**](ZoneModel.md) | Zone object that needs to be created | 

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


## DeleteZone

> ApiResponse DeleteZone(ctx, zoneId).Execute()

Delete a Zone



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
    zoneId := int64(789) // int64 | ID of zone to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.DeleteZone(context.Background(), zoneId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.DeleteZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteZone`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.DeleteZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int64** | ID of zone to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneRequest struct via the builder pattern


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


## RetrieveListZone

> ZoneListResponse RetrieveListZone(ctx).Arguments(arguments).Execute()

Retrieve List of Zone



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
    arguments := *openapiclient.NewZoneRetrieveListArguments() // ZoneRetrieveListArguments | Zone RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.RetrieveListZone(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.RetrieveListZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListZone`: ZoneListResponse
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.RetrieveListZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**ZoneRetrieveListArguments**](ZoneRetrieveListArguments.md) | Zone RetrieveList Arguments | 

### Return type

[**ZoneListResponse**](ZoneListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveZone

> ZoneModel RetrieveZone(ctx, zoneId).Execute()

Retrieve a Zone



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
    zoneId := int64(789) // int64 | ID of zone to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.RetrieveZone(context.Background(), zoneId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.RetrieveZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveZone`: ZoneModel
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.RetrieveZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int64** | ID of zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ZoneModel**](ZoneModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetZone

> ApiResponse SetZone(ctx).Body(body).Execute()

Set (create or update) a Zone



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
    body := *openapiclient.NewZoneModel() // ZoneModel | Zone object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.SetZone(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.SetZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetZone`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.SetZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ZoneModel**](ZoneModel.md) | Zone object | 

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


## UpdateZone

> ApiResponse UpdateZone(ctx, zoneId).Execute()

Update a Zone



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
    zoneId := int64(789) // int64 | ID of zone to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZoneApi.UpdateZone(context.Background(), zoneId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ZoneApi.UpdateZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateZone`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ZoneApi.UpdateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int64** | ID of zone to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateZoneRequest struct via the builder pattern


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

