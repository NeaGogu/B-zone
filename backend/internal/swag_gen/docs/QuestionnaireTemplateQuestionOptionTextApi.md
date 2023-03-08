# \QuestionnaireTemplateQuestionOptionTextApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuestionnaireTemplateQuestionOptionText**](QuestionnaireTemplateQuestionOptionTextApi.md#CreateQuestionnaireTemplateQuestionOptionText) | **Post** /questionnaire-template-question-option-text | Add a new QuestionnaireTemplateQuestionOptionText
[**DeleteQuestionnaireTemplateQuestionOptionText**](QuestionnaireTemplateQuestionOptionTextApi.md#DeleteQuestionnaireTemplateQuestionOptionText) | **Delete** /questionnaire-template-question-option-text/{questionnaire-template-question-option-textId} | Delete an QuestionnaireTemplateQuestionOptionText entry
[**RetrieveListQuestionnaireTemplateQuestionOptionText**](QuestionnaireTemplateQuestionOptionTextApi.md#RetrieveListQuestionnaireTemplateQuestionOptionText) | **Put** /questionnaire-template-question-option-text | Retrieve List of QuestionnaireTemplateQuestionOptionText
[**RetrieveQuestionnaireTemplateQuestionOptionText**](QuestionnaireTemplateQuestionOptionTextApi.md#RetrieveQuestionnaireTemplateQuestionOptionText) | **Get** /questionnaire-template-question-option-text/{questionnaire-template-question-option-textId} | Retrieve a QuestionnaireTemplateQuestionOptionText
[**SetQuestionnaireTemplateQuestionOptionText**](QuestionnaireTemplateQuestionOptionTextApi.md#SetQuestionnaireTemplateQuestionOptionText) | **Post** /questionnaire-template-question-option-text/set | Set (create or update) a QuestionnaireTemplateQuestionOptionText
[**UpdateQuestionnaireTemplateQuestionOptionText**](QuestionnaireTemplateQuestionOptionTextApi.md#UpdateQuestionnaireTemplateQuestionOptionText) | **Put** /questionnaire-template-question-option-text/{questionnaire-template-question-option-textId} | Update a specific QuestionnaireTemplateQuestionOptionText object



## CreateQuestionnaireTemplateQuestionOptionText

> ApiResponse47 CreateQuestionnaireTemplateQuestionOptionText(ctx).Body(body).Execute()

Add a new QuestionnaireTemplateQuestionOptionText



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionOptionTextModel() // QuestionnaireTemplateQuestionOptionTextModel | QuestionnaireTemplateQuestionOptionText object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionTextApi.CreateQuestionnaireTemplateQuestionOptionText(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionTextApi.CreateQuestionnaireTemplateQuestionOptionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuestionnaireTemplateQuestionOptionText`: ApiResponse47
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionTextApi.CreateQuestionnaireTemplateQuestionOptionText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestionnaireTemplateQuestionOptionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionOptionTextModel**](QuestionnaireTemplateQuestionOptionTextModel.md) | QuestionnaireTemplateQuestionOptionText object that needs to be created | 

### Return type

[**ApiResponse47**](ApiResponse_47.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionnaireTemplateQuestionOptionText

> ApiResponse46 DeleteQuestionnaireTemplateQuestionOptionText(ctx, questionnaireTemplateQuestionOptionTextId).Execute()

Delete an QuestionnaireTemplateQuestionOptionText entry



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
    questionnaireTemplateQuestionOptionTextId := int64(789) // int64 | ID of QuestionnaireTemplateQuestionOptionText to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionTextApi.DeleteQuestionnaireTemplateQuestionOptionText(context.Background(), questionnaireTemplateQuestionOptionTextId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionTextApi.DeleteQuestionnaireTemplateQuestionOptionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQuestionnaireTemplateQuestionOptionText`: ApiResponse46
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionTextApi.DeleteQuestionnaireTemplateQuestionOptionText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionOptionTextId** | **int64** | ID of QuestionnaireTemplateQuestionOptionText to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionnaireTemplateQuestionOptionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse46**](ApiResponse_46.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListQuestionnaireTemplateQuestionOptionText

> QuestionnaireTemplateQuestionOptionTextListResponse RetrieveListQuestionnaireTemplateQuestionOptionText(ctx).Arguments(arguments).Execute()

Retrieve List of QuestionnaireTemplateQuestionOptionText



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
    arguments := *openapiclient.NewQuestionnaireTemplateQuestionOptionTextRetrieveListArguments() // QuestionnaireTemplateQuestionOptionTextRetrieveListArguments | QuestionnaireTemplateQuestionOptionText RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionTextApi.RetrieveListQuestionnaireTemplateQuestionOptionText(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionTextApi.RetrieveListQuestionnaireTemplateQuestionOptionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaireTemplateQuestionOptionText`: QuestionnaireTemplateQuestionOptionTextListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionTextApi.RetrieveListQuestionnaireTemplateQuestionOptionText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireTemplateQuestionOptionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireTemplateQuestionOptionTextRetrieveListArguments**](QuestionnaireTemplateQuestionOptionTextRetrieveListArguments.md) | QuestionnaireTemplateQuestionOptionText RetrieveList Arguments | 

### Return type

[**QuestionnaireTemplateQuestionOptionTextListResponse**](QuestionnaireTemplateQuestionOptionTextListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaireTemplateQuestionOptionText

> QuestionnaireTemplateQuestionOptionTextModel RetrieveQuestionnaireTemplateQuestionOptionText(ctx, questionnaireTemplateQuestionOptionTextId).Execute()

Retrieve a QuestionnaireTemplateQuestionOptionText



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
    questionnaireTemplateQuestionOptionTextId := int64(789) // int64 | ID of QuestionnaireTemplateQuestionOptionText to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionTextApi.RetrieveQuestionnaireTemplateQuestionOptionText(context.Background(), questionnaireTemplateQuestionOptionTextId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionTextApi.RetrieveQuestionnaireTemplateQuestionOptionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaireTemplateQuestionOptionText`: QuestionnaireTemplateQuestionOptionTextModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionTextApi.RetrieveQuestionnaireTemplateQuestionOptionText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionOptionTextId** | **int64** | ID of QuestionnaireTemplateQuestionOptionText to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireTemplateQuestionOptionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireTemplateQuestionOptionTextModel**](QuestionnaireTemplateQuestionOptionTextModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetQuestionnaireTemplateQuestionOptionText

> ApiResponse SetQuestionnaireTemplateQuestionOptionText(ctx).Body(body).Execute()

Set (create or update) a QuestionnaireTemplateQuestionOptionText



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionOptionTextModel() // QuestionnaireTemplateQuestionOptionTextModel | QuestionnaireTemplateQuestionOptionText object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionTextApi.SetQuestionnaireTemplateQuestionOptionText(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionTextApi.SetQuestionnaireTemplateQuestionOptionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetQuestionnaireTemplateQuestionOptionText`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionTextApi.SetQuestionnaireTemplateQuestionOptionText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetQuestionnaireTemplateQuestionOptionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionOptionTextModel**](QuestionnaireTemplateQuestionOptionTextModel.md) | QuestionnaireTemplateQuestionOptionText object | 

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


## UpdateQuestionnaireTemplateQuestionOptionText

> ApiResponse45 UpdateQuestionnaireTemplateQuestionOptionText(ctx, questionnaireTemplateQuestionOptionTextId).Body(body).Execute()

Update a specific QuestionnaireTemplateQuestionOptionText object



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
    questionnaireTemplateQuestionOptionTextId := int64(789) // int64 | ID of the QuestionnaireTemplateQuestionOptionText object to update
    body := *openapiclient.NewQuestionnaireTemplateQuestionOptionTextModel() // QuestionnaireTemplateQuestionOptionTextModel | QuestionnaireTemplateQuestionOptionText object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionOptionTextApi.UpdateQuestionnaireTemplateQuestionOptionText(context.Background(), questionnaireTemplateQuestionOptionTextId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionOptionTextApi.UpdateQuestionnaireTemplateQuestionOptionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuestionnaireTemplateQuestionOptionText`: ApiResponse45
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionOptionTextApi.UpdateQuestionnaireTemplateQuestionOptionText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionOptionTextId** | **int64** | ID of the QuestionnaireTemplateQuestionOptionText object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuestionnaireTemplateQuestionOptionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**QuestionnaireTemplateQuestionOptionTextModel**](QuestionnaireTemplateQuestionOptionTextModel.md) | QuestionnaireTemplateQuestionOptionText object that needs to be updated | 

### Return type

[**ApiResponse45**](ApiResponse_45.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

