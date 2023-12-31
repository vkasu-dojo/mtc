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

// checks if the UpdatePageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePageRequest{}

// UpdatePageRequest struct for UpdatePageRequest
type UpdatePageRequest struct {
	// Id of the page.
	Id string `json:"id"`
	// The status of the page.
	Status string `json:"status"`
	// Title of the page.
	Title string `json:"title"`
	// ID of the containing space.
	SpaceId string                       `json:"spaceId"`
	Body    CreatePageRequestBody        `json:"body"`
	Version UpdateBlogPostRequestVersion `json:"version"`
}

// NewUpdatePageRequest instantiates a new UpdatePageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePageRequest(id string, status string, title string, spaceId string, body CreatePageRequestBody, version UpdateBlogPostRequestVersion) *UpdatePageRequest {
	this := UpdatePageRequest{}
	this.Id = id
	this.Status = status
	this.Title = title
	this.SpaceId = spaceId
	this.Body = body
	this.Version = version
	return &this
}

// NewUpdatePageRequestWithDefaults instantiates a new UpdatePageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePageRequestWithDefaults() *UpdatePageRequest {
	this := UpdatePageRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdatePageRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdatePageRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdatePageRequest) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *UpdatePageRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdatePageRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdatePageRequest) SetStatus(v string) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *UpdatePageRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UpdatePageRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UpdatePageRequest) SetTitle(v string) {
	o.Title = v
}

// GetSpaceId returns the SpaceId field value
func (o *UpdatePageRequest) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *UpdatePageRequest) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *UpdatePageRequest) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetBody returns the Body field value
func (o *UpdatePageRequest) GetBody() CreatePageRequestBody {
	if o == nil {
		var ret CreatePageRequestBody
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *UpdatePageRequest) GetBodyOk() (*CreatePageRequestBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *UpdatePageRequest) SetBody(v CreatePageRequestBody) {
	o.Body = v
}

// GetVersion returns the Version field value
func (o *UpdatePageRequest) GetVersion() UpdateBlogPostRequestVersion {
	if o == nil {
		var ret UpdateBlogPostRequestVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdatePageRequest) GetVersionOk() (*UpdateBlogPostRequestVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UpdatePageRequest) SetVersion(v UpdateBlogPostRequestVersion) {
	o.Version = v
}

func (o UpdatePageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["spaceId"] = o.SpaceId
	toSerialize["body"] = o.Body
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableUpdatePageRequest struct {
	value *UpdatePageRequest
	isSet bool
}

func (v NullableUpdatePageRequest) Get() *UpdatePageRequest {
	return v.value
}

func (v *NullableUpdatePageRequest) Set(val *UpdatePageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePageRequest(val *UpdatePageRequest) *NullableUpdatePageRequest {
	return &NullableUpdatePageRequest{value: val, isSet: true}
}

func (v NullableUpdatePageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
