# \NotificationApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotification**](NotificationApi.md#CreateNotification) | **Post** /notification | Add a new Notification
[**DeleteNotification**](NotificationApi.md#DeleteNotification) | **Delete** /notification/{notificationId} | Delete an Notification entry
[**RetrieveListNotification**](NotificationApi.md#RetrieveListNotification) | **Put** /notification | Retrieve List of Notification
[**RetrieveNotification**](NotificationApi.md#RetrieveNotification) | **Get** /notification/{notificationId} | Retrieve a Notification
[**SetNotification**](NotificationApi.md#SetNotification) | **Post** /notification/set | Set (create or update) a notification
[**UpdateNotification**](NotificationApi.md#UpdateNotification) | **Put** /notification/{notificationId} | Update a specific Notification object



## CreateNotification

> ApiResponse20 CreateNotification(ctx).Body(body).Execute()

Add a new Notification



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
    body := *openapiclient.NewNotificationModel() // NotificationModel | Notification object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.CreateNotification(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.CreateNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotification`: ApiResponse20
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.CreateNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationModel**](NotificationModel.md) | Notification object that needs to be created | 

### Return type

[**ApiResponse20**](ApiResponse_20.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotification

> ApiResponse17 DeleteNotification(ctx, notificationId).Execute()

Delete an Notification entry



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
    notificationId := int64(789) // int64 | ID of Notification to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.DeleteNotification(context.Background(), notificationId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.DeleteNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNotification`: ApiResponse17
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.DeleteNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **int64** | ID of Notification to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse17**](ApiResponse_17.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListNotification

> NotificationListResponse RetrieveListNotification(ctx).Arguments(arguments).Execute()

Retrieve List of Notification



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
    arguments := *openapiclient.NewNotificationRetrieveListArguments() // NotificationRetrieveListArguments | Notification RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.RetrieveListNotification(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.RetrieveListNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListNotification`: NotificationListResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.RetrieveListNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**NotificationRetrieveListArguments**](NotificationRetrieveListArguments.md) | Notification RetrieveList Arguments | 

### Return type

[**NotificationListResponse**](NotificationListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNotification

> NotificationModel RetrieveNotification(ctx, notificationId).IncludeCategoryTypeName(includeCategoryTypeName).IncludeRecordInfo(includeRecordInfo).Execute()

Retrieve a Notification



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
    notificationId := int64(789) // int64 | ID of Notification to retrieve
    includeCategoryTypeName := true // bool | Show the name of the category type (optional) (default to false)
    includeRecordInfo := true // bool | Show the record info (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.RetrieveNotification(context.Background(), notificationId).IncludeCategoryTypeName(includeCategoryTypeName).IncludeRecordInfo(includeRecordInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.RetrieveNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveNotification`: NotificationModel
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.RetrieveNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **int64** | ID of Notification to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCategoryTypeName** | **bool** | Show the name of the category type | [default to false]
 **includeRecordInfo** | **bool** | Show the record info | [default to false]

### Return type

[**NotificationModel**](NotificationModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNotification

> ApiResponse SetNotification(ctx).Body(body).Execute()

Set (create or update) a notification



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
    body := *openapiclient.NewNotificationModel() // NotificationModel | Notification object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.SetNotification(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.SetNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetNotification`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.SetNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationModel**](NotificationModel.md) | Notification object | 

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


## UpdateNotification

> ApiResponse16 UpdateNotification(ctx, notificationId).Body(body).Execute()

Update a specific Notification object



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
    notificationId := int64(789) // int64 | ID of the Notification object to update
    body := *openapiclient.NewNotificationModel() // NotificationModel | Notification object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.UpdateNotification(context.Background(), notificationId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.UpdateNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotification`: ApiResponse16
    fmt.Fprintf(os.Stdout, "Response from `NotificationApi.UpdateNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **int64** | ID of the Notification object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NotificationModel**](NotificationModel.md) | Notification object that needs to be updated | 

### Return type

[**ApiResponse16**](ApiResponse_16.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

