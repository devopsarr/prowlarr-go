# \FileSystemAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFileSystem**](FileSystemAPI.md#GetFileSystem) | **Get** /api/v1/filesystem | 
[**GetFileSystemType**](FileSystemAPI.md#GetFileSystemType) | **Get** /api/v1/filesystem/type | 



## GetFileSystem

> GetFileSystem(ctx).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()



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
	path := "path_example" // string |  (optional)
	includeFiles := true // bool |  (optional) (default to false)
	allowFoldersWithoutTrailingSlashes := true // bool |  (optional) (default to false)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	r, err := apiClient.FileSystemAPI.GetFileSystem(context.Background()).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileSystemAPI.GetFileSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **includeFiles** | **bool** |  | [default to false]
 **allowFoldersWithoutTrailingSlashes** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileSystemType

> GetFileSystemType(ctx).Path(path).Execute()



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
	path := "path_example" // string |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	r, err := apiClient.FileSystemAPI.GetFileSystemType(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileSystemAPI.GetFileSystemType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileSystemTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

