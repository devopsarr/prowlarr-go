# \IndexerEditorApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndexerEditor**](IndexerEditorApi.md#DeleteIndexerEditor) | **Delete** /api/v1/indexer/editor | 
[**PutIndexerEditor**](IndexerEditorApi.md#PutIndexerEditor) | **Put** /api/v1/indexer/editor | 



## DeleteIndexerEditor

> DeleteIndexerEditor(ctx).IndexerEditorResource(indexerEditorResource).Execute()



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
    resp, r, err := apiClient.IndexerEditorApi.DeleteIndexerEditor(context.Background()).IndexerEditorResource(indexerEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerEditorApi.DeleteIndexerEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndexerEditorRequest struct via the builder pattern


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


## PutIndexerEditor

> PutIndexerEditor(ctx).IndexerEditorResource(indexerEditorResource).Execute()



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
    resp, r, err := apiClient.IndexerEditorApi.PutIndexerEditor(context.Background()).IndexerEditorResource(indexerEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerEditorApi.PutIndexerEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIndexerEditorRequest struct via the builder pattern


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

