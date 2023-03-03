# \DriverApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDriver**](DriverApi.md#CreateDriver) | **Post** /driver | Add a driver
[**DeleteDriver**](DriverApi.md#DeleteDriver) | **Delete** /driver/{driverId} | Delete an driver
[**RetrieveDriver**](DriverApi.md#RetrieveDriver) | **Get** /driver/{driverId} | Find driver by ID
[**RetrieveListDriver**](DriverApi.md#RetrieveListDriver) | **Put** /driver | Retrieve List of Drivers
[**SetDriver**](DriverApi.md#SetDriver) | **Post** /driver/set | Set (create or update) a driver
[**UpdateDriver**](DriverApi.md#UpdateDriver) | **Put** /driver/{driverId} | Update a driver



## CreateDriver

> ApiResponse CreateDriver(ctx).Body(body).Execute()

Add a driver



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
    body := *openapiclient.NewDriverModel() // DriverModel | Driver object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverApi.CreateDriver(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverApi.CreateDriver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDriver`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DriverApi.CreateDriver`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DriverModel**](DriverModel.md) | Driver object that needs to be created | 

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


## DeleteDriver

> ApiResponse DeleteDriver(ctx, driverId).Execute()

Delete an driver



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
    driverId := int64(789) // int64 | ID of the driver to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverApi.DeleteDriver(context.Background(), driverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverApi.DeleteDriver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDriver`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DriverApi.DeleteDriver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driverId** | **int64** | ID of the driver to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDriverRequest struct via the builder pattern


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


## RetrieveDriver

> DriverModel RetrieveDriver(ctx, driverId).IncludeDriverTags(includeDriverTags).IncludeUpdatedBy(includeUpdatedBy).Execute()

Find driver by ID



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
    driverId := int64(789) // int64 | ID of driver to return
    includeDriverTags := true // bool | a list of tags bound to driver (default to true)
    includeUpdatedBy := true // bool | include updated_by_name (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverApi.RetrieveDriver(context.Background(), driverId).IncludeDriverTags(includeDriverTags).IncludeUpdatedBy(includeUpdatedBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverApi.RetrieveDriver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDriver`: DriverModel
    fmt.Fprintf(os.Stdout, "Response from `DriverApi.RetrieveDriver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driverId** | **int64** | ID of driver to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDriverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDriverTags** | **bool** | a list of tags bound to driver | [default to true]
 **includeUpdatedBy** | **bool** | include updated_by_name | [default to true]

### Return type

[**DriverModel**](DriverModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListDriver

> DriverListResponse RetrieveListDriver(ctx).Arguments(arguments).Execute()

Retrieve List of Drivers



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
    arguments := *openapiclient.NewDriverRetrieveListArguments() // DriverRetrieveListArguments | Driver RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverApi.RetrieveListDriver(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverApi.RetrieveListDriver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListDriver`: DriverListResponse
    fmt.Fprintf(os.Stdout, "Response from `DriverApi.RetrieveListDriver`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListDriverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**DriverRetrieveListArguments**](DriverRetrieveListArguments.md) | Driver RetrieveList Arguments | 

### Return type

[**DriverListResponse**](DriverListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDriver

> ApiResponse SetDriver(ctx).Body(body).Execute()

Set (create or update) a driver



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
    body := *openapiclient.NewDriverModel() // DriverModel | Driver object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverApi.SetDriver(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverApi.SetDriver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDriver`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DriverApi.SetDriver`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDriverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DriverModel**](DriverModel.md) | Driver object | 

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


## UpdateDriver

> ApiResponse UpdateDriver(ctx, driverId).Body(body).Execute()

Update a driver



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
    driverId := int64(789) // int64 | ID of driver to update
    body := *openapiclient.NewDriverModel() // DriverModel | Driver object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriverApi.UpdateDriver(context.Background(), driverId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DriverApi.UpdateDriver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDriver`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `DriverApi.UpdateDriver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driverId** | **int64** | ID of driver to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DriverModel**](DriverModel.md) | Driver object that needs to be updated | 

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

