# \LocalizationAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocalization**](LocalizationAPI.md#GetLocalization) | **Get** /api/v1/localization | 
[**ListLocalizationOptions**](LocalizationAPI.md#ListLocalizationOptions) | **Get** /api/v1/localization/options | 



## GetLocalization

> GetLocalization(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    prowlarrClient "./openapi"
)

func main() {

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalizationAPI.GetLocalization(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.GetLocalization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocalizationOptions

> []LocalizationOption ListLocalizationOptions(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    prowlarrClient "./openapi"
)

func main() {

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalizationAPI.ListLocalizationOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.ListLocalizationOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocalizationOptions`: []LocalizationOption
    fmt.Fprintf(os.Stdout, "Response from `LocalizationAPI.ListLocalizationOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLocalizationOptionsRequest struct via the builder pattern


### Return type

[**[]LocalizationOption**](LocalizationOption.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

