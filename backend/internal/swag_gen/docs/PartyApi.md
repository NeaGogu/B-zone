# \PartyApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateParty**](PartyApi.md#CreateParty) | **Post** /party | Create or update an Party
[**DeleteParty**](PartyApi.md#DeleteParty) | **Delete** /party/{partyId} | Delete an party
[**RetrieveListParty**](PartyApi.md#RetrieveListParty) | **Put** /party | Retrieve List of Parties
[**RetrieveParty**](PartyApi.md#RetrieveParty) | **Get** /party/{partyId} | Find party by ID
[**SetParty**](PartyApi.md#SetParty) | **Post** /party/set | Set (create or update) an Party
[**UpdateParty**](PartyApi.md#UpdateParty) | **Put** /party/{partyId} | Update a party



## CreateParty

> ApiResponse CreateParty(ctx).Body(body).Execute()

Create or update an Party



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
    body := *openapiclient.NewPartyModel() // PartyModel | Party object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartyApi.CreateParty(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PartyApi.CreateParty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateParty`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PartyApi.CreateParty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PartyModel**](PartyModel.md) | Party object | 

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


## DeleteParty

> ApiResponse DeleteParty(ctx, partyId).Execute()

Delete an party



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
    partyId := int64(789) // int64 | ID of the party to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartyApi.DeleteParty(context.Background(), partyId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PartyApi.DeleteParty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteParty`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PartyApi.DeleteParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partyId** | **int64** | ID of the party to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RetrieveListParty

> PartyListResponse RetrieveListParty(ctx).Arguments(arguments).Execute()

Retrieve List of Parties



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
    arguments := *openapiclient.NewPartyRetrieveListArguments() // PartyRetrieveListArguments | Party RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartyApi.RetrieveListParty(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PartyApi.RetrieveListParty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListParty`: PartyListResponse
    fmt.Fprintf(os.Stdout, "Response from `PartyApi.RetrieveListParty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**PartyRetrieveListArguments**](PartyRetrieveListArguments.md) | Party RetrieveList Arguments | 

### Return type

[**PartyListResponse**](PartyListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveParty

> PartyModel RetrieveParty(ctx, partyId).Execute()

Find party by ID



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
    partyId := int64(789) // int64 | ID of party to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartyApi.RetrieveParty(context.Background(), partyId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PartyApi.RetrieveParty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveParty`: PartyModel
    fmt.Fprintf(os.Stdout, "Response from `PartyApi.RetrieveParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partyId** | **int64** | ID of party to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PartyModel**](PartyModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetParty

> ApiResponse SetParty(ctx).Body(body).Execute()

Set (create or update) an Party



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
    body := *openapiclient.NewPartyModel() // PartyModel | Party object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartyApi.SetParty(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PartyApi.SetParty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetParty`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PartyApi.SetParty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PartyModel**](PartyModel.md) | Party object | 

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


## UpdateParty

> ApiResponse UpdateParty(ctx, partyId).Body(body).Execute()

Update a party



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
    partyId := int64(789) // int64 | ID of party to update
    body := *openapiclient.NewPartyModel() // PartyModel | Party object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartyApi.UpdateParty(context.Background(), partyId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PartyApi.UpdateParty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateParty`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PartyApi.UpdateParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partyId** | **int64** | ID of party to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PartyModel**](PartyModel.md) | Party object that needs to be updated | 

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

