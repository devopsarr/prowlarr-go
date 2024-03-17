/*
Prowlarr

Prowlarr API docs

API version: v1.14.3.4333
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"time"
)

// checks if the IndexerStatusResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexerStatusResource{}

// IndexerStatusResource struct for IndexerStatusResource
type IndexerStatusResource struct {
	Id *int32 `json:"id,omitempty"`
	IndexerId *int32 `json:"indexerId,omitempty"`
	DisabledTill NullableTime `json:"disabledTill,omitempty"`
	MostRecentFailure NullableTime `json:"mostRecentFailure,omitempty"`
	InitialFailure NullableTime `json:"initialFailure,omitempty"`
}

// NewIndexerStatusResource instantiates a new IndexerStatusResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerStatusResource() *IndexerStatusResource {
	this := IndexerStatusResource{}
	return &this
}

// NewIndexerStatusResourceWithDefaults instantiates a new IndexerStatusResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerStatusResourceWithDefaults() *IndexerStatusResource {
	this := IndexerStatusResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndexerStatusResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatusResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndexerStatusResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IndexerStatusResource) SetId(v int32) {
	o.Id = &v
}

// GetIndexerId returns the IndexerId field value if set, zero value otherwise.
func (o *IndexerStatusResource) GetIndexerId() int32 {
	if o == nil || IsNil(o.IndexerId) {
		var ret int32
		return ret
	}
	return *o.IndexerId
}

// GetIndexerIdOk returns a tuple with the IndexerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatusResource) GetIndexerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IndexerId) {
		return nil, false
	}
	return o.IndexerId, true
}

// HasIndexerId returns a boolean if a field has been set.
func (o *IndexerStatusResource) HasIndexerId() bool {
	if o != nil && !IsNil(o.IndexerId) {
		return true
	}

	return false
}

// SetIndexerId gets a reference to the given int32 and assigns it to the IndexerId field.
func (o *IndexerStatusResource) SetIndexerId(v int32) {
	o.IndexerId = &v
}

// GetDisabledTill returns the DisabledTill field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerStatusResource) GetDisabledTill() time.Time {
	if o == nil || IsNil(o.DisabledTill.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DisabledTill.Get()
}

// GetDisabledTillOk returns a tuple with the DisabledTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerStatusResource) GetDisabledTillOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisabledTill.Get(), o.DisabledTill.IsSet()
}

// HasDisabledTill returns a boolean if a field has been set.
func (o *IndexerStatusResource) HasDisabledTill() bool {
	if o != nil && o.DisabledTill.IsSet() {
		return true
	}

	return false
}

// SetDisabledTill gets a reference to the given NullableTime and assigns it to the DisabledTill field.
func (o *IndexerStatusResource) SetDisabledTill(v time.Time) {
	o.DisabledTill.Set(&v)
}
// SetDisabledTillNil sets the value for DisabledTill to be an explicit nil
func (o *IndexerStatusResource) SetDisabledTillNil() {
	o.DisabledTill.Set(nil)
}

// UnsetDisabledTill ensures that no value is present for DisabledTill, not even an explicit nil
func (o *IndexerStatusResource) UnsetDisabledTill() {
	o.DisabledTill.Unset()
}

// GetMostRecentFailure returns the MostRecentFailure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerStatusResource) GetMostRecentFailure() time.Time {
	if o == nil || IsNil(o.MostRecentFailure.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MostRecentFailure.Get()
}

// GetMostRecentFailureOk returns a tuple with the MostRecentFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerStatusResource) GetMostRecentFailureOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MostRecentFailure.Get(), o.MostRecentFailure.IsSet()
}

// HasMostRecentFailure returns a boolean if a field has been set.
func (o *IndexerStatusResource) HasMostRecentFailure() bool {
	if o != nil && o.MostRecentFailure.IsSet() {
		return true
	}

	return false
}

// SetMostRecentFailure gets a reference to the given NullableTime and assigns it to the MostRecentFailure field.
func (o *IndexerStatusResource) SetMostRecentFailure(v time.Time) {
	o.MostRecentFailure.Set(&v)
}
// SetMostRecentFailureNil sets the value for MostRecentFailure to be an explicit nil
func (o *IndexerStatusResource) SetMostRecentFailureNil() {
	o.MostRecentFailure.Set(nil)
}

// UnsetMostRecentFailure ensures that no value is present for MostRecentFailure, not even an explicit nil
func (o *IndexerStatusResource) UnsetMostRecentFailure() {
	o.MostRecentFailure.Unset()
}

// GetInitialFailure returns the InitialFailure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerStatusResource) GetInitialFailure() time.Time {
	if o == nil || IsNil(o.InitialFailure.Get()) {
		var ret time.Time
		return ret
	}
	return *o.InitialFailure.Get()
}

// GetInitialFailureOk returns a tuple with the InitialFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerStatusResource) GetInitialFailureOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitialFailure.Get(), o.InitialFailure.IsSet()
}

// HasInitialFailure returns a boolean if a field has been set.
func (o *IndexerStatusResource) HasInitialFailure() bool {
	if o != nil && o.InitialFailure.IsSet() {
		return true
	}

	return false
}

// SetInitialFailure gets a reference to the given NullableTime and assigns it to the InitialFailure field.
func (o *IndexerStatusResource) SetInitialFailure(v time.Time) {
	o.InitialFailure.Set(&v)
}
// SetInitialFailureNil sets the value for InitialFailure to be an explicit nil
func (o *IndexerStatusResource) SetInitialFailureNil() {
	o.InitialFailure.Set(nil)
}

// UnsetInitialFailure ensures that no value is present for InitialFailure, not even an explicit nil
func (o *IndexerStatusResource) UnsetInitialFailure() {
	o.InitialFailure.Unset()
}

func (o IndexerStatusResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexerStatusResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IndexerId) {
		toSerialize["indexerId"] = o.IndexerId
	}
	if o.DisabledTill.IsSet() {
		toSerialize["disabledTill"] = o.DisabledTill.Get()
	}
	if o.MostRecentFailure.IsSet() {
		toSerialize["mostRecentFailure"] = o.MostRecentFailure.Get()
	}
	if o.InitialFailure.IsSet() {
		toSerialize["initialFailure"] = o.InitialFailure.Get()
	}
	return toSerialize, nil
}

type NullableIndexerStatusResource struct {
	value *IndexerStatusResource
	isSet bool
}

func (v NullableIndexerStatusResource) Get() *IndexerStatusResource {
	return v.value
}

func (v *NullableIndexerStatusResource) Set(val *IndexerStatusResource) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerStatusResource) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerStatusResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerStatusResource(val *IndexerStatusResource) *NullableIndexerStatusResource {
	return &NullableIndexerStatusResource{value: val, isSet: true}
}

func (v NullableIndexerStatusResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerStatusResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


