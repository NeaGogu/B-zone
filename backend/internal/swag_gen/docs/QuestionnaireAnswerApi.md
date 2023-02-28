# \QuestionnaireAnswerApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuestionnaireAnswer**](QuestionnaireAnswerApi.md#CreateQuestionnaireAnswer) | **Post** /questionnaire-answer | Add a new QuestionnaireAnswer
[**DeleteQuestionnaireAnswer**](QuestionnaireAnswerApi.md#DeleteQuestionnaireAnswer) | **Delete** /questionnaire-answer/{questionnaire-answerId} | Delete an QuestionnaireAnswer entry
[**RetrieveListQuestionnaireAnswer**](QuestionnaireAnswerApi.md#RetrieveListQuestionnaireAnswer) | **Put** /questionnaire-answer | Retrieve List of QuestionnaireAnswer
[**RetrieveQuestionnaireAnswer**](QuestionnaireAnswerApi.md#RetrieveQuestionnaireAnswer) | **Get** /questionnaire-answer/{questionnaire-answerId} | Retrieve a QuestionnaireAnswer
[**SetQuestionnaireAnswer**](QuestionnaireAnswerApi.md#SetQuestionnaireAnswer) | **Post** /questionnaire-answer/set | Set (create or update) a QuestionnaireAnswer
[**UpdateQuestionnaireAnswer**](QuestionnaireAnswerApi.md#UpdateQuestionnaireAnswer) | **Put** /questionnaire-answer/{questionnaire-answerId} | Update a specific QuestionnaireAnswer object



## CreateQuestionnaireAnswer

> ApiResponse26 CreateQuestionnaireAnswer(ctx).Body(body).Execute()

Add a new QuestionnaireAnswer



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
    body := *openapiclient.NewQuestionnaireAnswerModel() // QuestionnaireAnswerModel | QuestionnaireAnswer object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireAnswerApi.CreateQuestionnaireAnswer(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnswerApi.CreateQuestionnaireAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuestionnaireAnswer`: ApiResponse26
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnswerApi.CreateQuestionnaireAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestionnaireAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireAnswerModel**](QuestionnaireAnswerModel.md) | QuestionnaireAnswer object that needs to be created | 

### Return type

[**ApiResponse26**](ApiResponse_26.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionnaireAnswer

> ApiResponse24 DeleteQuestionnaireAnswer(ctx, questionnaireAnswerId).Execute()

Delete an QuestionnaireAnswer entry



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
    questionnaireAnswerId := int64(789) // int64 | ID of QuestionnaireAnswer to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireAnswerApi.DeleteQuestionnaireAnswer(context.Background(), questionnaireAnswerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnswerApi.DeleteQuestionnaireAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQuestionnaireAnswer`: ApiResponse24
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnswerApi.DeleteQuestionnaireAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireAnswerId** | **int64** | ID of QuestionnaireAnswer to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionnaireAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse24**](ApiResponse_24.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListQuestionnaireAnswer

> QuestionnaireAnswerListResponse RetrieveListQuestionnaireAnswer(ctx).Arguments(arguments).Execute()

Retrieve List of QuestionnaireAnswer



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
    arguments := *openapiclient.NewQuestionnaireAnswerRetrieveListArguments() // QuestionnaireAnswerRetrieveListArguments | QuestionnaireAnswer RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireAnswerApi.RetrieveListQuestionnaireAnswer(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnswerApi.RetrieveListQuestionnaireAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaireAnswer`: QuestionnaireAnswerListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnswerApi.RetrieveListQuestionnaireAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireAnswerRetrieveListArguments**](QuestionnaireAnswerRetrieveListArguments.md) | QuestionnaireAnswer RetrieveList Arguments | 

### Return type

[**QuestionnaireAnswerListResponse**](QuestionnaireAnswerListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaireAnswer

> QuestionnaireAnswerModel RetrieveQuestionnaireAnswer(ctx, questionnaireAnswerId).Execute()

Retrieve a QuestionnaireAnswer



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
    questionnaireAnswerId := int64(789) // int64 | ID of QuestionnaireAnswer to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireAnswerApi.RetrieveQuestionnaireAnswer(context.Background(), questionnaireAnswerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnswerApi.RetrieveQuestionnaireAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaireAnswer`: QuestionnaireAnswerModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnswerApi.RetrieveQuestionnaireAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireAnswerId** | **int64** | ID of QuestionnaireAnswer to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireAnswerModel**](QuestionnaireAnswerModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetQuestionnaireAnswer

> ApiResponse SetQuestionnaireAnswer(ctx).Body(body).Execute()

Set (create or update) a QuestionnaireAnswer



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
    body := *openapiclient.NewQuestionnaireAnswerModel() // QuestionnaireAnswerModel | QuestionnaireAnswer object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireAnswerApi.SetQuestionnaireAnswer(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnswerApi.SetQuestionnaireAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetQuestionnaireAnswer`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnswerApi.SetQuestionnaireAnswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetQuestionnaireAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireAnswerModel**](QuestionnaireAnswerModel.md) | QuestionnaireAnswer object | 

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


## UpdateQuestionnaireAnswer

> ApiResponse23 UpdateQuestionnaireAnswer(ctx, questionnaireAnswerId).Body(body).Execute()

Update a specific QuestionnaireAnswer object



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
    questionnaireAnswerId := int64(789) // int64 | ID of the QuestionnaireAnswer object to update
    body := *openapiclient.NewQuestionnaireAnswerModel() // QuestionnaireAnswerModel | QuestionnaireAnswer object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireAnswerApi.UpdateQuestionnaireAnswer(context.Background(), questionnaireAnswerId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnswerApi.UpdateQuestionnaireAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuestionnaireAnswer`: ApiResponse23
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnswerApi.UpdateQuestionnaireAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireAnswerId** | **int64** | ID of the QuestionnaireAnswer object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuestionnaireAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**QuestionnaireAnswerModel**](QuestionnaireAnswerModel.md) | QuestionnaireAnswer object that needs to be updated | 

### Return type

[**ApiResponse23**](ApiResponse_23.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

