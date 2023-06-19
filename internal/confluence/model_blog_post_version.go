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

// checks if the BlogPostVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlogPostVersion{}

// BlogPostVersion struct for BlogPostVersion
type BlogPostVersion struct {
	// Date and time when the version was created. In format \"YYYY-MM-DDTHH:mm:ss.sssZ\".
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Message associated with the current version.
	Message *string `json:"message,omitempty"`
	// The version number.
	Number *int32 `json:"number,omitempty"`
	// Describes if this version is a minor version. Email notifications and activity stream updates are not created for minor versions.
	MinorEdit *bool `json:"minorEdit,omitempty"`
	// The account ID of the user who created this version.
	AuthorId *string          `json:"authorId,omitempty"`
	Blogpost *VersionedEntity `json:"blogpost,omitempty"`
}

// NewBlogPostVersion instantiates a new BlogPostVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlogPostVersion() *BlogPostVersion {
	this := BlogPostVersion{}
	return &this
}

// NewBlogPostVersionWithDefaults instantiates a new BlogPostVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlogPostVersionWithDefaults() *BlogPostVersion {
	this := BlogPostVersion{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BlogPostVersion) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BlogPostVersion) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BlogPostVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BlogPostVersion) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostVersion) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BlogPostVersion) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BlogPostVersion) SetMessage(v string) {
	o.Message = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *BlogPostVersion) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostVersion) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *BlogPostVersion) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *BlogPostVersion) SetNumber(v int32) {
	o.Number = &v
}

// GetMinorEdit returns the MinorEdit field value if set, zero value otherwise.
func (o *BlogPostVersion) GetMinorEdit() bool {
	if o == nil || IsNil(o.MinorEdit) {
		var ret bool
		return ret
	}
	return *o.MinorEdit
}

// GetMinorEditOk returns a tuple with the MinorEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostVersion) GetMinorEditOk() (*bool, bool) {
	if o == nil || IsNil(o.MinorEdit) {
		return nil, false
	}
	return o.MinorEdit, true
}

// HasMinorEdit returns a boolean if a field has been set.
func (o *BlogPostVersion) HasMinorEdit() bool {
	if o != nil && !IsNil(o.MinorEdit) {
		return true
	}

	return false
}

// SetMinorEdit gets a reference to the given bool and assigns it to the MinorEdit field.
func (o *BlogPostVersion) SetMinorEdit(v bool) {
	o.MinorEdit = &v
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *BlogPostVersion) GetAuthorId() string {
	if o == nil || IsNil(o.AuthorId) {
		var ret string
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostVersion) GetAuthorIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorId) {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *BlogPostVersion) HasAuthorId() bool {
	if o != nil && !IsNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given string and assigns it to the AuthorId field.
func (o *BlogPostVersion) SetAuthorId(v string) {
	o.AuthorId = &v
}

// GetBlogpost returns the Blogpost field value if set, zero value otherwise.
func (o *BlogPostVersion) GetBlogpost() VersionedEntity {
	if o == nil || IsNil(o.Blogpost) {
		var ret VersionedEntity
		return ret
	}
	return *o.Blogpost
}

// GetBlogpostOk returns a tuple with the Blogpost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostVersion) GetBlogpostOk() (*VersionedEntity, bool) {
	if o == nil || IsNil(o.Blogpost) {
		return nil, false
	}
	return o.Blogpost, true
}

// HasBlogpost returns a boolean if a field has been set.
func (o *BlogPostVersion) HasBlogpost() bool {
	if o != nil && !IsNil(o.Blogpost) {
		return true
	}

	return false
}

// SetBlogpost gets a reference to the given VersionedEntity and assigns it to the Blogpost field.
func (o *BlogPostVersion) SetBlogpost(v VersionedEntity) {
	o.Blogpost = &v
}

func (o BlogPostVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlogPostVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.MinorEdit) {
		toSerialize["minorEdit"] = o.MinorEdit
	}
	if !IsNil(o.AuthorId) {
		toSerialize["authorId"] = o.AuthorId
	}
	if !IsNil(o.Blogpost) {
		toSerialize["blogpost"] = o.Blogpost
	}
	return toSerialize, nil
}

type NullableBlogPostVersion struct {
	value *BlogPostVersion
	isSet bool
}

func (v NullableBlogPostVersion) Get() *BlogPostVersion {
	return v.value
}

func (v *NullableBlogPostVersion) Set(val *BlogPostVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableBlogPostVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableBlogPostVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlogPostVersion(val *BlogPostVersion) *NullableBlogPostVersion {
	return &NullableBlogPostVersion{value: val, isSet: true}
}

func (v NullableBlogPostVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlogPostVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}