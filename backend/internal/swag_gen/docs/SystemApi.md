# \SystemApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemGetConfig**](SystemApi.md#SystemGetConfig) | **Get** /system/get-config | Retrieve System Configuration
[**SystemGetSayWhenConfig**](SystemApi.md#SystemGetSayWhenConfig) | **Get** /system/get-say-when-config | Retrieve SayWhen System Configuration
[**SystemGetVariables**](SystemApi.md#SystemGetVariables) | **Get** /system/get-variables | Retrieve System Variables



## SystemGetConfig

> ConfigModel SystemGetConfig(ctx).Execute()

Retrieve System Configuration



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
    resp, r, err := api_client.SystemApi.SystemGetConfig(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.SystemGetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemGetConfig`: ConfigModel
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.SystemGetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemGetConfigRequest struct via the builder pattern


### Return type

[**ConfigModel**](ConfigModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemGetSayWhenConfig

> SayWhenConfigModel SystemGetSayWhenConfig(ctx).Execute()

Retrieve SayWhen System Configuration



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
    resp, r, err := api_client.SystemApi.SystemGetSayWhenConfig(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.SystemGetSayWhenConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemGetSayWhenConfig`: SayWhenConfigModel
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.SystemGetSayWhenConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemGetSayWhenConfigRequest struct via the builder pattern


### Return type

[**SayWhenConfigModel**](SayWhenConfigModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemGetVariables

> VariablesModel SystemGetVariables(ctx).Execute()

Retrieve System Variables



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
    resp, r, err := api_client.SystemApi.SystemGetVariables(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.SystemGetVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemGetVariables`: VariablesModel
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.SystemGetVariables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemGetVariablesRequest struct via the builder pattern


### Return type

[**VariablesModel**](VariablesModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

