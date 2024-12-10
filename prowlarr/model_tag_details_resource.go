/*
Prowlarr

Prowlarr API docs

API version: v1.27.0.4852
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the TagDetailsResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagDetailsResource{}

// TagDetailsResource struct for TagDetailsResource
type TagDetailsResource struct {
	Id *int32 `json:"id,omitempty"`
	Label NullableString `json:"label,omitempty"`
	NotificationIds []int32 `json:"notificationIds,omitempty"`
	IndexerIds []int32 `json:"indexerIds,omitempty"`
	IndexerProxyIds []int32 `json:"indexerProxyIds,omitempty"`
	ApplicationIds []int32 `json:"applicationIds,omitempty"`
}

// NewTagDetailsResource instantiates a new TagDetailsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagDetailsResource() *TagDetailsResource {
	this := TagDetailsResource{}
	return &this
}

// NewTagDetailsResourceWithDefaults instantiates a new TagDetailsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagDetailsResourceWithDefaults() *TagDetailsResource {
	this := TagDetailsResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TagDetailsResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDetailsResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TagDetailsResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TagDetailsResource) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *TagDetailsResource) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *TagDetailsResource) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *TagDetailsResource) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *TagDetailsResource) UnsetLabel() {
	o.Label.Unset()
}

// GetNotificationIds returns the NotificationIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetNotificationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.NotificationIds
}

// GetNotificationIdsOk returns a tuple with the NotificationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetNotificationIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.NotificationIds) {
		return nil, false
	}
	return o.NotificationIds, true
}

// HasNotificationIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasNotificationIds() bool {
	if o != nil && !IsNil(o.NotificationIds) {
		return true
	}

	return false
}

// SetNotificationIds gets a reference to the given []int32 and assigns it to the NotificationIds field.
func (o *TagDetailsResource) SetNotificationIds(v []int32) {
	o.NotificationIds = v
}

// GetIndexerIds returns the IndexerIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetIndexerIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.IndexerIds
}

// GetIndexerIdsOk returns a tuple with the IndexerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetIndexerIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.IndexerIds) {
		return nil, false
	}
	return o.IndexerIds, true
}

// HasIndexerIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasIndexerIds() bool {
	if o != nil && !IsNil(o.IndexerIds) {
		return true
	}

	return false
}

// SetIndexerIds gets a reference to the given []int32 and assigns it to the IndexerIds field.
func (o *TagDetailsResource) SetIndexerIds(v []int32) {
	o.IndexerIds = v
}

// GetIndexerProxyIds returns the IndexerProxyIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetIndexerProxyIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.IndexerProxyIds
}

// GetIndexerProxyIdsOk returns a tuple with the IndexerProxyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetIndexerProxyIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.IndexerProxyIds) {
		return nil, false
	}
	return o.IndexerProxyIds, true
}

// HasIndexerProxyIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasIndexerProxyIds() bool {
	if o != nil && !IsNil(o.IndexerProxyIds) {
		return true
	}

	return false
}

// SetIndexerProxyIds gets a reference to the given []int32 and assigns it to the IndexerProxyIds field.
func (o *TagDetailsResource) SetIndexerProxyIds(v []int32) {
	o.IndexerProxyIds = v
}

// GetApplicationIds returns the ApplicationIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetApplicationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ApplicationIds
}

// GetApplicationIdsOk returns a tuple with the ApplicationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetApplicationIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ApplicationIds) {
		return nil, false
	}
	return o.ApplicationIds, true
}

// HasApplicationIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasApplicationIds() bool {
	if o != nil && !IsNil(o.ApplicationIds) {
		return true
	}

	return false
}

// SetApplicationIds gets a reference to the given []int32 and assigns it to the ApplicationIds field.
func (o *TagDetailsResource) SetApplicationIds(v []int32) {
	o.ApplicationIds = v
}

func (o TagDetailsResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagDetailsResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.NotificationIds != nil {
		toSerialize["notificationIds"] = o.NotificationIds
	}
	if o.IndexerIds != nil {
		toSerialize["indexerIds"] = o.IndexerIds
	}
	if o.IndexerProxyIds != nil {
		toSerialize["indexerProxyIds"] = o.IndexerProxyIds
	}
	if o.ApplicationIds != nil {
		toSerialize["applicationIds"] = o.ApplicationIds
	}
	return toSerialize, nil
}

type NullableTagDetailsResource struct {
	value *TagDetailsResource
	isSet bool
}

func (v NullableTagDetailsResource) Get() *TagDetailsResource {
	return v.value
}

func (v *NullableTagDetailsResource) Set(val *TagDetailsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTagDetailsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTagDetailsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagDetailsResource(val *TagDetailsResource) *NullableTagDetailsResource {
	return &NullableTagDetailsResource{value: val, isSet: true}
}

func (v NullableTagDetailsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagDetailsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


