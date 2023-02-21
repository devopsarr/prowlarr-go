# ReleaseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Guid** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **int32** |  | [optional] 
**AgeHours** | Pointer to **float64** |  | [optional] 
**AgeMinutes** | Pointer to **float64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Files** | Pointer to **NullableInt32** |  | [optional] 
**Grabs** | Pointer to **NullableInt32** |  | [optional] 
**IndexerId** | Pointer to **int32** |  | [optional] 
**Indexer** | Pointer to **NullableString** |  | [optional] 
**SubGroup** | Pointer to **NullableString** |  | [optional] 
**ReleaseHash** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**ImdbId** | Pointer to **int32** |  | [optional] 
**PublishDate** | Pointer to **time.Time** |  | [optional] 
**CommentUrl** | Pointer to **NullableString** |  | [optional] 
**DownloadUrl** | Pointer to **NullableString** |  | [optional] 
**InfoUrl** | Pointer to **NullableString** |  | [optional] 
**PosterUrl** | Pointer to **NullableString** |  | [optional] 
**IndexerFlags** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to [**[]IndexerCategory**](IndexerCategory.md) |  | [optional] 
**MagnetUrl** | Pointer to **NullableString** |  | [optional] 
**InfoHash** | Pointer to **NullableString** |  | [optional] 
**Seeders** | Pointer to **NullableInt32** |  | [optional] 
**Leechers** | Pointer to **NullableInt32** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewReleaseResource

`func NewReleaseResource() *ReleaseResource`

NewReleaseResource instantiates a new ReleaseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseResourceWithDefaults

`func NewReleaseResourceWithDefaults() *ReleaseResource`

NewReleaseResourceWithDefaults instantiates a new ReleaseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGuid

`func (o *ReleaseResource) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ReleaseResource) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ReleaseResource) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ReleaseResource) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *ReleaseResource) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *ReleaseResource) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetAge

`func (o *ReleaseResource) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ReleaseResource) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ReleaseResource) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ReleaseResource) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAgeHours

`func (o *ReleaseResource) GetAgeHours() float64`

GetAgeHours returns the AgeHours field if non-nil, zero value otherwise.

### GetAgeHoursOk

`func (o *ReleaseResource) GetAgeHoursOk() (*float64, bool)`

GetAgeHoursOk returns a tuple with the AgeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeHours

`func (o *ReleaseResource) SetAgeHours(v float64)`

SetAgeHours sets AgeHours field to given value.

### HasAgeHours

`func (o *ReleaseResource) HasAgeHours() bool`

HasAgeHours returns a boolean if a field has been set.

### GetAgeMinutes

`func (o *ReleaseResource) GetAgeMinutes() float64`

GetAgeMinutes returns the AgeMinutes field if non-nil, zero value otherwise.

### GetAgeMinutesOk

`func (o *ReleaseResource) GetAgeMinutesOk() (*float64, bool)`

GetAgeMinutesOk returns a tuple with the AgeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMinutes

`func (o *ReleaseResource) SetAgeMinutes(v float64)`

SetAgeMinutes sets AgeMinutes field to given value.

### HasAgeMinutes

`func (o *ReleaseResource) HasAgeMinutes() bool`

HasAgeMinutes returns a boolean if a field has been set.

### GetSize

`func (o *ReleaseResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ReleaseResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ReleaseResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ReleaseResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetFiles

`func (o *ReleaseResource) GetFiles() int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ReleaseResource) GetFilesOk() (*int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ReleaseResource) SetFiles(v int32)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ReleaseResource) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *ReleaseResource) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *ReleaseResource) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetGrabs

`func (o *ReleaseResource) GetGrabs() int32`

GetGrabs returns the Grabs field if non-nil, zero value otherwise.

### GetGrabsOk

`func (o *ReleaseResource) GetGrabsOk() (*int32, bool)`

GetGrabsOk returns a tuple with the Grabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabs

`func (o *ReleaseResource) SetGrabs(v int32)`

SetGrabs sets Grabs field to given value.

### HasGrabs

`func (o *ReleaseResource) HasGrabs() bool`

HasGrabs returns a boolean if a field has been set.

### SetGrabsNil

`func (o *ReleaseResource) SetGrabsNil(b bool)`

 SetGrabsNil sets the value for Grabs to be an explicit nil

### UnsetGrabs
`func (o *ReleaseResource) UnsetGrabs()`

UnsetGrabs ensures that no value is present for Grabs, not even an explicit nil
### GetIndexerId

`func (o *ReleaseResource) GetIndexerId() int32`

GetIndexerId returns the IndexerId field if non-nil, zero value otherwise.

### GetIndexerIdOk

`func (o *ReleaseResource) GetIndexerIdOk() (*int32, bool)`

GetIndexerIdOk returns a tuple with the IndexerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerId

`func (o *ReleaseResource) SetIndexerId(v int32)`

SetIndexerId sets IndexerId field to given value.

### HasIndexerId

`func (o *ReleaseResource) HasIndexerId() bool`

HasIndexerId returns a boolean if a field has been set.

### GetIndexer

`func (o *ReleaseResource) GetIndexer() string`

GetIndexer returns the Indexer field if non-nil, zero value otherwise.

### GetIndexerOk

`func (o *ReleaseResource) GetIndexerOk() (*string, bool)`

GetIndexerOk returns a tuple with the Indexer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexer

`func (o *ReleaseResource) SetIndexer(v string)`

SetIndexer sets Indexer field to given value.

### HasIndexer

`func (o *ReleaseResource) HasIndexer() bool`

HasIndexer returns a boolean if a field has been set.

### SetIndexerNil

`func (o *ReleaseResource) SetIndexerNil(b bool)`

 SetIndexerNil sets the value for Indexer to be an explicit nil

### UnsetIndexer
`func (o *ReleaseResource) UnsetIndexer()`

UnsetIndexer ensures that no value is present for Indexer, not even an explicit nil
### GetSubGroup

`func (o *ReleaseResource) GetSubGroup() string`

GetSubGroup returns the SubGroup field if non-nil, zero value otherwise.

### GetSubGroupOk

`func (o *ReleaseResource) GetSubGroupOk() (*string, bool)`

GetSubGroupOk returns a tuple with the SubGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGroup

`func (o *ReleaseResource) SetSubGroup(v string)`

SetSubGroup sets SubGroup field to given value.

### HasSubGroup

`func (o *ReleaseResource) HasSubGroup() bool`

HasSubGroup returns a boolean if a field has been set.

### SetSubGroupNil

`func (o *ReleaseResource) SetSubGroupNil(b bool)`

 SetSubGroupNil sets the value for SubGroup to be an explicit nil

### UnsetSubGroup
`func (o *ReleaseResource) UnsetSubGroup()`

UnsetSubGroup ensures that no value is present for SubGroup, not even an explicit nil
### GetReleaseHash

`func (o *ReleaseResource) GetReleaseHash() string`

GetReleaseHash returns the ReleaseHash field if non-nil, zero value otherwise.

### GetReleaseHashOk

`func (o *ReleaseResource) GetReleaseHashOk() (*string, bool)`

GetReleaseHashOk returns a tuple with the ReleaseHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHash

`func (o *ReleaseResource) SetReleaseHash(v string)`

SetReleaseHash sets ReleaseHash field to given value.

### HasReleaseHash

`func (o *ReleaseResource) HasReleaseHash() bool`

HasReleaseHash returns a boolean if a field has been set.

### SetReleaseHashNil

`func (o *ReleaseResource) SetReleaseHashNil(b bool)`

 SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil

### UnsetReleaseHash
`func (o *ReleaseResource) UnsetReleaseHash()`

UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
### GetTitle

`func (o *ReleaseResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ReleaseResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ReleaseResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetApproved

`func (o *ReleaseResource) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ReleaseResource) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ReleaseResource) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *ReleaseResource) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetImdbId

`func (o *ReleaseResource) GetImdbId() int32`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *ReleaseResource) GetImdbIdOk() (*int32, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *ReleaseResource) SetImdbId(v int32)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *ReleaseResource) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### GetPublishDate

`func (o *ReleaseResource) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *ReleaseResource) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *ReleaseResource) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *ReleaseResource) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### GetCommentUrl

`func (o *ReleaseResource) GetCommentUrl() string`

GetCommentUrl returns the CommentUrl field if non-nil, zero value otherwise.

### GetCommentUrlOk

`func (o *ReleaseResource) GetCommentUrlOk() (*string, bool)`

GetCommentUrlOk returns a tuple with the CommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentUrl

`func (o *ReleaseResource) SetCommentUrl(v string)`

SetCommentUrl sets CommentUrl field to given value.

### HasCommentUrl

`func (o *ReleaseResource) HasCommentUrl() bool`

HasCommentUrl returns a boolean if a field has been set.

### SetCommentUrlNil

`func (o *ReleaseResource) SetCommentUrlNil(b bool)`

 SetCommentUrlNil sets the value for CommentUrl to be an explicit nil

### UnsetCommentUrl
`func (o *ReleaseResource) UnsetCommentUrl()`

UnsetCommentUrl ensures that no value is present for CommentUrl, not even an explicit nil
### GetDownloadUrl

`func (o *ReleaseResource) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ReleaseResource) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ReleaseResource) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *ReleaseResource) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### SetDownloadUrlNil

`func (o *ReleaseResource) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *ReleaseResource) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetInfoUrl

`func (o *ReleaseResource) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *ReleaseResource) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *ReleaseResource) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *ReleaseResource) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### SetInfoUrlNil

`func (o *ReleaseResource) SetInfoUrlNil(b bool)`

 SetInfoUrlNil sets the value for InfoUrl to be an explicit nil

### UnsetInfoUrl
`func (o *ReleaseResource) UnsetInfoUrl()`

UnsetInfoUrl ensures that no value is present for InfoUrl, not even an explicit nil
### GetPosterUrl

`func (o *ReleaseResource) GetPosterUrl() string`

GetPosterUrl returns the PosterUrl field if non-nil, zero value otherwise.

### GetPosterUrlOk

`func (o *ReleaseResource) GetPosterUrlOk() (*string, bool)`

GetPosterUrlOk returns a tuple with the PosterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosterUrl

`func (o *ReleaseResource) SetPosterUrl(v string)`

SetPosterUrl sets PosterUrl field to given value.

### HasPosterUrl

`func (o *ReleaseResource) HasPosterUrl() bool`

HasPosterUrl returns a boolean if a field has been set.

### SetPosterUrlNil

`func (o *ReleaseResource) SetPosterUrlNil(b bool)`

 SetPosterUrlNil sets the value for PosterUrl to be an explicit nil

### UnsetPosterUrl
`func (o *ReleaseResource) UnsetPosterUrl()`

UnsetPosterUrl ensures that no value is present for PosterUrl, not even an explicit nil
### GetIndexerFlags

`func (o *ReleaseResource) GetIndexerFlags() []string`

GetIndexerFlags returns the IndexerFlags field if non-nil, zero value otherwise.

### GetIndexerFlagsOk

`func (o *ReleaseResource) GetIndexerFlagsOk() (*[]string, bool)`

GetIndexerFlagsOk returns a tuple with the IndexerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerFlags

`func (o *ReleaseResource) SetIndexerFlags(v []string)`

SetIndexerFlags sets IndexerFlags field to given value.

### HasIndexerFlags

`func (o *ReleaseResource) HasIndexerFlags() bool`

HasIndexerFlags returns a boolean if a field has been set.

### SetIndexerFlagsNil

`func (o *ReleaseResource) SetIndexerFlagsNil(b bool)`

 SetIndexerFlagsNil sets the value for IndexerFlags to be an explicit nil

### UnsetIndexerFlags
`func (o *ReleaseResource) UnsetIndexerFlags()`

UnsetIndexerFlags ensures that no value is present for IndexerFlags, not even an explicit nil
### GetCategories

`func (o *ReleaseResource) GetCategories() []IndexerCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ReleaseResource) GetCategoriesOk() (*[]IndexerCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ReleaseResource) SetCategories(v []IndexerCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ReleaseResource) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *ReleaseResource) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *ReleaseResource) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetMagnetUrl

`func (o *ReleaseResource) GetMagnetUrl() string`

GetMagnetUrl returns the MagnetUrl field if non-nil, zero value otherwise.

### GetMagnetUrlOk

`func (o *ReleaseResource) GetMagnetUrlOk() (*string, bool)`

GetMagnetUrlOk returns a tuple with the MagnetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnetUrl

`func (o *ReleaseResource) SetMagnetUrl(v string)`

SetMagnetUrl sets MagnetUrl field to given value.

### HasMagnetUrl

`func (o *ReleaseResource) HasMagnetUrl() bool`

HasMagnetUrl returns a boolean if a field has been set.

### SetMagnetUrlNil

`func (o *ReleaseResource) SetMagnetUrlNil(b bool)`

 SetMagnetUrlNil sets the value for MagnetUrl to be an explicit nil

### UnsetMagnetUrl
`func (o *ReleaseResource) UnsetMagnetUrl()`

UnsetMagnetUrl ensures that no value is present for MagnetUrl, not even an explicit nil
### GetInfoHash

`func (o *ReleaseResource) GetInfoHash() string`

GetInfoHash returns the InfoHash field if non-nil, zero value otherwise.

### GetInfoHashOk

`func (o *ReleaseResource) GetInfoHashOk() (*string, bool)`

GetInfoHashOk returns a tuple with the InfoHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoHash

`func (o *ReleaseResource) SetInfoHash(v string)`

SetInfoHash sets InfoHash field to given value.

### HasInfoHash

`func (o *ReleaseResource) HasInfoHash() bool`

HasInfoHash returns a boolean if a field has been set.

### SetInfoHashNil

`func (o *ReleaseResource) SetInfoHashNil(b bool)`

 SetInfoHashNil sets the value for InfoHash to be an explicit nil

### UnsetInfoHash
`func (o *ReleaseResource) UnsetInfoHash()`

UnsetInfoHash ensures that no value is present for InfoHash, not even an explicit nil
### GetSeeders

`func (o *ReleaseResource) GetSeeders() int32`

GetSeeders returns the Seeders field if non-nil, zero value otherwise.

### GetSeedersOk

`func (o *ReleaseResource) GetSeedersOk() (*int32, bool)`

GetSeedersOk returns a tuple with the Seeders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeders

`func (o *ReleaseResource) SetSeeders(v int32)`

SetSeeders sets Seeders field to given value.

### HasSeeders

`func (o *ReleaseResource) HasSeeders() bool`

HasSeeders returns a boolean if a field has been set.

### SetSeedersNil

`func (o *ReleaseResource) SetSeedersNil(b bool)`

 SetSeedersNil sets the value for Seeders to be an explicit nil

### UnsetSeeders
`func (o *ReleaseResource) UnsetSeeders()`

UnsetSeeders ensures that no value is present for Seeders, not even an explicit nil
### GetLeechers

`func (o *ReleaseResource) GetLeechers() int32`

GetLeechers returns the Leechers field if non-nil, zero value otherwise.

### GetLeechersOk

`func (o *ReleaseResource) GetLeechersOk() (*int32, bool)`

GetLeechersOk returns a tuple with the Leechers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeechers

`func (o *ReleaseResource) SetLeechers(v int32)`

SetLeechers sets Leechers field to given value.

### HasLeechers

`func (o *ReleaseResource) HasLeechers() bool`

HasLeechers returns a boolean if a field has been set.

### SetLeechersNil

`func (o *ReleaseResource) SetLeechersNil(b bool)`

 SetLeechersNil sets the value for Leechers to be an explicit nil

### UnsetLeechers
`func (o *ReleaseResource) UnsetLeechers()`

UnsetLeechers ensures that no value is present for Leechers, not even an explicit nil
### GetProtocol

`func (o *ReleaseResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ReleaseResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ReleaseResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ReleaseResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetFileName

`func (o *ReleaseResource) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ReleaseResource) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ReleaseResource) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ReleaseResource) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ReleaseResource) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ReleaseResource) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


