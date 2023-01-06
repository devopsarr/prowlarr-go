# \NewznabApi

All URIs are relative to *http://localhost:9696*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1IndexeridDownload**](NewznabApi.md#GetApiV1IndexeridDownload) | **Get** /api/v1/indexer/{id}/download | 
[**GetApiV1IndexeridNewznab**](NewznabApi.md#GetApiV1IndexeridNewznab) | **Get** /api/v1/indexer/{id}/newznab | 
[**GetidApi**](NewznabApi.md#GetidApi) | **Get** /{id}/api | 
[**GetidDownload**](NewznabApi.md#GetidDownload) | **Get** /{id}/download | 



## GetApiV1IndexeridDownload

> GetApiV1IndexeridDownload(ctx, id).Link(link).File(file).Execute()



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
    id := int32(56) // int32 | 
    link := "link_example" // string |  (optional)
    file := "file_example" // string |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NewznabApi.GetApiV1IndexeridDownload(context.Background(), id).Link(link).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NewznabApi.GetApiV1IndexeridDownload``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiV1IndexeridDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **link** | **string** |  | 
 **file** | **string** |  | 

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


## GetApiV1IndexeridNewznab

> GetApiV1IndexeridNewznab(ctx, id).T(t).Q(q).Cat(cat).Imdbid(imdbid).Tmdbid(tmdbid).Extended(extended).Limit(limit).Offset(offset).Rid(rid).Tvmazeid(tvmazeid).Traktid(traktid).Tvdbid(tvdbid).Doubanid(doubanid).Season(season).Ep(ep).Album(album).Artist(artist).Label(label).Track(track).Year(year).Genre(genre).Author(author).Title(title).Publisher(publisher).Configured(configured).Source(source).Host(host).Server(server).Execute()



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
    id := int32(56) // int32 | 
    t := "t_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    cat := "cat_example" // string |  (optional)
    imdbid := "imdbid_example" // string |  (optional)
    tmdbid := int32(56) // int32 |  (optional)
    extended := "extended_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    rid := int32(56) // int32 |  (optional)
    tvmazeid := int32(56) // int32 |  (optional)
    traktid := int32(56) // int32 |  (optional)
    tvdbid := int32(56) // int32 |  (optional)
    doubanid := int32(56) // int32 |  (optional)
    season := int32(56) // int32 |  (optional)
    ep := "ep_example" // string |  (optional)
    album := "album_example" // string |  (optional)
    artist := "artist_example" // string |  (optional)
    label := "label_example" // string |  (optional)
    track := "track_example" // string |  (optional)
    year := int32(56) // int32 |  (optional)
    genre := "genre_example" // string |  (optional)
    author := "author_example" // string |  (optional)
    title := "title_example" // string |  (optional)
    publisher := "publisher_example" // string |  (optional)
    configured := "configured_example" // string |  (optional)
    source := "source_example" // string |  (optional)
    host := "host_example" // string |  (optional)
    server := "server_example" // string |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NewznabApi.GetApiV1IndexeridNewznab(context.Background(), id).T(t).Q(q).Cat(cat).Imdbid(imdbid).Tmdbid(tmdbid).Extended(extended).Limit(limit).Offset(offset).Rid(rid).Tvmazeid(tvmazeid).Traktid(traktid).Tvdbid(tvdbid).Doubanid(doubanid).Season(season).Ep(ep).Album(album).Artist(artist).Label(label).Track(track).Year(year).Genre(genre).Author(author).Title(title).Publisher(publisher).Configured(configured).Source(source).Host(host).Server(server).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NewznabApi.GetApiV1IndexeridNewznab``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiV1IndexeridNewznabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t** | **string** |  | 
 **q** | **string** |  | 
 **cat** | **string** |  | 
 **imdbid** | **string** |  | 
 **tmdbid** | **int32** |  | 
 **extended** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **rid** | **int32** |  | 
 **tvmazeid** | **int32** |  | 
 **traktid** | **int32** |  | 
 **tvdbid** | **int32** |  | 
 **doubanid** | **int32** |  | 
 **season** | **int32** |  | 
 **ep** | **string** |  | 
 **album** | **string** |  | 
 **artist** | **string** |  | 
 **label** | **string** |  | 
 **track** | **string** |  | 
 **year** | **int32** |  | 
 **genre** | **string** |  | 
 **author** | **string** |  | 
 **title** | **string** |  | 
 **publisher** | **string** |  | 
 **configured** | **string** |  | 
 **source** | **string** |  | 
 **host** | **string** |  | 
 **server** | **string** |  | 

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


## GetidApi

> GetidApi(ctx, id).T(t).Q(q).Cat(cat).Imdbid(imdbid).Tmdbid(tmdbid).Extended(extended).Limit(limit).Offset(offset).Rid(rid).Tvmazeid(tvmazeid).Traktid(traktid).Tvdbid(tvdbid).Doubanid(doubanid).Season(season).Ep(ep).Album(album).Artist(artist).Label(label).Track(track).Year(year).Genre(genre).Author(author).Title(title).Publisher(publisher).Configured(configured).Source(source).Host(host).Server(server).Execute()



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
    id := int32(56) // int32 | 
    t := "t_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    cat := "cat_example" // string |  (optional)
    imdbid := "imdbid_example" // string |  (optional)
    tmdbid := int32(56) // int32 |  (optional)
    extended := "extended_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    rid := int32(56) // int32 |  (optional)
    tvmazeid := int32(56) // int32 |  (optional)
    traktid := int32(56) // int32 |  (optional)
    tvdbid := int32(56) // int32 |  (optional)
    doubanid := int32(56) // int32 |  (optional)
    season := int32(56) // int32 |  (optional)
    ep := "ep_example" // string |  (optional)
    album := "album_example" // string |  (optional)
    artist := "artist_example" // string |  (optional)
    label := "label_example" // string |  (optional)
    track := "track_example" // string |  (optional)
    year := int32(56) // int32 |  (optional)
    genre := "genre_example" // string |  (optional)
    author := "author_example" // string |  (optional)
    title := "title_example" // string |  (optional)
    publisher := "publisher_example" // string |  (optional)
    configured := "configured_example" // string |  (optional)
    source := "source_example" // string |  (optional)
    host := "host_example" // string |  (optional)
    server := "server_example" // string |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NewznabApi.GetidApi(context.Background(), id).T(t).Q(q).Cat(cat).Imdbid(imdbid).Tmdbid(tmdbid).Extended(extended).Limit(limit).Offset(offset).Rid(rid).Tvmazeid(tvmazeid).Traktid(traktid).Tvdbid(tvdbid).Doubanid(doubanid).Season(season).Ep(ep).Album(album).Artist(artist).Label(label).Track(track).Year(year).Genre(genre).Author(author).Title(title).Publisher(publisher).Configured(configured).Source(source).Host(host).Server(server).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NewznabApi.GetidApi``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetidApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t** | **string** |  | 
 **q** | **string** |  | 
 **cat** | **string** |  | 
 **imdbid** | **string** |  | 
 **tmdbid** | **int32** |  | 
 **extended** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **rid** | **int32** |  | 
 **tvmazeid** | **int32** |  | 
 **traktid** | **int32** |  | 
 **tvdbid** | **int32** |  | 
 **doubanid** | **int32** |  | 
 **season** | **int32** |  | 
 **ep** | **string** |  | 
 **album** | **string** |  | 
 **artist** | **string** |  | 
 **label** | **string** |  | 
 **track** | **string** |  | 
 **year** | **int32** |  | 
 **genre** | **string** |  | 
 **author** | **string** |  | 
 **title** | **string** |  | 
 **publisher** | **string** |  | 
 **configured** | **string** |  | 
 **source** | **string** |  | 
 **host** | **string** |  | 
 **server** | **string** |  | 

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


## GetidDownload

> GetidDownload(ctx, id).Link(link).File(file).Execute()



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
    id := int32(56) // int32 | 
    link := "link_example" // string |  (optional)
    file := "file_example" // string |  (optional)

    configuration := prowlarrClient.NewConfiguration()
    apiClient := prowlarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.NewznabApi.GetidDownload(context.Background(), id).Link(link).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NewznabApi.GetidDownload``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetidDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **link** | **string** |  | 
 **file** | **string** |  | 

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

