# \QuestionnaireTemplateQuestionTextApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuestionnaireTemplateQuestionText**](QuestionnaireTemplateQuestionTextApi.md#CreateQuestionnaireTemplateQuestionText) | **Post** /questionnaire-template-question-text | Add a new QuestionnaireTemplateQuestionText
[**DeleteQuestionnaireTemplateQuestionText**](QuestionnaireTemplateQuestionTextApi.md#DeleteQuestionnaireTemplateQuestionText) | **Delete** /questionnaire-template-question-text/{questionnaire-template-question-textId} | Delete an QuestionnaireTemplateQuestionText entry
[**RetrieveListQuestionnaireTemplateQuestionText**](QuestionnaireTemplateQuestionTextApi.md#RetrieveListQuestionnaireTemplateQuestionText) | **Put** /questionnaire-template-question-text | Retrieve List of QuestionnaireTemplateQuestionText
[**RetrieveQuestionnaireTemplateQuestionText**](QuestionnaireTemplateQuestionTextApi.md#RetrieveQuestionnaireTemplateQuestionText) | **Get** /questionnaire-template-question-text/{questionnaire-template-question-textId} | Retrieve a QuestionnaireTemplateQuestionText
[**SetQuestionnaireTemplateQuestionText**](QuestionnaireTemplateQuestionTextApi.md#SetQuestionnaireTemplateQuestionText) | **Post** /questionnaire-template-question-text/set | Set (create or update) a QuestionnaireTemplateQuestionText
[**UpdateQuestionnaireTemplateQuestionText**](QuestionnaireTemplateQuestionTextApi.md#UpdateQuestionnaireTemplateQuestionText) | **Put** /questionnaire-template-question-text/{questionnaire-template-question-textId} | Update a specific QuestionnaireTemplateQuestionText object



## CreateQuestionnaireTemplateQuestionText

> ApiResponse50 CreateQuestionnaireTemplateQuestionText(ctx).Body(body).Execute()

Add a new QuestionnaireTemplateQuestionText



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionTextModel() // QuestionnaireTemplateQuestionTextModel | QuestionnaireTemplateQuestionText object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionTextApi.CreateQuestionnaireTemplateQuestionText(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionTextApi.CreateQuestionnaireTemplateQuestionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuestionnaireTemplateQuestionText`: ApiResponse50
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionTextApi.CreateQuestionnaireTemplateQuestionText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuestionnaireTemplateQuestionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionTextModel**](QuestionnaireTemplateQuestionTextModel.md) | QuestionnaireTemplateQuestionText object that needs to be created | 

### Return type

[**ApiResponse50**](ApiResponse_50.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuestionnaireTemplateQuestionText

> ApiResponse49 DeleteQuestionnaireTemplateQuestionText(ctx, questionnaireTemplateQuestionTextId).Execute()

Delete an QuestionnaireTemplateQuestionText entry



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
    questionnaireTemplateQuestionTextId := int64(789) // int64 | ID of QuestionnaireTemplateQuestionText to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionTextApi.DeleteQuestionnaireTemplateQuestionText(context.Background(), questionnaireTemplateQuestionTextId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionTextApi.DeleteQuestionnaireTemplateQuestionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteQuestionnaireTemplateQuestionText`: ApiResponse49
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionTextApi.DeleteQuestionnaireTemplateQuestionText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionTextId** | **int64** | ID of QuestionnaireTemplateQuestionText to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuestionnaireTemplateQuestionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse49**](ApiResponse_49.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListQuestionnaireTemplateQuestionText

> QuestionnaireTemplateQuestionTextListResponse RetrieveListQuestionnaireTemplateQuestionText(ctx).Arguments(arguments).Execute()

Retrieve List of QuestionnaireTemplateQuestionText



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
    arguments := *openapiclient.NewQuestionnaireTemplateQuestionTextRetrieveListArguments() // QuestionnaireTemplateQuestionTextRetrieveListArguments | QuestionnaireTemplateQuestionText RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionTextApi.RetrieveListQuestionnaireTemplateQuestionText(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionTextApi.RetrieveListQuestionnaireTemplateQuestionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaireTemplateQuestionText`: QuestionnaireTemplateQuestionTextListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionTextApi.RetrieveListQuestionnaireTemplateQuestionText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireTemplateQuestionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireTemplateQuestionTextRetrieveListArguments**](QuestionnaireTemplateQuestionTextRetrieveListArguments.md) | QuestionnaireTemplateQuestionText RetrieveList Arguments | 

### Return type

[**QuestionnaireTemplateQuestionTextListResponse**](QuestionnaireTemplateQuestionTextListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaireTemplateQuestionText

> QuestionnaireTemplateQuestionTextModel RetrieveQuestionnaireTemplateQuestionText(ctx, questionnaireTemplateQuestionTextId).Execute()

Retrieve a QuestionnaireTemplateQuestionText



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
    questionnaireTemplateQuestionTextId := int64(789) // int64 | ID of QuestionnaireTemplateQuestionText to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionTextApi.RetrieveQuestionnaireTemplateQuestionText(context.Background(), questionnaireTemplateQuestionTextId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionTextApi.RetrieveQuestionnaireTemplateQuestionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaireTemplateQuestionText`: QuestionnaireTemplateQuestionTextModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionTextApi.RetrieveQuestionnaireTemplateQuestionText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionTextId** | **int64** | ID of QuestionnaireTemplateQuestionText to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireTemplateQuestionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireTemplateQuestionTextModel**](QuestionnaireTemplateQuestionTextModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetQuestionnaireTemplateQuestionText

> ApiResponse SetQuestionnaireTemplateQuestionText(ctx).Body(body).Execute()

Set (create or update) a QuestionnaireTemplateQuestionText



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
    body := *openapiclient.NewQuestionnaireTemplateQuestionTextModel() // QuestionnaireTemplateQuestionTextModel | QuestionnaireTemplateQuestionText object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionTextApi.SetQuestionnaireTemplateQuestionText(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionTextApi.SetQuestionnaireTemplateQuestionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetQuestionnaireTemplateQuestionText`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionTextApi.SetQuestionnaireTemplateQuestionText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetQuestionnaireTemplateQuestionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QuestionnaireTemplateQuestionTextModel**](QuestionnaireTemplateQuestionTextModel.md) | QuestionnaireTemplateQuestionText object | 

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


## UpdateQuestionnaireTemplateQuestionText

> ApiResponse48 UpdateQuestionnaireTemplateQuestionText(ctx, questionnaireTemplateQuestionTextId).Body(body).Execute()

Update a specific QuestionnaireTemplateQuestionText object



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
    questionnaireTemplateQuestionTextId := int64(789) // int64 | ID of the QuestionnaireTemplateQuestionText object to update
    body := *openapiclient.NewQuestionnaireTemplateQuestionTextModel() // QuestionnaireTemplateQuestionTextModel | QuestionnaireTemplateQuestionText object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTemplateQuestionTextApi.UpdateQuestionnaireTemplateQuestionText(context.Background(), questionnaireTemplateQuestionTextId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTemplateQuestionTextApi.UpdateQuestionnaireTemplateQuestionText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuestionnaireTemplateQuestionText`: ApiResponse48
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTemplateQuestionTextApi.UpdateQuestionnaireTemplateQuestionText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTemplateQuestionTextId** | **int64** | ID of the QuestionnaireTemplateQuestionText object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuestionnaireTemplateQuestionTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**QuestionnaireTemplateQuestionTextModel**](QuestionnaireTemplateQuestionTextModel.md) | QuestionnaireTemplateQuestionText object that needs to be updated | 

### Return type

[**ApiResponse48**](ApiResponse_48.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

