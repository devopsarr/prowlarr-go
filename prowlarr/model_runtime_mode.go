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

// RuntimeMode the model 'RuntimeMode'
type RuntimeMode string

// List of RuntimeMode
const (
	RUNTIMEMODE_CONSOLE RuntimeMode = "console"
	RUNTIMEMODE_SERVICE RuntimeMode = "service"
	RUNTIMEMODE_TRAY RuntimeMode = "tray"
)

// All allowed values of RuntimeMode enum
var AllowedRuntimeModeEnumValues = []RuntimeMode{
	"console",
	"service",
	"tray",
}

func (v *RuntimeMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RuntimeMode(value)
	for _, existing := range AllowedRuntimeModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RuntimeMode", value)
}

// NewRuntimeModeFromValue returns a pointer to a valid RuntimeMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRuntimeModeFromValue(v string) (*RuntimeMode, error) {
	ev := RuntimeMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RuntimeMode: valid values are %v", v, AllowedRuntimeModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RuntimeMode) IsValid() bool {
	for _, existing := range AllowedRuntimeModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuntimeMode value
func (v RuntimeMode) Ptr() *RuntimeMode {
	return &v
}

type NullableRuntimeMode struct {
	value *RuntimeMode
	isSet bool
}

func (v NullableRuntimeMode) Get() *RuntimeMode {
	return v.value
}

func (v *NullableRuntimeMode) Set(val *RuntimeMode) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeMode) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeMode(val *RuntimeMode) *NullableRuntimeMode {
	return &NullableRuntimeMode{value: val, isSet: true}
}

func (v NullableRuntimeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

