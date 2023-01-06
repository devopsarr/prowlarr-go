# \IndexerEditorApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1IndexerEditor**](IndexerEditorApi.md#DeleteApiV1IndexerEditor) | **Delete** /api/v1/indexer/editor | 
[**PutApiV1IndexerEditor**](IndexerEditorApi.md#PutApiV1IndexerEditor) | **Put** /api/v1/indexer/editor | 



## DeleteApiV1IndexerEditor

> DeleteApiV1IndexerEditor(ctx).IndexerEditorResource(indexerEditorResource).Execute()



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
    indexerEditorResource := *prowlarrClient.NewIndexerEditorResource() // IndexerEditorResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerEditorApi.DeleteApiV1IndexerEditor(context.Background()).IndexerEditorResource(indexerEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerEditorApi.DeleteApiV1IndexerEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1IndexerEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexerEditorResource** | [**IndexerEditorResource**](IndexerEditorResource.md) |  | 

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


## PutApiV1IndexerEditor

> PutApiV1IndexerEditor(ctx).IndexerEditorResource(indexerEditorResource).Execute()



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
    indexerEditorResource := *prowlarrClient.NewIndexerEditorResource() // IndexerEditorResource |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerEditorApi.PutApiV1IndexerEditor(context.Background()).IndexerEditorResource(indexerEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerEditorApi.PutApiV1IndexerEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1IndexerEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexerEditorResource** | [**IndexerEditorResource**](IndexerEditorResource.md) |  | 

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

