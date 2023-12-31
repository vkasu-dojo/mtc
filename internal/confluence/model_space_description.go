/*
The Confluence Cloud REST API v2

This document describes Confluence's v2 APIs. This is intended to be an iteration on the existing Confluence Cloud REST API with improvements in both endpoint definitions and performance.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confluence

import (
	"encoding/json"
)

// checks if the SpaceDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceDescription{}

// SpaceDescription Contains fields for each representation type requested.
type SpaceDescription struct {
	Plain *BodyType `json:"plain,omitempty"`
	View  *BodyType `json:"view,omitempty"`
}

// NewSpaceDescription instantiates a new SpaceDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceDescription() *SpaceDescription {
	this := SpaceDescription{}
	return &this
}

// NewSpaceDescriptionWithDefaults instantiates a new SpaceDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceDescriptionWithDefaults() *SpaceDescription {
	this := SpaceDescription{}
	return &this
}

// GetPlain returns the Plain field value if set, zero value otherwise.
func (o *SpaceDescription) GetPlain() BodyType {
	if o == nil || IsNil(o.Plain) {
		var ret BodyType
		return ret
	}
	return *o.Plain
}

// GetPlainOk returns a tuple with the Plain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceDescription) GetPlainOk() (*BodyType, bool) {
	if o == nil || IsNil(o.Plain) {
		return nil, false
	}
	return o.Plain, true
}

// HasPlain returns a boolean if a field has been set.
func (o *SpaceDescription) HasPlain() bool {
	if o != nil && !IsNil(o.Plain) {
		return true
	}

	return false
}

// SetPlain gets a reference to the given BodyType and assigns it to the Plain field.
func (o *SpaceDescription) SetPlain(v BodyType) {
	o.Plain = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *SpaceDescription) GetView() BodyType {
	if o == nil || IsNil(o.View) {
		var ret BodyType
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceDescription) GetViewOk() (*BodyType, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *SpaceDescription) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given BodyType and assigns it to the View field.
func (o *SpaceDescription) SetView(v BodyType) {
	o.View = &v
}

func (o SpaceDescription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plain) {
		toSerialize["plain"] = o.Plain
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	return toSerialize, nil
}

type NullableSpaceDescription struct {
	value *SpaceDescription
	isSet bool
}

func (v NullableSpaceDescription) Get() *SpaceDescription {
	return v.value
}

func (v *NullableSpaceDescription) Set(val *SpaceDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceDescription(val *SpaceDescription) *NullableSpaceDescription {
	return &NullableSpaceDescription{value: val, isSet: true}
}

func (v NullableSpaceDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
