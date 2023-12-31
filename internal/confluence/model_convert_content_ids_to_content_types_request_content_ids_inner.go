/*
The Confluence Cloud REST API v2

This document describes Confluence's v2 APIs. This is intended to be an iteration on the existing Confluence Cloud REST API with improvements in both endpoint definitions and performance.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confluence

import (
	"encoding/json"
	"fmt"
)

// ConvertContentIdsToContentTypesRequestContentIdsInner struct for ConvertContentIdsToContentTypesRequestContentIdsInner
type ConvertContentIdsToContentTypesRequestContentIdsInner struct {
	float32 *float32
	string  *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ConvertContentIdsToContentTypesRequestContentIdsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into float32
	err = json.Unmarshal(data, &dst.float32)
	if err == nil {
		jsonfloat32, _ := json.Marshal(dst.float32)
		if string(jsonfloat32) == "{}" { // empty struct
			dst.float32 = nil
		} else {
			return nil // data stored in dst.float32, return on the first match
		}
	} else {
		dst.float32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ConvertContentIdsToContentTypesRequestContentIdsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ConvertContentIdsToContentTypesRequestContentIdsInner) MarshalJSON() ([]byte, error) {
	if src.float32 != nil {
		return json.Marshal(&src.float32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableConvertContentIdsToContentTypesRequestContentIdsInner struct {
	value *ConvertContentIdsToContentTypesRequestContentIdsInner
	isSet bool
}

func (v NullableConvertContentIdsToContentTypesRequestContentIdsInner) Get() *ConvertContentIdsToContentTypesRequestContentIdsInner {
	return v.value
}

func (v *NullableConvertContentIdsToContentTypesRequestContentIdsInner) Set(val *ConvertContentIdsToContentTypesRequestContentIdsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertContentIdsToContentTypesRequestContentIdsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertContentIdsToContentTypesRequestContentIdsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertContentIdsToContentTypesRequestContentIdsInner(val *ConvertContentIdsToContentTypesRequestContentIdsInner) *NullableConvertContentIdsToContentTypesRequestContentIdsInner {
	return &NullableConvertContentIdsToContentTypesRequestContentIdsInner{value: val, isSet: true}
}

func (v NullableConvertContentIdsToContentTypesRequestContentIdsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertContentIdsToContentTypesRequestContentIdsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
