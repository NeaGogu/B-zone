# \QuestionnaireTemplateQuestionOptionApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuestionnaireTemplateQuestionOption**](QuestionnaireTemplateQuestionOptionApi.md#CreateQuestionnaireTemplateQuestionOption) | **Post** /questionnaire-template-question-option | Add a new QuestionnaireTemplateQuestionOption
[**DeleteQuestionnaireTemplateQuestionOption**](QuestionnaireTemplateQuestionOptionApi.md#DeleteQuestionnaireTemplateQuestionOption) | **Delete** /questionnaire-template-question-option/{questionnaire-template-question-optionId} | Delete an QuestionnaireTemplateQuestionOption entry
[**FindPossibleFollowUpQuestions**](QuestionnaireTemplateQuestionOptionApi.md#FindPossibleFollowUpQuestions) | **Post** /questionnaire-template-question-option/find-possible-followup-questions | find possible follow up questions
[**GetPossibleFollowUpQuestions**](QuestionnaireTemplateQuestionOptionApi.md#GetPossibleFollowUpQuestions) | **Post** /questionnaire-template-question-option/get-possible-followup-questions | get possible follow up questions
[**RetrieveListQuestionnaireTemplateQuestionOption**](QuestionnaireTemplateQuestionOptionApi.md#RetrieveListQuestionnaireTemplateQuestionOption) | **Put** /questionnaire-template-question-option | Retrieve List of QuestionnaireTemplateQuestionOption
[**RetrieveQuestionnaireTemplateQuestionOption**](QuestionnaireTemplateQuestionOptionApi.md#RetrieveQuestionnaireTemplateQuestionOption) | **Get** /questionnaire-template-question-option/{questionnaire-template-question-optionId} | Retrieve a QuestionnaireTemplateQuestionOption
[**SetQuestionnaireTemplateQuestionOption**](QuestionnaireTemplateQuestionOptionApi.md#SetQuestionnaireTemplateQuestionOption) | **Post** /questionnaire-template-question-option/set | Set (create or update) a QuestionnaireTemplateQuestionOption
[**UpdateQuestionnaireTemplateQuestionOption**](QuestionnaireTemplateQuestionOptionApi.md#UpdateQuestionnaireTemplateQuestionOption) | **Put** /questionnaire-template-question-option/{questionnaire-template-question-optionId} | Update a specific QuestionnaireTemplateQuestionOption object



## CreateQuestionnaireTemplateQuestionOption

> ApiResponse42 CreateQuestionnaireTemplateQuestionOption(ctx).Body(body).Execute()

Add a new QuestionnaireTemplateQuestionOption



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionOptionModel() // QuestionnaireTemplateQuestionOptionModel | QuestionnaireTemplateQuestionOption object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionApi.CreateQuestionnaireTemplateQuestionOption(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionApi.CreateQuestionnaireTemplateQuestionOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuestionnaireTemplateQuestionOption`: ApiResponse42
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionApi.CreateQuestionnaireTemplateQuestionOption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestionnaireTemplateQuestionOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionOptionModel**](QuestionnaireTemplateQuestionOptionModel.md) | QuestionnaireTemplateQuestionOption object that needs to be created | 

### Return type

[**ApiResponse42**](ApiResponse_42.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionnaireTemplateQuestionOption

> ApiResponse40 DeleteQuestionnaireTemplateQuestionOption(ctx, questionnaireTemplateQuestionOptionId).Execute()

Delete an QuestionnaireTemplateQuestionOption entry



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
    questionnaireTemplateQuestionOptionId := int64(789) // int64 | ID of QuestionnaireTemplateQuestionOption to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionApi.DeleteQuestionnaireTemplateQuestionOption(context.Background(), questionnaireTemplateQuestionOptionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionApi.DeleteQuestionnaireTemplateQuestionOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQuestionnaireTemplateQuestionOption`: ApiResponse40
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionApi.DeleteQuestionnaireTemplateQuestionOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionOptionId** | **int64** | ID of QuestionnaireTemplateQuestionOption to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionnaireTemplateQuestionOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse40**](ApiResponse_40.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindPossibleFollowUpQuestions

> ApiResponse44 FindPossibleFollowUpQuestions(ctx).Body(body).Execute()

find possible follow up questions



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionOptionPossibleQuestionsModel() // QuestionnaireTemplateQuestionOptionPossibleQuestionsModel | QuestionnaireTemplate object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionApi.FindPossibleFollowUpQuestions(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionApi.FindPossibleFollowUpQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindPossibleFollowUpQuestions`: ApiResponse44
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionApi.FindPossibleFollowUpQuestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindPossibleFollowUpQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionOptionPossibleQuestionsModel**](QuestionnaireTemplateQuestionOptionPossibleQuestionsModel.md) | QuestionnaireTemplate object that needs to be updated | 

### Return type

[**ApiResponse44**](ApiResponse_44.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPossibleFollowUpQuestions

> ApiResponse40 GetPossibleFollowUpQuestions(ctx).Body(body).Execute()

get possible follow up questions



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionOptionPossibleQuestionsModel() // QuestionnaireTemplateQuestionOptionPossibleQuestionsModel | QuestionnaireTemplate object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionApi.GetPossibleFollowUpQuestions(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionApi.GetPossibleFollowUpQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPossibleFollowUpQuestions`: ApiResponse40
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionApi.GetPossibleFollowUpQuestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPossibleFollowUpQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionOptionPossibleQuestionsModel**](QuestionnaireTemplateQuestionOptionPossibleQuestionsModel.md) | QuestionnaireTemplate object that needs to be updated | 

### Return type

[**ApiResponse40**](ApiResponse_40.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListQuestionnaireTemplateQuestionOption

> QuestionnaireTemplateQuestionOptionListResponse RetrieveListQuestionnaireTemplateQuestionOption(ctx).Arguments(arguments).Execute()

Retrieve List of QuestionnaireTemplateQuestionOption



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
    arguments := *openapiclient.NewQuestionnaireTemplateQuestionOptionRetrieveListArguments() // QuestionnaireTemplateQuestionOptionRetrieveListArguments | QuestionnaireTemplateQuestionOption RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionApi.RetrieveListQuestionnaireTemplateQuestionOption(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionApi.RetrieveListQuestionnaireTemplateQuestionOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaireTemplateQuestionOption`: QuestionnaireTemplateQuestionOptionListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionApi.RetrieveListQuestionnaireTemplateQuestionOption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireTemplateQuestionOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireTemplateQuestionOptionRetrieveListArguments**](QuestionnaireTemplateQuestionOptionRetrieveListArguments.md) | QuestionnaireTemplateQuestionOption RetrieveList Arguments | 

### Return type

[**QuestionnaireTemplateQuestionOptionListResponse**](QuestionnaireTemplateQuestionOptionListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaireTemplateQuestionOption

> QuestionnaireTemplateQuestionOptionModel RetrieveQuestionnaireTemplateQuestionOption(ctx, questionnaireTemplateQuestionOptionId).Execute()

Retrieve a QuestionnaireTemplateQuestionOption



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
    questionnaireTemplateQuestionOptionId := int64(789) // int64 | ID of QuestionnaireTemplateQuestionOption to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionApi.RetrieveQuestionnaireTemplateQuestionOption(context.Background(), questionnaireTemplateQuestionOptionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionApi.RetrieveQuestionnaireTemplateQuestionOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaireTemplateQuestionOption`: QuestionnaireTemplateQuestionOptionModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionApi.RetrieveQuestionnaireTemplateQuestionOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionOptionId** | **int64** | ID of QuestionnaireTemplateQuestionOption to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireTemplateQuestionOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireTemplateQuestionOptionModel**](QuestionnaireTemplateQuestionOptionModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetQuestionnaireTemplateQuestionOption

> ApiResponse SetQuestionnaireTemplateQuestionOption(ctx).Body(body).Execute()

Set (create or update) a QuestionnaireTemplateQuestionOption



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionOptionModel() // QuestionnaireTemplateQuestionOptionModel | QuestionnaireTemplateQuestionOption object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionApi.SetQuestionnaireTemplateQuestionOption(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionApi.SetQuestionnaireTemplateQuestionOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetQuestionnaireTemplateQuestionOption`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionApi.SetQuestionnaireTemplateQuestionOption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetQuestionnaireTemplateQuestionOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionOptionModel**](QuestionnaireTemplateQuestionOptionModel.md) | QuestionnaireTemplateQuestionOption object | 

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


## UpdateQuestionnaireTemplateQuestionOption

> ApiResponse39 UpdateQuestionnaireTemplateQuestionOption(ctx, questionnaireTemplateQuestionOptionId).Body(body).Execute()

Update a specific QuestionnaireTemplateQuestionOption object



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
    questionnaireTemplateQuestionOptionId := int64(789) // int64 | ID of the QuestionnaireTemplateQuestionOption object to update
    body := *openapiclient.NewQuestionnaireTemplateQuestionOptionModel() // QuestionnaireTemplateQuestionOptionModel | QuestionnaireTemplateQuestionOption object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionApi.UpdateQuestionnaireTemplateQuestionOption(context.Background(), questionnaireTemplateQuestionOptionId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionApi.UpdateQuestionnaireTemplateQuestionOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuestionnaireTemplateQuestionOption`: ApiResponse39
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionApi.UpdateQuestionnaireTemplateQuestionOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionOptionId** | **int64** | ID of the QuestionnaireTemplateQuestionOption object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuestionnaireTemplateQuestionOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**QuestionnaireTemplateQuestionOptionModel**](QuestionnaireTemplateQuestionOptionModel.md) | QuestionnaireTemplateQuestionOption object that needs to be updated | 

### Return type

[**ApiResponse39**](ApiResponse_39.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

