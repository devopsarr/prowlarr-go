/*
Prowlarr

Prowlarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// ApplicationResource struct for ApplicationResource
type ApplicationResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Fields []*Field `json:"fields,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ConfigContract NullableString `json:"configContract,omitempty"`
	InfoLink NullableString `json:"infoLink,omitempty"`
	Message *ProviderMessage `json:"message,omitempty"`
	Tags []*int32 `json:"tags,omitempty"`
	Presets []*ApplicationResource `json:"presets,omitempty"`
	SyncLevel *ApplicationSyncLevel `json:"syncLevel,omitempty"`
	TestCommand NullableString `json:"testCommand,omitempty"`
}

// NewApplicationResource instantiates a new ApplicationResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResource() *ApplicationResource {
	this := ApplicationResource{}
	return &this
}

// NewApplicationResourceWithDefaults instantiates a new ApplicationResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResourceWithDefaults() *ApplicationResource {
	this := ApplicationResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApplicationResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApplicationResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ApplicationResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApplicationResource) UnsetName() {
	o.Name.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetFields() []*Field {
	if o == nil {
		var ret []*Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetFieldsOk() ([]*Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ApplicationResource) HasFields() bool {
	if o != nil && isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *ApplicationResource) SetFields(v []*Field) {
	o.Fields = v
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetImplementationName() string {
	if o == nil || isNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetImplementationNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *ApplicationResource) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *ApplicationResource) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *ApplicationResource) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *ApplicationResource) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetImplementation() string {
	if o == nil || isNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetImplementationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *ApplicationResource) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *ApplicationResource) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *ApplicationResource) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *ApplicationResource) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetConfigContract returns the ConfigContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetConfigContract() string {
	if o == nil || isNil(o.ConfigContract.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigContract.Get()
}

// GetConfigContractOk returns a tuple with the ConfigContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetConfigContractOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ConfigContract.Get(), o.ConfigContract.IsSet()
}

// HasConfigContract returns a boolean if a field has been set.
func (o *ApplicationResource) HasConfigContract() bool {
	if o != nil && o.ConfigContract.IsSet() {
		return true
	}

	return false
}

// SetConfigContract gets a reference to the given NullableString and assigns it to the ConfigContract field.
func (o *ApplicationResource) SetConfigContract(v string) {
	o.ConfigContract.Set(&v)
}
// SetConfigContractNil sets the value for ConfigContract to be an explicit nil
func (o *ApplicationResource) SetConfigContractNil() {
	o.ConfigContract.Set(nil)
}

// UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
func (o *ApplicationResource) UnsetConfigContract() {
	o.ConfigContract.Unset()
}

// GetInfoLink returns the InfoLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetInfoLink() string {
	if o == nil || isNil(o.InfoLink.Get()) {
		var ret string
		return ret
	}
	return *o.InfoLink.Get()
}

// GetInfoLinkOk returns a tuple with the InfoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetInfoLinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.InfoLink.Get(), o.InfoLink.IsSet()
}

// HasInfoLink returns a boolean if a field has been set.
func (o *ApplicationResource) HasInfoLink() bool {
	if o != nil && o.InfoLink.IsSet() {
		return true
	}

	return false
}

// SetInfoLink gets a reference to the given NullableString and assigns it to the InfoLink field.
func (o *ApplicationResource) SetInfoLink(v string) {
	o.InfoLink.Set(&v)
}
// SetInfoLinkNil sets the value for InfoLink to be an explicit nil
func (o *ApplicationResource) SetInfoLinkNil() {
	o.InfoLink.Set(nil)
}

// UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
func (o *ApplicationResource) UnsetInfoLink() {
	o.InfoLink.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplicationResource) GetMessage() ProviderMessage {
	if o == nil || isNil(o.Message) {
		var ret ProviderMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResource) GetMessageOk() (*ProviderMessage, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplicationResource) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ProviderMessage and assigns it to the Message field.
func (o *ApplicationResource) SetMessage(v ProviderMessage) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetTags() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetTagsOk() ([]*int32, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationResource) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ApplicationResource) SetTags(v []*int32) {
	o.Tags = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetPresets() []*ApplicationResource {
	if o == nil {
		var ret []*ApplicationResource
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetPresetsOk() ([]*ApplicationResource, bool) {
	if o == nil || isNil(o.Presets) {
    return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *ApplicationResource) HasPresets() bool {
	if o != nil && isNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []ApplicationResource and assigns it to the Presets field.
func (o *ApplicationResource) SetPresets(v []*ApplicationResource) {
	o.Presets = v
}

// GetSyncLevel returns the SyncLevel field value if set, zero value otherwise.
func (o *ApplicationResource) GetSyncLevel() ApplicationSyncLevel {
	if o == nil || isNil(o.SyncLevel) {
		var ret ApplicationSyncLevel
		return ret
	}
	return *o.SyncLevel
}

// GetSyncLevelOk returns a tuple with the SyncLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResource) GetSyncLevelOk() (*ApplicationSyncLevel, bool) {
	if o == nil || isNil(o.SyncLevel) {
    return nil, false
	}
	return o.SyncLevel, true
}

// HasSyncLevel returns a boolean if a field has been set.
func (o *ApplicationResource) HasSyncLevel() bool {
	if o != nil && !isNil(o.SyncLevel) {
		return true
	}

	return false
}

// SetSyncLevel gets a reference to the given ApplicationSyncLevel and assigns it to the SyncLevel field.
func (o *ApplicationResource) SetSyncLevel(v ApplicationSyncLevel) {
	o.SyncLevel = &v
}

// GetTestCommand returns the TestCommand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationResource) GetTestCommand() string {
	if o == nil || isNil(o.TestCommand.Get()) {
		var ret string
		return ret
	}
	return *o.TestCommand.Get()
}

// GetTestCommandOk returns a tuple with the TestCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResource) GetTestCommandOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TestCommand.Get(), o.TestCommand.IsSet()
}

// HasTestCommand returns a boolean if a field has been set.
func (o *ApplicationResource) HasTestCommand() bool {
	if o != nil && o.TestCommand.IsSet() {
		return true
	}

	return false
}

// SetTestCommand gets a reference to the given NullableString and assigns it to the TestCommand field.
func (o *ApplicationResource) SetTestCommand(v string) {
	o.TestCommand.Set(&v)
}
// SetTestCommandNil sets the value for TestCommand to be an explicit nil
func (o *ApplicationResource) SetTestCommandNil() {
	o.TestCommand.Set(nil)
}

// UnsetTestCommand ensures that no value is present for TestCommand, not even an explicit nil
func (o *ApplicationResource) UnsetTestCommand() {
	o.TestCommand.Unset()
}

func (o ApplicationResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
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
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	if !isNil(o.SyncLevel) {
		toSerialize["syncLevel"] = o.SyncLevel
	}
	if o.TestCommand.IsSet() {
		toSerialize["testCommand"] = o.TestCommand.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationResource struct {
	value *ApplicationResource
	isSet bool
}

func (v NullableApplicationResource) Get() *ApplicationResource {
	return v.value
}

func (v *NullableApplicationResource) Set(val *ApplicationResource) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResource) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResource(val *ApplicationResource) *NullableApplicationResource {
	return &NullableApplicationResource{value: val, isSet: true}
}

func (v NullableApplicationResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

