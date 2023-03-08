# \TransactionApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransaction**](TransactionApi.md#CreateTransaction) | **Post** /transaction | Add a new Transaction
[**DeleteTransaction**](TransactionApi.md#DeleteTransaction) | **Delete** /transaction/{transactionId} | Delete an Transaction entry
[**RetrieveListTransaction**](TransactionApi.md#RetrieveListTransaction) | **Put** /transaction | Retrieve List of Transaction
[**RetrieveTransaction**](TransactionApi.md#RetrieveTransaction) | **Get** /transaction/{transactionId} | Retrieve a Transaction
[**SetTransaction**](TransactionApi.md#SetTransaction) | **Post** /transaction/set | Set (create or update) a Transaction
[**Token**](TransactionApi.md#Token) | **Post** /transaction/token | get a transaction token
[**TokenHasFailed**](TransactionApi.md#TokenHasFailed) | **Post** /transaction/token-has-failed | set a transaction to failed with token
[**TokenIsCancelled**](TransactionApi.md#TokenIsCancelled) | **Post** /transaction/token-is-cancelled | set a transaction to cancelled with token
[**TokenIsPaid**](TransactionApi.md#TokenIsPaid) | **Post** /transaction/token-is-paid | set a transaction to paid with token
[**UpdateTransaction**](TransactionApi.md#UpdateTransaction) | **Put** /transaction/{transactionId} | Update a specific Transaction object



## CreateTransaction

> TransactionCreateResponse CreateTransaction(ctx).Body(body).Execute()

Add a new Transaction



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
    body := *openapiclient.NewTransactionModel() // TransactionModel | Transaction object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.CreateTransaction(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.CreateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransaction`: TransactionCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.CreateTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TransactionModel**](TransactionModel.md) | Transaction object that needs to be created | 

### Return type

[**TransactionCreateResponse**](TransactionCreateResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransaction

> TransactionDeleteResponse DeleteTransaction(ctx, transactionId).Execute()

Delete an Transaction entry



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
    transactionId := int64(789) // int64 | ID of Transaction to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.DeleteTransaction(context.Background(), transactionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.DeleteTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTransaction`: TransactionDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.DeleteTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** | ID of Transaction to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionDeleteResponse**](TransactionDeleteResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListTransaction

> TransactionListResponse RetrieveListTransaction(ctx).Arguments(arguments).Execute()

Retrieve List of Transaction



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
    arguments := *openapiclient.NewTransactionRetrieveListArguments() // TransactionRetrieveListArguments | Transaction RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.RetrieveListTransaction(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.RetrieveListTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListTransaction`: TransactionListResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.RetrieveListTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**TransactionRetrieveListArguments**](TransactionRetrieveListArguments.md) | Transaction RetrieveList Arguments | 

### Return type

[**TransactionListResponse**](TransactionListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveTransaction

> TransactionModel RetrieveTransaction(ctx, transactionId).Execute()

Retrieve a Transaction



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
    transactionId := int64(789) // int64 | ID of Transaction to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.RetrieveTransaction(context.Background(), transactionId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.RetrieveTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveTransaction`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.RetrieveTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** | ID of Transaction to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTransaction

> TransactionSetResponse SetTransaction(ctx).Body(body).Execute()

Set (create or update) a Transaction



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
    body := *openapiclient.NewTransactionModel() // TransactionModel | Transaction object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.SetTransaction(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.SetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTransaction`: TransactionSetResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.SetTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TransactionModel**](TransactionModel.md) | Transaction object | 

### Return type

[**TransactionSetResponse**](TransactionSetResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Token

> TransactionTokenResponse Token(ctx).Arguments(arguments).Execute()

get a transaction token



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
    arguments := *openapiclient.NewTransactionTokenArguments() // TransactionTokenArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.Token(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.Token``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Token`: TransactionTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.Token`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**TransactionTokenArguments**](TransactionTokenArguments.md) | Request Arguments | 

### Return type

[**TransactionTokenResponse**](TransactionTokenResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenHasFailed

> TransactionTokenHasFailedResponse TokenHasFailed(ctx).Arguments(arguments).Execute()

set a transaction to failed with token



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
    arguments := *openapiclient.NewTransactionTokenHasFailedArguments() // TransactionTokenHasFailedArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.TokenHasFailed(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.TokenHasFailed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenHasFailed`: TransactionTokenHasFailedResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.TokenHasFailed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenHasFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**TransactionTokenHasFailedArguments**](TransactionTokenHasFailedArguments.md) | Request Arguments | 

### Return type

[**TransactionTokenHasFailedResponse**](TransactionTokenHasFailedResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenIsCancelled

> TransactionTokenIsCancelledResponse TokenIsCancelled(ctx).Arguments(arguments).Execute()

set a transaction to cancelled with token



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
    arguments := *openapiclient.NewTransactionTokenIsCancelledArguments() // TransactionTokenIsCancelledArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.TokenIsCancelled(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.TokenIsCancelled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenIsCancelled`: TransactionTokenIsCancelledResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.TokenIsCancelled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenIsCancelledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**TransactionTokenIsCancelledArguments**](TransactionTokenIsCancelledArguments.md) | Request Arguments | 

### Return type

[**TransactionTokenIsCancelledResponse**](TransactionTokenIsCancelledResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenIsPaid

> TransactionTokenIsPaidResponse TokenIsPaid(ctx).Arguments(arguments).Execute()

set a transaction to paid with token



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
    arguments := *openapiclient.NewTransactionTokenIsPaidArguments() // TransactionTokenIsPaidArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.TokenIsPaid(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.TokenIsPaid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenIsPaid`: TransactionTokenIsPaidResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.TokenIsPaid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenIsPaidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**TransactionTokenIsPaidArguments**](TransactionTokenIsPaidArguments.md) | Request Arguments | 

### Return type

[**TransactionTokenIsPaidResponse**](TransactionTokenIsPaidResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransaction

> TransactionUpdateResponse UpdateTransaction(ctx, transactionId).Body(body).Execute()

Update a specific Transaction object



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
    transactionId := int64(789) // int64 | ID of the Transaction object to update
    body := *openapiclient.NewTransactionModel() // TransactionModel | Transaction object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.UpdateTransaction(context.Background(), transactionId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.UpdateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransaction`: TransactionUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.UpdateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** | ID of the Transaction object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TransactionModel**](TransactionModel.md) | Transaction object that needs to be updated | 

### Return type

[**TransactionUpdateResponse**](TransactionUpdateResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

