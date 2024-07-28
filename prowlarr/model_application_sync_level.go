/*
Prowlarr

Prowlarr API docs

API version: v1.21.2.4649
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"fmt"
)

// ApplicationSyncLevel the model 'ApplicationSyncLevel'
type ApplicationSyncLevel string

// List of ApplicationSyncLevel
const (
	APPLICATIONSYNCLEVEL_DISABLED ApplicationSyncLevel = "disabled"
	APPLICATIONSYNCLEVEL_ADD_ONLY ApplicationSyncLevel = "addOnly"
	APPLICATIONSYNCLEVEL_FULL_SYNC ApplicationSyncLevel = "fullSync"
)

// All allowed values of ApplicationSyncLevel enum
var AllowedApplicationSyncLevelEnumValues = []ApplicationSyncLevel{
	"disabled",
	"addOnly",
	"fullSync",
}

func (v *ApplicationSyncLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplicationSyncLevel(value)
	for _, existing := range AllowedApplicationSyncLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplicationSyncLevel", value)
}

// NewApplicationSyncLevelFromValue returns a pointer to a valid ApplicationSyncLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplicationSyncLevelFromValue(v string) (*ApplicationSyncLevel, error) {
	ev := ApplicationSyncLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplicationSyncLevel: valid values are %v", v, AllowedApplicationSyncLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplicationSyncLevel) IsValid() bool {
	for _, existing := range AllowedApplicationSyncLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSyncLevel value
func (v ApplicationSyncLevel) Ptr() *ApplicationSyncLevel {
	return &v
}

type NullableApplicationSyncLevel struct {
	value *ApplicationSyncLevel
	isSet bool
}

func (v NullableApplicationSyncLevel) Get() *ApplicationSyncLevel {
	return v.value
}

func (v *NullableApplicationSyncLevel) Set(val *ApplicationSyncLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSyncLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSyncLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSyncLevel(val *ApplicationSyncLevel) *NullableApplicationSyncLevel {
	return &NullableApplicationSyncLevel{value: val, isSet: true}
}

func (v NullableApplicationSyncLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSyncLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

