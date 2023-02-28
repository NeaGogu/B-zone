# \CapacityTypeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCapacityType**](CapacityTypeApi.md#DeleteCapacityType) | **Delete** /capacity-type/{capacityTypeId} | Delete a capacity-type
[**RetrieveCapacityType**](CapacityTypeApi.md#RetrieveCapacityType) | **Get** /capacity-type/{capacityTypeId} | Find capacity-type by ID
[**RetrieveListCapacityType**](CapacityTypeApi.md#RetrieveListCapacityType) | **Put** /capacity-type | Retrieve List of CapacityTypes
[**SetCapacityType**](CapacityTypeApi.md#SetCapacityType) | **Post** /capacity-type/set | Set (create or update) an CapacityType



## DeleteCapacityType

> ApiResponse DeleteCapacityType(ctx, capacityTypeId).Execute()

Delete a capacity-type



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
    capacityTypeId := int64(789) // int64 | ID of the capacity-type to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CapacityTypeApi.DeleteCapacityType(context.Background(), capacityTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityTypeApi.DeleteCapacityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCapacityType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CapacityTypeApi.DeleteCapacityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**capacityTypeId** | **int64** | ID of the capacity-type to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCapacityTypeRequest struct via the builder pattern


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


## RetrieveCapacityType

> CapacityTypeModel RetrieveCapacityType(ctx, capacityTypeId).IncludeUom(includeUom).IncludeUomGroup(includeUomGroup).IncludeUomName(includeUomName).Execute()

Find capacity-type by ID



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
    capacityTypeId := int64(789) // int64 | ID of capacity-type to return
    includeUom := true // bool | Include uom object (optional) (default to false)
    includeUomGroup := true // bool | Include uom group (optional) (default to false)
    includeUomName := true // bool | Include uom name (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CapacityTypeApi.RetrieveCapacityType(context.Background(), capacityTypeId).IncludeUom(includeUom).IncludeUomGroup(includeUomGroup).IncludeUomName(includeUomName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityTypeApi.RetrieveCapacityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCapacityType`: CapacityTypeModel
    fmt.Fprintf(os.Stdout, "Response from `CapacityTypeApi.RetrieveCapacityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**capacityTypeId** | **int64** | ID of capacity-type to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCapacityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUom** | **bool** | Include uom object | [default to false]
 **includeUomGroup** | **bool** | Include uom group | [default to false]
 **includeUomName** | **bool** | Include uom name | [default to false]

### Return type

[**CapacityTypeModel**](CapacityTypeModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListCapacityType

> CapacityTypeListResponse RetrieveListCapacityType(ctx).Arguments(arguments).Execute()

Retrieve List of CapacityTypes



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
    arguments := *openapiclient.NewCapacityTypeRetrieveListArguments() // CapacityTypeRetrieveListArguments | CapacityType RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CapacityTypeApi.RetrieveListCapacityType(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityTypeApi.RetrieveListCapacityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListCapacityType`: CapacityTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `CapacityTypeApi.RetrieveListCapacityType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListCapacityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**CapacityTypeRetrieveListArguments**](CapacityTypeRetrieveListArguments.md) | CapacityType RetrieveList Arguments | 

### Return type

[**CapacityTypeListResponse**](CapacityTypeListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCapacityType

> ApiResponse SetCapacityType(ctx).Body(body).Execute()

Set (create or update) an CapacityType



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
    body := *openapiclient.NewCapacityTypeModel() // CapacityTypeModel | CapacityType model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CapacityTypeApi.SetCapacityType(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CapacityTypeApi.SetCapacityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCapacityType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CapacityTypeApi.SetCapacityType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetCapacityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CapacityTypeModel**](CapacityTypeModel.md) | CapacityType model | 

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

