# \AssignmentApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAssignment**](AssignmentApi.md#DeleteAssignment) | **Delete** /assignment/{assignmentId} | Delete an assignment
[**RetrieveAssignment**](AssignmentApi.md#RetrieveAssignment) | **Get** /assignment/{assignmentId} | Find assignment by ID
[**RetrieveListAssignment**](AssignmentApi.md#RetrieveListAssignment) | **Put** /assignment | Retrieve List of Assignments
[**SetAssignment**](AssignmentApi.md#SetAssignment) | **Post** /assignment/set | Set (create or update) an Assignment
[**UpdateAssignment**](AssignmentApi.md#UpdateAssignment) | **Put** /assignment/{assignmentId} | Update a assignment



## DeleteAssignment

> ApiResponse DeleteAssignment(ctx, assignmentId).Execute()

Delete an assignment



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
    assignmentId := int64(789) // int64 | ID of the assignment to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssignmentApi.DeleteAssignment(context.Background(), assignmentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentApi.DeleteAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAssignment`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AssignmentApi.DeleteAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assignmentId** | **int64** | ID of the assignment to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssignmentRequest struct via the builder pattern


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


## RetrieveAssignment

> AssignmentModel RetrieveAssignment(ctx, assignmentId).IncludeActivities(includeActivities).IncludeMetaData(includeMetaData).IncludeLinks(includeLinks).IncludeFiles(includeFiles).IncludeTagIds(includeTagIds).IncludeTagNames(includeTagNames).IncludeBookingAccount(includeBookingAccount).IncludeRecordInfo(includeRecordInfo).IncludeRecordObject(includeRecordObject).Execute()

Find assignment by ID



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
    assignmentId := int64(789) // int64 | ID of assignment to return
    includeActivities := true // bool | Include activities belonging to assignment (optional) (default to false)
    includeMetaData := true // bool | Include meta data (optional) (default to false)
    includeLinks := true // bool | Include links (optional) (default to false)
    includeFiles := true // bool | Include files (optional) (default to false)
    includeTagIds := true // bool | Include tag ids (optional) (default to false)
    includeTagNames := true // bool | Include tag names (optional) (default to false)
    includeBookingAccount := true // bool | Include booking account (optional) (default to false)
    includeRecordInfo := true // bool | Include record info (optional) (default to false)
    includeRecordObject := true // bool | Include record object (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssignmentApi.RetrieveAssignment(context.Background(), assignmentId).IncludeActivities(includeActivities).IncludeMetaData(includeMetaData).IncludeLinks(includeLinks).IncludeFiles(includeFiles).IncludeTagIds(includeTagIds).IncludeTagNames(includeTagNames).IncludeBookingAccount(includeBookingAccount).IncludeRecordInfo(includeRecordInfo).IncludeRecordObject(includeRecordObject).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentApi.RetrieveAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAssignment`: AssignmentModel
    fmt.Fprintf(os.Stdout, "Response from `AssignmentApi.RetrieveAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assignmentId** | **int64** | ID of assignment to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeActivities** | **bool** | Include activities belonging to assignment | [default to false]
 **includeMetaData** | **bool** | Include meta data | [default to false]
 **includeLinks** | **bool** | Include links | [default to false]
 **includeFiles** | **bool** | Include files | [default to false]
 **includeTagIds** | **bool** | Include tag ids | [default to false]
 **includeTagNames** | **bool** | Include tag names | [default to false]
 **includeBookingAccount** | **bool** | Include booking account | [default to false]
 **includeRecordInfo** | **bool** | Include record info | [default to false]
 **includeRecordObject** | **bool** | Include record object | [default to false]

### Return type

[**AssignmentModel**](AssignmentModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListAssignment

> AssignmentListResponse RetrieveListAssignment(ctx).Arguments(arguments).Execute()

Retrieve List of Assignments



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
    arguments := *openapiclient.NewAssignmentRetrieveListArguments() // AssignmentRetrieveListArguments | Assignment RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssignmentApi.RetrieveListAssignment(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentApi.RetrieveListAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListAssignment`: AssignmentListResponse
    fmt.Fprintf(os.Stdout, "Response from `AssignmentApi.RetrieveListAssignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**AssignmentRetrieveListArguments**](AssignmentRetrieveListArguments.md) | Assignment RetrieveList Arguments | 

### Return type

[**AssignmentListResponse**](AssignmentListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAssignment

> ApiResponse SetAssignment(ctx).Body(body).Execute()

Set (create or update) an Assignment



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
    body := *openapiclient.NewAssignmentModel() // AssignmentModel | Assignment object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssignmentApi.SetAssignment(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentApi.SetAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAssignment`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AssignmentApi.SetAssignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AssignmentModel**](AssignmentModel.md) | Assignment object | 

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


## UpdateAssignment

> ApiResponse UpdateAssignment(ctx, assignmentId).Body(body).Execute()

Update a assignment



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
    assignmentId := int64(789) // int64 | ID of assignment to update
    body := *openapiclient.NewAssignmentModel() // AssignmentModel | Assignment object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssignmentApi.UpdateAssignment(context.Background(), assignmentId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentApi.UpdateAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssignment`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AssignmentApi.UpdateAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assignmentId** | **int64** | ID of assignment to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AssignmentModel**](AssignmentModel.md) | Assignment object that needs to be updated | 

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

