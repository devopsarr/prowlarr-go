# \CustomFilterApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1CustomFilter**](CustomFilterApi.md#CreateApiV1CustomFilter) | **Post** /api/v1/customfilter | 
[**DeleteApiV1CustomFilter**](CustomFilterApi.md#DeleteApiV1CustomFilter) | **Delete** /api/v1/customfilter/{id} | 
[**GetApiV1CustomFilterById**](CustomFilterApi.md#GetApiV1CustomFilterById) | **Get** /api/v1/customfilter/{id} | 
[**ListApiV1CustomFilter**](CustomFilterApi.md#ListApiV1CustomFilter) | **Get** /api/v1/customfilter | 
[**UpdateApiV1CustomFilter**](CustomFilterApi.md#UpdateApiV1CustomFilter) | **Put** /api/v1/customfilter/{id} | 



## CreateApiV1CustomFilter

> CustomFilterResource CreateApiV1CustomFilter(ctx).CustomFilterResource(customFilterResource).Execute()



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
    customFilterResource := *prowlarrClient.NewCustomFilterResource() // CustomFilterResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFilterApi.CreateApiV1CustomFilter(context.Background()).CustomFilterResource(customFilterResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.CreateApiV1CustomFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1CustomFilter`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterApi.CreateApiV1CustomFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1CustomFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customFilterResource** | [**CustomFilterResource**](CustomFilterResource.md) |  | 

### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1CustomFilter

> DeleteApiV1CustomFilter(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFilterApi.DeleteApiV1CustomFilter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.DeleteApiV1CustomFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1CustomFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1CustomFilterById

> CustomFilterResource GetApiV1CustomFilterById(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFilterApi.GetApiV1CustomFilterById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.GetApiV1CustomFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1CustomFilterById`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterApi.GetApiV1CustomFilterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1CustomFilterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1CustomFilter

> []CustomFilterResource ListApiV1CustomFilter(ctx).Execute()



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
    resp, r, err := apiClient.CustomFilterApi.ListApiV1CustomFilter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.ListApiV1CustomFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1CustomFilter`: []CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterApi.ListApiV1CustomFilter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1CustomFilterRequest struct via the builder pattern


### Return type

[**[]CustomFilterResource**](CustomFilterResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1CustomFilter

> CustomFilterResource UpdateApiV1CustomFilter(ctx, id).CustomFilterResource(customFilterResource).Execute()



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
    customFilterResource := *prowlarrClient.NewCustomFilterResource() // CustomFilterResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFilterApi.UpdateApiV1CustomFilter(context.Background(), id).CustomFilterResource(customFilterResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.UpdateApiV1CustomFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1CustomFilter`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterApi.UpdateApiV1CustomFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1CustomFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFilterResource** | [**CustomFilterResource**](CustomFilterResource.md) |  | 

### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

