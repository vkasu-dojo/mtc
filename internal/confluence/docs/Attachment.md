# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the attachment. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**PageId** | Pointer to [**AttachmentPageId**](AttachmentPageId.md) |  | [optional] 
**BlogPostId** | Pointer to [**AttachmentBlogPostId**](AttachmentBlogPostId.md) |  | [optional] 
**CustomContentId** | Pointer to [**AttachmentCustomContentId**](AttachmentCustomContentId.md) |  | [optional] 
**MediaType** | Pointer to **string** | Media Type for the attachment. | [optional] 
**MediaTypeDescription** | Pointer to **string** | Media Type description for the attachment. | [optional] 
**Comment** | Pointer to **string** | Comment for the attachment. | [optional] 
**FileSize** | Pointer to **int64** | File size of the attachment. | [optional] 
**WebuiLink** | Pointer to **string** | WebUI link of the attachment. | [optional] 
**DownloadLink** | Pointer to **string** | Download link of the attachment. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 

## Methods

### NewAttachment

`func NewAttachment() *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Attachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Attachment) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Attachment) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Attachment) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Attachment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *Attachment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Attachment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Attachment) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Attachment) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPageId

`func (o *Attachment) GetPageId() AttachmentPageId`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *Attachment) GetPageIdOk() (*AttachmentPageId, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *Attachment) SetPageId(v AttachmentPageId)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *Attachment) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *Attachment) GetBlogPostId() AttachmentBlogPostId`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *Attachment) GetBlogPostIdOk() (*AttachmentBlogPostId, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *Attachment) SetBlogPostId(v AttachmentBlogPostId)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *Attachment) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *Attachment) GetCustomContentId() AttachmentCustomContentId`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *Attachment) GetCustomContentIdOk() (*AttachmentCustomContentId, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *Attachment) SetCustomContentId(v AttachmentCustomContentId)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *Attachment) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetMediaType

`func (o *Attachment) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *Attachment) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *Attachment) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *Attachment) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMediaTypeDescription

`func (o *Attachment) GetMediaTypeDescription() string`

GetMediaTypeDescription returns the MediaTypeDescription field if non-nil, zero value otherwise.

### GetMediaTypeDescriptionOk

`func (o *Attachment) GetMediaTypeDescriptionOk() (*string, bool)`

GetMediaTypeDescriptionOk returns a tuple with the MediaTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypeDescription

`func (o *Attachment) SetMediaTypeDescription(v string)`

SetMediaTypeDescription sets MediaTypeDescription field to given value.

### HasMediaTypeDescription

`func (o *Attachment) HasMediaTypeDescription() bool`

HasMediaTypeDescription returns a boolean if a field has been set.

### GetComment

`func (o *Attachment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Attachment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Attachment) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Attachment) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFileSize

`func (o *Attachment) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *Attachment) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *Attachment) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *Attachment) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetWebuiLink

`func (o *Attachment) GetWebuiLink() string`

GetWebuiLink returns the WebuiLink field if non-nil, zero value otherwise.

### GetWebuiLinkOk

`func (o *Attachment) GetWebuiLinkOk() (*string, bool)`

GetWebuiLinkOk returns a tuple with the WebuiLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiLink

`func (o *Attachment) SetWebuiLink(v string)`

SetWebuiLink sets WebuiLink field to given value.

### HasWebuiLink

`func (o *Attachment) HasWebuiLink() bool`

HasWebuiLink returns a boolean if a field has been set.

### GetDownloadLink

`func (o *Attachment) GetDownloadLink() string`

GetDownloadLink returns the DownloadLink field if non-nil, zero value otherwise.

### GetDownloadLinkOk

`func (o *Attachment) GetDownloadLinkOk() (*string, bool)`

GetDownloadLinkOk returns a tuple with the DownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLink

`func (o *Attachment) SetDownloadLink(v string)`

SetDownloadLink sets DownloadLink field to given value.

### HasDownloadLink

`func (o *Attachment) HasDownloadLink() bool`

HasDownloadLink returns a boolean if a field has been set.

### GetVersion

`func (o *Attachment) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Attachment) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Attachment) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Attachment) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


