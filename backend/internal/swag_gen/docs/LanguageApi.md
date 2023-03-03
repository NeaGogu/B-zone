# \LanguageApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveListLanguage**](LanguageApi.md#RetrieveListLanguage) | **Put** /language | Retrieve List of Language



## RetrieveListLanguage

> LanguageListResponse RetrieveListLanguage(ctx).Arguments(arguments).Execute()

Retrieve List of Language



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
    arguments := *openapiclient.NewLanguageRetrieveListArguments() // LanguageRetrieveListArguments | Language RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LanguageApi.RetrieveListLanguage(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageApi.RetrieveListLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListLanguage`: LanguageListResponse
    fmt.Fprintf(os.Stdout, "Response from `LanguageApi.RetrieveListLanguage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**LanguageRetrieveListArguments**](LanguageRetrieveListArguments.md) | Language RetrieveList Arguments | 

### Return type

[**LanguageListResponse**](LanguageListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

