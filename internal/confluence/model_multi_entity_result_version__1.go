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

// checks if the MultiEntityResultVersion1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiEntityResultVersion1{}

// MultiEntityResultVersion1 struct for MultiEntityResultVersion1
type MultiEntityResultVersion1 struct {
	Results []BlogPostVersion            `json:"results,omitempty"`
	Links   *MultiEntityResultLabelLinks `json:"_links,omitempty"`
}

// NewMultiEntityResultVersion1 instantiates a new MultiEntityResultVersion1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiEntityResultVersion1() *MultiEntityResultVersion1 {
	this := MultiEntityResultVersion1{}
	return &this
}

// NewMultiEntityResultVersion1WithDefaults instantiates a new MultiEntityResultVersion1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiEntityResultVersion1WithDefaults() *MultiEntityResultVersion1 {
	this := MultiEntityResultVersion1{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MultiEntityResultVersion1) GetResults() []BlogPostVersion {
	if o == nil || IsNil(o.Results) {
		var ret []BlogPostVersion
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiEntityResultVersion1) GetResultsOk() ([]BlogPostVersion, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MultiEntityResultVersion1) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BlogPostVersion and assigns it to the Results field.
func (o *MultiEntityResultVersion1) SetResults(v []BlogPostVersion) {
	o.Results = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MultiEntityResultVersion1) GetLinks() MultiEntityResultLabelLinks {
	if o == nil || IsNil(o.Links) {
		var ret MultiEntityResultLabelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiEntityResultVersion1) GetLinksOk() (*MultiEntityResultLabelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MultiEntityResultVersion1) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given MultiEntityResultLabelLinks and assigns it to the Links field.
func (o *MultiEntityResultVersion1) SetLinks(v MultiEntityResultLabelLinks) {
	o.Links = &v
}

func (o MultiEntityResultVersion1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiEntityResultVersion1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableMultiEntityResultVersion1 struct {
	value *MultiEntityResultVersion1
	isSet bool
}

func (v NullableMultiEntityResultVersion1) Get() *MultiEntityResultVersion1 {
	return v.value
}

func (v *NullableMultiEntityResultVersion1) Set(val *MultiEntityResultVersion1) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiEntityResultVersion1) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiEntityResultVersion1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiEntityResultVersion1(val *MultiEntityResultVersion1) *NullableMultiEntityResultVersion1 {
	return &NullableMultiEntityResultVersion1{value: val, isSet: true}
}

func (v NullableMultiEntityResultVersion1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiEntityResultVersion1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
