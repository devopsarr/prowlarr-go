# \IndexerProxyApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndexerProxy**](IndexerProxyApi.md#CreateIndexerProxy) | **Post** /api/v1/indexerproxy | 
[**CreateIndexerProxyActionByName**](IndexerProxyApi.md#CreateIndexerProxyActionByName) | **Post** /api/v1/indexerproxy/action/{name} | 
[**DeleteIndexerProxy**](IndexerProxyApi.md#DeleteIndexerProxy) | **Delete** /api/v1/indexerproxy/{id} | 
[**GetIndexerProxyById**](IndexerProxyApi.md#GetIndexerProxyById) | **Get** /api/v1/indexerproxy/{id} | 
[**ListIndexerProxy**](IndexerProxyApi.md#ListIndexerProxy) | **Get** /api/v1/indexerproxy | 
[**ListIndexerProxySchema**](IndexerProxyApi.md#ListIndexerProxySchema) | **Get** /api/v1/indexerproxy/schema | 
[**TestIndexerProxy**](IndexerProxyApi.md#TestIndexerProxy) | **Post** /api/v1/indexerproxy/test | 
[**TestallIndexerProxy**](IndexerProxyApi.md#TestallIndexerProxy) | **Post** /api/v1/indexerproxy/testall | 
[**UpdateIndexerProxy**](IndexerProxyApi.md#UpdateIndexerProxy) | **Put** /api/v1/indexerproxy/{id} | 



## CreateIndexerProxy

> IndexerProxyResource CreateIndexerProxy(ctx).ForceSave(forceSave).IndexerProxyResource(indexerProxyResource).Execute()



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
    forceSave := true // bool |  (optional) (default to false)
    indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerProxyApi.CreateIndexerProxy(context.Background()).ForceSave(forceSave).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.CreateIndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndexerProxy`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.CreateIndexerProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceSave** | **bool** |  | [default to false]
 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIndexerProxyActionByName

> CreateIndexerProxyActionByName(ctx, name).IndexerProxyResource(indexerProxyResource).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.CreateIndexerProxyActionByName(context.Background(), name).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.CreateIndexerProxyActionByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateIndexerProxyActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndexerProxy

> DeleteIndexerProxy(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.DeleteIndexerProxy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.DeleteIndexerProxy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIndexerProxyById

> IndexerProxyResource GetIndexerProxyById(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.GetIndexerProxyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.GetIndexerProxyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndexerProxyById`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.GetIndexerProxyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexerProxyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexerProxy

> []IndexerProxyResource ListIndexerProxy(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.ListIndexerProxy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.ListIndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIndexerProxy`: []IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.ListIndexerProxy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerProxyRequest struct via the builder pattern


### Return type

[**[]IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexerProxySchema

> []IndexerProxyResource ListIndexerProxySchema(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.ListIndexerProxySchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.ListIndexerProxySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIndexerProxySchema`: []IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.ListIndexerProxySchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerProxySchemaRequest struct via the builder pattern


### Return type

[**[]IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIndexerProxy

> TestIndexerProxy(ctx).IndexerProxyResource(indexerProxyResource).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.TestIndexerProxy(context.Background()).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.TestIndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestIndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestallIndexerProxy

> TestallIndexerProxy(ctx).Execute()



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
    resp, r, err := apiClient.IndexerProxyApi.TestallIndexerProxy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.TestallIndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallIndexerProxyRequest struct via the builder pattern


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


## UpdateIndexerProxy

> IndexerProxyResource UpdateIndexerProxy(ctx, id).ForceSave(forceSave).IndexerProxyResource(indexerProxyResource).Execute()



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
    forceSave := true // bool |  (optional) (default to false)
    indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerProxyApi.UpdateIndexerProxy(context.Background(), id).ForceSave(forceSave).IndexerProxyResource(indexerProxyResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyApi.UpdateIndexerProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndexerProxy`: IndexerProxyResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerProxyApi.UpdateIndexerProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceSave** | **bool** |  | [default to false]
 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

