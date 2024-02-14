/*
Prowlarr

Prowlarr API docs

API version: v1.13.3.4273
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package prowlarr

import (
	"encoding/json"
	"fmt"
)

// MovieSearchParam the model 'MovieSearchParam'
type MovieSearchParam string

// List of MovieSearchParam
const (
	MOVIESEARCHPARAM_Q MovieSearchParam = "q"
	MOVIESEARCHPARAM_IMDB_ID MovieSearchParam = "imdbId"
	MOVIESEARCHPARAM_TMDB_ID MovieSearchParam = "tmdbId"
	MOVIESEARCHPARAM_IMDB_TITLE MovieSearchParam = "imdbTitle"
	MOVIESEARCHPARAM_IMDB_YEAR MovieSearchParam = "imdbYear"
	MOVIESEARCHPARAM_TRAKT_ID MovieSearchParam = "traktId"
	MOVIESEARCHPARAM_GENRE MovieSearchParam = "genre"
	MOVIESEARCHPARAM_DOUBAN_ID MovieSearchParam = "doubanId"
	MOVIESEARCHPARAM_YEAR MovieSearchParam = "year"
)

// All allowed values of MovieSearchParam enum
var AllowedMovieSearchParamEnumValues = []MovieSearchParam{
	"q",
	"imdbId",
	"tmdbId",
	"imdbTitle",
	"imdbYear",
	"traktId",
	"genre",
	"doubanId",
	"year",
}

func (v *MovieSearchParam) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MovieSearchParam(value)
	for _, existing := range AllowedMovieSearchParamEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MovieSearchParam", value)
}

// NewMovieSearchParamFromValue returns a pointer to a valid MovieSearchParam
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMovieSearchParamFromValue(v string) (*MovieSearchParam, error) {
	ev := MovieSearchParam(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MovieSearchParam: valid values are %v", v, AllowedMovieSearchParamEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MovieSearchParam) IsValid() bool {
	for _, existing := range AllowedMovieSearchParamEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MovieSearchParam value
func (v MovieSearchParam) Ptr() *MovieSearchParam {
	return &v
}

type NullableMovieSearchParam struct {
	value *MovieSearchParam
	isSet bool
}

func (v NullableMovieSearchParam) Get() *MovieSearchParam {
	return v.value
}

func (v *NullableMovieSearchParam) Set(val *MovieSearchParam) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieSearchParam) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieSearchParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieSearchParam(val *MovieSearchParam) *NullableMovieSearchParam {
	return &NullableMovieSearchParam{value: val, isSet: true}
}

func (v NullableMovieSearchParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieSearchParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

