# \QuestionnaireQuestionTypeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveListQuestionnaireQuestionType**](QuestionnaireQuestionTypeApi.md#RetrieveListQuestionnaireQuestionType) | **Put** /questionnaire-question-type | Retrieve List of QuestionnaireQuestionType
[**RetrieveQuestionnaireQuestionType**](QuestionnaireQuestionTypeApi.md#RetrieveQuestionnaireQuestionType) | **Get** /questionnaire-question-type/{questionnaire-question-typeId} | Retrieve a QuestionnaireQuestionType



## RetrieveListQuestionnaireQuestionType

> QuestionnaireQuestionTypeListResponse RetrieveListQuestionnaireQuestionType(ctx).Arguments(arguments).Execute()

Retrieve List of QuestionnaireQuestionType



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
    arguments := *openapiclient.NewQuestionnaireQuestionTypeRetrieveListArguments() // QuestionnaireQuestionTypeRetrieveListArguments | QuestionnaireQuestionType RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireQuestionTypeApi.RetrieveListQuestionnaireQuestionType(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireQuestionTypeApi.RetrieveListQuestionnaireQuestionType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaireQuestionType`: QuestionnaireQuestionTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireQuestionTypeApi.RetrieveListQuestionnaireQuestionType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireQuestionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireQuestionTypeRetrieveListArguments**](QuestionnaireQuestionTypeRetrieveListArguments.md) | QuestionnaireQuestionType RetrieveList Arguments | 

### Return type

[**QuestionnaireQuestionTypeListResponse**](QuestionnaireQuestionTypeListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaireQuestionType

> QuestionnaireQuestionTypeModel RetrieveQuestionnaireQuestionType(ctx, questionnaireQuestionTypeId).Execute()

Retrieve a QuestionnaireQuestionType



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
    questionnaireQuestionTypeId := int64(789) // int64 | ID of QuestionnaireQuestionType to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireQuestionTypeApi.RetrieveQuestionnaireQuestionType(context.Background(), questionnaireQuestionTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireQuestionTypeApi.RetrieveQuestionnaireQuestionType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaireQuestionType`: QuestionnaireQuestionTypeModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireQuestionTypeApi.RetrieveQuestionnaireQuestionType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireQuestionTypeId** | **int64** | ID of QuestionnaireQuestionType to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireQuestionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireQuestionTypeModel**](QuestionnaireQuestionTypeModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

