# \NotificationCategoryApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationCategory**](NotificationCategoryApi.md#CreateNotificationCategory) | **Post** /notification-category | Add a new NotificationCategory
[**DeleteNotificationCategory**](NotificationCategoryApi.md#DeleteNotificationCategory) | **Delete** /notification-category/{notification-categoryId} | Delete an NotificationCategory entry
[**RetrieveListNotificationCategory**](NotificationCategoryApi.md#RetrieveListNotificationCategory) | **Put** /notification-category | Retrieve List of NotificationCategory
[**RetrieveNotificationCategory**](NotificationCategoryApi.md#RetrieveNotificationCategory) | **Get** /notification-category/{notification-categoryId} | Retrieve a NotificationCategory
[**SetNotificationCategory**](NotificationCategoryApi.md#SetNotificationCategory) | **Post** /notification-category/set | Create a new NotificationCategory or update an existing one
[**UpdateNotificationCategory**](NotificationCategoryApi.md#UpdateNotificationCategory) | **Put** /notification-category/{notification-categoryId} | Update a specific NotificationCategory object



## CreateNotificationCategory

> ApiResponse11 CreateNotificationCategory(ctx).Body(body).Execute()

Add a new NotificationCategory



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
    body := *openapiclient.NewNotificationCategoryModel() // NotificationCategoryModel | NotificationCategory object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationCategoryApi.CreateNotificationCategory(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationCategoryApi.CreateNotificationCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationCategory`: ApiResponse11
    fmt.Fprintf(os.Stdout, "Response from `NotificationCategoryApi.CreateNotificationCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationCategoryModel**](NotificationCategoryModel.md) | NotificationCategory object that needs to be created | 

### Return type

[**ApiResponse11**](ApiResponse_11.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationCategory

> ApiResponse9 DeleteNotificationCategory(ctx, notificationCategoryId).Execute()

Delete an NotificationCategory entry



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
    notificationCategoryId := int64(789) // int64 | ID of NotificationCategory to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationCategoryApi.DeleteNotificationCategory(context.Background(), notificationCategoryId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationCategoryApi.DeleteNotificationCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNotificationCategory`: ApiResponse9
    fmt.Fprintf(os.Stdout, "Response from `NotificationCategoryApi.DeleteNotificationCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationCategoryId** | **int64** | ID of NotificationCategory to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse9**](ApiResponse_9.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListNotificationCategory

> NotificationCategoryListResponse RetrieveListNotificationCategory(ctx).Arguments(arguments).Execute()

Retrieve List of NotificationCategory



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
    arguments := *openapiclient.NewNotificationCategoryRetrieveListArguments() // NotificationCategoryRetrieveListArguments | NotificationCategory RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationCategoryApi.RetrieveListNotificationCategory(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationCategoryApi.RetrieveListNotificationCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListNotificationCategory`: NotificationCategoryListResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationCategoryApi.RetrieveListNotificationCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListNotificationCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**NotificationCategoryRetrieveListArguments**](NotificationCategoryRetrieveListArguments.md) | NotificationCategory RetrieveList Arguments | 

### Return type

[**NotificationCategoryListResponse**](NotificationCategoryListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNotificationCategory

> NotificationCategoryModel RetrieveNotificationCategory(ctx, notificationCategoryId).IncludeObjectTypeName(includeObjectTypeName).IncludeRecordInfo(includeRecordInfo).Execute()

Retrieve a NotificationCategory



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
    notificationCategoryId := int64(789) // int64 | ID of NotificationCategory to retrieve
    includeObjectTypeName := true // bool | Show the name of the object type (optional) (default to false)
    includeRecordInfo := true // bool | Show the record info (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationCategoryApi.RetrieveNotificationCategory(context.Background(), notificationCategoryId).IncludeObjectTypeName(includeObjectTypeName).IncludeRecordInfo(includeRecordInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationCategoryApi.RetrieveNotificationCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveNotificationCategory`: NotificationCategoryModel
    fmt.Fprintf(os.Stdout, "Response from `NotificationCategoryApi.RetrieveNotificationCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationCategoryId** | **int64** | ID of NotificationCategory to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNotificationCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeObjectTypeName** | **bool** | Show the name of the object type | [default to false]
 **includeRecordInfo** | **bool** | Show the record info | [default to false]

### Return type

[**NotificationCategoryModel**](NotificationCategoryModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNotificationCategory

> ApiResponse12 SetNotificationCategory(ctx).Body(body).Execute()

Create a new NotificationCategory or update an existing one



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
    body := *openapiclient.NewNotificationCategoryModel() // NotificationCategoryModel | NotificationCategory object that needs to be set (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationCategoryApi.SetNotificationCategory(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationCategoryApi.SetNotificationCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetNotificationCategory`: ApiResponse12
    fmt.Fprintf(os.Stdout, "Response from `NotificationCategoryApi.SetNotificationCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNotificationCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationCategoryModel**](NotificationCategoryModel.md) | NotificationCategory object that needs to be set | 

### Return type

[**ApiResponse12**](ApiResponse_12.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationCategory

> ApiResponse8 UpdateNotificationCategory(ctx, notificationCategoryId).Body(body).Execute()

Update a specific NotificationCategory object



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
    notificationCategoryId := int64(789) // int64 | ID of the NotificationCategory object to update
    body := *openapiclient.NewNotificationCategoryModel() // NotificationCategoryModel | NotificationCategory object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationCategoryApi.UpdateNotificationCategory(context.Background(), notificationCategoryId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationCategoryApi.UpdateNotificationCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationCategory`: ApiResponse8
    fmt.Fprintf(os.Stdout, "Response from `NotificationCategoryApi.UpdateNotificationCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationCategoryId** | **int64** | ID of the NotificationCategory object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NotificationCategoryModel**](NotificationCategoryModel.md) | NotificationCategory object that needs to be updated | 

### Return type

[**ApiResponse8**](ApiResponse_8.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

