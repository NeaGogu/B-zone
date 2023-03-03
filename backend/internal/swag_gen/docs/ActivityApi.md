# \ActivityApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateActivity**](ActivityApi.md#BulkUpdateActivity) | **Post** /activity/bulk-update | Update multiple activities
[**DeleteActivity**](ActivityApi.md#DeleteActivity) | **Delete** /activity/{activityId} | Delete an activity
[**LockActivity**](ActivityApi.md#LockActivity) | **Post** /activity/lock | Lock activities which satisfy set filters
[**LockActivityOnRoute**](ActivityApi.md#LockActivityOnRoute) | **Post** /activity/lock-on-route | Lock activities on route which satisfy set filters
[**LockActivityOnRouteAndTime**](ActivityApi.md#LockActivityOnRouteAndTime) | **Post** /activity/lock-on-route-and-time | Lock activities on route and time which satisfy set filters
[**RetrieveActivity**](ActivityApi.md#RetrieveActivity) | **Get** /activity/{activityId} | Find activity by ID
[**RetrieveListActivity**](ActivityApi.md#RetrieveListActivity) | **Put** /activity | Retrieve List of Activities
[**SetActivity**](ActivityApi.md#SetActivity) | **Post** /activity/set | Set (create or update) an Activity
[**UnlockActivity**](ActivityApi.md#UnlockActivity) | **Post** /activity/unlock | Unlock activities which satisfy set filters
[**Unsuccessful**](ActivityApi.md#Unsuccessful) | **Post** /activity/unsuccessful | Report an unsuccessful activity
[**UpdateActivity**](ActivityApi.md#UpdateActivity) | **Put** /activity/{activityId} | Update a activity



## BulkUpdateActivity

> ApiResponse BulkUpdateActivity(ctx).Body(body).Execute()

Update multiple activities



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
    body := *openapiclient.NewActivityBulkUpdateArguments() // ActivityBulkUpdateArguments | Activity object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.BulkUpdateActivity(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.BulkUpdateActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateActivity`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.BulkUpdateActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ActivityBulkUpdateArguments**](ActivityBulkUpdateArguments.md) | Activity object | 

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


## DeleteActivity

> ApiResponse DeleteActivity(ctx, activityId).Execute()

Delete an activity



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
    activityId := int64(789) // int64 | ID of the activity to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.DeleteActivity(context.Background(), activityId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.DeleteActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteActivity`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.DeleteActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityId** | **int64** | ID of the activity to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActivityRequest struct via the builder pattern


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


## LockActivity

> ApiResponse LockActivity(ctx).Filters(filters).Execute()

Lock activities which satisfy set filters



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
    filters := *openapiclient.NewActivityFiltersModel() // ActivityFiltersModel | Request Filters

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.LockActivity(context.Background()).Filters(filters).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.LockActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockActivity`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.LockActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLockActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**ActivityFiltersModel**](ActivityFiltersModel.md) | Request Filters | 

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


## LockActivityOnRoute

> ApiResponse LockActivityOnRoute(ctx).Filters(filters).Execute()

Lock activities on route which satisfy set filters



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
    filters := *openapiclient.NewActivityFiltersModel() // ActivityFiltersModel | Request Filters

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.LockActivityOnRoute(context.Background()).Filters(filters).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.LockActivityOnRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockActivityOnRoute`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.LockActivityOnRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLockActivityOnRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**ActivityFiltersModel**](ActivityFiltersModel.md) | Request Filters | 

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


## LockActivityOnRouteAndTime

> ApiResponse LockActivityOnRouteAndTime(ctx).Filters(filters).Execute()

Lock activities on route and time which satisfy set filters



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
    filters := *openapiclient.NewActivityFiltersModel() // ActivityFiltersModel | Request Filters

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.LockActivityOnRouteAndTime(context.Background()).Filters(filters).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.LockActivityOnRouteAndTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockActivityOnRouteAndTime`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.LockActivityOnRouteAndTime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLockActivityOnRouteAndTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**ActivityFiltersModel**](ActivityFiltersModel.md) | Request Filters | 

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


## RetrieveActivity

> ActivityModel RetrieveActivity(ctx, activityId).IncludeActivityStatus(includeActivityStatus).IncludeActivityTypeName(includeActivityTypeName).IncludeActivityMetaData(includeActivityMetaData).IncludeAddressObject(includeAddressObject).IncludeTimeSlots(includeTimeSlots).IncludeRouteInfo(includeRouteInfo).IncludeRoute(includeRoute).IncludePackageLines(includePackageLines).IncludePackageLinesInfo(includePackageLinesInfo).IncludeDriverInfo(includeDriverInfo).IncludeCommunication(includeCommunication).IncludeCommunicationObject(includeCommunicationObject).IncludeActivityLinks(includeActivityLinks).IncludeActivityFiles(includeActivityFiles).IncludeActivityFilesMetaData(includeActivityFilesMetaData).IncludeAssignmentNr(includeAssignmentNr).IncludeAssignment(includeAssignment).IncludeActivityTags(includeActivityTags).IncludeTagTypeName(includeTagTypeName).IncludeActivityRecordInfo(includeActivityRecordInfo).IncludeActivityNotes(includeActivityNotes).IncludeActivityNoteTags(includeActivityNoteTags).IncludeDepotAddressObject(includeDepotAddressObject).IncludeCapacityObject(includeCapacityObject).IncludeZones(includeZones).IncludeBrand(includeBrand).IncludeBrandColours(includeBrandColours).IncludeBrandFiles(includeBrandFiles).IncludeRelations(includeRelations).Execute()

Find activity by ID



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
    activityId := int64(789) // int64 | ID of activity to return
    includeActivityStatus := true // bool | Show the text value of the status (default to true)
    includeActivityTypeName := true // bool | Show the text value of the activity type (default to true)
    includeActivityMetaData := true // bool | Include meta data connected to this Activity (default to true)
    includeAddressObject := true // bool | Include address data (default to true)
    includeTimeSlots := true // bool | Include TimeSlots (default to true)
    includeRouteInfo := true // bool | Include route data (default to true)
    includeRoute := true // bool | Include Route (default to true)
    includePackageLines := true // bool | Include package lines (default to true)
    includePackageLinesInfo := true // bool | Include PackageLines (default to true)
    includeDriverInfo := true // bool | Include driver data (default to true)
    includeCommunication := true // bool | Include Communication Settings (default to true)
    includeCommunicationObject := true // bool | Include Communication Object (default to true)
    includeActivityLinks := true // bool | Include Link Data (default to true)
    includeActivityFiles := true // bool | Include files (default to true)
    includeActivityFilesMetaData := true // bool | Include files meta data (default to true)
    includeAssignmentNr := true // bool | Include Assignment Nr (default to true)
    includeAssignment := true // bool | Include Assignment (default to true)
    includeActivityTags := true // bool | Include Activity Tags (default to true)
    includeTagTypeName := true // bool | Include Activity Tag type names (default to true)
    includeActivityRecordInfo := true // bool | Include Activity Info (default to true)
    includeActivityNotes := true // bool | Include Activity Notes (default to true)
    includeActivityNoteTags := true // bool | Include Activity Note Tags (default to true)
    includeDepotAddressObject := true // bool | Include Depot Address Object (default to true)
    includeCapacityObject := true // bool | Include Capacity Object (default to true)
    includeZones := true // bool | Include Zones (default to true)
    includeBrand := true // bool | Include brand (default to true)
    includeBrandColours := true // bool | Include brand colours (default to true)
    includeBrandFiles := true // bool | Include brand files (default to true)
    includeRelations := true // bool | Include activity_before and activity_after (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.RetrieveActivity(context.Background(), activityId).IncludeActivityStatus(includeActivityStatus).IncludeActivityTypeName(includeActivityTypeName).IncludeActivityMetaData(includeActivityMetaData).IncludeAddressObject(includeAddressObject).IncludeTimeSlots(includeTimeSlots).IncludeRouteInfo(includeRouteInfo).IncludeRoute(includeRoute).IncludePackageLines(includePackageLines).IncludePackageLinesInfo(includePackageLinesInfo).IncludeDriverInfo(includeDriverInfo).IncludeCommunication(includeCommunication).IncludeCommunicationObject(includeCommunicationObject).IncludeActivityLinks(includeActivityLinks).IncludeActivityFiles(includeActivityFiles).IncludeActivityFilesMetaData(includeActivityFilesMetaData).IncludeAssignmentNr(includeAssignmentNr).IncludeAssignment(includeAssignment).IncludeActivityTags(includeActivityTags).IncludeTagTypeName(includeTagTypeName).IncludeActivityRecordInfo(includeActivityRecordInfo).IncludeActivityNotes(includeActivityNotes).IncludeActivityNoteTags(includeActivityNoteTags).IncludeDepotAddressObject(includeDepotAddressObject).IncludeCapacityObject(includeCapacityObject).IncludeZones(includeZones).IncludeBrand(includeBrand).IncludeBrandColours(includeBrandColours).IncludeBrandFiles(includeBrandFiles).IncludeRelations(includeRelations).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.RetrieveActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveActivity`: ActivityModel
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.RetrieveActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityId** | **int64** | ID of activity to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeActivityStatus** | **bool** | Show the text value of the status | [default to true]
 **includeActivityTypeName** | **bool** | Show the text value of the activity type | [default to true]
 **includeActivityMetaData** | **bool** | Include meta data connected to this Activity | [default to true]
 **includeAddressObject** | **bool** | Include address data | [default to true]
 **includeTimeSlots** | **bool** | Include TimeSlots | [default to true]
 **includeRouteInfo** | **bool** | Include route data | [default to true]
 **includeRoute** | **bool** | Include Route | [default to true]
 **includePackageLines** | **bool** | Include package lines | [default to true]
 **includePackageLinesInfo** | **bool** | Include PackageLines | [default to true]
 **includeDriverInfo** | **bool** | Include driver data | [default to true]
 **includeCommunication** | **bool** | Include Communication Settings | [default to true]
 **includeCommunicationObject** | **bool** | Include Communication Object | [default to true]
 **includeActivityLinks** | **bool** | Include Link Data | [default to true]
 **includeActivityFiles** | **bool** | Include files | [default to true]
 **includeActivityFilesMetaData** | **bool** | Include files meta data | [default to true]
 **includeAssignmentNr** | **bool** | Include Assignment Nr | [default to true]
 **includeAssignment** | **bool** | Include Assignment | [default to true]
 **includeActivityTags** | **bool** | Include Activity Tags | [default to true]
 **includeTagTypeName** | **bool** | Include Activity Tag type names | [default to true]
 **includeActivityRecordInfo** | **bool** | Include Activity Info | [default to true]
 **includeActivityNotes** | **bool** | Include Activity Notes | [default to true]
 **includeActivityNoteTags** | **bool** | Include Activity Note Tags | [default to true]
 **includeDepotAddressObject** | **bool** | Include Depot Address Object | [default to true]
 **includeCapacityObject** | **bool** | Include Capacity Object | [default to true]
 **includeZones** | **bool** | Include Zones | [default to true]
 **includeBrand** | **bool** | Include brand | [default to true]
 **includeBrandColours** | **bool** | Include brand colours | [default to true]
 **includeBrandFiles** | **bool** | Include brand files | [default to true]
 **includeRelations** | **bool** | Include activity_before and activity_after | [default to true]

### Return type

[**ActivityModel**](ActivityModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListActivity

> ActivityListResponse RetrieveListActivity(ctx).Arguments(arguments).Execute()

Retrieve List of Activities



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
    arguments := *openapiclient.NewActivityRetrieveListArguments() // ActivityRetrieveListArguments | Activity RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.RetrieveListActivity(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.RetrieveListActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListActivity`: ActivityListResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.RetrieveListActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**ActivityRetrieveListArguments**](ActivityRetrieveListArguments.md) | Activity RetrieveList Arguments | 

### Return type

[**ActivityListResponse**](ActivityListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetActivity

> ApiResponse SetActivity(ctx).Body(body).Execute()

Set (create or update) an Activity



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
    body := *openapiclient.NewActivityModel() // ActivityModel | Activity object (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.SetActivity(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.SetActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetActivity`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.SetActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ActivityModel**](ActivityModel.md) | Activity object | 

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


## UnlockActivity

> ApiResponse UnlockActivity(ctx).Filters(filters).Execute()

Unlock activities which satisfy set filters



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
    filters := *openapiclient.NewActivityFiltersModel() // ActivityFiltersModel | Request Filters

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.UnlockActivity(context.Background()).Filters(filters).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.UnlockActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnlockActivity`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.UnlockActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlockActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | [**ActivityFiltersModel**](ActivityFiltersModel.md) | Request Filters | 

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


## Unsuccessful

> UnsuccessfulResponseModel Unsuccessful(ctx).Arguments(arguments).Execute()

Report an unsuccessful activity



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
    arguments := *openapiclient.NewUnsuccessfulModel() // UnsuccessfulModel | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.Unsuccessful(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.Unsuccessful``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Unsuccessful`: UnsuccessfulResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.Unsuccessful`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsuccessfulRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**UnsuccessfulModel**](UnsuccessfulModel.md) | Request Arguments | 

### Return type

[**UnsuccessfulResponseModel**](UnsuccessfulResponseModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActivity

> ApiResponse UpdateActivity(ctx, activityId).Body(body).Execute()

Update a activity



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
    activityId := int64(789) // int64 | ID of activity to update
    body := *openapiclient.NewActivityModel() // ActivityModel | Activity object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivityApi.UpdateActivity(context.Background(), activityId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.UpdateActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateActivity`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.UpdateActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityId** | **int64** | ID of activity to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ActivityModel**](ActivityModel.md) | Activity object that needs to be updated | 

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

