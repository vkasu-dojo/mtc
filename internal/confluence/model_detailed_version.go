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

// checks if the DetailedVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedVersion{}

// DetailedVersion struct for DetailedVersion
type DetailedVersion struct {
	// The current version number.
	Number *int32 `json:"number,omitempty"`
	// The account ID of the user who created this version.
	AuthorId *string `json:"authorId,omitempty"`
	// Message associated with the current version.
	Message *string `json:"message,omitempty"`
	// Date and time when the version was created. In format \"YYYY-MM-DDTHH:mm:ss.sssZ\".
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Describes if this version is a minor version. Email notifications and activity stream updates are not created for minor versions.
	MinorEdit *bool `json:"minorEdit,omitempty"`
	// Describes if the content type is modified in this version (e.g. page to blog)
	ContentTypeModified *bool `json:"contentTypeModified,omitempty"`
	// The account IDs of users that collaborated on this version.
	Collaborators []string `json:"collaborators,omitempty"`
	// The version number of the version prior to this current content update.
	PrevVersion *int32 `json:"prevVersion,omitempty"`
	// The version number of the version after this current content update.
	NextVersion *int32 `json:"nextVersion,omitempty"`
}

// NewDetailedVersion instantiates a new DetailedVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedVersion() *DetailedVersion {
	this := DetailedVersion{}
	return &this
}

// NewDetailedVersionWithDefaults instantiates a new DetailedVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedVersionWithDefaults() *DetailedVersion {
	this := DetailedVersion{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *DetailedVersion) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *DetailedVersion) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *DetailedVersion) SetNumber(v int32) {
	o.Number = &v
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *DetailedVersion) GetAuthorId() string {
	if o == nil || IsNil(o.AuthorId) {
		var ret string
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetAuthorIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorId) {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *DetailedVersion) HasAuthorId() bool {
	if o != nil && !IsNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given string and assigns it to the AuthorId field.
func (o *DetailedVersion) SetAuthorId(v string) {
	o.AuthorId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DetailedVersion) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DetailedVersion) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DetailedVersion) SetMessage(v string) {
	o.Message = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DetailedVersion) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DetailedVersion) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DetailedVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetMinorEdit returns the MinorEdit field value if set, zero value otherwise.
func (o *DetailedVersion) GetMinorEdit() bool {
	if o == nil || IsNil(o.MinorEdit) {
		var ret bool
		return ret
	}
	return *o.MinorEdit
}

// GetMinorEditOk returns a tuple with the MinorEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetMinorEditOk() (*bool, bool) {
	if o == nil || IsNil(o.MinorEdit) {
		return nil, false
	}
	return o.MinorEdit, true
}

// HasMinorEdit returns a boolean if a field has been set.
func (o *DetailedVersion) HasMinorEdit() bool {
	if o != nil && !IsNil(o.MinorEdit) {
		return true
	}

	return false
}

// SetMinorEdit gets a reference to the given bool and assigns it to the MinorEdit field.
func (o *DetailedVersion) SetMinorEdit(v bool) {
	o.MinorEdit = &v
}

// GetContentTypeModified returns the ContentTypeModified field value if set, zero value otherwise.
func (o *DetailedVersion) GetContentTypeModified() bool {
	if o == nil || IsNil(o.ContentTypeModified) {
		var ret bool
		return ret
	}
	return *o.ContentTypeModified
}

// GetContentTypeModifiedOk returns a tuple with the ContentTypeModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetContentTypeModifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.ContentTypeModified) {
		return nil, false
	}
	return o.ContentTypeModified, true
}

// HasContentTypeModified returns a boolean if a field has been set.
func (o *DetailedVersion) HasContentTypeModified() bool {
	if o != nil && !IsNil(o.ContentTypeModified) {
		return true
	}

	return false
}

// SetContentTypeModified gets a reference to the given bool and assigns it to the ContentTypeModified field.
func (o *DetailedVersion) SetContentTypeModified(v bool) {
	o.ContentTypeModified = &v
}

// GetCollaborators returns the Collaborators field value if set, zero value otherwise.
func (o *DetailedVersion) GetCollaborators() []string {
	if o == nil || IsNil(o.Collaborators) {
		var ret []string
		return ret
	}
	return o.Collaborators
}

// GetCollaboratorsOk returns a tuple with the Collaborators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetCollaboratorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Collaborators) {
		return nil, false
	}
	return o.Collaborators, true
}

// HasCollaborators returns a boolean if a field has been set.
func (o *DetailedVersion) HasCollaborators() bool {
	if o != nil && !IsNil(o.Collaborators) {
		return true
	}

	return false
}

// SetCollaborators gets a reference to the given []string and assigns it to the Collaborators field.
func (o *DetailedVersion) SetCollaborators(v []string) {
	o.Collaborators = v
}

// GetPrevVersion returns the PrevVersion field value if set, zero value otherwise.
func (o *DetailedVersion) GetPrevVersion() int32 {
	if o == nil || IsNil(o.PrevVersion) {
		var ret int32
		return ret
	}
	return *o.PrevVersion
}

// GetPrevVersionOk returns a tuple with the PrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetPrevVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.PrevVersion) {
		return nil, false
	}
	return o.PrevVersion, true
}

// HasPrevVersion returns a boolean if a field has been set.
func (o *DetailedVersion) HasPrevVersion() bool {
	if o != nil && !IsNil(o.PrevVersion) {
		return true
	}

	return false
}

// SetPrevVersion gets a reference to the given int32 and assigns it to the PrevVersion field.
func (o *DetailedVersion) SetPrevVersion(v int32) {
	o.PrevVersion = &v
}

// GetNextVersion returns the NextVersion field value if set, zero value otherwise.
func (o *DetailedVersion) GetNextVersion() int32 {
	if o == nil || IsNil(o.NextVersion) {
		var ret int32
		return ret
	}
	return *o.NextVersion
}

// GetNextVersionOk returns a tuple with the NextVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedVersion) GetNextVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.NextVersion) {
		return nil, false
	}
	return o.NextVersion, true
}

// HasNextVersion returns a boolean if a field has been set.
func (o *DetailedVersion) HasNextVersion() bool {
	if o != nil && !IsNil(o.NextVersion) {
		return true
	}

	return false
}

// SetNextVersion gets a reference to the given int32 and assigns it to the NextVersion field.
func (o *DetailedVersion) SetNextVersion(v int32) {
	o.NextVersion = &v
}

func (o DetailedVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.AuthorId) {
		toSerialize["authorId"] = o.AuthorId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.MinorEdit) {
		toSerialize["minorEdit"] = o.MinorEdit
	}
	if !IsNil(o.ContentTypeModified) {
		toSerialize["contentTypeModified"] = o.ContentTypeModified
	}
	if !IsNil(o.Collaborators) {
		toSerialize["collaborators"] = o.Collaborators
	}
	if !IsNil(o.PrevVersion) {
		toSerialize["prevVersion"] = o.PrevVersion
	}
	if !IsNil(o.NextVersion) {
		toSerialize["nextVersion"] = o.NextVersion
	}
	return toSerialize, nil
}

type NullableDetailedVersion struct {
	value *DetailedVersion
	isSet bool
}

func (v NullableDetailedVersion) Get() *DetailedVersion {
	return v.value
}

func (v *NullableDetailedVersion) Set(val *DetailedVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedVersion(val *DetailedVersion) *NullableDetailedVersion {
	return &NullableDetailedVersion{value: val, isSet: true}
}

func (v NullableDetailedVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
