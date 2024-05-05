/*
Prowlarr

Prowlarr API docs

API version: v1.16.2.4435
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"fmt"
)

// BookSearchParam the model 'BookSearchParam'
type BookSearchParam string

// List of BookSearchParam
const (
	BOOKSEARCHPARAM_Q BookSearchParam = "q"
	BOOKSEARCHPARAM_TITLE BookSearchParam = "title"
	BOOKSEARCHPARAM_AUTHOR BookSearchParam = "author"
	BOOKSEARCHPARAM_PUBLISHER BookSearchParam = "publisher"
	BOOKSEARCHPARAM_GENRE BookSearchParam = "genre"
	BOOKSEARCHPARAM_YEAR BookSearchParam = "year"
)

// All allowed values of BookSearchParam enum
var AllowedBookSearchParamEnumValues = []BookSearchParam{
	"q",
	"title",
	"author",
	"publisher",
	"genre",
	"year",
}

func (v *BookSearchParam) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BookSearchParam(value)
	for _, existing := range AllowedBookSearchParamEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BookSearchParam", value)
}

// NewBookSearchParamFromValue returns a pointer to a valid BookSearchParam
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBookSearchParamFromValue(v string) (*BookSearchParam, error) {
	ev := BookSearchParam(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BookSearchParam: valid values are %v", v, AllowedBookSearchParamEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BookSearchParam) IsValid() bool {
	for _, existing := range AllowedBookSearchParamEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BookSearchParam value
func (v BookSearchParam) Ptr() *BookSearchParam {
	return &v
}

type NullableBookSearchParam struct {
	value *BookSearchParam
	isSet bool
}

func (v NullableBookSearchParam) Get() *BookSearchParam {
	return v.value
}

func (v *NullableBookSearchParam) Set(val *BookSearchParam) {
	v.value = val
	v.isSet = true
}

func (v NullableBookSearchParam) IsSet() bool {
	return v.isSet
}

func (v *NullableBookSearchParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookSearchParam(val *BookSearchParam) *NullableBookSearchParam {
	return &NullableBookSearchParam{value: val, isSet: true}
}

func (v NullableBookSearchParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookSearchParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

