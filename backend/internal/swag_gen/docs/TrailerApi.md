# \TrailerApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrailer**](TrailerApi.md#CreateTrailer) | **Post** /trailer | Add a trailer
[**DeleteTrailer**](TrailerApi.md#DeleteTrailer) | **Delete** /trailer/{trailerId} | Delete an trailer
[**RetrieveListTrailer**](TrailerApi.md#RetrieveListTrailer) | **Put** /trailer | Retrieve List of Trailers
[**RetrieveTrailer**](TrailerApi.md#RetrieveTrailer) | **Get** /trailer/{trailerId} | Find trailer by ID
[**SetTrailer**](TrailerApi.md#SetTrailer) | **Post** /trailer/set | Set (create or update) a trailer
[**UpdateTrailer**](TrailerApi.md#UpdateTrailer) | **Put** /trailer/{trailerId} | Update a trailer



## CreateTrailer

> ApiResponse CreateTrailer(ctx).Body(body).Execute()

Add a trailer



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
    body := *openapiclient.NewTrailerModel() // TrailerModel | Trailer object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrailerApi.CreateTrailer(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TrailerApi.CreateTrailer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrailer`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TrailerApi.CreateTrailer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrailerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TrailerModel**](TrailerModel.md) | Trailer object that needs to be created | 

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


## DeleteTrailer

> ApiResponse DeleteTrailer(ctx, trailerId).Execute()

Delete an trailer



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
    trailerId := int64(789) // int64 | ID of the trailer to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrailerApi.DeleteTrailer(context.Background(), trailerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TrailerApi.DeleteTrailer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTrailer`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TrailerApi.DeleteTrailer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailerId** | **int64** | ID of the trailer to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrailerRequest struct via the builder pattern


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


## RetrieveListTrailer

> TrailerListResponse RetrieveListTrailer(ctx).Arguments(arguments).Execute()

Retrieve List of Trailers



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
    arguments := *openapiclient.NewTrailerRetrieveListArguments() // TrailerRetrieveListArguments | Trailer RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrailerApi.RetrieveListTrailer(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TrailerApi.RetrieveListTrailer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListTrailer`: TrailerListResponse
    fmt.Fprintf(os.Stdout, "Response from `TrailerApi.RetrieveListTrailer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListTrailerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**TrailerRetrieveListArguments**](TrailerRetrieveListArguments.md) | Trailer RetrieveList Arguments | 

### Return type

[**TrailerListResponse**](TrailerListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveTrailer

> TrailerModel RetrieveTrailer(ctx, trailerId).IncludeTrailerTags(includeTrailerTags).IncludeUpdatedBy(includeUpdatedBy).Execute()

Find trailer by ID



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
    trailerId := int64(789) // int64 | ID of trailer to return
    includeTrailerTags := true // bool | a list of tags bound to trailer (default to true)
    includeUpdatedBy := true // bool | include updated_by_name (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrailerApi.RetrieveTrailer(context.Background(), trailerId).IncludeTrailerTags(includeTrailerTags).IncludeUpdatedBy(includeUpdatedBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TrailerApi.RetrieveTrailer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveTrailer`: TrailerModel
    fmt.Fprintf(os.Stdout, "Response from `TrailerApi.RetrieveTrailer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailerId** | **int64** | ID of trailer to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTrailerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTrailerTags** | **bool** | a list of tags bound to trailer | [default to true]
 **includeUpdatedBy** | **bool** | include updated_by_name | [default to true]

### Return type

[**TrailerModel**](TrailerModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTrailer

> ApiResponse SetTrailer(ctx).Body(body).Execute()

Set (create or update) a trailer



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
    body := *openapiclient.NewTrailerModel() // TrailerModel | Trailer object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrailerApi.SetTrailer(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TrailerApi.SetTrailer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTrailer`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TrailerApi.SetTrailer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTrailerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TrailerModel**](TrailerModel.md) | Trailer object | 

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


## UpdateTrailer

> ApiResponse UpdateTrailer(ctx, trailerId).Body(body).Execute()

Update a trailer



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
    trailerId := int64(789) // int64 | ID of trailer to update
    body := *openapiclient.NewTrailerModel() // TrailerModel | Trailer object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrailerApi.UpdateTrailer(context.Background(), trailerId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TrailerApi.UpdateTrailer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrailer`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TrailerApi.UpdateTrailer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailerId** | **int64** | ID of trailer to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrailerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TrailerModel**](TrailerModel.md) | Trailer object that needs to be updated | 

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

