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

// PagingResourceFilter struct for PagingResourceFilter
type PagingResourceFilter struct {
	Key NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewPagingResourceFilter instantiates a new PagingResourceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingResourceFilter() *PagingResourceFilter {
	this := PagingResourceFilter{}
	return &this
}

// NewPagingResourceFilterWithDefaults instantiates a new PagingResourceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingResourceFilterWithDefaults() *PagingResourceFilter {
	this := PagingResourceFilter{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PagingResourceFilter) GetKey() string {
	if o == nil || isNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PagingResourceFilter) GetKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *PagingResourceFilter) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *PagingResourceFilter) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *PagingResourceFilter) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *PagingResourceFilter) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PagingResourceFilter) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PagingResourceFilter) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *PagingResourceFilter) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *PagingResourceFilter) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *PagingResourceFilter) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *PagingResourceFilter) UnsetValue() {
	o.Value.Unset()
}

func (o PagingResourceFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePagingResourceFilter struct {
	value *PagingResourceFilter
	isSet bool
}

func (v NullablePagingResourceFilter) Get() *PagingResourceFilter {
	return v.value
}

func (v *NullablePagingResourceFilter) Set(val *PagingResourceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingResourceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingResourceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingResourceFilter(val *PagingResourceFilter) *NullablePagingResourceFilter {
	return &NullablePagingResourceFilter{value: val, isSet: true}
}

func (v NullablePagingResourceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingResourceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


