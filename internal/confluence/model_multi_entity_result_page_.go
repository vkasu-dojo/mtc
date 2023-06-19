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

// checks if the MultiEntityResultPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiEntityResultPage{}

// MultiEntityResultPage struct for MultiEntityResultPage
type MultiEntityResultPage struct {
	Results []Page                       `json:"results,omitempty"`
	Links   *MultiEntityResultLabelLinks `json:"_links,omitempty"`
}

// NewMultiEntityResultPage instantiates a new MultiEntityResultPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiEntityResultPage() *MultiEntityResultPage {
	this := MultiEntityResultPage{}
	return &this
}

// NewMultiEntityResultPageWithDefaults instantiates a new MultiEntityResultPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiEntityResultPageWithDefaults() *MultiEntityResultPage {
	this := MultiEntityResultPage{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MultiEntityResultPage) GetResults() []Page {
	if o == nil || IsNil(o.Results) {
		var ret []Page
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiEntityResultPage) GetResultsOk() ([]Page, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MultiEntityResultPage) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Page and assigns it to the Results field.
func (o *MultiEntityResultPage) SetResults(v []Page) {
	o.Results = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MultiEntityResultPage) GetLinks() MultiEntityResultLabelLinks {
	if o == nil || IsNil(o.Links) {
		var ret MultiEntityResultLabelLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiEntityResultPage) GetLinksOk() (*MultiEntityResultLabelLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MultiEntityResultPage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given MultiEntityResultLabelLinks and assigns it to the Links field.
func (o *MultiEntityResultPage) SetLinks(v MultiEntityResultLabelLinks) {
	o.Links = &v
}

func (o MultiEntityResultPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiEntityResultPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableMultiEntityResultPage struct {
	value *MultiEntityResultPage
	isSet bool
}

func (v NullableMultiEntityResultPage) Get() *MultiEntityResultPage {
	return v.value
}

func (v *NullableMultiEntityResultPage) Set(val *MultiEntityResultPage) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiEntityResultPage) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiEntityResultPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiEntityResultPage(val *MultiEntityResultPage) *NullableMultiEntityResultPage {
	return &NullableMultiEntityResultPage{value: val, isSet: true}
}

func (v NullableMultiEntityResultPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiEntityResultPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
