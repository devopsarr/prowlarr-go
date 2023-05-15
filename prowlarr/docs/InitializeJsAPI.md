# \InitializeJsAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInitializeJs**](InitializeJsAPI.md#GetInitializeJs) | **Get** /initialize.js | 



## GetInitializeJs

> GetInitializeJs(ctx).Execute()



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
    resp, r, err := apiClient.InitializeJsAPI.GetInitializeJs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InitializeJsAPI.GetInitializeJs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInitializeJsRequest struct via the builder pattern


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

