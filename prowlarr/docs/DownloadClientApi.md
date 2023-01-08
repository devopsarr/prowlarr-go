# \DownloadClientApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1DownloadClient**](DownloadClientApi.md#CreateApiV1DownloadClient) | **Post** /api/v1/downloadclient | 
[**CreateApiV1DownloadClientActionByName**](DownloadClientApi.md#CreateApiV1DownloadClientActionByName) | **Post** /api/v1/downloadclient/action/{name} | 
[**DeleteApiV1DownloadClient**](DownloadClientApi.md#DeleteApiV1DownloadClient) | **Delete** /api/v1/downloadclient/{id} | 
[**GetApiV1DownloadClientById**](DownloadClientApi.md#GetApiV1DownloadClientById) | **Get** /api/v1/downloadclient/{id} | 
[**ListApiV1DownloadClient**](DownloadClientApi.md#ListApiV1DownloadClient) | **Get** /api/v1/downloadclient | 
[**ListApiV1DownloadClientSchema**](DownloadClientApi.md#ListApiV1DownloadClientSchema) | **Get** /api/v1/downloadclient/schema | 
[**TestApiV1DownloadClient**](DownloadClientApi.md#TestApiV1DownloadClient) | **Post** /api/v1/downloadclient/test | 
[**TestallApiV1DownloadClient**](DownloadClientApi.md#TestallApiV1DownloadClient) | **Post** /api/v1/downloadclient/testall | 
[**UpdateApiV1DownloadClient**](DownloadClientApi.md#UpdateApiV1DownloadClient) | **Put** /api/v1/downloadclient/{id} | 



## CreateApiV1DownloadClient

> DownloadClientResource CreateApiV1DownloadClient(ctx).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *prowlarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.CreateApiV1DownloadClient(context.Background()).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.CreateApiV1DownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1DownloadClient`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.CreateApiV1DownloadClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1DownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiV1DownloadClientActionByName

> CreateApiV1DownloadClientActionByName(ctx, name).DownloadClientResource(downloadClientResource).Execute()



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
    name := "name_example" // string | 
    downloadClientResource := *prowlarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.CreateApiV1DownloadClientActionByName(context.Background(), name).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.CreateApiV1DownloadClientActionByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1DownloadClientActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiV1DownloadClient

> DeleteApiV1DownloadClient(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.DeleteApiV1DownloadClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.DeleteApiV1DownloadClient``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1DownloadClientRequest struct via the builder pattern


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


## GetApiV1DownloadClientById

> DownloadClientResource GetApiV1DownloadClientById(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.GetApiV1DownloadClientById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.GetApiV1DownloadClientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1DownloadClientById`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.GetApiV1DownloadClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1DownloadClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1DownloadClient

> []DownloadClientResource ListApiV1DownloadClient(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ListApiV1DownloadClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ListApiV1DownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1DownloadClient`: []DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ListApiV1DownloadClient`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1DownloadClientRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1DownloadClientSchema

> []DownloadClientResource ListApiV1DownloadClientSchema(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ListApiV1DownloadClientSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ListApiV1DownloadClientSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1DownloadClientSchema`: []DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ListApiV1DownloadClientSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1DownloadClientSchemaRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestApiV1DownloadClient

> TestApiV1DownloadClient(ctx).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *prowlarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.TestApiV1DownloadClient(context.Background()).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.TestApiV1DownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestApiV1DownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestallApiV1DownloadClient

> TestallApiV1DownloadClient(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.TestallApiV1DownloadClient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.TestallApiV1DownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallApiV1DownloadClientRequest struct via the builder pattern


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


## UpdateApiV1DownloadClient

> DownloadClientResource UpdateApiV1DownloadClient(ctx, id).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *prowlarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.UpdateApiV1DownloadClient(context.Background(), id).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.UpdateApiV1DownloadClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1DownloadClient`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.UpdateApiV1DownloadClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1DownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

