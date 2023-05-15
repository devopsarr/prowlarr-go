# \IndexerStatusAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIndexerStatusById**](IndexerStatusAPI.md#GetIndexerStatusById) | **Get** /api/v1/indexerstatus/{id} | 
[**ListIndexerStatus**](IndexerStatusAPI.md#ListIndexerStatus) | **Get** /api/v1/indexerstatus | 



## GetIndexerStatusById

> IndexerStatusResource GetIndexerStatusById(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerStatusAPI.GetIndexerStatusById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerStatusAPI.GetIndexerStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndexerStatusById`: IndexerStatusResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerStatusAPI.GetIndexerStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexerStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexerStatusResource**](IndexerStatusResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexerStatus

> []IndexerStatusResource ListIndexerStatus(ctx).Execute()



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
    resp, r, err := apiClient.IndexerStatusAPI.ListIndexerStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerStatusAPI.ListIndexerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIndexerStatus`: []IndexerStatusResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerStatusAPI.ListIndexerStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerStatusRequest struct via the builder pattern


### Return type

[**[]IndexerStatusResource**](IndexerStatusResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

