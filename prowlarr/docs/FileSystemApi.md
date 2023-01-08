# \FileSystemApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1FileSystem**](FileSystemApi.md#GetApiV1FileSystem) | **Get** /api/v1/filesystem | 
[**GetApiV1FileSystemType**](FileSystemApi.md#GetApiV1FileSystemType) | **Get** /api/v1/filesystem/type | 



## GetApiV1FileSystem

> GetApiV1FileSystem(ctx).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()



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
    path := "path_example" // string |  (optional)
    includeFiles := true // bool |  (optional) (default to false)
    allowFoldersWithoutTrailingSlashes := true // bool |  (optional) (default to false)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemApi.GetApiV1FileSystem(context.Background()).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.GetApiV1FileSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1FileSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **includeFiles** | **bool** |  | [default to false]
 **allowFoldersWithoutTrailingSlashes** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1FileSystemType

> GetApiV1FileSystemType(ctx).Path(path).Execute()



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
    path := "path_example" // string |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemApi.GetApiV1FileSystemType(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.GetApiV1FileSystemType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1FileSystemTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

