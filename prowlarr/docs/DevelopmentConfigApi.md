# \DevelopmentConfigApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1ConfigDevelopment**](DevelopmentConfigApi.md#GetApiV1ConfigDevelopment) | **Get** /api/v1/config/development | 
[**GetApiV1ConfigDevelopmentById**](DevelopmentConfigApi.md#GetApiV1ConfigDevelopmentById) | **Get** /api/v1/config/development/{id} | 
[**UpdateApiV1ConfigDevelopment**](DevelopmentConfigApi.md#UpdateApiV1ConfigDevelopment) | **Put** /api/v1/config/development/{id} | 



## GetApiV1ConfigDevelopment

> DevelopmentConfigResource GetApiV1ConfigDevelopment(ctx).Execute()



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
    resp, r, err := apiClient.DevelopmentConfigApi.GetApiV1ConfigDevelopment(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigApi.GetApiV1ConfigDevelopment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigDevelopment`: DevelopmentConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigApi.GetApiV1ConfigDevelopment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigDevelopmentRequest struct via the builder pattern


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


## GetApiV1ConfigDevelopmentById

> DevelopmentConfigResource GetApiV1ConfigDevelopmentById(ctx, id).Execute()



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
    resp, r, err := apiClient.DevelopmentConfigApi.GetApiV1ConfigDevelopmentById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigApi.GetApiV1ConfigDevelopmentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigDevelopmentById`: DevelopmentConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigApi.GetApiV1ConfigDevelopmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigDevelopmentByIdRequest struct via the builder pattern


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


## UpdateApiV1ConfigDevelopment

> DevelopmentConfigResource UpdateApiV1ConfigDevelopment(ctx, id).DevelopmentConfigResource(developmentConfigResource).Execute()



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
    resp, r, err := apiClient.DevelopmentConfigApi.UpdateApiV1ConfigDevelopment(context.Background(), id).DevelopmentConfigResource(developmentConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigApi.UpdateApiV1ConfigDevelopment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ConfigDevelopment`: DevelopmentConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigApi.UpdateApiV1ConfigDevelopment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ConfigDevelopmentRequest struct via the builder pattern


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

