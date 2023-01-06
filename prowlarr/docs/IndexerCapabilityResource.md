# IndexerCapabilityResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LimitsMax** | Pointer to **NullableInt32** |  | [optional] 
**LimitsDefault** | Pointer to **NullableInt32** |  | [optional] 
**Categories** | Pointer to [**[]IndexerCategory**](IndexerCategory.md) |  | [optional] 
**SupportsRawSearch** | Pointer to **bool** |  | [optional] 
**SearchParams** | Pointer to [**[]SearchParam**](SearchParam.md) |  | [optional] 
**TvSearchParams** | Pointer to [**[]TvSearchParam**](TvSearchParam.md) |  | [optional] 
**MovieSearchParams** | Pointer to [**[]MovieSearchParam**](MovieSearchParam.md) |  | [optional] 
**MusicSearchParams** | Pointer to [**[]MusicSearchParam**](MusicSearchParam.md) |  | [optional] 
**BookSearchParams** | Pointer to [**[]BookSearchParam**](BookSearchParam.md) |  | [optional] 

## Methods

### NewIndexerCapabilityResource

`func NewIndexerCapabilityResource() *IndexerCapabilityResource`

NewIndexerCapabilityResource instantiates a new IndexerCapabilityResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerCapabilityResourceWithDefaults

`func NewIndexerCapabilityResourceWithDefaults() *IndexerCapabilityResource`

NewIndexerCapabilityResourceWithDefaults instantiates a new IndexerCapabilityResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexerCapabilityResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerCapabilityResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerCapabilityResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerCapabilityResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimitsMax

`func (o *IndexerCapabilityResource) GetLimitsMax() int32`

GetLimitsMax returns the LimitsMax field if non-nil, zero value otherwise.

### GetLimitsMaxOk

`func (o *IndexerCapabilityResource) GetLimitsMaxOk() (*int32, bool)`

GetLimitsMaxOk returns a tuple with the LimitsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitsMax

`func (o *IndexerCapabilityResource) SetLimitsMax(v int32)`

SetLimitsMax sets LimitsMax field to given value.

### HasLimitsMax

`func (o *IndexerCapabilityResource) HasLimitsMax() bool`

HasLimitsMax returns a boolean if a field has been set.

### SetLimitsMaxNil

`func (o *IndexerCapabilityResource) SetLimitsMaxNil(b bool)`

 SetLimitsMaxNil sets the value for LimitsMax to be an explicit nil

### UnsetLimitsMax
`func (o *IndexerCapabilityResource) UnsetLimitsMax()`

UnsetLimitsMax ensures that no value is present for LimitsMax, not even an explicit nil
### GetLimitsDefault

`func (o *IndexerCapabilityResource) GetLimitsDefault() int32`

GetLimitsDefault returns the LimitsDefault field if non-nil, zero value otherwise.

### GetLimitsDefaultOk

`func (o *IndexerCapabilityResource) GetLimitsDefaultOk() (*int32, bool)`

GetLimitsDefaultOk returns a tuple with the LimitsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitsDefault

`func (o *IndexerCapabilityResource) SetLimitsDefault(v int32)`

SetLimitsDefault sets LimitsDefault field to given value.

### HasLimitsDefault

`func (o *IndexerCapabilityResource) HasLimitsDefault() bool`

HasLimitsDefault returns a boolean if a field has been set.

### SetLimitsDefaultNil

`func (o *IndexerCapabilityResource) SetLimitsDefaultNil(b bool)`

 SetLimitsDefaultNil sets the value for LimitsDefault to be an explicit nil

### UnsetLimitsDefault
`func (o *IndexerCapabilityResource) UnsetLimitsDefault()`

UnsetLimitsDefault ensures that no value is present for LimitsDefault, not even an explicit nil
### GetCategories

`func (o *IndexerCapabilityResource) GetCategories() []IndexerCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *IndexerCapabilityResource) GetCategoriesOk() (*[]IndexerCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *IndexerCapabilityResource) SetCategories(v []IndexerCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *IndexerCapabilityResource) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *IndexerCapabilityResource) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *IndexerCapabilityResource) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetSupportsRawSearch

`func (o *IndexerCapabilityResource) GetSupportsRawSearch() bool`

GetSupportsRawSearch returns the SupportsRawSearch field if non-nil, zero value otherwise.

### GetSupportsRawSearchOk

`func (o *IndexerCapabilityResource) GetSupportsRawSearchOk() (*bool, bool)`

GetSupportsRawSearchOk returns a tuple with the SupportsRawSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRawSearch

`func (o *IndexerCapabilityResource) SetSupportsRawSearch(v bool)`

SetSupportsRawSearch sets SupportsRawSearch field to given value.

### HasSupportsRawSearch

`func (o *IndexerCapabilityResource) HasSupportsRawSearch() bool`

HasSupportsRawSearch returns a boolean if a field has been set.

### GetSearchParams

`func (o *IndexerCapabilityResource) GetSearchParams() []SearchParam`

GetSearchParams returns the SearchParams field if non-nil, zero value otherwise.

### GetSearchParamsOk

`func (o *IndexerCapabilityResource) GetSearchParamsOk() (*[]SearchParam, bool)`

GetSearchParamsOk returns a tuple with the SearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchParams

`func (o *IndexerCapabilityResource) SetSearchParams(v []SearchParam)`

SetSearchParams sets SearchParams field to given value.

### HasSearchParams

`func (o *IndexerCapabilityResource) HasSearchParams() bool`

HasSearchParams returns a boolean if a field has been set.

### SetSearchParamsNil

`func (o *IndexerCapabilityResource) SetSearchParamsNil(b bool)`

 SetSearchParamsNil sets the value for SearchParams to be an explicit nil

### UnsetSearchParams
`func (o *IndexerCapabilityResource) UnsetSearchParams()`

UnsetSearchParams ensures that no value is present for SearchParams, not even an explicit nil
### GetTvSearchParams

`func (o *IndexerCapabilityResource) GetTvSearchParams() []TvSearchParam`

GetTvSearchParams returns the TvSearchParams field if non-nil, zero value otherwise.

### GetTvSearchParamsOk

`func (o *IndexerCapabilityResource) GetTvSearchParamsOk() (*[]TvSearchParam, bool)`

GetTvSearchParamsOk returns a tuple with the TvSearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvSearchParams

`func (o *IndexerCapabilityResource) SetTvSearchParams(v []TvSearchParam)`

SetTvSearchParams sets TvSearchParams field to given value.

### HasTvSearchParams

`func (o *IndexerCapabilityResource) HasTvSearchParams() bool`

HasTvSearchParams returns a boolean if a field has been set.

### SetTvSearchParamsNil

`func (o *IndexerCapabilityResource) SetTvSearchParamsNil(b bool)`

 SetTvSearchParamsNil sets the value for TvSearchParams to be an explicit nil

### UnsetTvSearchParams
`func (o *IndexerCapabilityResource) UnsetTvSearchParams()`

UnsetTvSearchParams ensures that no value is present for TvSearchParams, not even an explicit nil
### GetMovieSearchParams

`func (o *IndexerCapabilityResource) GetMovieSearchParams() []MovieSearchParam`

GetMovieSearchParams returns the MovieSearchParams field if non-nil, zero value otherwise.

### GetMovieSearchParamsOk

`func (o *IndexerCapabilityResource) GetMovieSearchParamsOk() (*[]MovieSearchParam, bool)`

GetMovieSearchParamsOk returns a tuple with the MovieSearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieSearchParams

`func (o *IndexerCapabilityResource) SetMovieSearchParams(v []MovieSearchParam)`

SetMovieSearchParams sets MovieSearchParams field to given value.

### HasMovieSearchParams

`func (o *IndexerCapabilityResource) HasMovieSearchParams() bool`

HasMovieSearchParams returns a boolean if a field has been set.

### SetMovieSearchParamsNil

`func (o *IndexerCapabilityResource) SetMovieSearchParamsNil(b bool)`

 SetMovieSearchParamsNil sets the value for MovieSearchParams to be an explicit nil

### UnsetMovieSearchParams
`func (o *IndexerCapabilityResource) UnsetMovieSearchParams()`

UnsetMovieSearchParams ensures that no value is present for MovieSearchParams, not even an explicit nil
### GetMusicSearchParams

`func (o *IndexerCapabilityResource) GetMusicSearchParams() []MusicSearchParam`

GetMusicSearchParams returns the MusicSearchParams field if non-nil, zero value otherwise.

### GetMusicSearchParamsOk

`func (o *IndexerCapabilityResource) GetMusicSearchParamsOk() (*[]MusicSearchParam, bool)`

GetMusicSearchParamsOk returns a tuple with the MusicSearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicSearchParams

`func (o *IndexerCapabilityResource) SetMusicSearchParams(v []MusicSearchParam)`

SetMusicSearchParams sets MusicSearchParams field to given value.

### HasMusicSearchParams

`func (o *IndexerCapabilityResource) HasMusicSearchParams() bool`

HasMusicSearchParams returns a boolean if a field has been set.

### SetMusicSearchParamsNil

`func (o *IndexerCapabilityResource) SetMusicSearchParamsNil(b bool)`

 SetMusicSearchParamsNil sets the value for MusicSearchParams to be an explicit nil

### UnsetMusicSearchParams
`func (o *IndexerCapabilityResource) UnsetMusicSearchParams()`

UnsetMusicSearchParams ensures that no value is present for MusicSearchParams, not even an explicit nil
### GetBookSearchParams

`func (o *IndexerCapabilityResource) GetBookSearchParams() []BookSearchParam`

GetBookSearchParams returns the BookSearchParams field if non-nil, zero value otherwise.

### GetBookSearchParamsOk

`func (o *IndexerCapabilityResource) GetBookSearchParamsOk() (*[]BookSearchParam, bool)`

GetBookSearchParamsOk returns a tuple with the BookSearchParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookSearchParams

`func (o *IndexerCapabilityResource) SetBookSearchParams(v []BookSearchParam)`

SetBookSearchParams sets BookSearchParams field to given value.

### HasBookSearchParams

`func (o *IndexerCapabilityResource) HasBookSearchParams() bool`

HasBookSearchParams returns a boolean if a field has been set.

### SetBookSearchParamsNil

`func (o *IndexerCapabilityResource) SetBookSearchParamsNil(b bool)`

 SetBookSearchParamsNil sets the value for BookSearchParams to be an explicit nil

### UnsetBookSearchParams
`func (o *IndexerCapabilityResource) UnsetBookSearchParams()`

UnsetBookSearchParams ensures that no value is present for BookSearchParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


