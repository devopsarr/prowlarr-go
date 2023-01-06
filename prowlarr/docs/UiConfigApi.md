# \UiConfigApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1ConfigUi**](UiConfigApi.md#GetApiV1ConfigUi) | **Get** /api/v1/config/ui | 
[**GetApiV1ConfigUiById**](UiConfigApi.md#GetApiV1ConfigUiById) | **Get** /api/v1/config/ui/{id} | 
[**UpdateApiV1ConfigUi**](UiConfigApi.md#UpdateApiV1ConfigUi) | **Put** /api/v1/config/ui/{id} | 



## GetApiV1ConfigUi

> UiConfigResource GetApiV1ConfigUi(ctx).Execute()



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
    resp, r, err := apiClient.UiConfigApi.GetApiV1ConfigUi(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiConfigApi.GetApiV1ConfigUi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigUi`: UiConfigResource
    fmt.Fprintf(os.Stdout, "Response from `UiConfigApi.GetApiV1ConfigUi`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigUiRequest struct via the builder pattern


### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1ConfigUiById

> UiConfigResource GetApiV1ConfigUiById(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UiConfigApi.GetApiV1ConfigUiById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiConfigApi.GetApiV1ConfigUiById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigUiById`: UiConfigResource
    fmt.Fprintf(os.Stdout, "Response from `UiConfigApi.GetApiV1ConfigUiById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigUiByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1ConfigUi

> UiConfigResource UpdateApiV1ConfigUi(ctx, id).UiConfigResource(uiConfigResource).Execute()



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
    id := "id_example" // string | 
    uiConfigResource := *prowlarrClient.NewUiConfigResource() // UiConfigResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UiConfigApi.UpdateApiV1ConfigUi(context.Background(), id).UiConfigResource(uiConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiConfigApi.UpdateApiV1ConfigUi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ConfigUi`: UiConfigResource
    fmt.Fprintf(os.Stdout, "Response from `UiConfigApi.UpdateApiV1ConfigUi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ConfigUiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uiConfigResource** | [**UiConfigResource**](UiConfigResource.md) |  | 

### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

