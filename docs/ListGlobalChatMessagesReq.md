# ListGlobalChatMessagesReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeforeId** | Pointer to **string** |  | [optional] 
**Page** | Pointer to [**PageReq**](PageReq.md) |  | [optional] 

## Methods

### NewListGlobalChatMessagesReq

`func NewListGlobalChatMessagesReq() *ListGlobalChatMessagesReq`

NewListGlobalChatMessagesReq instantiates a new ListGlobalChatMessagesReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGlobalChatMessagesReqWithDefaults

`func NewListGlobalChatMessagesReqWithDefaults() *ListGlobalChatMessagesReq`

NewListGlobalChatMessagesReqWithDefaults instantiates a new ListGlobalChatMessagesReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeforeId

`func (o *ListGlobalChatMessagesReq) GetBeforeId() string`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *ListGlobalChatMessagesReq) GetBeforeIdOk() (*string, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *ListGlobalChatMessagesReq) SetBeforeId(v string)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *ListGlobalChatMessagesReq) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetPage

`func (o *ListGlobalChatMessagesReq) GetPage() PageReq`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListGlobalChatMessagesReq) GetPageOk() (*PageReq, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListGlobalChatMessagesReq) SetPage(v PageReq)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListGlobalChatMessagesReq) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


