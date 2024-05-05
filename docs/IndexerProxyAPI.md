# \IndexerProxyAPI

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndexerProxy**](IndexerProxyAPI.md#CreateIndexerProxy) | **Post** /api/v1/indexerproxy | 
[**CreateIndexerProxyActionByName**](IndexerProxyAPI.md#CreateIndexerProxyActionByName) | **Post** /api/v1/indexerproxy/action/{name} | 
[**DeleteIndexerProxy**](IndexerProxyAPI.md#DeleteIndexerProxy) | **Delete** /api/v1/indexerproxy/{id} | 
[**GetIndexerProxyById**](IndexerProxyAPI.md#GetIndexerProxyById) | **Get** /api/v1/indexerproxy/{id} | 
[**ListIndexerProxy**](IndexerProxyAPI.md#ListIndexerProxy) | **Get** /api/v1/indexerproxy | 
[**ListIndexerProxySchema**](IndexerProxyAPI.md#ListIndexerProxySchema) | **Get** /api/v1/indexerproxy/schema | 
[**TestIndexerProxy**](IndexerProxyAPI.md#TestIndexerProxy) | **Post** /api/v1/indexerproxy/test | 
[**TestallIndexerProxy**](IndexerProxyAPI.md#TestallIndexerProxy) | **Post** /api/v1/indexerproxy/testall | 
[**UpdateIndexerProxy**](IndexerProxyAPI.md#UpdateIndexerProxy) | **Put** /api/v1/indexerproxy/{id} | 



## CreateIndexerProxy

> IndexerProxyResource CreateIndexerProxy(ctx).ForceSave(forceSave).IndexerProxyResource(indexerProxyResource).Execute()



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
	forceSave := true // bool |  (optional) (default to false)
	indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexerProxyAPI.CreateIndexerProxy(context.Background()).ForceSave(forceSave).IndexerProxyResource(indexerProxyResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.CreateIndexerProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIndexerProxy`: IndexerProxyResource
	fmt.Fprintf(os.Stdout, "Response from `IndexerProxyAPI.CreateIndexerProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceSave** | **bool** |  | [default to false]
 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIndexerProxyActionByName

> CreateIndexerProxyActionByName(ctx, name).IndexerProxyResource(indexerProxyResource).Execute()



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
	name := "name_example" // string | 
	indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	r, err := apiClient.IndexerProxyAPI.CreateIndexerProxyActionByName(context.Background(), name).IndexerProxyResource(indexerProxyResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.CreateIndexerProxyActionByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndexerProxyActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndexerProxy

> DeleteIndexerProxy(ctx, id).Execute()



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
	r, err := apiClient.IndexerProxyAPI.DeleteIndexerProxy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.DeleteIndexerProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIndexerProxyById

> IndexerProxyResource GetIndexerProxyById(ctx, id).Execute()



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
	resp, r, err := apiClient.IndexerProxyAPI.GetIndexerProxyById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.GetIndexerProxyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexerProxyById`: IndexerProxyResource
	fmt.Fprintf(os.Stdout, "Response from `IndexerProxyAPI.GetIndexerProxyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexerProxyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexerProxy

> []IndexerProxyResource ListIndexerProxy(ctx).Execute()



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
	resp, r, err := apiClient.IndexerProxyAPI.ListIndexerProxy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.ListIndexerProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIndexerProxy`: []IndexerProxyResource
	fmt.Fprintf(os.Stdout, "Response from `IndexerProxyAPI.ListIndexerProxy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerProxyRequest struct via the builder pattern


### Return type

[**[]IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexerProxySchema

> []IndexerProxyResource ListIndexerProxySchema(ctx).Execute()



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
	resp, r, err := apiClient.IndexerProxyAPI.ListIndexerProxySchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.ListIndexerProxySchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIndexerProxySchema`: []IndexerProxyResource
	fmt.Fprintf(os.Stdout, "Response from `IndexerProxyAPI.ListIndexerProxySchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerProxySchemaRequest struct via the builder pattern


### Return type

[**[]IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIndexerProxy

> TestIndexerProxy(ctx).ForceTest(forceTest).IndexerProxyResource(indexerProxyResource).Execute()



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
	forceTest := true // bool |  (optional) (default to false)
	indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	r, err := apiClient.IndexerProxyAPI.TestIndexerProxy(context.Background()).ForceTest(forceTest).IndexerProxyResource(indexerProxyResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.TestIndexerProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestIndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceTest** | **bool** |  | [default to false]
 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestallIndexerProxy

> TestallIndexerProxy(ctx).Execute()



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
	r, err := apiClient.IndexerProxyAPI.TestallIndexerProxy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.TestallIndexerProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallIndexerProxyRequest struct via the builder pattern


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


## UpdateIndexerProxy

> IndexerProxyResource UpdateIndexerProxy(ctx, id).ForceSave(forceSave).IndexerProxyResource(indexerProxyResource).Execute()



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
	forceSave := true // bool |  (optional) (default to false)
	indexerProxyResource := *prowlarrClient.NewIndexerProxyResource() // IndexerProxyResource |  (optional)

	configuration := prowlarrClient.NewConfiguration()
	apiClient := prowlarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexerProxyAPI.UpdateIndexerProxy(context.Background(), id).ForceSave(forceSave).IndexerProxyResource(indexerProxyResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerProxyAPI.UpdateIndexerProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIndexerProxy`: IndexerProxyResource
	fmt.Fprintf(os.Stdout, "Response from `IndexerProxyAPI.UpdateIndexerProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndexerProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceSave** | **bool** |  | [default to false]
 **indexerProxyResource** | [**IndexerProxyResource**](IndexerProxyResource.md) |  | 

### Return type

[**IndexerProxyResource**](IndexerProxyResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

