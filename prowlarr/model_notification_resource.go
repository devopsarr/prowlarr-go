/*
Prowlarr

Prowlarr API docs

API version: v1.31.2.4975
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the NotificationResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationResource{}

// NotificationResource struct for NotificationResource
type NotificationResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ConfigContract NullableString `json:"configContract,omitempty"`
	InfoLink NullableString `json:"infoLink,omitempty"`
	Message *ProviderMessage `json:"message,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	Presets []NotificationResource `json:"presets,omitempty"`
	Link NullableString `json:"link,omitempty"`
	OnGrab *bool `json:"onGrab,omitempty"`
	OnHealthIssue *bool `json:"onHealthIssue,omitempty"`
	OnHealthRestored *bool `json:"onHealthRestored,omitempty"`
	OnApplicationUpdate *bool `json:"onApplicationUpdate,omitempty"`
	SupportsOnGrab *bool `json:"supportsOnGrab,omitempty"`
	IncludeManualGrabs *bool `json:"includeManualGrabs,omitempty"`
	SupportsOnHealthIssue *bool `json:"supportsOnHealthIssue,omitempty"`
	SupportsOnHealthRestored *bool `json:"supportsOnHealthRestored,omitempty"`
	IncludeHealthWarnings *bool `json:"includeHealthWarnings,omitempty"`
	SupportsOnApplicationUpdate *bool `json:"supportsOnApplicationUpdate,omitempty"`
	TestCommand NullableString `json:"testCommand,omitempty"`
}

// NewNotificationResource instantiates a new NotificationResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationResource() *NotificationResource {
	this := NotificationResource{}
	return &this
}

// NewNotificationResourceWithDefaults instantiates a new NotificationResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationResourceWithDefaults() *NotificationResource {
	this := NotificationResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NotificationResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NotificationResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NotificationResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NotificationResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NotificationResource) UnsetName() {
	o.Name.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetFields() []Field {
	if o == nil {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetFieldsOk() ([]Field, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *NotificationResource) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *NotificationResource) SetFields(v []Field) {
	o.Fields = v
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetImplementationName() string {
	if o == nil || IsNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetImplementationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *NotificationResource) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *NotificationResource) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *NotificationResource) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *NotificationResource) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetImplementation() string {
	if o == nil || IsNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetImplementationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *NotificationResource) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *NotificationResource) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *NotificationResource) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *NotificationResource) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetConfigContract returns the ConfigContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetConfigContract() string {
	if o == nil || IsNil(o.ConfigContract.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigContract.Get()
}

// GetConfigContractOk returns a tuple with the ConfigContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetConfigContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigContract.Get(), o.ConfigContract.IsSet()
}

// HasConfigContract returns a boolean if a field has been set.
func (o *NotificationResource) HasConfigContract() bool {
	if o != nil && o.ConfigContract.IsSet() {
		return true
	}

	return false
}

// SetConfigContract gets a reference to the given NullableString and assigns it to the ConfigContract field.
func (o *NotificationResource) SetConfigContract(v string) {
	o.ConfigContract.Set(&v)
}
// SetConfigContractNil sets the value for ConfigContract to be an explicit nil
func (o *NotificationResource) SetConfigContractNil() {
	o.ConfigContract.Set(nil)
}

// UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
func (o *NotificationResource) UnsetConfigContract() {
	o.ConfigContract.Unset()
}

// GetInfoLink returns the InfoLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetInfoLink() string {
	if o == nil || IsNil(o.InfoLink.Get()) {
		var ret string
		return ret
	}
	return *o.InfoLink.Get()
}

// GetInfoLinkOk returns a tuple with the InfoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetInfoLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfoLink.Get(), o.InfoLink.IsSet()
}

// HasInfoLink returns a boolean if a field has been set.
func (o *NotificationResource) HasInfoLink() bool {
	if o != nil && o.InfoLink.IsSet() {
		return true
	}

	return false
}

// SetInfoLink gets a reference to the given NullableString and assigns it to the InfoLink field.
func (o *NotificationResource) SetInfoLink(v string) {
	o.InfoLink.Set(&v)
}
// SetInfoLinkNil sets the value for InfoLink to be an explicit nil
func (o *NotificationResource) SetInfoLinkNil() {
	o.InfoLink.Set(nil)
}

// UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
func (o *NotificationResource) UnsetInfoLink() {
	o.InfoLink.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotificationResource) GetMessage() ProviderMessage {
	if o == nil || IsNil(o.Message) {
		var ret ProviderMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetMessageOk() (*ProviderMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotificationResource) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ProviderMessage and assigns it to the Message field.
func (o *NotificationResource) SetMessage(v ProviderMessage) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NotificationResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *NotificationResource) SetTags(v []int32) {
	o.Tags = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetPresets() []NotificationResource {
	if o == nil {
		var ret []NotificationResource
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetPresetsOk() ([]NotificationResource, bool) {
	if o == nil || IsNil(o.Presets) {
		return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *NotificationResource) HasPresets() bool {
	if o != nil && !IsNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []NotificationResource and assigns it to the Presets field.
func (o *NotificationResource) SetPresets(v []NotificationResource) {
	o.Presets = v
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *NotificationResource) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *NotificationResource) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *NotificationResource) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *NotificationResource) UnsetLink() {
	o.Link.Unset()
}

// GetOnGrab returns the OnGrab field value if set, zero value otherwise.
func (o *NotificationResource) GetOnGrab() bool {
	if o == nil || IsNil(o.OnGrab) {
		var ret bool
		return ret
	}
	return *o.OnGrab
}

// GetOnGrabOk returns a tuple with the OnGrab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetOnGrabOk() (*bool, bool) {
	if o == nil || IsNil(o.OnGrab) {
		return nil, false
	}
	return o.OnGrab, true
}

// HasOnGrab returns a boolean if a field has been set.
func (o *NotificationResource) HasOnGrab() bool {
	if o != nil && !IsNil(o.OnGrab) {
		return true
	}

	return false
}

// SetOnGrab gets a reference to the given bool and assigns it to the OnGrab field.
func (o *NotificationResource) SetOnGrab(v bool) {
	o.OnGrab = &v
}

// GetOnHealthIssue returns the OnHealthIssue field value if set, zero value otherwise.
func (o *NotificationResource) GetOnHealthIssue() bool {
	if o == nil || IsNil(o.OnHealthIssue) {
		var ret bool
		return ret
	}
	return *o.OnHealthIssue
}

// GetOnHealthIssueOk returns a tuple with the OnHealthIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetOnHealthIssueOk() (*bool, bool) {
	if o == nil || IsNil(o.OnHealthIssue) {
		return nil, false
	}
	return o.OnHealthIssue, true
}

// HasOnHealthIssue returns a boolean if a field has been set.
func (o *NotificationResource) HasOnHealthIssue() bool {
	if o != nil && !IsNil(o.OnHealthIssue) {
		return true
	}

	return false
}

// SetOnHealthIssue gets a reference to the given bool and assigns it to the OnHealthIssue field.
func (o *NotificationResource) SetOnHealthIssue(v bool) {
	o.OnHealthIssue = &v
}

// GetOnHealthRestored returns the OnHealthRestored field value if set, zero value otherwise.
func (o *NotificationResource) GetOnHealthRestored() bool {
	if o == nil || IsNil(o.OnHealthRestored) {
		var ret bool
		return ret
	}
	return *o.OnHealthRestored
}

// GetOnHealthRestoredOk returns a tuple with the OnHealthRestored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetOnHealthRestoredOk() (*bool, bool) {
	if o == nil || IsNil(o.OnHealthRestored) {
		return nil, false
	}
	return o.OnHealthRestored, true
}

// HasOnHealthRestored returns a boolean if a field has been set.
func (o *NotificationResource) HasOnHealthRestored() bool {
	if o != nil && !IsNil(o.OnHealthRestored) {
		return true
	}

	return false
}

// SetOnHealthRestored gets a reference to the given bool and assigns it to the OnHealthRestored field.
func (o *NotificationResource) SetOnHealthRestored(v bool) {
	o.OnHealthRestored = &v
}

// GetOnApplicationUpdate returns the OnApplicationUpdate field value if set, zero value otherwise.
func (o *NotificationResource) GetOnApplicationUpdate() bool {
	if o == nil || IsNil(o.OnApplicationUpdate) {
		var ret bool
		return ret
	}
	return *o.OnApplicationUpdate
}

// GetOnApplicationUpdateOk returns a tuple with the OnApplicationUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetOnApplicationUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.OnApplicationUpdate) {
		return nil, false
	}
	return o.OnApplicationUpdate, true
}

// HasOnApplicationUpdate returns a boolean if a field has been set.
func (o *NotificationResource) HasOnApplicationUpdate() bool {
	if o != nil && !IsNil(o.OnApplicationUpdate) {
		return true
	}

	return false
}

// SetOnApplicationUpdate gets a reference to the given bool and assigns it to the OnApplicationUpdate field.
func (o *NotificationResource) SetOnApplicationUpdate(v bool) {
	o.OnApplicationUpdate = &v
}

// GetSupportsOnGrab returns the SupportsOnGrab field value if set, zero value otherwise.
func (o *NotificationResource) GetSupportsOnGrab() bool {
	if o == nil || IsNil(o.SupportsOnGrab) {
		var ret bool
		return ret
	}
	return *o.SupportsOnGrab
}

// GetSupportsOnGrabOk returns a tuple with the SupportsOnGrab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetSupportsOnGrabOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsOnGrab) {
		return nil, false
	}
	return o.SupportsOnGrab, true
}

// HasSupportsOnGrab returns a boolean if a field has been set.
func (o *NotificationResource) HasSupportsOnGrab() bool {
	if o != nil && !IsNil(o.SupportsOnGrab) {
		return true
	}

	return false
}

// SetSupportsOnGrab gets a reference to the given bool and assigns it to the SupportsOnGrab field.
func (o *NotificationResource) SetSupportsOnGrab(v bool) {
	o.SupportsOnGrab = &v
}

// GetIncludeManualGrabs returns the IncludeManualGrabs field value if set, zero value otherwise.
func (o *NotificationResource) GetIncludeManualGrabs() bool {
	if o == nil || IsNil(o.IncludeManualGrabs) {
		var ret bool
		return ret
	}
	return *o.IncludeManualGrabs
}

// GetIncludeManualGrabsOk returns a tuple with the IncludeManualGrabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetIncludeManualGrabsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeManualGrabs) {
		return nil, false
	}
	return o.IncludeManualGrabs, true
}

// HasIncludeManualGrabs returns a boolean if a field has been set.
func (o *NotificationResource) HasIncludeManualGrabs() bool {
	if o != nil && !IsNil(o.IncludeManualGrabs) {
		return true
	}

	return false
}

// SetIncludeManualGrabs gets a reference to the given bool and assigns it to the IncludeManualGrabs field.
func (o *NotificationResource) SetIncludeManualGrabs(v bool) {
	o.IncludeManualGrabs = &v
}

// GetSupportsOnHealthIssue returns the SupportsOnHealthIssue field value if set, zero value otherwise.
func (o *NotificationResource) GetSupportsOnHealthIssue() bool {
	if o == nil || IsNil(o.SupportsOnHealthIssue) {
		var ret bool
		return ret
	}
	return *o.SupportsOnHealthIssue
}

// GetSupportsOnHealthIssueOk returns a tuple with the SupportsOnHealthIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetSupportsOnHealthIssueOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsOnHealthIssue) {
		return nil, false
	}
	return o.SupportsOnHealthIssue, true
}

// HasSupportsOnHealthIssue returns a boolean if a field has been set.
func (o *NotificationResource) HasSupportsOnHealthIssue() bool {
	if o != nil && !IsNil(o.SupportsOnHealthIssue) {
		return true
	}

	return false
}

// SetSupportsOnHealthIssue gets a reference to the given bool and assigns it to the SupportsOnHealthIssue field.
func (o *NotificationResource) SetSupportsOnHealthIssue(v bool) {
	o.SupportsOnHealthIssue = &v
}

// GetSupportsOnHealthRestored returns the SupportsOnHealthRestored field value if set, zero value otherwise.
func (o *NotificationResource) GetSupportsOnHealthRestored() bool {
	if o == nil || IsNil(o.SupportsOnHealthRestored) {
		var ret bool
		return ret
	}
	return *o.SupportsOnHealthRestored
}

// GetSupportsOnHealthRestoredOk returns a tuple with the SupportsOnHealthRestored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetSupportsOnHealthRestoredOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsOnHealthRestored) {
		return nil, false
	}
	return o.SupportsOnHealthRestored, true
}

// HasSupportsOnHealthRestored returns a boolean if a field has been set.
func (o *NotificationResource) HasSupportsOnHealthRestored() bool {
	if o != nil && !IsNil(o.SupportsOnHealthRestored) {
		return true
	}

	return false
}

// SetSupportsOnHealthRestored gets a reference to the given bool and assigns it to the SupportsOnHealthRestored field.
func (o *NotificationResource) SetSupportsOnHealthRestored(v bool) {
	o.SupportsOnHealthRestored = &v
}

// GetIncludeHealthWarnings returns the IncludeHealthWarnings field value if set, zero value otherwise.
func (o *NotificationResource) GetIncludeHealthWarnings() bool {
	if o == nil || IsNil(o.IncludeHealthWarnings) {
		var ret bool
		return ret
	}
	return *o.IncludeHealthWarnings
}

// GetIncludeHealthWarningsOk returns a tuple with the IncludeHealthWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetIncludeHealthWarningsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHealthWarnings) {
		return nil, false
	}
	return o.IncludeHealthWarnings, true
}

// HasIncludeHealthWarnings returns a boolean if a field has been set.
func (o *NotificationResource) HasIncludeHealthWarnings() bool {
	if o != nil && !IsNil(o.IncludeHealthWarnings) {
		return true
	}

	return false
}

// SetIncludeHealthWarnings gets a reference to the given bool and assigns it to the IncludeHealthWarnings field.
func (o *NotificationResource) SetIncludeHealthWarnings(v bool) {
	o.IncludeHealthWarnings = &v
}

// GetSupportsOnApplicationUpdate returns the SupportsOnApplicationUpdate field value if set, zero value otherwise.
func (o *NotificationResource) GetSupportsOnApplicationUpdate() bool {
	if o == nil || IsNil(o.SupportsOnApplicationUpdate) {
		var ret bool
		return ret
	}
	return *o.SupportsOnApplicationUpdate
}

// GetSupportsOnApplicationUpdateOk returns a tuple with the SupportsOnApplicationUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationResource) GetSupportsOnApplicationUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsOnApplicationUpdate) {
		return nil, false
	}
	return o.SupportsOnApplicationUpdate, true
}

// HasSupportsOnApplicationUpdate returns a boolean if a field has been set.
func (o *NotificationResource) HasSupportsOnApplicationUpdate() bool {
	if o != nil && !IsNil(o.SupportsOnApplicationUpdate) {
		return true
	}

	return false
}

// SetSupportsOnApplicationUpdate gets a reference to the given bool and assigns it to the SupportsOnApplicationUpdate field.
func (o *NotificationResource) SetSupportsOnApplicationUpdate(v bool) {
	o.SupportsOnApplicationUpdate = &v
}

// GetTestCommand returns the TestCommand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationResource) GetTestCommand() string {
	if o == nil || IsNil(o.TestCommand.Get()) {
		var ret string
		return ret
	}
	return *o.TestCommand.Get()
}

// GetTestCommandOk returns a tuple with the TestCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationResource) GetTestCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestCommand.Get(), o.TestCommand.IsSet()
}

// HasTestCommand returns a boolean if a field has been set.
func (o *NotificationResource) HasTestCommand() bool {
	if o != nil && o.TestCommand.IsSet() {
		return true
	}

	return false
}

// SetTestCommand gets a reference to the given NullableString and assigns it to the TestCommand field.
func (o *NotificationResource) SetTestCommand(v string) {
	o.TestCommand.Set(&v)
}
// SetTestCommandNil sets the value for TestCommand to be an explicit nil
func (o *NotificationResource) SetTestCommandNil() {
	o.TestCommand.Set(nil)
}

// UnsetTestCommand ensures that no value is present for TestCommand, not even an explicit nil
func (o *NotificationResource) UnsetTestCommand() {
	o.TestCommand.Unset()
}

func (o NotificationResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationResource) ToMap() (map[string]interface{}, error) {
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
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if !IsNil(o.OnGrab) {
		toSerialize["onGrab"] = o.OnGrab
	}
	if !IsNil(o.OnHealthIssue) {
		toSerialize["onHealthIssue"] = o.OnHealthIssue
	}
	if !IsNil(o.OnHealthRestored) {
		toSerialize["onHealthRestored"] = o.OnHealthRestored
	}
	if !IsNil(o.OnApplicationUpdate) {
		toSerialize["onApplicationUpdate"] = o.OnApplicationUpdate
	}
	if !IsNil(o.SupportsOnGrab) {
		toSerialize["supportsOnGrab"] = o.SupportsOnGrab
	}
	if !IsNil(o.IncludeManualGrabs) {
		toSerialize["includeManualGrabs"] = o.IncludeManualGrabs
	}
	if !IsNil(o.SupportsOnHealthIssue) {
		toSerialize["supportsOnHealthIssue"] = o.SupportsOnHealthIssue
	}
	if !IsNil(o.SupportsOnHealthRestored) {
		toSerialize["supportsOnHealthRestored"] = o.SupportsOnHealthRestored
	}
	if !IsNil(o.IncludeHealthWarnings) {
		toSerialize["includeHealthWarnings"] = o.IncludeHealthWarnings
	}
	if !IsNil(o.SupportsOnApplicationUpdate) {
		toSerialize["supportsOnApplicationUpdate"] = o.SupportsOnApplicationUpdate
	}
	if o.TestCommand.IsSet() {
		toSerialize["testCommand"] = o.TestCommand.Get()
	}
	return toSerialize, nil
}

type NullableNotificationResource struct {
	value *NotificationResource
	isSet bool
}

func (v NullableNotificationResource) Get() *NotificationResource {
	return v.value
}

func (v *NullableNotificationResource) Set(val *NotificationResource) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationResource) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationResource(val *NotificationResource) *NullableNotificationResource {
	return &NullableNotificationResource{value: val, isSet: true}
}

func (v NullableNotificationResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


