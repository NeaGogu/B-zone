# \NoteCategoryApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNoteCategory**](NoteCategoryApi.md#DeleteNoteCategory) | **Delete** /note-category/{noteCategoryId} | Delete a note-category
[**RetrieveListNoteCategory**](NoteCategoryApi.md#RetrieveListNoteCategory) | **Put** /note-category | Retrieve List of NoteCategories
[**RetrieveNoteCategory**](NoteCategoryApi.md#RetrieveNoteCategory) | **Get** /note-category/{noteCategoryId} | Find note-category by ID
[**SetNoteCategory**](NoteCategoryApi.md#SetNoteCategory) | **Post** /note-category/set | Set (create or update) an NoteCategory



## DeleteNoteCategory

> ApiResponse DeleteNoteCategory(ctx, noteCategoryId).Execute()

Delete a note-category



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
    noteCategoryId := int64(789) // int64 | ID of the note-category to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteCategoryApi.DeleteNoteCategory(context.Background(), noteCategoryId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteCategoryApi.DeleteNoteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNoteCategory`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `NoteCategoryApi.DeleteNoteCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteCategoryId** | **int64** | ID of the note-category to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNoteCategoryRequest struct via the builder pattern


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


## RetrieveListNoteCategory

> NoteCategoryListResponse RetrieveListNoteCategory(ctx).Arguments(arguments).Execute()

Retrieve List of NoteCategories



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
    arguments := *openapiclient.NewNoteCategoryRetrieveListArguments() // NoteCategoryRetrieveListArguments | NoteCategory RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteCategoryApi.RetrieveListNoteCategory(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteCategoryApi.RetrieveListNoteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListNoteCategory`: NoteCategoryListResponse
    fmt.Fprintf(os.Stdout, "Response from `NoteCategoryApi.RetrieveListNoteCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListNoteCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**NoteCategoryRetrieveListArguments**](NoteCategoryRetrieveListArguments.md) | NoteCategory RetrieveList Arguments | 

### Return type

[**NoteCategoryListResponse**](NoteCategoryListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNoteCategory

> NoteCategoryModel RetrieveNoteCategory(ctx, noteCategoryId).Execute()

Find note-category by ID



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
    noteCategoryId := int64(789) // int64 | ID of note-category to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteCategoryApi.RetrieveNoteCategory(context.Background(), noteCategoryId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteCategoryApi.RetrieveNoteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveNoteCategory`: NoteCategoryModel
    fmt.Fprintf(os.Stdout, "Response from `NoteCategoryApi.RetrieveNoteCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteCategoryId** | **int64** | ID of note-category to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNoteCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NoteCategoryModel**](NoteCategoryModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNoteCategory

> ApiResponse SetNoteCategory(ctx).Body(body).Execute()

Set (create or update) an NoteCategory



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
    body := *openapiclient.NewNoteCategoryModel() // NoteCategoryModel | NoteCategory model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteCategoryApi.SetNoteCategory(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteCategoryApi.SetNoteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetNoteCategory`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `NoteCategoryApi.SetNoteCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNoteCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NoteCategoryModel**](NoteCategoryModel.md) | NoteCategory model | 

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

