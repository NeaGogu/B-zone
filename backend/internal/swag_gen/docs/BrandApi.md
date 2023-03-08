# \BrandApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrand**](BrandApi.md#CreateBrand) | **Post** /brand | Add a new Brand
[**DeleteBrand**](BrandApi.md#DeleteBrand) | **Delete** /brand/{brandId} | Delete a Brand
[**RetrieveBrand**](BrandApi.md#RetrieveBrand) | **Get** /brand/{brandId} | Retrieve a Brand
[**RetrieveListBrand**](BrandApi.md#RetrieveListBrand) | **Put** /brand | Retrieve List of Brands
[**SetBrand**](BrandApi.md#SetBrand) | **Post** /brand/set | Set (create or update) a Brand
[**UpdateBrand**](BrandApi.md#UpdateBrand) | **Put** /brand/{brandId} | Update a Brand



## CreateBrand

> ApiResponse CreateBrand(ctx).Body(body).Execute()

Add a new Brand



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
    body := *openapiclient.NewBrandModel() // BrandModel | Brand object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandApi.CreateBrand(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandApi.CreateBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrand`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandApi.CreateBrand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BrandModel**](BrandModel.md) | Brand object that needs to be created | 

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


## DeleteBrand

> ApiResponse DeleteBrand(ctx, brandId).Execute()

Delete a Brand



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
    brandId := int64(789) // int64 | ID of brand to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandApi.DeleteBrand(context.Background(), brandId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandApi.DeleteBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBrand`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandApi.DeleteBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int64** | ID of brand to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandRequest struct via the builder pattern


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


## RetrieveBrand

> BrandModel RetrieveBrand(ctx, brandId).Execute()

Retrieve a Brand



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
    brandId := int64(789) // int64 | ID of brand to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandApi.RetrieveBrand(context.Background(), brandId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandApi.RetrieveBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveBrand`: BrandModel
    fmt.Fprintf(os.Stdout, "Response from `BrandApi.RetrieveBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int64** | ID of brand to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrandModel**](BrandModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListBrand

> BrandListResponse RetrieveListBrand(ctx).Arguments(arguments).Execute()

Retrieve List of Brands



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
    arguments := *openapiclient.NewBrandRetrieveListArguments() // BrandRetrieveListArguments | Brand RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandApi.RetrieveListBrand(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandApi.RetrieveListBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListBrand`: BrandListResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandApi.RetrieveListBrand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**BrandRetrieveListArguments**](BrandRetrieveListArguments.md) | Brand RetrieveList Arguments | 

### Return type

[**BrandListResponse**](BrandListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBrand

> ApiResponse SetBrand(ctx).Body(body).Execute()

Set (create or update) a Brand



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
    body := *openapiclient.NewBrandModel() // BrandModel | Brand object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandApi.SetBrand(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandApi.SetBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetBrand`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandApi.SetBrand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BrandModel**](BrandModel.md) | Brand object | 

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


## UpdateBrand

> ApiResponse UpdateBrand(ctx, brandId).Execute()

Update a Brand



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
    brandId := int64(789) // int64 | ID of brand to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandApi.UpdateBrand(context.Background(), brandId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandApi.UpdateBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrand`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandApi.UpdateBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int64** | ID of brand to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandRequest struct via the builder pattern


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

