# \VehicleTypeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVehicleType**](VehicleTypeApi.md#CreateVehicleType) | **Post** /vehicle-type | Add a new VehicleType
[**DeleteVehicleType**](VehicleTypeApi.md#DeleteVehicleType) | **Delete** /vehicle-type/{vehicleTypeId} | Delete a VehicleType entry
[**RetrieveListVehicleType**](VehicleTypeApi.md#RetrieveListVehicleType) | **Put** /vehicle-type | Retrieve List of VehicleTypes
[**RetrieveVehicleType**](VehicleTypeApi.md#RetrieveVehicleType) | **Get** /vehicle-type/{vehicleTypeId} | Retrieve a VehicleType
[**UpdateVehicleType**](VehicleTypeApi.md#UpdateVehicleType) | **Put** /vehicle-type/{vehicleTypeId} | Update a specific VehicleType object



## CreateVehicleType

> VehicleTypeCreateResponse CreateVehicleType(ctx).Body(body).Execute()

Add a new VehicleType



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
    body := *openapiclient.NewVehicleTypeModel() // VehicleTypeModel | VehicleType object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleTypeApi.CreateVehicleType(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleTypeApi.CreateVehicleType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVehicleType`: VehicleTypeCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `VehicleTypeApi.CreateVehicleType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVehicleTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VehicleTypeModel**](VehicleTypeModel.md) | VehicleType object that needs to be created | 

### Return type

[**VehicleTypeCreateResponse**](VehicleTypeCreateResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVehicleType

> VehicleTypeDeleteResponse DeleteVehicleType(ctx, vehicleTypeId).Execute()

Delete a VehicleType entry



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
    vehicleTypeId := int64(789) // int64 | ID of VehicleType to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleTypeApi.DeleteVehicleType(context.Background(), vehicleTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleTypeApi.DeleteVehicleType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVehicleType`: VehicleTypeDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `VehicleTypeApi.DeleteVehicleType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleTypeId** | **int64** | ID of VehicleType to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVehicleTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VehicleTypeDeleteResponse**](VehicleTypeDeleteResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListVehicleType

> VehicleTypeListResponse RetrieveListVehicleType(ctx).Arguments(arguments).Execute()

Retrieve List of VehicleTypes



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
    arguments := *openapiclient.NewVehicleTypeRetrieveListArguments() // VehicleTypeRetrieveListArguments | VehicleType RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleTypeApi.RetrieveListVehicleType(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleTypeApi.RetrieveListVehicleType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListVehicleType`: VehicleTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `VehicleTypeApi.RetrieveListVehicleType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListVehicleTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**VehicleTypeRetrieveListArguments**](VehicleTypeRetrieveListArguments.md) | VehicleType RetrieveList Arguments | 

### Return type

[**VehicleTypeListResponse**](VehicleTypeListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVehicleType

> VehicleTypeModel RetrieveVehicleType(ctx, vehicleTypeId).Execute()

Retrieve a VehicleType



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
    vehicleTypeId := int64(789) // int64 | ID of VehicleType to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleTypeApi.RetrieveVehicleType(context.Background(), vehicleTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleTypeApi.RetrieveVehicleType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveVehicleType`: VehicleTypeModel
    fmt.Fprintf(os.Stdout, "Response from `VehicleTypeApi.RetrieveVehicleType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleTypeId** | **int64** | ID of VehicleType to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVehicleTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VehicleTypeModel**](VehicleTypeModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVehicleType

> VehicleTypeUpdateResponse UpdateVehicleType(ctx, vehicleTypeId).Body(body).Execute()

Update a specific VehicleType object



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
    vehicleTypeId := int64(789) // int64 | ID of the VehicleType object to update
    body := *openapiclient.NewVehicleTypeModel() // VehicleTypeModel | VehicleType object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VehicleTypeApi.UpdateVehicleType(context.Background(), vehicleTypeId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VehicleTypeApi.UpdateVehicleType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVehicleType`: VehicleTypeUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `VehicleTypeApi.UpdateVehicleType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vehicleTypeId** | **int64** | ID of the VehicleType object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVehicleTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VehicleTypeModel**](VehicleTypeModel.md) | VehicleType object that needs to be updated | 

### Return type

[**VehicleTypeUpdateResponse**](VehicleTypeUpdateResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

