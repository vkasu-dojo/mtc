# BlogPostCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**PageCommentModelId**](PageCommentModelId.md) |  | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**BlogPostId** | Pointer to [**BlogPostCommentModelBlogPostId**](BlogPostCommentModelBlogPostId.md) |  | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**Body**](Body.md) |  | [optional] 

## Methods

### NewBlogPostCommentModel

`func NewBlogPostCommentModel() *BlogPostCommentModel`

NewBlogPostCommentModel instantiates a new BlogPostCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostCommentModelWithDefaults

`func NewBlogPostCommentModelWithDefaults() *BlogPostCommentModel`

NewBlogPostCommentModelWithDefaults instantiates a new BlogPostCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPostCommentModel) GetId() PageCommentModelId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPostCommentModel) GetIdOk() (*PageCommentModelId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPostCommentModel) SetId(v PageCommentModelId)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPostCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BlogPostCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlogPostCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlogPostCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlogPostCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BlogPostCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BlogPostCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BlogPostCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BlogPostCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBlogPostId

`func (o *BlogPostCommentModel) GetBlogPostId() BlogPostCommentModelBlogPostId`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *BlogPostCommentModel) GetBlogPostIdOk() (*BlogPostCommentModelBlogPostId, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *BlogPostCommentModel) SetBlogPostId(v BlogPostCommentModelBlogPostId)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *BlogPostCommentModel) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetVersion

`func (o *BlogPostCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlogPostCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlogPostCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlogPostCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *BlogPostCommentModel) GetBody() Body`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BlogPostCommentModel) GetBodyOk() (*Body, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BlogPostCommentModel) SetBody(v Body)`

SetBody sets Body field to given value.

### HasBody

`func (o *BlogPostCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


