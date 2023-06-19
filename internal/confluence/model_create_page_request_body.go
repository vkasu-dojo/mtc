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

// CreatePageRequestBody - struct for CreatePageRequestBody
type CreatePageRequestBody struct {
	PageBodyWrite       *PageBodyWrite
	PageNestedBodyWrite *PageNestedBodyWrite
}

// PageBodyWriteAsCreatePageRequestBody is a convenience function that returns PageBodyWrite wrapped in CreatePageRequestBody
func PageBodyWriteAsCreatePageRequestBody(v *PageBodyWrite) CreatePageRequestBody {
	return CreatePageRequestBody{
		PageBodyWrite: v,
	}
}

// PageNestedBodyWriteAsCreatePageRequestBody is a convenience function that returns PageNestedBodyWrite wrapped in CreatePageRequestBody
func PageNestedBodyWriteAsCreatePageRequestBody(v *PageNestedBodyWrite) CreatePageRequestBody {
	return CreatePageRequestBody{
		PageNestedBodyWrite: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreatePageRequestBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PageBodyWrite
	err = newStrictDecoder(data).Decode(&dst.PageBodyWrite)
	if err == nil {
		jsonPageBodyWrite, _ := json.Marshal(dst.PageBodyWrite)
		if string(jsonPageBodyWrite) == "{}" { // empty struct
			dst.PageBodyWrite = nil
		} else {
			match++
		}
	} else {
		dst.PageBodyWrite = nil
	}

	// try to unmarshal data into PageNestedBodyWrite
	err = newStrictDecoder(data).Decode(&dst.PageNestedBodyWrite)
	if err == nil {
		jsonPageNestedBodyWrite, _ := json.Marshal(dst.PageNestedBodyWrite)
		if string(jsonPageNestedBodyWrite) == "{}" { // empty struct
			dst.PageNestedBodyWrite = nil
		} else {
			match++
		}
	} else {
		dst.PageNestedBodyWrite = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PageBodyWrite = nil
		dst.PageNestedBodyWrite = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreatePageRequestBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreatePageRequestBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreatePageRequestBody) MarshalJSON() ([]byte, error) {
	if src.PageBodyWrite != nil {
		return json.Marshal(&src.PageBodyWrite)
	}

	if src.PageNestedBodyWrite != nil {
		return json.Marshal(&src.PageNestedBodyWrite)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreatePageRequestBody) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PageBodyWrite != nil {
		return obj.PageBodyWrite
	}

	if obj.PageNestedBodyWrite != nil {
		return obj.PageNestedBodyWrite
	}

	// all schemas are nil
	return nil
}

type NullableCreatePageRequestBody struct {
	value *CreatePageRequestBody
	isSet bool
}

func (v NullableCreatePageRequestBody) Get() *CreatePageRequestBody {
	return v.value
}

func (v *NullableCreatePageRequestBody) Set(val *CreatePageRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePageRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePageRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePageRequestBody(val *CreatePageRequestBody) *NullableCreatePageRequestBody {
	return &NullableCreatePageRequestBody{value: val, isSet: true}
}

func (v NullableCreatePageRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePageRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
