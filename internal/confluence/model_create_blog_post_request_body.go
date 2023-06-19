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

// CreateBlogPostRequestBody - struct for CreateBlogPostRequestBody
type CreateBlogPostRequestBody struct {
	BlogPostBodyWrite       *BlogPostBodyWrite
	BlogPostNestedBodyWrite *BlogPostNestedBodyWrite
}

// BlogPostBodyWriteAsCreateBlogPostRequestBody is a convenience function that returns BlogPostBodyWrite wrapped in CreateBlogPostRequestBody
func BlogPostBodyWriteAsCreateBlogPostRequestBody(v *BlogPostBodyWrite) CreateBlogPostRequestBody {
	return CreateBlogPostRequestBody{
		BlogPostBodyWrite: v,
	}
}

// BlogPostNestedBodyWriteAsCreateBlogPostRequestBody is a convenience function that returns BlogPostNestedBodyWrite wrapped in CreateBlogPostRequestBody
func BlogPostNestedBodyWriteAsCreateBlogPostRequestBody(v *BlogPostNestedBodyWrite) CreateBlogPostRequestBody {
	return CreateBlogPostRequestBody{
		BlogPostNestedBodyWrite: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateBlogPostRequestBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlogPostBodyWrite
	err = newStrictDecoder(data).Decode(&dst.BlogPostBodyWrite)
	if err == nil {
		jsonBlogPostBodyWrite, _ := json.Marshal(dst.BlogPostBodyWrite)
		if string(jsonBlogPostBodyWrite) == "{}" { // empty struct
			dst.BlogPostBodyWrite = nil
		} else {
			match++
		}
	} else {
		dst.BlogPostBodyWrite = nil
	}

	// try to unmarshal data into BlogPostNestedBodyWrite
	err = newStrictDecoder(data).Decode(&dst.BlogPostNestedBodyWrite)
	if err == nil {
		jsonBlogPostNestedBodyWrite, _ := json.Marshal(dst.BlogPostNestedBodyWrite)
		if string(jsonBlogPostNestedBodyWrite) == "{}" { // empty struct
			dst.BlogPostNestedBodyWrite = nil
		} else {
			match++
		}
	} else {
		dst.BlogPostNestedBodyWrite = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlogPostBodyWrite = nil
		dst.BlogPostNestedBodyWrite = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateBlogPostRequestBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateBlogPostRequestBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateBlogPostRequestBody) MarshalJSON() ([]byte, error) {
	if src.BlogPostBodyWrite != nil {
		return json.Marshal(&src.BlogPostBodyWrite)
	}

	if src.BlogPostNestedBodyWrite != nil {
		return json.Marshal(&src.BlogPostNestedBodyWrite)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateBlogPostRequestBody) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BlogPostBodyWrite != nil {
		return obj.BlogPostBodyWrite
	}

	if obj.BlogPostNestedBodyWrite != nil {
		return obj.BlogPostNestedBodyWrite
	}

	// all schemas are nil
	return nil
}

type NullableCreateBlogPostRequestBody struct {
	value *CreateBlogPostRequestBody
	isSet bool
}

func (v NullableCreateBlogPostRequestBody) Get() *CreateBlogPostRequestBody {
	return v.value
}

func (v *NullableCreateBlogPostRequestBody) Set(val *CreateBlogPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBlogPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBlogPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBlogPostRequestBody(val *CreateBlogPostRequestBody) *NullableCreateBlogPostRequestBody {
	return &NullableCreateBlogPostRequestBody{value: val, isSet: true}
}

func (v NullableCreateBlogPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBlogPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}