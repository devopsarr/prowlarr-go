/*
Prowlarr

Prowlarr API docs

API version: v1.21.2.4649
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the DevelopmentConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevelopmentConfigResource{}

// DevelopmentConfigResource struct for DevelopmentConfigResource
type DevelopmentConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	ConsoleLogLevel NullableString `json:"consoleLogLevel,omitempty"`
	LogSql *bool `json:"logSql,omitempty"`
	LogIndexerResponse *bool `json:"logIndexerResponse,omitempty"`
	LogRotate *int32 `json:"logRotate,omitempty"`
	FilterSentryEvents *bool `json:"filterSentryEvents,omitempty"`
}

// NewDevelopmentConfigResource instantiates a new DevelopmentConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevelopmentConfigResource() *DevelopmentConfigResource {
	this := DevelopmentConfigResource{}
	return &this
}

// NewDevelopmentConfigResourceWithDefaults instantiates a new DevelopmentConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevelopmentConfigResourceWithDefaults() *DevelopmentConfigResource {
	this := DevelopmentConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevelopmentConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevelopmentConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DevelopmentConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetConsoleLogLevel returns the ConsoleLogLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DevelopmentConfigResource) GetConsoleLogLevel() string {
	if o == nil || IsNil(o.ConsoleLogLevel.Get()) {
		var ret string
		return ret
	}
	return *o.ConsoleLogLevel.Get()
}

// GetConsoleLogLevelOk returns a tuple with the ConsoleLogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DevelopmentConfigResource) GetConsoleLogLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsoleLogLevel.Get(), o.ConsoleLogLevel.IsSet()
}

// HasConsoleLogLevel returns a boolean if a field has been set.
func (o *DevelopmentConfigResource) HasConsoleLogLevel() bool {
	if o != nil && o.ConsoleLogLevel.IsSet() {
		return true
	}

	return false
}

// SetConsoleLogLevel gets a reference to the given NullableString and assigns it to the ConsoleLogLevel field.
func (o *DevelopmentConfigResource) SetConsoleLogLevel(v string) {
	o.ConsoleLogLevel.Set(&v)
}
// SetConsoleLogLevelNil sets the value for ConsoleLogLevel to be an explicit nil
func (o *DevelopmentConfigResource) SetConsoleLogLevelNil() {
	o.ConsoleLogLevel.Set(nil)
}

// UnsetConsoleLogLevel ensures that no value is present for ConsoleLogLevel, not even an explicit nil
func (o *DevelopmentConfigResource) UnsetConsoleLogLevel() {
	o.ConsoleLogLevel.Unset()
}

// GetLogSql returns the LogSql field value if set, zero value otherwise.
func (o *DevelopmentConfigResource) GetLogSql() bool {
	if o == nil || IsNil(o.LogSql) {
		var ret bool
		return ret
	}
	return *o.LogSql
}

// GetLogSqlOk returns a tuple with the LogSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentConfigResource) GetLogSqlOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSql) {
		return nil, false
	}
	return o.LogSql, true
}

// HasLogSql returns a boolean if a field has been set.
func (o *DevelopmentConfigResource) HasLogSql() bool {
	if o != nil && !IsNil(o.LogSql) {
		return true
	}

	return false
}

// SetLogSql gets a reference to the given bool and assigns it to the LogSql field.
func (o *DevelopmentConfigResource) SetLogSql(v bool) {
	o.LogSql = &v
}

// GetLogIndexerResponse returns the LogIndexerResponse field value if set, zero value otherwise.
func (o *DevelopmentConfigResource) GetLogIndexerResponse() bool {
	if o == nil || IsNil(o.LogIndexerResponse) {
		var ret bool
		return ret
	}
	return *o.LogIndexerResponse
}

// GetLogIndexerResponseOk returns a tuple with the LogIndexerResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentConfigResource) GetLogIndexerResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.LogIndexerResponse) {
		return nil, false
	}
	return o.LogIndexerResponse, true
}

// HasLogIndexerResponse returns a boolean if a field has been set.
func (o *DevelopmentConfigResource) HasLogIndexerResponse() bool {
	if o != nil && !IsNil(o.LogIndexerResponse) {
		return true
	}

	return false
}

// SetLogIndexerResponse gets a reference to the given bool and assigns it to the LogIndexerResponse field.
func (o *DevelopmentConfigResource) SetLogIndexerResponse(v bool) {
	o.LogIndexerResponse = &v
}

// GetLogRotate returns the LogRotate field value if set, zero value otherwise.
func (o *DevelopmentConfigResource) GetLogRotate() int32 {
	if o == nil || IsNil(o.LogRotate) {
		var ret int32
		return ret
	}
	return *o.LogRotate
}

// GetLogRotateOk returns a tuple with the LogRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentConfigResource) GetLogRotateOk() (*int32, bool) {
	if o == nil || IsNil(o.LogRotate) {
		return nil, false
	}
	return o.LogRotate, true
}

// HasLogRotate returns a boolean if a field has been set.
func (o *DevelopmentConfigResource) HasLogRotate() bool {
	if o != nil && !IsNil(o.LogRotate) {
		return true
	}

	return false
}

// SetLogRotate gets a reference to the given int32 and assigns it to the LogRotate field.
func (o *DevelopmentConfigResource) SetLogRotate(v int32) {
	o.LogRotate = &v
}

// GetFilterSentryEvents returns the FilterSentryEvents field value if set, zero value otherwise.
func (o *DevelopmentConfigResource) GetFilterSentryEvents() bool {
	if o == nil || IsNil(o.FilterSentryEvents) {
		var ret bool
		return ret
	}
	return *o.FilterSentryEvents
}

// GetFilterSentryEventsOk returns a tuple with the FilterSentryEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentConfigResource) GetFilterSentryEventsOk() (*bool, bool) {
	if o == nil || IsNil(o.FilterSentryEvents) {
		return nil, false
	}
	return o.FilterSentryEvents, true
}

// HasFilterSentryEvents returns a boolean if a field has been set.
func (o *DevelopmentConfigResource) HasFilterSentryEvents() bool {
	if o != nil && !IsNil(o.FilterSentryEvents) {
		return true
	}

	return false
}

// SetFilterSentryEvents gets a reference to the given bool and assigns it to the FilterSentryEvents field.
func (o *DevelopmentConfigResource) SetFilterSentryEvents(v bool) {
	o.FilterSentryEvents = &v
}

func (o DevelopmentConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevelopmentConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ConsoleLogLevel.IsSet() {
		toSerialize["consoleLogLevel"] = o.ConsoleLogLevel.Get()
	}
	if !IsNil(o.LogSql) {
		toSerialize["logSql"] = o.LogSql
	}
	if !IsNil(o.LogIndexerResponse) {
		toSerialize["logIndexerResponse"] = o.LogIndexerResponse
	}
	if !IsNil(o.LogRotate) {
		toSerialize["logRotate"] = o.LogRotate
	}
	if !IsNil(o.FilterSentryEvents) {
		toSerialize["filterSentryEvents"] = o.FilterSentryEvents
	}
	return toSerialize, nil
}

type NullableDevelopmentConfigResource struct {
	value *DevelopmentConfigResource
	isSet bool
}

func (v NullableDevelopmentConfigResource) Get() *DevelopmentConfigResource {
	return v.value
}

func (v *NullableDevelopmentConfigResource) Set(val *DevelopmentConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDevelopmentConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDevelopmentConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevelopmentConfigResource(val *DevelopmentConfigResource) *NullableDevelopmentConfigResource {
	return &NullableDevelopmentConfigResource{value: val, isSet: true}
}

func (v NullableDevelopmentConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevelopmentConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


