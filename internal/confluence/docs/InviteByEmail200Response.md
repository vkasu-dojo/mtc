# InviteByEmail200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailsInvited** | Pointer to **[]string** | List of emails invited to site. | [optional] 
**EmailsNotInvited** | Pointer to [**InviteByEmail200ResponseEmailsNotInvited**](InviteByEmail200ResponseEmailsNotInvited.md) |  | [optional] 

## Methods

### NewInviteByEmail200Response

`func NewInviteByEmail200Response() *InviteByEmail200Response`

NewInviteByEmail200Response instantiates a new InviteByEmail200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteByEmail200ResponseWithDefaults

`func NewInviteByEmail200ResponseWithDefaults() *InviteByEmail200Response`

NewInviteByEmail200ResponseWithDefaults instantiates a new InviteByEmail200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailsInvited

`func (o *InviteByEmail200Response) GetEmailsInvited() []string`

GetEmailsInvited returns the EmailsInvited field if non-nil, zero value otherwise.

### GetEmailsInvitedOk

`func (o *InviteByEmail200Response) GetEmailsInvitedOk() (*[]string, bool)`

GetEmailsInvitedOk returns a tuple with the EmailsInvited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsInvited

`func (o *InviteByEmail200Response) SetEmailsInvited(v []string)`

SetEmailsInvited sets EmailsInvited field to given value.

### HasEmailsInvited

`func (o *InviteByEmail200Response) HasEmailsInvited() bool`

HasEmailsInvited returns a boolean if a field has been set.

### GetEmailsNotInvited

`func (o *InviteByEmail200Response) GetEmailsNotInvited() InviteByEmail200ResponseEmailsNotInvited`

GetEmailsNotInvited returns the EmailsNotInvited field if non-nil, zero value otherwise.

### GetEmailsNotInvitedOk

`func (o *InviteByEmail200Response) GetEmailsNotInvitedOk() (*InviteByEmail200ResponseEmailsNotInvited, bool)`

GetEmailsNotInvitedOk returns a tuple with the EmailsNotInvited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsNotInvited

`func (o *InviteByEmail200Response) SetEmailsNotInvited(v InviteByEmail200ResponseEmailsNotInvited)`

SetEmailsNotInvited sets EmailsNotInvited field to given value.

### HasEmailsNotInvited

`func (o *InviteByEmail200Response) HasEmailsNotInvited() bool`

HasEmailsNotInvited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


