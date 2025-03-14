/*
Prowlarr

Prowlarr API docs

API version: v1.31.2.4975
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
)

// checks if the IndexerStatsResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexerStatsResource{}

// IndexerStatsResource struct for IndexerStatsResource
type IndexerStatsResource struct {
	Id *int32 `json:"id,omitempty"`
	Indexers []IndexerStatistics `json:"indexers,omitempty"`
	UserAgents []UserAgentStatistics `json:"userAgents,omitempty"`
	Hosts []HostStatistics `json:"hosts,omitempty"`
}

// NewIndexerStatsResource instantiates a new IndexerStatsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerStatsResource() *IndexerStatsResource {
	this := IndexerStatsResource{}
	return &this
}

// NewIndexerStatsResourceWithDefaults instantiates a new IndexerStatsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerStatsResourceWithDefaults() *IndexerStatsResource {
	this := IndexerStatsResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndexerStatsResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerStatsResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndexerStatsResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IndexerStatsResource) SetId(v int32) {
	o.Id = &v
}

// GetIndexers returns the Indexers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerStatsResource) GetIndexers() []IndexerStatistics {
	if o == nil {
		var ret []IndexerStatistics
		return ret
	}
	return o.Indexers
}

// GetIndexersOk returns a tuple with the Indexers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerStatsResource) GetIndexersOk() ([]IndexerStatistics, bool) {
	if o == nil || IsNil(o.Indexers) {
		return nil, false
	}
	return o.Indexers, true
}

// HasIndexers returns a boolean if a field has been set.
func (o *IndexerStatsResource) HasIndexers() bool {
	if o != nil && !IsNil(o.Indexers) {
		return true
	}

	return false
}

// SetIndexers gets a reference to the given []IndexerStatistics and assigns it to the Indexers field.
func (o *IndexerStatsResource) SetIndexers(v []IndexerStatistics) {
	o.Indexers = v
}

// GetUserAgents returns the UserAgents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerStatsResource) GetUserAgents() []UserAgentStatistics {
	if o == nil {
		var ret []UserAgentStatistics
		return ret
	}
	return o.UserAgents
}

// GetUserAgentsOk returns a tuple with the UserAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerStatsResource) GetUserAgentsOk() ([]UserAgentStatistics, bool) {
	if o == nil || IsNil(o.UserAgents) {
		return nil, false
	}
	return o.UserAgents, true
}

// HasUserAgents returns a boolean if a field has been set.
func (o *IndexerStatsResource) HasUserAgents() bool {
	if o != nil && !IsNil(o.UserAgents) {
		return true
	}

	return false
}

// SetUserAgents gets a reference to the given []UserAgentStatistics and assigns it to the UserAgents field.
func (o *IndexerStatsResource) SetUserAgents(v []UserAgentStatistics) {
	o.UserAgents = v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerStatsResource) GetHosts() []HostStatistics {
	if o == nil {
		var ret []HostStatistics
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerStatsResource) GetHostsOk() ([]HostStatistics, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *IndexerStatsResource) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []HostStatistics and assigns it to the Hosts field.
func (o *IndexerStatsResource) SetHosts(v []HostStatistics) {
	o.Hosts = v
}

func (o IndexerStatsResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexerStatsResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Indexers != nil {
		toSerialize["indexers"] = o.Indexers
	}
	if o.UserAgents != nil {
		toSerialize["userAgents"] = o.UserAgents
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	return toSerialize, nil
}

type NullableIndexerStatsResource struct {
	value *IndexerStatsResource
	isSet bool
}

func (v NullableIndexerStatsResource) Get() *IndexerStatsResource {
	return v.value
}

func (v *NullableIndexerStatsResource) Set(val *IndexerStatsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerStatsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerStatsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerStatsResource(val *IndexerStatsResource) *NullableIndexerStatsResource {
	return &NullableIndexerStatsResource{value: val, isSet: true}
}

func (v NullableIndexerStatsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerStatsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


