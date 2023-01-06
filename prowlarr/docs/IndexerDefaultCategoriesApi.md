# \IndexerDefaultCategoriesApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApiV1IndexerCategories**](IndexerDefaultCategoriesApi.md#ListApiV1IndexerCategories) | **Get** /api/v1/indexer/categories | 



## ListApiV1IndexerCategories

> []IndexerCategory ListApiV1IndexerCategories(ctx).Execute()



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
    resp, r, err := apiClient.IndexerDefaultCategoriesApi.ListApiV1IndexerCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerDefaultCategoriesApi.ListApiV1IndexerCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiV1IndexerCategories`: []IndexerCategory
    fmt.Fprintf(os.Stdout, "Response from `IndexerDefaultCategoriesApi.ListApiV1IndexerCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1IndexerCategoriesRequest struct via the builder pattern


### Return type

[**[]IndexerCategory**](IndexerCategory.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

