# \UomApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveListUom**](UomApi.md#RetrieveListUom) | **Put** /uom | Retrieve List of Uom&#39;s



## RetrieveListUom

> UomListResponse RetrieveListUom(ctx).Arguments(arguments).Execute()

Retrieve List of Uom's



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
    arguments := *openapiclient.NewUomRetrieveListArguments() // UomRetrieveListArguments | Uom RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UomApi.RetrieveListUom(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UomApi.RetrieveListUom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListUom`: UomListResponse
    fmt.Fprintf(os.Stdout, "Response from `UomApi.RetrieveListUom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListUomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**UomRetrieveListArguments**](UomRetrieveListArguments.md) | Uom RetrieveList Arguments | 

### Return type

[**UomListResponse**](UomListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

