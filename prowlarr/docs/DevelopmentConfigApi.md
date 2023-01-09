# \DevelopmentConfigApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDevelopmentConfig**](DevelopmentConfigApi.md#GetDevelopmentConfig) | **Get** /api/v1/config/development | 
[**GetDevelopmentConfigById**](DevelopmentConfigApi.md#GetDevelopmentConfigById) | **Get** /api/v1/config/development/{id} | 
[**UpdateDevelopmentConfig**](DevelopmentConfigApi.md#UpdateDevelopmentConfig) | **Put** /api/v1/config/development/{id} | 



## GetDevelopmentConfig

> DevelopmentConfigResource GetDevelopmentConfig(ctx).Execute()



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
    resp, r, err := apiClient.DevelopmentConfigApi.GetDevelopmentConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigApi.GetDevelopmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevelopmentConfig`: DevelopmentConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigApi.GetDevelopmentConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevelopmentConfigRequest struct via the builder pattern


### Return type

[**DevelopmentConfigResource**](DevelopmentConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevelopmentConfigById

> DevelopmentConfigResource GetDevelopmentConfigById(ctx, id).Execute()



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
    resp, r, err := apiClient.DevelopmentConfigApi.GetDevelopmentConfigById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigApi.GetDevelopmentConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevelopmentConfigById`: DevelopmentConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigApi.GetDevelopmentConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevelopmentConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DevelopmentConfigResource**](DevelopmentConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevelopmentConfig

> DevelopmentConfigResource UpdateDevelopmentConfig(ctx, id).DevelopmentConfigResource(developmentConfigResource).Execute()



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
    developmentConfigResource := *prowlarrClient.NewDevelopmentConfigResource() // DevelopmentConfigResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevelopmentConfigApi.UpdateDevelopmentConfig(context.Background(), id).DevelopmentConfigResource(developmentConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigApi.UpdateDevelopmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevelopmentConfig`: DevelopmentConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigApi.UpdateDevelopmentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDevelopmentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **developmentConfigResource** | [**DevelopmentConfigResource**](DevelopmentConfigResource.md) |  | 

### Return type

[**DevelopmentConfigResource**](DevelopmentConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

