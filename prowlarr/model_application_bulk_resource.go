/*
Prowlarr

Prowlarr API docs

API version: v1.28.2.4885
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the ApplicationBulkResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationBulkResource{}

// ApplicationBulkResource struct for ApplicationBulkResource
type ApplicationBulkResource struct {
	Ids []int32 `json:"ids,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	ApplyTags *ApplyTags `json:"applyTags,omitempty"`
	SyncLevel *ApplicationSyncLevel `json:"syncLevel,omitempty"`
}

// NewApplicationBulkResource instantiates a new ApplicationBulkResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationBulkResource() *ApplicationBulkResource {
	this := ApplicationBulkResource{}
	return &this
}

// NewApplicationBulkResourceWithDefaults instantiates a new ApplicationBulkResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationBulkResourceWithDefaults() *ApplicationBulkResource {
	this := ApplicationBulkResource{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationBulkResource) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationBulkResource) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ApplicationBulkResource) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *ApplicationBulkResource) SetIds(v []int32) {
	o.Ids = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationBulkResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationBulkResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationBulkResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ApplicationBulkResource) SetTags(v []int32) {
	o.Tags = v
}

// GetApplyTags returns the ApplyTags field value if set, zero value otherwise.
func (o *ApplicationBulkResource) GetApplyTags() ApplyTags {
	if o == nil || IsNil(o.ApplyTags) {
		var ret ApplyTags
		return ret
	}
	return *o.ApplyTags
}

// GetApplyTagsOk returns a tuple with the ApplyTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationBulkResource) GetApplyTagsOk() (*ApplyTags, bool) {
	if o == nil || IsNil(o.ApplyTags) {
		return nil, false
	}
	return o.ApplyTags, true
}

// HasApplyTags returns a boolean if a field has been set.
func (o *ApplicationBulkResource) HasApplyTags() bool {
	if o != nil && !IsNil(o.ApplyTags) {
		return true
	}

	return false
}

// SetApplyTags gets a reference to the given ApplyTags and assigns it to the ApplyTags field.
func (o *ApplicationBulkResource) SetApplyTags(v ApplyTags) {
	o.ApplyTags = &v
}

// GetSyncLevel returns the SyncLevel field value if set, zero value otherwise.
func (o *ApplicationBulkResource) GetSyncLevel() ApplicationSyncLevel {
	if o == nil || IsNil(o.SyncLevel) {
		var ret ApplicationSyncLevel
		return ret
	}
	return *o.SyncLevel
}

// GetSyncLevelOk returns a tuple with the SyncLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationBulkResource) GetSyncLevelOk() (*ApplicationSyncLevel, bool) {
	if o == nil || IsNil(o.SyncLevel) {
		return nil, false
	}
	return o.SyncLevel, true
}

// HasSyncLevel returns a boolean if a field has been set.
func (o *ApplicationBulkResource) HasSyncLevel() bool {
	if o != nil && !IsNil(o.SyncLevel) {
		return true
	}

	return false
}

// SetSyncLevel gets a reference to the given ApplicationSyncLevel and assigns it to the SyncLevel field.
func (o *ApplicationBulkResource) SetSyncLevel(v ApplicationSyncLevel) {
	o.SyncLevel = &v
}

func (o ApplicationBulkResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationBulkResource) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SyncLevel) {
		toSerialize["syncLevel"] = o.SyncLevel
	}
	return toSerialize, nil
}

type NullableApplicationBulkResource struct {
	value *ApplicationBulkResource
	isSet bool
}

func (v NullableApplicationBulkResource) Get() *ApplicationBulkResource {
	return v.value
}

func (v *NullableApplicationBulkResource) Set(val *ApplicationBulkResource) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationBulkResource) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationBulkResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationBulkResource(val *ApplicationBulkResource) *NullableApplicationBulkResource {
	return &NullableApplicationBulkResource{value: val, isSet: true}
}

func (v NullableApplicationBulkResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationBulkResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


