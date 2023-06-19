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

// checks if the ContentProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentProperty{}

// ContentProperty struct for ContentProperty
type ContentProperty struct {
	Id *ContentPropertyId `json:"id,omitempty"`
	// Key of the property
	Key *string `json:"key,omitempty"`
	// Value of the property. Must be a valid JSON value.
	Value   interface{} `json:"value,omitempty"`
	Version *Version    `json:"version,omitempty"`
}

// NewContentProperty instantiates a new ContentProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentProperty() *ContentProperty {
	this := ContentProperty{}
	return &this
}

// NewContentPropertyWithDefaults instantiates a new ContentProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentPropertyWithDefaults() *ContentProperty {
	this := ContentProperty{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContentProperty) GetId() ContentPropertyId {
	if o == nil || IsNil(o.Id) {
		var ret ContentPropertyId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentProperty) GetIdOk() (*ContentPropertyId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContentProperty) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given ContentPropertyId and assigns it to the Id field.
func (o *ContentProperty) SetId(v ContentPropertyId) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ContentProperty) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentProperty) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ContentProperty) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ContentProperty) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentProperty) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentProperty) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ContentProperty) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ContentProperty) SetValue(v interface{}) {
	o.Value = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ContentProperty) GetVersion() Version {
	if o == nil || IsNil(o.Version) {
		var ret Version
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentProperty) GetVersionOk() (*Version, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ContentProperty) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given Version and assigns it to the Version field.
func (o *ContentProperty) SetVersion(v Version) {
	o.Version = &v
}

func (o ContentProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableContentProperty struct {
	value *ContentProperty
	isSet bool
}

func (v NullableContentProperty) Get() *ContentProperty {
	return v.value
}

func (v *NullableContentProperty) Set(val *ContentProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableContentProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableContentProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentProperty(val *ContentProperty) *NullableContentProperty {
	return &NullableContentProperty{value: val, isSet: true}
}

func (v NullableContentProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
