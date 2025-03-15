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

// checks if the IndexerBulkResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexerBulkResource{}

// IndexerBulkResource struct for IndexerBulkResource
type IndexerBulkResource struct {
	Ids []int32 `json:"ids,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	ApplyTags *ApplyTags `json:"applyTags,omitempty"`
	Enable NullableBool `json:"enable,omitempty"`
	AppProfileId NullableInt32 `json:"appProfileId,omitempty"`
	Priority NullableInt32 `json:"priority,omitempty"`
	MinimumSeeders NullableInt32 `json:"minimumSeeders,omitempty"`
	SeedRatio NullableFloat64 `json:"seedRatio,omitempty"`
	SeedTime NullableInt32 `json:"seedTime,omitempty"`
	PackSeedTime NullableInt32 `json:"packSeedTime,omitempty"`
	PreferMagnetUrl NullableBool `json:"preferMagnetUrl,omitempty"`
}

// NewIndexerBulkResource instantiates a new IndexerBulkResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerBulkResource() *IndexerBulkResource {
	this := IndexerBulkResource{}
	return &this
}

// NewIndexerBulkResourceWithDefaults instantiates a new IndexerBulkResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerBulkResourceWithDefaults() *IndexerBulkResource {
	this := IndexerBulkResource{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *IndexerBulkResource) SetIds(v []int32) {
	o.Ids = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *IndexerBulkResource) SetTags(v []int32) {
	o.Tags = v
}

// GetApplyTags returns the ApplyTags field value if set, zero value otherwise.
func (o *IndexerBulkResource) GetApplyTags() ApplyTags {
	if o == nil || IsNil(o.ApplyTags) {
		var ret ApplyTags
		return ret
	}
	return *o.ApplyTags
}

// GetApplyTagsOk returns a tuple with the ApplyTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerBulkResource) GetApplyTagsOk() (*ApplyTags, bool) {
	if o == nil || IsNil(o.ApplyTags) {
		return nil, false
	}
	return o.ApplyTags, true
}

// HasApplyTags returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasApplyTags() bool {
	if o != nil && !IsNil(o.ApplyTags) {
		return true
	}

	return false
}

// SetApplyTags gets a reference to the given ApplyTags and assigns it to the ApplyTags field.
func (o *IndexerBulkResource) SetApplyTags(v ApplyTags) {
	o.ApplyTags = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetEnable() bool {
	if o == nil || IsNil(o.Enable.Get()) {
		var ret bool
		return ret
	}
	return *o.Enable.Get()
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enable.Get(), o.Enable.IsSet()
}

// HasEnable returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasEnable() bool {
	if o != nil && o.Enable.IsSet() {
		return true
	}

	return false
}

// SetEnable gets a reference to the given NullableBool and assigns it to the Enable field.
func (o *IndexerBulkResource) SetEnable(v bool) {
	o.Enable.Set(&v)
}
// SetEnableNil sets the value for Enable to be an explicit nil
func (o *IndexerBulkResource) SetEnableNil() {
	o.Enable.Set(nil)
}

// UnsetEnable ensures that no value is present for Enable, not even an explicit nil
func (o *IndexerBulkResource) UnsetEnable() {
	o.Enable.Unset()
}

// GetAppProfileId returns the AppProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetAppProfileId() int32 {
	if o == nil || IsNil(o.AppProfileId.Get()) {
		var ret int32
		return ret
	}
	return *o.AppProfileId.Get()
}

// GetAppProfileIdOk returns a tuple with the AppProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetAppProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppProfileId.Get(), o.AppProfileId.IsSet()
}

// HasAppProfileId returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasAppProfileId() bool {
	if o != nil && o.AppProfileId.IsSet() {
		return true
	}

	return false
}

// SetAppProfileId gets a reference to the given NullableInt32 and assigns it to the AppProfileId field.
func (o *IndexerBulkResource) SetAppProfileId(v int32) {
	o.AppProfileId.Set(&v)
}
// SetAppProfileIdNil sets the value for AppProfileId to be an explicit nil
func (o *IndexerBulkResource) SetAppProfileIdNil() {
	o.AppProfileId.Set(nil)
}

// UnsetAppProfileId ensures that no value is present for AppProfileId, not even an explicit nil
func (o *IndexerBulkResource) UnsetAppProfileId() {
	o.AppProfileId.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *IndexerBulkResource) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *IndexerBulkResource) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *IndexerBulkResource) UnsetPriority() {
	o.Priority.Unset()
}

// GetMinimumSeeders returns the MinimumSeeders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetMinimumSeeders() int32 {
	if o == nil || IsNil(o.MinimumSeeders.Get()) {
		var ret int32
		return ret
	}
	return *o.MinimumSeeders.Get()
}

// GetMinimumSeedersOk returns a tuple with the MinimumSeeders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetMinimumSeedersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumSeeders.Get(), o.MinimumSeeders.IsSet()
}

// HasMinimumSeeders returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasMinimumSeeders() bool {
	if o != nil && o.MinimumSeeders.IsSet() {
		return true
	}

	return false
}

// SetMinimumSeeders gets a reference to the given NullableInt32 and assigns it to the MinimumSeeders field.
func (o *IndexerBulkResource) SetMinimumSeeders(v int32) {
	o.MinimumSeeders.Set(&v)
}
// SetMinimumSeedersNil sets the value for MinimumSeeders to be an explicit nil
func (o *IndexerBulkResource) SetMinimumSeedersNil() {
	o.MinimumSeeders.Set(nil)
}

// UnsetMinimumSeeders ensures that no value is present for MinimumSeeders, not even an explicit nil
func (o *IndexerBulkResource) UnsetMinimumSeeders() {
	o.MinimumSeeders.Unset()
}

// GetSeedRatio returns the SeedRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetSeedRatio() float64 {
	if o == nil || IsNil(o.SeedRatio.Get()) {
		var ret float64
		return ret
	}
	return *o.SeedRatio.Get()
}

// GetSeedRatioOk returns a tuple with the SeedRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetSeedRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeedRatio.Get(), o.SeedRatio.IsSet()
}

// HasSeedRatio returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasSeedRatio() bool {
	if o != nil && o.SeedRatio.IsSet() {
		return true
	}

	return false
}

// SetSeedRatio gets a reference to the given NullableFloat64 and assigns it to the SeedRatio field.
func (o *IndexerBulkResource) SetSeedRatio(v float64) {
	o.SeedRatio.Set(&v)
}
// SetSeedRatioNil sets the value for SeedRatio to be an explicit nil
func (o *IndexerBulkResource) SetSeedRatioNil() {
	o.SeedRatio.Set(nil)
}

// UnsetSeedRatio ensures that no value is present for SeedRatio, not even an explicit nil
func (o *IndexerBulkResource) UnsetSeedRatio() {
	o.SeedRatio.Unset()
}

// GetSeedTime returns the SeedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetSeedTime() int32 {
	if o == nil || IsNil(o.SeedTime.Get()) {
		var ret int32
		return ret
	}
	return *o.SeedTime.Get()
}

// GetSeedTimeOk returns a tuple with the SeedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetSeedTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeedTime.Get(), o.SeedTime.IsSet()
}

// HasSeedTime returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasSeedTime() bool {
	if o != nil && o.SeedTime.IsSet() {
		return true
	}

	return false
}

// SetSeedTime gets a reference to the given NullableInt32 and assigns it to the SeedTime field.
func (o *IndexerBulkResource) SetSeedTime(v int32) {
	o.SeedTime.Set(&v)
}
// SetSeedTimeNil sets the value for SeedTime to be an explicit nil
func (o *IndexerBulkResource) SetSeedTimeNil() {
	o.SeedTime.Set(nil)
}

// UnsetSeedTime ensures that no value is present for SeedTime, not even an explicit nil
func (o *IndexerBulkResource) UnsetSeedTime() {
	o.SeedTime.Unset()
}

// GetPackSeedTime returns the PackSeedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetPackSeedTime() int32 {
	if o == nil || IsNil(o.PackSeedTime.Get()) {
		var ret int32
		return ret
	}
	return *o.PackSeedTime.Get()
}

// GetPackSeedTimeOk returns a tuple with the PackSeedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetPackSeedTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackSeedTime.Get(), o.PackSeedTime.IsSet()
}

// HasPackSeedTime returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasPackSeedTime() bool {
	if o != nil && o.PackSeedTime.IsSet() {
		return true
	}

	return false
}

// SetPackSeedTime gets a reference to the given NullableInt32 and assigns it to the PackSeedTime field.
func (o *IndexerBulkResource) SetPackSeedTime(v int32) {
	o.PackSeedTime.Set(&v)
}
// SetPackSeedTimeNil sets the value for PackSeedTime to be an explicit nil
func (o *IndexerBulkResource) SetPackSeedTimeNil() {
	o.PackSeedTime.Set(nil)
}

// UnsetPackSeedTime ensures that no value is present for PackSeedTime, not even an explicit nil
func (o *IndexerBulkResource) UnsetPackSeedTime() {
	o.PackSeedTime.Unset()
}

// GetPreferMagnetUrl returns the PreferMagnetUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerBulkResource) GetPreferMagnetUrl() bool {
	if o == nil || IsNil(o.PreferMagnetUrl.Get()) {
		var ret bool
		return ret
	}
	return *o.PreferMagnetUrl.Get()
}

// GetPreferMagnetUrlOk returns a tuple with the PreferMagnetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerBulkResource) GetPreferMagnetUrlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferMagnetUrl.Get(), o.PreferMagnetUrl.IsSet()
}

// HasPreferMagnetUrl returns a boolean if a field has been set.
func (o *IndexerBulkResource) HasPreferMagnetUrl() bool {
	if o != nil && o.PreferMagnetUrl.IsSet() {
		return true
	}

	return false
}

// SetPreferMagnetUrl gets a reference to the given NullableBool and assigns it to the PreferMagnetUrl field.
func (o *IndexerBulkResource) SetPreferMagnetUrl(v bool) {
	o.PreferMagnetUrl.Set(&v)
}
// SetPreferMagnetUrlNil sets the value for PreferMagnetUrl to be an explicit nil
func (o *IndexerBulkResource) SetPreferMagnetUrlNil() {
	o.PreferMagnetUrl.Set(nil)
}

// UnsetPreferMagnetUrl ensures that no value is present for PreferMagnetUrl, not even an explicit nil
func (o *IndexerBulkResource) UnsetPreferMagnetUrl() {
	o.PreferMagnetUrl.Unset()
}

func (o IndexerBulkResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexerBulkResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ApplyTags) {
		toSerialize["applyTags"] = o.ApplyTags
	}
	if o.Enable.IsSet() {
		toSerialize["enable"] = o.Enable.Get()
	}
	if o.AppProfileId.IsSet() {
		toSerialize["appProfileId"] = o.AppProfileId.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.MinimumSeeders.IsSet() {
		toSerialize["minimumSeeders"] = o.MinimumSeeders.Get()
	}
	if o.SeedRatio.IsSet() {
		toSerialize["seedRatio"] = o.SeedRatio.Get()
	}
	if o.SeedTime.IsSet() {
		toSerialize["seedTime"] = o.SeedTime.Get()
	}
	if o.PackSeedTime.IsSet() {
		toSerialize["packSeedTime"] = o.PackSeedTime.Get()
	}
	if o.PreferMagnetUrl.IsSet() {
		toSerialize["preferMagnetUrl"] = o.PreferMagnetUrl.Get()
	}
	return toSerialize, nil
}

type NullableIndexerBulkResource struct {
	value *IndexerBulkResource
	isSet bool
}

func (v NullableIndexerBulkResource) Get() *IndexerBulkResource {
	return v.value
}

func (v *NullableIndexerBulkResource) Set(val *IndexerBulkResource) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerBulkResource) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerBulkResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerBulkResource(val *IndexerBulkResource) *NullableIndexerBulkResource {
	return &NullableIndexerBulkResource{value: val, isSet: true}
}

func (v NullableIndexerBulkResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerBulkResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


