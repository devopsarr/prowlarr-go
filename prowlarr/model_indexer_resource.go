/*
Prowlarr

Prowlarr API docs

API version: v1.28.2.4885
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"time"
)

// checks if the IndexerResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexerResource{}

// IndexerResource struct for IndexerResource
type IndexerResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ConfigContract NullableString `json:"configContract,omitempty"`
	InfoLink NullableString `json:"infoLink,omitempty"`
	Message *ProviderMessage `json:"message,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	Presets []IndexerResource `json:"presets,omitempty"`
	IndexerUrls []string `json:"indexerUrls,omitempty"`
	LegacyUrls []string `json:"legacyUrls,omitempty"`
	DefinitionName NullableString `json:"definitionName,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Language NullableString `json:"language,omitempty"`
	Encoding NullableString `json:"encoding,omitempty"`
	Enable *bool `json:"enable,omitempty"`
	Redirect *bool `json:"redirect,omitempty"`
	SupportsRss *bool `json:"supportsRss,omitempty"`
	SupportsSearch *bool `json:"supportsSearch,omitempty"`
	SupportsRedirect *bool `json:"supportsRedirect,omitempty"`
	SupportsPagination *bool `json:"supportsPagination,omitempty"`
	AppProfileId *int32 `json:"appProfileId,omitempty"`
	Protocol *DownloadProtocol `json:"protocol,omitempty"`
	Privacy *IndexerPrivacy `json:"privacy,omitempty"`
	Capabilities *IndexerCapabilityResource `json:"capabilities,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	DownloadClientId *int32 `json:"downloadClientId,omitempty"`
	Added *time.Time `json:"added,omitempty"`
	Status *IndexerStatusResource `json:"status,omitempty"`
	SortName NullableString `json:"sortName,omitempty"`
}

// NewIndexerResource instantiates a new IndexerResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerResource() *IndexerResource {
	this := IndexerResource{}
	return &this
}

// NewIndexerResourceWithDefaults instantiates a new IndexerResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerResourceWithDefaults() *IndexerResource {
	this := IndexerResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndexerResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndexerResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IndexerResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IndexerResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IndexerResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IndexerResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IndexerResource) UnsetName() {
	o.Name.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetFields() []Field {
	if o == nil {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetFieldsOk() ([]Field, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IndexerResource) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *IndexerResource) SetFields(v []Field) {
	o.Fields = v
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetImplementationName() string {
	if o == nil || IsNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetImplementationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *IndexerResource) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *IndexerResource) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *IndexerResource) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *IndexerResource) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetImplementation() string {
	if o == nil || IsNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetImplementationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *IndexerResource) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *IndexerResource) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *IndexerResource) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *IndexerResource) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetConfigContract returns the ConfigContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetConfigContract() string {
	if o == nil || IsNil(o.ConfigContract.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigContract.Get()
}

// GetConfigContractOk returns a tuple with the ConfigContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetConfigContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigContract.Get(), o.ConfigContract.IsSet()
}

// HasConfigContract returns a boolean if a field has been set.
func (o *IndexerResource) HasConfigContract() bool {
	if o != nil && o.ConfigContract.IsSet() {
		return true
	}

	return false
}

// SetConfigContract gets a reference to the given NullableString and assigns it to the ConfigContract field.
func (o *IndexerResource) SetConfigContract(v string) {
	o.ConfigContract.Set(&v)
}
// SetConfigContractNil sets the value for ConfigContract to be an explicit nil
func (o *IndexerResource) SetConfigContractNil() {
	o.ConfigContract.Set(nil)
}

// UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
func (o *IndexerResource) UnsetConfigContract() {
	o.ConfigContract.Unset()
}

// GetInfoLink returns the InfoLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetInfoLink() string {
	if o == nil || IsNil(o.InfoLink.Get()) {
		var ret string
		return ret
	}
	return *o.InfoLink.Get()
}

// GetInfoLinkOk returns a tuple with the InfoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetInfoLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfoLink.Get(), o.InfoLink.IsSet()
}

// HasInfoLink returns a boolean if a field has been set.
func (o *IndexerResource) HasInfoLink() bool {
	if o != nil && o.InfoLink.IsSet() {
		return true
	}

	return false
}

// SetInfoLink gets a reference to the given NullableString and assigns it to the InfoLink field.
func (o *IndexerResource) SetInfoLink(v string) {
	o.InfoLink.Set(&v)
}
// SetInfoLinkNil sets the value for InfoLink to be an explicit nil
func (o *IndexerResource) SetInfoLinkNil() {
	o.InfoLink.Set(nil)
}

// UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
func (o *IndexerResource) UnsetInfoLink() {
	o.InfoLink.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IndexerResource) GetMessage() ProviderMessage {
	if o == nil || IsNil(o.Message) {
		var ret ProviderMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetMessageOk() (*ProviderMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IndexerResource) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ProviderMessage and assigns it to the Message field.
func (o *IndexerResource) SetMessage(v ProviderMessage) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IndexerResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *IndexerResource) SetTags(v []int32) {
	o.Tags = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetPresets() []IndexerResource {
	if o == nil {
		var ret []IndexerResource
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetPresetsOk() ([]IndexerResource, bool) {
	if o == nil || IsNil(o.Presets) {
		return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *IndexerResource) HasPresets() bool {
	if o != nil && !IsNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []IndexerResource and assigns it to the Presets field.
func (o *IndexerResource) SetPresets(v []IndexerResource) {
	o.Presets = v
}

// GetIndexerUrls returns the IndexerUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetIndexerUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IndexerUrls
}

// GetIndexerUrlsOk returns a tuple with the IndexerUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetIndexerUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.IndexerUrls) {
		return nil, false
	}
	return o.IndexerUrls, true
}

// HasIndexerUrls returns a boolean if a field has been set.
func (o *IndexerResource) HasIndexerUrls() bool {
	if o != nil && !IsNil(o.IndexerUrls) {
		return true
	}

	return false
}

// SetIndexerUrls gets a reference to the given []string and assigns it to the IndexerUrls field.
func (o *IndexerResource) SetIndexerUrls(v []string) {
	o.IndexerUrls = v
}

// GetLegacyUrls returns the LegacyUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetLegacyUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LegacyUrls
}

// GetLegacyUrlsOk returns a tuple with the LegacyUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetLegacyUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.LegacyUrls) {
		return nil, false
	}
	return o.LegacyUrls, true
}

// HasLegacyUrls returns a boolean if a field has been set.
func (o *IndexerResource) HasLegacyUrls() bool {
	if o != nil && !IsNil(o.LegacyUrls) {
		return true
	}

	return false
}

// SetLegacyUrls gets a reference to the given []string and assigns it to the LegacyUrls field.
func (o *IndexerResource) SetLegacyUrls(v []string) {
	o.LegacyUrls = v
}

// GetDefinitionName returns the DefinitionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetDefinitionName() string {
	if o == nil || IsNil(o.DefinitionName.Get()) {
		var ret string
		return ret
	}
	return *o.DefinitionName.Get()
}

// GetDefinitionNameOk returns a tuple with the DefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetDefinitionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefinitionName.Get(), o.DefinitionName.IsSet()
}

// HasDefinitionName returns a boolean if a field has been set.
func (o *IndexerResource) HasDefinitionName() bool {
	if o != nil && o.DefinitionName.IsSet() {
		return true
	}

	return false
}

// SetDefinitionName gets a reference to the given NullableString and assigns it to the DefinitionName field.
func (o *IndexerResource) SetDefinitionName(v string) {
	o.DefinitionName.Set(&v)
}
// SetDefinitionNameNil sets the value for DefinitionName to be an explicit nil
func (o *IndexerResource) SetDefinitionNameNil() {
	o.DefinitionName.Set(nil)
}

// UnsetDefinitionName ensures that no value is present for DefinitionName, not even an explicit nil
func (o *IndexerResource) UnsetDefinitionName() {
	o.DefinitionName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IndexerResource) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IndexerResource) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IndexerResource) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IndexerResource) UnsetDescription() {
	o.Description.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetLanguage() string {
	if o == nil || IsNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *IndexerResource) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *IndexerResource) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *IndexerResource) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *IndexerResource) UnsetLanguage() {
	o.Language.Unset()
}

// GetEncoding returns the Encoding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetEncoding() string {
	if o == nil || IsNil(o.Encoding.Get()) {
		var ret string
		return ret
	}
	return *o.Encoding.Get()
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Encoding.Get(), o.Encoding.IsSet()
}

// HasEncoding returns a boolean if a field has been set.
func (o *IndexerResource) HasEncoding() bool {
	if o != nil && o.Encoding.IsSet() {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given NullableString and assigns it to the Encoding field.
func (o *IndexerResource) SetEncoding(v string) {
	o.Encoding.Set(&v)
}
// SetEncodingNil sets the value for Encoding to be an explicit nil
func (o *IndexerResource) SetEncodingNil() {
	o.Encoding.Set(nil)
}

// UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
func (o *IndexerResource) UnsetEncoding() {
	o.Encoding.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *IndexerResource) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *IndexerResource) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *IndexerResource) SetEnable(v bool) {
	o.Enable = &v
}

// GetRedirect returns the Redirect field value if set, zero value otherwise.
func (o *IndexerResource) GetRedirect() bool {
	if o == nil || IsNil(o.Redirect) {
		var ret bool
		return ret
	}
	return *o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetRedirectOk() (*bool, bool) {
	if o == nil || IsNil(o.Redirect) {
		return nil, false
	}
	return o.Redirect, true
}

// HasRedirect returns a boolean if a field has been set.
func (o *IndexerResource) HasRedirect() bool {
	if o != nil && !IsNil(o.Redirect) {
		return true
	}

	return false
}

// SetRedirect gets a reference to the given bool and assigns it to the Redirect field.
func (o *IndexerResource) SetRedirect(v bool) {
	o.Redirect = &v
}

// GetSupportsRss returns the SupportsRss field value if set, zero value otherwise.
func (o *IndexerResource) GetSupportsRss() bool {
	if o == nil || IsNil(o.SupportsRss) {
		var ret bool
		return ret
	}
	return *o.SupportsRss
}

// GetSupportsRssOk returns a tuple with the SupportsRss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetSupportsRssOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsRss) {
		return nil, false
	}
	return o.SupportsRss, true
}

// HasSupportsRss returns a boolean if a field has been set.
func (o *IndexerResource) HasSupportsRss() bool {
	if o != nil && !IsNil(o.SupportsRss) {
		return true
	}

	return false
}

// SetSupportsRss gets a reference to the given bool and assigns it to the SupportsRss field.
func (o *IndexerResource) SetSupportsRss(v bool) {
	o.SupportsRss = &v
}

// GetSupportsSearch returns the SupportsSearch field value if set, zero value otherwise.
func (o *IndexerResource) GetSupportsSearch() bool {
	if o == nil || IsNil(o.SupportsSearch) {
		var ret bool
		return ret
	}
	return *o.SupportsSearch
}

// GetSupportsSearchOk returns a tuple with the SupportsSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetSupportsSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsSearch) {
		return nil, false
	}
	return o.SupportsSearch, true
}

// HasSupportsSearch returns a boolean if a field has been set.
func (o *IndexerResource) HasSupportsSearch() bool {
	if o != nil && !IsNil(o.SupportsSearch) {
		return true
	}

	return false
}

// SetSupportsSearch gets a reference to the given bool and assigns it to the SupportsSearch field.
func (o *IndexerResource) SetSupportsSearch(v bool) {
	o.SupportsSearch = &v
}

// GetSupportsRedirect returns the SupportsRedirect field value if set, zero value otherwise.
func (o *IndexerResource) GetSupportsRedirect() bool {
	if o == nil || IsNil(o.SupportsRedirect) {
		var ret bool
		return ret
	}
	return *o.SupportsRedirect
}

// GetSupportsRedirectOk returns a tuple with the SupportsRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetSupportsRedirectOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsRedirect) {
		return nil, false
	}
	return o.SupportsRedirect, true
}

// HasSupportsRedirect returns a boolean if a field has been set.
func (o *IndexerResource) HasSupportsRedirect() bool {
	if o != nil && !IsNil(o.SupportsRedirect) {
		return true
	}

	return false
}

// SetSupportsRedirect gets a reference to the given bool and assigns it to the SupportsRedirect field.
func (o *IndexerResource) SetSupportsRedirect(v bool) {
	o.SupportsRedirect = &v
}

// GetSupportsPagination returns the SupportsPagination field value if set, zero value otherwise.
func (o *IndexerResource) GetSupportsPagination() bool {
	if o == nil || IsNil(o.SupportsPagination) {
		var ret bool
		return ret
	}
	return *o.SupportsPagination
}

// GetSupportsPaginationOk returns a tuple with the SupportsPagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetSupportsPaginationOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsPagination) {
		return nil, false
	}
	return o.SupportsPagination, true
}

// HasSupportsPagination returns a boolean if a field has been set.
func (o *IndexerResource) HasSupportsPagination() bool {
	if o != nil && !IsNil(o.SupportsPagination) {
		return true
	}

	return false
}

// SetSupportsPagination gets a reference to the given bool and assigns it to the SupportsPagination field.
func (o *IndexerResource) SetSupportsPagination(v bool) {
	o.SupportsPagination = &v
}

// GetAppProfileId returns the AppProfileId field value if set, zero value otherwise.
func (o *IndexerResource) GetAppProfileId() int32 {
	if o == nil || IsNil(o.AppProfileId) {
		var ret int32
		return ret
	}
	return *o.AppProfileId
}

// GetAppProfileIdOk returns a tuple with the AppProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetAppProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AppProfileId) {
		return nil, false
	}
	return o.AppProfileId, true
}

// HasAppProfileId returns a boolean if a field has been set.
func (o *IndexerResource) HasAppProfileId() bool {
	if o != nil && !IsNil(o.AppProfileId) {
		return true
	}

	return false
}

// SetAppProfileId gets a reference to the given int32 and assigns it to the AppProfileId field.
func (o *IndexerResource) SetAppProfileId(v int32) {
	o.AppProfileId = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *IndexerResource) GetProtocol() DownloadProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret DownloadProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetProtocolOk() (*DownloadProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *IndexerResource) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given DownloadProtocol and assigns it to the Protocol field.
func (o *IndexerResource) SetProtocol(v DownloadProtocol) {
	o.Protocol = &v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *IndexerResource) GetPrivacy() IndexerPrivacy {
	if o == nil || IsNil(o.Privacy) {
		var ret IndexerPrivacy
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetPrivacyOk() (*IndexerPrivacy, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *IndexerResource) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given IndexerPrivacy and assigns it to the Privacy field.
func (o *IndexerResource) SetPrivacy(v IndexerPrivacy) {
	o.Privacy = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *IndexerResource) GetCapabilities() IndexerCapabilityResource {
	if o == nil || IsNil(o.Capabilities) {
		var ret IndexerCapabilityResource
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetCapabilitiesOk() (*IndexerCapabilityResource, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *IndexerResource) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given IndexerCapabilityResource and assigns it to the Capabilities field.
func (o *IndexerResource) SetCapabilities(v IndexerCapabilityResource) {
	o.Capabilities = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IndexerResource) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IndexerResource) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *IndexerResource) SetPriority(v int32) {
	o.Priority = &v
}

// GetDownloadClientId returns the DownloadClientId field value if set, zero value otherwise.
func (o *IndexerResource) GetDownloadClientId() int32 {
	if o == nil || IsNil(o.DownloadClientId) {
		var ret int32
		return ret
	}
	return *o.DownloadClientId
}

// GetDownloadClientIdOk returns a tuple with the DownloadClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetDownloadClientIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DownloadClientId) {
		return nil, false
	}
	return o.DownloadClientId, true
}

// HasDownloadClientId returns a boolean if a field has been set.
func (o *IndexerResource) HasDownloadClientId() bool {
	if o != nil && !IsNil(o.DownloadClientId) {
		return true
	}

	return false
}

// SetDownloadClientId gets a reference to the given int32 and assigns it to the DownloadClientId field.
func (o *IndexerResource) SetDownloadClientId(v int32) {
	o.DownloadClientId = &v
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *IndexerResource) GetAdded() time.Time {
	if o == nil || IsNil(o.Added) {
		var ret time.Time
		return ret
	}
	return *o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetAddedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Added) {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *IndexerResource) HasAdded() bool {
	if o != nil && !IsNil(o.Added) {
		return true
	}

	return false
}

// SetAdded gets a reference to the given time.Time and assigns it to the Added field.
func (o *IndexerResource) SetAdded(v time.Time) {
	o.Added = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IndexerResource) GetStatus() IndexerStatusResource {
	if o == nil || IsNil(o.Status) {
		var ret IndexerStatusResource
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerResource) GetStatusOk() (*IndexerStatusResource, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IndexerResource) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IndexerStatusResource and assigns it to the Status field.
func (o *IndexerResource) SetStatus(v IndexerStatusResource) {
	o.Status = &v
}

// GetSortName returns the SortName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerResource) GetSortName() string {
	if o == nil || IsNil(o.SortName.Get()) {
		var ret string
		return ret
	}
	return *o.SortName.Get()
}

// GetSortNameOk returns a tuple with the SortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerResource) GetSortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortName.Get(), o.SortName.IsSet()
}

// HasSortName returns a boolean if a field has been set.
func (o *IndexerResource) HasSortName() bool {
	if o != nil && o.SortName.IsSet() {
		return true
	}

	return false
}

// SetSortName gets a reference to the given NullableString and assigns it to the SortName field.
func (o *IndexerResource) SetSortName(v string) {
	o.SortName.Set(&v)
}
// SetSortNameNil sets the value for SortName to be an explicit nil
func (o *IndexerResource) SetSortNameNil() {
	o.SortName.Set(nil)
}

// UnsetSortName ensures that no value is present for SortName, not even an explicit nil
func (o *IndexerResource) UnsetSortName() {
	o.SortName.Unset()
}

func (o IndexerResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexerResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.ImplementationName.IsSet() {
		toSerialize["implementationName"] = o.ImplementationName.Get()
	}
	if o.Implementation.IsSet() {
		toSerialize["implementation"] = o.Implementation.Get()
	}
	if o.ConfigContract.IsSet() {
		toSerialize["configContract"] = o.ConfigContract.Get()
	}
	if o.InfoLink.IsSet() {
		toSerialize["infoLink"] = o.InfoLink.Get()
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	if o.IndexerUrls != nil {
		toSerialize["indexerUrls"] = o.IndexerUrls
	}
	if o.LegacyUrls != nil {
		toSerialize["legacyUrls"] = o.LegacyUrls
	}
	if o.DefinitionName.IsSet() {
		toSerialize["definitionName"] = o.DefinitionName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.Encoding.IsSet() {
		toSerialize["encoding"] = o.Encoding.Get()
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Redirect) {
		toSerialize["redirect"] = o.Redirect
	}
	if !IsNil(o.SupportsRss) {
		toSerialize["supportsRss"] = o.SupportsRss
	}
	if !IsNil(o.SupportsSearch) {
		toSerialize["supportsSearch"] = o.SupportsSearch
	}
	if !IsNil(o.SupportsRedirect) {
		toSerialize["supportsRedirect"] = o.SupportsRedirect
	}
	if !IsNil(o.SupportsPagination) {
		toSerialize["supportsPagination"] = o.SupportsPagination
	}
	if !IsNil(o.AppProfileId) {
		toSerialize["appProfileId"] = o.AppProfileId
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.DownloadClientId) {
		toSerialize["downloadClientId"] = o.DownloadClientId
	}
	if !IsNil(o.Added) {
		toSerialize["added"] = o.Added
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.SortName.IsSet() {
		toSerialize["sortName"] = o.SortName.Get()
	}
	return toSerialize, nil
}

type NullableIndexerResource struct {
	value *IndexerResource
	isSet bool
}

func (v NullableIndexerResource) Get() *IndexerResource {
	return v.value
}

func (v *NullableIndexerResource) Set(val *IndexerResource) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerResource) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerResource(val *IndexerResource) *NullableIndexerResource {
	return &NullableIndexerResource{value: val, isSet: true}
}

func (v NullableIndexerResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


