/*
Prowlarr

Prowlarr API docs

API version: v1.16.2.4435
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"time"
)

// checks if the HistoryResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryResource{}

// HistoryResource struct for HistoryResource
type HistoryResource struct {
	Id *int32 `json:"id,omitempty"`
	IndexerId *int32 `json:"indexerId,omitempty"`
	Date *time.Time `json:"date,omitempty"`
	DownloadId NullableString `json:"downloadId,omitempty"`
	Successful *bool `json:"successful,omitempty"`
	EventType *HistoryEventType `json:"eventType,omitempty"`
	Data map[string]string `json:"data,omitempty"`
}

// NewHistoryResource instantiates a new HistoryResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryResource() *HistoryResource {
	this := HistoryResource{}
	return &this
}

// NewHistoryResourceWithDefaults instantiates a new HistoryResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryResourceWithDefaults() *HistoryResource {
	this := HistoryResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HistoryResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HistoryResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *HistoryResource) SetId(v int32) {
	o.Id = &v
}

// GetIndexerId returns the IndexerId field value if set, zero value otherwise.
func (o *HistoryResource) GetIndexerId() int32 {
	if o == nil || IsNil(o.IndexerId) {
		var ret int32
		return ret
	}
	return *o.IndexerId
}

// GetIndexerIdOk returns a tuple with the IndexerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetIndexerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IndexerId) {
		return nil, false
	}
	return o.IndexerId, true
}

// HasIndexerId returns a boolean if a field has been set.
func (o *HistoryResource) HasIndexerId() bool {
	if o != nil && !IsNil(o.IndexerId) {
		return true
	}

	return false
}

// SetIndexerId gets a reference to the given int32 and assigns it to the IndexerId field.
func (o *HistoryResource) SetIndexerId(v int32) {
	o.IndexerId = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *HistoryResource) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *HistoryResource) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *HistoryResource) SetDate(v time.Time) {
	o.Date = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetDownloadId() string {
	if o == nil || IsNil(o.DownloadId.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadId.Get()
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetDownloadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadId.Get(), o.DownloadId.IsSet()
}

// HasDownloadId returns a boolean if a field has been set.
func (o *HistoryResource) HasDownloadId() bool {
	if o != nil && o.DownloadId.IsSet() {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given NullableString and assigns it to the DownloadId field.
func (o *HistoryResource) SetDownloadId(v string) {
	o.DownloadId.Set(&v)
}
// SetDownloadIdNil sets the value for DownloadId to be an explicit nil
func (o *HistoryResource) SetDownloadIdNil() {
	o.DownloadId.Set(nil)
}

// UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
func (o *HistoryResource) UnsetDownloadId() {
	o.DownloadId.Unset()
}

// GetSuccessful returns the Successful field value if set, zero value otherwise.
func (o *HistoryResource) GetSuccessful() bool {
	if o == nil || IsNil(o.Successful) {
		var ret bool
		return ret
	}
	return *o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetSuccessfulOk() (*bool, bool) {
	if o == nil || IsNil(o.Successful) {
		return nil, false
	}
	return o.Successful, true
}

// HasSuccessful returns a boolean if a field has been set.
func (o *HistoryResource) HasSuccessful() bool {
	if o != nil && !IsNil(o.Successful) {
		return true
	}

	return false
}

// SetSuccessful gets a reference to the given bool and assigns it to the Successful field.
func (o *HistoryResource) SetSuccessful(v bool) {
	o.Successful = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *HistoryResource) GetEventType() HistoryEventType {
	if o == nil || IsNil(o.EventType) {
		var ret HistoryEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetEventTypeOk() (*HistoryEventType, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *HistoryResource) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given HistoryEventType and assigns it to the EventType field.
func (o *HistoryResource) SetEventType(v HistoryEventType) {
	o.EventType = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoryResource) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *HistoryResource) SetData(v map[string]string) {
	o.Data = v
}

func (o HistoryResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IndexerId) {
		toSerialize["indexerId"] = o.IndexerId
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.DownloadId.IsSet() {
		toSerialize["downloadId"] = o.DownloadId.Get()
	}
	if !IsNil(o.Successful) {
		toSerialize["successful"] = o.Successful
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableHistoryResource struct {
	value *HistoryResource
	isSet bool
}

func (v NullableHistoryResource) Get() *HistoryResource {
	return v.value
}

func (v *NullableHistoryResource) Set(val *HistoryResource) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryResource) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryResource(val *HistoryResource) *NullableHistoryResource {
	return &NullableHistoryResource{value: val, isSet: true}
}

func (v NullableHistoryResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


