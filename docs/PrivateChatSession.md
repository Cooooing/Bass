# PrivateChatSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerUserId** | Pointer to **string** |  | [optional] 
**LastReadMessageId** | Pointer to **string** |  | [optional] 
**UnreadCount** | Pointer to **int32** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**Pinned** | Pointer to **bool** |  | [optional] 
**LastMessage** | Pointer to [**PrivateChatMessage**](PrivateChatMessage.md) |  | [optional] 
**LastMessageAt** | Pointer to **time.Time** |  | [optional] 
**Peer** | Pointer to [**PrivateChatMember**](PrivateChatMember.md) |  | [optional] 

## Methods

### NewPrivateChatSession

`func NewPrivateChatSession() *PrivateChatSession`

NewPrivateChatSession instantiates a new PrivateChatSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateChatSessionWithDefaults

`func NewPrivateChatSessionWithDefaults() *PrivateChatSession`

NewPrivateChatSessionWithDefaults instantiates a new PrivateChatSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerUserId

`func (o *PrivateChatSession) GetPeerUserId() string`

GetPeerUserId returns the PeerUserId field if non-nil, zero value otherwise.

### GetPeerUserIdOk

`func (o *PrivateChatSession) GetPeerUserIdOk() (*string, bool)`

GetPeerUserIdOk returns a tuple with the PeerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerUserId

`func (o *PrivateChatSession) SetPeerUserId(v string)`

SetPeerUserId sets PeerUserId field to given value.

### HasPeerUserId

`func (o *PrivateChatSession) HasPeerUserId() bool`

HasPeerUserId returns a boolean if a field has been set.

### GetLastReadMessageId

`func (o *PrivateChatSession) GetLastReadMessageId() string`

GetLastReadMessageId returns the LastReadMessageId field if non-nil, zero value otherwise.

### GetLastReadMessageIdOk

`func (o *PrivateChatSession) GetLastReadMessageIdOk() (*string, bool)`

GetLastReadMessageIdOk returns a tuple with the LastReadMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadMessageId

`func (o *PrivateChatSession) SetLastReadMessageId(v string)`

SetLastReadMessageId sets LastReadMessageId field to given value.

### HasLastReadMessageId

`func (o *PrivateChatSession) HasLastReadMessageId() bool`

HasLastReadMessageId returns a boolean if a field has been set.

### GetUnreadCount

`func (o *PrivateChatSession) GetUnreadCount() int32`

GetUnreadCount returns the UnreadCount field if non-nil, zero value otherwise.

### GetUnreadCountOk

`func (o *PrivateChatSession) GetUnreadCountOk() (*int32, bool)`

GetUnreadCountOk returns a tuple with the UnreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadCount

`func (o *PrivateChatSession) SetUnreadCount(v int32)`

SetUnreadCount sets UnreadCount field to given value.

### HasUnreadCount

`func (o *PrivateChatSession) HasUnreadCount() bool`

HasUnreadCount returns a boolean if a field has been set.

### GetMuted

`func (o *PrivateChatSession) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *PrivateChatSession) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *PrivateChatSession) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *PrivateChatSession) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetPinned

`func (o *PrivateChatSession) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *PrivateChatSession) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *PrivateChatSession) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *PrivateChatSession) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetLastMessage

`func (o *PrivateChatSession) GetLastMessage() PrivateChatMessage`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *PrivateChatSession) GetLastMessageOk() (*PrivateChatMessage, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *PrivateChatSession) SetLastMessage(v PrivateChatMessage)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *PrivateChatSession) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### GetLastMessageAt

`func (o *PrivateChatSession) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *PrivateChatSession) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *PrivateChatSession) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.

### HasLastMessageAt

`func (o *PrivateChatSession) HasLastMessageAt() bool`

HasLastMessageAt returns a boolean if a field has been set.

### GetPeer

`func (o *PrivateChatSession) GetPeer() PrivateChatMember`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *PrivateChatSession) GetPeerOk() (*PrivateChatMember, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *PrivateChatSession) SetPeer(v PrivateChatMember)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *PrivateChatSession) HasPeer() bool`

HasPeer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


