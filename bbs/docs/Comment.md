# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | 评论 ID。 | [optional] 
**ArticleId** | Pointer to **string** | 文章 ID。 | [optional] 
**Content** | Pointer to **string** | 原始内容。 | [optional] 
**ContentRender** | Pointer to **string** | 渲染后的内容。 | [optional] 
**Level** | Pointer to **int32** | 评论层级。 | [optional] 
**ParentId** | Pointer to **string** | 父评论 ID。 | [optional] 
**ReplyId** | Pointer to **string** | 回复的评论 ID。 | [optional] 
**Status** | Pointer to **string** | 评论状态。 | [optional] 
**ReplyCount** | Pointer to **int32** | 回复数量。 | [optional] 
**LikeCount** | Pointer to **int32** | 点赞数量。 | [optional] 
**ThankCount** | Pointer to **int32** | 感谢数量。 | [optional] 
**User** | Pointer to [**AccountProfile**](AccountProfile.md) | 评论账号展示资料。 | [optional] 
**ReplyUser** | Pointer to [**AccountProfile**](AccountProfile.md) | 被回复账号展示资料。 | [optional] 
**Article** | Pointer to [**Article**](Article.md) | 所属文章。 | [optional] 
**CreatedBy** | Pointer to **string** | 创建账号 ID。 | [optional] 
**UpdatedBy** | Pointer to **string** | 更新账号 ID。 | [optional] 
**CreatedAt** | Pointer to **string** | 创建时间。 | [optional] 
**UpdatedAt** | Pointer to **string** | 更新时间。 | [optional] 

## Methods

### NewComment

`func NewComment() *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Comment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArticleId

`func (o *Comment) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *Comment) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *Comment) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.

### HasArticleId

`func (o *Comment) HasArticleId() bool`

HasArticleId returns a boolean if a field has been set.

### GetContent

`func (o *Comment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Comment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Comment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Comment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentRender

`func (o *Comment) GetContentRender() string`

GetContentRender returns the ContentRender field if non-nil, zero value otherwise.

### GetContentRenderOk

`func (o *Comment) GetContentRenderOk() (*string, bool)`

GetContentRenderOk returns a tuple with the ContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRender

`func (o *Comment) SetContentRender(v string)`

SetContentRender sets ContentRender field to given value.

### HasContentRender

`func (o *Comment) HasContentRender() bool`

HasContentRender returns a boolean if a field has been set.

### GetLevel

`func (o *Comment) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Comment) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Comment) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Comment) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetParentId

`func (o *Comment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Comment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Comment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Comment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReplyId

`func (o *Comment) GetReplyId() string`

GetReplyId returns the ReplyId field if non-nil, zero value otherwise.

### GetReplyIdOk

`func (o *Comment) GetReplyIdOk() (*string, bool)`

GetReplyIdOk returns a tuple with the ReplyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyId

`func (o *Comment) SetReplyId(v string)`

SetReplyId sets ReplyId field to given value.

### HasReplyId

`func (o *Comment) HasReplyId() bool`

HasReplyId returns a boolean if a field has been set.

### GetStatus

`func (o *Comment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Comment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Comment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Comment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReplyCount

`func (o *Comment) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *Comment) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *Comment) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *Comment) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *Comment) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *Comment) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *Comment) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *Comment) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetThankCount

`func (o *Comment) GetThankCount() int32`

GetThankCount returns the ThankCount field if non-nil, zero value otherwise.

### GetThankCountOk

`func (o *Comment) GetThankCountOk() (*int32, bool)`

GetThankCountOk returns a tuple with the ThankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankCount

`func (o *Comment) SetThankCount(v int32)`

SetThankCount sets ThankCount field to given value.

### HasThankCount

`func (o *Comment) HasThankCount() bool`

HasThankCount returns a boolean if a field has been set.

### GetUser

`func (o *Comment) GetUser() AccountProfile`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Comment) GetUserOk() (*AccountProfile, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Comment) SetUser(v AccountProfile)`

SetUser sets User field to given value.

### HasUser

`func (o *Comment) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReplyUser

`func (o *Comment) GetReplyUser() AccountProfile`

GetReplyUser returns the ReplyUser field if non-nil, zero value otherwise.

### GetReplyUserOk

`func (o *Comment) GetReplyUserOk() (*AccountProfile, bool)`

GetReplyUserOk returns a tuple with the ReplyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUser

`func (o *Comment) SetReplyUser(v AccountProfile)`

SetReplyUser sets ReplyUser field to given value.

### HasReplyUser

`func (o *Comment) HasReplyUser() bool`

HasReplyUser returns a boolean if a field has been set.

### GetArticle

`func (o *Comment) GetArticle() Article`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *Comment) GetArticleOk() (*Article, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *Comment) SetArticle(v Article)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *Comment) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Comment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Comment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Comment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Comment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Comment) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Comment) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Comment) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Comment) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Comment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Comment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Comment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Comment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Comment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Comment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Comment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Comment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


