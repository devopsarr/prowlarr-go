/*
Prowlarr

Prowlarr API docs

API version: v1.23.1.4708
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"fmt"
)

// HistoryEventType the model 'HistoryEventType'
type HistoryEventType string

// List of HistoryEventType
const (
	HISTORYEVENTTYPE_UNKNOWN HistoryEventType = "unknown"
	HISTORYEVENTTYPE_RELEASE_GRABBED HistoryEventType = "releaseGrabbed"
	HISTORYEVENTTYPE_INDEXER_QUERY HistoryEventType = "indexerQuery"
	HISTORYEVENTTYPE_INDEXER_RSS HistoryEventType = "indexerRss"
	HISTORYEVENTTYPE_INDEXER_AUTH HistoryEventType = "indexerAuth"
	HISTORYEVENTTYPE_INDEXER_INFO HistoryEventType = "indexerInfo"
)

// All allowed values of HistoryEventType enum
var AllowedHistoryEventTypeEnumValues = []HistoryEventType{
	"unknown",
	"releaseGrabbed",
	"indexerQuery",
	"indexerRss",
	"indexerAuth",
	"indexerInfo",
}

func (v *HistoryEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HistoryEventType(value)
	for _, existing := range AllowedHistoryEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HistoryEventType", value)
}

// NewHistoryEventTypeFromValue returns a pointer to a valid HistoryEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHistoryEventTypeFromValue(v string) (*HistoryEventType, error) {
	ev := HistoryEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HistoryEventType: valid values are %v", v, AllowedHistoryEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HistoryEventType) IsValid() bool {
	for _, existing := range AllowedHistoryEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HistoryEventType value
func (v HistoryEventType) Ptr() *HistoryEventType {
	return &v
}

type NullableHistoryEventType struct {
	value *HistoryEventType
	isSet bool
}

func (v NullableHistoryEventType) Get() *HistoryEventType {
	return v.value
}

func (v *NullableHistoryEventType) Set(val *HistoryEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryEventType(val *HistoryEventType) *NullableHistoryEventType {
	return &NullableHistoryEventType{value: val, isSet: true}
}

func (v NullableHistoryEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

