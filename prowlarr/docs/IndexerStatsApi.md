# \IndexerStatsApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIndexerStats**](IndexerStatsApi.md#GetIndexerStats) | **Get** /api/v1/indexerstats | 



## GetIndexerStats

> IndexerStatsResource GetIndexerStats(ctx).StartDate(startDate).EndDate(endDate).Execute()



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
    startDate := time.Now() // time.Time |  (optional)
    endDate := time.Now() // time.Time |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerStatsApi.GetIndexerStats(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerStatsApi.GetIndexerStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndexerStats`: IndexerStatsResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerStatsApi.GetIndexerStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexerStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

### Return type

[**IndexerStatsResource**](IndexerStatsResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

