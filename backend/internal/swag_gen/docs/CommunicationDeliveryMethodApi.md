# \CommunicationDeliveryMethodApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommunicationDeliveryMethod**](CommunicationDeliveryMethodApi.md#CreateCommunicationDeliveryMethod) | **Post** /communication-delivery-method/set | Create a communication delivery method



## CreateCommunicationDeliveryMethod

> ApiResponse CreateCommunicationDeliveryMethod(ctx).Body(body).Execute()

Create a communication delivery method



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
    body := *openapiclient.NewCommunicationDeliveryMethodModel() // CommunicationDeliveryMethodModel | CommunicationDeliveryMethod object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationDeliveryMethodApi.CreateCommunicationDeliveryMethod(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationDeliveryMethodApi.CreateCommunicationDeliveryMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommunicationDeliveryMethod`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationDeliveryMethodApi.CreateCommunicationDeliveryMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommunicationDeliveryMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommunicationDeliveryMethodModel**](CommunicationDeliveryMethodModel.md) | CommunicationDeliveryMethod object | 

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

