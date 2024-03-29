# \LogFileAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogFileByFilename**](LogFileAPI.md#GetLogFileByFilename) | **Get** /api/v1/log/file/{filename} | 
[**ListLogFile**](LogFileAPI.md#ListLogFile) | **Get** /api/v1/log/file | 



## GetLogFileByFilename

> map[string]interface{} GetLogFileByFilename(ctx, filename).Execute()



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
	filename := "filename_example" // string | 

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogFileAPI.GetLogFileByFilename(context.Background(), filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogFileAPI.GetLogFileByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogFileByFilename`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LogFileAPI.GetLogFileByFilename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFileByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogFile

> []LogFileResource ListLogFile(ctx).Execute()



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

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogFileAPI.ListLogFile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogFileAPI.ListLogFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogFile`: []LogFileResource
	fmt.Fprintf(os.Stdout, "Response from `LogFileAPI.ListLogFile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogFileRequest struct via the builder pattern


### Return type

[**[]LogFileResource**](LogFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

