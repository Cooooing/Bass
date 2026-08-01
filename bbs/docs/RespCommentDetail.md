# RespCommentDetail

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

### NewRespCommentDetail

`func NewRespCommentDetail() *RespCommentDetail`

NewRespCommentDetail instantiates a new RespCommentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespCommentDetailWithDefaults

`func NewRespCommentDetailWithDefaults() *RespCommentDetail`

NewRespCommentDetailWithDefaults instantiates a new RespCommentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RespCommentDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RespCommentDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RespCommentDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RespCommentDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArticleId

`func (o *RespCommentDetail) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *RespCommentDetail) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *RespCommentDetail) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.

### HasArticleId

`func (o *RespCommentDetail) HasArticleId() bool`

HasArticleId returns a boolean if a field has been set.

### GetContent

`func (o *RespCommentDetail) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RespCommentDetail) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RespCommentDetail) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *RespCommentDetail) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentRender

`func (o *RespCommentDetail) GetContentRender() string`

GetContentRender returns the ContentRender field if non-nil, zero value otherwise.

### GetContentRenderOk

`func (o *RespCommentDetail) GetContentRenderOk() (*string, bool)`

GetContentRenderOk returns a tuple with the ContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRender

`func (o *RespCommentDetail) SetContentRender(v string)`

SetContentRender sets ContentRender field to given value.

### HasContentRender

`func (o *RespCommentDetail) HasContentRender() bool`

HasContentRender returns a boolean if a field has been set.

### GetLevel

`func (o *RespCommentDetail) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RespCommentDetail) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RespCommentDetail) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RespCommentDetail) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetParentId

`func (o *RespCommentDetail) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *RespCommentDetail) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *RespCommentDetail) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *RespCommentDetail) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReplyId

`func (o *RespCommentDetail) GetReplyId() string`

GetReplyId returns the ReplyId field if non-nil, zero value otherwise.

### GetReplyIdOk

`func (o *RespCommentDetail) GetReplyIdOk() (*string, bool)`

GetReplyIdOk returns a tuple with the ReplyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyId

`func (o *RespCommentDetail) SetReplyId(v string)`

SetReplyId sets ReplyId field to given value.

### HasReplyId

`func (o *RespCommentDetail) HasReplyId() bool`

HasReplyId returns a boolean if a field has been set.

### GetReplyCount

`func (o *RespCommentDetail) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *RespCommentDetail) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *RespCommentDetail) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *RespCommentDetail) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *RespCommentDetail) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *RespCommentDetail) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *RespCommentDetail) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *RespCommentDetail) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetThankCount

`func (o *RespCommentDetail) GetThankCount() int32`

GetThankCount returns the ThankCount field if non-nil, zero value otherwise.

### GetThankCountOk

`func (o *RespCommentDetail) GetThankCountOk() (*int32, bool)`

GetThankCountOk returns a tuple with the ThankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankCount

`func (o *RespCommentDetail) SetThankCount(v int32)`

SetThankCount sets ThankCount field to given value.

### HasThankCount

`func (o *RespCommentDetail) HasThankCount() bool`

HasThankCount returns a boolean if a field has been set.

### GetUser

`func (o *RespCommentDetail) GetUser() RespAccountProfile`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RespCommentDetail) GetUserOk() (*RespAccountProfile, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RespCommentDetail) SetUser(v RespAccountProfile)`

SetUser sets User field to given value.

### HasUser

`func (o *RespCommentDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReplyUser

`func (o *RespCommentDetail) GetReplyUser() RespAccountProfile`

GetReplyUser returns the ReplyUser field if non-nil, zero value otherwise.

### GetReplyUserOk

`func (o *RespCommentDetail) GetReplyUserOk() (*RespAccountProfile, bool)`

GetReplyUserOk returns a tuple with the ReplyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUser

`func (o *RespCommentDetail) SetReplyUser(v RespAccountProfile)`

SetReplyUser sets ReplyUser field to given value.

### HasReplyUser

`func (o *RespCommentDetail) HasReplyUser() bool`

HasReplyUser returns a boolean if a field has been set.

### GetArticle

`func (o *RespCommentDetail) GetArticle() RespArticleBrief`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *RespCommentDetail) GetArticleOk() (*RespArticleBrief, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *RespCommentDetail) SetArticle(v RespArticleBrief)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *RespCommentDetail) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetViewerActionState

`func (o *RespCommentDetail) GetViewerActionState() RespCommentViewerActionState`

GetViewerActionState returns the ViewerActionState field if non-nil, zero value otherwise.

### GetViewerActionStateOk

`func (o *RespCommentDetail) GetViewerActionStateOk() (*RespCommentViewerActionState, bool)`

GetViewerActionStateOk returns a tuple with the ViewerActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerActionState

`func (o *RespCommentDetail) SetViewerActionState(v RespCommentViewerActionState)`

SetViewerActionState sets ViewerActionState field to given value.

### HasViewerActionState

`func (o *RespCommentDetail) HasViewerActionState() bool`

HasViewerActionState returns a boolean if a field has been set.

### GetRestriction

`func (o *RespCommentDetail) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *RespCommentDetail) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *RespCommentDetail) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *RespCommentDetail) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetDeletedAt

`func (o *RespCommentDetail) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *RespCommentDetail) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *RespCommentDetail) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *RespCommentDetail) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RespCommentDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RespCommentDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RespCommentDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RespCommentDetail) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RespCommentDetail) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RespCommentDetail) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RespCommentDetail) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RespCommentDetail) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RespCommentDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespCommentDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespCommentDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespCommentDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RespCommentDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RespCommentDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RespCommentDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RespCommentDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


