# \NoteApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNote**](NoteApi.md#DeleteNote) | **Delete** /note/{noteId} | Delete an note
[**RetrieveListNote**](NoteApi.md#RetrieveListNote) | **Put** /note | Retrieve List of Notes
[**RetrieveNote**](NoteApi.md#RetrieveNote) | **Get** /note/{noteId} | Find note by ID
[**SetNote**](NoteApi.md#SetNote) | **Post** /note/set | Set (create or update) a note



## DeleteNote

> ApiResponse DeleteNote(ctx, noteId).Execute()

Delete an note



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
    noteId := int64(789) // int64 | ID of the note to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteApi.DeleteNote(context.Background(), noteId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteApi.DeleteNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNote`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `NoteApi.DeleteNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **int64** | ID of the note to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNoteRequest struct via the builder pattern


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


## RetrieveListNote

> []NoteModel RetrieveListNote(ctx).Arguments(arguments).Execute()

Retrieve List of Notes



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
    arguments := *openapiclient.NewNoteRetrieveListArguments() // NoteRetrieveListArguments | Note RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteApi.RetrieveListNote(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteApi.RetrieveListNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListNote`: []NoteModel
    fmt.Fprintf(os.Stdout, "Response from `NoteApi.RetrieveListNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**NoteRetrieveListArguments**](NoteRetrieveListArguments.md) | Note RetrieveList Arguments | 

### Return type

[**[]NoteModel**](NoteModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNote

> NoteModel RetrieveNote(ctx, noteId).IncludeNoteTags(includeNoteTags).IncludeNoteTagTypeLinkIds(includeNoteTagTypeLinkIds).IncludeNoteObjectLinkIds(includeNoteObjectLinkIds).IncludeUpdatedBy(includeUpdatedBy).Execute()

Find note by ID



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
    noteId := int64(789) // int64 | ID of note to return
    includeNoteTags := true // bool | a list of tags bound to note (default to true)
    includeNoteTagTypeLinkIds := true // bool | link ids of the tag types (default to true)
    includeNoteObjectLinkIds := true // bool | Include teh link ids bound to teh object where teh note belongs to (default to true)
    includeUpdatedBy := true // bool | include updated_by_name (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteApi.RetrieveNote(context.Background(), noteId).IncludeNoteTags(includeNoteTags).IncludeNoteTagTypeLinkIds(includeNoteTagTypeLinkIds).IncludeNoteObjectLinkIds(includeNoteObjectLinkIds).IncludeUpdatedBy(includeUpdatedBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteApi.RetrieveNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveNote`: NoteModel
    fmt.Fprintf(os.Stdout, "Response from `NoteApi.RetrieveNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **int64** | ID of note to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeNoteTags** | **bool** | a list of tags bound to note | [default to true]
 **includeNoteTagTypeLinkIds** | **bool** | link ids of the tag types | [default to true]
 **includeNoteObjectLinkIds** | **bool** | Include teh link ids bound to teh object where teh note belongs to | [default to true]
 **includeUpdatedBy** | **bool** | include updated_by_name | [default to true]

### Return type

[**NoteModel**](NoteModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNote

> ApiResponse SetNote(ctx).Body(body).Execute()

Set (create or update) a note



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
    body := *openapiclient.NewNoteModel() // NoteModel | Note object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NoteApi.SetNote(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NoteApi.SetNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetNote`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `NoteApi.SetNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NoteModel**](NoteModel.md) | Note object | 

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

