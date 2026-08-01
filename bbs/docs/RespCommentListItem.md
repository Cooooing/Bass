# RespCommentListItem

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
**User** | Pointer to [**RespAccountProfile**](RespAccountProfile.md) |  | [optional] 
**ReplyUser** | Pointer to [**RespAccountProfile**](RespAccountProfile.md) |  | [optional] 
**Article** | Pointer to [**RespArticleBrief**](RespArticleBrief.md) |  | [optional] 
**ViewerActionState** | Pointer to [**RespCommentViewerActionState**](RespCommentViewerActionState.md) |  | [optional] 
**Restriction** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRespCommentListItem

`func NewRespCommentListItem() *RespCommentListItem`

NewRespCommentListItem instantiates a new RespCommentListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespCommentListItemWithDefaults

`func NewRespCommentListItemWithDefaults() *RespCommentListItem`

NewRespCommentListItemWithDefaults instantiates a new RespCommentListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RespCommentListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RespCommentListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RespCommentListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RespCommentListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArticleId

`func (o *RespCommentListItem) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *RespCommentListItem) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *RespCommentListItem) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.

### HasArticleId

`func (o *RespCommentListItem) HasArticleId() bool`

HasArticleId returns a boolean if a field has been set.

### GetContent

`func (o *RespCommentListItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RespCommentListItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RespCommentListItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *RespCommentListItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentRender

`func (o *RespCommentListItem) GetContentRender() string`

GetContentRender returns the ContentRender field if non-nil, zero value otherwise.

### GetContentRenderOk

`func (o *RespCommentListItem) GetContentRenderOk() (*string, bool)`

GetContentRenderOk returns a tuple with the ContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRender

`func (o *RespCommentListItem) SetContentRender(v string)`

SetContentRender sets ContentRender field to given value.

### HasContentRender

`func (o *RespCommentListItem) HasContentRender() bool`

HasContentRender returns a boolean if a field has been set.

### GetLevel

`func (o *RespCommentListItem) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RespCommentListItem) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RespCommentListItem) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RespCommentListItem) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetParentId

`func (o *RespCommentListItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *RespCommentListItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *RespCommentListItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *RespCommentListItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReplyId

`func (o *RespCommentListItem) GetReplyId() string`

GetReplyId returns the ReplyId field if non-nil, zero value otherwise.

### GetReplyIdOk

`func (o *RespCommentListItem) GetReplyIdOk() (*string, bool)`

GetReplyIdOk returns a tuple with the ReplyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyId

`func (o *RespCommentListItem) SetReplyId(v string)`

SetReplyId sets ReplyId field to given value.

### HasReplyId

`func (o *RespCommentListItem) HasReplyId() bool`

HasReplyId returns a boolean if a field has been set.

### GetReplyCount

`func (o *RespCommentListItem) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *RespCommentListItem) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *RespCommentListItem) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *RespCommentListItem) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *RespCommentListItem) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *RespCommentListItem) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *RespCommentListItem) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *RespCommentListItem) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetThankCount

`func (o *RespCommentListItem) GetThankCount() int32`

GetThankCount returns the ThankCount field if non-nil, zero value otherwise.

### GetThankCountOk

`func (o *RespCommentListItem) GetThankCountOk() (*int32, bool)`

GetThankCountOk returns a tuple with the ThankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankCount

`func (o *RespCommentListItem) SetThankCount(v int32)`

SetThankCount sets ThankCount field to given value.

### HasThankCount

`func (o *RespCommentListItem) HasThankCount() bool`

HasThankCount returns a boolean if a field has been set.

### GetUser

`func (o *RespCommentListItem) GetUser() RespAccountProfile`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RespCommentListItem) GetUserOk() (*RespAccountProfile, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RespCommentListItem) SetUser(v RespAccountProfile)`

SetUser sets User field to given value.

### HasUser

`func (o *RespCommentListItem) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReplyUser

`func (o *RespCommentListItem) GetReplyUser() RespAccountProfile`

GetReplyUser returns the ReplyUser field if non-nil, zero value otherwise.

### GetReplyUserOk

`func (o *RespCommentListItem) GetReplyUserOk() (*RespAccountProfile, bool)`

GetReplyUserOk returns a tuple with the ReplyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUser

`func (o *RespCommentListItem) SetReplyUser(v RespAccountProfile)`

SetReplyUser sets ReplyUser field to given value.

### HasReplyUser

`func (o *RespCommentListItem) HasReplyUser() bool`

HasReplyUser returns a boolean if a field has been set.

### GetArticle

`func (o *RespCommentListItem) GetArticle() RespArticleBrief`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *RespCommentListItem) GetArticleOk() (*RespArticleBrief, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *RespCommentListItem) SetArticle(v RespArticleBrief)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *RespCommentListItem) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetViewerActionState

`func (o *RespCommentListItem) GetViewerActionState() RespCommentViewerActionState`

GetViewerActionState returns the ViewerActionState field if non-nil, zero value otherwise.

### GetViewerActionStateOk

`func (o *RespCommentListItem) GetViewerActionStateOk() (*RespCommentViewerActionState, bool)`

GetViewerActionStateOk returns a tuple with the ViewerActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerActionState

`func (o *RespCommentListItem) SetViewerActionState(v RespCommentViewerActionState)`

SetViewerActionState sets ViewerActionState field to given value.

### HasViewerActionState

`func (o *RespCommentListItem) HasViewerActionState() bool`

HasViewerActionState returns a boolean if a field has been set.

### GetRestriction

`func (o *RespCommentListItem) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *RespCommentListItem) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *RespCommentListItem) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *RespCommentListItem) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RespCommentListItem) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RespCommentListItem) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RespCommentListItem) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RespCommentListItem) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RespCommentListItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RespCommentListItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RespCommentListItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RespCommentListItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RespCommentListItem) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RespCommentListItem) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RespCommentListItem) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RespCommentListItem) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RespCommentListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespCommentListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespCommentListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespCommentListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RespCommentListItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RespCommentListItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RespCommentListItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RespCommentListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


