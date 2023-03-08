# \ManualWebHookApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManualWebHookTrigger**](ManualWebHookApi.md#ManualWebHookTrigger) | **Post** /manual-web-hook/trigger | Trigger a webhook



## ManualWebHookTrigger

> ApiResponse ManualWebHookTrigger(ctx).Body(body).Execute()

Trigger a webhook



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
    body := *openapiclient.NewManualWebHookModel() // ManualWebHookModel | WebHook that has to be triggered manually (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManualWebHookApi.ManualWebHookTrigger(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualWebHookApi.ManualWebHookTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManualWebHookTrigger`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ManualWebHookApi.ManualWebHookTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManualWebHookTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ManualWebHookModel**](ManualWebHookModel.md) | WebHook that has to be triggered manually | 

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

