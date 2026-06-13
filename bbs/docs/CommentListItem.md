# CommentListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ArticleId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ContentRender** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ReplyId** | Pointer to **string** |  | [optional] 
**ReplyCount** | Pointer to **int32** |  | [optional] 
**LikeCount** | Pointer to **int32** |  | [optional] 
**ThankCount** | Pointer to **int32** |  | [optional] 
**User** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**ReplyUser** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**Article** | Pointer to [**ArticleBrief**](ArticleBrief.md) |  | [optional] 
**ViewerActionState** | Pointer to [**CommentViewerActionState**](CommentViewerActionState.md) |  | [optional] 
**Restriction** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCommentListItem

`func NewCommentListItem() *CommentListItem`

NewCommentListItem instantiates a new CommentListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentListItemWithDefaults

`func NewCommentListItemWithDefaults() *CommentListItem`

NewCommentListItemWithDefaults instantiates a new CommentListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommentListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArticleId

`func (o *CommentListItem) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *CommentListItem) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *CommentListItem) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.

### HasArticleId

`func (o *CommentListItem) HasArticleId() bool`

HasArticleId returns a boolean if a field has been set.

### GetContent

`func (o *CommentListItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CommentListItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CommentListItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CommentListItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentRender

`func (o *CommentListItem) GetContentRender() string`

GetContentRender returns the ContentRender field if non-nil, zero value otherwise.

### GetContentRenderOk

`func (o *CommentListItem) GetContentRenderOk() (*string, bool)`

GetContentRenderOk returns a tuple with the ContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRender

`func (o *CommentListItem) SetContentRender(v string)`

SetContentRender sets ContentRender field to given value.

### HasContentRender

`func (o *CommentListItem) HasContentRender() bool`

HasContentRender returns a boolean if a field has been set.

### GetLevel

`func (o *CommentListItem) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CommentListItem) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CommentListItem) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CommentListItem) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetParentId

`func (o *CommentListItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CommentListItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CommentListItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CommentListItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReplyId

`func (o *CommentListItem) GetReplyId() string`

GetReplyId returns the ReplyId field if non-nil, zero value otherwise.

### GetReplyIdOk

`func (o *CommentListItem) GetReplyIdOk() (*string, bool)`

GetReplyIdOk returns a tuple with the ReplyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyId

`func (o *CommentListItem) SetReplyId(v string)`

SetReplyId sets ReplyId field to given value.

### HasReplyId

`func (o *CommentListItem) HasReplyId() bool`

HasReplyId returns a boolean if a field has been set.

### GetReplyCount

`func (o *CommentListItem) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *CommentListItem) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *CommentListItem) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *CommentListItem) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *CommentListItem) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *CommentListItem) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *CommentListItem) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *CommentListItem) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetThankCount

`func (o *CommentListItem) GetThankCount() int32`

GetThankCount returns the ThankCount field if non-nil, zero value otherwise.

### GetThankCountOk

`func (o *CommentListItem) GetThankCountOk() (*int32, bool)`

GetThankCountOk returns a tuple with the ThankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankCount

`func (o *CommentListItem) SetThankCount(v int32)`

SetThankCount sets ThankCount field to given value.

### HasThankCount

`func (o *CommentListItem) HasThankCount() bool`

HasThankCount returns a boolean if a field has been set.

### GetUser

`func (o *CommentListItem) GetUser() AccountProfile`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CommentListItem) GetUserOk() (*AccountProfile, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CommentListItem) SetUser(v AccountProfile)`

SetUser sets User field to given value.

### HasUser

`func (o *CommentListItem) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReplyUser

`func (o *CommentListItem) GetReplyUser() AccountProfile`

GetReplyUser returns the ReplyUser field if non-nil, zero value otherwise.

### GetReplyUserOk

`func (o *CommentListItem) GetReplyUserOk() (*AccountProfile, bool)`

GetReplyUserOk returns a tuple with the ReplyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUser

`func (o *CommentListItem) SetReplyUser(v AccountProfile)`

SetReplyUser sets ReplyUser field to given value.

### HasReplyUser

`func (o *CommentListItem) HasReplyUser() bool`

HasReplyUser returns a boolean if a field has been set.

### GetArticle

`func (o *CommentListItem) GetArticle() ArticleBrief`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *CommentListItem) GetArticleOk() (*ArticleBrief, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *CommentListItem) SetArticle(v ArticleBrief)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *CommentListItem) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetViewerActionState

`func (o *CommentListItem) GetViewerActionState() CommentViewerActionState`

GetViewerActionState returns the ViewerActionState field if non-nil, zero value otherwise.

### GetViewerActionStateOk

`func (o *CommentListItem) GetViewerActionStateOk() (*CommentViewerActionState, bool)`

GetViewerActionStateOk returns a tuple with the ViewerActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerActionState

`func (o *CommentListItem) SetViewerActionState(v CommentViewerActionState)`

SetViewerActionState sets ViewerActionState field to given value.

### HasViewerActionState

`func (o *CommentListItem) HasViewerActionState() bool`

HasViewerActionState returns a boolean if a field has been set.

### GetRestriction

`func (o *CommentListItem) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *CommentListItem) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *CommentListItem) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *CommentListItem) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CommentListItem) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CommentListItem) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CommentListItem) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CommentListItem) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CommentListItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CommentListItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CommentListItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CommentListItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CommentListItem) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CommentListItem) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CommentListItem) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CommentListItem) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommentListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommentListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommentListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommentListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommentListItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommentListItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommentListItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommentListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


