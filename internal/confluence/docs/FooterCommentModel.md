# FooterCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**PageCommentModelId**](PageCommentModelId.md) |  | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**BlogPostId** | Pointer to [**FooterCommentModelBlogPostId**](FooterCommentModelBlogPostId.md) |  | [optional] 
**PageId** | Pointer to [**FooterCommentModelPageId**](FooterCommentModelPageId.md) |  | [optional] 
**ParentCommentId** | Pointer to [**FooterCommentModelParentCommentId**](FooterCommentModelParentCommentId.md) |  | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**Body**](Body.md) |  | [optional] 

## Methods

### NewFooterCommentModel

`func NewFooterCommentModel() *FooterCommentModel`

NewFooterCommentModel instantiates a new FooterCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFooterCommentModelWithDefaults

`func NewFooterCommentModelWithDefaults() *FooterCommentModel`

NewFooterCommentModelWithDefaults instantiates a new FooterCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FooterCommentModel) GetId() PageCommentModelId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FooterCommentModel) GetIdOk() (*PageCommentModelId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FooterCommentModel) SetId(v PageCommentModelId)`

SetId sets Id field to given value.

### HasId

`func (o *FooterCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *FooterCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FooterCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FooterCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FooterCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *FooterCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FooterCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FooterCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FooterCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBlogPostId

`func (o *FooterCommentModel) GetBlogPostId() FooterCommentModelBlogPostId`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *FooterCommentModel) GetBlogPostIdOk() (*FooterCommentModelBlogPostId, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *FooterCommentModel) SetBlogPostId(v FooterCommentModelBlogPostId)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *FooterCommentModel) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetPageId

`func (o *FooterCommentModel) GetPageId() FooterCommentModelPageId`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *FooterCommentModel) GetPageIdOk() (*FooterCommentModelPageId, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *FooterCommentModel) SetPageId(v FooterCommentModelPageId)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *FooterCommentModel) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetParentCommentId

`func (o *FooterCommentModel) GetParentCommentId() FooterCommentModelParentCommentId`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *FooterCommentModel) GetParentCommentIdOk() (*FooterCommentModelParentCommentId, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *FooterCommentModel) SetParentCommentId(v FooterCommentModelParentCommentId)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *FooterCommentModel) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetVersion

`func (o *FooterCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FooterCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FooterCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FooterCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *FooterCommentModel) GetBody() Body`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *FooterCommentModel) GetBodyOk() (*Body, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *FooterCommentModel) SetBody(v Body)`

SetBody sets Body field to given value.

### HasBody

`func (o *FooterCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


