# CommentDetail

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

### NewCommentDetail

`func NewCommentDetail() *CommentDetail`

NewCommentDetail instantiates a new CommentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentDetailWithDefaults

`func NewCommentDetailWithDefaults() *CommentDetail`

NewCommentDetailWithDefaults instantiates a new CommentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommentDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArticleId

`func (o *CommentDetail) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *CommentDetail) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *CommentDetail) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.

### HasArticleId

`func (o *CommentDetail) HasArticleId() bool`

HasArticleId returns a boolean if a field has been set.

### GetContent

`func (o *CommentDetail) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CommentDetail) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CommentDetail) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CommentDetail) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentRender

`func (o *CommentDetail) GetContentRender() string`

GetContentRender returns the ContentRender field if non-nil, zero value otherwise.

### GetContentRenderOk

`func (o *CommentDetail) GetContentRenderOk() (*string, bool)`

GetContentRenderOk returns a tuple with the ContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRender

`func (o *CommentDetail) SetContentRender(v string)`

SetContentRender sets ContentRender field to given value.

### HasContentRender

`func (o *CommentDetail) HasContentRender() bool`

HasContentRender returns a boolean if a field has been set.

### GetLevel

`func (o *CommentDetail) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CommentDetail) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CommentDetail) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CommentDetail) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetParentId

`func (o *CommentDetail) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CommentDetail) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CommentDetail) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CommentDetail) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReplyId

`func (o *CommentDetail) GetReplyId() string`

GetReplyId returns the ReplyId field if non-nil, zero value otherwise.

### GetReplyIdOk

`func (o *CommentDetail) GetReplyIdOk() (*string, bool)`

GetReplyIdOk returns a tuple with the ReplyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyId

`func (o *CommentDetail) SetReplyId(v string)`

SetReplyId sets ReplyId field to given value.

### HasReplyId

`func (o *CommentDetail) HasReplyId() bool`

HasReplyId returns a boolean if a field has been set.

### GetReplyCount

`func (o *CommentDetail) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *CommentDetail) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *CommentDetail) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *CommentDetail) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *CommentDetail) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *CommentDetail) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *CommentDetail) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *CommentDetail) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetThankCount

`func (o *CommentDetail) GetThankCount() int32`

GetThankCount returns the ThankCount field if non-nil, zero value otherwise.

### GetThankCountOk

`func (o *CommentDetail) GetThankCountOk() (*int32, bool)`

GetThankCountOk returns a tuple with the ThankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankCount

`func (o *CommentDetail) SetThankCount(v int32)`

SetThankCount sets ThankCount field to given value.

### HasThankCount

`func (o *CommentDetail) HasThankCount() bool`

HasThankCount returns a boolean if a field has been set.

### GetUser

`func (o *CommentDetail) GetUser() AccountProfile`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CommentDetail) GetUserOk() (*AccountProfile, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CommentDetail) SetUser(v AccountProfile)`

SetUser sets User field to given value.

### HasUser

`func (o *CommentDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReplyUser

`func (o *CommentDetail) GetReplyUser() AccountProfile`

GetReplyUser returns the ReplyUser field if non-nil, zero value otherwise.

### GetReplyUserOk

`func (o *CommentDetail) GetReplyUserOk() (*AccountProfile, bool)`

GetReplyUserOk returns a tuple with the ReplyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUser

`func (o *CommentDetail) SetReplyUser(v AccountProfile)`

SetReplyUser sets ReplyUser field to given value.

### HasReplyUser

`func (o *CommentDetail) HasReplyUser() bool`

HasReplyUser returns a boolean if a field has been set.

### GetArticle

`func (o *CommentDetail) GetArticle() ArticleBrief`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *CommentDetail) GetArticleOk() (*ArticleBrief, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *CommentDetail) SetArticle(v ArticleBrief)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *CommentDetail) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetViewerActionState

`func (o *CommentDetail) GetViewerActionState() CommentViewerActionState`

GetViewerActionState returns the ViewerActionState field if non-nil, zero value otherwise.

### GetViewerActionStateOk

`func (o *CommentDetail) GetViewerActionStateOk() (*CommentViewerActionState, bool)`

GetViewerActionStateOk returns a tuple with the ViewerActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerActionState

`func (o *CommentDetail) SetViewerActionState(v CommentViewerActionState)`

SetViewerActionState sets ViewerActionState field to given value.

### HasViewerActionState

`func (o *CommentDetail) HasViewerActionState() bool`

HasViewerActionState returns a boolean if a field has been set.

### GetRestriction

`func (o *CommentDetail) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *CommentDetail) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *CommentDetail) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *CommentDetail) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CommentDetail) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CommentDetail) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CommentDetail) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CommentDetail) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CommentDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CommentDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CommentDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CommentDetail) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CommentDetail) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CommentDetail) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CommentDetail) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CommentDetail) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommentDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommentDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommentDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommentDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommentDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommentDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommentDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommentDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


