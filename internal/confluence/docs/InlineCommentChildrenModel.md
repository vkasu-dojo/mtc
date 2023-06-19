# InlineCommentChildrenModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**PageCommentModelId**](PageCommentModelId.md) |  | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**ParentCommentId** | Pointer to [**ChildrenCommentModelParentCommentId**](ChildrenCommentModelParentCommentId.md) |  | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**Body**](Body.md) |  | [optional] 
**ResolutionStatus** | Pointer to [**InlineCommentResolutionStatus**](InlineCommentResolutionStatus.md) |  | [optional] 
**Properties** | Pointer to [**InlineCommentProperties**](InlineCommentProperties.md) |  | [optional] 

## Methods

### NewInlineCommentChildrenModel

`func NewInlineCommentChildrenModel() *InlineCommentChildrenModel`

NewInlineCommentChildrenModel instantiates a new InlineCommentChildrenModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineCommentChildrenModelWithDefaults

`func NewInlineCommentChildrenModelWithDefaults() *InlineCommentChildrenModel`

NewInlineCommentChildrenModelWithDefaults instantiates a new InlineCommentChildrenModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineCommentChildrenModel) GetId() PageCommentModelId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineCommentChildrenModel) GetIdOk() (*PageCommentModelId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineCommentChildrenModel) SetId(v PageCommentModelId)`

SetId sets Id field to given value.

### HasId

`func (o *InlineCommentChildrenModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineCommentChildrenModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineCommentChildrenModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineCommentChildrenModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineCommentChildrenModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *InlineCommentChildrenModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineCommentChildrenModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineCommentChildrenModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineCommentChildrenModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentCommentId

`func (o *InlineCommentChildrenModel) GetParentCommentId() ChildrenCommentModelParentCommentId`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *InlineCommentChildrenModel) GetParentCommentIdOk() (*ChildrenCommentModelParentCommentId, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *InlineCommentChildrenModel) SetParentCommentId(v ChildrenCommentModelParentCommentId)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *InlineCommentChildrenModel) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetVersion

`func (o *InlineCommentChildrenModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineCommentChildrenModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineCommentChildrenModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineCommentChildrenModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *InlineCommentChildrenModel) GetBody() Body`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InlineCommentChildrenModel) GetBodyOk() (*Body, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InlineCommentChildrenModel) SetBody(v Body)`

SetBody sets Body field to given value.

### HasBody

`func (o *InlineCommentChildrenModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetResolutionStatus

`func (o *InlineCommentChildrenModel) GetResolutionStatus() InlineCommentResolutionStatus`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *InlineCommentChildrenModel) GetResolutionStatusOk() (*InlineCommentResolutionStatus, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *InlineCommentChildrenModel) SetResolutionStatus(v InlineCommentResolutionStatus)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *InlineCommentChildrenModel) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### GetProperties

`func (o *InlineCommentChildrenModel) GetProperties() InlineCommentProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InlineCommentChildrenModel) GetPropertiesOk() (*InlineCommentProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InlineCommentChildrenModel) SetProperties(v InlineCommentProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InlineCommentChildrenModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


