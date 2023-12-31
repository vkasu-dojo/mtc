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

// checks if the InviteByEmail200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteByEmail200Response{}

// InviteByEmail200Response struct for InviteByEmail200Response
type InviteByEmail200Response struct {
	// List of emails invited to site.
	EmailsInvited    []string                                  `json:"emailsInvited,omitempty"`
	EmailsNotInvited *InviteByEmail200ResponseEmailsNotInvited `json:"emailsNotInvited,omitempty"`
}

// NewInviteByEmail200Response instantiates a new InviteByEmail200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteByEmail200Response() *InviteByEmail200Response {
	this := InviteByEmail200Response{}
	return &this
}

// NewInviteByEmail200ResponseWithDefaults instantiates a new InviteByEmail200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteByEmail200ResponseWithDefaults() *InviteByEmail200Response {
	this := InviteByEmail200Response{}
	return &this
}

// GetEmailsInvited returns the EmailsInvited field value if set, zero value otherwise.
func (o *InviteByEmail200Response) GetEmailsInvited() []string {
	if o == nil || IsNil(o.EmailsInvited) {
		var ret []string
		return ret
	}
	return o.EmailsInvited
}

// GetEmailsInvitedOk returns a tuple with the EmailsInvited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteByEmail200Response) GetEmailsInvitedOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailsInvited) {
		return nil, false
	}
	return o.EmailsInvited, true
}

// HasEmailsInvited returns a boolean if a field has been set.
func (o *InviteByEmail200Response) HasEmailsInvited() bool {
	if o != nil && !IsNil(o.EmailsInvited) {
		return true
	}

	return false
}

// SetEmailsInvited gets a reference to the given []string and assigns it to the EmailsInvited field.
func (o *InviteByEmail200Response) SetEmailsInvited(v []string) {
	o.EmailsInvited = v
}

// GetEmailsNotInvited returns the EmailsNotInvited field value if set, zero value otherwise.
func (o *InviteByEmail200Response) GetEmailsNotInvited() InviteByEmail200ResponseEmailsNotInvited {
	if o == nil || IsNil(o.EmailsNotInvited) {
		var ret InviteByEmail200ResponseEmailsNotInvited
		return ret
	}
	return *o.EmailsNotInvited
}

// GetEmailsNotInvitedOk returns a tuple with the EmailsNotInvited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteByEmail200Response) GetEmailsNotInvitedOk() (*InviteByEmail200ResponseEmailsNotInvited, bool) {
	if o == nil || IsNil(o.EmailsNotInvited) {
		return nil, false
	}
	return o.EmailsNotInvited, true
}

// HasEmailsNotInvited returns a boolean if a field has been set.
func (o *InviteByEmail200Response) HasEmailsNotInvited() bool {
	if o != nil && !IsNil(o.EmailsNotInvited) {
		return true
	}

	return false
}

// SetEmailsNotInvited gets a reference to the given InviteByEmail200ResponseEmailsNotInvited and assigns it to the EmailsNotInvited field.
func (o *InviteByEmail200Response) SetEmailsNotInvited(v InviteByEmail200ResponseEmailsNotInvited) {
	o.EmailsNotInvited = &v
}

func (o InviteByEmail200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteByEmail200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailsInvited) {
		toSerialize["emailsInvited"] = o.EmailsInvited
	}
	if !IsNil(o.EmailsNotInvited) {
		toSerialize["emailsNotInvited"] = o.EmailsNotInvited
	}
	return toSerialize, nil
}

type NullableInviteByEmail200Response struct {
	value *InviteByEmail200Response
	isSet bool
}

func (v NullableInviteByEmail200Response) Get() *InviteByEmail200Response {
	return v.value
}

func (v *NullableInviteByEmail200Response) Set(val *InviteByEmail200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteByEmail200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteByEmail200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteByEmail200Response(val *InviteByEmail200Response) *NullableInviteByEmail200Response {
	return &NullableInviteByEmail200Response{value: val, isSet: true}
}

func (v NullableInviteByEmail200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteByEmail200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
