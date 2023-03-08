# \AddressApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAddress**](AddressApi.md#DeleteAddress) | **Delete** /address/{addressId} | Delete an address
[**RetrieveAddress**](AddressApi.md#RetrieveAddress) | **Get** /address/{addressId} | Retrieve a Address
[**RetrieveListAddress**](AddressApi.md#RetrieveListAddress) | **Put** /address | Retrieve List of Addresses
[**ReverseGeoCodeAddress**](AddressApi.md#ReverseGeoCodeAddress) | **Post** /address/reverse-geo-code | Reverse Geo Code an address
[**SetAddress**](AddressApi.md#SetAddress) | **Post** /address/set | Add a new Address
[**SuggestAddress**](AddressApi.md#SuggestAddress) | **Post** /address/suggest-address | Suggest an address
[**UpdateAddress**](AddressApi.md#UpdateAddress) | **Put** /address/{addressId} | Update a address
[**ValidateAddress**](AddressApi.md#ValidateAddress) | **Get** /address/validate | Validate an address



## DeleteAddress

> ApiResponse DeleteAddress(ctx, addressId).Execute()

Delete an address



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
    addressId := int64(789) // int64 | ID of the address to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.DeleteAddress(context.Background(), addressId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.DeleteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAddress`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.DeleteAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int64** | ID of the address to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddressRequest struct via the builder pattern


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


## RetrieveAddress

> AddressModel RetrieveAddress(ctx, addressId).Execute()

Retrieve a Address



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
    addressId := int64(789) // int64 | ID of address to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.RetrieveAddress(context.Background(), addressId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.RetrieveAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAddress`: AddressModel
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.RetrieveAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int64** | ID of address to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressModel**](AddressModel.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveListAddress

> AddressListResponse RetrieveListAddress(ctx).Arguments(arguments).Execute()

Retrieve List of Addresses



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
    arguments := *openapiclient.NewAddressRetrieveListArguments() // AddressRetrieveListArguments | Address RetrieveList Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.RetrieveListAddress(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.RetrieveListAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveListAddress`: AddressListResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.RetrieveListAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveListAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**AddressRetrieveListArguments**](AddressRetrieveListArguments.md) | Address RetrieveList Arguments | 

### Return type

[**AddressListResponse**](AddressListResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReverseGeoCodeAddress

> AddressReverseGeoCodeResponse ReverseGeoCodeAddress(ctx).Arguments(arguments).Execute()

Reverse Geo Code an address



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
    arguments := *openapiclient.NewAddressReverseGeoCodeArguments() // AddressReverseGeoCodeArguments | Address Reverse GeoCode Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.ReverseGeoCodeAddress(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.ReverseGeoCodeAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReverseGeoCodeAddress`: AddressReverseGeoCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.ReverseGeoCodeAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReverseGeoCodeAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**AddressReverseGeoCodeArguments**](AddressReverseGeoCodeArguments.md) | Address Reverse GeoCode Arguments | 

### Return type

[**AddressReverseGeoCodeResponse**](AddressReverseGeoCodeResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAddress

> ApiResponse SetAddress(ctx).Body(body).Execute()

Add a new Address



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
    body := *openapiclient.NewAddressModel() // AddressModel | Address object that needs to be created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.SetAddress(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.SetAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAddress`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.SetAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AddressModel**](AddressModel.md) | Address object that needs to be created | 

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


## SuggestAddress

> AddressSuggestionResponse SuggestAddress(ctx).Arguments(arguments).Execute()

Suggest an address



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
    arguments := *openapiclient.NewAddressSuggestionArguments() // AddressSuggestionArguments | Address Suggestion Arguments

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.SuggestAddress(context.Background()).Arguments(arguments).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.SuggestAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuggestAddress`: AddressSuggestionResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.SuggestAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuggestAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arguments** | [**AddressSuggestionArguments**](AddressSuggestionArguments.md) | Address Suggestion Arguments | 

### Return type

[**AddressSuggestionResponse**](AddressSuggestionResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddress

> ApiResponse UpdateAddress(ctx, addressId).Body(body).Execute()

Update a address



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
    addressId := int64(789) // int64 | ID of address to update
    body := *openapiclient.NewAddressModel() // AddressModel | Address object that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.UpdateAddress(context.Background(), addressId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.UpdateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAddress`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.UpdateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **int64** | ID of address to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AddressModel**](AddressModel.md) | Address object that needs to be updated | 

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


## ValidateAddress

> AddressValidationResponse ValidateAddress(ctx).City(city).IsoCountry(isoCountry).Street(street).HouseNr(houseNr).HouseNrAddendum(houseNrAddendum).Zipcode(zipcode).MinimumCertainty(minimumCertainty).Execute()

Validate an address



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
    city := "city_example" // string | City
    isoCountry := "isoCountry_example" // string | Country in ISO 3166-1 alpha 2
    street := "street_example" // string | Street (optional)
    houseNr := "houseNr_example" // string | House Number (optional)
    houseNrAddendum := "houseNrAddendum_example" // string | House Number Annex (optional)
    zipcode := "zipcode_example" // string | Zipcode (optional)
    minimumCertainty := int32(56) // int32 | Minimum required certainty as an int between 0 and 100 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.ValidateAddress(context.Background()).City(city).IsoCountry(isoCountry).Street(street).HouseNr(houseNr).HouseNrAddendum(houseNrAddendum).Zipcode(zipcode).MinimumCertainty(minimumCertainty).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.ValidateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAddress`: AddressValidationResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.ValidateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **city** | **string** | City | 
 **isoCountry** | **string** | Country in ISO 3166-1 alpha 2 | 
 **street** | **string** | Street | 
 **houseNr** | **string** | House Number | 
 **houseNrAddendum** | **string** | House Number Annex | 
 **zipcode** | **string** | Zipcode | 
 **minimumCertainty** | **int32** | Minimum required certainty as an int between 0 and 100 | 

### Return type

[**AddressValidationResponse**](AddressValidationResponse.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

