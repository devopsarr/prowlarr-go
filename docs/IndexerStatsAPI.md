# \IndexerStatsAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIndexerStats**](IndexerStatsAPI.md#GetIndexerStats) | **Get** /api/v1/indexerstats | 



## GetIndexerStats

> IndexerStatsResource GetIndexerStats(ctx).StartDate(startDate).EndDate(endDate).Indexers(indexers).Protocols(protocols).Tags(tags).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	prowlarrClient "github.com/devopsarr/prowlarr-go/prowlarr"
)

func main() {
	startDate := time.Now() // time.Time |  (optional)
	endDate := time.Now() // time.Time |  (optional)
	indexers := "indexers_example" // string |  (optional)
	protocols := "protocols_example" // string |  (optional)
	tags := "tags_example" // string |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexerStatsAPI.GetIndexerStats(context.Background()).StartDate(startDate).EndDate(endDate).Indexers(indexers).Protocols(protocols).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerStatsAPI.GetIndexerStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexerStats`: IndexerStatsResource
	fmt.Fprintf(os.Stdout, "Response from `IndexerStatsAPI.GetIndexerStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexerStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 
 **indexers** | **string** |  | 
 **protocols** | **string** |  | 
 **tags** | **string** |  | 

### Return type

[**IndexerStatsResource**](IndexerStatsResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

