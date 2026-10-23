# PrivateChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SenderId** | Pointer to **string** |  | [optional] 
**ReceiverId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**RevokedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPrivateChatMessage

`func NewPrivateChatMessage() *PrivateChatMessage`

NewPrivateChatMessage instantiates a new PrivateChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateChatMessageWithDefaults

`func NewPrivateChatMessageWithDefaults() *PrivateChatMessage`

NewPrivateChatMessageWithDefaults instantiates a new PrivateChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateChatMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateChatMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateChatMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateChatMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSenderId

`func (o *PrivateChatMessage) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *PrivateChatMessage) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *PrivateChatMessage) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *PrivateChatMessage) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetReceiverId

`func (o *PrivateChatMessage) GetReceiverId() string`

GetReceiverId returns the ReceiverId field if non-nil, zero value otherwise.

### GetReceiverIdOk

`func (o *PrivateChatMessage) GetReceiverIdOk() (*string, bool)`

GetReceiverIdOk returns a tuple with the ReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverId

`func (o *PrivateChatMessage) SetReceiverId(v string)`

SetReceiverId sets ReceiverId field to given value.

### HasReceiverId

`func (o *PrivateChatMessage) HasReceiverId() bool`

HasReceiverId returns a boolean if a field has been set.

### GetContent

`func (o *PrivateChatMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PrivateChatMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PrivateChatMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PrivateChatMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMessageType

`func (o *PrivateChatMessage) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *PrivateChatMessage) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *PrivateChatMessage) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *PrivateChatMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetRevoked

`func (o *PrivateChatMessage) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *PrivateChatMessage) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *PrivateChatMessage) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *PrivateChatMessage) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PrivateChatMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrivateChatMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrivateChatMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PrivateChatMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRevokedAt

`func (o *PrivateChatMessage) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *PrivateChatMessage) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *PrivateChatMessage) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *PrivateChatMessage) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


