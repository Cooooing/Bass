# RespCommentThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Root** | Pointer to [**RespCommentListItem**](RespCommentListItem.md) |  | [optional] 
**PreviewReplies** | Pointer to [**[]RespCommentListItem**](RespCommentListItem.md) |  | [optional] 
**ReplyCount** | Pointer to **int32** |  | [optional] 
**HasMoreReplies** | Pointer to **bool** |  | [optional] 

## Methods

### NewRespCommentThread

`func NewRespCommentThread() *RespCommentThread`

NewRespCommentThread instantiates a new RespCommentThread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespCommentThreadWithDefaults

`func NewRespCommentThreadWithDefaults() *RespCommentThread`

NewRespCommentThreadWithDefaults instantiates a new RespCommentThread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoot

`func (o *RespCommentThread) GetRoot() RespCommentListItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *RespCommentThread) GetRootOk() (*RespCommentListItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *RespCommentThread) SetRoot(v RespCommentListItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *RespCommentThread) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetPreviewReplies

`func (o *RespCommentThread) GetPreviewReplies() []RespCommentListItem`

GetPreviewReplies returns the PreviewReplies field if non-nil, zero value otherwise.

### GetPreviewRepliesOk

`func (o *RespCommentThread) GetPreviewRepliesOk() (*[]RespCommentListItem, bool)`

GetPreviewRepliesOk returns a tuple with the PreviewReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewReplies

`func (o *RespCommentThread) SetPreviewReplies(v []RespCommentListItem)`

SetPreviewReplies sets PreviewReplies field to given value.

### HasPreviewReplies

`func (o *RespCommentThread) HasPreviewReplies() bool`

HasPreviewReplies returns a boolean if a field has been set.

### GetReplyCount

`func (o *RespCommentThread) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *RespCommentThread) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *RespCommentThread) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *RespCommentThread) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetHasMoreReplies

`func (o *RespCommentThread) GetHasMoreReplies() bool`

GetHasMoreReplies returns the HasMoreReplies field if non-nil, zero value otherwise.

### GetHasMoreRepliesOk

`func (o *RespCommentThread) GetHasMoreRepliesOk() (*bool, bool)`

GetHasMoreRepliesOk returns a tuple with the HasMoreReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreReplies

`func (o *RespCommentThread) SetHasMoreReplies(v bool)`

SetHasMoreReplies sets HasMoreReplies field to given value.

### HasHasMoreReplies

`func (o *RespCommentThread) HasHasMoreReplies() bool`

HasHasMoreReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


