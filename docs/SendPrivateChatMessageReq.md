# SendPrivateChatMessageReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiverId** | Pointer to **string** |  | [optional] 
**ClientMessageId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewSendPrivateChatMessageReq

`func NewSendPrivateChatMessageReq() *SendPrivateChatMessageReq`

NewSendPrivateChatMessageReq instantiates a new SendPrivateChatMessageReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPrivateChatMessageReqWithDefaults

`func NewSendPrivateChatMessageReqWithDefaults() *SendPrivateChatMessageReq`

NewSendPrivateChatMessageReqWithDefaults instantiates a new SendPrivateChatMessageReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiverId

`func (o *SendPrivateChatMessageReq) GetReceiverId() string`

GetReceiverId returns the ReceiverId field if non-nil, zero value otherwise.

### GetReceiverIdOk

`func (o *SendPrivateChatMessageReq) GetReceiverIdOk() (*string, bool)`

GetReceiverIdOk returns a tuple with the ReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverId

`func (o *SendPrivateChatMessageReq) SetReceiverId(v string)`

SetReceiverId sets ReceiverId field to given value.

### HasReceiverId

`func (o *SendPrivateChatMessageReq) HasReceiverId() bool`

HasReceiverId returns a boolean if a field has been set.

### GetClientMessageId

`func (o *SendPrivateChatMessageReq) GetClientMessageId() string`

GetClientMessageId returns the ClientMessageId field if non-nil, zero value otherwise.

### GetClientMessageIdOk

`func (o *SendPrivateChatMessageReq) GetClientMessageIdOk() (*string, bool)`

GetClientMessageIdOk returns a tuple with the ClientMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMessageId

`func (o *SendPrivateChatMessageReq) SetClientMessageId(v string)`

SetClientMessageId sets ClientMessageId field to given value.

### HasClientMessageId

`func (o *SendPrivateChatMessageReq) HasClientMessageId() bool`

HasClientMessageId returns a boolean if a field has been set.

### GetContent

`func (o *SendPrivateChatMessageReq) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SendPrivateChatMessageReq) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SendPrivateChatMessageReq) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SendPrivateChatMessageReq) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


