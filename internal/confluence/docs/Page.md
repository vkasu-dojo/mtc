# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**PageId**](PageId.md) |  | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the page. | [optional] 
**SpaceId** | Pointer to [**PageSpaceId**](PageSpaceId.md) |  | [optional] 
**ParentId** | Pointer to [**PageParentId**](PageParentId.md) |  | [optional] 
**AuthorId** | Pointer to **string** | The account ID of the user who created this page originally. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the page was created. In format \&quot;YYYY-MM-DDTHH:mm:ss.sssZ\&quot;. | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**Body**](Body.md) |  | [optional] 

## Methods

### NewPage

`func NewPage() *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Page) GetId() PageId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Page) GetIdOk() (*PageId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Page) SetId(v PageId)`

SetId sets Id field to given value.

### HasId

`func (o *Page) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Page) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Page) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Page) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Page) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *Page) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Page) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Page) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Page) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSpaceId

`func (o *Page) GetSpaceId() PageSpaceId`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *Page) GetSpaceIdOk() (*PageSpaceId, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *Page) SetSpaceId(v PageSpaceId)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *Page) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetParentId

`func (o *Page) GetParentId() PageParentId`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Page) GetParentIdOk() (*PageParentId, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Page) SetParentId(v PageParentId)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Page) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetAuthorId

`func (o *Page) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Page) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Page) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *Page) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Page) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Page) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Page) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Page) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *Page) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Page) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Page) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Page) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *Page) GetBody() Body`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Page) GetBodyOk() (*Body, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Page) SetBody(v Body)`

SetBody sets Body field to given value.

### HasBody

`func (o *Page) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


