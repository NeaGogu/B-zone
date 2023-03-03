# \PortalSettingApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveListPortalSetting**](PortalSettingApi.md#RetrieveListPortalSetting) | **Put** /portal-setting | Retrieve List of Portal Settings
[**RetrievePortalSetting**](PortalSettingApi.md#RetrievePortalSetting) | **Get** /portal-setting/{portal-settingId} | Retrieve a PortalSetting
[**SetPortalSetting**](PortalSettingApi.md#SetPortalSetting) | **Post** /portal-setting/set | Set (create or update) a PortalSetting
[**UpdatePortalSetting**](PortalSettingApi.md#UpdatePortalSetting) | **Put** /portal-setting/{portal-settingId} | Update a PortalSetting



## RetrieveListPortalSetting

> PortalSettingListResponse RetrieveListPortalSetting(ctx).Arguments(arguments).Execute()

Retrieve List of Portal Settings



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
    arguments := *openapiclient.NewPortalSettingRetrieveListArguments() // PortalSettingRetrieveListArguments | PortalSetting RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortalSettingApi.RetrieveListPortalSetting(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PortalSettingApi.RetrieveListPortalSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListPortalSetting`: PortalSettingListResponse
    fmt.Fprintf(os.Stdout, "Response from `PortalSettingApi.RetrieveListPortalSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListPortalSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**PortalSettingRetrieveListArguments**](PortalSettingRetrieveListArguments.md) | PortalSetting RetrieveList Arguments | 

### Return type

[**PortalSettingListResponse**](PortalSettingListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePortalSetting

> PortalSettingModel RetrievePortalSetting(ctx, portalSettingId).Execute()

Retrieve a PortalSetting



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
    portalSettingId := int64(789) // int64 | ID of portal-setting to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortalSettingApi.RetrievePortalSetting(context.Background(), portalSettingId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PortalSettingApi.RetrievePortalSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePortalSetting`: PortalSettingModel
    fmt.Fprintf(os.Stdout, "Response from `PortalSettingApi.RetrievePortalSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalSettingId** | **int64** | ID of portal-setting to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePortalSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortalSettingModel**](PortalSettingModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPortalSetting

> ApiResponse SetPortalSetting(ctx).Body(body).Execute()

Set (create or update) a PortalSetting



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
    body := *openapiclient.NewPortalSettingModel() // PortalSettingModel | PortalSetting object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortalSettingApi.SetPortalSetting(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PortalSettingApi.SetPortalSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPortalSetting`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PortalSettingApi.SetPortalSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPortalSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PortalSettingModel**](PortalSettingModel.md) | PortalSetting object | 

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


## UpdatePortalSetting

> ApiResponse UpdatePortalSetting(ctx, portalSettingId).Execute()

Update a PortalSetting



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
    portalSettingId := int64(789) // int64 | ID of portal-setting to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PortalSettingApi.UpdatePortalSetting(context.Background(), portalSettingId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PortalSettingApi.UpdatePortalSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePortalSetting`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PortalSettingApi.UpdatePortalSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalSettingId** | **int64** | ID of portal-setting to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortalSettingRequest struct via the builder pattern


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

