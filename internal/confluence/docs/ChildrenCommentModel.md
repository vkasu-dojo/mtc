# ChildrenCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**PageCommentModelId**](PageCommentModelId.md) |  | [optional] 
**Status** | Pointer to [**ContentStatus**](ContentStatus.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the comment. | [optional] 
**ParentCommentId** | Pointer to [**ChildrenCommentModelParentCommentId**](ChildrenCommentModelParentCommentId.md) |  | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Body** | Pointer to [**Body**](Body.md) |  | [optional] 

## Methods

### NewChildrenCommentModel

`func NewChildrenCommentModel() *ChildrenCommentModel`

NewChildrenCommentModel instantiates a new ChildrenCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildrenCommentModelWithDefaults

`func NewChildrenCommentModelWithDefaults() *ChildrenCommentModel`

NewChildrenCommentModelWithDefaults instantiates a new ChildrenCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChildrenCommentModel) GetId() PageCommentModelId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChildrenCommentModel) GetIdOk() (*PageCommentModelId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChildrenCommentModel) SetId(v PageCommentModelId)`

SetId sets Id field to given value.

### HasId

`func (o *ChildrenCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ChildrenCommentModel) GetStatus() ContentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChildrenCommentModel) GetStatusOk() (*ContentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChildrenCommentModel) SetStatus(v ContentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChildrenCommentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ChildrenCommentModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChildrenCommentModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChildrenCommentModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChildrenCommentModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentCommentId

`func (o *ChildrenCommentModel) GetParentCommentId() ChildrenCommentModelParentCommentId`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *ChildrenCommentModel) GetParentCommentIdOk() (*ChildrenCommentModelParentCommentId, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *ChildrenCommentModel) SetParentCommentId(v ChildrenCommentModelParentCommentId)`

SetParentCommentId sets ParentCommentId field to given value.

### HasParentCommentId

`func (o *ChildrenCommentModel) HasParentCommentId() bool`

HasParentCommentId returns a boolean if a field has been set.

### GetVersion

`func (o *ChildrenCommentModel) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ChildrenCommentModel) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ChildrenCommentModel) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ChildrenCommentModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBody

`func (o *ChildrenCommentModel) GetBody() Body`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ChildrenCommentModel) GetBodyOk() (*Body, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ChildrenCommentModel) SetBody(v Body)`

SetBody sets Body field to given value.

### HasBody

`func (o *ChildrenCommentModel) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


