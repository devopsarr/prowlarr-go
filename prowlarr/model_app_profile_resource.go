/*
Prowlarr

Prowlarr API docs

API version: v1.20.1.4603
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the AppProfileResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppProfileResource{}

// AppProfileResource struct for AppProfileResource
type AppProfileResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	EnableRss *bool `json:"enableRss,omitempty"`
	EnableAutomaticSearch *bool `json:"enableAutomaticSearch,omitempty"`
	EnableInteractiveSearch *bool `json:"enableInteractiveSearch,omitempty"`
	MinimumSeeders *int32 `json:"minimumSeeders,omitempty"`
}

// NewAppProfileResource instantiates a new AppProfileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppProfileResource() *AppProfileResource {
	this := AppProfileResource{}
	return &this
}

// NewAppProfileResourceWithDefaults instantiates a new AppProfileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppProfileResourceWithDefaults() *AppProfileResource {
	this := AppProfileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppProfileResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppProfileResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppProfileResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AppProfileResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppProfileResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppProfileResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AppProfileResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AppProfileResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AppProfileResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AppProfileResource) UnsetName() {
	o.Name.Unset()
}

// GetEnableRss returns the EnableRss field value if set, zero value otherwise.
func (o *AppProfileResource) GetEnableRss() bool {
	if o == nil || IsNil(o.EnableRss) {
		var ret bool
		return ret
	}
	return *o.EnableRss
}

// GetEnableRssOk returns a tuple with the EnableRss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppProfileResource) GetEnableRssOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRss) {
		return nil, false
	}
	return o.EnableRss, true
}

// HasEnableRss returns a boolean if a field has been set.
func (o *AppProfileResource) HasEnableRss() bool {
	if o != nil && !IsNil(o.EnableRss) {
		return true
	}

	return false
}

// SetEnableRss gets a reference to the given bool and assigns it to the EnableRss field.
func (o *AppProfileResource) SetEnableRss(v bool) {
	o.EnableRss = &v
}

// GetEnableAutomaticSearch returns the EnableAutomaticSearch field value if set, zero value otherwise.
func (o *AppProfileResource) GetEnableAutomaticSearch() bool {
	if o == nil || IsNil(o.EnableAutomaticSearch) {
		var ret bool
		return ret
	}
	return *o.EnableAutomaticSearch
}

// GetEnableAutomaticSearchOk returns a tuple with the EnableAutomaticSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppProfileResource) GetEnableAutomaticSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutomaticSearch) {
		return nil, false
	}
	return o.EnableAutomaticSearch, true
}

// HasEnableAutomaticSearch returns a boolean if a field has been set.
func (o *AppProfileResource) HasEnableAutomaticSearch() bool {
	if o != nil && !IsNil(o.EnableAutomaticSearch) {
		return true
	}

	return false
}

// SetEnableAutomaticSearch gets a reference to the given bool and assigns it to the EnableAutomaticSearch field.
func (o *AppProfileResource) SetEnableAutomaticSearch(v bool) {
	o.EnableAutomaticSearch = &v
}

// GetEnableInteractiveSearch returns the EnableInteractiveSearch field value if set, zero value otherwise.
func (o *AppProfileResource) GetEnableInteractiveSearch() bool {
	if o == nil || IsNil(o.EnableInteractiveSearch) {
		var ret bool
		return ret
	}
	return *o.EnableInteractiveSearch
}

// GetEnableInteractiveSearchOk returns a tuple with the EnableInteractiveSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppProfileResource) GetEnableInteractiveSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableInteractiveSearch) {
		return nil, false
	}
	return o.EnableInteractiveSearch, true
}

// HasEnableInteractiveSearch returns a boolean if a field has been set.
func (o *AppProfileResource) HasEnableInteractiveSearch() bool {
	if o != nil && !IsNil(o.EnableInteractiveSearch) {
		return true
	}

	return false
}

// SetEnableInteractiveSearch gets a reference to the given bool and assigns it to the EnableInteractiveSearch field.
func (o *AppProfileResource) SetEnableInteractiveSearch(v bool) {
	o.EnableInteractiveSearch = &v
}

// GetMinimumSeeders returns the MinimumSeeders field value if set, zero value otherwise.
func (o *AppProfileResource) GetMinimumSeeders() int32 {
	if o == nil || IsNil(o.MinimumSeeders) {
		var ret int32
		return ret
	}
	return *o.MinimumSeeders
}

// GetMinimumSeedersOk returns a tuple with the MinimumSeeders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppProfileResource) GetMinimumSeedersOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumSeeders) {
		return nil, false
	}
	return o.MinimumSeeders, true
}

// HasMinimumSeeders returns a boolean if a field has been set.
func (o *AppProfileResource) HasMinimumSeeders() bool {
	if o != nil && !IsNil(o.MinimumSeeders) {
		return true
	}

	return false
}

// SetMinimumSeeders gets a reference to the given int32 and assigns it to the MinimumSeeders field.
func (o *AppProfileResource) SetMinimumSeeders(v int32) {
	o.MinimumSeeders = &v
}

func (o AppProfileResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppProfileResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.EnableRss) {
		toSerialize["enableRss"] = o.EnableRss
	}
	if !IsNil(o.EnableAutomaticSearch) {
		toSerialize["enableAutomaticSearch"] = o.EnableAutomaticSearch
	}
	if !IsNil(o.EnableInteractiveSearch) {
		toSerialize["enableInteractiveSearch"] = o.EnableInteractiveSearch
	}
	if !IsNil(o.MinimumSeeders) {
		toSerialize["minimumSeeders"] = o.MinimumSeeders
	}
	return toSerialize, nil
}

type NullableAppProfileResource struct {
	value *AppProfileResource
	isSet bool
}

func (v NullableAppProfileResource) Get() *AppProfileResource {
	return v.value
}

func (v *NullableAppProfileResource) Set(val *AppProfileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAppProfileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAppProfileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppProfileResource(val *AppProfileResource) *NullableAppProfileResource {
	return &NullableAppProfileResource{value: val, isSet: true}
}

func (v NullableAppProfileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppProfileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


