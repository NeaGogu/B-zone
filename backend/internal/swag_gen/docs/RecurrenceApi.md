# \RecurrenceApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActivityRecurrence**](RecurrenceApi.md#CreateActivityRecurrence) | **Post** /recurrence/create-activity-recurrence | create a activity recurrence
[**CreateRouteRecurrence**](RecurrenceApi.md#CreateRouteRecurrence) | **Post** /recurrence/create-route-recurrence | create a route recurrence
[**DeleteRecurrenceObject**](RecurrenceApi.md#DeleteRecurrenceObject) | **Delete** /recurrence/delete-recurrence | Delete a Recurrence
[**Finish**](RecurrenceApi.md#Finish) | **Post** /recurrence/finish | Cleans up after the process run
[**GetRuns**](RecurrenceApi.md#GetRuns) | **Post** /recurrence/get-runs | Returns the given runs for the next recurrences!
[**ProcessRuns**](RecurrenceApi.md#ProcessRuns) | **Post** /recurrence/process-runs | Executes the the processes for the ids retrieved with get-runs
[**RetrieveListRecurrence**](RecurrenceApi.md#RetrieveListRecurrence) | **Put** /recurrence | Retrieve List of Recurrences
[**RetrieveRecurrence**](RecurrenceApi.md#RetrieveRecurrence) | **Get** /recurrence/{recurrenceId} | Retrieve a Recurrence
[**UpdateRecurrence**](RecurrenceApi.md#UpdateRecurrence) | **Put** /recurrence/{recurrenceId} | Update a recurrence



## CreateActivityRecurrence

> ApiResponse CreateActivityRecurrence(ctx).Body(body).Execute()

create a activity recurrence



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
    body := *openapiclient.NewRecurrenceModel(int32(333)) // RecurrenceModel | Recurrence object that needs to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.CreateActivityRecurrence(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.CreateActivityRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateActivityRecurrence`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.CreateActivityRecurrence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActivityRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RecurrenceModel**](RecurrenceModel.md) | Recurrence object that needs to be created | 

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


## CreateRouteRecurrence

> ApiResponse CreateRouteRecurrence(ctx).Body(body).Execute()

create a route recurrence



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
    body := *openapiclient.NewRecurrenceModel(int32(333)) // RecurrenceModel | Recurrence object that needs to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.CreateRouteRecurrence(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.CreateRouteRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRouteRecurrence`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.CreateRouteRecurrence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RecurrenceModel**](RecurrenceModel.md) | Recurrence object that needs to be created | 

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


## DeleteRecurrenceObject

> ApiResponse DeleteRecurrenceObject(ctx).Body(body).Execute()

Delete a Recurrence



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
    body := *openapiclient.NewRecurrenceDeleteModel() // RecurrenceDeleteModel | Delete recurrence model

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.DeleteRecurrenceObject(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.DeleteRecurrenceObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRecurrenceObject`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.DeleteRecurrenceObject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecurrenceObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RecurrenceDeleteModel**](RecurrenceDeleteModel.md) | Delete recurrence model | 

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


## Finish

> RecurrenceFinishResponse Finish(ctx).Execute()

Cleans up after the process run

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.Finish(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.Finish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Finish`: RecurrenceFinishResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.Finish`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFinishRequest struct via the builder pattern


### Return type

[**RecurrenceFinishResponse**](RecurrenceFinishResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuns

> RecurrenceGetRunsResponse GetRuns(ctx).Arguments(arguments).Execute()

Returns the given runs for the next recurrences!

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
    arguments := *openapiclient.NewRecurrenceGetRunsArguments() // RecurrenceGetRunsArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.GetRuns(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.GetRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuns`: RecurrenceGetRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.GetRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**RecurrenceGetRunsArguments**](RecurrenceGetRunsArguments.md) | Request Arguments | 

### Return type

[**RecurrenceGetRunsResponse**](RecurrenceGetRunsResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessRuns

> RecurrenceProcessRunsResponse ProcessRuns(ctx).Arguments(arguments).Execute()

Executes the the processes for the ids retrieved with get-runs

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
    arguments := *openapiclient.NewRecurrenceProcessRunsArguments() // RecurrenceProcessRunsArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.ProcessRuns(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.ProcessRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessRuns`: RecurrenceProcessRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.ProcessRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**RecurrenceProcessRunsArguments**](RecurrenceProcessRunsArguments.md) | Request Arguments | 

### Return type

[**RecurrenceProcessRunsResponse**](RecurrenceProcessRunsResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListRecurrence

> RecurrenceListResponse RetrieveListRecurrence(ctx).Arguments(arguments).Execute()

Retrieve List of Recurrences



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
    arguments := *openapiclient.NewRecurrenceRetrieveListArguments() // RecurrenceRetrieveListArguments | Recurrence RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.RetrieveListRecurrence(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.RetrieveListRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListRecurrence`: RecurrenceListResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.RetrieveListRecurrence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**RecurrenceRetrieveListArguments**](RecurrenceRetrieveListArguments.md) | Recurrence RetrieveList Arguments | 

### Return type

[**RecurrenceListResponse**](RecurrenceListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRecurrence

> RecurrenceModel RetrieveRecurrence(ctx, recurrenceId).Execute()

Retrieve a Recurrence



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
    recurrenceId := int64(789) // int64 | ID of recurrence to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.RetrieveRecurrence(context.Background(), recurrenceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.RetrieveRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRecurrence`: RecurrenceModel
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.RetrieveRecurrence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurrenceId** | **int64** | ID of recurrence to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecurrenceModel**](RecurrenceModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecurrence

> ApiResponse UpdateRecurrence(ctx, recurrenceId).Body(body).Execute()

Update a recurrence



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
    recurrenceId := int64(789) // int64 | ID of recurrence to update
    body := *openapiclient.NewRecurrenceModel(int32(333)) // RecurrenceModel | Recurrence object that needs to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrenceApi.UpdateRecurrence(context.Background(), recurrenceId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrenceApi.UpdateRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecurrence`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurrenceApi.UpdateRecurrence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurrenceId** | **int64** | ID of recurrence to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RecurrenceModel**](RecurrenceModel.md) | Recurrence object that needs to be updated | 

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

