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

// checks if the MultiEntityResultSpaceProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiEntityResultSpaceProperty{}

// MultiEntityResultSpaceProperty struct for MultiEntityResultSpaceProperty
type MultiEntityResultSpaceProperty struct {
	Results []SpaceProperty `json:"results,omitempty"`
}

// NewMultiEntityResultSpaceProperty instantiates a new MultiEntityResultSpaceProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiEntityResultSpaceProperty() *MultiEntityResultSpaceProperty {
	this := MultiEntityResultSpaceProperty{}
	return &this
}

// NewMultiEntityResultSpacePropertyWithDefaults instantiates a new MultiEntityResultSpaceProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiEntityResultSpacePropertyWithDefaults() *MultiEntityResultSpaceProperty {
	this := MultiEntityResultSpaceProperty{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MultiEntityResultSpaceProperty) GetResults() []SpaceProperty {
	if o == nil || IsNil(o.Results) {
		var ret []SpaceProperty
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiEntityResultSpaceProperty) GetResultsOk() ([]SpaceProperty, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MultiEntityResultSpaceProperty) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SpaceProperty and assigns it to the Results field.
func (o *MultiEntityResultSpaceProperty) SetResults(v []SpaceProperty) {
	o.Results = v
}

func (o MultiEntityResultSpaceProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiEntityResultSpaceProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableMultiEntityResultSpaceProperty struct {
	value *MultiEntityResultSpaceProperty
	isSet bool
}

func (v NullableMultiEntityResultSpaceProperty) Get() *MultiEntityResultSpaceProperty {
	return v.value
}

func (v *NullableMultiEntityResultSpaceProperty) Set(val *MultiEntityResultSpaceProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiEntityResultSpaceProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiEntityResultSpaceProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiEntityResultSpaceProperty(val *MultiEntityResultSpaceProperty) *NullableMultiEntityResultSpaceProperty {
	return &NullableMultiEntityResultSpaceProperty{value: val, isSet: true}
}

func (v NullableMultiEntityResultSpaceProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiEntityResultSpaceProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
