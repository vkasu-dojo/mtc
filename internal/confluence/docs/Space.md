# Space

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**SpaceId**](SpaceId.md) |  | [optional] 
**Key** | Pointer to **string** | Key of the space. | [optional] 
**Name** | Pointer to **string** | Name of the space. | [optional] 
**Type** | Pointer to [**SpaceType**](SpaceType.md) |  | [optional] 
**Status** | Pointer to [**SpaceStatus**](SpaceStatus.md) |  | [optional] 
**HomepageId** | Pointer to [**SpaceHomepageId**](SpaceHomepageId.md) |  | [optional] 
**Description** | Pointer to [**SpaceDescription**](SpaceDescription.md) |  | [optional] 

## Methods

### NewSpace

`func NewSpace() *Space`

NewSpace instantiates a new Space object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceWithDefaults

`func NewSpaceWithDefaults() *Space`

NewSpaceWithDefaults instantiates a new Space object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Space) GetId() SpaceId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Space) GetIdOk() (*SpaceId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Space) SetId(v SpaceId)`

SetId sets Id field to given value.

### HasId

`func (o *Space) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *Space) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Space) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Space) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Space) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *Space) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Space) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Space) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Space) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Space) GetType() SpaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Space) GetTypeOk() (*SpaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Space) SetType(v SpaceType)`

SetType sets Type field to given value.

### HasType

`func (o *Space) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Space) GetStatus() SpaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Space) GetStatusOk() (*SpaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Space) SetStatus(v SpaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Space) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHomepageId

`func (o *Space) GetHomepageId() SpaceHomepageId`

GetHomepageId returns the HomepageId field if non-nil, zero value otherwise.

### GetHomepageIdOk

`func (o *Space) GetHomepageIdOk() (*SpaceHomepageId, bool)`

GetHomepageIdOk returns a tuple with the HomepageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageId

`func (o *Space) SetHomepageId(v SpaceHomepageId)`

SetHomepageId sets HomepageId field to given value.

### HasHomepageId

`func (o *Space) HasHomepageId() bool`

HasHomepageId returns a boolean if a field has been set.

### GetDescription

`func (o *Space) GetDescription() SpaceDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Space) GetDescriptionOk() (*SpaceDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Space) SetDescription(v SpaceDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Space) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


