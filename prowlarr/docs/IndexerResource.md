# IndexerResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**ImplementationName** | Pointer to **NullableString** |  | [optional] 
**Implementation** | Pointer to **NullableString** |  | [optional] 
**ConfigContract** | Pointer to **NullableString** |  | [optional] 
**InfoLink** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to [**ProviderMessage**](ProviderMessage.md) |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Presets** | Pointer to [**[]IndexerResource**](IndexerResource.md) |  | [optional] 
**IndexerUrls** | Pointer to **[]string** |  | [optional] 
**LegacyUrls** | Pointer to **[]string** |  | [optional] 
**DefinitionName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**Encoding** | Pointer to **NullableString** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**Redirect** | Pointer to **bool** |  | [optional] 
**SupportsRss** | Pointer to **bool** |  | [optional] 
**SupportsSearch** | Pointer to **bool** |  | [optional] 
**SupportsRedirect** | Pointer to **bool** |  | [optional] 
**SupportsPagination** | Pointer to **bool** |  | [optional] 
**AppProfileId** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**Privacy** | Pointer to [**IndexerPrivacy**](IndexerPrivacy.md) |  | [optional] 
**Capabilities** | Pointer to [**IndexerCapabilityResource**](IndexerCapabilityResource.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**IndexerStatusResource**](IndexerStatusResource.md) |  | [optional] 
**SortName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIndexerResource

`func NewIndexerResource() *IndexerResource`

NewIndexerResource instantiates a new IndexerResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerResourceWithDefaults

`func NewIndexerResourceWithDefaults() *IndexerResource`

NewIndexerResourceWithDefaults instantiates a new IndexerResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexerResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IndexerResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndexerResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndexerResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IndexerResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IndexerResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IndexerResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFields

`func (o *IndexerResource) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IndexerResource) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IndexerResource) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IndexerResource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *IndexerResource) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *IndexerResource) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetImplementationName

`func (o *IndexerResource) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *IndexerResource) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *IndexerResource) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *IndexerResource) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *IndexerResource) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *IndexerResource) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetImplementation

`func (o *IndexerResource) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *IndexerResource) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *IndexerResource) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *IndexerResource) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *IndexerResource) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *IndexerResource) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetConfigContract

`func (o *IndexerResource) GetConfigContract() string`

GetConfigContract returns the ConfigContract field if non-nil, zero value otherwise.

### GetConfigContractOk

`func (o *IndexerResource) GetConfigContractOk() (*string, bool)`

GetConfigContractOk returns a tuple with the ConfigContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContract

`func (o *IndexerResource) SetConfigContract(v string)`

SetConfigContract sets ConfigContract field to given value.

### HasConfigContract

`func (o *IndexerResource) HasConfigContract() bool`

HasConfigContract returns a boolean if a field has been set.

### SetConfigContractNil

`func (o *IndexerResource) SetConfigContractNil(b bool)`

 SetConfigContractNil sets the value for ConfigContract to be an explicit nil

### UnsetConfigContract
`func (o *IndexerResource) UnsetConfigContract()`

UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
### GetInfoLink

`func (o *IndexerResource) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *IndexerResource) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *IndexerResource) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *IndexerResource) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *IndexerResource) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *IndexerResource) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetMessage

`func (o *IndexerResource) GetMessage() ProviderMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IndexerResource) GetMessageOk() (*ProviderMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IndexerResource) SetMessage(v ProviderMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IndexerResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *IndexerResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IndexerResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IndexerResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IndexerResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *IndexerResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *IndexerResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPresets

`func (o *IndexerResource) GetPresets() []IndexerResource`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *IndexerResource) GetPresetsOk() (*[]IndexerResource, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *IndexerResource) SetPresets(v []IndexerResource)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *IndexerResource) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *IndexerResource) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *IndexerResource) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil
### GetIndexerUrls

`func (o *IndexerResource) GetIndexerUrls() []string`

GetIndexerUrls returns the IndexerUrls field if non-nil, zero value otherwise.

### GetIndexerUrlsOk

`func (o *IndexerResource) GetIndexerUrlsOk() (*[]string, bool)`

GetIndexerUrlsOk returns a tuple with the IndexerUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerUrls

`func (o *IndexerResource) SetIndexerUrls(v []string)`

SetIndexerUrls sets IndexerUrls field to given value.

### HasIndexerUrls

`func (o *IndexerResource) HasIndexerUrls() bool`

HasIndexerUrls returns a boolean if a field has been set.

### SetIndexerUrlsNil

`func (o *IndexerResource) SetIndexerUrlsNil(b bool)`

 SetIndexerUrlsNil sets the value for IndexerUrls to be an explicit nil

### UnsetIndexerUrls
`func (o *IndexerResource) UnsetIndexerUrls()`

UnsetIndexerUrls ensures that no value is present for IndexerUrls, not even an explicit nil
### GetLegacyUrls

`func (o *IndexerResource) GetLegacyUrls() []string`

GetLegacyUrls returns the LegacyUrls field if non-nil, zero value otherwise.

### GetLegacyUrlsOk

`func (o *IndexerResource) GetLegacyUrlsOk() (*[]string, bool)`

GetLegacyUrlsOk returns a tuple with the LegacyUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyUrls

`func (o *IndexerResource) SetLegacyUrls(v []string)`

SetLegacyUrls sets LegacyUrls field to given value.

### HasLegacyUrls

`func (o *IndexerResource) HasLegacyUrls() bool`

HasLegacyUrls returns a boolean if a field has been set.

### SetLegacyUrlsNil

`func (o *IndexerResource) SetLegacyUrlsNil(b bool)`

 SetLegacyUrlsNil sets the value for LegacyUrls to be an explicit nil

### UnsetLegacyUrls
`func (o *IndexerResource) UnsetLegacyUrls()`

UnsetLegacyUrls ensures that no value is present for LegacyUrls, not even an explicit nil
### GetDefinitionName

`func (o *IndexerResource) GetDefinitionName() string`

GetDefinitionName returns the DefinitionName field if non-nil, zero value otherwise.

### GetDefinitionNameOk

`func (o *IndexerResource) GetDefinitionNameOk() (*string, bool)`

GetDefinitionNameOk returns a tuple with the DefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionName

`func (o *IndexerResource) SetDefinitionName(v string)`

SetDefinitionName sets DefinitionName field to given value.

### HasDefinitionName

`func (o *IndexerResource) HasDefinitionName() bool`

HasDefinitionName returns a boolean if a field has been set.

### SetDefinitionNameNil

`func (o *IndexerResource) SetDefinitionNameNil(b bool)`

 SetDefinitionNameNil sets the value for DefinitionName to be an explicit nil

### UnsetDefinitionName
`func (o *IndexerResource) UnsetDefinitionName()`

UnsetDefinitionName ensures that no value is present for DefinitionName, not even an explicit nil
### GetDescription

`func (o *IndexerResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IndexerResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IndexerResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IndexerResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IndexerResource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IndexerResource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLanguage

`func (o *IndexerResource) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *IndexerResource) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *IndexerResource) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *IndexerResource) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *IndexerResource) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *IndexerResource) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetEncoding

`func (o *IndexerResource) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *IndexerResource) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *IndexerResource) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *IndexerResource) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### SetEncodingNil

`func (o *IndexerResource) SetEncodingNil(b bool)`

 SetEncodingNil sets the value for Encoding to be an explicit nil

### UnsetEncoding
`func (o *IndexerResource) UnsetEncoding()`

UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
### GetEnable

`func (o *IndexerResource) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IndexerResource) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IndexerResource) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IndexerResource) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetRedirect

`func (o *IndexerResource) GetRedirect() bool`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *IndexerResource) GetRedirectOk() (*bool, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *IndexerResource) SetRedirect(v bool)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *IndexerResource) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetSupportsRss

`func (o *IndexerResource) GetSupportsRss() bool`

GetSupportsRss returns the SupportsRss field if non-nil, zero value otherwise.

### GetSupportsRssOk

`func (o *IndexerResource) GetSupportsRssOk() (*bool, bool)`

GetSupportsRssOk returns a tuple with the SupportsRss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRss

`func (o *IndexerResource) SetSupportsRss(v bool)`

SetSupportsRss sets SupportsRss field to given value.

### HasSupportsRss

`func (o *IndexerResource) HasSupportsRss() bool`

HasSupportsRss returns a boolean if a field has been set.

### GetSupportsSearch

`func (o *IndexerResource) GetSupportsSearch() bool`

GetSupportsSearch returns the SupportsSearch field if non-nil, zero value otherwise.

### GetSupportsSearchOk

`func (o *IndexerResource) GetSupportsSearchOk() (*bool, bool)`

GetSupportsSearchOk returns a tuple with the SupportsSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSearch

`func (o *IndexerResource) SetSupportsSearch(v bool)`

SetSupportsSearch sets SupportsSearch field to given value.

### HasSupportsSearch

`func (o *IndexerResource) HasSupportsSearch() bool`

HasSupportsSearch returns a boolean if a field has been set.

### GetSupportsRedirect

`func (o *IndexerResource) GetSupportsRedirect() bool`

GetSupportsRedirect returns the SupportsRedirect field if non-nil, zero value otherwise.

### GetSupportsRedirectOk

`func (o *IndexerResource) GetSupportsRedirectOk() (*bool, bool)`

GetSupportsRedirectOk returns a tuple with the SupportsRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRedirect

`func (o *IndexerResource) SetSupportsRedirect(v bool)`

SetSupportsRedirect sets SupportsRedirect field to given value.

### HasSupportsRedirect

`func (o *IndexerResource) HasSupportsRedirect() bool`

HasSupportsRedirect returns a boolean if a field has been set.

### GetSupportsPagination

`func (o *IndexerResource) GetSupportsPagination() bool`

GetSupportsPagination returns the SupportsPagination field if non-nil, zero value otherwise.

### GetSupportsPaginationOk

`func (o *IndexerResource) GetSupportsPaginationOk() (*bool, bool)`

GetSupportsPaginationOk returns a tuple with the SupportsPagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPagination

`func (o *IndexerResource) SetSupportsPagination(v bool)`

SetSupportsPagination sets SupportsPagination field to given value.

### HasSupportsPagination

`func (o *IndexerResource) HasSupportsPagination() bool`

HasSupportsPagination returns a boolean if a field has been set.

### GetAppProfileId

`func (o *IndexerResource) GetAppProfileId() int32`

GetAppProfileId returns the AppProfileId field if non-nil, zero value otherwise.

### GetAppProfileIdOk

`func (o *IndexerResource) GetAppProfileIdOk() (*int32, bool)`

GetAppProfileIdOk returns a tuple with the AppProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileId

`func (o *IndexerResource) SetAppProfileId(v int32)`

SetAppProfileId sets AppProfileId field to given value.

### HasAppProfileId

`func (o *IndexerResource) HasAppProfileId() bool`

HasAppProfileId returns a boolean if a field has been set.

### GetProtocol

`func (o *IndexerResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IndexerResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IndexerResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IndexerResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPrivacy

`func (o *IndexerResource) GetPrivacy() IndexerPrivacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *IndexerResource) GetPrivacyOk() (*IndexerPrivacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *IndexerResource) SetPrivacy(v IndexerPrivacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *IndexerResource) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetCapabilities

`func (o *IndexerResource) GetCapabilities() IndexerCapabilityResource`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *IndexerResource) GetCapabilitiesOk() (*IndexerCapabilityResource, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *IndexerResource) SetCapabilities(v IndexerCapabilityResource)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *IndexerResource) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetPriority

`func (o *IndexerResource) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IndexerResource) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IndexerResource) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IndexerResource) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAdded

`func (o *IndexerResource) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *IndexerResource) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *IndexerResource) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *IndexerResource) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetStatus

`func (o *IndexerResource) GetStatus() IndexerStatusResource`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndexerResource) GetStatusOk() (*IndexerStatusResource, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndexerResource) SetStatus(v IndexerStatusResource)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IndexerResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSortName

`func (o *IndexerResource) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *IndexerResource) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *IndexerResource) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *IndexerResource) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### SetSortNameNil

`func (o *IndexerResource) SetSortNameNil(b bool)`

 SetSortNameNil sets the value for SortName to be an explicit nil

### UnsetSortName
`func (o *IndexerResource) UnsetSortName()`

UnsetSortName ensures that no value is present for SortName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


