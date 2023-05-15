# \HistoryAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistory**](HistoryAPI.md#GetHistory) | **Get** /api/v1/history | 
[**ListHistoryIndexer**](HistoryAPI.md#ListHistoryIndexer) | **Get** /api/v1/history/indexer | 
[**ListHistorySince**](HistoryAPI.md#ListHistorySince) | **Get** /api/v1/history/since | 



## GetHistory

> HistoryResourcePagingResource GetHistory(ctx).Execute()



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
    resp, r, err := apiClient.HistoryAPI.GetHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.GetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistory`: HistoryResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.GetHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistoryIndexer

> []HistoryResource ListHistoryIndexer(ctx).IndexerId(indexerId).EventType(eventType).Execute()



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
    indexerId := int32(56) // int32 |  (optional)
    eventType := prowlarrClient.HistoryEventType("unknown") // HistoryEventType |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryAPI.ListHistoryIndexer(context.Background()).IndexerId(indexerId).EventType(eventType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.ListHistoryIndexer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryIndexer`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.ListHistoryIndexer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryIndexerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexerId** | **int32** |  | 
 **eventType** | [**HistoryEventType**](HistoryEventType.md) |  | 

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistorySince

> []HistoryResource ListHistorySince(ctx).Date(date).EventType(eventType).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    prowlarrClient "./openapi"
)

func main() {
    date := time.Now() // time.Time |  (optional)
    eventType := prowlarrClient.HistoryEventType("unknown") // HistoryEventType |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryAPI.ListHistorySince(context.Background()).Date(date).EventType(eventType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.ListHistorySince``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistorySince`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.ListHistorySince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistorySinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **eventType** | [**HistoryEventType**](HistoryEventType.md) |  | 

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

