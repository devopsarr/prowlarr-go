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

// IndexerCategory struct for IndexerCategory
type IndexerCategory struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	SubCategories []*IndexerCategory `json:"subCategories,omitempty"`
}

// NewIndexerCategory instantiates a new IndexerCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerCategory() *IndexerCategory {
	this := IndexerCategory{}
	return &this
}

// NewIndexerCategoryWithDefaults instantiates a new IndexerCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerCategoryWithDefaults() *IndexerCategory {
	this := IndexerCategory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndexerCategory) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerCategory) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndexerCategory) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IndexerCategory) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerCategory) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerCategory) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IndexerCategory) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IndexerCategory) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IndexerCategory) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IndexerCategory) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerCategory) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerCategory) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IndexerCategory) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IndexerCategory) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IndexerCategory) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IndexerCategory) UnsetDescription() {
	o.Description.Unset()
}

// GetSubCategories returns the SubCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerCategory) GetSubCategories() []*IndexerCategory {
	if o == nil {
		var ret []*IndexerCategory
		return ret
	}
	return o.SubCategories
}

// GetSubCategoriesOk returns a tuple with the SubCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerCategory) GetSubCategoriesOk() ([]*IndexerCategory, bool) {
	if o == nil || isNil(o.SubCategories) {
    return nil, false
	}
	return o.SubCategories, true
}

// HasSubCategories returns a boolean if a field has been set.
func (o *IndexerCategory) HasSubCategories() bool {
	if o != nil && isNil(o.SubCategories) {
		return true
	}

	return false
}

// SetSubCategories gets a reference to the given []IndexerCategory and assigns it to the SubCategories field.
func (o *IndexerCategory) SetSubCategories(v []*IndexerCategory) {
	o.SubCategories = v
}

func (o IndexerCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.SubCategories != nil {
		toSerialize["subCategories"] = o.SubCategories
	}
	return json.Marshal(toSerialize)
}

type NullableIndexerCategory struct {
	value *IndexerCategory
	isSet bool
}

func (v NullableIndexerCategory) Get() *IndexerCategory {
	return v.value
}

func (v *NullableIndexerCategory) Set(val *IndexerCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerCategory(val *IndexerCategory) *NullableIndexerCategory {
	return &NullableIndexerCategory{value: val, isSet: true}
}

func (v NullableIndexerCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

