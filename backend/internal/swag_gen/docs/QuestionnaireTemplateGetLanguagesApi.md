# \QuestionnaireTemplateGetLanguagesApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuestionnaireTemplateLanguages**](QuestionnaireTemplateGetLanguagesApi.md#GetQuestionnaireTemplateLanguages) | **Get** /questionnaire-template/get-languages/{questionnaire-templateId} | Retrieves all set languages for an QuestionnaireTemplate entry



## GetQuestionnaireTemplateLanguages

> ApiResponse34 GetQuestionnaireTemplateLanguages(ctx, questionnaireTemplateId).Execute()

Retrieves all set languages for an QuestionnaireTemplate entry



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
    questionnaireTemplateId := int64(789) // int64 | ID of QuestionnaireTemplate

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateGetLanguagesApi.GetQuestionnaireTemplateLanguages(context.Background(), questionnaireTemplateId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateGetLanguagesApi.GetQuestionnaireTemplateLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuestionnaireTemplateLanguages`: ApiResponse34
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateGetLanguagesApi.GetQuestionnaireTemplateLanguages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateId** | **int64** | ID of QuestionnaireTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuestionnaireTemplateLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse34**](ApiResponse_34.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

