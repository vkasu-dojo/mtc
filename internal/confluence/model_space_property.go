/*
The Confluence Cloud REST API v2

This document describes Confluence's v2 APIs. This is intended to be an iteration on the existing Confluence Cloud REST API with improvements in both endpoint definitions and performance.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confluence

import (
	"encoding/json"
	"time"
)

// checks if the SpaceProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceProperty{}

// SpaceProperty struct for SpaceProperty
type SpaceProperty struct {
	Id *SpacePropertyId `json:"id,omitempty"`
	// Key of the space property.
	Key *string `json:"key,omitempty"`
	// Value of the space property.
	Value interface{} `json:"value,omitempty"`
	// RFC3339 compliant date time at which the property was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Atlassian account ID of the user that created the space property.
	CreatedBy *string               `json:"createdBy,omitempty"`
	Version   *SpacePropertyVersion `json:"version,omitempty"`
}

// NewSpaceProperty instantiates a new SpaceProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceProperty() *SpaceProperty {
	this := SpaceProperty{}
	return &this
}

// NewSpacePropertyWithDefaults instantiates a new SpaceProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpacePropertyWithDefaults() *SpaceProperty {
	this := SpaceProperty{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpaceProperty) GetId() SpacePropertyId {
	if o == nil || IsNil(o.Id) {
		var ret SpacePropertyId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceProperty) GetIdOk() (*SpacePropertyId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpaceProperty) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given SpacePropertyId and assigns it to the Id field.
func (o *SpaceProperty) SetId(v SpacePropertyId) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SpaceProperty) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceProperty) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SpaceProperty) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SpaceProperty) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpaceProperty) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpaceProperty) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SpaceProperty) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *SpaceProperty) SetValue(v interface{}) {
	o.Value = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SpaceProperty) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceProperty) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SpaceProperty) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SpaceProperty) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SpaceProperty) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceProperty) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SpaceProperty) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SpaceProperty) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SpaceProperty) GetVersion() SpacePropertyVersion {
	if o == nil || IsNil(o.Version) {
		var ret SpacePropertyVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceProperty) GetVersionOk() (*SpacePropertyVersion, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SpaceProperty) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given SpacePropertyVersion and assigns it to the Version field.
func (o *SpaceProperty) SetVersion(v SpacePropertyVersion) {
	o.Version = &v
}

func (o SpaceProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceProperty) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableSpaceProperty struct {
	value *SpaceProperty
	isSet bool
}

func (v NullableSpaceProperty) Get() *SpaceProperty {
	return v.value
}

func (v *NullableSpaceProperty) Set(val *SpaceProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceProperty(val *SpaceProperty) *NullableSpaceProperty {
	return &NullableSpaceProperty{value: val, isSet: true}
}

func (v NullableSpaceProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
