# \DriverUnavailabilityApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDriverUnavailability**](DriverUnavailabilityApi.md#CreateDriverUnavailability) | **Post** /driver-unavailability | Add a new DriverUnavailability
[**DeleteDriverUnavailability**](DriverUnavailabilityApi.md#DeleteDriverUnavailability) | **Delete** /driver-unavailability/{driverunavailabilityId} | Delete a DriverUnavailability entry
[**RetrieveDriverUnavailability**](DriverUnavailabilityApi.md#RetrieveDriverUnavailability) | **Get** /driver-unavailability/{driverunavailabilityId} | Retrieve a DriverUnavailability
[**RetrieveListDriverUnavailability**](DriverUnavailabilityApi.md#RetrieveListDriverUnavailability) | **Put** /driver-unavailability | Retrieve List of DriverUnavailability
[**SetDriverUnavailability**](DriverUnavailabilityApi.md#SetDriverUnavailability) | **Post** /driver-unavailability/set | Set (create or update) a driver unavailability
[**UpdateDriverUnavailability**](DriverUnavailabilityApi.md#UpdateDriverUnavailability) | **Put** /driver-unavailability/{driverunavailabilityId} | Update a specific DriverUnavailability object



## CreateDriverUnavailability

> ApiResponse4 CreateDriverUnavailability(ctx).Body(body).Execute()

Add a new DriverUnavailability



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
    body := *openapiclient.NewDriverUnavailabilityModel() // DriverUnavailabilityModel | DriverUnavailability object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverUnavailabilityApi.CreateDriverUnavailability(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverUnavailabilityApi.CreateDriverUnavailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDriverUnavailability`: ApiResponse4
    fmt.Fprintf(os.Stdout, "Response from `DriverUnavailabilityApi.CreateDriverUnavailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriverUnavailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DriverUnavailabilityModel**](DriverUnavailabilityModel.md) | DriverUnavailability object that needs to be created | 

### Return type

[**ApiResponse4**](ApiResponse_4.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDriverUnavailability

> ApiResponse2 DeleteDriverUnavailability(ctx, driverunavailabilityId).Execute()

Delete a DriverUnavailability entry



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
    driverunavailabilityId := int64(789) // int64 | ID of DriverUnavailability to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverUnavailabilityApi.DeleteDriverUnavailability(context.Background(), driverunavailabilityId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverUnavailabilityApi.DeleteDriverUnavailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDriverUnavailability`: ApiResponse2
    fmt.Fprintf(os.Stdout, "Response from `DriverUnavailabilityApi.DeleteDriverUnavailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driverunavailabilityId** | **int64** | ID of DriverUnavailability to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDriverUnavailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse2**](ApiResponse_2.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDriverUnavailability

> DriverUnavailabilityModel RetrieveDriverUnavailability(ctx, driverunavailabilityId).IncludeObjectTypeName(includeObjectTypeName).IncludeRecordInfo(includeRecordInfo).Execute()

Retrieve a DriverUnavailability



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
    driverunavailabilityId := int64(789) // int64 | ID of DriverUnavailability to retrieve
    includeObjectTypeName := true // bool | Show the name of the object type (optional) (default to false)
    includeRecordInfo := true // bool | Show the record info (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverUnavailabilityApi.RetrieveDriverUnavailability(context.Background(), driverunavailabilityId).IncludeObjectTypeName(includeObjectTypeName).IncludeRecordInfo(includeRecordInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverUnavailabilityApi.RetrieveDriverUnavailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDriverUnavailability`: DriverUnavailabilityModel
    fmt.Fprintf(os.Stdout, "Response from `DriverUnavailabilityApi.RetrieveDriverUnavailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driverunavailabilityId** | **int64** | ID of DriverUnavailability to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDriverUnavailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeObjectTypeName** | **bool** | Show the name of the object type | [default to false]
 **includeRecordInfo** | **bool** | Show the record info | [default to false]

### Return type

[**DriverUnavailabilityModel**](DriverUnavailabilityModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListDriverUnavailability

> DriverUnavailabilityListResponse RetrieveListDriverUnavailability(ctx).Arguments(arguments).Execute()

Retrieve List of DriverUnavailability



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
    arguments := *openapiclient.NewDriverUnavailabilityRetrieveListArguments() // DriverUnavailabilityRetrieveListArguments | DriverUnavailability RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverUnavailabilityApi.RetrieveListDriverUnavailability(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverUnavailabilityApi.RetrieveListDriverUnavailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListDriverUnavailability`: DriverUnavailabilityListResponse
    fmt.Fprintf(os.Stdout, "Response from `DriverUnavailabilityApi.RetrieveListDriverUnavailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListDriverUnavailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**DriverUnavailabilityRetrieveListArguments**](DriverUnavailabilityRetrieveListArguments.md) | DriverUnavailability RetrieveList Arguments | 

### Return type

[**DriverUnavailabilityListResponse**](DriverUnavailabilityListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDriverUnavailability

> ApiResponse SetDriverUnavailability(ctx).Body(body).Execute()

Set (create or update) a driver unavailability



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
    body := *openapiclient.NewDriverUnavailabilityModel() // DriverUnavailabilityModel | DriverUnavailability object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverUnavailabilityApi.SetDriverUnavailability(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverUnavailabilityApi.SetDriverUnavailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDriverUnavailability`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DriverUnavailabilityApi.SetDriverUnavailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDriverUnavailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DriverUnavailabilityModel**](DriverUnavailabilityModel.md) | DriverUnavailability object | 

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


## UpdateDriverUnavailability

> ApiResponse1 UpdateDriverUnavailability(ctx, driverunavailabilityId).Body(body).Execute()

Update a specific DriverUnavailability object



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
    driverunavailabilityId := int64(789) // int64 | ID of the DriverUnavailability object to update
    body := *openapiclient.NewDriverUnavailabilityModel() // DriverUnavailabilityModel | DriverUnavailability object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverUnavailabilityApi.UpdateDriverUnavailability(context.Background(), driverunavailabilityId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverUnavailabilityApi.UpdateDriverUnavailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDriverUnavailability`: ApiResponse1
    fmt.Fprintf(os.Stdout, "Response from `DriverUnavailabilityApi.UpdateDriverUnavailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driverunavailabilityId** | **int64** | ID of the DriverUnavailability object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriverUnavailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DriverUnavailabilityModel**](DriverUnavailabilityModel.md) | DriverUnavailability object that needs to be updated | 

### Return type

[**ApiResponse1**](ApiResponse_1.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

