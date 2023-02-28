# \CommunicationTemplateApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommunicationTemplate**](CommunicationTemplateApi.md#CreateCommunicationTemplate) | **Post** /communication-template/set | Create a Communication Template
[**DeleteCommunicationTemplate**](CommunicationTemplateApi.md#DeleteCommunicationTemplate) | **Delete** /communication-template/{communicationTemplateId} | Delete a communicationtemplate
[**RetrieveCommunicationTemplate**](CommunicationTemplateApi.md#RetrieveCommunicationTemplate) | **Get** /communication-template/{communicationTemplateId} | Retrieve a communicationtemplate
[**RetrieveListTemplate**](CommunicationTemplateApi.md#RetrieveListTemplate) | **Put** /communication-template | Retrieve List of Templates



## CreateCommunicationTemplate

> ApiResponse CreateCommunicationTemplate(ctx).Body(body).Execute()

Create a Communication Template



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
    body := *openapiclient.NewCommunicationTemplateModel() // CommunicationTemplateModel | CommunicationTemplate object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationTemplateApi.CreateCommunicationTemplate(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationTemplateApi.CreateCommunicationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommunicationTemplate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationTemplateApi.CreateCommunicationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommunicationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommunicationTemplateModel**](CommunicationTemplateModel.md) | CommunicationTemplate object | 

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


## DeleteCommunicationTemplate

> ApiResponse DeleteCommunicationTemplate(ctx, communicationTemplateId).Execute()

Delete a communicationtemplate



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
    communicationTemplateId := int64(789) // int64 | ID of the communication template to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationTemplateApi.DeleteCommunicationTemplate(context.Background(), communicationTemplateId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationTemplateApi.DeleteCommunicationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCommunicationTemplate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationTemplateApi.DeleteCommunicationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationTemplateId** | **int64** | ID of the communication template to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommunicationTemplateRequest struct via the builder pattern


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


## RetrieveCommunicationTemplate

> CommunicationTemplateModel RetrieveCommunicationTemplate(ctx, communicationTemplateId).Execute()

Retrieve a communicationtemplate



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
    communicationTemplateId := int64(789) // int64 | ID of template to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationTemplateApi.RetrieveCommunicationTemplate(context.Background(), communicationTemplateId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationTemplateApi.RetrieveCommunicationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCommunicationTemplate`: CommunicationTemplateModel
    fmt.Fprintf(os.Stdout, "Response from `CommunicationTemplateApi.RetrieveCommunicationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**communicationTemplateId** | **int64** | ID of template to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCommunicationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommunicationTemplateModel**](CommunicationTemplateModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListTemplate

> CommunicationTemplateListResponse RetrieveListTemplate(ctx).Arguments(arguments).Execute()

Retrieve List of Templates



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
    arguments := *openapiclient.NewCommunicationTemplateRetrieveListArguments() // CommunicationTemplateRetrieveListArguments | Template RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationTemplateApi.RetrieveListTemplate(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationTemplateApi.RetrieveListTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListTemplate`: CommunicationTemplateListResponse
    fmt.Fprintf(os.Stdout, "Response from `CommunicationTemplateApi.RetrieveListTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**CommunicationTemplateRetrieveListArguments**](CommunicationTemplateRetrieveListArguments.md) | Template RetrieveList Arguments | 

### Return type

[**CommunicationTemplateListResponse**](CommunicationTemplateListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

