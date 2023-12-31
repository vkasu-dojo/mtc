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

// checks if the CreateInlineCommentModelInlineCommentProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInlineCommentModelInlineCommentProperties{}

// CreateInlineCommentModelInlineCommentProperties Object describing the text to highlight on the page/blog post. Only applicable for top level inline comments (not replies) and required in that case.
type CreateInlineCommentModelInlineCommentProperties struct {
	// The text to highlight
	TextSelection *string `json:"textSelection,omitempty"`
	// The number of matches for the selected text on the page (should be strictly greater than textSelectionMatchIndex)
	TextSelectionMatchCount *int32 `json:"textSelectionMatchCount,omitempty"`
	// The match index to highlight. This is zero-based. E.g. if you have 3 occurrences of \"hello world\" on a page  and you want to highlight the second occurrence, you should pass 1 for textSelectionMatchIndex and 3 for textSelectionMatchCount.
	TextSelectionMatchIndex *int32 `json:"textSelectionMatchIndex,omitempty"`
}

// NewCreateInlineCommentModelInlineCommentProperties instantiates a new CreateInlineCommentModelInlineCommentProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInlineCommentModelInlineCommentProperties() *CreateInlineCommentModelInlineCommentProperties {
	this := CreateInlineCommentModelInlineCommentProperties{}
	return &this
}

// NewCreateInlineCommentModelInlineCommentPropertiesWithDefaults instantiates a new CreateInlineCommentModelInlineCommentProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInlineCommentModelInlineCommentPropertiesWithDefaults() *CreateInlineCommentModelInlineCommentProperties {
	this := CreateInlineCommentModelInlineCommentProperties{}
	return &this
}

// GetTextSelection returns the TextSelection field value if set, zero value otherwise.
func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelection() string {
	if o == nil || IsNil(o.TextSelection) {
		var ret string
		return ret
	}
	return *o.TextSelection
}

// GetTextSelectionOk returns a tuple with the TextSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionOk() (*string, bool) {
	if o == nil || IsNil(o.TextSelection) {
		return nil, false
	}
	return o.TextSelection, true
}

// HasTextSelection returns a boolean if a field has been set.
func (o *CreateInlineCommentModelInlineCommentProperties) HasTextSelection() bool {
	if o != nil && !IsNil(o.TextSelection) {
		return true
	}

	return false
}

// SetTextSelection gets a reference to the given string and assigns it to the TextSelection field.
func (o *CreateInlineCommentModelInlineCommentProperties) SetTextSelection(v string) {
	o.TextSelection = &v
}

// GetTextSelectionMatchCount returns the TextSelectionMatchCount field value if set, zero value otherwise.
func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionMatchCount() int32 {
	if o == nil || IsNil(o.TextSelectionMatchCount) {
		var ret int32
		return ret
	}
	return *o.TextSelectionMatchCount
}

// GetTextSelectionMatchCountOk returns a tuple with the TextSelectionMatchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionMatchCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TextSelectionMatchCount) {
		return nil, false
	}
	return o.TextSelectionMatchCount, true
}

// HasTextSelectionMatchCount returns a boolean if a field has been set.
func (o *CreateInlineCommentModelInlineCommentProperties) HasTextSelectionMatchCount() bool {
	if o != nil && !IsNil(o.TextSelectionMatchCount) {
		return true
	}

	return false
}

// SetTextSelectionMatchCount gets a reference to the given int32 and assigns it to the TextSelectionMatchCount field.
func (o *CreateInlineCommentModelInlineCommentProperties) SetTextSelectionMatchCount(v int32) {
	o.TextSelectionMatchCount = &v
}

// GetTextSelectionMatchIndex returns the TextSelectionMatchIndex field value if set, zero value otherwise.
func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionMatchIndex() int32 {
	if o == nil || IsNil(o.TextSelectionMatchIndex) {
		var ret int32
		return ret
	}
	return *o.TextSelectionMatchIndex
}

// GetTextSelectionMatchIndexOk returns a tuple with the TextSelectionMatchIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInlineCommentModelInlineCommentProperties) GetTextSelectionMatchIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.TextSelectionMatchIndex) {
		return nil, false
	}
	return o.TextSelectionMatchIndex, true
}

// HasTextSelectionMatchIndex returns a boolean if a field has been set.
func (o *CreateInlineCommentModelInlineCommentProperties) HasTextSelectionMatchIndex() bool {
	if o != nil && !IsNil(o.TextSelectionMatchIndex) {
		return true
	}

	return false
}

// SetTextSelectionMatchIndex gets a reference to the given int32 and assigns it to the TextSelectionMatchIndex field.
func (o *CreateInlineCommentModelInlineCommentProperties) SetTextSelectionMatchIndex(v int32) {
	o.TextSelectionMatchIndex = &v
}

func (o CreateInlineCommentModelInlineCommentProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInlineCommentModelInlineCommentProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TextSelection) {
		toSerialize["textSelection"] = o.TextSelection
	}
	if !IsNil(o.TextSelectionMatchCount) {
		toSerialize["textSelectionMatchCount"] = o.TextSelectionMatchCount
	}
	if !IsNil(o.TextSelectionMatchIndex) {
		toSerialize["textSelectionMatchIndex"] = o.TextSelectionMatchIndex
	}
	return toSerialize, nil
}

type NullableCreateInlineCommentModelInlineCommentProperties struct {
	value *CreateInlineCommentModelInlineCommentProperties
	isSet bool
}

func (v NullableCreateInlineCommentModelInlineCommentProperties) Get() *CreateInlineCommentModelInlineCommentProperties {
	return v.value
}

func (v *NullableCreateInlineCommentModelInlineCommentProperties) Set(val *CreateInlineCommentModelInlineCommentProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInlineCommentModelInlineCommentProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInlineCommentModelInlineCommentProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInlineCommentModelInlineCommentProperties(val *CreateInlineCommentModelInlineCommentProperties) *NullableCreateInlineCommentModelInlineCommentProperties {
	return &NullableCreateInlineCommentModelInlineCommentProperties{value: val, isSet: true}
}

func (v NullableCreateInlineCommentModelInlineCommentProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInlineCommentModelInlineCommentProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
