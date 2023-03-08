# \LinkApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLink**](LinkApi.md#DeleteLink) | **Delete** /link/{linkId} | Delete a link
[**RetrieveListLink**](LinkApi.md#RetrieveListLink) | **Put** /link | Retrieve List of Links
[**UpdateLink**](LinkApi.md#UpdateLink) | **Put** /link/{linkId} | Update a specific link object



## DeleteLink

> ApiResponse DeleteLink(ctx, linkId).Execute()

Delete a link



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
    linkId := int64(789) // int64 | ID of the link to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinkApi.DeleteLink(context.Background(), linkId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkApi.DeleteLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLink`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `LinkApi.DeleteLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **int64** | ID of the link to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkRequest struct via the builder pattern


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


## RetrieveListLink

> []LinkListResponse RetrieveListLink(ctx).Arguments(arguments).Execute()

Retrieve List of Links



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
    arguments := *openapiclient.NewLinkRetrieveListArguments() // LinkRetrieveListArguments | Link RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinkApi.RetrieveListLink(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkApi.RetrieveListLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListLink`: []LinkListResponse
    fmt.Fprintf(os.Stdout, "Response from `LinkApi.RetrieveListLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**LinkRetrieveListArguments**](LinkRetrieveListArguments.md) | Link RetrieveList Arguments | 

### Return type

[**[]LinkListResponse**](LinkListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLink

> ApiResponse UpdateLink(ctx, linkId).Body(body).Execute()

Update a specific link object



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
    linkId := int64(789) // int64 | ID of the link object to update
    body := *openapiclient.NewLinkModel() // LinkModel | Link object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinkApi.UpdateLink(context.Background(), linkId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkApi.UpdateLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLink`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `LinkApi.UpdateLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **int64** | ID of the link object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LinkModel**](LinkModel.md) | Link object that needs to be updated | 

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

