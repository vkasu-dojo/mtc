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

// OnlyArchivedAndCurrentContentStatus The status of the content.
type OnlyArchivedAndCurrentContentStatus string

// List of OnlyArchivedAndCurrentContentStatus
const (
	ONLYARCHIVEDANDCURRENTCONTENTSTATUS_CURRENT  OnlyArchivedAndCurrentContentStatus = "current"
	ONLYARCHIVEDANDCURRENTCONTENTSTATUS_ARCHIVED OnlyArchivedAndCurrentContentStatus = "archived"
)

// All allowed values of OnlyArchivedAndCurrentContentStatus enum
var AllowedOnlyArchivedAndCurrentContentStatusEnumValues = []OnlyArchivedAndCurrentContentStatus{
	"current",
	"archived",
}

func (v *OnlyArchivedAndCurrentContentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OnlyArchivedAndCurrentContentStatus(value)
	for _, existing := range AllowedOnlyArchivedAndCurrentContentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OnlyArchivedAndCurrentContentStatus", value)
}

// NewOnlyArchivedAndCurrentContentStatusFromValue returns a pointer to a valid OnlyArchivedAndCurrentContentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOnlyArchivedAndCurrentContentStatusFromValue(v string) (*OnlyArchivedAndCurrentContentStatus, error) {
	ev := OnlyArchivedAndCurrentContentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OnlyArchivedAndCurrentContentStatus: valid values are %v", v, AllowedOnlyArchivedAndCurrentContentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OnlyArchivedAndCurrentContentStatus) IsValid() bool {
	for _, existing := range AllowedOnlyArchivedAndCurrentContentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnlyArchivedAndCurrentContentStatus value
func (v OnlyArchivedAndCurrentContentStatus) Ptr() *OnlyArchivedAndCurrentContentStatus {
	return &v
}

type NullableOnlyArchivedAndCurrentContentStatus struct {
	value *OnlyArchivedAndCurrentContentStatus
	isSet bool
}

func (v NullableOnlyArchivedAndCurrentContentStatus) Get() *OnlyArchivedAndCurrentContentStatus {
	return v.value
}

func (v *NullableOnlyArchivedAndCurrentContentStatus) Set(val *OnlyArchivedAndCurrentContentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOnlyArchivedAndCurrentContentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOnlyArchivedAndCurrentContentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnlyArchivedAndCurrentContentStatus(val *OnlyArchivedAndCurrentContentStatus) *NullableOnlyArchivedAndCurrentContentStatus {
	return &NullableOnlyArchivedAndCurrentContentStatus{value: val, isSet: true}
}

func (v NullableOnlyArchivedAndCurrentContentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnlyArchivedAndCurrentContentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
