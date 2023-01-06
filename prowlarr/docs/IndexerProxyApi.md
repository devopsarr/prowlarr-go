# \IndexerProxyApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiV1Indexerproxy**](IndexerProxyApi.md#CreateApiV1Indexerproxy) | **Post** /api/v1/indexerproxy | 
[**CreateApiV1IndexerproxyActionByName**](IndexerProxyApi.md#CreateApiV1IndexerproxyActionByName) | **Post** /api/v1/indexerproxy/action/{name} | 
[**DeleteApiV1Indexerproxy**](IndexerProxyApi.md#DeleteApiV1Indexerproxy) | **Delete** /api/v1/indexerproxy/{id} | 
[**GetApiV1IndexerproxyById**](IndexerProxyApi.md#GetApiV1IndexerproxyById) | **Get** /api/v1/indexerproxy/{id} | 
[**ListApiV1Indexerproxy**](IndexerProxyApi.md#ListApiV1Indexerproxy) | **Get** /api/v1/indexerproxy | 
[**ListApiV1IndexerproxySchema**](IndexerProxyApi.md#ListApiV1IndexerproxySchema) | **Get** /api/v1/indexerproxy/schema | 
[**TestApiV1Indexerproxy**](IndexerProxyApi.md#TestApiV1Indexerproxy) | **Post** /api/v1/indexerproxy/test | 
[**TestallApiV1Indexerproxy**](IndexerProxyApi.md#TestallApiV1Indexerproxy) | **Post** /api/v1/indexerproxy/testall | 
[**UpdateApiV1Indexerproxy**](IndexerProxyApi.md#UpdateApiV1Indexerproxy) | **Put** /api/v1/indexerproxy/{id} | 



## CreateApiV1Indexerproxy

> IndexerProxyResource CreateApiV1Indexerproxy(ctx).IndexerProxyResource(indexerProxyResource).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.CreateApiV1Indexerproxy(context.Background()).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.CreateApiV1Indexerproxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiV1Indexerproxy`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.CreateApiV1Indexerproxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiV1IndexerproxyRequest struct via the builder pattern


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


## CreateApiV1IndexerproxyActionByName

> CreateApiV1IndexerproxyActionByName(ctx, name).IndexerProxyResource(indexerProxyResource).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.CreateApiV1IndexerproxyActionByName(context.Background(), name).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.CreateApiV1IndexerproxyActionByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateApiV1IndexerproxyActionByNameRequest struct via the builder pattern


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


## DeleteApiV1Indexerproxy

> DeleteApiV1Indexerproxy(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.DeleteApiV1Indexerproxy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.DeleteApiV1Indexerproxy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiV1IndexerproxyRequest struct via the builder pattern


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


## GetApiV1IndexerproxyById

> IndexerProxyResource GetApiV1IndexerproxyById(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.GetApiV1IndexerproxyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.GetApiV1IndexerproxyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiV1IndexerproxyById`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.GetApiV1IndexerproxyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1IndexerproxyByIdRequest struct via the builder pattern


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


## ListApiV1Indexerproxy

> []IndexerProxyResource ListApiV1Indexerproxy(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.ListApiV1Indexerproxy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.ListApiV1Indexerproxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1Indexerproxy`: []IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.ListApiV1Indexerproxy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1IndexerproxyRequest struct via the builder pattern


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


## ListApiV1IndexerproxySchema

> []IndexerProxyResource ListApiV1IndexerproxySchema(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.ListApiV1IndexerproxySchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.ListApiV1IndexerproxySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1IndexerproxySchema`: []IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.ListApiV1IndexerproxySchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1IndexerproxySchemaRequest struct via the builder pattern


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


## TestApiV1Indexerproxy

> TestApiV1Indexerproxy(ctx).IndexerProxyResource(indexerProxyResource).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.TestApiV1Indexerproxy(context.Background()).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.TestApiV1Indexerproxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestApiV1IndexerproxyRequest struct via the builder pattern


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


## TestallApiV1Indexerproxy

> TestallApiV1Indexerproxy(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.TestallApiV1Indexerproxy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.TestallApiV1Indexerproxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallApiV1IndexerproxyRequest struct via the builder pattern


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


## UpdateApiV1Indexerproxy

> IndexerProxyResource UpdateApiV1Indexerproxy(ctx, id).IndexerProxyResource(indexerProxyResource).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.UpdateApiV1Indexerproxy(context.Background(), id).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.UpdateApiV1Indexerproxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiV1Indexerproxy`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.UpdateApiV1Indexerproxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiV1IndexerproxyRequest struct via the builder pattern


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

