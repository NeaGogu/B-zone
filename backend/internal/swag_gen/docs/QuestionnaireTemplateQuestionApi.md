# \QuestionnaireTemplateQuestionApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuestionnaireTemplateQuestion**](QuestionnaireTemplateQuestionApi.md#CreateQuestionnaireTemplateQuestion) | **Post** /questionnaire-template-question | Add a new QuestionnaireTemplateQuestion
[**DeleteQuestionnaireTemplateQuestion**](QuestionnaireTemplateQuestionApi.md#DeleteQuestionnaireTemplateQuestion) | **Delete** /questionnaire-template-question/{questionnaireTemplateQuestionId} | Delete an QuestionnaireTemplateQuestion entry
[**RetrieveListQuestionnaireTemplateQuestion**](QuestionnaireTemplateQuestionApi.md#RetrieveListQuestionnaireTemplateQuestion) | **Put** /questionnaire-template-question | Retrieve List of QuestionnaireTemplateQuestion
[**RetrieveQuestionnaireTemplateQuestion**](QuestionnaireTemplateQuestionApi.md#RetrieveQuestionnaireTemplateQuestion) | **Get** /questionnaire-template-question/{questionnaireTemplateQuestionId} | Retrieve a QuestionnaireTemplateQuestion
[**SetQuestionnaireTemplateQuestion**](QuestionnaireTemplateQuestionApi.md#SetQuestionnaireTemplateQuestion) | **Post** /questionnaire-template-question/set | Set (create or update) a QuestionnaireTemplateQuestion
[**UpdateQuestionnaireTemplateQuestion**](QuestionnaireTemplateQuestionApi.md#UpdateQuestionnaireTemplateQuestion) | **Put** /questionnaire-template-question/{questionnaireTemplateQuestionId} | Update a specific QuestionnaireTemplateQuestion object



## CreateQuestionnaireTemplateQuestion

> ApiResponse38 CreateQuestionnaireTemplateQuestion(ctx).Body(body).Execute()

Add a new QuestionnaireTemplateQuestion



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionModel() // QuestionnaireTemplateQuestionModel | QuestionnaireTemplateQuestion object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionApi.CreateQuestionnaireTemplateQuestion(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionApi.CreateQuestionnaireTemplateQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuestionnaireTemplateQuestion`: ApiResponse38
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionApi.CreateQuestionnaireTemplateQuestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestionnaireTemplateQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionModel**](QuestionnaireTemplateQuestionModel.md) | QuestionnaireTemplateQuestion object that needs to be created | 

### Return type

[**ApiResponse38**](ApiResponse_38.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionnaireTemplateQuestion

> ApiResponse36 DeleteQuestionnaireTemplateQuestion(ctx, questionnaireTemplateQuestionId).Execute()

Delete an QuestionnaireTemplateQuestion entry



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
    questionnaireTemplateQuestionId := int64(789) // int64 | ID of QuestionnaireTemplateQuestion to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionApi.DeleteQuestionnaireTemplateQuestion(context.Background(), questionnaireTemplateQuestionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionApi.DeleteQuestionnaireTemplateQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQuestionnaireTemplateQuestion`: ApiResponse36
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionApi.DeleteQuestionnaireTemplateQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionId** | **int64** | ID of QuestionnaireTemplateQuestion to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionnaireTemplateQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse36**](ApiResponse_36.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListQuestionnaireTemplateQuestion

> QuestionnaireTemplateQuestionListResponse RetrieveListQuestionnaireTemplateQuestion(ctx).Arguments(arguments).Execute()

Retrieve List of QuestionnaireTemplateQuestion



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
    arguments := *openapiclient.NewQuestionnaireTemplateQuestionRetrieveListArguments() // QuestionnaireTemplateQuestionRetrieveListArguments | QuestionnaireTemplateQuestion RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionApi.RetrieveListQuestionnaireTemplateQuestion(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionApi.RetrieveListQuestionnaireTemplateQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaireTemplateQuestion`: QuestionnaireTemplateQuestionListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionApi.RetrieveListQuestionnaireTemplateQuestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireTemplateQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireTemplateQuestionRetrieveListArguments**](QuestionnaireTemplateQuestionRetrieveListArguments.md) | QuestionnaireTemplateQuestion RetrieveList Arguments | 

### Return type

[**QuestionnaireTemplateQuestionListResponse**](QuestionnaireTemplateQuestionListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaireTemplateQuestion

> QuestionnaireTemplateQuestionModel RetrieveQuestionnaireTemplateQuestion(ctx, questionnaireTemplateQuestionId).Execute()

Retrieve a QuestionnaireTemplateQuestion



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
    questionnaireTemplateQuestionId := int64(789) // int64 | ID of QuestionnaireTemplateQuestion to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionApi.RetrieveQuestionnaireTemplateQuestion(context.Background(), questionnaireTemplateQuestionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionApi.RetrieveQuestionnaireTemplateQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaireTemplateQuestion`: QuestionnaireTemplateQuestionModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionApi.RetrieveQuestionnaireTemplateQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionId** | **int64** | ID of QuestionnaireTemplateQuestion to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireTemplateQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireTemplateQuestionModel**](QuestionnaireTemplateQuestionModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetQuestionnaireTemplateQuestion

> ApiResponse SetQuestionnaireTemplateQuestion(ctx).Body(body).Execute()

Set (create or update) a QuestionnaireTemplateQuestion



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionModel() // QuestionnaireTemplateQuestionModel | QuestionnaireTemplateQuestion object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionApi.SetQuestionnaireTemplateQuestion(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionApi.SetQuestionnaireTemplateQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetQuestionnaireTemplateQuestion`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionApi.SetQuestionnaireTemplateQuestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetQuestionnaireTemplateQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionModel**](QuestionnaireTemplateQuestionModel.md) | QuestionnaireTemplateQuestion object | 

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


## UpdateQuestionnaireTemplateQuestion

> ApiResponse35 UpdateQuestionnaireTemplateQuestion(ctx, questionnaireTemplateQuestionId).Body(body).Execute()

Update a specific QuestionnaireTemplateQuestion object



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
    questionnaireTemplateQuestionId := int64(789) // int64 | ID of the QuestionnaireTemplateQuestion object to update
    body := *openapiclient.NewQuestionnaireTemplateQuestionModel() // QuestionnaireTemplateQuestionModel | QuestionnaireTemplateQuestion object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionApi.UpdateQuestionnaireTemplateQuestion(context.Background(), questionnaireTemplateQuestionId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionApi.UpdateQuestionnaireTemplateQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuestionnaireTemplateQuestion`: ApiResponse35
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionApi.UpdateQuestionnaireTemplateQuestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionId** | **int64** | ID of the QuestionnaireTemplateQuestion object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuestionnaireTemplateQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**QuestionnaireTemplateQuestionModel**](QuestionnaireTemplateQuestionModel.md) | QuestionnaireTemplateQuestion object that needs to be updated | 

### Return type

[**ApiResponse35**](ApiResponse_35.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

