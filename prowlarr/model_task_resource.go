/*
Prowlarr

Prowlarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"time"
)

// TaskResource struct for TaskResource
type TaskResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	TaskName NullableString `json:"taskName,omitempty"`
	Interval *int32 `json:"interval,omitempty"`
	LastExecution *time.Time `json:"lastExecution,omitempty"`
	LastStartTime *time.Time `json:"lastStartTime,omitempty"`
	NextExecution *time.Time `json:"nextExecution,omitempty"`
	LastDuration *TimeSpan `json:"lastDuration,omitempty"`
}

// NewTaskResource instantiates a new TaskResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResource() *TaskResource {
	this := TaskResource{}
	return &this
}

// NewTaskResourceWithDefaults instantiates a new TaskResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskResourceWithDefaults() *TaskResource {
	this := TaskResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TaskResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskResource) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskResource) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TaskResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TaskResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TaskResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TaskResource) UnsetName() {
	o.Name.Unset()
}

// GetTaskName returns the TaskName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskResource) GetTaskName() string {
	if o == nil || isNil(o.TaskName.Get()) {
		var ret string
		return ret
	}
	return *o.TaskName.Get()
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskResource) GetTaskNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TaskName.Get(), o.TaskName.IsSet()
}

// HasTaskName returns a boolean if a field has been set.
func (o *TaskResource) HasTaskName() bool {
	if o != nil && o.TaskName.IsSet() {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given NullableString and assigns it to the TaskName field.
func (o *TaskResource) SetTaskName(v string) {
	o.TaskName.Set(&v)
}
// SetTaskNameNil sets the value for TaskName to be an explicit nil
func (o *TaskResource) SetTaskNameNil() {
	o.TaskName.Set(nil)
}

// UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
func (o *TaskResource) UnsetTaskName() {
	o.TaskName.Unset()
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *TaskResource) GetInterval() int32 {
	if o == nil || isNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResource) GetIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.Interval) {
    return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *TaskResource) HasInterval() bool {
	if o != nil && !isNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *TaskResource) SetInterval(v int32) {
	o.Interval = &v
}

// GetLastExecution returns the LastExecution field value if set, zero value otherwise.
func (o *TaskResource) GetLastExecution() time.Time {
	if o == nil || isNil(o.LastExecution) {
		var ret time.Time
		return ret
	}
	return *o.LastExecution
}

// GetLastExecutionOk returns a tuple with the LastExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResource) GetLastExecutionOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastExecution) {
    return nil, false
	}
	return o.LastExecution, true
}

// HasLastExecution returns a boolean if a field has been set.
func (o *TaskResource) HasLastExecution() bool {
	if o != nil && !isNil(o.LastExecution) {
		return true
	}

	return false
}

// SetLastExecution gets a reference to the given time.Time and assigns it to the LastExecution field.
func (o *TaskResource) SetLastExecution(v time.Time) {
	o.LastExecution = &v
}

// GetLastStartTime returns the LastStartTime field value if set, zero value otherwise.
func (o *TaskResource) GetLastStartTime() time.Time {
	if o == nil || isNil(o.LastStartTime) {
		var ret time.Time
		return ret
	}
	return *o.LastStartTime
}

// GetLastStartTimeOk returns a tuple with the LastStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResource) GetLastStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastStartTime) {
    return nil, false
	}
	return o.LastStartTime, true
}

// HasLastStartTime returns a boolean if a field has been set.
func (o *TaskResource) HasLastStartTime() bool {
	if o != nil && !isNil(o.LastStartTime) {
		return true
	}

	return false
}

// SetLastStartTime gets a reference to the given time.Time and assigns it to the LastStartTime field.
func (o *TaskResource) SetLastStartTime(v time.Time) {
	o.LastStartTime = &v
}

// GetNextExecution returns the NextExecution field value if set, zero value otherwise.
func (o *TaskResource) GetNextExecution() time.Time {
	if o == nil || isNil(o.NextExecution) {
		var ret time.Time
		return ret
	}
	return *o.NextExecution
}

// GetNextExecutionOk returns a tuple with the NextExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResource) GetNextExecutionOk() (*time.Time, bool) {
	if o == nil || isNil(o.NextExecution) {
    return nil, false
	}
	return o.NextExecution, true
}

// HasNextExecution returns a boolean if a field has been set.
func (o *TaskResource) HasNextExecution() bool {
	if o != nil && !isNil(o.NextExecution) {
		return true
	}

	return false
}

// SetNextExecution gets a reference to the given time.Time and assigns it to the NextExecution field.
func (o *TaskResource) SetNextExecution(v time.Time) {
	o.NextExecution = &v
}

// GetLastDuration returns the LastDuration field value if set, zero value otherwise.
func (o *TaskResource) GetLastDuration() TimeSpan {
	if o == nil || isNil(o.LastDuration) {
		var ret TimeSpan
		return ret
	}
	return *o.LastDuration
}

// GetLastDurationOk returns a tuple with the LastDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResource) GetLastDurationOk() (*TimeSpan, bool) {
	if o == nil || isNil(o.LastDuration) {
    return nil, false
	}
	return o.LastDuration, true
}

// HasLastDuration returns a boolean if a field has been set.
func (o *TaskResource) HasLastDuration() bool {
	if o != nil && !isNil(o.LastDuration) {
		return true
	}

	return false
}

// SetLastDuration gets a reference to the given TimeSpan and assigns it to the LastDuration field.
func (o *TaskResource) SetLastDuration(v TimeSpan) {
	o.LastDuration = &v
}

func (o TaskResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TaskName.IsSet() {
		toSerialize["taskName"] = o.TaskName.Get()
	}
	if !isNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !isNil(o.LastExecution) {
		toSerialize["lastExecution"] = o.LastExecution
	}
	if !isNil(o.LastStartTime) {
		toSerialize["lastStartTime"] = o.LastStartTime
	}
	if !isNil(o.NextExecution) {
		toSerialize["nextExecution"] = o.NextExecution
	}
	if !isNil(o.LastDuration) {
		toSerialize["lastDuration"] = o.LastDuration
	}
	return json.Marshal(toSerialize)
}

type NullableTaskResource struct {
	value *TaskResource
	isSet bool
}

func (v NullableTaskResource) Get() *TaskResource {
	return v.value
}

func (v *NullableTaskResource) Set(val *TaskResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskResource(val *TaskResource) *NullableTaskResource {
	return &NullableTaskResource{value: val, isSet: true}
}

func (v NullableTaskResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

