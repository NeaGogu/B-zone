# \PaymentApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayment**](PaymentApi.md#CreatePayment) | **Post** /payment | Add a new Payment
[**DeletePayment**](PaymentApi.md#DeletePayment) | **Delete** /payment/{paymentId} | Delete an Payment entry
[**RetrieveListPayment**](PaymentApi.md#RetrieveListPayment) | **Put** /payment | Retrieve List of Payment
[**RetrievePayment**](PaymentApi.md#RetrievePayment) | **Get** /payment/{paymentId} | Retrieve a Payment
[**SetPayment**](PaymentApi.md#SetPayment) | **Post** /payment/set | Set (create or update) a Payment
[**UpdatePayment**](PaymentApi.md#UpdatePayment) | **Put** /payment/{paymentId} | Update a specific Payment object



## CreatePayment

> PaymentCreateResponse CreatePayment(ctx).Body(body).Execute()

Add a new Payment



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
    body := *openapiclient.NewPaymentModel() // PaymentModel | Payment object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.CreatePayment(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.CreatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayment`: PaymentCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.CreatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentModel**](PaymentModel.md) | Payment object that needs to be created | 

### Return type

[**PaymentCreateResponse**](PaymentCreateResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePayment

> PaymentDeleteResponse DeletePayment(ctx, paymentId).Execute()

Delete an Payment entry



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
    paymentId := int64(789) // int64 | ID of Payment to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.DeletePayment(context.Background(), paymentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.DeletePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePayment`: PaymentDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.DeletePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **int64** | ID of Payment to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentDeleteResponse**](PaymentDeleteResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListPayment

> PaymentListResponse RetrieveListPayment(ctx).Arguments(arguments).Execute()

Retrieve List of Payment



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
    arguments := *openapiclient.NewPaymentRetrieveListArguments() // PaymentRetrieveListArguments | Payment RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.RetrieveListPayment(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.RetrieveListPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListPayment`: PaymentListResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.RetrieveListPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**PaymentRetrieveListArguments**](PaymentRetrieveListArguments.md) | Payment RetrieveList Arguments | 

### Return type

[**PaymentListResponse**](PaymentListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePayment

> PaymentModel RetrievePayment(ctx, paymentId).Execute()

Retrieve a Payment



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
    paymentId := int64(789) // int64 | ID of Payment to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.RetrievePayment(context.Background(), paymentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.RetrievePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePayment`: PaymentModel
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.RetrievePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **int64** | ID of Payment to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentModel**](PaymentModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPayment

> PaymentSetResponse SetPayment(ctx).Body(body).Execute()

Set (create or update) a Payment



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
    body := *openapiclient.NewPaymentModel() // PaymentModel | Payment object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.SetPayment(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.SetPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPayment`: PaymentSetResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.SetPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaymentModel**](PaymentModel.md) | Payment object | 

### Return type

[**PaymentSetResponse**](PaymentSetResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePayment

> PaymentUpdateResponse UpdatePayment(ctx, paymentId).Body(body).Execute()

Update a specific Payment object



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
    paymentId := int64(789) // int64 | ID of the Payment object to update
    body := *openapiclient.NewPaymentModel() // PaymentModel | Payment object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.UpdatePayment(context.Background(), paymentId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.UpdatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePayment`: PaymentUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.UpdatePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **int64** | ID of the Payment object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PaymentModel**](PaymentModel.md) | Payment object that needs to be updated | 

### Return type

[**PaymentUpdateResponse**](PaymentUpdateResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

