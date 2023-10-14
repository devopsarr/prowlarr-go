# \QualityProfileSchemaAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppprofileSchema**](QualityProfileSchemaAPI.md#GetAppprofileSchema) | **Get** /api/v1/appprofile/schema | 



## GetAppprofileSchema

> AppProfileResource GetAppprofileSchema(ctx).Execute()



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
    resp, r, err := apiClient.QualityProfileSchemaAPI.GetAppprofileSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileSchemaAPI.GetAppprofileSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppprofileSchema`: AppProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileSchemaAPI.GetAppprofileSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppprofileSchemaRequest struct via the builder pattern


### Return type

[**AppProfileResource**](AppProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

