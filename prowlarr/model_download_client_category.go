/*
Prowlarr

Prowlarr API docs

API version: v1.14.3.4333
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the DownloadClientCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadClientCategory{}

// DownloadClientCategory struct for DownloadClientCategory
type DownloadClientCategory struct {
	ClientCategory NullableString `json:"clientCategory,omitempty"`
	Categories []int32 `json:"categories,omitempty"`
}

// NewDownloadClientCategory instantiates a new DownloadClientCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadClientCategory() *DownloadClientCategory {
	this := DownloadClientCategory{}
	return &this
}

// NewDownloadClientCategoryWithDefaults instantiates a new DownloadClientCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadClientCategoryWithDefaults() *DownloadClientCategory {
	this := DownloadClientCategory{}
	return &this
}

// GetClientCategory returns the ClientCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientCategory) GetClientCategory() string {
	if o == nil || IsNil(o.ClientCategory.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCategory.Get()
}

// GetClientCategoryOk returns a tuple with the ClientCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientCategory) GetClientCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCategory.Get(), o.ClientCategory.IsSet()
}

// HasClientCategory returns a boolean if a field has been set.
func (o *DownloadClientCategory) HasClientCategory() bool {
	if o != nil && o.ClientCategory.IsSet() {
		return true
	}

	return false
}

// SetClientCategory gets a reference to the given NullableString and assigns it to the ClientCategory field.
func (o *DownloadClientCategory) SetClientCategory(v string) {
	o.ClientCategory.Set(&v)
}
// SetClientCategoryNil sets the value for ClientCategory to be an explicit nil
func (o *DownloadClientCategory) SetClientCategoryNil() {
	o.ClientCategory.Set(nil)
}

// UnsetClientCategory ensures that no value is present for ClientCategory, not even an explicit nil
func (o *DownloadClientCategory) UnsetClientCategory() {
	o.ClientCategory.Unset()
}

// GetCategories returns the Categories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientCategory) GetCategories() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientCategory) GetCategoriesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *DownloadClientCategory) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []int32 and assigns it to the Categories field.
func (o *DownloadClientCategory) SetCategories(v []int32) {
	o.Categories = v
}

func (o DownloadClientCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadClientCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientCategory.IsSet() {
		toSerialize["clientCategory"] = o.ClientCategory.Get()
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableDownloadClientCategory struct {
	value *DownloadClientCategory
	isSet bool
}

func (v NullableDownloadClientCategory) Get() *DownloadClientCategory {
	return v.value
}

func (v *NullableDownloadClientCategory) Set(val *DownloadClientCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadClientCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadClientCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadClientCategory(val *DownloadClientCategory) *NullableDownloadClientCategory {
	return &NullableDownloadClientCategory{value: val, isSet: true}
}

func (v NullableDownloadClientCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadClientCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


