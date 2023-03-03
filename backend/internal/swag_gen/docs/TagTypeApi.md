# \TagTypeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTagType**](TagTypeApi.md#CreateTagType) | **Post** /tag-type | Add a new Tag type
[**DeleteTagType**](TagTypeApi.md#DeleteTagType) | **Delete** /tag-type/{tagTypeId} | Delete a Tag type
[**RetrieveListTagType**](TagTypeApi.md#RetrieveListTagType) | **Put** /tag-type | Retrieve List of Tag types
[**RetrieveTagType**](TagTypeApi.md#RetrieveTagType) | **Get** /tag-type/{tagTypeId} | Retrieve a Tag type
[**SetTagType**](TagTypeApi.md#SetTagType) | **Post** /tag-type/set | Set (create or update) Tag type
[**UpdateTagType**](TagTypeApi.md#UpdateTagType) | **Put** /tag-type/{tagTypeId} | Update a Tag type



## CreateTagType

> ApiResponse CreateTagType(ctx).Body(body).Execute()

Add a new Tag type



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
    body := *openapiclient.NewTagTypeModel() // TagTypeModel | Tag type object that needs to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagTypeApi.CreateTagType(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TagTypeApi.CreateTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TagTypeApi.CreateTagType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TagTypeModel**](TagTypeModel.md) | Tag type object that needs to be created | 

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


## DeleteTagType

> ApiResponse DeleteTagType(ctx, tagTypeId).Execute()

Delete a Tag type



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
    tagTypeId := int64(789) // int64 | ID of tag type to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagTypeApi.DeleteTagType(context.Background(), tagTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TagTypeApi.DeleteTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTagType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TagTypeApi.DeleteTagType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagTypeId** | **int64** | ID of tag type to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagTypeRequest struct via the builder pattern


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


## RetrieveListTagType

> TagTypeRetrieveResponse RetrieveListTagType(ctx).Arguments(arguments).Execute()

Retrieve List of Tag types



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
    arguments := *openapiclient.NewTagTypeRetrieveListArguments() // TagTypeRetrieveListArguments | Tag types RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagTypeApi.RetrieveListTagType(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TagTypeApi.RetrieveListTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListTagType`: TagTypeRetrieveResponse
    fmt.Fprintf(os.Stdout, "Response from `TagTypeApi.RetrieveListTagType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**TagTypeRetrieveListArguments**](TagTypeRetrieveListArguments.md) | Tag types RetrieveList Arguments | 

### Return type

[**TagTypeRetrieveResponse**](TagTypeRetrieveResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveTagType

> TagTypeModel RetrieveTagType(ctx, tagTypeId).IncludeObjectTypes(includeObjectTypes).Execute()

Retrieve a Tag type



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
    tagTypeId := int64(789) // int64 | ID of tag type to retrieve
    includeObjectTypes := true // bool | Show the text value of the status (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagTypeApi.RetrieveTagType(context.Background(), tagTypeId).IncludeObjectTypes(includeObjectTypes).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TagTypeApi.RetrieveTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveTagType`: TagTypeModel
    fmt.Fprintf(os.Stdout, "Response from `TagTypeApi.RetrieveTagType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagTypeId** | **int64** | ID of tag type to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeObjectTypes** | **bool** | Show the text value of the status | 

### Return type

[**TagTypeModel**](TagTypeModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTagType

> ApiResponse SetTagType(ctx).Body(body).Execute()

Set (create or update) Tag type



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
    body := *openapiclient.NewTagTypeModel() // TagTypeModel | tag type data (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagTypeApi.SetTagType(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TagTypeApi.SetTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTagType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TagTypeApi.SetTagType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TagTypeModel**](TagTypeModel.md) | tag type data | 

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


## UpdateTagType

> ApiResponse UpdateTagType(ctx, tagTypeId).Execute()

Update a Tag type



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
    tagTypeId := int64(789) // int64 | ID of tag type to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagTypeApi.UpdateTagType(context.Background(), tagTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TagTypeApi.UpdateTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagType`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TagTypeApi.UpdateTagType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagTypeId** | **int64** | ID of tag type to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagTypeRequest struct via the builder pattern


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

