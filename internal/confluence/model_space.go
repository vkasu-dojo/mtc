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

// checks if the Space type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Space{}

// Space struct for Space
type Space struct {
	Id *SpaceId `json:"id,omitempty"`
	// Key of the space.
	Key *string `json:"key,omitempty"`
	// Name of the space.
	Name        *string           `json:"name,omitempty"`
	Type        *SpaceType        `json:"type,omitempty"`
	Status      *SpaceStatus      `json:"status,omitempty"`
	HomepageId  *SpaceHomepageId  `json:"homepageId,omitempty"`
	Description *SpaceDescription `json:"description,omitempty"`
}

// NewSpace instantiates a new Space object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpace() *Space {
	this := Space{}
	return &this
}

// NewSpaceWithDefaults instantiates a new Space object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceWithDefaults() *Space {
	this := Space{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Space) GetId() SpaceId {
	if o == nil || IsNil(o.Id) {
		var ret SpaceId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetIdOk() (*SpaceId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Space) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given SpaceId and assigns it to the Id field.
func (o *Space) SetId(v SpaceId) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Space) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Space) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Space) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Space) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Space) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Space) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Space) GetType() SpaceType {
	if o == nil || IsNil(o.Type) {
		var ret SpaceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetTypeOk() (*SpaceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Space) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SpaceType and assigns it to the Type field.
func (o *Space) SetType(v SpaceType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Space) GetStatus() SpaceStatus {
	if o == nil || IsNil(o.Status) {
		var ret SpaceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetStatusOk() (*SpaceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Space) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SpaceStatus and assigns it to the Status field.
func (o *Space) SetStatus(v SpaceStatus) {
	o.Status = &v
}

// GetHomepageId returns the HomepageId field value if set, zero value otherwise.
func (o *Space) GetHomepageId() SpaceHomepageId {
	if o == nil || IsNil(o.HomepageId) {
		var ret SpaceHomepageId
		return ret
	}
	return *o.HomepageId
}

// GetHomepageIdOk returns a tuple with the HomepageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetHomepageIdOk() (*SpaceHomepageId, bool) {
	if o == nil || IsNil(o.HomepageId) {
		return nil, false
	}
	return o.HomepageId, true
}

// HasHomepageId returns a boolean if a field has been set.
func (o *Space) HasHomepageId() bool {
	if o != nil && !IsNil(o.HomepageId) {
		return true
	}

	return false
}

// SetHomepageId gets a reference to the given SpaceHomepageId and assigns it to the HomepageId field.
func (o *Space) SetHomepageId(v SpaceHomepageId) {
	o.HomepageId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Space) GetDescription() SpaceDescription {
	if o == nil || IsNil(o.Description) {
		var ret SpaceDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetDescriptionOk() (*SpaceDescription, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Space) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given SpaceDescription and assigns it to the Description field.
func (o *Space) SetDescription(v SpaceDescription) {
	o.Description = &v
}

func (o Space) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Space) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.HomepageId) {
		toSerialize["homepageId"] = o.HomepageId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSpace struct {
	value *Space
	isSet bool
}

func (v NullableSpace) Get() *Space {
	return v.value
}

func (v *NullableSpace) Set(val *Space) {
	v.value = val
	v.isSet = true
}

func (v NullableSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpace(val *Space) *NullableSpace {
	return &NullableSpace{value: val, isSet: true}
}

func (v NullableSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}