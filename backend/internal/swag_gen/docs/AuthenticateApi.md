# \AuthenticateApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateCheckToken**](AuthenticateApi.md#AuthenticateCheckToken) | **Get** /authenticate/check-token | Check a token for validity
[**AuthenticateSignIn**](AuthenticateApi.md#AuthenticateSignIn) | **Post** /authenticate/sign-in | Sign In with your user credentials
[**AuthenticateSignOut**](AuthenticateApi.md#AuthenticateSignOut) | **Get** /authenticate/sign-out | Sign out



## AuthenticateCheckToken

> ApiResponse AuthenticateCheckToken(ctx).Execute()

Check a token for validity



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticateApi.AuthenticateCheckToken(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateApi.AuthenticateCheckToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateCheckToken`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticateApi.AuthenticateCheckToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateCheckTokenRequest struct via the builder pattern


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateSignIn

> AuthenticateModel AuthenticateSignIn(ctx).Body(body).Execute()

Sign In with your user credentials



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
    body := *openapiclient.NewCredentialsModel() // CredentialsModel | Credentials object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticateApi.AuthenticateSignIn(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateApi.AuthenticateSignIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateSignIn`: AuthenticateModel
    fmt.Fprintf(os.Stdout, "Response from `AuthenticateApi.AuthenticateSignIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CredentialsModel**](CredentialsModel.md) | Credentials object | 

### Return type

[**AuthenticateModel**](AuthenticateModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateSignOut

> ApiResponse AuthenticateSignOut(ctx).Token(token).Execute()

Sign out



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
    token := "token_example" // string | Token

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticateApi.AuthenticateSignOut(context.Background()).Token(token).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticateApi.AuthenticateSignOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateSignOut`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticateApi.AuthenticateSignOut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateSignOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

