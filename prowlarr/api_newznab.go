/*
Prowlarr

Prowlarr API docs

API version: v1.23.1.4708
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// NewznabAPIService NewznabAPI service
type NewznabAPIService service

type ApiGetIndexerDownloadRequest struct {
	ctx context.Context
	ApiService *NewznabAPIService
	id int32
	link *string
	file *string
}

func (r ApiGetIndexerDownloadRequest) Link(link string) ApiGetIndexerDownloadRequest {
	r.link = &link
	return r
}

func (r ApiGetIndexerDownloadRequest) File(file string) ApiGetIndexerDownloadRequest {
	r.file = &file
	return r
}

func (r ApiGetIndexerDownloadRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetIndexerDownloadExecute(r)
}

/*
GetIndexerDownload Method for GetIndexerDownload

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetIndexerDownloadRequest
*/
func (a *NewznabAPIService) GetIndexerDownload(ctx context.Context, id int32) ApiGetIndexerDownloadRequest {
	return ApiGetIndexerDownloadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *NewznabAPIService) GetIndexerDownloadExecute(r ApiGetIndexerDownloadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewznabAPIService.GetIndexerDownload")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/indexer/{id}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.link != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "link", r.link, "form", "")
	}
	if r.file != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "file", r.file, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetIndexerNewznabRequest struct {
	ctx context.Context
	ApiService *NewznabAPIService
	id int32
	t *string
	q *string
	cat *string
	imdbid *string
	tmdbid *int32
	extended *string
	limit *int32
	offset *int32
	minage *int32
	maxage *int32
	minsize *int64
	maxsize *int64
	rid *int32
	tvmazeid *int32
	traktid *int32
	tvdbid *int32
	doubanid *int32
	season *int32
	ep *string
	album *string
	artist *string
	label *string
	track *string
	year *int32
	genre *string
	author *string
	title *string
	publisher *string
	configured *string
	source *string
	host *string
	server *string
}

func (r ApiGetIndexerNewznabRequest) T(t string) ApiGetIndexerNewznabRequest {
	r.t = &t
	return r
}

func (r ApiGetIndexerNewznabRequest) Q(q string) ApiGetIndexerNewznabRequest {
	r.q = &q
	return r
}

func (r ApiGetIndexerNewznabRequest) Cat(cat string) ApiGetIndexerNewznabRequest {
	r.cat = &cat
	return r
}

func (r ApiGetIndexerNewznabRequest) Imdbid(imdbid string) ApiGetIndexerNewznabRequest {
	r.imdbid = &imdbid
	return r
}

func (r ApiGetIndexerNewznabRequest) Tmdbid(tmdbid int32) ApiGetIndexerNewznabRequest {
	r.tmdbid = &tmdbid
	return r
}

func (r ApiGetIndexerNewznabRequest) Extended(extended string) ApiGetIndexerNewznabRequest {
	r.extended = &extended
	return r
}

func (r ApiGetIndexerNewznabRequest) Limit(limit int32) ApiGetIndexerNewznabRequest {
	r.limit = &limit
	return r
}

func (r ApiGetIndexerNewznabRequest) Offset(offset int32) ApiGetIndexerNewznabRequest {
	r.offset = &offset
	return r
}

func (r ApiGetIndexerNewznabRequest) Minage(minage int32) ApiGetIndexerNewznabRequest {
	r.minage = &minage
	return r
}

func (r ApiGetIndexerNewznabRequest) Maxage(maxage int32) ApiGetIndexerNewznabRequest {
	r.maxage = &maxage
	return r
}

func (r ApiGetIndexerNewznabRequest) Minsize(minsize int64) ApiGetIndexerNewznabRequest {
	r.minsize = &minsize
	return r
}

func (r ApiGetIndexerNewznabRequest) Maxsize(maxsize int64) ApiGetIndexerNewznabRequest {
	r.maxsize = &maxsize
	return r
}

func (r ApiGetIndexerNewznabRequest) Rid(rid int32) ApiGetIndexerNewznabRequest {
	r.rid = &rid
	return r
}

func (r ApiGetIndexerNewznabRequest) Tvmazeid(tvmazeid int32) ApiGetIndexerNewznabRequest {
	r.tvmazeid = &tvmazeid
	return r
}

func (r ApiGetIndexerNewznabRequest) Traktid(traktid int32) ApiGetIndexerNewznabRequest {
	r.traktid = &traktid
	return r
}

func (r ApiGetIndexerNewznabRequest) Tvdbid(tvdbid int32) ApiGetIndexerNewznabRequest {
	r.tvdbid = &tvdbid
	return r
}

func (r ApiGetIndexerNewznabRequest) Doubanid(doubanid int32) ApiGetIndexerNewznabRequest {
	r.doubanid = &doubanid
	return r
}

func (r ApiGetIndexerNewznabRequest) Season(season int32) ApiGetIndexerNewznabRequest {
	r.season = &season
	return r
}

func (r ApiGetIndexerNewznabRequest) Ep(ep string) ApiGetIndexerNewznabRequest {
	r.ep = &ep
	return r
}

func (r ApiGetIndexerNewznabRequest) Album(album string) ApiGetIndexerNewznabRequest {
	r.album = &album
	return r
}

func (r ApiGetIndexerNewznabRequest) Artist(artist string) ApiGetIndexerNewznabRequest {
	r.artist = &artist
	return r
}

func (r ApiGetIndexerNewznabRequest) Label(label string) ApiGetIndexerNewznabRequest {
	r.label = &label
	return r
}

func (r ApiGetIndexerNewznabRequest) Track(track string) ApiGetIndexerNewznabRequest {
	r.track = &track
	return r
}

func (r ApiGetIndexerNewznabRequest) Year(year int32) ApiGetIndexerNewznabRequest {
	r.year = &year
	return r
}

func (r ApiGetIndexerNewznabRequest) Genre(genre string) ApiGetIndexerNewznabRequest {
	r.genre = &genre
	return r
}

func (r ApiGetIndexerNewznabRequest) Author(author string) ApiGetIndexerNewznabRequest {
	r.author = &author
	return r
}

func (r ApiGetIndexerNewznabRequest) Title(title string) ApiGetIndexerNewznabRequest {
	r.title = &title
	return r
}

func (r ApiGetIndexerNewznabRequest) Publisher(publisher string) ApiGetIndexerNewznabRequest {
	r.publisher = &publisher
	return r
}

func (r ApiGetIndexerNewznabRequest) Configured(configured string) ApiGetIndexerNewznabRequest {
	r.configured = &configured
	return r
}

func (r ApiGetIndexerNewznabRequest) Source(source string) ApiGetIndexerNewznabRequest {
	r.source = &source
	return r
}

func (r ApiGetIndexerNewznabRequest) Host(host string) ApiGetIndexerNewznabRequest {
	r.host = &host
	return r
}

func (r ApiGetIndexerNewznabRequest) Server(server string) ApiGetIndexerNewznabRequest {
	r.server = &server
	return r
}

func (r ApiGetIndexerNewznabRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetIndexerNewznabExecute(r)
}

/*
GetIndexerNewznab Method for GetIndexerNewznab

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetIndexerNewznabRequest
*/
func (a *NewznabAPIService) GetIndexerNewznab(ctx context.Context, id int32) ApiGetIndexerNewznabRequest {
	return ApiGetIndexerNewznabRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *NewznabAPIService) GetIndexerNewznabExecute(r ApiGetIndexerNewznabRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewznabAPIService.GetIndexerNewznab")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/indexer/{id}/newznab"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t", r.t, "form", "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.cat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cat", r.cat, "form", "")
	}
	if r.imdbid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "imdbid", r.imdbid, "form", "")
	}
	if r.tmdbid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tmdbid", r.tmdbid, "form", "")
	}
	if r.extended != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "extended", r.extended, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.minage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minage", r.minage, "form", "")
	}
	if r.maxage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxage", r.maxage, "form", "")
	}
	if r.minsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minsize", r.minsize, "form", "")
	}
	if r.maxsize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxsize", r.maxsize, "form", "")
	}
	if r.rid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rid", r.rid, "form", "")
	}
	if r.tvmazeid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tvmazeid", r.tvmazeid, "form", "")
	}
	if r.traktid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "traktid", r.traktid, "form", "")
	}
	if r.tvdbid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tvdbid", r.tvdbid, "form", "")
	}
	if r.doubanid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "doubanid", r.doubanid, "form", "")
	}
	if r.season != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "season", r.season, "form", "")
	}
	if r.ep != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ep", r.ep, "form", "")
	}
	if r.album != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "album", r.album, "form", "")
	}
	if r.artist != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "artist", r.artist, "form", "")
	}
	if r.label != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "label", r.label, "form", "")
	}
	if r.track != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "track", r.track, "form", "")
	}
	if r.year != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "year", r.year, "form", "")
	}
	if r.genre != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "genre", r.genre, "form", "")
	}
	if r.author != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "author", r.author, "form", "")
	}
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "title", r.title, "form", "")
	}
	if r.publisher != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "publisher", r.publisher, "form", "")
	}
	if r.configured != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "configured", r.configured, "form", "")
	}
	if r.source != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "source", r.source, "form", "")
	}
	if r.host != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "host", r.host, "form", "")
	}
	if r.server != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "server", r.server, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
