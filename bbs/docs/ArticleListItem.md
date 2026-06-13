# ArticleListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ContentRender** | Pointer to **string** |  | [optional] 
**HasPostscript** | Pointer to **bool** |  | [optional] 
**HasReward** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Statement** | Pointer to **string** |  | [optional] 
**Commentable** | Pointer to **bool** |  | [optional] 
**Anonymous** | Pointer to **bool** |  | [optional] 
**ViewCount** | Pointer to **int32** |  | [optional] 
**ThankCount** | Pointer to **int32** |  | [optional] 
**LikeCount** | Pointer to **int32** |  | [optional] 
**CollectCount** | Pointer to **int32** |  | [optional] 
**WatchCount** | Pointer to **int32** |  | [optional] 
**ReplyCount** | Pointer to **int32** |  | [optional] 
**BountyPoints** | Pointer to **int32** |  | [optional] 
**AcceptedAnswerId** | Pointer to **string** |  | [optional] 
**AuthorUser** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**LastReplyUser** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**LastReplyAt** | Pointer to **string** |  | [optional] 
**CoverImageUrl** | Pointer to **string** |  | [optional] 
**ViewerActionState** | Pointer to [**ArticleViewerActionState**](ArticleViewerActionState.md) |  | [optional] 
**PublishedAt** | Pointer to **string** |  | [optional] 
**PublishStatus** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Restriction** | Pointer to **string** |  | [optional] 
**EditedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewArticleListItem

`func NewArticleListItem() *ArticleListItem`

NewArticleListItem instantiates a new ArticleListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleListItemWithDefaults

`func NewArticleListItemWithDefaults() *ArticleListItem`

NewArticleListItemWithDefaults instantiates a new ArticleListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArticleListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticleListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticleListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArticleListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ArticleListItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticleListItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticleListItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ArticleListItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *ArticleListItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ArticleListItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ArticleListItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ArticleListItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentRender

`func (o *ArticleListItem) GetContentRender() string`

GetContentRender returns the ContentRender field if non-nil, zero value otherwise.

### GetContentRenderOk

`func (o *ArticleListItem) GetContentRenderOk() (*string, bool)`

GetContentRenderOk returns a tuple with the ContentRender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRender

`func (o *ArticleListItem) SetContentRender(v string)`

SetContentRender sets ContentRender field to given value.

### HasContentRender

`func (o *ArticleListItem) HasContentRender() bool`

HasContentRender returns a boolean if a field has been set.

### GetHasPostscript

`func (o *ArticleListItem) GetHasPostscript() bool`

GetHasPostscript returns the HasPostscript field if non-nil, zero value otherwise.

### GetHasPostscriptOk

`func (o *ArticleListItem) GetHasPostscriptOk() (*bool, bool)`

GetHasPostscriptOk returns a tuple with the HasPostscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPostscript

`func (o *ArticleListItem) SetHasPostscript(v bool)`

SetHasPostscript sets HasPostscript field to given value.

### HasHasPostscript

`func (o *ArticleListItem) HasHasPostscript() bool`

HasHasPostscript returns a boolean if a field has been set.

### GetHasReward

`func (o *ArticleListItem) GetHasReward() bool`

GetHasReward returns the HasReward field if non-nil, zero value otherwise.

### GetHasRewardOk

`func (o *ArticleListItem) GetHasRewardOk() (*bool, bool)`

GetHasRewardOk returns a tuple with the HasReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReward

`func (o *ArticleListItem) SetHasReward(v bool)`

SetHasReward sets HasReward field to given value.

### HasHasReward

`func (o *ArticleListItem) HasHasReward() bool`

HasHasReward returns a boolean if a field has been set.

### GetType

`func (o *ArticleListItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArticleListItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArticleListItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ArticleListItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatement

`func (o *ArticleListItem) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ArticleListItem) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ArticleListItem) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *ArticleListItem) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetCommentable

`func (o *ArticleListItem) GetCommentable() bool`

GetCommentable returns the Commentable field if non-nil, zero value otherwise.

### GetCommentableOk

`func (o *ArticleListItem) GetCommentableOk() (*bool, bool)`

GetCommentableOk returns a tuple with the Commentable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentable

`func (o *ArticleListItem) SetCommentable(v bool)`

SetCommentable sets Commentable field to given value.

### HasCommentable

`func (o *ArticleListItem) HasCommentable() bool`

HasCommentable returns a boolean if a field has been set.

### GetAnonymous

`func (o *ArticleListItem) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *ArticleListItem) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *ArticleListItem) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *ArticleListItem) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetViewCount

`func (o *ArticleListItem) GetViewCount() int32`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *ArticleListItem) GetViewCountOk() (*int32, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *ArticleListItem) SetViewCount(v int32)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *ArticleListItem) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetThankCount

`func (o *ArticleListItem) GetThankCount() int32`

GetThankCount returns the ThankCount field if non-nil, zero value otherwise.

### GetThankCountOk

`func (o *ArticleListItem) GetThankCountOk() (*int32, bool)`

GetThankCountOk returns a tuple with the ThankCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankCount

`func (o *ArticleListItem) SetThankCount(v int32)`

SetThankCount sets ThankCount field to given value.

### HasThankCount

`func (o *ArticleListItem) HasThankCount() bool`

HasThankCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *ArticleListItem) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *ArticleListItem) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *ArticleListItem) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *ArticleListItem) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetCollectCount

`func (o *ArticleListItem) GetCollectCount() int32`

GetCollectCount returns the CollectCount field if non-nil, zero value otherwise.

### GetCollectCountOk

`func (o *ArticleListItem) GetCollectCountOk() (*int32, bool)`

GetCollectCountOk returns a tuple with the CollectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectCount

`func (o *ArticleListItem) SetCollectCount(v int32)`

SetCollectCount sets CollectCount field to given value.

### HasCollectCount

`func (o *ArticleListItem) HasCollectCount() bool`

HasCollectCount returns a boolean if a field has been set.

### GetWatchCount

`func (o *ArticleListItem) GetWatchCount() int32`

GetWatchCount returns the WatchCount field if non-nil, zero value otherwise.

### GetWatchCountOk

`func (o *ArticleListItem) GetWatchCountOk() (*int32, bool)`

GetWatchCountOk returns a tuple with the WatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchCount

`func (o *ArticleListItem) SetWatchCount(v int32)`

SetWatchCount sets WatchCount field to given value.

### HasWatchCount

`func (o *ArticleListItem) HasWatchCount() bool`

HasWatchCount returns a boolean if a field has been set.

### GetReplyCount

`func (o *ArticleListItem) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *ArticleListItem) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *ArticleListItem) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *ArticleListItem) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetBountyPoints

`func (o *ArticleListItem) GetBountyPoints() int32`

GetBountyPoints returns the BountyPoints field if non-nil, zero value otherwise.

### GetBountyPointsOk

`func (o *ArticleListItem) GetBountyPointsOk() (*int32, bool)`

GetBountyPointsOk returns a tuple with the BountyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBountyPoints

`func (o *ArticleListItem) SetBountyPoints(v int32)`

SetBountyPoints sets BountyPoints field to given value.

### HasBountyPoints

`func (o *ArticleListItem) HasBountyPoints() bool`

HasBountyPoints returns a boolean if a field has been set.

### GetAcceptedAnswerId

`func (o *ArticleListItem) GetAcceptedAnswerId() string`

GetAcceptedAnswerId returns the AcceptedAnswerId field if non-nil, zero value otherwise.

### GetAcceptedAnswerIdOk

`func (o *ArticleListItem) GetAcceptedAnswerIdOk() (*string, bool)`

GetAcceptedAnswerIdOk returns a tuple with the AcceptedAnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAnswerId

`func (o *ArticleListItem) SetAcceptedAnswerId(v string)`

SetAcceptedAnswerId sets AcceptedAnswerId field to given value.

### HasAcceptedAnswerId

`func (o *ArticleListItem) HasAcceptedAnswerId() bool`

HasAcceptedAnswerId returns a boolean if a field has been set.

### GetAuthorUser

`func (o *ArticleListItem) GetAuthorUser() AccountProfile`

GetAuthorUser returns the AuthorUser field if non-nil, zero value otherwise.

### GetAuthorUserOk

`func (o *ArticleListItem) GetAuthorUserOk() (*AccountProfile, bool)`

GetAuthorUserOk returns a tuple with the AuthorUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUser

`func (o *ArticleListItem) SetAuthorUser(v AccountProfile)`

SetAuthorUser sets AuthorUser field to given value.

### HasAuthorUser

`func (o *ArticleListItem) HasAuthorUser() bool`

HasAuthorUser returns a boolean if a field has been set.

### GetLastReplyUser

`func (o *ArticleListItem) GetLastReplyUser() AccountProfile`

GetLastReplyUser returns the LastReplyUser field if non-nil, zero value otherwise.

### GetLastReplyUserOk

`func (o *ArticleListItem) GetLastReplyUserOk() (*AccountProfile, bool)`

GetLastReplyUserOk returns a tuple with the LastReplyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyUser

`func (o *ArticleListItem) SetLastReplyUser(v AccountProfile)`

SetLastReplyUser sets LastReplyUser field to given value.

### HasLastReplyUser

`func (o *ArticleListItem) HasLastReplyUser() bool`

HasLastReplyUser returns a boolean if a field has been set.

### GetLastReplyAt

`func (o *ArticleListItem) GetLastReplyAt() string`

GetLastReplyAt returns the LastReplyAt field if non-nil, zero value otherwise.

### GetLastReplyAtOk

`func (o *ArticleListItem) GetLastReplyAtOk() (*string, bool)`

GetLastReplyAtOk returns a tuple with the LastReplyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplyAt

`func (o *ArticleListItem) SetLastReplyAt(v string)`

SetLastReplyAt sets LastReplyAt field to given value.

### HasLastReplyAt

`func (o *ArticleListItem) HasLastReplyAt() bool`

HasLastReplyAt returns a boolean if a field has been set.

### GetCoverImageUrl

`func (o *ArticleListItem) GetCoverImageUrl() string`

GetCoverImageUrl returns the CoverImageUrl field if non-nil, zero value otherwise.

### GetCoverImageUrlOk

`func (o *ArticleListItem) GetCoverImageUrlOk() (*string, bool)`

GetCoverImageUrlOk returns a tuple with the CoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageUrl

`func (o *ArticleListItem) SetCoverImageUrl(v string)`

SetCoverImageUrl sets CoverImageUrl field to given value.

### HasCoverImageUrl

`func (o *ArticleListItem) HasCoverImageUrl() bool`

HasCoverImageUrl returns a boolean if a field has been set.

### GetViewerActionState

`func (o *ArticleListItem) GetViewerActionState() ArticleViewerActionState`

GetViewerActionState returns the ViewerActionState field if non-nil, zero value otherwise.

### GetViewerActionStateOk

`func (o *ArticleListItem) GetViewerActionStateOk() (*ArticleViewerActionState, bool)`

GetViewerActionStateOk returns a tuple with the ViewerActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerActionState

`func (o *ArticleListItem) SetViewerActionState(v ArticleViewerActionState)`

SetViewerActionState sets ViewerActionState field to given value.

### HasViewerActionState

`func (o *ArticleListItem) HasViewerActionState() bool`

HasViewerActionState returns a boolean if a field has been set.

### GetPublishedAt

`func (o *ArticleListItem) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *ArticleListItem) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *ArticleListItem) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *ArticleListItem) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetPublishStatus

`func (o *ArticleListItem) GetPublishStatus() string`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *ArticleListItem) GetPublishStatusOk() (*string, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *ArticleListItem) SetPublishStatus(v string)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *ArticleListItem) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *ArticleListItem) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ArticleListItem) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ArticleListItem) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ArticleListItem) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetRestriction

`func (o *ArticleListItem) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ArticleListItem) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ArticleListItem) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *ArticleListItem) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetEditedAt

`func (o *ArticleListItem) GetEditedAt() string`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *ArticleListItem) GetEditedAtOk() (*string, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *ArticleListItem) SetEditedAt(v string)`

SetEditedAt sets EditedAt field to given value.

### HasEditedAt

`func (o *ArticleListItem) HasEditedAt() bool`

HasEditedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArticleListItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArticleListItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArticleListItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArticleListItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ArticleListItem) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ArticleListItem) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ArticleListItem) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ArticleListItem) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArticleListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArticleListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArticleListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArticleListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ArticleListItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArticleListItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArticleListItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArticleListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


