/*
Prowlarr

Prowlarr API docs

API version: v1.30.2.4939
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"time"
)

// checks if the LogFileResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogFileResource{}

// LogFileResource struct for LogFileResource
type LogFileResource struct {
	Id *int32 `json:"id,omitempty"`
	Filename NullableString `json:"filename,omitempty"`
	LastWriteTime *time.Time `json:"lastWriteTime,omitempty"`
	ContentsUrl NullableString `json:"contentsUrl,omitempty"`
	DownloadUrl NullableString `json:"downloadUrl,omitempty"`
}

// NewLogFileResource instantiates a new LogFileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogFileResource() *LogFileResource {
	this := LogFileResource{}
	return &this
}

// NewLogFileResourceWithDefaults instantiates a new LogFileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogFileResourceWithDefaults() *LogFileResource {
	this := LogFileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogFileResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogFileResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogFileResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LogFileResource) SetId(v int32) {
	o.Id = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogFileResource) GetFilename() string {
	if o == nil || IsNil(o.Filename.Get()) {
		var ret string
		return ret
	}
	return *o.Filename.Get()
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogFileResource) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filename.Get(), o.Filename.IsSet()
}

// HasFilename returns a boolean if a field has been set.
func (o *LogFileResource) HasFilename() bool {
	if o != nil && o.Filename.IsSet() {
		return true
	}

	return false
}

// SetFilename gets a reference to the given NullableString and assigns it to the Filename field.
func (o *LogFileResource) SetFilename(v string) {
	o.Filename.Set(&v)
}
// SetFilenameNil sets the value for Filename to be an explicit nil
func (o *LogFileResource) SetFilenameNil() {
	o.Filename.Set(nil)
}

// UnsetFilename ensures that no value is present for Filename, not even an explicit nil
func (o *LogFileResource) UnsetFilename() {
	o.Filename.Unset()
}

// GetLastWriteTime returns the LastWriteTime field value if set, zero value otherwise.
func (o *LogFileResource) GetLastWriteTime() time.Time {
	if o == nil || IsNil(o.LastWriteTime) {
		var ret time.Time
		return ret
	}
	return *o.LastWriteTime
}

// GetLastWriteTimeOk returns a tuple with the LastWriteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogFileResource) GetLastWriteTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastWriteTime) {
		return nil, false
	}
	return o.LastWriteTime, true
}

// HasLastWriteTime returns a boolean if a field has been set.
func (o *LogFileResource) HasLastWriteTime() bool {
	if o != nil && !IsNil(o.LastWriteTime) {
		return true
	}

	return false
}

// SetLastWriteTime gets a reference to the given time.Time and assigns it to the LastWriteTime field.
func (o *LogFileResource) SetLastWriteTime(v time.Time) {
	o.LastWriteTime = &v
}

// GetContentsUrl returns the ContentsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogFileResource) GetContentsUrl() string {
	if o == nil || IsNil(o.ContentsUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ContentsUrl.Get()
}

// GetContentsUrlOk returns a tuple with the ContentsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogFileResource) GetContentsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentsUrl.Get(), o.ContentsUrl.IsSet()
}

// HasContentsUrl returns a boolean if a field has been set.
func (o *LogFileResource) HasContentsUrl() bool {
	if o != nil && o.ContentsUrl.IsSet() {
		return true
	}

	return false
}

// SetContentsUrl gets a reference to the given NullableString and assigns it to the ContentsUrl field.
func (o *LogFileResource) SetContentsUrl(v string) {
	o.ContentsUrl.Set(&v)
}
// SetContentsUrlNil sets the value for ContentsUrl to be an explicit nil
func (o *LogFileResource) SetContentsUrlNil() {
	o.ContentsUrl.Set(nil)
}

// UnsetContentsUrl ensures that no value is present for ContentsUrl, not even an explicit nil
func (o *LogFileResource) UnsetContentsUrl() {
	o.ContentsUrl.Unset()
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogFileResource) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadUrl.Get()
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogFileResource) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadUrl.Get(), o.DownloadUrl.IsSet()
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *LogFileResource) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl.IsSet() {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given NullableString and assigns it to the DownloadUrl field.
func (o *LogFileResource) SetDownloadUrl(v string) {
	o.DownloadUrl.Set(&v)
}
// SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil
func (o *LogFileResource) SetDownloadUrlNil() {
	o.DownloadUrl.Set(nil)
}

// UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
func (o *LogFileResource) UnsetDownloadUrl() {
	o.DownloadUrl.Unset()
}

func (o LogFileResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogFileResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Filename.IsSet() {
		toSerialize["filename"] = o.Filename.Get()
	}
	if !IsNil(o.LastWriteTime) {
		toSerialize["lastWriteTime"] = o.LastWriteTime
	}
	if o.ContentsUrl.IsSet() {
		toSerialize["contentsUrl"] = o.ContentsUrl.Get()
	}
	if o.DownloadUrl.IsSet() {
		toSerialize["downloadUrl"] = o.DownloadUrl.Get()
	}
	return toSerialize, nil
}

type NullableLogFileResource struct {
	value *LogFileResource
	isSet bool
}

func (v NullableLogFileResource) Get() *LogFileResource {
	return v.value
}

func (v *NullableLogFileResource) Set(val *LogFileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableLogFileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableLogFileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogFileResource(val *LogFileResource) *NullableLogFileResource {
	return &NullableLogFileResource{value: val, isSet: true}
}

func (v NullableLogFileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogFileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


