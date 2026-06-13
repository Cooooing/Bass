# CommentThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Root** | Pointer to [**CommentListItem**](CommentListItem.md) |  | [optional] 
**PreviewReplies** | Pointer to [**[]CommentListItem**](CommentListItem.md) |  | [optional] 
**ReplyCount** | Pointer to **int32** |  | [optional] 
**HasMoreReplies** | Pointer to **bool** |  | [optional] 

## Methods

### NewCommentThread

`func NewCommentThread() *CommentThread`

NewCommentThread instantiates a new CommentThread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentThreadWithDefaults

`func NewCommentThreadWithDefaults() *CommentThread`

NewCommentThreadWithDefaults instantiates a new CommentThread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoot

`func (o *CommentThread) GetRoot() CommentListItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *CommentThread) GetRootOk() (*CommentListItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *CommentThread) SetRoot(v CommentListItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *CommentThread) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetPreviewReplies

`func (o *CommentThread) GetPreviewReplies() []CommentListItem`

GetPreviewReplies returns the PreviewReplies field if non-nil, zero value otherwise.

### GetPreviewRepliesOk

`func (o *CommentThread) GetPreviewRepliesOk() (*[]CommentListItem, bool)`

GetPreviewRepliesOk returns a tuple with the PreviewReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewReplies

`func (o *CommentThread) SetPreviewReplies(v []CommentListItem)`

SetPreviewReplies sets PreviewReplies field to given value.

### HasPreviewReplies

`func (o *CommentThread) HasPreviewReplies() bool`

HasPreviewReplies returns a boolean if a field has been set.

### GetReplyCount

`func (o *CommentThread) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *CommentThread) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *CommentThread) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *CommentThread) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetHasMoreReplies

`func (o *CommentThread) GetHasMoreReplies() bool`

GetHasMoreReplies returns the HasMoreReplies field if non-nil, zero value otherwise.

### GetHasMoreRepliesOk

`func (o *CommentThread) GetHasMoreRepliesOk() (*bool, bool)`

GetHasMoreRepliesOk returns a tuple with the HasMoreReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreReplies

`func (o *CommentThread) SetHasMoreReplies(v bool)`

SetHasMoreReplies sets HasMoreReplies field to given value.

### HasHasMoreReplies

`func (o *CommentThread) HasHasMoreReplies() bool`

HasHasMoreReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


