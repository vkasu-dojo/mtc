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

// SpaceHomepageId - ID of the space's homepage.  Due to JavaScript's max integer representation of 2^53-1, the type of this field will be changed from a numeric type to a string type at the end of the deprecation period. In the meantime, `serialize-ids-as-strings=true` can be passed as a query param to any v2 endpoint to opt-in to this change now. See this [changelog](https://developer.atlassian.com/cloud/confluence/changelog/#CHANGE-905) for more detail.
type SpaceHomepageId struct {
	Int64  *int64
	String *string
}

// int64AsSpaceHomepageId is a convenience function that returns int64 wrapped in SpaceHomepageId
func Int64AsSpaceHomepageId(v *int64) SpaceHomepageId {
	return SpaceHomepageId{
		Int64: v,
	}
}

// stringAsSpaceHomepageId is a convenience function that returns string wrapped in SpaceHomepageId
func StringAsSpaceHomepageId(v *string) SpaceHomepageId {
	return SpaceHomepageId{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SpaceHomepageId) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int64
	err = newStrictDecoder(data).Decode(&dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			match++
		}
	} else {
		dst.Int64 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SpaceHomepageId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SpaceHomepageId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SpaceHomepageId) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SpaceHomepageId) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableSpaceHomepageId struct {
	value *SpaceHomepageId
	isSet bool
}

func (v NullableSpaceHomepageId) Get() *SpaceHomepageId {
	return v.value
}

func (v *NullableSpaceHomepageId) Set(val *SpaceHomepageId) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceHomepageId) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceHomepageId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceHomepageId(val *SpaceHomepageId) *NullableSpaceHomepageId {
	return &NullableSpaceHomepageId{value: val, isSet: true}
}

func (v NullableSpaceHomepageId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceHomepageId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}