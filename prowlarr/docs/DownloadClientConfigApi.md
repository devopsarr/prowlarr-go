# \DownloadClientConfigApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1ConfigDownloadclient**](DownloadClientConfigApi.md#GetApiV1ConfigDownloadclient) | **Get** /api/v1/config/downloadclient | 
[**GetApiV1ConfigDownloadclientById**](DownloadClientConfigApi.md#GetApiV1ConfigDownloadclientById) | **Get** /api/v1/config/downloadclient/{id} | 
[**UpdateApiV1ConfigDownloadclient**](DownloadClientConfigApi.md#UpdateApiV1ConfigDownloadclient) | **Put** /api/v1/config/downloadclient/{id} | 



## GetApiV1ConfigDownloadclient

> DownloadClientConfigResource GetApiV1ConfigDownloadclient(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientConfigApi.GetApiV1ConfigDownloadclient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.GetApiV1ConfigDownloadclient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigDownloadclient`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.GetApiV1ConfigDownloadclient`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigDownloadclientRequest struct via the builder pattern


### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1ConfigDownloadclientById

> DownloadClientConfigResource GetApiV1ConfigDownloadclientById(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientConfigApi.GetApiV1ConfigDownloadclientById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.GetApiV1ConfigDownloadclientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1ConfigDownloadclientById`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.GetApiV1ConfigDownloadclientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1ConfigDownloadclientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiV1ConfigDownloadclient

> DownloadClientConfigResource UpdateApiV1ConfigDownloadclient(ctx, id).DownloadClientConfigResource(downloadClientConfigResource).Execute()



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
    downloadClientConfigResource := *prowlarrClient.NewDownloadClientConfigResource() // DownloadClientConfigResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientConfigApi.UpdateApiV1ConfigDownloadclient(context.Background(), id).DownloadClientConfigResource(downloadClientConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.UpdateApiV1ConfigDownloadclient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1ConfigDownloadclient`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.UpdateApiV1ConfigDownloadclient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1ConfigDownloadclientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientConfigResource** | [**DownloadClientConfigResource**](DownloadClientConfigResource.md) |  | 

### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

