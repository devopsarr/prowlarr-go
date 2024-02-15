# \DevelopmentConfigAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDevelopmentConfig**](DevelopmentConfigAPI.md#GetDevelopmentConfig) | **Get** /api/v1/config/development | 
[**GetDevelopmentConfigById**](DevelopmentConfigAPI.md#GetDevelopmentConfigById) | **Get** /api/v1/config/development/{id} | 
[**UpdateDevelopmentConfig**](DevelopmentConfigAPI.md#UpdateDevelopmentConfig) | **Put** /api/v1/config/development/{id} | 



## GetDevelopmentConfig

> DevelopmentConfigResource GetDevelopmentConfig(ctx).Execute()



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
	resp, r, err := apiClient.DevelopmentConfigAPI.GetDevelopmentConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigAPI.GetDevelopmentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevelopmentConfig`: DevelopmentConfigResource
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigAPI.GetDevelopmentConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevelopmentConfigRequest struct via the builder pattern


### Return type

[**DevelopmentConfigResource**](DevelopmentConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevelopmentConfigById

> DevelopmentConfigResource GetDevelopmentConfigById(ctx, id).Execute()



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
	id := int32(56) // int32 | 

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentConfigAPI.GetDevelopmentConfigById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigAPI.GetDevelopmentConfigById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevelopmentConfigById`: DevelopmentConfigResource
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigAPI.GetDevelopmentConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevelopmentConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DevelopmentConfigResource**](DevelopmentConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevelopmentConfig

> DevelopmentConfigResource UpdateDevelopmentConfig(ctx, id).DevelopmentConfigResource(developmentConfigResource).Execute()



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
	id := "id_example" // string | 
	developmentConfigResource := *prowlarrClient.NewDevelopmentConfigResource() // DevelopmentConfigResource |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentConfigAPI.UpdateDevelopmentConfig(context.Background(), id).DevelopmentConfigResource(developmentConfigResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentConfigAPI.UpdateDevelopmentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDevelopmentConfig`: DevelopmentConfigResource
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentConfigAPI.UpdateDevelopmentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDevelopmentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **developmentConfigResource** | [**DevelopmentConfigResource**](DevelopmentConfigResource.md) |  | 

### Return type

[**DevelopmentConfigResource**](DevelopmentConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

