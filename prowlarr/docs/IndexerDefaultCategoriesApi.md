# \IndexerDefaultCategoriesApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIndexerCategories**](IndexerDefaultCategoriesApi.md#ListIndexerCategories) | **Get** /api/v1/indexer/categories | 



## ListIndexerCategories

> []IndexerCategory ListIndexerCategories(ctx).Execute()



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
    resp, r, err := apiClient.IndexerDefaultCategoriesApi.ListIndexerCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerDefaultCategoriesApi.ListIndexerCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIndexerCategories`: []IndexerCategory
    fmt.Fprintf(os.Stdout, "Response from `IndexerDefaultCategoriesApi.ListIndexerCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerCategoriesRequest struct via the builder pattern


### Return type

[**[]IndexerCategory**](IndexerCategory.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

