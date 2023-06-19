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

// checks if the BlogPostInlineCommentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlogPostInlineCommentModel{}

// BlogPostInlineCommentModel struct for BlogPostInlineCommentModel
type BlogPostInlineCommentModel struct {
	Id     *PageCommentModelId `json:"id,omitempty"`
	Status *ContentStatus      `json:"status,omitempty"`
	// Title of the comment.
	Title            *string                         `json:"title,omitempty"`
	BlogPostId       *BlogPostCommentModelBlogPostId `json:"blogPostId,omitempty"`
	Version          *Version                        `json:"version,omitempty"`
	Body             *Body                           `json:"body,omitempty"`
	ResolutionStatus *InlineCommentResolutionStatus  `json:"resolutionStatus,omitempty"`
	Properties       *InlineCommentProperties        `json:"properties,omitempty"`
}

// NewBlogPostInlineCommentModel instantiates a new BlogPostInlineCommentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlogPostInlineCommentModel() *BlogPostInlineCommentModel {
	this := BlogPostInlineCommentModel{}
	return &this
}

// NewBlogPostInlineCommentModelWithDefaults instantiates a new BlogPostInlineCommentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlogPostInlineCommentModelWithDefaults() *BlogPostInlineCommentModel {
	this := BlogPostInlineCommentModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlogPostInlineCommentModel) GetId() PageCommentModelId {
	if o == nil || IsNil(o.Id) {
		var ret PageCommentModelId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostInlineCommentModel) GetIdOk() (*PageCommentModelId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlogPostInlineCommentModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given PageCommentModelId and assigns it to the Id field.
func (o *BlogPostInlineCommentModel) SetId(v PageCommentModelId) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BlogPostInlineCommentModel) GetStatus() ContentStatus {
	if o == nil || IsNil(o.Status) {
		var ret ContentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostInlineCommentModel) GetStatusOk() (*ContentStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BlogPostInlineCommentModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ContentStatus and assigns it to the Status field.
func (o *BlogPostInlineCommentModel) SetStatus(v ContentStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BlogPostInlineCommentModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostInlineCommentModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BlogPostInlineCommentModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BlogPostInlineCommentModel) SetTitle(v string) {
	o.Title = &v
}

// GetBlogPostId returns the BlogPostId field value if set, zero value otherwise.
func (o *BlogPostInlineCommentModel) GetBlogPostId() BlogPostCommentModelBlogPostId {
	if o == nil || IsNil(o.BlogPostId) {
		var ret BlogPostCommentModelBlogPostId
		return ret
	}
	return *o.BlogPostId
}

// GetBlogPostIdOk returns a tuple with the BlogPostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostInlineCommentModel) GetBlogPostIdOk() (*BlogPostCommentModelBlogPostId, bool) {
	if o == nil || IsNil(o.BlogPostId) {
		return nil, false
	}
	return o.BlogPostId, true
}

// HasBlogPostId returns a boolean if a field has been set.
func (o *BlogPostInlineCommentModel) HasBlogPostId() bool {
	if o != nil && !IsNil(o.BlogPostId) {
		return true
	}

	return false
}

// SetBlogPostId gets a reference to the given BlogPostCommentModelBlogPostId and assigns it to the BlogPostId field.
func (o *BlogPostInlineCommentModel) SetBlogPostId(v BlogPostCommentModelBlogPostId) {
	o.BlogPostId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BlogPostInlineCommentModel) GetVersion() Version {
	if o == nil || IsNil(o.Version) {
		var ret Version
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostInlineCommentModel) GetVersionOk() (*Version, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BlogPostInlineCommentModel) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given Version and assigns it to the Version field.
func (o *BlogPostInlineCommentModel) SetVersion(v Version) {
	o.Version = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BlogPostInlineCommentModel) GetBody() Body {
	if o == nil || IsNil(o.Body) {
		var ret Body
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostInlineCommentModel) GetBodyOk() (*Body, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BlogPostInlineCommentModel) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given Body and assigns it to the Body field.
func (o *BlogPostInlineCommentModel) SetBody(v Body) {
	o.Body = &v
}

// GetResolutionStatus returns the ResolutionStatus field value if set, zero value otherwise.
func (o *BlogPostInlineCommentModel) GetResolutionStatus() InlineCommentResolutionStatus {
	if o == nil || IsNil(o.ResolutionStatus) {
		var ret InlineCommentResolutionStatus
		return ret
	}
	return *o.ResolutionStatus
}

// GetResolutionStatusOk returns a tuple with the ResolutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostInlineCommentModel) GetResolutionStatusOk() (*InlineCommentResolutionStatus, bool) {
	if o == nil || IsNil(o.ResolutionStatus) {
		return nil, false
	}
	return o.ResolutionStatus, true
}

// HasResolutionStatus returns a boolean if a field has been set.
func (o *BlogPostInlineCommentModel) HasResolutionStatus() bool {
	if o != nil && !IsNil(o.ResolutionStatus) {
		return true
	}

	return false
}

// SetResolutionStatus gets a reference to the given InlineCommentResolutionStatus and assigns it to the ResolutionStatus field.
func (o *BlogPostInlineCommentModel) SetResolutionStatus(v InlineCommentResolutionStatus) {
	o.ResolutionStatus = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BlogPostInlineCommentModel) GetProperties() InlineCommentProperties {
	if o == nil || IsNil(o.Properties) {
		var ret InlineCommentProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlogPostInlineCommentModel) GetPropertiesOk() (*InlineCommentProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BlogPostInlineCommentModel) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given InlineCommentProperties and assigns it to the Properties field.
func (o *BlogPostInlineCommentModel) SetProperties(v InlineCommentProperties) {
	o.Properties = &v
}

func (o BlogPostInlineCommentModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlogPostInlineCommentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.BlogPostId) {
		toSerialize["blogPostId"] = o.BlogPostId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.ResolutionStatus) {
		toSerialize["resolutionStatus"] = o.ResolutionStatus
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableBlogPostInlineCommentModel struct {
	value *BlogPostInlineCommentModel
	isSet bool
}

func (v NullableBlogPostInlineCommentModel) Get() *BlogPostInlineCommentModel {
	return v.value
}

func (v *NullableBlogPostInlineCommentModel) Set(val *BlogPostInlineCommentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBlogPostInlineCommentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBlogPostInlineCommentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlogPostInlineCommentModel(val *BlogPostInlineCommentModel) *NullableBlogPostInlineCommentModel {
	return &NullableBlogPostInlineCommentModel{value: val, isSet: true}
}

func (v NullableBlogPostInlineCommentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlogPostInlineCommentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
