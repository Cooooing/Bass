# ArticleDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ContentRender** | Pointer to **string** |  | [optional] 
**HasPostscript** | Pointer to **bool** |  | [optional] 
**HasReward** | Pointer to **bool** |  | [optional] 
**RewardContent** | Pointer to **string** |  | [optional] 
**RewardContentRender** | Pointer to **string** |  | [optional] 
**RewardPoints** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Statement** | Pointer to **string** |  | [optional] 
**Commentable** | Pointer to **bool** |  | [optional] 
**ViewCount** | Pointer to **int32** |  | [optional] 
**ThankCount** | Pointer to **int32** |  | [optional] 
**LikeCount** | Pointer to **int32** |  | [optional] 
**CollectCount** | Pointer to **int32** |  | [optional] 
**RewardCount** | Pointer to **int32** |  | [optional] 
**ReplyCount** | Pointer to **int32** |  | [optional] 
**AuthorUser** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**LastReplyUser** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**LastReplyAt** | Pointer to **time.Time** |  | [optional] 
**CoverImageUrl** | Pointer to **string** |  | [optional] 
**ViewerActionState** | Pointer to [**ArticleViewerActionState**](ArticleViewerActionState.md) |  | [optional] 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**Postscripts** | Pointer to [**[]ArticlePostscript**](ArticlePostscript.md) |  | [optional] 
**PublishStatus** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Restriction** | Pointer to **string** |  | [optional] 
**EditedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewArticleDetail

`func NewArticleDetail() *ArticleDetail`

NewArticleDetail instantiates a new ArticleDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleDetailWithDefaults

`func NewArticleDetailWithDefaults() *ArticleDetail`

NewArticleDetailWithDefaults instantiates a new ArticleDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArticleDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticleDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticleDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArticleDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ArticleDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticleDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticleDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ArticleDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *ArticleDetail) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ArticleDetail) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ArticleDetail) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ArticleDetail) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentRender

`func (o *ArticleDetail) GetContentRender() string`

GetContentRender returns the ContentRender field if non-nil, zero value otherwise.

### GetContentRenderOk

`func (o *ArticleDetail) GetContentRenderOk() (*string, bool)`

GetContentRenderOk returns a tuple with the ContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRender

`func (o *ArticleDetail) SetContentRender(v string)`

SetContentRender sets ContentRender field to given value.

### HasContentRender

`func (o *ArticleDetail) HasContentRender() bool`

HasContentRender returns a boolean if a field has been set.

### GetHasPostscript

`func (o *ArticleDetail) GetHasPostscript() bool`

GetHasPostscript returns the HasPostscript field if non-nil, zero value otherwise.

### GetHasPostscriptOk

`func (o *ArticleDetail) GetHasPostscriptOk() (*bool, bool)`

GetHasPostscriptOk returns a tuple with the HasPostscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPostscript

`func (o *ArticleDetail) SetHasPostscript(v bool)`

SetHasPostscript sets HasPostscript field to given value.

### HasHasPostscript

`func (o *ArticleDetail) HasHasPostscript() bool`

HasHasPostscript returns a boolean if a field has been set.

### GetHasReward

`func (o *ArticleDetail) GetHasReward() bool`

GetHasReward returns the HasReward field if non-nil, zero value otherwise.

### GetHasRewardOk

`func (o *ArticleDetail) GetHasRewardOk() (*bool, bool)`

GetHasRewardOk returns a tuple with the HasReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReward

`func (o *ArticleDetail) SetHasReward(v bool)`

SetHasReward sets HasReward field to given value.

### HasHasReward

`func (o *ArticleDetail) HasHasReward() bool`

HasHasReward returns a boolean if a field has been set.

### GetRewardContent

`func (o *ArticleDetail) GetRewardContent() string`

GetRewardContent returns the RewardContent field if non-nil, zero value otherwise.

### GetRewardContentOk

`func (o *ArticleDetail) GetRewardContentOk() (*string, bool)`

GetRewardContentOk returns a tuple with the RewardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardContent

`func (o *ArticleDetail) SetRewardContent(v string)`

SetRewardContent sets RewardContent field to given value.

### HasRewardContent

`func (o *ArticleDetail) HasRewardContent() bool`

HasRewardContent returns a boolean if a field has been set.

### GetRewardContentRender

`func (o *ArticleDetail) GetRewardContentRender() string`

GetRewardContentRender returns the RewardContentRender field if non-nil, zero value otherwise.

### GetRewardContentRenderOk

`func (o *ArticleDetail) GetRewardContentRenderOk() (*string, bool)`

GetRewardContentRenderOk returns a tuple with the RewardContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardContentRender

`func (o *ArticleDetail) SetRewardContentRender(v string)`

SetRewardContentRender sets RewardContentRender field to given value.

### HasRewardContentRender

`func (o *ArticleDetail) HasRewardContentRender() bool`

HasRewardContentRender returns a boolean if a field has been set.

### GetRewardPoints

`func (o *ArticleDetail) GetRewardPoints() int32`

GetRewardPoints returns the RewardPoints field if non-nil, zero value otherwise.

### GetRewardPointsOk

`func (o *ArticleDetail) GetRewardPointsOk() (*int32, bool)`

GetRewardPointsOk returns a tuple with the RewardPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPoints

`func (o *ArticleDetail) SetRewardPoints(v int32)`

SetRewardPoints sets RewardPoints field to given value.

### HasRewardPoints

`func (o *ArticleDetail) HasRewardPoints() bool`

HasRewardPoints returns a boolean if a field has been set.

### GetType

`func (o *ArticleDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArticleDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArticleDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ArticleDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatement

`func (o *ArticleDetail) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ArticleDetail) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ArticleDetail) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *ArticleDetail) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetCommentable

`func (o *ArticleDetail) GetCommentable() bool`

GetCommentable returns the Commentable field if non-nil, zero value otherwise.

### GetCommentableOk

`func (o *ArticleDetail) GetCommentableOk() (*bool, bool)`

GetCommentableOk returns a tuple with the Commentable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentable

`func (o *ArticleDetail) SetCommentable(v bool)`

SetCommentable sets Commentable field to given value.

### HasCommentable

`func (o *ArticleDetail) HasCommentable() bool`

HasCommentable returns a boolean if a field has been set.

### GetViewCount

`func (o *ArticleDetail) GetViewCount() int32`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *ArticleDetail) GetViewCountOk() (*int32, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *ArticleDetail) SetViewCount(v int32)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *ArticleDetail) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetThankCount

`func (o *ArticleDetail) GetThankCount() int32`

GetThankCount returns the ThankCount field if non-nil, zero value otherwise.

### GetThankCountOk

`func (o *ArticleDetail) GetThankCountOk() (*int32, bool)`

GetThankCountOk returns a tuple with the ThankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankCount

`func (o *ArticleDetail) SetThankCount(v int32)`

SetThankCount sets ThankCount field to given value.

### HasThankCount

`func (o *ArticleDetail) HasThankCount() bool`

HasThankCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *ArticleDetail) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *ArticleDetail) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *ArticleDetail) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *ArticleDetail) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetCollectCount

`func (o *ArticleDetail) GetCollectCount() int32`

GetCollectCount returns the CollectCount field if non-nil, zero value otherwise.

### GetCollectCountOk

`func (o *ArticleDetail) GetCollectCountOk() (*int32, bool)`

GetCollectCountOk returns a tuple with the CollectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectCount

`func (o *ArticleDetail) SetCollectCount(v int32)`

SetCollectCount sets CollectCount field to given value.

### HasCollectCount

`func (o *ArticleDetail) HasCollectCount() bool`

HasCollectCount returns a boolean if a field has been set.

### GetRewardCount

`func (o *ArticleDetail) GetRewardCount() int32`

GetRewardCount returns the RewardCount field if non-nil, zero value otherwise.

### GetRewardCountOk

`func (o *ArticleDetail) GetRewardCountOk() (*int32, bool)`

GetRewardCountOk returns a tuple with the RewardCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardCount

`func (o *ArticleDetail) SetRewardCount(v int32)`

SetRewardCount sets RewardCount field to given value.

### HasRewardCount

`func (o *ArticleDetail) HasRewardCount() bool`

HasRewardCount returns a boolean if a field has been set.

### GetReplyCount

`func (o *ArticleDetail) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *ArticleDetail) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *ArticleDetail) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *ArticleDetail) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetAuthorUser

`func (o *ArticleDetail) GetAuthorUser() AccountProfile`

GetAuthorUser returns the AuthorUser field if non-nil, zero value otherwise.

### GetAuthorUserOk

`func (o *ArticleDetail) GetAuthorUserOk() (*AccountProfile, bool)`

GetAuthorUserOk returns a tuple with the AuthorUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUser

`func (o *ArticleDetail) SetAuthorUser(v AccountProfile)`

SetAuthorUser sets AuthorUser field to given value.

### HasAuthorUser

`func (o *ArticleDetail) HasAuthorUser() bool`

HasAuthorUser returns a boolean if a field has been set.

### GetLastReplyUser

`func (o *ArticleDetail) GetLastReplyUser() AccountProfile`

GetLastReplyUser returns the LastReplyUser field if non-nil, zero value otherwise.

### GetLastReplyUserOk

`func (o *ArticleDetail) GetLastReplyUserOk() (*AccountProfile, bool)`

GetLastReplyUserOk returns a tuple with the LastReplyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyUser

`func (o *ArticleDetail) SetLastReplyUser(v AccountProfile)`

SetLastReplyUser sets LastReplyUser field to given value.

### HasLastReplyUser

`func (o *ArticleDetail) HasLastReplyUser() bool`

HasLastReplyUser returns a boolean if a field has been set.

### GetLastReplyAt

`func (o *ArticleDetail) GetLastReplyAt() time.Time`

GetLastReplyAt returns the LastReplyAt field if non-nil, zero value otherwise.

### GetLastReplyAtOk

`func (o *ArticleDetail) GetLastReplyAtOk() (*time.Time, bool)`

GetLastReplyAtOk returns a tuple with the LastReplyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyAt

`func (o *ArticleDetail) SetLastReplyAt(v time.Time)`

SetLastReplyAt sets LastReplyAt field to given value.

### HasLastReplyAt

`func (o *ArticleDetail) HasLastReplyAt() bool`

HasLastReplyAt returns a boolean if a field has been set.

### GetCoverImageUrl

`func (o *ArticleDetail) GetCoverImageUrl() string`

GetCoverImageUrl returns the CoverImageUrl field if non-nil, zero value otherwise.

### GetCoverImageUrlOk

`func (o *ArticleDetail) GetCoverImageUrlOk() (*string, bool)`

GetCoverImageUrlOk returns a tuple with the CoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageUrl

`func (o *ArticleDetail) SetCoverImageUrl(v string)`

SetCoverImageUrl sets CoverImageUrl field to given value.

### HasCoverImageUrl

`func (o *ArticleDetail) HasCoverImageUrl() bool`

HasCoverImageUrl returns a boolean if a field has been set.

### GetViewerActionState

`func (o *ArticleDetail) GetViewerActionState() ArticleViewerActionState`

GetViewerActionState returns the ViewerActionState field if non-nil, zero value otherwise.

### GetViewerActionStateOk

`func (o *ArticleDetail) GetViewerActionStateOk() (*ArticleViewerActionState, bool)`

GetViewerActionStateOk returns a tuple with the ViewerActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerActionState

`func (o *ArticleDetail) SetViewerActionState(v ArticleViewerActionState)`

SetViewerActionState sets ViewerActionState field to given value.

### HasViewerActionState

`func (o *ArticleDetail) HasViewerActionState() bool`

HasViewerActionState returns a boolean if a field has been set.

### GetPublishedAt

`func (o *ArticleDetail) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *ArticleDetail) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *ArticleDetail) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *ArticleDetail) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetPostscripts

`func (o *ArticleDetail) GetPostscripts() []ArticlePostscript`

GetPostscripts returns the Postscripts field if non-nil, zero value otherwise.

### GetPostscriptsOk

`func (o *ArticleDetail) GetPostscriptsOk() (*[]ArticlePostscript, bool)`

GetPostscriptsOk returns a tuple with the Postscripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostscripts

`func (o *ArticleDetail) SetPostscripts(v []ArticlePostscript)`

SetPostscripts sets Postscripts field to given value.

### HasPostscripts

`func (o *ArticleDetail) HasPostscripts() bool`

HasPostscripts returns a boolean if a field has been set.

### GetPublishStatus

`func (o *ArticleDetail) GetPublishStatus() string`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *ArticleDetail) GetPublishStatusOk() (*string, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *ArticleDetail) SetPublishStatus(v string)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *ArticleDetail) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *ArticleDetail) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ArticleDetail) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ArticleDetail) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ArticleDetail) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetRestriction

`func (o *ArticleDetail) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ArticleDetail) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ArticleDetail) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *ArticleDetail) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetEditedAt

`func (o *ArticleDetail) GetEditedAt() time.Time`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *ArticleDetail) GetEditedAtOk() (*time.Time, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *ArticleDetail) SetEditedAt(v time.Time)`

SetEditedAt sets EditedAt field to given value.

### HasEditedAt

`func (o *ArticleDetail) HasEditedAt() bool`

HasEditedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArticleDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArticleDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArticleDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArticleDetail) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ArticleDetail) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ArticleDetail) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ArticleDetail) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ArticleDetail) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArticleDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArticleDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArticleDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArticleDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArticleDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArticleDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArticleDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArticleDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


