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

// checks if the MultiEntityResultAttachment1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiEntityResultAttachment1{}

// MultiEntityResultAttachment1 struct for MultiEntityResultAttachment1
type MultiEntityResultAttachment1 struct {
	Results []Attachment `json:"results,omitempty"`
}

// NewMultiEntityResultAttachment1 instantiates a new MultiEntityResultAttachment1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiEntityResultAttachment1() *MultiEntityResultAttachment1 {
	this := MultiEntityResultAttachment1{}
	return &this
}

// NewMultiEntityResultAttachment1WithDefaults instantiates a new MultiEntityResultAttachment1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiEntityResultAttachment1WithDefaults() *MultiEntityResultAttachment1 {
	this := MultiEntityResultAttachment1{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MultiEntityResultAttachment1) GetResults() []Attachment {
	if o == nil || IsNil(o.Results) {
		var ret []Attachment
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiEntityResultAttachment1) GetResultsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MultiEntityResultAttachment1) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Attachment and assigns it to the Results field.
func (o *MultiEntityResultAttachment1) SetResults(v []Attachment) {
	o.Results = v
}

func (o MultiEntityResultAttachment1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiEntityResultAttachment1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableMultiEntityResultAttachment1 struct {
	value *MultiEntityResultAttachment1
	isSet bool
}

func (v NullableMultiEntityResultAttachment1) Get() *MultiEntityResultAttachment1 {
	return v.value
}

func (v *NullableMultiEntityResultAttachment1) Set(val *MultiEntityResultAttachment1) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiEntityResultAttachment1) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiEntityResultAttachment1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiEntityResultAttachment1(val *MultiEntityResultAttachment1) *NullableMultiEntityResultAttachment1 {
	return &NullableMultiEntityResultAttachment1{value: val, isSet: true}
}

func (v NullableMultiEntityResultAttachment1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiEntityResultAttachment1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
