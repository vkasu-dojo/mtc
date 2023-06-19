# MultiEntityResultAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**Links** | Pointer to [**MultiEntityResultLabelLinks**](MultiEntityResultLabelLinks.md) |  | [optional] 

## Methods

### NewMultiEntityResultAttachment

`func NewMultiEntityResultAttachment() *MultiEntityResultAttachment`

NewMultiEntityResultAttachment instantiates a new MultiEntityResultAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEntityResultAttachmentWithDefaults

`func NewMultiEntityResultAttachmentWithDefaults() *MultiEntityResultAttachment`

NewMultiEntityResultAttachmentWithDefaults instantiates a new MultiEntityResultAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *MultiEntityResultAttachment) GetResults() []Attachment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MultiEntityResultAttachment) GetResultsOk() (*[]Attachment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MultiEntityResultAttachment) SetResults(v []Attachment)`

SetResults sets Results field to given value.

### HasResults

`func (o *MultiEntityResultAttachment) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLinks

`func (o *MultiEntityResultAttachment) GetLinks() MultiEntityResultLabelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MultiEntityResultAttachment) GetLinksOk() (*MultiEntityResultLabelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MultiEntityResultAttachment) SetLinks(v MultiEntityResultLabelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MultiEntityResultAttachment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


