/*
Prowlarr

Prowlarr API docs

API version: v1.28.2.4885
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"fmt"
)

// IndexerPrivacy the model 'IndexerPrivacy'
type IndexerPrivacy string

// List of IndexerPrivacy
const (
	INDEXERPRIVACY_PUBLIC IndexerPrivacy = "public"
	INDEXERPRIVACY_SEMI_PRIVATE IndexerPrivacy = "semiPrivate"
	INDEXERPRIVACY_PRIVATE IndexerPrivacy = "private"
)

// All allowed values of IndexerPrivacy enum
var AllowedIndexerPrivacyEnumValues = []IndexerPrivacy{
	"public",
	"semiPrivate",
	"private",
}

func (v *IndexerPrivacy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IndexerPrivacy(value)
	for _, existing := range AllowedIndexerPrivacyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IndexerPrivacy", value)
}

// NewIndexerPrivacyFromValue returns a pointer to a valid IndexerPrivacy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIndexerPrivacyFromValue(v string) (*IndexerPrivacy, error) {
	ev := IndexerPrivacy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IndexerPrivacy: valid values are %v", v, AllowedIndexerPrivacyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IndexerPrivacy) IsValid() bool {
	for _, existing := range AllowedIndexerPrivacyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IndexerPrivacy value
func (v IndexerPrivacy) Ptr() *IndexerPrivacy {
	return &v
}

type NullableIndexerPrivacy struct {
	value *IndexerPrivacy
	isSet bool
}

func (v NullableIndexerPrivacy) Get() *IndexerPrivacy {
	return v.value
}

func (v *NullableIndexerPrivacy) Set(val *IndexerPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerPrivacy(val *IndexerPrivacy) *NullableIndexerPrivacy {
	return &NullableIndexerPrivacy{value: val, isSet: true}
}

func (v NullableIndexerPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

