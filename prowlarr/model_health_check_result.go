/*
Prowlarr

Prowlarr API docs

API version: v1.31.2.4975
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"fmt"
)

// HealthCheckResult the model 'HealthCheckResult'
type HealthCheckResult string

// List of HealthCheckResult
const (
	HEALTHCHECKRESULT_OK HealthCheckResult = "ok"
	HEALTHCHECKRESULT_NOTICE HealthCheckResult = "notice"
	HEALTHCHECKRESULT_WARNING HealthCheckResult = "warning"
	HEALTHCHECKRESULT_ERROR HealthCheckResult = "error"
)

// All allowed values of HealthCheckResult enum
var AllowedHealthCheckResultEnumValues = []HealthCheckResult{
	"ok",
	"notice",
	"warning",
	"error",
}

func (v *HealthCheckResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HealthCheckResult(value)
	for _, existing := range AllowedHealthCheckResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HealthCheckResult", value)
}

// NewHealthCheckResultFromValue returns a pointer to a valid HealthCheckResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHealthCheckResultFromValue(v string) (*HealthCheckResult, error) {
	ev := HealthCheckResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HealthCheckResult: valid values are %v", v, AllowedHealthCheckResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HealthCheckResult) IsValid() bool {
	for _, existing := range AllowedHealthCheckResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HealthCheckResult value
func (v HealthCheckResult) Ptr() *HealthCheckResult {
	return &v
}

type NullableHealthCheckResult struct {
	value *HealthCheckResult
	isSet bool
}

func (v NullableHealthCheckResult) Get() *HealthCheckResult {
	return v.value
}

func (v *NullableHealthCheckResult) Set(val *HealthCheckResult) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckResult) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckResult(val *HealthCheckResult) *NullableHealthCheckResult {
	return &NullableHealthCheckResult{value: val, isSet: true}
}

func (v NullableHealthCheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

