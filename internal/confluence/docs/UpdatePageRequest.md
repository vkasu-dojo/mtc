# UpdatePageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the page. | 
**Status** | **string** | The status of the page. | 
**Title** | **string** | Title of the page. | 
**SpaceId** | **string** | ID of the containing space. | 
**Body** | [**CreatePageRequestBody**](CreatePageRequestBody.md) |  | 
**Version** | [**UpdateBlogPostRequestVersion**](UpdateBlogPostRequestVersion.md) |  | 

## Methods

### NewUpdatePageRequest

`func NewUpdatePageRequest(id string, status string, title string, spaceId string, body CreatePageRequestBody, version UpdateBlogPostRequestVersion, ) *UpdatePageRequest`

NewUpdatePageRequest instantiates a new UpdatePageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePageRequestWithDefaults

`func NewUpdatePageRequestWithDefaults() *UpdatePageRequest`

NewUpdatePageRequestWithDefaults instantiates a new UpdatePageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatePageRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePageRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePageRequest) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *UpdatePageRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdatePageRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdatePageRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *UpdatePageRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePageRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePageRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSpaceId

`func (o *UpdatePageRequest) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *UpdatePageRequest) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *UpdatePageRequest) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetBody

`func (o *UpdatePageRequest) GetBody() CreatePageRequestBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdatePageRequest) GetBodyOk() (*CreatePageRequestBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdatePageRequest) SetBody(v CreatePageRequestBody)`

SetBody sets Body field to given value.


### GetVersion

`func (o *UpdatePageRequest) GetVersion() UpdateBlogPostRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdatePageRequest) GetVersionOk() (*UpdateBlogPostRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdatePageRequest) SetVersion(v UpdateBlogPostRequestVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


