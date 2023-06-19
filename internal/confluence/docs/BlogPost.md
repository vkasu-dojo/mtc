# BlogPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**BlogPostId**](BlogPostId.md) |  | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the blog post. | [optional] 
**SpaceId** | Pointer to [**BlogPostSpaceId**](BlogPostSpaceId.md) |  | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this blog post originally. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the blog post was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**Body**](Body.md) |  | [optional] 

## Methods

### NewBlogPost

`func NewBlogPost() *BlogPost`

NewBlogPost instantiates a new BlogPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogPostWithDefaults

`func NewBlogPostWithDefaults() *BlogPost`

NewBlogPostWithDefaults instantiates a new BlogPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogPost) GetId() BlogPostId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogPost) GetIdOk() (*BlogPostId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogPost) SetId(v BlogPostId)`

SetId sets Id field to given value.

### HasId

`func (o *BlogPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BlogPost) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlogPost) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlogPost) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlogPost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BlogPost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BlogPost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BlogPost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BlogPost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *BlogPost) GetSpaceId() BlogPostSpaceId`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *BlogPost) GetSpaceIdOk() (*BlogPostSpaceId, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *BlogPost) SetSpaceId(v BlogPostSpaceId)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *BlogPost) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetAuthorId

`func (o *BlogPost) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *BlogPost) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *BlogPost) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *BlogPost) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BlogPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlogPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlogPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BlogPost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *BlogPost) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlogPost) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlogPost) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlogPost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *BlogPost) GetBody() Body`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *BlogPost) GetBodyOk() (*Body, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *BlogPost) SetBody(v Body)`

SetBody sets Body field to given value.

### HasBody

`func (o *BlogPost) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


