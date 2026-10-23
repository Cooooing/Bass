# GlobalChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SenderId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 
**Sender** | Pointer to [**GlobalChatMember**](GlobalChatMember.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**RevokedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGlobalChatMessage

`func NewGlobalChatMessage() *GlobalChatMessage`

NewGlobalChatMessage instantiates a new GlobalChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalChatMessageWithDefaults

`func NewGlobalChatMessageWithDefaults() *GlobalChatMessage`

NewGlobalChatMessageWithDefaults instantiates a new GlobalChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalChatMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalChatMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalChatMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalChatMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSenderId

`func (o *GlobalChatMessage) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *GlobalChatMessage) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *GlobalChatMessage) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *GlobalChatMessage) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetContent

`func (o *GlobalChatMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GlobalChatMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GlobalChatMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GlobalChatMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMessageType

`func (o *GlobalChatMessage) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *GlobalChatMessage) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *GlobalChatMessage) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *GlobalChatMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetRevoked

`func (o *GlobalChatMessage) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *GlobalChatMessage) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *GlobalChatMessage) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *GlobalChatMessage) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetSender

`func (o *GlobalChatMessage) GetSender() GlobalChatMember`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GlobalChatMessage) GetSenderOk() (*GlobalChatMember, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GlobalChatMessage) SetSender(v GlobalChatMember)`

SetSender sets Sender field to given value.

### HasSender

`func (o *GlobalChatMessage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalChatMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalChatMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalChatMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalChatMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRevokedAt

`func (o *GlobalChatMessage) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *GlobalChatMessage) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *GlobalChatMessage) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *GlobalChatMessage) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


