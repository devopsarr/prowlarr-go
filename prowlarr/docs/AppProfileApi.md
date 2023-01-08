# \AppProfileApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1AppProfile**](AppProfileApi.md#CreateApiV1AppProfile) | **Post** /api/v1/appprofile | 
[**DeleteApiV1AppProfile**](AppProfileApi.md#DeleteApiV1AppProfile) | **Delete** /api/v1/appprofile/{id} | 
[**GetApiV1AppProfileById**](AppProfileApi.md#GetApiV1AppProfileById) | **Get** /api/v1/appprofile/{id} | 
[**ListApiV1AppProfile**](AppProfileApi.md#ListApiV1AppProfile) | **Get** /api/v1/appprofile | 
[**UpdateApiV1AppProfile**](AppProfileApi.md#UpdateApiV1AppProfile) | **Put** /api/v1/appprofile/{id} | 



## CreateApiV1AppProfile

> AppProfileResource CreateApiV1AppProfile(ctx).AppProfileResource(appProfileResource).Execute()



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
    appProfileResource := *prowlarrClient.NewAppProfileResource() // AppProfileResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppProfileApi.CreateApiV1AppProfile(context.Background()).AppProfileResource(appProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppProfileApi.CreateApiV1AppProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1AppProfile`: AppProfileResource
    fmt.Fprintf(os.Stdout, "Response from `AppProfileApi.CreateApiV1AppProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1AppProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appProfileResource** | [**AppProfileResource**](AppProfileResource.md) |  | 

### Return type

[**AppProfileResource**](AppProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1AppProfile

> DeleteApiV1AppProfile(ctx, id).Execute()



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
    resp, r, err := apiClient.AppProfileApi.DeleteApiV1AppProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppProfileApi.DeleteApiV1AppProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1AppProfileRequest struct via the builder pattern


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


## GetApiV1AppProfileById

> AppProfileResource GetApiV1AppProfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.AppProfileApi.GetApiV1AppProfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppProfileApi.GetApiV1AppProfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1AppProfileById`: AppProfileResource
    fmt.Fprintf(os.Stdout, "Response from `AppProfileApi.GetApiV1AppProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AppProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppProfileResource**](AppProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1AppProfile

> []AppProfileResource ListApiV1AppProfile(ctx).Execute()



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
    resp, r, err := apiClient.AppProfileApi.ListApiV1AppProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppProfileApi.ListApiV1AppProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1AppProfile`: []AppProfileResource
    fmt.Fprintf(os.Stdout, "Response from `AppProfileApi.ListApiV1AppProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1AppProfileRequest struct via the builder pattern


### Return type

[**[]AppProfileResource**](AppProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1AppProfile

> AppProfileResource UpdateApiV1AppProfile(ctx, id).AppProfileResource(appProfileResource).Execute()



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
    appProfileResource := *prowlarrClient.NewAppProfileResource() // AppProfileResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppProfileApi.UpdateApiV1AppProfile(context.Background(), id).AppProfileResource(appProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppProfileApi.UpdateApiV1AppProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1AppProfile`: AppProfileResource
    fmt.Fprintf(os.Stdout, "Response from `AppProfileApi.UpdateApiV1AppProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1AppProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appProfileResource** | [**AppProfileResource**](AppProfileResource.md) |  | 

### Return type

[**AppProfileResource**](AppProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

