# \QuestionnaireApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeLanguage**](QuestionnaireApi.md#ChangeLanguage) | **Post** /questionnaire/change-language | change language of a Questionnaire
[**DeleteQuestionnaire**](QuestionnaireApi.md#DeleteQuestionnaire) | **Delete** /questionnaire/{questionnaireId} | Delete an Questionnaire entry
[**GetNextQuestion**](QuestionnaireApi.md#GetNextQuestion) | **Post** /questionnaire/get-next-question | getNextQuestion of an Questionnaire
[**GetPreviousQuestion**](QuestionnaireApi.md#GetPreviousQuestion) | **Post** /questionnaire/get-previous-question | getPreviousQuestion of an Questionnaire
[**RetrieveListQuestionnaire**](QuestionnaireApi.md#RetrieveListQuestionnaire) | **Put** /questionnaire | Retrieve List of Questionnaire
[**RetrieveQuestionnaire**](QuestionnaireApi.md#RetrieveQuestionnaire) | **Get** /questionnaire/{questionnaireId} | Retrieve a Questionnaire



## ChangeLanguage

> QuestionnaireChangeLanguageResponse ChangeLanguage(ctx).Arguments(arguments).Execute()

change language of a Questionnaire



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
    arguments := *openapiclient.NewQuestionnaireChangeLanguageArguments() // QuestionnaireChangeLanguageArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireApi.ChangeLanguage(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireApi.ChangeLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeLanguage`: QuestionnaireChangeLanguageResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireApi.ChangeLanguage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireChangeLanguageArguments**](QuestionnaireChangeLanguageArguments.md) | Request Arguments | 

### Return type

[**QuestionnaireChangeLanguageResponse**](QuestionnaireChangeLanguageResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionnaire

> ApiResponse27 DeleteQuestionnaire(ctx, questionnaireId).Execute()

Delete an Questionnaire entry



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
    questionnaireId := int64(789) // int64 | ID of Questionnaire to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireApi.DeleteQuestionnaire(context.Background(), questionnaireId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireApi.DeleteQuestionnaire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQuestionnaire`: ApiResponse27
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireApi.DeleteQuestionnaire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireId** | **int64** | ID of Questionnaire to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionnaireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse27**](ApiResponse_27.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextQuestion

> QuestionnaireQuestionModel GetNextQuestion(ctx).Arguments(arguments).Execute()

getNextQuestion of an Questionnaire



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
    arguments := *openapiclient.NewQuestionnaireGetNextQuestionArguments() // QuestionnaireGetNextQuestionArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireApi.GetNextQuestion(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireApi.GetNextQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextQuestion`: QuestionnaireQuestionModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireApi.GetNextQuestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNextQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireGetNextQuestionArguments**](QuestionnaireGetNextQuestionArguments.md) | Request Arguments | 

### Return type

[**QuestionnaireQuestionModel**](QuestionnaireQuestionModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreviousQuestion

> QuestionnaireQuestionModel GetPreviousQuestion(ctx).Arguments(arguments).Execute()

getPreviousQuestion of an Questionnaire



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
    arguments := *openapiclient.NewQuestionnaireGetPreviousQuestionArguments() // QuestionnaireGetPreviousQuestionArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireApi.GetPreviousQuestion(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireApi.GetPreviousQuestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviousQuestion`: QuestionnaireQuestionModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireApi.GetPreviousQuestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPreviousQuestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireGetPreviousQuestionArguments**](QuestionnaireGetPreviousQuestionArguments.md) | Request Arguments | 

### Return type

[**QuestionnaireQuestionModel**](QuestionnaireQuestionModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListQuestionnaire

> QuestionnaireListResponse RetrieveListQuestionnaire(ctx).Arguments(arguments).Execute()

Retrieve List of Questionnaire



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
    arguments := *openapiclient.NewQuestionnaireRetrieveListArguments() // QuestionnaireRetrieveListArguments | Questionnaire RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireApi.RetrieveListQuestionnaire(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireApi.RetrieveListQuestionnaire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaire`: QuestionnaireListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireApi.RetrieveListQuestionnaire`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireRetrieveListArguments**](QuestionnaireRetrieveListArguments.md) | Questionnaire RetrieveList Arguments | 

### Return type

[**QuestionnaireListResponse**](QuestionnaireListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaire

> QuestionnaireModel RetrieveQuestionnaire(ctx, questionnaireId).IncludeAnswers(includeAnswers).Execute()

Retrieve a Questionnaire



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
    questionnaireId := int64(789) // int64 | ID of Questionnaire to retrieve
    includeAnswers := true // bool | Include answers (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireApi.RetrieveQuestionnaire(context.Background(), questionnaireId).IncludeAnswers(includeAnswers).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireApi.RetrieveQuestionnaire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaire`: QuestionnaireModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireApi.RetrieveQuestionnaire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireId** | **int64** | ID of Questionnaire to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAnswers** | **bool** | Include answers | [default to false]

### Return type

[**QuestionnaireModel**](QuestionnaireModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

