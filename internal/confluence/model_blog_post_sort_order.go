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

// BlogPostSortOrder The sort fields for blog posts. The default sort direction is ascending. To sort in descending order, append a `-` character before the sort field. For example, `fieldName` or `-fieldName`.
type BlogPostSortOrder string

// List of BlogPostSortOrder
const (
	BLOGPOSTSORTORDER_ID_ASC             BlogPostSortOrder = "id"
	BLOGPOSTSORTORDER_ID_DESC            BlogPostSortOrder = "-id"
	BLOGPOSTSORTORDER_CREATED_DATE_ASC   BlogPostSortOrder = "created-date"
	BLOGPOSTSORTORDER_CREATED_DATE_DESC  BlogPostSortOrder = "-created-date"
	BLOGPOSTSORTORDER_MODIFIED_DATE_ASC  BlogPostSortOrder = "modified-date"
	BLOGPOSTSORTORDER_MODIFIED_DATE_DESC BlogPostSortOrder = "-modified-date"
)

// All allowed values of BlogPostSortOrder enum
var AllowedBlogPostSortOrderEnumValues = []BlogPostSortOrder{
	"id",
	"-id",
	"created-date",
	"-created-date",
	"modified-date",
	"-modified-date",
}

func (v *BlogPostSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BlogPostSortOrder(value)
	for _, existing := range AllowedBlogPostSortOrderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BlogPostSortOrder", value)
}

// NewBlogPostSortOrderFromValue returns a pointer to a valid BlogPostSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBlogPostSortOrderFromValue(v string) (*BlogPostSortOrder, error) {
	ev := BlogPostSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BlogPostSortOrder: valid values are %v", v, AllowedBlogPostSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BlogPostSortOrder) IsValid() bool {
	for _, existing := range AllowedBlogPostSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BlogPostSortOrder value
func (v BlogPostSortOrder) Ptr() *BlogPostSortOrder {
	return &v
}

type NullableBlogPostSortOrder struct {
	value *BlogPostSortOrder
	isSet bool
}

func (v NullableBlogPostSortOrder) Get() *BlogPostSortOrder {
	return v.value
}

func (v *NullableBlogPostSortOrder) Set(val *BlogPostSortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableBlogPostSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableBlogPostSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlogPostSortOrder(val *BlogPostSortOrder) *NullableBlogPostSortOrder {
	return &NullableBlogPostSortOrder{value: val, isSet: true}
}

func (v NullableBlogPostSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlogPostSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
