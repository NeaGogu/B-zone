# \MetadataApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetaData**](MetadataApi.md#CreateMetaData) | **Post** /metadata | Add a new MetaData
[**DeleteMetaData**](MetadataApi.md#DeleteMetaData) | **Delete** /metadata/{metadataId} | Delete a MetaData entry
[**RetrieveListMetaData**](MetadataApi.md#RetrieveListMetaData) | **Put** /metadata | Retrieve List of MetaData
[**RetrieveMetaData**](MetadataApi.md#RetrieveMetaData) | **Get** /metadata/{metadataId} | Retrieve a MetaData
[**SetMetaData**](MetadataApi.md#SetMetaData) | **Post** /metadata/set | Set (create or update) a MetaData record
[**UpdateMetaData**](MetadataApi.md#UpdateMetaData) | **Put** /metadata/{metadataId} | Update a specific MetaData object



## CreateMetaData

> ApiResponse CreateMetaData(ctx).Body(body).Execute()

Add a new MetaData



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
    body := *openapiclient.NewMetaDataModel() // MetaDataModel | MetaData object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.CreateMetaData(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.CreateMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetaData`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.CreateMetaData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MetaDataModel**](MetaDataModel.md) | MetaData object that needs to be created | 

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


## DeleteMetaData

> ApiResponse6 DeleteMetaData(ctx, metadataId).Execute()

Delete a MetaData entry



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
    metadataId := int64(789) // int64 | ID of MetaData to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.DeleteMetaData(context.Background(), metadataId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.DeleteMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMetaData`: ApiResponse6
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.DeleteMetaData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataId** | **int64** | ID of MetaData to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResponse6**](ApiResponse_6.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListMetaData

> MetaDataListResponse RetrieveListMetaData(ctx).Arguments(arguments).Execute()

Retrieve List of MetaData



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
    arguments := *openapiclient.NewMetaDataRetrieveListArguments() // MetaDataRetrieveListArguments | MetaData RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.RetrieveListMetaData(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.RetrieveListMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListMetaData`: MetaDataListResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.RetrieveListMetaData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**MetaDataRetrieveListArguments**](MetaDataRetrieveListArguments.md) | MetaData RetrieveList Arguments | 

### Return type

[**MetaDataListResponse**](MetaDataListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMetaData

> MetaDataModel RetrieveMetaData(ctx, metadataId).IncludeObjectTypeName(includeObjectTypeName).IncludeRecordInfo(includeRecordInfo).Execute()

Retrieve a MetaData



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
    metadataId := int64(789) // int64 | ID of MetaData to retrieve
    includeObjectTypeName := true // bool | Show the name of the object type (optional) (default to false)
    includeRecordInfo := true // bool | Show the record info (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.RetrieveMetaData(context.Background(), metadataId).IncludeObjectTypeName(includeObjectTypeName).IncludeRecordInfo(includeRecordInfo).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.RetrieveMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMetaData`: MetaDataModel
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.RetrieveMetaData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataId** | **int64** | ID of MetaData to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeObjectTypeName** | **bool** | Show the name of the object type | [default to false]
 **includeRecordInfo** | **bool** | Show the record info | [default to false]

### Return type

[**MetaDataModel**](MetaDataModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMetaData

> ApiResponse SetMetaData(ctx).Body(body).Execute()

Set (create or update) a MetaData record



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
    body := *openapiclient.NewMetaDataModel() // MetaDataModel | MetaData object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.SetMetaData(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.SetMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMetaData`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.SetMetaData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MetaDataModel**](MetaDataModel.md) | MetaData object | 

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


## UpdateMetaData

> ApiResponse5 UpdateMetaData(ctx, metadataId).Body(body).Execute()

Update a specific MetaData object



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
    metadataId := int64(789) // int64 | ID of the MetaData object to update
    body := *openapiclient.NewMetaDataModel() // MetaDataModel | MetaData object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.UpdateMetaData(context.Background(), metadataId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetaData`: ApiResponse5
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.UpdateMetaData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataId** | **int64** | ID of the MetaData object to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetaDataModel**](MetaDataModel.md) | MetaData object that needs to be updated | 

### Return type

[**ApiResponse5**](ApiResponse_5.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

