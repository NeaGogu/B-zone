# \QuestionnaireTypeApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveListQuestionnaireType**](QuestionnaireTypeApi.md#RetrieveListQuestionnaireType) | **Put** /questionnaire-type | Retrieve List of QuestionnaireType
[**RetrieveQuestionnaireType**](QuestionnaireTypeApi.md#RetrieveQuestionnaireType) | **Get** /questionnaire-type/{questionnaire-typeId} | Retrieve a QuestionnaireType



## RetrieveListQuestionnaireType

> QuestionnaireTypeListResponse RetrieveListQuestionnaireType(ctx).Arguments(arguments).Execute()

Retrieve List of QuestionnaireType



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
    arguments := *openapiclient.NewQuestionnaireTypeRetrieveListArguments() // QuestionnaireTypeRetrieveListArguments | QuestionnaireType RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTypeApi.RetrieveListQuestionnaireType(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTypeApi.RetrieveListQuestionnaireType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListQuestionnaireType`: QuestionnaireTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTypeApi.RetrieveListQuestionnaireType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListQuestionnaireTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**QuestionnaireTypeRetrieveListArguments**](QuestionnaireTypeRetrieveListArguments.md) | QuestionnaireType RetrieveList Arguments | 

### Return type

[**QuestionnaireTypeListResponse**](QuestionnaireTypeListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveQuestionnaireType

> QuestionnaireTypeModel RetrieveQuestionnaireType(ctx, questionnaireTypeId).Execute()

Retrieve a QuestionnaireType



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
    questionnaireTypeId := int64(789) // int64 | ID of QuestionnaireType to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.QuestionnaireTypeApi.RetrieveQuestionnaireType(context.Background(), questionnaireTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireTypeApi.RetrieveQuestionnaireType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveQuestionnaireType`: QuestionnaireTypeModel
    fmt.Fprintf(os.Stdout, "Response from `QuestionnaireTypeApi.RetrieveQuestionnaireType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**questionnaireTypeId** | **int64** | ID of QuestionnaireType to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveQuestionnaireTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireTypeModel**](QuestionnaireTypeModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

