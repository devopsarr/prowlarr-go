/*
Prowlarr

Prowlarr API docs

API version: v1.24.3.4754
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the HostStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostStatistics{}

// HostStatistics struct for HostStatistics
type HostStatistics struct {
	Host NullableString `json:"host,omitempty"`
	NumberOfQueries *int32 `json:"numberOfQueries,omitempty"`
	NumberOfGrabs *int32 `json:"numberOfGrabs,omitempty"`
}

// NewHostStatistics instantiates a new HostStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostStatistics() *HostStatistics {
	this := HostStatistics{}
	return &this
}

// NewHostStatisticsWithDefaults instantiates a new HostStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostStatisticsWithDefaults() *HostStatistics {
	this := HostStatistics{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostStatistics) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostStatistics) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *HostStatistics) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *HostStatistics) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *HostStatistics) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *HostStatistics) UnsetHost() {
	o.Host.Unset()
}

// GetNumberOfQueries returns the NumberOfQueries field value if set, zero value otherwise.
func (o *HostStatistics) GetNumberOfQueries() int32 {
	if o == nil || IsNil(o.NumberOfQueries) {
		var ret int32
		return ret
	}
	return *o.NumberOfQueries
}

// GetNumberOfQueriesOk returns a tuple with the NumberOfQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostStatistics) GetNumberOfQueriesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfQueries) {
		return nil, false
	}
	return o.NumberOfQueries, true
}

// HasNumberOfQueries returns a boolean if a field has been set.
func (o *HostStatistics) HasNumberOfQueries() bool {
	if o != nil && !IsNil(o.NumberOfQueries) {
		return true
	}

	return false
}

// SetNumberOfQueries gets a reference to the given int32 and assigns it to the NumberOfQueries field.
func (o *HostStatistics) SetNumberOfQueries(v int32) {
	o.NumberOfQueries = &v
}

// GetNumberOfGrabs returns the NumberOfGrabs field value if set, zero value otherwise.
func (o *HostStatistics) GetNumberOfGrabs() int32 {
	if o == nil || IsNil(o.NumberOfGrabs) {
		var ret int32
		return ret
	}
	return *o.NumberOfGrabs
}

// GetNumberOfGrabsOk returns a tuple with the NumberOfGrabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostStatistics) GetNumberOfGrabsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfGrabs) {
		return nil, false
	}
	return o.NumberOfGrabs, true
}

// HasNumberOfGrabs returns a boolean if a field has been set.
func (o *HostStatistics) HasNumberOfGrabs() bool {
	if o != nil && !IsNil(o.NumberOfGrabs) {
		return true
	}

	return false
}

// SetNumberOfGrabs gets a reference to the given int32 and assigns it to the NumberOfGrabs field.
func (o *HostStatistics) SetNumberOfGrabs(v int32) {
	o.NumberOfGrabs = &v
}

func (o HostStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if !IsNil(o.NumberOfQueries) {
		toSerialize["numberOfQueries"] = o.NumberOfQueries
	}
	if !IsNil(o.NumberOfGrabs) {
		toSerialize["numberOfGrabs"] = o.NumberOfGrabs
	}
	return toSerialize, nil
}

type NullableHostStatistics struct {
	value *HostStatistics
	isSet bool
}

func (v NullableHostStatistics) Get() *HostStatistics {
	return v.value
}

func (v *NullableHostStatistics) Set(val *HostStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableHostStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableHostStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostStatistics(val *HostStatistics) *NullableHostStatistics {
	return &NullableHostStatistics{value: val, isSet: true}
}

func (v NullableHostStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


