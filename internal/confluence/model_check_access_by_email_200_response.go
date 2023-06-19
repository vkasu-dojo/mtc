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

// checks if the CheckAccessByEmail200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckAccessByEmail200Response{}

// CheckAccessByEmail200Response struct for CheckAccessByEmail200Response
type CheckAccessByEmail200Response struct {
	// List of emails that do not have access to site.
	EmailsWithoutAccess []string `json:"emailsWithoutAccess,omitempty"`
	// List of invalid emails provided in the request.
	InvalidEmails []string `json:"invalidEmails,omitempty"`
}

// NewCheckAccessByEmail200Response instantiates a new CheckAccessByEmail200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAccessByEmail200Response() *CheckAccessByEmail200Response {
	this := CheckAccessByEmail200Response{}
	return &this
}

// NewCheckAccessByEmail200ResponseWithDefaults instantiates a new CheckAccessByEmail200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAccessByEmail200ResponseWithDefaults() *CheckAccessByEmail200Response {
	this := CheckAccessByEmail200Response{}
	return &this
}

// GetEmailsWithoutAccess returns the EmailsWithoutAccess field value if set, zero value otherwise.
func (o *CheckAccessByEmail200Response) GetEmailsWithoutAccess() []string {
	if o == nil || IsNil(o.EmailsWithoutAccess) {
		var ret []string
		return ret
	}
	return o.EmailsWithoutAccess
}

// GetEmailsWithoutAccessOk returns a tuple with the EmailsWithoutAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAccessByEmail200Response) GetEmailsWithoutAccessOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailsWithoutAccess) {
		return nil, false
	}
	return o.EmailsWithoutAccess, true
}

// HasEmailsWithoutAccess returns a boolean if a field has been set.
func (o *CheckAccessByEmail200Response) HasEmailsWithoutAccess() bool {
	if o != nil && !IsNil(o.EmailsWithoutAccess) {
		return true
	}

	return false
}

// SetEmailsWithoutAccess gets a reference to the given []string and assigns it to the EmailsWithoutAccess field.
func (o *CheckAccessByEmail200Response) SetEmailsWithoutAccess(v []string) {
	o.EmailsWithoutAccess = v
}

// GetInvalidEmails returns the InvalidEmails field value if set, zero value otherwise.
func (o *CheckAccessByEmail200Response) GetInvalidEmails() []string {
	if o == nil || IsNil(o.InvalidEmails) {
		var ret []string
		return ret
	}
	return o.InvalidEmails
}

// GetInvalidEmailsOk returns a tuple with the InvalidEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAccessByEmail200Response) GetInvalidEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.InvalidEmails) {
		return nil, false
	}
	return o.InvalidEmails, true
}

// HasInvalidEmails returns a boolean if a field has been set.
func (o *CheckAccessByEmail200Response) HasInvalidEmails() bool {
	if o != nil && !IsNil(o.InvalidEmails) {
		return true
	}

	return false
}

// SetInvalidEmails gets a reference to the given []string and assigns it to the InvalidEmails field.
func (o *CheckAccessByEmail200Response) SetInvalidEmails(v []string) {
	o.InvalidEmails = v
}

func (o CheckAccessByEmail200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckAccessByEmail200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailsWithoutAccess) {
		toSerialize["emailsWithoutAccess"] = o.EmailsWithoutAccess
	}
	if !IsNil(o.InvalidEmails) {
		toSerialize["invalidEmails"] = o.InvalidEmails
	}
	return toSerialize, nil
}

type NullableCheckAccessByEmail200Response struct {
	value *CheckAccessByEmail200Response
	isSet bool
}

func (v NullableCheckAccessByEmail200Response) Get() *CheckAccessByEmail200Response {
	return v.value
}

func (v *NullableCheckAccessByEmail200Response) Set(val *CheckAccessByEmail200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAccessByEmail200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAccessByEmail200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAccessByEmail200Response(val *CheckAccessByEmail200Response) *NullableCheckAccessByEmail200Response {
	return &NullableCheckAccessByEmail200Response{value: val, isSet: true}
}

func (v NullableCheckAccessByEmail200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAccessByEmail200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
