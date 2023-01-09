# \SearchApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearch**](SearchApi.md#CreateSearch) | **Post** /api/v1/search | 
[**CreateSearchBulk**](SearchApi.md#CreateSearchBulk) | **Post** /api/v1/search/bulk | 
[**GetSearchById**](SearchApi.md#GetSearchById) | **Get** /api/v1/search/{id} | 
[**ListSearch**](SearchApi.md#ListSearch) | **Get** /api/v1/search | 



## CreateSearch

> ReleaseResource CreateSearch(ctx).ReleaseResource(releaseResource).Execute()



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
    releaseResource := *prowlarrClient.NewReleaseResource() // ReleaseResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.CreateSearch(context.Background()).ReleaseResource(releaseResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CreateSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSearch`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CreateSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseResource** | [**ReleaseResource**](ReleaseResource.md) |  | 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSearchBulk

> ReleaseResource CreateSearchBulk(ctx).ReleaseResource(releaseResource).Execute()



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
    releaseResource := []prowlarrClient.ReleaseResource{*prowlarrClient.NewReleaseResource()} // []ReleaseResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.CreateSearchBulk(context.Background()).ReleaseResource(releaseResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CreateSearchBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSearchBulk`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CreateSearchBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSearchBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseResource** | [**[]ReleaseResource**](ReleaseResource.md) |  | 

### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchById

> ReleaseResource GetSearchById(ctx, id).Execute()



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
    resp, r, err := apiClient.SearchApi.GetSearchById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSearchById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchById`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSearchById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSearch

> []ReleaseResource ListSearch(ctx).Query(query).Type_(type_).IndexerIds(indexerIds).Categories(categories).Limit(limit).Offset(offset).Execute()



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
    query := "query_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    indexerIds := []int32{int32(123)} // []int32 |  (optional)
    categories := []int32{int32(123)} // []int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.ListSearch(context.Background()).Query(query).Type_(type_).IndexerIds(indexerIds).Categories(categories).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.ListSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSearch`: []ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.ListSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **type_** | **string** |  | 
 **indexerIds** | **[]int32** |  | 
 **categories** | **[]int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**[]ReleaseResource**](ReleaseResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

