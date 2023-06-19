# Body

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | Pointer to [**BodyType**](BodyType.md) |  | [optional] 
**AtlasDocFormat** | Pointer to [**BodyType**](BodyType.md) |  | [optional] 

## Methods

### NewBody

`func NewBody() *Body`

NewBody instantiates a new Body object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyWithDefaults

`func NewBodyWithDefaults() *Body`

NewBodyWithDefaults instantiates a new Body object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *Body) GetStorage() BodyType`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Body) GetStorageOk() (*BodyType, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Body) SetStorage(v BodyType)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Body) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetAtlasDocFormat

`func (o *Body) GetAtlasDocFormat() BodyType`

GetAtlasDocFormat returns the AtlasDocFormat field if non-nil, zero value otherwise.

### GetAtlasDocFormatOk

`func (o *Body) GetAtlasDocFormatOk() (*BodyType, bool)`

GetAtlasDocFormatOk returns a tuple with the AtlasDocFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasDocFormat

`func (o *Body) SetAtlasDocFormat(v BodyType)`

SetAtlasDocFormat sets AtlasDocFormat field to given value.

### HasAtlasDocFormat

`func (o *Body) HasAtlasDocFormat() bool`

HasAtlasDocFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


