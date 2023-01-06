/*
Prowlarr

Prowlarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"fmt"
)

// AuthenticationRequiredType the model 'AuthenticationRequiredType'
type AuthenticationRequiredType string

// List of AuthenticationRequiredType
const (
	AUTHENTICATIONREQUIREDTYPE_ENABLED AuthenticationRequiredType = "enabled"
	AUTHENTICATIONREQUIREDTYPE_DISABLED_FOR_LOCAL_ADDRESSES AuthenticationRequiredType = "disabledForLocalAddresses"
)

// All allowed values of AuthenticationRequiredType enum
var AllowedAuthenticationRequiredTypeEnumValues = []AuthenticationRequiredType{
	"enabled",
	"disabledForLocalAddresses",
}

func (v *AuthenticationRequiredType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticationRequiredType(value)
	for _, existing := range AllowedAuthenticationRequiredTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticationRequiredType", value)
}

// NewAuthenticationRequiredTypeFromValue returns a pointer to a valid AuthenticationRequiredType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticationRequiredTypeFromValue(v string) (*AuthenticationRequiredType, error) {
	ev := AuthenticationRequiredType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticationRequiredType: valid values are %v", v, AllowedAuthenticationRequiredTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticationRequiredType) IsValid() bool {
	for _, existing := range AllowedAuthenticationRequiredTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthenticationRequiredType value
func (v AuthenticationRequiredType) Ptr() *AuthenticationRequiredType {
	return &v
}

type NullableAuthenticationRequiredType struct {
	value *AuthenticationRequiredType
	isSet bool
}

func (v NullableAuthenticationRequiredType) Get() *AuthenticationRequiredType {
	return v.value
}

func (v *NullableAuthenticationRequiredType) Set(val *AuthenticationRequiredType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationRequiredType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationRequiredType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationRequiredType(val *AuthenticationRequiredType) *NullableAuthenticationRequiredType {
	return &NullableAuthenticationRequiredType{value: val, isSet: true}
}

func (v NullableAuthenticationRequiredType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationRequiredType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

