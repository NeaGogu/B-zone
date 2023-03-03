# \ActivityTypeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveActivityType**](ActivityTypeApi.md#RetrieveActivityType) | **Get** /activity-type/{activityTypeId} | Find ActivityType by ID
[**RetrieveListActivityType**](ActivityTypeApi.md#RetrieveListActivityType) | **Put** /activity-type | Retrieve List of ActivityTypes



## RetrieveActivityType

> ActivityTypeModel RetrieveActivityType(ctx, activityTypeId).Execute()

Find ActivityType by ID



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
    activityTypeId := int64(789) // int64 | ID of ActivityType to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityTypeApi.RetrieveActivityType(context.Background(), activityTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityTypeApi.RetrieveActivityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveActivityType`: ActivityTypeModel
    fmt.Fprintf(os.Stdout, "Response from `ActivityTypeApi.RetrieveActivityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityTypeId** | **int64** | ID of ActivityType to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveActivityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivityTypeModel**](ActivityTypeModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListActivityType

> ActivityTypeListResponse RetrieveListActivityType(ctx).Arguments(arguments).Execute()

Retrieve List of ActivityTypes



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
    arguments := *openapiclient.NewActivityTypeRetrieveListArguments() // ActivityTypeRetrieveListArguments | ActivityType RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityTypeApi.RetrieveListActivityType(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityTypeApi.RetrieveListActivityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListActivityType`: ActivityTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityTypeApi.RetrieveListActivityType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListActivityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**ActivityTypeRetrieveListArguments**](ActivityTypeRetrieveListArguments.md) | ActivityType RetrieveList Arguments | 

### Return type

[**ActivityTypeListResponse**](ActivityTypeListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

