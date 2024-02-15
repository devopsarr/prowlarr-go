/*
Prowlarr

Prowlarr API docs

API version: v1.13.3.4273
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the ApiInfoResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInfoResource{}

// ApiInfoResource struct for ApiInfoResource
type ApiInfoResource struct {
	Current NullableString `json:"current,omitempty"`
	Deprecated []string `json:"deprecated,omitempty"`
}

// NewApiInfoResource instantiates a new ApiInfoResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInfoResource() *ApiInfoResource {
	this := ApiInfoResource{}
	return &this
}

// NewApiInfoResourceWithDefaults instantiates a new ApiInfoResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInfoResourceWithDefaults() *ApiInfoResource {
	this := ApiInfoResource{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiInfoResource) GetCurrent() string {
	if o == nil || IsNil(o.Current.Get()) {
		var ret string
		return ret
	}
	return *o.Current.Get()
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiInfoResource) GetCurrentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Current.Get(), o.Current.IsSet()
}

// HasCurrent returns a boolean if a field has been set.
func (o *ApiInfoResource) HasCurrent() bool {
	if o != nil && o.Current.IsSet() {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given NullableString and assigns it to the Current field.
func (o *ApiInfoResource) SetCurrent(v string) {
	o.Current.Set(&v)
}
// SetCurrentNil sets the value for Current to be an explicit nil
func (o *ApiInfoResource) SetCurrentNil() {
	o.Current.Set(nil)
}

// UnsetCurrent ensures that no value is present for Current, not even an explicit nil
func (o *ApiInfoResource) UnsetCurrent() {
	o.Current.Unset()
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiInfoResource) GetDeprecated() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiInfoResource) GetDeprecatedOk() ([]string, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *ApiInfoResource) HasDeprecated() bool {
	if o != nil && IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given []string and assigns it to the Deprecated field.
func (o *ApiInfoResource) SetDeprecated(v []string) {
	o.Deprecated = v
}

func (o ApiInfoResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInfoResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Current.IsSet() {
		toSerialize["current"] = o.Current.Get()
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	return toSerialize, nil
}

type NullableApiInfoResource struct {
	value *ApiInfoResource
	isSet bool
}

func (v NullableApiInfoResource) Get() *ApiInfoResource {
	return v.value
}

func (v *NullableApiInfoResource) Set(val *ApiInfoResource) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInfoResource) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInfoResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInfoResource(val *ApiInfoResource) *NullableApiInfoResource {
	return &NullableApiInfoResource{value: val, isSet: true}
}

func (v NullableApiInfoResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInfoResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


