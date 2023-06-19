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

// checks if the MultiEntityResultTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiEntityResultTask{}

// MultiEntityResultTask struct for MultiEntityResultTask
type MultiEntityResultTask struct {
	Results []Task                       `json:"results,omitempty"`
	Links   *MultiEntityResultLabelLinks `json:"_links,omitempty"`
}

// NewMultiEntityResultTask instantiates a new MultiEntityResultTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiEntityResultTask() *MultiEntityResultTask {
	this := MultiEntityResultTask{}
	return &this
}

// NewMultiEntityResultTaskWithDefaults instantiates a new MultiEntityResultTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiEntityResultTaskWithDefaults() *MultiEntityResultTask {
	this := MultiEntityResultTask{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MultiEntityResultTask) GetResults() []Task {
	if o == nil || IsNil(o.Results) {
		var ret []Task
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiEntityResultTask) GetResultsOk() ([]Task, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MultiEntityResultTask) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Task and assigns it to the Results field.
func (o *MultiEntityResultTask) SetResults(v []Task) {
	o.Results = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MultiEntityResultTask) GetLinks() MultiEntityResultLabelLinks {
	if o == nil || IsNil(o.Links) {
		var ret MultiEntityResultLabelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiEntityResultTask) GetLinksOk() (*MultiEntityResultLabelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MultiEntityResultTask) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given MultiEntityResultLabelLinks and assigns it to the Links field.
func (o *MultiEntityResultTask) SetLinks(v MultiEntityResultLabelLinks) {
	o.Links = &v
}

func (o MultiEntityResultTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiEntityResultTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableMultiEntityResultTask struct {
	value *MultiEntityResultTask
	isSet bool
}

func (v NullableMultiEntityResultTask) Get() *MultiEntityResultTask {
	return v.value
}

func (v *NullableMultiEntityResultTask) Set(val *MultiEntityResultTask) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiEntityResultTask) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiEntityResultTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiEntityResultTask(val *MultiEntityResultTask) *NullableMultiEntityResultTask {
	return &NullableMultiEntityResultTask{value: val, isSet: true}
}

func (v NullableMultiEntityResultTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiEntityResultTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}