# \VehicleApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVehicle**](VehicleApi.md#CreateVehicle) | **Post** /vehicle | Add a new vehicle
[**DeleteVehicle**](VehicleApi.md#DeleteVehicle) | **Delete** /vehicle/{vehicleId} | Delete a vehicle entry
[**RetrieveListVehicle**](VehicleApi.md#RetrieveListVehicle) | **Put** /vehicle | Retrieve List of Vehicles
[**RetrieveVehicle**](VehicleApi.md#RetrieveVehicle) | **Get** /vehicle/{vehicleId} | Find vehicle by ID
[**SetVehicle**](VehicleApi.md#SetVehicle) | **Post** /vehicle/set | Set (create or update) a vehicle
[**UpdateVehicle**](VehicleApi.md#UpdateVehicle) | **Put** /vehicle/{vehicleId} | Update a vehicle



## CreateVehicle

> ApiResponse57 CreateVehicle(ctx).Body(body).Execute()

Add a new vehicle



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
    body := *openapiclient.NewVehicleModel() // VehicleModel | Vehicle object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleApi.CreateVehicle(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleApi.CreateVehicle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVehicle`: ApiResponse57
    fmt.Fprintf(os.Stdout, "Response from `VehicleApi.CreateVehicle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VehicleModel**](VehicleModel.md) | Vehicle object that needs to be created | 

### Return type

[**ApiResponse57**](ApiResponse_57.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVehicle

> ApiResponse58 DeleteVehicle(ctx, vehicleId).Execute()

Delete a vehicle entry



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
    vehicleId := int64(789) // int64 | ID of vehicle to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleApi.DeleteVehicle(context.Background(), vehicleId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleApi.DeleteVehicle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVehicle`: ApiResponse58
    fmt.Fprintf(os.Stdout, "Response from `VehicleApi.DeleteVehicle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleId** | **int64** | ID of vehicle to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse58**](ApiResponse_58.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListVehicle

> VehicleListResponse RetrieveListVehicle(ctx).Arguments(arguments).Execute()

Retrieve List of Vehicles



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
    arguments := *openapiclient.NewVehicleRetrieveListArguments() // VehicleRetrieveListArguments | Vehicle RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleApi.RetrieveListVehicle(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleApi.RetrieveListVehicle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListVehicle`: VehicleListResponse
    fmt.Fprintf(os.Stdout, "Response from `VehicleApi.RetrieveListVehicle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**VehicleRetrieveListArguments**](VehicleRetrieveListArguments.md) | Vehicle RetrieveList Arguments | 

### Return type

[**VehicleListResponse**](VehicleListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVehicle

> VehicleModel RetrieveVehicle(ctx, vehicleId).IncludeVehicleTags(includeVehicleTags).IncludeUpdatedBy(includeUpdatedBy).Execute()

Find vehicle by ID



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
    vehicleId := int64(789) // int64 | ID of vehicle to return
    includeVehicleTags := true // bool | a list of tags bound to vehicle (default to true)
    includeUpdatedBy := true // bool | include updated_by_name (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleApi.RetrieveVehicle(context.Background(), vehicleId).IncludeVehicleTags(includeVehicleTags).IncludeUpdatedBy(includeUpdatedBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleApi.RetrieveVehicle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveVehicle`: VehicleModel
    fmt.Fprintf(os.Stdout, "Response from `VehicleApi.RetrieveVehicle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleId** | **int64** | ID of vehicle to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeVehicleTags** | **bool** | a list of tags bound to vehicle | [default to true]
 **includeUpdatedBy** | **bool** | include updated_by_name | [default to true]

### Return type

[**VehicleModel**](VehicleModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVehicle

> ApiResponse SetVehicle(ctx).Body(body).Execute()

Set (create or update) a vehicle



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
    body := *openapiclient.NewVehicleModel() // VehicleModel | Vehicle object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleApi.SetVehicle(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleApi.SetVehicle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetVehicle`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `VehicleApi.SetVehicle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VehicleModel**](VehicleModel.md) | Vehicle object | 

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


## UpdateVehicle

> ApiResponse UpdateVehicle(ctx, vehicleId).Body(body).Execute()

Update a vehicle



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
    vehicleId := int64(789) // int64 | ID of vehicle to update
    body := *openapiclient.NewVehicleModel() // VehicleModel | Vehicle object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleApi.UpdateVehicle(context.Background(), vehicleId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleApi.UpdateVehicle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVehicle`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `VehicleApi.UpdateVehicle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleId** | **int64** | ID of vehicle to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVehicleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VehicleModel**](VehicleModel.md) | Vehicle object that needs to be updated | 

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

