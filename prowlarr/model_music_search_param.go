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

// MusicSearchParam the model 'MusicSearchParam'
type MusicSearchParam string

// List of MusicSearchParam
const (
	MUSICSEARCHPARAM_Q MusicSearchParam = "q"
	MUSICSEARCHPARAM_ALBUM MusicSearchParam = "album"
	MUSICSEARCHPARAM_ARTIST MusicSearchParam = "artist"
	MUSICSEARCHPARAM_LABEL MusicSearchParam = "label"
	MUSICSEARCHPARAM_YEAR MusicSearchParam = "year"
	MUSICSEARCHPARAM_GENRE MusicSearchParam = "genre"
	MUSICSEARCHPARAM_TRACK MusicSearchParam = "track"
)

// All allowed values of MusicSearchParam enum
var AllowedMusicSearchParamEnumValues = []MusicSearchParam{
	"q",
	"album",
	"artist",
	"label",
	"year",
	"genre",
	"track",
}

func (v *MusicSearchParam) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MusicSearchParam(value)
	for _, existing := range AllowedMusicSearchParamEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MusicSearchParam", value)
}

// NewMusicSearchParamFromValue returns a pointer to a valid MusicSearchParam
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMusicSearchParamFromValue(v string) (*MusicSearchParam, error) {
	ev := MusicSearchParam(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MusicSearchParam: valid values are %v", v, AllowedMusicSearchParamEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MusicSearchParam) IsValid() bool {
	for _, existing := range AllowedMusicSearchParamEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MusicSearchParam value
func (v MusicSearchParam) Ptr() *MusicSearchParam {
	return &v
}

type NullableMusicSearchParam struct {
	value *MusicSearchParam
	isSet bool
}

func (v NullableMusicSearchParam) Get() *MusicSearchParam {
	return v.value
}

func (v *NullableMusicSearchParam) Set(val *MusicSearchParam) {
	v.value = val
	v.isSet = true
}

func (v NullableMusicSearchParam) IsSet() bool {
	return v.isSet
}

func (v *NullableMusicSearchParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMusicSearchParam(val *MusicSearchParam) *NullableMusicSearchParam {
	return &NullableMusicSearchParam{value: val, isSet: true}
}

func (v NullableMusicSearchParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMusicSearchParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

