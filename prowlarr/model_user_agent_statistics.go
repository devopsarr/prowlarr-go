/*
Prowlarr

Prowlarr API docs

API version: v1.16.2.4435
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the UserAgentStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAgentStatistics{}

// UserAgentStatistics struct for UserAgentStatistics
type UserAgentStatistics struct {
	UserAgent NullableString `json:"userAgent,omitempty"`
	NumberOfQueries *int32 `json:"numberOfQueries,omitempty"`
	NumberOfGrabs *int32 `json:"numberOfGrabs,omitempty"`
}

// NewUserAgentStatistics instantiates a new UserAgentStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAgentStatistics() *UserAgentStatistics {
	this := UserAgentStatistics{}
	return &this
}

// NewUserAgentStatisticsWithDefaults instantiates a new UserAgentStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAgentStatisticsWithDefaults() *UserAgentStatistics {
	this := UserAgentStatistics{}
	return &this
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAgentStatistics) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent.Get()) {
		var ret string
		return ret
	}
	return *o.UserAgent.Get()
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAgentStatistics) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAgent.Get(), o.UserAgent.IsSet()
}

// HasUserAgent returns a boolean if a field has been set.
func (o *UserAgentStatistics) HasUserAgent() bool {
	if o != nil && o.UserAgent.IsSet() {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given NullableString and assigns it to the UserAgent field.
func (o *UserAgentStatistics) SetUserAgent(v string) {
	o.UserAgent.Set(&v)
}
// SetUserAgentNil sets the value for UserAgent to be an explicit nil
func (o *UserAgentStatistics) SetUserAgentNil() {
	o.UserAgent.Set(nil)
}

// UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
func (o *UserAgentStatistics) UnsetUserAgent() {
	o.UserAgent.Unset()
}

// GetNumberOfQueries returns the NumberOfQueries field value if set, zero value otherwise.
func (o *UserAgentStatistics) GetNumberOfQueries() int32 {
	if o == nil || IsNil(o.NumberOfQueries) {
		var ret int32
		return ret
	}
	return *o.NumberOfQueries
}

// GetNumberOfQueriesOk returns a tuple with the NumberOfQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentStatistics) GetNumberOfQueriesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfQueries) {
		return nil, false
	}
	return o.NumberOfQueries, true
}

// HasNumberOfQueries returns a boolean if a field has been set.
func (o *UserAgentStatistics) HasNumberOfQueries() bool {
	if o != nil && !IsNil(o.NumberOfQueries) {
		return true
	}

	return false
}

// SetNumberOfQueries gets a reference to the given int32 and assigns it to the NumberOfQueries field.
func (o *UserAgentStatistics) SetNumberOfQueries(v int32) {
	o.NumberOfQueries = &v
}

// GetNumberOfGrabs returns the NumberOfGrabs field value if set, zero value otherwise.
func (o *UserAgentStatistics) GetNumberOfGrabs() int32 {
	if o == nil || IsNil(o.NumberOfGrabs) {
		var ret int32
		return ret
	}
	return *o.NumberOfGrabs
}

// GetNumberOfGrabsOk returns a tuple with the NumberOfGrabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentStatistics) GetNumberOfGrabsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfGrabs) {
		return nil, false
	}
	return o.NumberOfGrabs, true
}

// HasNumberOfGrabs returns a boolean if a field has been set.
func (o *UserAgentStatistics) HasNumberOfGrabs() bool {
	if o != nil && !IsNil(o.NumberOfGrabs) {
		return true
	}

	return false
}

// SetNumberOfGrabs gets a reference to the given int32 and assigns it to the NumberOfGrabs field.
func (o *UserAgentStatistics) SetNumberOfGrabs(v int32) {
	o.NumberOfGrabs = &v
}

func (o UserAgentStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAgentStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserAgent.IsSet() {
		toSerialize["userAgent"] = o.UserAgent.Get()
	}
	if !IsNil(o.NumberOfQueries) {
		toSerialize["numberOfQueries"] = o.NumberOfQueries
	}
	if !IsNil(o.NumberOfGrabs) {
		toSerialize["numberOfGrabs"] = o.NumberOfGrabs
	}
	return toSerialize, nil
}

type NullableUserAgentStatistics struct {
	value *UserAgentStatistics
	isSet bool
}

func (v NullableUserAgentStatistics) Get() *UserAgentStatistics {
	return v.value
}

func (v *NullableUserAgentStatistics) Set(val *UserAgentStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAgentStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAgentStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAgentStatistics(val *UserAgentStatistics) *NullableUserAgentStatistics {
	return &NullableUserAgentStatistics{value: val, isSet: true}
}

func (v NullableUserAgentStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAgentStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


