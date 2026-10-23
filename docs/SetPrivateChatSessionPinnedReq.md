# SetPrivateChatSessionPinnedReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerUserId** | Pointer to **string** |  | [optional] 
**Pinned** | Pointer to **bool** |  | [optional] 

## Methods

### NewSetPrivateChatSessionPinnedReq

`func NewSetPrivateChatSessionPinnedReq() *SetPrivateChatSessionPinnedReq`

NewSetPrivateChatSessionPinnedReq instantiates a new SetPrivateChatSessionPinnedReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPrivateChatSessionPinnedReqWithDefaults

`func NewSetPrivateChatSessionPinnedReqWithDefaults() *SetPrivateChatSessionPinnedReq`

NewSetPrivateChatSessionPinnedReqWithDefaults instantiates a new SetPrivateChatSessionPinnedReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerUserId

`func (o *SetPrivateChatSessionPinnedReq) GetPeerUserId() string`

GetPeerUserId returns the PeerUserId field if non-nil, zero value otherwise.

### GetPeerUserIdOk

`func (o *SetPrivateChatSessionPinnedReq) GetPeerUserIdOk() (*string, bool)`

GetPeerUserIdOk returns a tuple with the PeerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerUserId

`func (o *SetPrivateChatSessionPinnedReq) SetPeerUserId(v string)`

SetPeerUserId sets PeerUserId field to given value.

### HasPeerUserId

`func (o *SetPrivateChatSessionPinnedReq) HasPeerUserId() bool`

HasPeerUserId returns a boolean if a field has been set.

### GetPinned

`func (o *SetPrivateChatSessionPinnedReq) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *SetPrivateChatSessionPinnedReq) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *SetPrivateChatSessionPinnedReq) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *SetPrivateChatSessionPinnedReq) HasPinned() bool`

HasPinned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


