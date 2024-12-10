/*
Prowlarr

Prowlarr API docs

API version: v1.27.0.4852
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"fmt"
)

// TvSearchParam the model 'TvSearchParam'
type TvSearchParam string

// List of TvSearchParam
const (
	TVSEARCHPARAM_Q TvSearchParam = "q"
	TVSEARCHPARAM_SEASON TvSearchParam = "season"
	TVSEARCHPARAM_EP TvSearchParam = "ep"
	TVSEARCHPARAM_IMDB_ID TvSearchParam = "imdbId"
	TVSEARCHPARAM_TVDB_ID TvSearchParam = "tvdbId"
	TVSEARCHPARAM_R_ID TvSearchParam = "rId"
	TVSEARCHPARAM_TV_MAZE_ID TvSearchParam = "tvMazeId"
	TVSEARCHPARAM_TRAKT_ID TvSearchParam = "traktId"
	TVSEARCHPARAM_TMDB_ID TvSearchParam = "tmdbId"
	TVSEARCHPARAM_DOUBAN_ID TvSearchParam = "doubanId"
	TVSEARCHPARAM_GENRE TvSearchParam = "genre"
	TVSEARCHPARAM_YEAR TvSearchParam = "year"
)

// All allowed values of TvSearchParam enum
var AllowedTvSearchParamEnumValues = []TvSearchParam{
	"q",
	"season",
	"ep",
	"imdbId",
	"tvdbId",
	"rId",
	"tvMazeId",
	"traktId",
	"tmdbId",
	"doubanId",
	"genre",
	"year",
}

func (v *TvSearchParam) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TvSearchParam(value)
	for _, existing := range AllowedTvSearchParamEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TvSearchParam", value)
}

// NewTvSearchParamFromValue returns a pointer to a valid TvSearchParam
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTvSearchParamFromValue(v string) (*TvSearchParam, error) {
	ev := TvSearchParam(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TvSearchParam: valid values are %v", v, AllowedTvSearchParamEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TvSearchParam) IsValid() bool {
	for _, existing := range AllowedTvSearchParamEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TvSearchParam value
func (v TvSearchParam) Ptr() *TvSearchParam {
	return &v
}

type NullableTvSearchParam struct {
	value *TvSearchParam
	isSet bool
}

func (v NullableTvSearchParam) Get() *TvSearchParam {
	return v.value
}

func (v *NullableTvSearchParam) Set(val *TvSearchParam) {
	v.value = val
	v.isSet = true
}

func (v NullableTvSearchParam) IsSet() bool {
	return v.isSet
}

func (v *NullableTvSearchParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTvSearchParam(val *TvSearchParam) *NullableTvSearchParam {
	return &NullableTvSearchParam{value: val, isSet: true}
}

func (v NullableTvSearchParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTvSearchParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

