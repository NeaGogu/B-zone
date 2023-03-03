# \PlannerApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyPlanning**](PlannerApi.md#ApplyPlanning) | **Post** /planner/apply-planning | Apply a planning schema
[**AutoPlan**](PlannerApi.md#AutoPlan) | **Post** /planner/auto-plan | Plan a certain activity in any fitting route
[**AutoPlanResult**](PlannerApi.md#AutoPlanResult) | **Post** /planner/auto-plan-result | Fetch current result for a auto plan Request. This could be done, in progress or cancelled.
[**CheckAvailability**](PlannerApi.md#CheckAvailability) | **Post** /planner/check-availability | check availability in planning for a certain set of activity properties
[**CheckAvailabilityCaching**](PlannerApi.md#CheckAvailabilityCaching) | **Post** /planner/check-availability-caching | Retrieve the cached check-availability
[**CheckAvailabilityResult**](PlannerApi.md#CheckAvailabilityResult) | **Post** /planner/check-availability-result | Fetch current result for a checkAvailability Request. This could be done, in progress or cancelled.
[**PlannerAddActivitiesToRoute**](PlannerApi.md#PlannerAddActivitiesToRoute) | **Post** /planner/add-activities-to-route | Add Activities to Route
[**RemoveActivitiesFromRoute**](PlannerApi.md#RemoveActivitiesFromRoute) | **Post** /planner/remove-activities-from-route | Remove Activities From Route
[**UpdateRecurrenceRelations**](PlannerApi.md#UpdateRecurrenceRelations) | **Post** /planner/update-recurrence-relations | Update recurrence relations for follow routes



## ApplyPlanning

> ApiResponse ApplyPlanning(ctx).Arguments(arguments).Execute()

Apply a planning schema



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
    arguments := *openapiclient.NewApplyPlanningArguments() // ApplyPlanningArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.ApplyPlanning(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.ApplyPlanning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyPlanning`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.ApplyPlanning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyPlanningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**ApplyPlanningArguments**](ApplyPlanningArguments.md) | Request Arguments | 

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


## AutoPlan

> ApiResponse AutoPlan(ctx).Arguments(arguments).Execute()

Plan a certain activity in any fitting route



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
    arguments := *openapiclient.NewAutoPlanArguments() // AutoPlanArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.AutoPlan(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.AutoPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoPlan`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.AutoPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**AutoPlanArguments**](AutoPlanArguments.md) | Request Arguments | 

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


## AutoPlanResult

> ApiResponse AutoPlanResult(ctx).Arguments(arguments).Execute()

Fetch current result for a auto plan Request. This could be done, in progress or cancelled.



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
    arguments := *openapiclient.NewAutoPlanArguments() // AutoPlanArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.AutoPlanResult(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.AutoPlanResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoPlanResult`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.AutoPlanResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoPlanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**AutoPlanArguments**](AutoPlanArguments.md) | Request Arguments | 

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


## CheckAvailability

> ApiResponse CheckAvailability(ctx).Arguments(arguments).Execute()

check availability in planning for a certain set of activity properties



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
    arguments := *openapiclient.NewCheckAvailabilityArguments() // CheckAvailabilityArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.CheckAvailability(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.CheckAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAvailability`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.CheckAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**CheckAvailabilityArguments**](CheckAvailabilityArguments.md) | Request Arguments | 

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


## CheckAvailabilityCaching

> ApiResponse CheckAvailabilityCaching(ctx).Arguments(arguments).Execute()

Retrieve the cached check-availability



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
    arguments := *openapiclient.NewCheckAvailabilityCacheArguments() // CheckAvailabilityCacheArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.CheckAvailabilityCaching(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.CheckAvailabilityCaching``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAvailabilityCaching`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.CheckAvailabilityCaching`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAvailabilityCachingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**CheckAvailabilityCacheArguments**](CheckAvailabilityCacheArguments.md) | Request Arguments | 

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


## CheckAvailabilityResult

> ApiResponse CheckAvailabilityResult(ctx).Arguments(arguments).Execute()

Fetch current result for a checkAvailability Request. This could be done, in progress or cancelled.



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
    arguments := *openapiclient.NewCheckAvailabilityArguments() // CheckAvailabilityArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.CheckAvailabilityResult(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.CheckAvailabilityResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAvailabilityResult`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.CheckAvailabilityResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAvailabilityResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**CheckAvailabilityArguments**](CheckAvailabilityArguments.md) | Request Arguments | 

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


## PlannerAddActivitiesToRoute

> AddActivitiesToRouteResponse PlannerAddActivitiesToRoute(ctx).Arguments(arguments).Execute()

Add Activities to Route



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
    arguments := *openapiclient.NewAddActivitiesToRouteArguments() // AddActivitiesToRouteArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.PlannerAddActivitiesToRoute(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.PlannerAddActivitiesToRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerAddActivitiesToRoute`: AddActivitiesToRouteResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.PlannerAddActivitiesToRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlannerAddActivitiesToRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**AddActivitiesToRouteArguments**](AddActivitiesToRouteArguments.md) | Request Arguments | 

### Return type

[**AddActivitiesToRouteResponse**](AddActivitiesToRouteResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveActivitiesFromRoute

> ApiResponse RemoveActivitiesFromRoute(ctx).Arguments(arguments).Execute()

Remove Activities From Route



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
    arguments := *openapiclient.NewRemoveActivitiesFromRouteArguments() // RemoveActivitiesFromRouteArguments | Request Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.RemoveActivitiesFromRoute(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.RemoveActivitiesFromRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveActivitiesFromRoute`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.RemoveActivitiesFromRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveActivitiesFromRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**RemoveActivitiesFromRouteArguments**](RemoveActivitiesFromRouteArguments.md) | Request Arguments | 

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


## UpdateRecurrenceRelations

> ApiResponse UpdateRecurrenceRelations(ctx).Arguments(arguments).Execute()

Update recurrence relations for follow routes



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
    arguments := *openapiclient.NewUpdateRecurrenceRelations() // UpdateRecurrenceRelations | Update recurrence relations arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerApi.UpdateRecurrenceRelations(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerApi.UpdateRecurrenceRelations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecurrenceRelations`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `PlannerApi.UpdateRecurrenceRelations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecurrenceRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**UpdateRecurrenceRelations**](UpdateRecurrenceRelations.md) | Update recurrence relations arguments | 

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

