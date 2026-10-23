# SendGlobalChatMessageResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to [**GlobalChatMessage**](GlobalChatMessage.md) |  | [optional] 

## Methods

### NewSendGlobalChatMessageResp

`func NewSendGlobalChatMessageResp() *SendGlobalChatMessageResp`

NewSendGlobalChatMessageResp instantiates a new SendGlobalChatMessageResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendGlobalChatMessageRespWithDefaults

`func NewSendGlobalChatMessageRespWithDefaults() *SendGlobalChatMessageResp`

NewSendGlobalChatMessageRespWithDefaults instantiates a new SendGlobalChatMessageResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SendGlobalChatMessageResp) GetMessage() GlobalChatMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendGlobalChatMessageResp) GetMessageOk() (*GlobalChatMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendGlobalChatMessageResp) SetMessage(v GlobalChatMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendGlobalChatMessageResp) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


