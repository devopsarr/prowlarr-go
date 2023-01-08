# \IndexerProxyApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1IndexerProxy**](IndexerProxyApi.md#CreateApiV1IndexerProxy) | **Post** /api/v1/indexerproxy | 
[**CreateApiV1IndexerProxyActionByName**](IndexerProxyApi.md#CreateApiV1IndexerProxyActionByName) | **Post** /api/v1/indexerproxy/action/{name} | 
[**DeleteApiV1IndexerProxy**](IndexerProxyApi.md#DeleteApiV1IndexerProxy) | **Delete** /api/v1/indexerproxy/{id} | 
[**GetApiV1IndexerProxyById**](IndexerProxyApi.md#GetApiV1IndexerProxyById) | **Get** /api/v1/indexerproxy/{id} | 
[**ListApiV1IndexerProxy**](IndexerProxyApi.md#ListApiV1IndexerProxy) | **Get** /api/v1/indexerproxy | 
[**ListApiV1IndexerProxySchema**](IndexerProxyApi.md#ListApiV1IndexerProxySchema) | **Get** /api/v1/indexerproxy/schema | 
[**TestApiV1IndexerProxy**](IndexerProxyApi.md#TestApiV1IndexerProxy) | **Post** /api/v1/indexerproxy/test | 
[**TestallApiV1IndexerProxy**](IndexerProxyApi.md#TestallApiV1IndexerProxy) | **Post** /api/v1/indexerproxy/testall | 
[**UpdateApiV1IndexerProxy**](IndexerProxyApi.md#UpdateApiV1IndexerProxy) | **Put** /api/v1/indexerproxy/{id} | 



## CreateApiV1IndexerProxy

> IndexerProxyResource CreateApiV1IndexerProxy(ctx).IndexerProxyResource(indexerProxyResource).Execute()



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
    indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerProxyApi.CreateApiV1IndexerProxy(context.Background()).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.CreateApiV1IndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1IndexerProxy`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.CreateApiV1IndexerProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1IndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiV1IndexerProxyActionByName

> CreateApiV1IndexerProxyActionByName(ctx, name).IndexerProxyResource(indexerProxyResource).Execute()



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
    indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerProxyApi.CreateApiV1IndexerProxyActionByName(context.Background(), name).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.CreateApiV1IndexerProxyActionByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateApiV1IndexerProxyActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

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


## DeleteApiV1IndexerProxy

> DeleteApiV1IndexerProxy(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.DeleteApiV1IndexerProxy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.DeleteApiV1IndexerProxy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1IndexerProxyRequest struct via the builder pattern


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


## GetApiV1IndexerProxyById

> IndexerProxyResource GetApiV1IndexerProxyById(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.GetApiV1IndexerProxyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.GetApiV1IndexerProxyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1IndexerProxyById`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.GetApiV1IndexerProxyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1IndexerProxyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1IndexerProxy

> []IndexerProxyResource ListApiV1IndexerProxy(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.ListApiV1IndexerProxy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.ListApiV1IndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1IndexerProxy`: []IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.ListApiV1IndexerProxy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1IndexerProxyRequest struct via the builder pattern


### Return type

[**[]IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1IndexerProxySchema

> []IndexerProxyResource ListApiV1IndexerProxySchema(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.ListApiV1IndexerProxySchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.ListApiV1IndexerProxySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1IndexerProxySchema`: []IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.ListApiV1IndexerProxySchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1IndexerProxySchemaRequest struct via the builder pattern


### Return type

[**[]IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestApiV1IndexerProxy

> TestApiV1IndexerProxy(ctx).IndexerProxyResource(indexerProxyResource).Execute()



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
    indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerProxyApi.TestApiV1IndexerProxy(context.Background()).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.TestApiV1IndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestApiV1IndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

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


## TestallApiV1IndexerProxy

> TestallApiV1IndexerProxy(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.TestallApiV1IndexerProxy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.TestallApiV1IndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallApiV1IndexerProxyRequest struct via the builder pattern


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


## UpdateApiV1IndexerProxy

> IndexerProxyResource UpdateApiV1IndexerProxy(ctx, id).IndexerProxyResource(indexerProxyResource).Execute()



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
    indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerProxyApi.UpdateApiV1IndexerProxy(context.Background(), id).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.UpdateApiV1IndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1IndexerProxy`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.UpdateApiV1IndexerProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1IndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

