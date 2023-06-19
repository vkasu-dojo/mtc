# CustomContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**CustomContentId**](CustomContentId.md) |  | [optional] 
**Type** | Pointer to **string** | The type of custom content. | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the custom content. | [optional] 
**SpaceId** | Pointer to [**CustomContentSpaceId**](CustomContentSpaceId.md) |  | [optional] 
**PageId** | Pointer to [**CustomContentPageId**](CustomContentPageId.md) |  | [optional] 
**BlogPostId** | Pointer to [**CustomContentBlogPostId**](CustomContentBlogPostId.md) |  | [optional] 
**CustomContentId** | Pointer to [**CustomContentCustomContentId**](CustomContentCustomContentId.md) |  | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this custom content originally. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the custom content was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Body** | Pointer to [**CustomContentBody**](CustomContentBody.md) |  | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 

## Methods

### NewCustomContent

`func NewCustomContent() *CustomContent`

NewCustomContent instantiates a new CustomContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomContentWithDefaults

`func NewCustomContentWithDefaults() *CustomContent`

NewCustomContentWithDefaults instantiates a new CustomContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomContent) GetId() CustomContentId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomContent) GetIdOk() (*CustomContentId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomContent) SetId(v CustomContentId)`

SetId sets Id field to given value.

### HasId

`func (o *CustomContent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomContent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomContent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CustomContent) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomContent) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomContent) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomContent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CustomContent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomContent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomContent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomContent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *CustomContent) GetSpaceId() CustomContentSpaceId`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *CustomContent) GetSpaceIdOk() (*CustomContentSpaceId, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *CustomContent) SetSpaceId(v CustomContentSpaceId)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *CustomContent) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetPageId

`func (o *CustomContent) GetPageId() CustomContentPageId`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *CustomContent) GetPageIdOk() (*CustomContentPageId, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *CustomContent) SetPageId(v CustomContentPageId)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *CustomContent) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetBlogPostId

`func (o *CustomContent) GetBlogPostId() CustomContentBlogPostId`

GetBlogPostId returns the BlogPostId field if non-nil, zero value otherwise.

### GetBlogPostIdOk

`func (o *CustomContent) GetBlogPostIdOk() (*CustomContentBlogPostId, bool)`

GetBlogPostIdOk returns a tuple with the BlogPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogPostId

`func (o *CustomContent) SetBlogPostId(v CustomContentBlogPostId)`

SetBlogPostId sets BlogPostId field to given value.

### HasBlogPostId

`func (o *CustomContent) HasBlogPostId() bool`

HasBlogPostId returns a boolean if a field has been set.

### GetCustomContentId

`func (o *CustomContent) GetCustomContentId() CustomContentCustomContentId`

GetCustomContentId returns the CustomContentId field if non-nil, zero value otherwise.

### GetCustomContentIdOk

`func (o *CustomContent) GetCustomContentIdOk() (*CustomContentCustomContentId, bool)`

GetCustomContentIdOk returns a tuple with the CustomContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContentId

`func (o *CustomContent) SetCustomContentId(v CustomContentCustomContentId)`

SetCustomContentId sets CustomContentId field to given value.

### HasCustomContentId

`func (o *CustomContent) HasCustomContentId() bool`

HasCustomContentId returns a boolean if a field has been set.

### GetAuthorId

`func (o *CustomContent) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CustomContent) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CustomContent) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CustomContent) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomContent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomContent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomContent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomContent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetBody

`func (o *CustomContent) GetBody() CustomContentBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CustomContent) GetBodyOk() (*CustomContentBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CustomContent) SetBody(v CustomContentBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *CustomContent) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetVersion

`func (o *CustomContent) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CustomContent) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CustomContent) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CustomContent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


