# \EquipmentApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEquipment**](EquipmentApi.md#CreateEquipment) | **Post** /equipment | Add a new Equipment
[**DeleteEquipment**](EquipmentApi.md#DeleteEquipment) | **Delete** /equipment/{equipmentId} | Delete an Equipment entry
[**RetrieveEquipment**](EquipmentApi.md#RetrieveEquipment) | **Get** /equipment/{equipmentId} | Retrieve a Equipment
[**RetrieveListEquipment**](EquipmentApi.md#RetrieveListEquipment) | **Put** /equipment | Retrieve List of Equipment
[**SetEquipment**](EquipmentApi.md#SetEquipment) | **Post** /equipment/set | Set (create or update) a Equipment
[**UpdateEquipment**](EquipmentApi.md#UpdateEquipment) | **Put** /equipment/{equipmentId} | Update a specific Equipment object



## CreateEquipment

> EquipmentCreateResponse CreateEquipment(ctx).Body(body).Execute()

Add a new Equipment



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
    body := *openapiclient.NewEquipmentModel() // EquipmentModel | Equipment object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.CreateEquipment(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.CreateEquipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEquipment`: EquipmentCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.CreateEquipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EquipmentModel**](EquipmentModel.md) | Equipment object that needs to be created | 

### Return type

[**EquipmentCreateResponse**](EquipmentCreateResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEquipment

> EquipmentDeleteResponse DeleteEquipment(ctx, equipmentId).Execute()

Delete an Equipment entry



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
    equipmentId := int64(789) // int64 | ID of Equipment to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.DeleteEquipment(context.Background(), equipmentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.DeleteEquipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEquipment`: EquipmentDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.DeleteEquipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**equipmentId** | **int64** | ID of Equipment to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EquipmentDeleteResponse**](EquipmentDeleteResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEquipment

> EquipmentModel RetrieveEquipment(ctx, equipmentId).Execute()

Retrieve a Equipment



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
    equipmentId := int64(789) // int64 | ID of Equipment to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.RetrieveEquipment(context.Background(), equipmentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.RetrieveEquipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveEquipment`: EquipmentModel
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.RetrieveEquipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**equipmentId** | **int64** | ID of Equipment to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EquipmentModel**](EquipmentModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListEquipment

> EquipmentListResponse RetrieveListEquipment(ctx).Arguments(arguments).Execute()

Retrieve List of Equipment



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
    arguments := *openapiclient.NewEquipmentRetrieveListArguments() // EquipmentRetrieveListArguments | Equipment RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.RetrieveListEquipment(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.RetrieveListEquipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListEquipment`: EquipmentListResponse
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.RetrieveListEquipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**EquipmentRetrieveListArguments**](EquipmentRetrieveListArguments.md) | Equipment RetrieveList Arguments | 

### Return type

[**EquipmentListResponse**](EquipmentListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEquipment

> EquipmentSetResponse SetEquipment(ctx).Body(body).Execute()

Set (create or update) a Equipment



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
    body := *openapiclient.NewEquipmentModel() // EquipmentModel | Equipment object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.SetEquipment(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.SetEquipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEquipment`: EquipmentSetResponse
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.SetEquipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EquipmentModel**](EquipmentModel.md) | Equipment object | 

### Return type

[**EquipmentSetResponse**](EquipmentSetResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEquipment

> EquipmentUpdateResponse UpdateEquipment(ctx, equipmentId).Body(body).Execute()

Update a specific Equipment object



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
    equipmentId := int64(789) // int64 | ID of the Equipment object to update
    body := *openapiclient.NewEquipmentModel() // EquipmentModel | Equipment object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EquipmentApi.UpdateEquipment(context.Background(), equipmentId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentApi.UpdateEquipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEquipment`: EquipmentUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `EquipmentApi.UpdateEquipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**equipmentId** | **int64** | ID of the Equipment object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EquipmentModel**](EquipmentModel.md) | Equipment object that needs to be updated | 

### Return type

[**EquipmentUpdateResponse**](EquipmentUpdateResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

