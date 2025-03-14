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

// SearchParam the model 'SearchParam'
type SearchParam string

// List of SearchParam
const (
	SEARCHPARAM_Q SearchParam = "q"
)

// All allowed values of SearchParam enum
var AllowedSearchParamEnumValues = []SearchParam{
	"q",
}

func (v *SearchParam) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchParam(value)
	for _, existing := range AllowedSearchParamEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchParam", value)
}

// NewSearchParamFromValue returns a pointer to a valid SearchParam
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchParamFromValue(v string) (*SearchParam, error) {
	ev := SearchParam(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchParam: valid values are %v", v, AllowedSearchParamEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchParam) IsValid() bool {
	for _, existing := range AllowedSearchParamEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SearchParam value
func (v SearchParam) Ptr() *SearchParam {
	return &v
}

type NullableSearchParam struct {
	value *SearchParam
	isSet bool
}

func (v NullableSearchParam) Get() *SearchParam {
	return v.value
}

func (v *NullableSearchParam) Set(val *SearchParam) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchParam) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchParam(val *SearchParam) *NullableSearchParam {
	return &NullableSearchParam{value: val, isSet: true}
}

func (v NullableSearchParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

