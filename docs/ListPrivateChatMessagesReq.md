# ListPrivateChatMessagesReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerUserId** | Pointer to **string** |  | [optional] 
**BeforeId** | Pointer to **string** |  | [optional] 
**Page** | Pointer to [**PageReq**](PageReq.md) |  | [optional] 

## Methods

### NewListPrivateChatMessagesReq

`func NewListPrivateChatMessagesReq() *ListPrivateChatMessagesReq`

NewListPrivateChatMessagesReq instantiates a new ListPrivateChatMessagesReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPrivateChatMessagesReqWithDefaults

`func NewListPrivateChatMessagesReqWithDefaults() *ListPrivateChatMessagesReq`

NewListPrivateChatMessagesReqWithDefaults instantiates a new ListPrivateChatMessagesReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerUserId

`func (o *ListPrivateChatMessagesReq) GetPeerUserId() string`

GetPeerUserId returns the PeerUserId field if non-nil, zero value otherwise.

### GetPeerUserIdOk

`func (o *ListPrivateChatMessagesReq) GetPeerUserIdOk() (*string, bool)`

GetPeerUserIdOk returns a tuple with the PeerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerUserId

`func (o *ListPrivateChatMessagesReq) SetPeerUserId(v string)`

SetPeerUserId sets PeerUserId field to given value.

### HasPeerUserId

`func (o *ListPrivateChatMessagesReq) HasPeerUserId() bool`

HasPeerUserId returns a boolean if a field has been set.

### GetBeforeId

`func (o *ListPrivateChatMessagesReq) GetBeforeId() string`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *ListPrivateChatMessagesReq) GetBeforeIdOk() (*string, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *ListPrivateChatMessagesReq) SetBeforeId(v string)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *ListPrivateChatMessagesReq) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetPage

`func (o *ListPrivateChatMessagesReq) GetPage() PageReq`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListPrivateChatMessagesReq) GetPageOk() (*PageReq, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListPrivateChatMessagesReq) SetPage(v PageReq)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListPrivateChatMessagesReq) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


