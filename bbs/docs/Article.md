# Article

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | 文章 ID。 | [optional] 
**Title** | Pointer to **string** | 标题。 | [optional] 
**Content** | Pointer to **string** | 原始内容。 | [optional] 
**ContentRender** | Pointer to **string** | 渲染后的内容。 | [optional] 
**HasPostscript** | Pointer to **bool** | 是否有附言。 | [optional] 
**HasReward** | Pointer to **bool** | 是否有打赏内容。 | [optional] 
**RewardContent** | Pointer to **string** | 打赏后可见内容。 | [optional] 
**RewardContentRender** | Pointer to **string** | 渲染后的打赏内容。 | [optional] 
**RewardPoints** | Pointer to **int32** | 打赏所需积分。 | [optional] 
**Status** | Pointer to **string** | 文章状态。 | [optional] 
**Type** | Pointer to **string** | 文章类型。 | [optional] 
**Statement** | Pointer to **string** | 文章声明。 | [optional] 
**Commentable** | Pointer to **bool** | 是否允许评论。 | [optional] 
**Anonymous** | Pointer to **bool** | 是否匿名展示。 | [optional] 
**Listable** | Pointer to **bool** | 是否在列表中展示。 | [optional] 
**ViewCount** | Pointer to **int32** | 浏览数量。 | [optional] 
**ThankCount** | Pointer to **int32** | 感谢数量。 | [optional] 
**LikeCount** | Pointer to **int32** | 点赞数量。 | [optional] 
**CollectCount** | Pointer to **int32** | 收藏数量。 | [optional] 
**WatchCount** | Pointer to **int32** | 关注数量。 | [optional] 
**ReplyCount** | Pointer to **int32** | 回复数量。 | [optional] 
**BountyPoints** | Pointer to **int32** | 悬赏积分。 | [optional] 
**AcceptedAnswerId** | Pointer to **string** | 已采纳答案评论 ID。 | [optional] 
**AuthorUser** | Pointer to [**AccountProfile**](AccountProfile.md) | 作者账号展示资料。 | [optional] 
**LastReplyUser** | Pointer to [**AccountProfile**](AccountProfile.md) | 最后回复账号展示资料。 | [optional] 
**LastReplyAt** | Pointer to **string** | 最后回复时间。 | [optional] 
**CoverImageUrl** | Pointer to **string** | 封面图片 URL。 | [optional] 
**CreatedBy** | Pointer to **string** | 创建账号 ID。 | [optional] 
**UpdatedBy** | Pointer to **string** | 更新账号 ID。 | [optional] 
**CreatedAt** | Pointer to **string** | 创建时间。 | [optional] 
**UpdatedAt** | Pointer to **string** | 更新时间。 | [optional] 

## Methods

### NewArticle

`func NewArticle() *Article`

NewArticle instantiates a new Article object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleWithDefaults

`func NewArticleWithDefaults() *Article`

NewArticleWithDefaults instantiates a new Article object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Article) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Article) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Article) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Article) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Article) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Article) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Article) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Article) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *Article) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Article) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Article) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Article) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentRender

`func (o *Article) GetContentRender() string`

GetContentRender returns the ContentRender field if non-nil, zero value otherwise.

### GetContentRenderOk

`func (o *Article) GetContentRenderOk() (*string, bool)`

GetContentRenderOk returns a tuple with the ContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRender

`func (o *Article) SetContentRender(v string)`

SetContentRender sets ContentRender field to given value.

### HasContentRender

`func (o *Article) HasContentRender() bool`

HasContentRender returns a boolean if a field has been set.

### GetHasPostscript

`func (o *Article) GetHasPostscript() bool`

GetHasPostscript returns the HasPostscript field if non-nil, zero value otherwise.

### GetHasPostscriptOk

`func (o *Article) GetHasPostscriptOk() (*bool, bool)`

GetHasPostscriptOk returns a tuple with the HasPostscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPostscript

`func (o *Article) SetHasPostscript(v bool)`

SetHasPostscript sets HasPostscript field to given value.

### HasHasPostscript

`func (o *Article) HasHasPostscript() bool`

HasHasPostscript returns a boolean if a field has been set.

### GetHasReward

`func (o *Article) GetHasReward() bool`

GetHasReward returns the HasReward field if non-nil, zero value otherwise.

### GetHasRewardOk

`func (o *Article) GetHasRewardOk() (*bool, bool)`

GetHasRewardOk returns a tuple with the HasReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReward

`func (o *Article) SetHasReward(v bool)`

SetHasReward sets HasReward field to given value.

### HasHasReward

`func (o *Article) HasHasReward() bool`

HasHasReward returns a boolean if a field has been set.

### GetRewardContent

`func (o *Article) GetRewardContent() string`

GetRewardContent returns the RewardContent field if non-nil, zero value otherwise.

### GetRewardContentOk

`func (o *Article) GetRewardContentOk() (*string, bool)`

GetRewardContentOk returns a tuple with the RewardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardContent

`func (o *Article) SetRewardContent(v string)`

SetRewardContent sets RewardContent field to given value.

### HasRewardContent

`func (o *Article) HasRewardContent() bool`

HasRewardContent returns a boolean if a field has been set.

### GetRewardContentRender

`func (o *Article) GetRewardContentRender() string`

GetRewardContentRender returns the RewardContentRender field if non-nil, zero value otherwise.

### GetRewardContentRenderOk

`func (o *Article) GetRewardContentRenderOk() (*string, bool)`

GetRewardContentRenderOk returns a tuple with the RewardContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardContentRender

`func (o *Article) SetRewardContentRender(v string)`

SetRewardContentRender sets RewardContentRender field to given value.

### HasRewardContentRender

`func (o *Article) HasRewardContentRender() bool`

HasRewardContentRender returns a boolean if a field has been set.

### GetRewardPoints

`func (o *Article) GetRewardPoints() int32`

GetRewardPoints returns the RewardPoints field if non-nil, zero value otherwise.

### GetRewardPointsOk

`func (o *Article) GetRewardPointsOk() (*int32, bool)`

GetRewardPointsOk returns a tuple with the RewardPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPoints

`func (o *Article) SetRewardPoints(v int32)`

SetRewardPoints sets RewardPoints field to given value.

### HasRewardPoints

`func (o *Article) HasRewardPoints() bool`

HasRewardPoints returns a boolean if a field has been set.

### GetStatus

`func (o *Article) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Article) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Article) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Article) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Article) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Article) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Article) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Article) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatement

`func (o *Article) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *Article) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *Article) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *Article) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetCommentable

`func (o *Article) GetCommentable() bool`

GetCommentable returns the Commentable field if non-nil, zero value otherwise.

### GetCommentableOk

`func (o *Article) GetCommentableOk() (*bool, bool)`

GetCommentableOk returns a tuple with the Commentable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentable

`func (o *Article) SetCommentable(v bool)`

SetCommentable sets Commentable field to given value.

### HasCommentable

`func (o *Article) HasCommentable() bool`

HasCommentable returns a boolean if a field has been set.

### GetAnonymous

`func (o *Article) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *Article) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *Article) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *Article) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetListable

`func (o *Article) GetListable() bool`

GetListable returns the Listable field if non-nil, zero value otherwise.

### GetListableOk

`func (o *Article) GetListableOk() (*bool, bool)`

GetListableOk returns a tuple with the Listable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListable

`func (o *Article) SetListable(v bool)`

SetListable sets Listable field to given value.

### HasListable

`func (o *Article) HasListable() bool`

HasListable returns a boolean if a field has been set.

### GetViewCount

`func (o *Article) GetViewCount() int32`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *Article) GetViewCountOk() (*int32, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *Article) SetViewCount(v int32)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *Article) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetThankCount

`func (o *Article) GetThankCount() int32`

GetThankCount returns the ThankCount field if non-nil, zero value otherwise.

### GetThankCountOk

`func (o *Article) GetThankCountOk() (*int32, bool)`

GetThankCountOk returns a tuple with the ThankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankCount

`func (o *Article) SetThankCount(v int32)`

SetThankCount sets ThankCount field to given value.

### HasThankCount

`func (o *Article) HasThankCount() bool`

HasThankCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *Article) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *Article) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *Article) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *Article) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetCollectCount

`func (o *Article) GetCollectCount() int32`

GetCollectCount returns the CollectCount field if non-nil, zero value otherwise.

### GetCollectCountOk

`func (o *Article) GetCollectCountOk() (*int32, bool)`

GetCollectCountOk returns a tuple with the CollectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectCount

`func (o *Article) SetCollectCount(v int32)`

SetCollectCount sets CollectCount field to given value.

### HasCollectCount

`func (o *Article) HasCollectCount() bool`

HasCollectCount returns a boolean if a field has been set.

### GetWatchCount

`func (o *Article) GetWatchCount() int32`

GetWatchCount returns the WatchCount field if non-nil, zero value otherwise.

### GetWatchCountOk

`func (o *Article) GetWatchCountOk() (*int32, bool)`

GetWatchCountOk returns a tuple with the WatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchCount

`func (o *Article) SetWatchCount(v int32)`

SetWatchCount sets WatchCount field to given value.

### HasWatchCount

`func (o *Article) HasWatchCount() bool`

HasWatchCount returns a boolean if a field has been set.

### GetReplyCount

`func (o *Article) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *Article) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *Article) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *Article) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetBountyPoints

`func (o *Article) GetBountyPoints() int32`

GetBountyPoints returns the BountyPoints field if non-nil, zero value otherwise.

### GetBountyPointsOk

`func (o *Article) GetBountyPointsOk() (*int32, bool)`

GetBountyPointsOk returns a tuple with the BountyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBountyPoints

`func (o *Article) SetBountyPoints(v int32)`

SetBountyPoints sets BountyPoints field to given value.

### HasBountyPoints

`func (o *Article) HasBountyPoints() bool`

HasBountyPoints returns a boolean if a field has been set.

### GetAcceptedAnswerId

`func (o *Article) GetAcceptedAnswerId() string`

GetAcceptedAnswerId returns the AcceptedAnswerId field if non-nil, zero value otherwise.

### GetAcceptedAnswerIdOk

`func (o *Article) GetAcceptedAnswerIdOk() (*string, bool)`

GetAcceptedAnswerIdOk returns a tuple with the AcceptedAnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAnswerId

`func (o *Article) SetAcceptedAnswerId(v string)`

SetAcceptedAnswerId sets AcceptedAnswerId field to given value.

### HasAcceptedAnswerId

`func (o *Article) HasAcceptedAnswerId() bool`

HasAcceptedAnswerId returns a boolean if a field has been set.

### GetAuthorUser

`func (o *Article) GetAuthorUser() AccountProfile`

GetAuthorUser returns the AuthorUser field if non-nil, zero value otherwise.

### GetAuthorUserOk

`func (o *Article) GetAuthorUserOk() (*AccountProfile, bool)`

GetAuthorUserOk returns a tuple with the AuthorUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUser

`func (o *Article) SetAuthorUser(v AccountProfile)`

SetAuthorUser sets AuthorUser field to given value.

### HasAuthorUser

`func (o *Article) HasAuthorUser() bool`

HasAuthorUser returns a boolean if a field has been set.

### GetLastReplyUser

`func (o *Article) GetLastReplyUser() AccountProfile`

GetLastReplyUser returns the LastReplyUser field if non-nil, zero value otherwise.

### GetLastReplyUserOk

`func (o *Article) GetLastReplyUserOk() (*AccountProfile, bool)`

GetLastReplyUserOk returns a tuple with the LastReplyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyUser

`func (o *Article) SetLastReplyUser(v AccountProfile)`

SetLastReplyUser sets LastReplyUser field to given value.

### HasLastReplyUser

`func (o *Article) HasLastReplyUser() bool`

HasLastReplyUser returns a boolean if a field has been set.

### GetLastReplyAt

`func (o *Article) GetLastReplyAt() string`

GetLastReplyAt returns the LastReplyAt field if non-nil, zero value otherwise.

### GetLastReplyAtOk

`func (o *Article) GetLastReplyAtOk() (*string, bool)`

GetLastReplyAtOk returns a tuple with the LastReplyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyAt

`func (o *Article) SetLastReplyAt(v string)`

SetLastReplyAt sets LastReplyAt field to given value.

### HasLastReplyAt

`func (o *Article) HasLastReplyAt() bool`

HasLastReplyAt returns a boolean if a field has been set.

### GetCoverImageUrl

`func (o *Article) GetCoverImageUrl() string`

GetCoverImageUrl returns the CoverImageUrl field if non-nil, zero value otherwise.

### GetCoverImageUrlOk

`func (o *Article) GetCoverImageUrlOk() (*string, bool)`

GetCoverImageUrlOk returns a tuple with the CoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageUrl

`func (o *Article) SetCoverImageUrl(v string)`

SetCoverImageUrl sets CoverImageUrl field to given value.

### HasCoverImageUrl

`func (o *Article) HasCoverImageUrl() bool`

HasCoverImageUrl returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Article) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Article) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Article) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Article) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Article) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Article) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Article) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Article) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Article) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Article) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Article) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Article) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Article) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Article) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Article) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Article) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


