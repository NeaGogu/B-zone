# \PackageTypeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePackageType**](PackageTypeApi.md#CreatePackageType) | **Post** /package-type | Create or update a Package Line
[**DeletePackageType**](PackageTypeApi.md#DeletePackageType) | **Delete** /package-type/{packageTypeId} | Delete an package-type
[**RetrieveListPackageType**](PackageTypeApi.md#RetrieveListPackageType) | **Put** /package-type | Retrieve List of PackageTypes
[**RetrievePackageType**](PackageTypeApi.md#RetrievePackageType) | **Get** /package-type/{packageTypeId} | Find package-type by ID
[**SetPackageType**](PackageTypeApi.md#SetPackageType) | **Post** /package-type/set | Set (create or update) an PackageType
[**UpdatePackageType**](PackageTypeApi.md#UpdatePackageType) | **Put** /package-type/{packageTypeId} | Update a package-type



## CreatePackageType

> ApiResponse CreatePackageType(ctx).Body(body).Execute()

Create or update a Package Line



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
    body := *openapiclient.NewPackageTypeModel() // PackageTypeModel | PackageType model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageTypeApi.CreatePackageType(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageTypeApi.CreatePackageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePackageType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageTypeApi.CreatePackageType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePackageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PackageTypeModel**](PackageTypeModel.md) | PackageType model | 

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


## DeletePackageType

> ApiResponse DeletePackageType(ctx, packageTypeId).Execute()

Delete an package-type



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
    packageTypeId := int64(789) // int64 | ID of the package-type to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageTypeApi.DeletePackageType(context.Background(), packageTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageTypeApi.DeletePackageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePackageType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageTypeApi.DeletePackageType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageTypeId** | **int64** | ID of the package-type to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePackageTypeRequest struct via the builder pattern


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


## RetrieveListPackageType

> PackageTypeListResponse RetrieveListPackageType(ctx).Arguments(arguments).Execute()

Retrieve List of PackageTypes



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
    arguments := *openapiclient.NewPackageTypeRetrieveListArguments() // PackageTypeRetrieveListArguments | PackageType RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageTypeApi.RetrieveListPackageType(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageTypeApi.RetrieveListPackageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListPackageType`: PackageTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageTypeApi.RetrieveListPackageType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListPackageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**PackageTypeRetrieveListArguments**](PackageTypeRetrieveListArguments.md) | PackageType RetrieveList Arguments | 

### Return type

[**PackageTypeListResponse**](PackageTypeListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePackageType

> PackageTypeModel RetrievePackageType(ctx, packageTypeId).Execute()

Find package-type by ID



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
    packageTypeId := int64(789) // int64 | ID of package-type to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageTypeApi.RetrievePackageType(context.Background(), packageTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageTypeApi.RetrievePackageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePackageType`: PackageTypeModel
    fmt.Fprintf(os.Stdout, "Response from `PackageTypeApi.RetrievePackageType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageTypeId** | **int64** | ID of package-type to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePackageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PackageTypeModel**](PackageTypeModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPackageType

> ApiResponse SetPackageType(ctx).Body(body).Execute()

Set (create or update) an PackageType



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
    body := *openapiclient.NewPackageTypeModel() // PackageTypeModel | PackageType model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageTypeApi.SetPackageType(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageTypeApi.SetPackageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPackageType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageTypeApi.SetPackageType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPackageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PackageTypeModel**](PackageTypeModel.md) | PackageType model | 

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


## UpdatePackageType

> ApiResponse UpdatePackageType(ctx, packageTypeId).Body(body).Execute()

Update a package-type



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
    packageTypeId := int64(789) // int64 | ID of package-type to update
    body := *openapiclient.NewPackageTypeModel() // PackageTypeModel | PackageType object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageTypeApi.UpdatePackageType(context.Background(), packageTypeId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageTypeApi.UpdatePackageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePackageType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageTypeApi.UpdatePackageType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageTypeId** | **int64** | ID of package-type to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePackageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PackageTypeModel**](PackageTypeModel.md) | PackageType object that needs to be updated | 

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

