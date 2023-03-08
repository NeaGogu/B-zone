# \CommunicationMessageTypesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCommunicationMessageTypes**](CommunicationMessageTypesApi.md#ListCommunicationMessageTypes) | **Put** /communication-message-type/get-list | List all CommunicationMessageTypes



## ListCommunicationMessageTypes

> CommunicationMessageTypeListResponse ListCommunicationMessageTypes(ctx).Execute()

List all CommunicationMessageTypes



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
    resp, r, err := api_client.CommunicationMessageTypesApi.ListCommunicationMessageTypes(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationMessageTypesApi.ListCommunicationMessageTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommunicationMessageTypes`: CommunicationMessageTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationMessageTypesApi.ListCommunicationMessageTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCommunicationMessageTypesRequest struct via the builder pattern


### Return type

[**CommunicationMessageTypeListResponse**](CommunicationMessageTypeListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

