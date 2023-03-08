# \PackageLineApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePackageLine**](PackageLineApi.md#CreatePackageLine) | **Post** /package-line | Create or update an Package Line
[**DeletePackageLine**](PackageLineApi.md#DeletePackageLine) | **Delete** /package-line/{packageLineId} | Delete an package-line
[**RetrieveListPackageLine**](PackageLineApi.md#RetrieveListPackageLine) | **Put** /package-line | Retrieve List of PackageLines
[**RetrievePackageLine**](PackageLineApi.md#RetrievePackageLine) | **Get** /package-line/{packageLineId} | Find package-line by ID
[**SetPackageLine**](PackageLineApi.md#SetPackageLine) | **Post** /package-line/set | Set (create or update) an PackageLine
[**UpdatePackageLine**](PackageLineApi.md#UpdatePackageLine) | **Put** /package-line/update | Update package-lines in bulk



## CreatePackageLine

> ApiResponse CreatePackageLine(ctx).Body(body).Execute()

Create or update an Package Line



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
    body := *openapiclient.NewPackageLineModel() // PackageLineModel | PackageLine model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageLineApi.CreatePackageLine(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageLineApi.CreatePackageLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePackageLine`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageLineApi.CreatePackageLine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePackageLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PackageLineModel**](PackageLineModel.md) | PackageLine model | 

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


## DeletePackageLine

> ApiResponse DeletePackageLine(ctx, packageLineId).Execute()

Delete an package-line



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
    packageLineId := int64(789) // int64 | ID of the package-line to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageLineApi.DeletePackageLine(context.Background(), packageLineId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageLineApi.DeletePackageLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePackageLine`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageLineApi.DeletePackageLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageLineId** | **int64** | ID of the package-line to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePackageLineRequest struct via the builder pattern


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


## RetrieveListPackageLine

> PackageLineListResponse RetrieveListPackageLine(ctx).Arguments(arguments).Execute()

Retrieve List of PackageLines



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
    arguments := *openapiclient.NewPackageLineRetrieveListArguments() // PackageLineRetrieveListArguments | PackageLine RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageLineApi.RetrieveListPackageLine(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageLineApi.RetrieveListPackageLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListPackageLine`: PackageLineListResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageLineApi.RetrieveListPackageLine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListPackageLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**PackageLineRetrieveListArguments**](PackageLineRetrieveListArguments.md) | PackageLine RetrieveList Arguments | 

### Return type

[**PackageLineListResponse**](PackageLineListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePackageLine

> PackageLineModel RetrievePackageLine(ctx, packageLineId).IncludePackageLineStatus(includePackageLineStatus).IncludePackageLineTypeName(includePackageLineTypeName).IncludePackageLineMetaData(includePackageLineMetaData).IncludeAddressObject(includeAddressObject).IncludeTimeSlots(includeTimeSlots).IncludeTimeSlotTags(includeTimeSlotTags).IncludeRouteInfo(includeRouteInfo).IncludeDriverInfo(includeDriverInfo).IncludeCommunication(includeCommunication).IncludePackageLineLinks(includePackageLineLinks).IncludePackageLinesInfo(includePackageLinesInfo).IncludePackageLineFiles(includePackageLineFiles).IncludePackageLineFilesMetaData(includePackageLineFilesMetaData).Execute()

Find package-line by ID



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
    packageLineId := int64(789) // int64 | ID of package-line to return
    includePackageLineStatus := true // bool | Show the text value of the status (default to true)
    includePackageLineTypeName := true // bool | Show the text value of the package-line type (default to true)
    includePackageLineMetaData := true // bool | Include meta data connected to this PackageLine (default to true)
    includeAddressObject := true // bool | Include address data (default to true)
    includeTimeSlots := true // bool | Include TimeSlots (default to true)
    includeTimeSlotTags := true // bool | Include tags from TimeSlots (default to true)
    includeRouteInfo := true // bool | Include route data (default to true)
    includeDriverInfo := true // bool | Include driver data (default to true)
    includeCommunication := true // bool | Include Communication Settings (default to true)
    includePackageLineLinks := true // bool | Include Link Data (default to true)
    includePackageLinesInfo := true // bool | Include PackageLines (default to true)
    includePackageLineFiles := true // bool | Include files (default to true)
    includePackageLineFilesMetaData := true // bool | Include files meta data (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageLineApi.RetrievePackageLine(context.Background(), packageLineId).IncludePackageLineStatus(includePackageLineStatus).IncludePackageLineTypeName(includePackageLineTypeName).IncludePackageLineMetaData(includePackageLineMetaData).IncludeAddressObject(includeAddressObject).IncludeTimeSlots(includeTimeSlots).IncludeTimeSlotTags(includeTimeSlotTags).IncludeRouteInfo(includeRouteInfo).IncludeDriverInfo(includeDriverInfo).IncludeCommunication(includeCommunication).IncludePackageLineLinks(includePackageLineLinks).IncludePackageLinesInfo(includePackageLinesInfo).IncludePackageLineFiles(includePackageLineFiles).IncludePackageLineFilesMetaData(includePackageLineFilesMetaData).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageLineApi.RetrievePackageLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePackageLine`: PackageLineModel
    fmt.Fprintf(os.Stdout, "Response from `PackageLineApi.RetrievePackageLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageLineId** | **int64** | ID of package-line to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePackageLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePackageLineStatus** | **bool** | Show the text value of the status | [default to true]
 **includePackageLineTypeName** | **bool** | Show the text value of the package-line type | [default to true]
 **includePackageLineMetaData** | **bool** | Include meta data connected to this PackageLine | [default to true]
 **includeAddressObject** | **bool** | Include address data | [default to true]
 **includeTimeSlots** | **bool** | Include TimeSlots | [default to true]
 **includeTimeSlotTags** | **bool** | Include tags from TimeSlots | [default to true]
 **includeRouteInfo** | **bool** | Include route data | [default to true]
 **includeDriverInfo** | **bool** | Include driver data | [default to true]
 **includeCommunication** | **bool** | Include Communication Settings | [default to true]
 **includePackageLineLinks** | **bool** | Include Link Data | [default to true]
 **includePackageLinesInfo** | **bool** | Include PackageLines | [default to true]
 **includePackageLineFiles** | **bool** | Include files | [default to true]
 **includePackageLineFilesMetaData** | **bool** | Include files meta data | [default to true]

### Return type

[**PackageLineModel**](PackageLineModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPackageLine

> ApiResponse SetPackageLine(ctx).Body(body).Execute()

Set (create or update) an PackageLine



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
    body := *openapiclient.NewPackageLineModel() // PackageLineModel | PackageLine model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageLineApi.SetPackageLine(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageLineApi.SetPackageLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPackageLine`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageLineApi.SetPackageLine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPackageLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PackageLineModel**](PackageLineModel.md) | PackageLine model | 

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


## UpdatePackageLine

> ApiResponse UpdatePackageLine(ctx).Body(body).Execute()

Update package-lines in bulk



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
    body := *openapiclient.NewPackageLineUpdateArguments() // PackageLineUpdateArguments | PackageLine Update object that contains all information about this update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackageLineApi.UpdatePackageLine(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PackageLineApi.UpdatePackageLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePackageLine`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PackageLineApi.UpdatePackageLine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePackageLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PackageLineUpdateArguments**](PackageLineUpdateArguments.md) | PackageLine Update object that contains all information about this update | 

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

