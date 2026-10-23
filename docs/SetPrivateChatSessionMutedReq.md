# SetPrivateChatSessionMutedReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerUserId** | Pointer to **string** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 

## Methods

### NewSetPrivateChatSessionMutedReq

`func NewSetPrivateChatSessionMutedReq() *SetPrivateChatSessionMutedReq`

NewSetPrivateChatSessionMutedReq instantiates a new SetPrivateChatSessionMutedReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPrivateChatSessionMutedReqWithDefaults

`func NewSetPrivateChatSessionMutedReqWithDefaults() *SetPrivateChatSessionMutedReq`

NewSetPrivateChatSessionMutedReqWithDefaults instantiates a new SetPrivateChatSessionMutedReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerUserId

`func (o *SetPrivateChatSessionMutedReq) GetPeerUserId() string`

GetPeerUserId returns the PeerUserId field if non-nil, zero value otherwise.

### GetPeerUserIdOk

`func (o *SetPrivateChatSessionMutedReq) GetPeerUserIdOk() (*string, bool)`

GetPeerUserIdOk returns a tuple with the PeerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerUserId

`func (o *SetPrivateChatSessionMutedReq) SetPeerUserId(v string)`

SetPeerUserId sets PeerUserId field to given value.

### HasPeerUserId

`func (o *SetPrivateChatSessionMutedReq) HasPeerUserId() bool`

HasPeerUserId returns a boolean if a field has been set.

### GetMuted

`func (o *SetPrivateChatSessionMutedReq) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *SetPrivateChatSessionMutedReq) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *SetPrivateChatSessionMutedReq) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *SetPrivateChatSessionMutedReq) HasMuted() bool`

HasMuted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


