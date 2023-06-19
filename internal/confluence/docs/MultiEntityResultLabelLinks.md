# MultiEntityResultLabelLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | Used for pagination. Contains the relative URL for the next set of results, using a cursor query parameter. This property will not be present if there is no additional data available. | [optional] 

## Methods

### NewMultiEntityResultLabelLinks

`func NewMultiEntityResultLabelLinks() *MultiEntityResultLabelLinks`

NewMultiEntityResultLabelLinks instantiates a new MultiEntityResultLabelLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEntityResultLabelLinksWithDefaults

`func NewMultiEntityResultLabelLinksWithDefaults() *MultiEntityResultLabelLinks`

NewMultiEntityResultLabelLinksWithDefaults instantiates a new MultiEntityResultLabelLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *MultiEntityResultLabelLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *MultiEntityResultLabelLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *MultiEntityResultLabelLinks) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *MultiEntityResultLabelLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


