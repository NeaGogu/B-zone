# \QuestionnaireTemplateApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuestionnaireTemplate**](QuestionnaireTemplateApi.md#CreateQuestionnaireTemplate) | **Post** /questionnaire-template | Add a new QuestionnaireTemplate
[**DeleteQuestionnaireTemplate**](QuestionnaireTemplateApi.md#DeleteQuestionnaireTemplate) | **Delete** /questionnaire-template/{questionnaire-templateId} | Delete an QuestionnaireTemplate entry
[**RetrieveListQuestionnaireTemplate**](QuestionnaireTemplateApi.md#RetrieveListQuestionnaireTemplate) | **Put** /questionnaire-template | Retrieve List of QuestionnaireTemplate
[**RetrieveQuestionnaireTemplate**](QuestionnaireTemplateApi.md#RetrieveQuestionnaireTemplate) | **Get** /questionnaire-template/{questionnaire-templateId} | Retrieve a QuestionnaireTemplate
[**SetQuestionnaireTemplate**](QuestionnaireTemplateApi.md#SetQuestionnaireTemplate) | **Post** /questionnaire-template/set | Set (create or update) a QuestionnaireTemplate
[**UpdateQuestionnaireTemplate**](QuestionnaireTemplateApi.md#UpdateQuestionnaireTemplate) | **Put** /questionnaire-template/{questionnaire-templateId} | Update a specific QuestionnaireTemplate object



## CreateQuestionnaireTemplate

> ApiResponse33 CreateQuestionnaireTemplate(ctx).Body(body).Execute()

Add a new QuestionnaireTemplate



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
    body := *openapiclient.NewQuestionnaireTemplateModel() // QuestionnaireTemplateModel | QuestionnaireTemplate object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateApi.CreateQuestionnaireTemplate(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateApi.CreateQuestionnaireTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuestionnaireTemplate`: ApiResponse33
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateApi.CreateQuestionnaireTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestionnaireTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateModel**](QuestionnaireTemplateModel.md) | QuestionnaireTemplate object that needs to be created | 

### Return type

[**ApiResponse33**](ApiResponse_33.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionnaireTemplate

> ApiResponse31 DeleteQuestionnaireTemplate(ctx, questionnaireTemplateId).Execute()

Delete an QuestionnaireTemplate entry



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
    questionnaireTemplateId := int64(789) // int64 | ID of QuestionnaireTemplate to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateApi.DeleteQuestionnaireTemplate(context.Background(), questionnaireTemplateId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateApi.DeleteQuestionnaireTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQuestionnaireTemplate`: ApiResponse31
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateApi.DeleteQuestionnaireTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateId** | **int64** | ID of QuestionnaireTemplate to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionnaireTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse31**](ApiResponse_31.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListQuestionnaireTemplate

> QuestionnaireTemplateListResponse RetrieveListQuestionnaireTemplate(ctx).Arguments(arguments).Execute()

Retrieve List of QuestionnaireTemplate



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
    arguments := *openapiclient.NewQuestionnaireTemplateRetrieveListArguments() // QuestionnaireTemplateRetrieveListArguments | QuestionnaireTemplate RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateApi.RetrieveListQuestionnaireTemplate(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateApi.RetrieveListQuestionnaireTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaireTemplate`: QuestionnaireTemplateListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateApi.RetrieveListQuestionnaireTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireTemplateRetrieveListArguments**](QuestionnaireTemplateRetrieveListArguments.md) | QuestionnaireTemplate RetrieveList Arguments | 

### Return type

[**QuestionnaireTemplateListResponse**](QuestionnaireTemplateListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaireTemplate

> QuestionnaireTemplateModel RetrieveQuestionnaireTemplate(ctx, questionnaireTemplateId).Execute()

Retrieve a QuestionnaireTemplate



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
    questionnaireTemplateId := int64(789) // int64 | ID of QuestionnaireTemplate to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateApi.RetrieveQuestionnaireTemplate(context.Background(), questionnaireTemplateId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateApi.RetrieveQuestionnaireTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaireTemplate`: QuestionnaireTemplateModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateApi.RetrieveQuestionnaireTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateId** | **int64** | ID of QuestionnaireTemplate to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireTemplateModel**](QuestionnaireTemplateModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetQuestionnaireTemplate

> ApiResponse SetQuestionnaireTemplate(ctx).Body(body).Execute()

Set (create or update) a QuestionnaireTemplate



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
    body := *openapiclient.NewQuestionnaireTemplateModel() // QuestionnaireTemplateModel | QuestionnaireTemplate object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateApi.SetQuestionnaireTemplate(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateApi.SetQuestionnaireTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetQuestionnaireTemplate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateApi.SetQuestionnaireTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetQuestionnaireTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateModel**](QuestionnaireTemplateModel.md) | QuestionnaireTemplate object | 

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


## UpdateQuestionnaireTemplate

> ApiResponse30 UpdateQuestionnaireTemplate(ctx, questionnaireTemplateId).Body(body).Execute()

Update a specific QuestionnaireTemplate object



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
    questionnaireTemplateId := int64(789) // int64 | ID of the QuestionnaireTemplate object to update
    body := *openapiclient.NewQuestionnaireTemplateModel() // QuestionnaireTemplateModel | QuestionnaireTemplate object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateApi.UpdateQuestionnaireTemplate(context.Background(), questionnaireTemplateId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateApi.UpdateQuestionnaireTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuestionnaireTemplate`: ApiResponse30
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateApi.UpdateQuestionnaireTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateId** | **int64** | ID of the QuestionnaireTemplate object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuestionnaireTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**QuestionnaireTemplateModel**](QuestionnaireTemplateModel.md) | QuestionnaireTemplate object that needs to be updated | 

### Return type

[**ApiResponse30**](ApiResponse_30.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

