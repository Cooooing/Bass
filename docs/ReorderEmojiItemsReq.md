# ReorderEmojiItemsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | 可选目标分组 ID；为空时重排未分组表情。 | [optional] 
**ItemIds** | Pointer to **[]string** | 按最终展示顺序排列的完整表情项 ID 序列。 | [optional] 

## Methods

### NewReorderEmojiItemsReq

`func NewReorderEmojiItemsReq() *ReorderEmojiItemsReq`

NewReorderEmojiItemsReq instantiates a new ReorderEmojiItemsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReorderEmojiItemsReqWithDefaults

`func NewReorderEmojiItemsReqWithDefaults() *ReorderEmojiItemsReq`

NewReorderEmojiItemsReqWithDefaults instantiates a new ReorderEmojiItemsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ReorderEmojiItemsReq) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ReorderEmojiItemsReq) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ReorderEmojiItemsReq) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ReorderEmojiItemsReq) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetItemIds

`func (o *ReorderEmojiItemsReq) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *ReorderEmojiItemsReq) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *ReorderEmojiItemsReq) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *ReorderEmojiItemsReq) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


