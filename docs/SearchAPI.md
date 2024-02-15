# \SearchAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearch**](SearchAPI.md#CreateSearch) | **Post** /api/v1/search | 
[**CreateSearchBulk**](SearchAPI.md#CreateSearchBulk) | **Post** /api/v1/search/bulk | 
[**ListSearch**](SearchAPI.md#ListSearch) | **Get** /api/v1/search | 



## CreateSearch

> ReleaseResource CreateSearch(ctx).ReleaseResource(releaseResource).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	prowlarrClient "github.com/devopsarr/prowlarr-go/prowlarr"
)

func main() {
	releaseResource := *prowlarrClient.NewReleaseResource() // ReleaseResource |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CreateSearch(context.Background()).ReleaseResource(releaseResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CreateSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSearch`: ReleaseResource
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CreateSearch`: %v\n", resp)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
	prowlarrClient "github.com/devopsarr/prowlarr-go/prowlarr"
)

func main() {
	releaseResource := []prowlarrClient.ReleaseResource{*prowlarrClient.NewReleaseResource()} // []ReleaseResource |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CreateSearchBulk(context.Background()).ReleaseResource(releaseResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CreateSearchBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSearchBulk`: ReleaseResource
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CreateSearchBulk`: %v\n", resp)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
	prowlarrClient "github.com/devopsarr/prowlarr-go/prowlarr"
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
	resp, r, err := apiClient.SearchAPI.ListSearch(context.Background()).Query(query).Type_(type_).IndexerIds(indexerIds).Categories(categories).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.ListSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSearch`: []ReleaseResource
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.ListSearch`: %v\n", resp)
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

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

