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

// IndexerStatistics struct for IndexerStatistics
type IndexerStatistics struct {
	IndexerId *int32 `json:"indexerId,omitempty"`
	IndexerName NullableString `json:"indexerName,omitempty"`
	AverageResponseTime *int32 `json:"averageResponseTime,omitempty"`
	NumberOfQueries *int32 `json:"numberOfQueries,omitempty"`
	NumberOfGrabs *int32 `json:"numberOfGrabs,omitempty"`
	NumberOfRssQueries *int32 `json:"numberOfRssQueries,omitempty"`
	NumberOfAuthQueries *int32 `json:"numberOfAuthQueries,omitempty"`
	NumberOfFailedQueries *int32 `json:"numberOfFailedQueries,omitempty"`
	NumberOfFailedGrabs *int32 `json:"numberOfFailedGrabs,omitempty"`
	NumberOfFailedRssQueries *int32 `json:"numberOfFailedRssQueries,omitempty"`
	NumberOfFailedAuthQueries *int32 `json:"numberOfFailedAuthQueries,omitempty"`
}

// NewIndexerStatistics instantiates a new IndexerStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerStatistics() *IndexerStatistics {
	this := IndexerStatistics{}
	return &this
}

// NewIndexerStatisticsWithDefaults instantiates a new IndexerStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerStatisticsWithDefaults() *IndexerStatistics {
	this := IndexerStatistics{}
	return &this
}

// GetIndexerId returns the IndexerId field value if set, zero value otherwise.
func (o *IndexerStatistics) GetIndexerId() int32 {
	if o == nil || isNil(o.IndexerId) {
		var ret int32
		return ret
	}
	return *o.IndexerId
}

// GetIndexerIdOk returns a tuple with the IndexerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetIndexerIdOk() (*int32, bool) {
	if o == nil || isNil(o.IndexerId) {
    return nil, false
	}
	return o.IndexerId, true
}

// HasIndexerId returns a boolean if a field has been set.
func (o *IndexerStatistics) HasIndexerId() bool {
	if o != nil && !isNil(o.IndexerId) {
		return true
	}

	return false
}

// SetIndexerId gets a reference to the given int32 and assigns it to the IndexerId field.
func (o *IndexerStatistics) SetIndexerId(v int32) {
	o.IndexerId = &v
}

// GetIndexerName returns the IndexerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerStatistics) GetIndexerName() string {
	if o == nil || isNil(o.IndexerName.Get()) {
		var ret string
		return ret
	}
	return *o.IndexerName.Get()
}

// GetIndexerNameOk returns a tuple with the IndexerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerStatistics) GetIndexerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.IndexerName.Get(), o.IndexerName.IsSet()
}

// HasIndexerName returns a boolean if a field has been set.
func (o *IndexerStatistics) HasIndexerName() bool {
	if o != nil && o.IndexerName.IsSet() {
		return true
	}

	return false
}

// SetIndexerName gets a reference to the given NullableString and assigns it to the IndexerName field.
func (o *IndexerStatistics) SetIndexerName(v string) {
	o.IndexerName.Set(&v)
}
// SetIndexerNameNil sets the value for IndexerName to be an explicit nil
func (o *IndexerStatistics) SetIndexerNameNil() {
	o.IndexerName.Set(nil)
}

// UnsetIndexerName ensures that no value is present for IndexerName, not even an explicit nil
func (o *IndexerStatistics) UnsetIndexerName() {
	o.IndexerName.Unset()
}

// GetAverageResponseTime returns the AverageResponseTime field value if set, zero value otherwise.
func (o *IndexerStatistics) GetAverageResponseTime() int32 {
	if o == nil || isNil(o.AverageResponseTime) {
		var ret int32
		return ret
	}
	return *o.AverageResponseTime
}

// GetAverageResponseTimeOk returns a tuple with the AverageResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetAverageResponseTimeOk() (*int32, bool) {
	if o == nil || isNil(o.AverageResponseTime) {
    return nil, false
	}
	return o.AverageResponseTime, true
}

// HasAverageResponseTime returns a boolean if a field has been set.
func (o *IndexerStatistics) HasAverageResponseTime() bool {
	if o != nil && !isNil(o.AverageResponseTime) {
		return true
	}

	return false
}

// SetAverageResponseTime gets a reference to the given int32 and assigns it to the AverageResponseTime field.
func (o *IndexerStatistics) SetAverageResponseTime(v int32) {
	o.AverageResponseTime = &v
}

// GetNumberOfQueries returns the NumberOfQueries field value if set, zero value otherwise.
func (o *IndexerStatistics) GetNumberOfQueries() int32 {
	if o == nil || isNil(o.NumberOfQueries) {
		var ret int32
		return ret
	}
	return *o.NumberOfQueries
}

// GetNumberOfQueriesOk returns a tuple with the NumberOfQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetNumberOfQueriesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfQueries) {
    return nil, false
	}
	return o.NumberOfQueries, true
}

// HasNumberOfQueries returns a boolean if a field has been set.
func (o *IndexerStatistics) HasNumberOfQueries() bool {
	if o != nil && !isNil(o.NumberOfQueries) {
		return true
	}

	return false
}

// SetNumberOfQueries gets a reference to the given int32 and assigns it to the NumberOfQueries field.
func (o *IndexerStatistics) SetNumberOfQueries(v int32) {
	o.NumberOfQueries = &v
}

// GetNumberOfGrabs returns the NumberOfGrabs field value if set, zero value otherwise.
func (o *IndexerStatistics) GetNumberOfGrabs() int32 {
	if o == nil || isNil(o.NumberOfGrabs) {
		var ret int32
		return ret
	}
	return *o.NumberOfGrabs
}

// GetNumberOfGrabsOk returns a tuple with the NumberOfGrabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetNumberOfGrabsOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfGrabs) {
    return nil, false
	}
	return o.NumberOfGrabs, true
}

// HasNumberOfGrabs returns a boolean if a field has been set.
func (o *IndexerStatistics) HasNumberOfGrabs() bool {
	if o != nil && !isNil(o.NumberOfGrabs) {
		return true
	}

	return false
}

// SetNumberOfGrabs gets a reference to the given int32 and assigns it to the NumberOfGrabs field.
func (o *IndexerStatistics) SetNumberOfGrabs(v int32) {
	o.NumberOfGrabs = &v
}

// GetNumberOfRssQueries returns the NumberOfRssQueries field value if set, zero value otherwise.
func (o *IndexerStatistics) GetNumberOfRssQueries() int32 {
	if o == nil || isNil(o.NumberOfRssQueries) {
		var ret int32
		return ret
	}
	return *o.NumberOfRssQueries
}

// GetNumberOfRssQueriesOk returns a tuple with the NumberOfRssQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetNumberOfRssQueriesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfRssQueries) {
    return nil, false
	}
	return o.NumberOfRssQueries, true
}

// HasNumberOfRssQueries returns a boolean if a field has been set.
func (o *IndexerStatistics) HasNumberOfRssQueries() bool {
	if o != nil && !isNil(o.NumberOfRssQueries) {
		return true
	}

	return false
}

// SetNumberOfRssQueries gets a reference to the given int32 and assigns it to the NumberOfRssQueries field.
func (o *IndexerStatistics) SetNumberOfRssQueries(v int32) {
	o.NumberOfRssQueries = &v
}

// GetNumberOfAuthQueries returns the NumberOfAuthQueries field value if set, zero value otherwise.
func (o *IndexerStatistics) GetNumberOfAuthQueries() int32 {
	if o == nil || isNil(o.NumberOfAuthQueries) {
		var ret int32
		return ret
	}
	return *o.NumberOfAuthQueries
}

// GetNumberOfAuthQueriesOk returns a tuple with the NumberOfAuthQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetNumberOfAuthQueriesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfAuthQueries) {
    return nil, false
	}
	return o.NumberOfAuthQueries, true
}

// HasNumberOfAuthQueries returns a boolean if a field has been set.
func (o *IndexerStatistics) HasNumberOfAuthQueries() bool {
	if o != nil && !isNil(o.NumberOfAuthQueries) {
		return true
	}

	return false
}

// SetNumberOfAuthQueries gets a reference to the given int32 and assigns it to the NumberOfAuthQueries field.
func (o *IndexerStatistics) SetNumberOfAuthQueries(v int32) {
	o.NumberOfAuthQueries = &v
}

// GetNumberOfFailedQueries returns the NumberOfFailedQueries field value if set, zero value otherwise.
func (o *IndexerStatistics) GetNumberOfFailedQueries() int32 {
	if o == nil || isNil(o.NumberOfFailedQueries) {
		var ret int32
		return ret
	}
	return *o.NumberOfFailedQueries
}

// GetNumberOfFailedQueriesOk returns a tuple with the NumberOfFailedQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetNumberOfFailedQueriesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfFailedQueries) {
    return nil, false
	}
	return o.NumberOfFailedQueries, true
}

// HasNumberOfFailedQueries returns a boolean if a field has been set.
func (o *IndexerStatistics) HasNumberOfFailedQueries() bool {
	if o != nil && !isNil(o.NumberOfFailedQueries) {
		return true
	}

	return false
}

// SetNumberOfFailedQueries gets a reference to the given int32 and assigns it to the NumberOfFailedQueries field.
func (o *IndexerStatistics) SetNumberOfFailedQueries(v int32) {
	o.NumberOfFailedQueries = &v
}

// GetNumberOfFailedGrabs returns the NumberOfFailedGrabs field value if set, zero value otherwise.
func (o *IndexerStatistics) GetNumberOfFailedGrabs() int32 {
	if o == nil || isNil(o.NumberOfFailedGrabs) {
		var ret int32
		return ret
	}
	return *o.NumberOfFailedGrabs
}

// GetNumberOfFailedGrabsOk returns a tuple with the NumberOfFailedGrabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetNumberOfFailedGrabsOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfFailedGrabs) {
    return nil, false
	}
	return o.NumberOfFailedGrabs, true
}

// HasNumberOfFailedGrabs returns a boolean if a field has been set.
func (o *IndexerStatistics) HasNumberOfFailedGrabs() bool {
	if o != nil && !isNil(o.NumberOfFailedGrabs) {
		return true
	}

	return false
}

// SetNumberOfFailedGrabs gets a reference to the given int32 and assigns it to the NumberOfFailedGrabs field.
func (o *IndexerStatistics) SetNumberOfFailedGrabs(v int32) {
	o.NumberOfFailedGrabs = &v
}

// GetNumberOfFailedRssQueries returns the NumberOfFailedRssQueries field value if set, zero value otherwise.
func (o *IndexerStatistics) GetNumberOfFailedRssQueries() int32 {
	if o == nil || isNil(o.NumberOfFailedRssQueries) {
		var ret int32
		return ret
	}
	return *o.NumberOfFailedRssQueries
}

// GetNumberOfFailedRssQueriesOk returns a tuple with the NumberOfFailedRssQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetNumberOfFailedRssQueriesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfFailedRssQueries) {
    return nil, false
	}
	return o.NumberOfFailedRssQueries, true
}

// HasNumberOfFailedRssQueries returns a boolean if a field has been set.
func (o *IndexerStatistics) HasNumberOfFailedRssQueries() bool {
	if o != nil && !isNil(o.NumberOfFailedRssQueries) {
		return true
	}

	return false
}

// SetNumberOfFailedRssQueries gets a reference to the given int32 and assigns it to the NumberOfFailedRssQueries field.
func (o *IndexerStatistics) SetNumberOfFailedRssQueries(v int32) {
	o.NumberOfFailedRssQueries = &v
}

// GetNumberOfFailedAuthQueries returns the NumberOfFailedAuthQueries field value if set, zero value otherwise.
func (o *IndexerStatistics) GetNumberOfFailedAuthQueries() int32 {
	if o == nil || isNil(o.NumberOfFailedAuthQueries) {
		var ret int32
		return ret
	}
	return *o.NumberOfFailedAuthQueries
}

// GetNumberOfFailedAuthQueriesOk returns a tuple with the NumberOfFailedAuthQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatistics) GetNumberOfFailedAuthQueriesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfFailedAuthQueries) {
    return nil, false
	}
	return o.NumberOfFailedAuthQueries, true
}

// HasNumberOfFailedAuthQueries returns a boolean if a field has been set.
func (o *IndexerStatistics) HasNumberOfFailedAuthQueries() bool {
	if o != nil && !isNil(o.NumberOfFailedAuthQueries) {
		return true
	}

	return false
}

// SetNumberOfFailedAuthQueries gets a reference to the given int32 and assigns it to the NumberOfFailedAuthQueries field.
func (o *IndexerStatistics) SetNumberOfFailedAuthQueries(v int32) {
	o.NumberOfFailedAuthQueries = &v
}

func (o IndexerStatistics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IndexerId) {
		toSerialize["indexerId"] = o.IndexerId
	}
	if o.IndexerName.IsSet() {
		toSerialize["indexerName"] = o.IndexerName.Get()
	}
	if !isNil(o.AverageResponseTime) {
		toSerialize["averageResponseTime"] = o.AverageResponseTime
	}
	if !isNil(o.NumberOfQueries) {
		toSerialize["numberOfQueries"] = o.NumberOfQueries
	}
	if !isNil(o.NumberOfGrabs) {
		toSerialize["numberOfGrabs"] = o.NumberOfGrabs
	}
	if !isNil(o.NumberOfRssQueries) {
		toSerialize["numberOfRssQueries"] = o.NumberOfRssQueries
	}
	if !isNil(o.NumberOfAuthQueries) {
		toSerialize["numberOfAuthQueries"] = o.NumberOfAuthQueries
	}
	if !isNil(o.NumberOfFailedQueries) {
		toSerialize["numberOfFailedQueries"] = o.NumberOfFailedQueries
	}
	if !isNil(o.NumberOfFailedGrabs) {
		toSerialize["numberOfFailedGrabs"] = o.NumberOfFailedGrabs
	}
	if !isNil(o.NumberOfFailedRssQueries) {
		toSerialize["numberOfFailedRssQueries"] = o.NumberOfFailedRssQueries
	}
	if !isNil(o.NumberOfFailedAuthQueries) {
		toSerialize["numberOfFailedAuthQueries"] = o.NumberOfFailedAuthQueries
	}
	return json.Marshal(toSerialize)
}

type NullableIndexerStatistics struct {
	value *IndexerStatistics
	isSet bool
}

func (v NullableIndexerStatistics) Get() *IndexerStatistics {
	return v.value
}

func (v *NullableIndexerStatistics) Set(val *IndexerStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerStatistics(val *IndexerStatistics) *NullableIndexerStatistics {
	return &NullableIndexerStatistics{value: val, isSet: true}
}

func (v NullableIndexerStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


