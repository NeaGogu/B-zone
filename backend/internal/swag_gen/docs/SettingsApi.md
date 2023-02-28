# \SettingsApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConstants**](SettingsApi.md#GetConstants) | **Get** /settings/get-constants | getConstants, returns in structure &#39;string&#39;:&#39;string&#39;, can contain rounding errors for floating points
[**GetConstantsReversed**](SettingsApi.md#GetConstantsReversed) | **Get** /settings/get-constants-reversed | getConstantsReversed, returns in structure &#39;string&#39;:int/float
[**RetrieveListSettings**](SettingsApi.md#RetrieveListSettings) | **Put** /settings | Retrieve List of Settings
[**RetrieveSettings**](SettingsApi.md#RetrieveSettings) | **Get** /settings/{settingsId} | Retrieve a Settings
[**SetSetting**](SettingsApi.md#SetSetting) | **Post** /settings/set | Set (update) Setting value
[**UpdateSettings**](SettingsApi.md#UpdateSettings) | **Put** /settings/{settingsId} | Update a Settings



## GetConstants

> SettingsGetConstantsResponse GetConstants(ctx).Execute()

getConstants, returns in structure 'string':'string', can contain rounding errors for floating points



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
    resp, r, err := api_client.SettingsApi.GetConstants(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetConstants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConstants`: SettingsGetConstantsResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetConstants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConstantsRequest struct via the builder pattern


### Return type

[**SettingsGetConstantsResponse**](SettingsGetConstantsResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConstantsReversed

> SettingsGetConstantsReversedResponse GetConstantsReversed(ctx).Execute()

getConstantsReversed, returns in structure 'string':int/float



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
    resp, r, err := api_client.SettingsApi.GetConstantsReversed(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetConstantsReversed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConstantsReversed`: SettingsGetConstantsReversedResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetConstantsReversed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConstantsReversedRequest struct via the builder pattern


### Return type

[**SettingsGetConstantsReversedResponse**](SettingsGetConstantsReversedResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListSettings

> []SettingsModel RetrieveListSettings(ctx).Arguments(arguments).Execute()

Retrieve List of Settings



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
    arguments := *openapiclient.NewSettingsRetrieveListArguments() // SettingsRetrieveListArguments | Settings RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.RetrieveListSettings(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.RetrieveListSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListSettings`: []SettingsModel
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.RetrieveListSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**SettingsRetrieveListArguments**](SettingsRetrieveListArguments.md) | Settings RetrieveList Arguments | 

### Return type

[**[]SettingsModel**](SettingsModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSettings

> SettingsModel RetrieveSettings(ctx, settingsId).Execute()

Retrieve a Settings



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
    settingsId := int64(789) // int64 | ID of settings to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.RetrieveSettings(context.Background(), settingsId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.RetrieveSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSettings`: SettingsModel
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.RetrieveSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingsId** | **int64** | ID of settings to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsModel**](SettingsModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSetting

> ApiResponse SetSetting(ctx).Body(body).Execute()

Set (update) Setting value



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
    body := *openapiclient.NewSettingsModel() // SettingsModel | Setting object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.SetSetting(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.SetSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSetting`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.SetSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SettingsModel**](SettingsModel.md) | Setting object | 

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


## UpdateSettings

> ApiResponse UpdateSettings(ctx, settingsId).Execute()

Update a Settings



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
    settingsId := int64(789) // int64 | ID of settings to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SettingsApi.UpdateSettings(context.Background(), settingsId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.UpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSettings`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.UpdateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingsId** | **int64** | ID of settings to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


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

